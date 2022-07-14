package vmmanagerapi

import (
	"encoding/json"
	"fmt"

	"github.com/gadost/go-vmmanager-api/types"
)

func (a *Api) NodeScriptDelete(nsid int) (*types.DeletedResponse, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/node_script/%d", nsid),
		requestTypeDelete,
		DefaultService)

	var t *types.DeletedResponse
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) NodeScripts() (*types.NodeScriptResponse, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		"/node_script",
		requestTypeGet,
		DefaultService)

	var t *types.NodeScriptResponse
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) NodeScript(nsid int) (*types.NodeScriptElement, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/node_script/%d", nsid),
		requestTypeGet,
		DefaultService)

	var t *types.NodeScriptElement
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) NodeScriptNew(params *types.NodeScriptRequest) (*types.ID, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		"/node_script",
		requestTypePost,
		DefaultService)
	var t *types.ID
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) NodeScriptUpdate(nsid int, params *types.NodeScriptRequest) (*types.Task, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/node_script/%d", nsid),
		requestTypePost,
		DefaultService)
	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) NodeScriptRun(nid int, params *types.Script) (*types.Task, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/node/%d/run_script", nid),
		requestTypePost,
		DefaultService)
	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}
