package vmmanagerapi

import (
	"encoding/json"
	"fmt"

	"github.com/gadost/go-vmmanager-api/types"
)

func (a *Api) Tasks() (*types.TasksList, error) {
	bodyResp, err := a.NewRequest(NilPayload, "/task", requestTypeGet, DefaultService)

	var r *types.TasksList
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) Task(tid int) (*types.TaskElement, error) {
	uri := fmt.Sprintf("/task/%d", tid)
	bodyResp, err := a.NewRequest(NilPayload, uri, requestTypeGet, DefaultService)

	var r *types.TaskElement
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) TagAbort(tid int) (*types.Task, error) {
	uri := fmt.Sprintf("/task/%d/abort", tid)
	bodyResp, err := a.NewRequest(NilPayload, uri, requestTypePost, DefaultService)

	var r *types.Task
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) TaskInternalUpdate(tid int, params *types.TaskInternalUpdate) (*types.ID, error) {
	payload, _ := json.Marshal(params)
	uri := fmt.Sprintf("/task/%d/internal_edit", tid)
	bodyResp, err := a.NewRequest(payload, uri, requestTypePost, DefaultService)

	var r *types.ID
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) TaskHASync(params *types.TaskSyncRequest) (*types.Tasks, error) {
	payload, _ := json.Marshal(params)
	uri := "/task/ha/sync"
	bodyResp, err := a.NewRequest(payload, uri, requestTypePost, DefaultService)

	var r *types.Tasks
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) TaskDelete(tid int) (*types.DeletedResponse, error) {
	uri := fmt.Sprintf("/task/%d/delete", tid)
	bodyResp, err := a.NewRequest(NilPayload, uri, requestTypeDelete, DefaultService)

	var r *types.DeletedResponse
	json.Unmarshal(bodyResp, &r)
	return r, err
}
