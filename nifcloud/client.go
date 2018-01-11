package nifcloud

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
)

// Client interface
type Client interface {
	CallComputingAPI(params map[string]string) (res string, err error)
}

type client struct {
	region     Region
	credential Credential
	debug      bool
}

// NewClient returns Client interface
func NewClient(regionName, accessKey, secretKey string, debug bool) Client {
	return &client{
		region:     NewRegion(regionName),
		credential: NewCredential(accessKey, secretKey),
		debug:      debug,
	}
}

func (c *client) CallComputingAPI(params map[string]string) (res string, err error) {
	endpoint, _ := url.Parse(c.region.ComputingEndpoint)

	req, err := http.NewRequest("GET", endpoint.String(), nil)

	q := req.URL.Query()
	for k, v := range params {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()

	if err != nil {
		return
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	c.credential.Sign2(req)

	result, err := http.DefaultClient.Do(req)

	if c.debug {
		dumpReq, _ := httputil.DumpRequest(req, true)
		dumpRes, _ := httputil.DumpResponse(result, true)

		println("########## Request ##########")
		fmt.Printf("%s", dumpReq)

		println("########## Response ##########")
		fmt.Printf("%s", dumpRes)
	}

	if err != nil {
		return
	}

	defer result.Body.Close()

	bodyBytes, err := ioutil.ReadAll(result.Body)
	res = string(bodyBytes)
	return
}
