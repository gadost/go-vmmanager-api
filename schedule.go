package vmmanagerapi

import (
	"encoding/json"
	"fmt"

	"github.com/gadost/go-vmmanager-api/types"
)

func (a *Api) MetricsSendReport() (*types.Task, error) {
	bodyResp, err := a.NewRequest(NilPayload, "/metrics/sendreport", requestTypePost, DefaultService)

	var r *types.Task
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) Schedules() (*types.Schedules, error) {
	bodyResp, err := a.NewRequest(NilPayload, "/schedule", requestTypeGet, DefaultService)

	var r *types.Schedules
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) ScheduleNew(params *types.Schedules) (*types.Task, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(payload, "/schedule", requestTypePost, DefaultService)

	var r *types.Task
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) ScheduleUpdate(schid int, params *types.Schedules) (*types.Task, error) {
	payload, _ := json.Marshal(params)
	uri := fmt.Sprintf("/schedule/%d", schid)
	bodyResp, err := a.NewRequest(payload, uri, requestTypePost, DefaultService)

	var r *types.Task
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) ScheduleDelete(schid int) (*types.DeletedResponse, error) {
	uri := fmt.Sprintf("/schedule/%d", schid)
	bodyResp, err := a.NewRequest(NilPayload, uri, requestTypeDelete, DefaultService)

	var r *types.DeletedResponse
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) ScheduleHostsAffected(schid int) (*types.AffectedHostsCount, error) {
	uri := fmt.Sprintf("/schedule/%d/hosts", schid)
	bodyResp, err := a.NewRequest(NilPayload, uri, requestTypeGet, DefaultService)

	var r *types.AffectedHostsCount
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) ScheduleRun(schid int) (*types.ID, error) {
	uri := fmt.Sprintf("/schedule/%d/run", schid)
	bodyResp, err := a.NewRequest(NilPayload, uri, requestTypePost, DefaultService)

	var r *types.ID
	json.Unmarshal(bodyResp, &r)
	return r, err
}
