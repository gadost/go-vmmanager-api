package vmmanagerapi

import (
	"crypto/tls"
	"fmt"
	"net/http"
)

const (
	ApiVersion                = "v3"
	AuthApiV4                 = "v4"
	DefaultService            = "vm"
	AuthService               = "auth"
	AuthByEmailAndPasswordUri = "/public/auth"
	AuthByKeyUri              = "/public/auth_by_key"
	AuthV4Uri                 = "/public/token"
	requestTypePost           = "POST"
	requestTypeGet            = "GET"
	requestTypeDelete         = "DELETE"
	NilPayload                = ""
	BackupUri                 = "/backup"
)

type Api struct {
	Host     string
	conn     *http.Client
	AuthData AuthData
}

func New(host string) *Api {
	return &Api{
		Host: host,
		conn: connect(),
	}
}

func (a *Api) entrypoint(service string) string {
	switch service {
	case AuthService:
		return fmt.Sprintf("https://%s/%s/%s", a.Host, AuthService, AuthApiV4)
	default:
		return fmt.Sprintf("https://%s/%s/%s", a.Host, DefaultService, ApiVersion)
	}

}

func connect() *http.Client {
	tl := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	return &http.Client{Transport: tl}
}
