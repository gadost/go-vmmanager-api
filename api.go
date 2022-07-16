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
	BackupUri                 = "/backup"
	DiskUri                   = "/disk"
	HostUri                   = "/host"
	ClusterUri                = "/cluster"
	SSHUri                    = "/ssh_key"
)

var (
	NilPayload = []byte("")
)

type Api struct {
	Host     string
	conn     *http.Client
	AuthData AuthData
}

type ParamsQuery struct {
	Query string
}

//New Api
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

// NewRequest make  requests to API.
// payload: marshaled json object, uri: endpoit uri, reqType: POST/GET/DELETE, service vm/auth
func (a *Api) NewRequest(payload []byte, uri string, reqType string, service string) ([]byte, error) {
	body := bytes.NewReader(payload)
	req, err := http.NewRequest(reqType, a.entrypoint(service)+uri, body)
	if err != nil {
		return nil, err
	}

	req = a.SetHeaders(req)

	resp, err := a.conn.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	bodyResp, err := io.ReadAll(resp.Body)

	return bodyResp, err
}

// SetHeaders set default headers for API requests
func (a *Api) SetHeaders(req *http.Request) *http.Request {
	req.Proto = "HTTP/2"
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	if a.AuthData.Token != "" {
		req.Header.Set("X-XSRF-TOKEN", a.AuthData.Token)
	}

	return req
}

// Make query returns ?x=y&z=t
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

//NilQuery return nil query
func NilQuery() *ParamsQuery {
	return &ParamsQuery{}
}
