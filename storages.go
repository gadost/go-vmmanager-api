package vmmanagerapi

import (
	"encoding/json"
	"fmt"

	"github.com/gadost/go-vmmanager-api/types"
)

// Storages returns all storages
func (a *Api) Storages() (*types.Storages, error) {
	bodyResp, err := a.NewRequest(NilPayload, "/storage", requestTypeGet, DefaultService)

	var r *types.Storages
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// Storage returns storage params
func (a *Api) Storage(sid int) (*types.StorageParams, error) {
	uri := fmt.Sprintf("/storage/%d", sid)
	bodyResp, err := a.NewRequest(NilPayload, uri, requestTypeGet, DefaultService)

	var r *types.StorageParams
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// StorageNode return node storage params
func (a *Api) StorageNode(sid int) ([]byte, error) {
	uri := fmt.Sprintf("/storage/%d/node", sid)
	bodyResp, err := a.NewRequest(NilPayload, uri, requestTypeGet, DefaultService)

	return bodyResp, err
}

// StorageLocal return local storage params
func (a *Api) StorageLocal() ([]byte, error) {
	bodyResp, err := a.NewRequest(NilPayload, "/storage/local", requestTypeGet, DefaultService)

	return bodyResp, err
}

// StorageNew create new storage
func (a *Api) StorageNew(params *types.StorageNew) (*types.ID, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(payload, "/storage", requestTypePost, DefaultService)

	var r *types.ID
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// StorageUpdate update storage params
func (a *Api) StorageUpdate(sid int, params *types.StorageUpdate) (*types.ID, error) {
	payload, _ := json.Marshal(params)
	uri := fmt.Sprintf("/storage/%d", sid)
	bodyResp, err := a.NewRequest(payload, uri, requestTypePost, DefaultService)

	var r *types.ID
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// StorageStateUpdate update storage state
func (a *Api) StorageStateUpdate(sid int, params *types.State) (*types.ID, error) {
	payload, _ := json.Marshal(params)
	uri := fmt.Sprintf("/storage/%d/change_state", sid)
	bodyResp, err := a.NewRequest(payload, uri, requestTypePost, DefaultService)

	var r *types.ID
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// update node storage state
func (a *Api) NodeStorageStateUpdate(sid int, nid int, params *types.State) (*types.ID, error) {
	payload, _ := json.Marshal(params)
	uri := fmt.Sprintf("/storage/%d/node/%d/change_state", sid, nid)
	bodyResp, err := a.NewRequest(payload, uri, requestTypePost, DefaultService)

	var r *types.ID
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// StorageDelete delete storage
func (a *Api) StorageDelete(sid int) (*types.DeletedResponse, error) {
	uri := fmt.Sprintf("/storage/%d", sid)
	bodyResp, err := a.NewRequest(NilPayload, uri, requestTypeDelete, DefaultService)

	var r *types.DeletedResponse
	json.Unmarshal(bodyResp, &r)
	return r, err
}
