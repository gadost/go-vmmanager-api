package vmmanagerapi

import (
	"encoding/json"
	"fmt"

	"github.com/gadost/go-vmmanager-api/types"
)

// Presets returns preset list
func (a *Api) Presets() (*types.PresetListResponse, error) {
	bodyResp, err := a.NewRequest(
		NilPayload,
		"/preset",
		requestTypeGet,
		DefaultService)

	var t *types.PresetListResponse
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// Preset returns preset params
func (a *Api) Preset(pid int) (*types.PresetListElement, error) {
	bodyResp, err := a.NewRequest(
		NilPayload,
		fmt.Sprintf("/preset/%d", pid),
		requestTypeGet,
		DefaultService)

	var t *types.PresetListElement
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// PresetNew create new preset
func (a *Api) PresetNew(params *types.PresetListElement) (*types.ID, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		"/preset",
		requestTypePost,
		DefaultService)

	var t *types.ID
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// PresetUpdate update preset params
func (a *Api) PresetUpdate(pid int,
	params *types.PresetListElement) (*types.ID, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/preset/%d", pid),
		requestTypePost,
		DefaultService)

	var t *types.ID
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// PresetDelete delete preset
func (a *Api) PresetDelete(pid int) (*types.ID, error) {
	bodyResp, err := a.NewRequest(
		NilPayload,
		fmt.Sprintf("/preset/%d", pid),
		requestTypeDelete,
		DefaultService)

	var t *types.ID
	json.Unmarshal(bodyResp, &t)
	return t, err
}
