package vmmanagerapi

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gadost/go-vmmanager-api/types"
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

// ChangePasswordByToken POST request to "/public/token/{token_id}/change_password"
func (a *Api) ChangePasswordByToken(password string) error {
	payload, _ := json.Marshal(&types.Password{Password: password})
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
	payload, _ := json.Marshal(&types.TokenLifetime{
		ExpiresAt:   expiresAt.Format("2006-01-02 15:04:05"),
		Description: desc,
	})

	bodyResp, _ := a.NewRequest(payload, "/token", requestTypePost, AuthService)
	json.Unmarshal(bodyResp, &a.AuthData)
	return a
}

// InviteUserByToken  POST request to /public/token/{token_id}/invite_user
func (a *Api) InviteUserByToken(password string, lang string) error {
	payload, _ := json.Marshal(&types.InviteUser{Password: password, Lang: lang})
	_, err := a.NewRequest(
		payload,
		fmt.Sprintf("/public/token/%s/invite_user", a.AuthData.Token),
		requestTypePost, DefaultService)
	if err != nil {
		return err
	}

	return nil
}
