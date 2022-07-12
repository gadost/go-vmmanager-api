package vmmanagerapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"log"
)

type AuthByEmailAndPassword struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthByKey struct {
	Key string `json:"key"`
}

type AuthData struct {
	Confirmed bool   `json:"confirmed"`
	ID        string `json:"id"`
	Token     string `json:"token"`
}

type Auth struct {
	AuthByKey
	AuthByEmailAndPassword
	authUri string
}

type Password struct {
	Password string `json:"password"`
}

type InviteUser struct {
	Password        string `json:"password"`
	Lang            string `json:"lang"`
	AdditionalProp1 struct {
	} `json:"additionalProp1"`
}

type TokenLifetime struct {
	ExpiresAt   string `json:"expires_at"`
	Description string `json:"description"`
}

// Auth get session and token
func (a *Api) Auth(auth *Auth) *Api {
	var payload []byte

	if auth.AuthByEmailAndPassword.Email != "" && auth.AuthByEmailAndPassword.Password != "" {
		auth.authUri = AuthV4Uri
		payload, _ = json.Marshal(auth.AuthByEmailAndPassword)
	} else if auth.AuthByKey.Key != "" {
		auth.authUri = AuthByKeyUri
		payload, _ = json.Marshal(auth.AuthByKey)
	} else {
		return nil
	}

	bodyResp, err := a.NewRequest(payload, auth.authUri, requestTypePost, AuthService)
	if err != nil {
		return nil
	}

	json.Unmarshal(bodyResp, &a.AuthData)

	return a
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

// ChangePasswordByToken POST request to "/public/token/{token_id}/change_password"
func (a *Api) ChangePasswordByToken(password string) error {
	payload, _ := json.Marshal(&Password{Password: password})
	_, err := a.NewRequest(
		payload,
		fmt.Sprintf("/public/token/%s/change_password", a.AuthData.Token),
		requestTypePost, DefaultService)
	if err != nil {
		return err
	}

	return nil
}

// WithExtendedTokenLifetime extend token lifetime. default 1h
func (a *Api) WithExtendedTokenLifetime(expiresAt time.Time, desc string) *Api {
	payload, _ := json.Marshal(&TokenLifetime{
		ExpiresAt:   expiresAt.Format("2006-01-02 15:04:05"),
		Description: desc,
	})
	log.Println(string(payload))
	bodyResp, _ := a.NewRequest(payload, "/token", requestTypePost, AuthService)
	json.Unmarshal(bodyResp, &a.AuthData)
	return a
}

// InviteUserByToken  POST request to /public/token/{token_id}/invite_user
func (a *Api) InviteUserByToken(password string, lang string) error {
	payload, _ := json.Marshal(&InviteUser{Password: password, Lang: lang})
	_, err := a.NewRequest(
		payload,
		fmt.Sprintf("/public/token/%s/invite_user", a.AuthData.Token),
		requestTypePost, DefaultService)
	if err != nil {
		return err
	}

	return nil
}
