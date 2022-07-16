package vmmanagerapi

import (
	"encoding/json"
	"fmt"

	"github.com/gadost/go-vmmanager-api/types"
)

// MetricsSendReport send metrics report
func (a *Api) MetricsSendReport() (*types.Task, error) {
	bodyResp, err := a.NewRequest(NilPayload, "/metrics/sendreport", requestTypePost, DefaultService)

	var r *types.Task
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// Schedules return schedules
func (a *Api) Schedules() (*types.Schedules, error) {
	bodyResp, err := a.NewRequest(NilPayload, "/schedule", requestTypeGet, DefaultService)

	var r *types.Schedules
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// ScheduleNew add new schedule
func (a *Api) ScheduleNew(params *types.Schedule) (*types.Task, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(payload, "/schedule", requestTypePost, DefaultService)

	var r *types.Task
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// ScheduleUpdate update schedule
func (a *Api) ScheduleUpdate(schid int, params *types.Schedule) (*types.Task, error) {
	payload, _ := json.Marshal(params)
	uri := fmt.Sprintf("/schedule/%d", schid)
	bodyResp, err := a.NewRequest(payload, uri, requestTypePost, DefaultService)

	var r *types.Task
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// ScheduleDelete delete schedule
func (a *Api) ScheduleDelete(schid int) (*types.DeletedResponse, error) {
	uri := fmt.Sprintf("/schedule/%d", schid)
	bodyResp, err := a.NewRequest(NilPayload, uri, requestTypeDelete, DefaultService)

	var r *types.DeletedResponse
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// ScheduleHostsAffected returns affected host by schedule
func (a *Api) ScheduleHostsAffected(schid int) (*types.AffectedHostsCount, error) {
	uri := fmt.Sprintf("/schedule/%d/hosts", schid)
	bodyResp, err := a.NewRequest(NilPayload, uri, requestTypeGet, DefaultService)

	var r *types.AffectedHostsCount
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// ScheduleRun runs schedule
func (a *Api) ScheduleRun(schid int) (*types.ID, error) {
	uri := fmt.Sprintf("/schedule/%d/run", schid)
	bodyResp, err := a.NewRequest(NilPayload, uri, requestTypePost, DefaultService)

	var r *types.ID
	json.Unmarshal(bodyResp, &r)
	return r, err
}
