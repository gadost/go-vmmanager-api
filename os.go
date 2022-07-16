package vmmanagerapi

import (
	"encoding/json"
	"fmt"

	"github.com/gadost/go-vmmanager-api/types"
)

// OSes get os list
func (a *Api) OSes() (*types.OSList, error) {
	bodyResp, err := a.NewRequest(
		NilPayload,
		"/os",
		requestTypeGet,
		DefaultService)

	var t *types.OSList
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// OSUpdate update os params
func (a *Api) OSUpdate(oid int, params *types.OSUpdateParams) (*types.Tasks, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/os/%d", oid),
		requestTypePost,
		DefaultService)
	var t *types.Tasks
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// OSDelete delete os
func (a *Api) OSDelete(oid int) (*types.DeletedResponse, error) {
	bodyResp, err := a.NewRequest(
		NilPayload,
		fmt.Sprintf("/os/%d", oid),
		requestTypeDelete,
		DefaultService)

	var t *types.DeletedResponse
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// OSNodeDelete delete os on node
func (a *Api) OSNodeDelete(oid int, nid int) (*types.DeletedResponse, error) {
	bodyResp, err := a.NewRequest(
		NilPayload,
		fmt.Sprintf("/os/%d/node/%d", oid, nid),
		requestTypeDelete,
		DefaultService)

	var t *types.DeletedResponse
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// OSNodeDisable disable os on node
func (a *Api) OSNodeDisable(oid int, nid int) ([]byte, error) {
	bodyResp, err := a.NewRequest(
		NilPayload,
		fmt.Sprintf("/os/%d/node/%d/disable", oid, nid),
		requestTypePost,
		DefaultService)

	return bodyResp, err
}

// OSNodeEnable enable os on node
func (a *Api) OSNodeEnable(oid int, nid int) ([]byte, error) {
	bodyResp, err := a.NewRequest(
		NilPayload,
		fmt.Sprintf("/os/%d/node/%d/enable", oid, nid),
		requestTypePost,
		DefaultService)

	return bodyResp, err
}

// OSesAll get all available os
func (a *Api) OSesAll() (*types.OSList, error) {
	bodyResp, err := a.NewRequest(
		NilPayload,
		"/os/all",
		requestTypeGet,
		DefaultService)

	var t *types.OSList
	json.Unmarshal(bodyResp, &t)
	return t, err
}
