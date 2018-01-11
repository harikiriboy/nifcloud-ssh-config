package nifcloud

import (
	"net/http"
	"os"

	"github.com/fuku2014/go-niftycloud-auth"
)

// Credential interface
type Credential interface {
	Sign2(req *http.Request)
}

type credential struct {
	credentials awsauth.Credentials
}

func NewCredential(accesskey, secretkey string) Credential {
	if accesskey == "" {
		accesskey = os.Getenv("NIFCLOUD_ACCESS_KEY_ID")
	}

	if accesskey == "" {
		accesskey = os.Getenv("NIFTY_ACCESS_KEY_ID")
	}

	if accesskey == "" {
		accesskey = os.Getenv("AWS_ACCESS_KEY_ID")
	}

	if secretkey == "" {
		secretkey = os.Getenv("NIFCLOUD_SECRET_ACCESS_KEY")
	}

	if secretkey == "" {
		secretkey = os.Getenv("NIFTY_SECRET_KEY")
	}

	if secretkey == "" {
		secretkey = os.Getenv("AWS_SECRET_ACCESS_KEY")
	}

	return credential{
		credentials: awsauth.Credentials{AccessKeyID: accesskey, SecretAccessKey: secretkey},
	}
}

func (c credential) Sign2(req *http.Request) {
	awsauth.Sign2(req, c.credentials)
}
