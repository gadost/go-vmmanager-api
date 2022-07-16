package vmmanagerapi

import (
	"encoding/json"
	"fmt"

	"github.com/gadost/go-vmmanager-api/types"
)

// ImportHistory import history
func (a *Api) ImportHistory() (*types.ImportHistoryResponse, error) {
	bodyResp, err := a.NewRequest(
		NilPayload,
		"/import_history",
		requestTypeGet,
		DefaultService)

	var t *types.ImportHistoryResponse
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// ImportHistoryNew create new history entry
func (a *Api) ImportHistoryNew(params *types.ImportHistoryRequest) (*types.ID, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		"/import_history",
		requestTypePost,
		DefaultService)

	var t *types.ID
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// ImportHistoryUpdate update history entry
func (a *Api) ImportHistoryUpdate(hsid int, params *types.State) (*types.ID, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/import_history/%d", hsid),
		requestTypePost,
		DefaultService)

	var t *types.ID
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// ImportHistoryResult returns history results
func (a *Api) ImportHistoryResult(hsid int, params *types.ImportHistoryResult) (*types.ID, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/import_history/%d/result", hsid),
		requestTypePost,
		DefaultService)

	var t *types.ID
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// ImportCluster import cluster
func (a *Api) ImportCluster(hsid int, params *types.ImportClusterParams) (*types.ID, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		"/import/cluster",
		requestTypePost,
		DefaultService)

	var t *types.ID
	json.Unmarshal(bodyResp, &t)
	return t, err
}
