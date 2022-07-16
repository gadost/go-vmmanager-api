package vmmanagerapi

import (
	"encoding/json"
	"fmt"

	"github.com/gadost/go-vmmanager-api/types"
)

func (a *Api) Repositories() (*types.Repositories, error) {
	bodyResp, err := a.NewRequest(NilPayload, "/repository", requestTypeGet, DefaultService)

	var r *types.Repositories
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) Repository(repoid int) (*types.RepositoryName, error) {
	uri := fmt.Sprintf("/repository/%d", repoid)
	bodyResp, err := a.NewRequest(NilPayload, uri, requestTypeGet, DefaultService)

	var r *types.RepositoryName
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) RepositoryNew(params *types.RepositoryNew) (*types.Tasks, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(payload, "/repository", requestTypePost, DefaultService)

	var r *types.Tasks
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) RepositoryUpdate(repoid int, params *types.RepositoryNew) (*types.Task, error) {
	uri := fmt.Sprintf("/repository/%d", repoid)
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(payload, uri, requestTypePost, DefaultService)

	var r *types.Task
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) RepositoryDelete(repoid int) (*types.DeletedResponse, error) {
	uri := fmt.Sprintf("/repository/%d", repoid)
	bodyResp, err := a.NewRequest(NilPayload, uri, requestTypeDelete, DefaultService)

	var r *types.DeletedResponse
	json.Unmarshal(bodyResp, &r)
	return r, err
}
