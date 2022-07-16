package vmmanagerapi

import (
	"encoding/json"
	"fmt"

	"github.com/gadost/go-vmmanager-api/types"
)

func (a *Api) Accounts() (*types.Accounts, error) {
	bodyResp, err := a.NewRequest(NilPayload, "/account", requestTypeGet, DefaultService)

	var r *types.Accounts
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) User(uid int) ([]byte, error) {
	uri := fmt.Sprintf("/user/%d", uid)
	bodyResp, err := a.NewRequest(NilPayload, uri, requestTypeGet, DefaultService)
	return bodyResp, err
}

func (a *Api) AccountNew(params *types.AccountNew) (*types.ID, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(payload, "/account", requestTypePost, DefaultService)

	var r *types.ID
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) AccountSync(params *types.AccountSyncingParams) (*types.ID, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(payload, "/account/sync", requestTypePost, DefaultService)

	var r *types.ID
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) InviteUser(params *types.InviteUserParams) (*types.ID, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(payload, "/invite", requestTypePost, DefaultService)

	var r *types.ID
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) ReInviteUser(params *types.InviteUserParams) (*types.ID, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(payload, "/reinvite", requestTypePost, DefaultService)

	var r *types.ID
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) UserUpdate(uid int, params *types.UserParams) (*types.ID, error) {
	payload, _ := json.Marshal(params)
	uri := fmt.Sprintf("/user/%d", uid)
	bodyResp, err := a.NewRequest(payload, uri, requestTypePost, DefaultService)
	var r *types.ID
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) UserResume(uid int) (*types.ID, error) {
	uri := fmt.Sprintf("/user/%d/%s", uid, "resume")
	bodyResp, err := a.NewRequest(NilPayload, uri, requestTypePost, DefaultService)
	var r *types.ID
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) UserSuspend(uid int) (*types.ID, error) {
	uri := fmt.Sprintf("/user/%d/%s", uid, "suspend")
	bodyResp, err := a.NewRequest(NilPayload, uri, requestTypePost, DefaultService)
	var r *types.ID
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) UserKeyUpdate(uid int) (*types.ID, error) {
	uri := fmt.Sprintf("/user/%d/key", uid)
	bodyResp, err := a.NewRequest(NilPayload, uri, requestTypePost, DefaultService)
	var r *types.ID
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) UserPasswordUpdate(uid int, params *types.UserParams) (*types.ID, error) {
	payload, _ := json.Marshal(params)
	uri := fmt.Sprintf("/user/%d/password", uid)
	bodyResp, err := a.NewRequest(payload, uri, requestTypePost, DefaultService)
	var r *types.ID
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) UserDelete(uid int) (*types.ID, error) {
	uri := fmt.Sprintf("/user/%d", uid)
	bodyResp, err := a.NewRequest(NilPayload, uri, requestTypeDelete, DefaultService)
	var r *types.ID
	json.Unmarshal(bodyResp, &r)
	return r, err
}
