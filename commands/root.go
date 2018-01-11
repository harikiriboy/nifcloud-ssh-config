package commands

import (
	"bytes"
	"encoding/xml"
	"os"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/harikiriboy/nifcloud-ssh-config/nifcloud"
)

var sshConfigTemplate = template.Must(template.New("sshconfig").Parse(`
{{- range .Instances}}
Host {{.InstanceID}}
  HostName {{.IPAddress}}
  User {{$.User}}
  Port {{$.Port}}
  IdentityFile {{.KeyName}}
  IdentitiesOnly {{$.IdentitiesOnly}}
  StrictHostKeyChecking {{$.StrictHostKeyChecking }}
{{end}}
`))

type DescribeInstancesResponse struct {
	RequestID      string `xml:"requestId"`
	ReservationSet struct {
		Item []ReservationSetItem `xml:"item"`
	} `xml:"reservationSet"`
}

type ReservationSetItem struct {
	InstancesSet struct {
		Item struct {
			InstanceID       string `xml:"instanceId"`
			IPAddress        string `xml:"ipAddress"`
			PrivateIpAddress string `xml:"privateIpAddress"`
			KeyName          string `xml:"keyName"`
			Platform         string `xml:"platform"`
		} `xml:"item"`
	} `xml:"instancesSet"`
}

var RootCmd = &cobra.Command{
	Short: "A very simple tool that generates SSH config file using NIFCLOUD Comoutigng API.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.SetOutput(os.Stdout)
		var sshConfigs = ""

		regionNames := viper.GetStringSlice("region")

		if len(regionNames) == 0 {
			for name, _ := range nifcloud.Regions {
				regionNames = append(regionNames, name)
			}
		}

		for _, regionName := range regionNames {
			res, err := describeInstances(regionName)

			if err != nil {
				cmd.Printf("Failed to describeInstance: %s", err.Error())
				os.Exit(1)
			}

			sshConfig, err := generateSSHConfig(res)

			if err != nil {
				cmd.Printf("Failed to generateSSHConfig: %s", err.Error())
				os.Exit(1)
			}

			sshConfigs = sshConfigs + sshConfig
		}

		cmd.Print(sshConfigs)
	},
}

func init() {
	flags := RootCmd.Flags()
	flags.StringP("access-key", "", "", "NIFCLOUD API ACCESS KEY (default NIFCLOUD_ACCESS_KEY_ID environment variable")
	flags.StringP("secret-key", "", "", "NIFCLOUD API SECRET KEY (default NIFCLOUD_SECRET_ACCESS_KEY environment variable")
	flags.StringP("keydir", "", "~/.ssh", "Location of private keys")
	flags.StringP("user", "", "root", "SSH username")
	flags.StringP("prefix", "", "", "Specify a prefix to prepend to all host names")
	flags.StringP("ssh-key-name", "", "", "SSH key name (default use keyName get from API response)")
	flags.StringP("port", "", "22", "SSH port")
	flags.BoolP("strict-hostkey-checking", "", false, "Do not include StrictHostKeyChecking=no in ssh config")
	flags.BoolP("no-identities-only", "", false, "Do not include IdentitiesOnly=yes in ssh config; may cause connection refused if using ssh-agent")
	flags.BoolP("private", "", false, "Use private IP addresses (public are used by default)")
	flags.BoolP("debug", "", false, "Use debug mode")
	flags.StringSliceP("region", "", nil, "List of NIFCLOUD Regions (default all region)")
	flags.StringSliceP("exclude-instance-id", "", nil, "List of Exclude instanceID")

	viper.BindPFlag("access-key", flags.Lookup("access-key"))
	viper.BindPFlag("secret-key", flags.Lookup("secret-key"))
	viper.BindPFlag("keydir", flags.Lookup("keydir"))
	viper.BindPFlag("user", flags.Lookup("user"))
	viper.BindPFlag("prefix", flags.Lookup("prefix"))
	viper.BindPFlag("port", flags.Lookup("port"))
	viper.BindPFlag("ssh-key-name", flags.Lookup("ssh-key-name"))
	viper.BindPFlag("strict-hostkey-checking", flags.Lookup("strict-hostkey-checking"))
	viper.BindPFlag("no-identities-only", flags.Lookup("no-identities-only"))
	viper.BindPFlag("private", flags.Lookup("private"))
	viper.BindPFlag("region", flags.Lookup("region"))
	viper.BindPFlag("debug", flags.Lookup("debug"))
	viper.BindPFlag("exclude-instance-id", flags.Lookup("exclude-instance-id"))
}

func describeInstances(regionName string) (res string, err error) {
	client := nifcloud.NewClient(regionName, viper.GetString("access-key"), viper.GetString("secret-key"), viper.GetBool("debug"))
	params := map[string]string{"Action": "DescribeInstances"}
	res, err = client.CallComputingAPI(params)
	return
}

func generateSSHConfig(describeInstancesResponse string) (sshConfig string, err error) {
	obj := new(DescribeInstancesResponse)
	err = xml.NewDecoder(strings.NewReader(describeInstancesResponse)).Decode(&obj)
	if err != nil {
		return
	}

	type instance = struct {
		InstanceID string
		IPAddress  string
		KeyName    string
	}

	instances := []instance{}

	for _, item := range obj.ReservationSet.Item {
		if item.InstancesSet.Item.Platform == "windows" {
			continue
		}

		exclude := false
		for _, excludeInstanceID := range viper.GetStringSlice("exclude-instance-id") {
			if item.InstancesSet.Item.InstanceID == excludeInstanceID {
				exclude = true
				break
			}
		}

		if exclude {
			continue
		}

		var instanceID, ipaddress, keyName string

		instanceID = viper.GetString("prefix") + item.InstancesSet.Item.InstanceID

		if item.InstancesSet.Item.IPAddress == "" || viper.GetBool("private") {
			ipaddress = item.InstancesSet.Item.PrivateIpAddress
		} else {
			ipaddress = item.InstancesSet.Item.IPAddress
		}

		if item.InstancesSet.Item.KeyName == "" || viper.GetString("ssh-key-name") != "" {
			keyName = viper.GetString("keydir") + "/" + viper.GetString("ssh-key-name")
		} else {
			keyName = viper.GetString("keydir") + "/" + item.InstancesSet.Item.KeyName + "_private.pem"
		}

		instances = append(instances, instance{InstanceID: instanceID, IPAddress: ipaddress, KeyName: keyName})
	}

	var user, port, identitiesOnly, strictHostKeyChecking string

	user = viper.GetString("user")
	port = viper.GetString("port")

	if viper.GetBool("no-identities-only") {
		identitiesOnly = "no"
	} else {
		identitiesOnly = "yes"
	}

	if viper.GetBool("strict-hostkey-checking") {
		strictHostKeyChecking = "yes"
	} else {
		strictHostKeyChecking = "no"
	}

	buffer := new(bytes.Buffer)
	sshConfigTemplate.Execute(buffer, struct {
		Instances             []instance
		User                  string
		Port                  string
		IdentitiesOnly        string
		StrictHostKeyChecking string
	}{
		Instances:             instances,
		User:                  user,
		Port:                  port,
		IdentitiesOnly:        identitiesOnly,
		StrictHostKeyChecking: strictHostKeyChecking,
	})
	sshConfig = buffer.String()
	return
}
