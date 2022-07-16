package vmmanagerapi

import (
	"encoding/json"
	"fmt"

	"github.com/gadost/go-vmmanager-api/types"
)

// VMAPI returns API params
func (a *Api) VMAPI() ([]byte, error) {
	bodyResp, err := a.NewRequest(NilPayload, "/api", requestTypeGet, DefaultService)
	return bodyResp, err
}

// Changelog returns changelog
func (a *Api) Changelog() ([]byte, error) {
	bodyResp, err := a.NewRequest(NilPayload, "/changelog", requestTypeGet, DefaultService)
	return bodyResp, err
}

// ChangelogLang return changelog translated EN/RU/etc
func (a *Api) ChangelogLang(lang string) ([]byte, error) {
	uri := fmt.Sprintf("/changelog/%s", lang)
	bodyResp, err := a.NewRequest(NilPayload, uri, requestTypeGet, DefaultService)
	return bodyResp, err
}

// LicenseInfo returns license info
func (a *Api) LicenseInfo() (*types.LicenseInfo, error) {
	bodyResp, err := a.NewRequest(NilPayload, "/get_license", requestTypeGet, DefaultService)

	var r *types.LicenseInfo
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// InstanceInfo returns instance info
func (a *Api) InstanceInfo() ([]byte, error) {
	bodyResp, err := a.NewRequest(NilPayload, "/instance_info", requestTypeGet, DefaultService)
	return bodyResp, err
}

// LastVersion returns last version
func (a *Api) LastVersion() ([]byte, error) {
	bodyResp, err := a.NewRequest(NilPayload, "/last_version", requestTypeGet, DefaultService)
	return bodyResp, err
}

// PrivateKey return private key
func (a *Api) PrivateKey() ([]byte, error) {
	bodyResp, err := a.NewRequest(NilPayload, "/private_key", requestTypeGet, DefaultService)
	return bodyResp, err
}

// Version returns current version
func (a *Api) Version() ([]byte, error) {
	bodyResp, err := a.NewRequest(NilPayload, "/version", requestTypeGet, DefaultService)
	return bodyResp, err
}

// LicenseSet set license key
func (a *Api) LicenseSet(params *types.LicenseSet) ([]byte, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(payload, "/set_license", requestTypePost, DefaultService)
	return bodyResp, err
}

// LicenseUpdate update license key
func (a *Api) LicenseUpdate(params *types.LicenseSet) ([]byte, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(payload, "/update_license", requestTypePost, DefaultService)
	return bodyResp, err
}

// VMUpdate update vmmanager
func (a *Api) VMUpdate() ([]byte, error) {
	bodyResp, err := a.NewRequest(NilPayload, "/vmupdate", requestTypePost, DefaultService)
	return bodyResp, err
}

// WriterStatus return writer status
func (a *Api) WriterStatus() ([]byte, error) {
	bodyResp, err := a.NewRequest(NilPayload, "/writer_status", requestTypePost, DefaultService)
	return bodyResp, err
}
