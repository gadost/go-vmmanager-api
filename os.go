package vmmanagerapi

import (
	"encoding/json"
	"fmt"

	"github.com/gadost/go-vmmanager-api/types"
)

func (a *Api) OSList() (*types.OSList, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		"/os",
		requestTypeGet,
		DefaultService)

	var t *types.OSList
	json.Unmarshal(bodyResp, &t)
	return t, err
}

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

func (a *Api) OSDelete(oid int) (*types.DeletedResponse, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/os/%d", oid),
		requestTypeDelete,
		DefaultService)

	var t *types.DeletedResponse
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) NodeOSDelete(oid int, nid int) (*types.DeletedResponse, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/os/%d/node/%d", oid, nid),
		requestTypeDelete,
		DefaultService)

	var t *types.DeletedResponse
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) NodeOSDisable(oid int, nid int) ([]byte, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/os/%d/node/%d/disable", oid, nid),
		requestTypePost,
		DefaultService)

	return bodyResp, err
}
func (a *Api) NodeOSEnable(oid int, nid int) ([]byte, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/os/%d/node/%d/enable", oid, nid),
		requestTypePost,
		DefaultService)

	return bodyResp, err
}

func (a *Api) OSListAll() (*types.OSList, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		"/os/all",
		requestTypeGet,
		DefaultService)

	var t *types.OSList
	json.Unmarshal(bodyResp, &t)
	return t, err
}