package vmmanagerapi

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
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

type ParamsQuery struct {
	Query string
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

// NewRequest to API
func (a *Api) NewRequest(payload []byte, uri string, reqType string, service string) ([]byte, error) {
	body := bytes.NewReader(payload)
	req, err := http.NewRequest(reqType, a.entrypoint(service)+uri, body)
	if err != nil {
		return nil, err
	}

	req.Proto = "HTTP/2"
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	if a.AuthData.Token != "" {
		req.Header.Set("X-XSRF-TOKEN", a.AuthData.Token)
	}

	resp, err := a.conn.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	bodyResp, err := io.ReadAll(resp.Body)

	return bodyResp, err
}

func MakeQuery(params map[string]string) *ParamsQuery {
	query := "?"
	for k, v := range params {
		query += fmt.Sprintf("%s=%s&", k, v)
	}

	if query[len(query)-1:] == "&" {
		query = query[:len(query)-1]
	}

	return &ParamsQuery{Query: query}
}

func NilQuery() *ParamsQuery {
	return &ParamsQuery{}
}
