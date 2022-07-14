package vmmanagerapi

import (
	"encoding/json"
	"fmt"

	"github.com/gadost/go-vmmanager-api/types"
)

func (a *Api) Presets() (*types.PresetListResponse, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		"/preset",
		requestTypeGet,
		DefaultService)

	var t *types.PresetListResponse
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) Preset(pid int) (*types.PresetListElement, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/preset/%d", pid),
		requestTypeGet,
		DefaultService)

	var t *types.PresetListElement
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) PresetNew(
	params *types.PresetListElement) (*types.ID, error) {
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

func (a *Api) PresetDelete(pid int) (*types.ID, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/preset/%d", pid),
		requestTypeDelete,
		DefaultService)

	var t *types.ID
	json.Unmarshal(bodyResp, &t)
	return t, err
}
