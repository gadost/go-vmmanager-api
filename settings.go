package vmmanagerapi

import (
	"encoding/json"
	"fmt"

	"github.com/gadost/go-vmmanager-api/types"
)

// Settings returns settings
func (a *Api) Settings() ([]byte, error) {
	bodyResp, err := a.NewRequest(NilPayload, "/settings", requestTypeGet, DefaultService)
	return bodyResp, err
}

// Setting returns setting params
func (a *Api) Setting(name string) ([]byte, error) {
	uri := fmt.Sprintf("/settings/%s", name)
	bodyResp, err := a.NewRequest(NilPayload, uri, requestTypeGet, DefaultService)
	return bodyResp, err
}

// SettingsDomain return domain
func (a *Api) SettingsDomain() ([]byte, error) {
	bodyResp, err := a.NewRequest(NilPayload, "/settings/domain", requestTypeGet, DefaultService)
	return bodyResp, err
}

// SettingsSendVMMessages return messages IDs
func (a *Api) SettingsSendVMMessages() ([]byte, error) {
	bodyResp, err := a.NewRequest(NilPayload, "/settings/send_vm_messages", requestTypeGet, DefaultService)
	return bodyResp, err
}

// SettingsSSHKey return ssh key settings
func (a *Api) SettingsSSHKey() ([]byte, error) {
	bodyResp, err := a.NewRequest(NilPayload, "/settings/ssh_key", requestTypeGet, DefaultService)
	return bodyResp, err
}

// SettingTaskTimeout return task timeout
func (a *Api) SettingTaskTimeout(tid int) ([]byte, error) {
	uri := fmt.Sprintf("/settings/task/%d/timeout", tid)
	bodyResp, err := a.NewRequest(NilPayload, uri, requestTypeGet, DefaultService)
	return bodyResp, err
}

// UsersLimits returns user limits
func (a *Api) UsersLimits() ([]byte, error) {
	bodyResp, err := a.NewRequest(NilPayload, "/user_limits", requestTypeGet, DefaultService)
	return bodyResp, err
}

// UserLimit return user limit params
func (a *Api) UserLimit(user string) (*types.UserLimits, error) {
	uri := fmt.Sprintf("/user_limits/account/%s", user)
	bodyResp, err := a.NewRequest(NilPayload, uri, requestTypeGet, DefaultService)

	var r *types.UserLimits
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// UserRoleLimit update role limits
func (a *Api) UserRoleLimit(rid int) (*types.UserLimits, error) {
	uri := fmt.Sprintf("/user_limits/role/%d", rid)
	bodyResp, err := a.NewRequest(NilPayload, uri, requestTypeGet, DefaultService)

	var r *types.UserLimits
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// SettingNew add new setting
func (a *Api) SettingNew(params *types.KeyValue) (*types.ID, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(payload, "/settings", requestTypePost, DefaultService)

	var r *types.ID
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// SettingUpdate update setting
func (a *Api) SettingUpdate(setting string, params *types.Value) (*types.ID, error) {
	payload, _ := json.Marshal(params)
	uri := fmt.Sprintf("/settings/%s", setting)
	bodyResp, err := a.NewRequest(payload, uri, requestTypePost, DefaultService)

	var r *types.ID
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// SettingsDomainUpdate update domain
func (a *Api) SettingsDomainUpdate(params *types.Name) ([]byte, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(payload, "/settings/domain", requestTypePost, DefaultService)

	return bodyResp, err
}

// SettingsEmailRestore restore email
func (a *Api) SettingsEmailRestore() ([]byte, error) {

	bodyResp, err := a.NewRequest(NilPayload, "/settings/email_restore", requestTypePost, DefaultService)

	return bodyResp, err
}

// SettingsTaskTimeoutSet set task timeout
func (a *Api) SettingsTaskTimeoutSet(task string, params *types.TimeoutSeconds) ([]byte, error) {
	uri := fmt.Sprintf("/settings/task/%s/timeout", task)
	bodyResp, err := a.NewRequest(NilPayload, uri, requestTypePost, DefaultService)

	return bodyResp, err
}

// UserLimitUpdate update user limits
func (a *Api) UserLimitUpdate(user string, params *types.UserLimits) ([]byte, error) {
	payload, _ := json.Marshal(params)
	uri := fmt.Sprintf("/user_limits/account/%s", user)
	bodyResp, err := a.NewRequest(payload, uri, requestTypePost, DefaultService)

	return bodyResp, err
}

// UserRoleLimitUpdate update user role limits
func (a *Api) UserRoleLimitUpdate(rid int, params *types.UserLimits) ([]byte, error) {
	payload, _ := json.Marshal(params)
	uri := fmt.Sprintf("/user_limits/role/%d", rid)
	bodyResp, err := a.NewRequest(payload, uri, requestTypePost, DefaultService)

	return bodyResp, err
}

// UserLimitDelete remove user limits
func (a *Api) UserLimitDelete(user string, params *types.UserLimits) ([]byte, error) {
	payload, _ := json.Marshal(params)
	uri := fmt.Sprintf("/user_limits/account/%s", user)
	bodyResp, err := a.NewRequest(payload, uri, requestTypeDelete, DefaultService)

	return bodyResp, err
}

// UserRoleLimitDelete delete user role limit
func (a *Api) UserRoleLimitDelete(rid int, params *types.UserLimits) ([]byte, error) {
	payload, _ := json.Marshal(params)
	uri := fmt.Sprintf("/user_limits/role/%d", rid)
	bodyResp, err := a.NewRequest(payload, uri, requestTypeDelete, DefaultService)

	return bodyResp, err
}

// SettingsTaskTimeoutDelete delete task timeout
func (a *Api) SettingsTaskTimeoutDelete(task string, params *types.TimeoutSeconds) ([]byte, error) {
	uri := fmt.Sprintf("/settings/task/%s/timeout", task)
	bodyResp, err := a.NewRequest(NilPayload, uri, requestTypeDelete, DefaultService)

	return bodyResp, err
}
