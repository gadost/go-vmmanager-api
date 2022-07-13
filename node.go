package vmmanagerapi

import (
	"encoding/json"
	"fmt"

	"github.com/gadost/go-vmmanager-api/types"
)

func (a *Api) Nodes(where *ParamsQuery) (*types.NodesResponse, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/node%s", where),
		requestTypeGet,
		DefaultService)

	var t *types.NodesResponse
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) NodeCheck() ([]byte, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		"/node_check",
		requestTypeGet,
		DefaultService)

	return bodyResp, err
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

func (a *Api) NodedShortClusterList() (*types.NodeReducedClusterListResponse, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		"/node_short",
		requestTypeGet,
		DefaultService)

	var t *types.NodeReducedClusterListResponse
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) Node(nid int) (*types.NodeResponse, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/node/%d", nid),
		requestTypeGet,
		DefaultService)

	var t *types.NodeResponse
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) NodeFiles(nid int) (*types.NodeFileListResponse, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/node/%d/file", nid),
		requestTypeGet,
		DefaultService)

	var t *types.NodeFileListResponse
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) NodeHeader(nid int) ([]byte, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/node/%d/header", nid),
		requestTypeGet,
		DefaultService)

	return bodyResp, err
}

func (a *Api) NodeHistory(nid int) ([]byte, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/node/%d/history", nid),
		requestTypeGet,
		DefaultService)

	return bodyResp, err
}

func (a *Api) NodeHostIFaces(nid int) (*types.NodeHostIFaces, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/node/%d/host/ifaces", nid),
		requestTypeGet,
		DefaultService)

	var t *types.NodeHostIFaces
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) NodeHosts(nid int) ([]byte, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/node/%d/hosts", nid),
		requestTypeGet,
		DefaultService)

	return bodyResp, err
}

func (a *Api) NodeLocalStorage(nid int) ([]byte, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/node/%d/local_storage", nid),
		requestTypeGet,
		DefaultService)

	return bodyResp, err
}

func (a *Api) NodeMetrics(nid int) ([]byte, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/node/%d/metrics", nid),
		requestTypeGet,
		DefaultService)

	return bodyResp, err
}

func (a *Api) NodeNetworkIFaces(nid int) (*types.NodeNetworkIFaces, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/node/%d/network/iface", nid),
		requestTypeGet,
		DefaultService)

	var t *types.NodeNetworkIFaces
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) NodeStorageConnectionParams(nid int, sid int) ([]byte, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/node/%d/storage/%d/connect_params", nid, sid),
		requestTypeGet,
		DefaultService)

	return bodyResp, err
}

func (a *Api) NodeNew(params *types.NodeNewRequest) (*types.StorageTask, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		"/node",
		requestTypePost,
		DefaultService)
	var t *types.StorageTask
	json.Unmarshal(bodyResp, &t)
	return t, err
}
