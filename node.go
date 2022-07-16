package vmmanagerapi

import (
	"encoding/json"
	"fmt"

	"github.com/gadost/go-vmmanager-api/types"
)

// Nodes returns Nodes params
func (a *Api) Nodes(where *ParamsQuery) (*types.NodesResponse, error) {
	bodyResp, err := a.NewRequest(
		NilPayload,
		fmt.Sprintf("/node%s", where),
		requestTypeGet,
		DefaultService)

	var t *types.NodesResponse
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// NodeCheck check Node
func (a *Api) NodeCheck() ([]byte, error) {
	bodyResp, err := a.NewRequest(
		NilPayload,
		"/node_check",
		requestTypeGet,
		DefaultService)

	return bodyResp, err
}

// NodeClustersShort returns short node clusters info
func (a *Api) NodeClustersShort() (*types.NodeReducedClusterListResponse, error) {
	bodyResp, err := a.NewRequest(
		NilPayload,
		"/node_short",
		requestTypeGet,
		DefaultService)

	var t *types.NodeReducedClusterListResponse
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// Node return node params
func (a *Api) Node(nid int) (*types.NodeResponse, error) {
	bodyResp, err := a.NewRequest(
		NilPayload,
		fmt.Sprintf("/node/%d", nid),
		requestTypeGet,
		DefaultService)

	var t *types.NodeResponse
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// NodeFiles returns filelist on node
func (a *Api) NodeFiles(nid int) (*types.NodeFileListResponse, error) {
	bodyResp, err := a.NewRequest(
		NilPayload,
		fmt.Sprintf("/node/%d/file", nid),
		requestTypeGet,
		DefaultService)

	var t *types.NodeFileListResponse
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// NodeHeader return node header
func (a *Api) NodeHeader(nid int) ([]byte, error) {
	bodyResp, err := a.NewRequest(
		NilPayload,
		fmt.Sprintf("/node/%d/header", nid),
		requestTypeGet,
		DefaultService)

	return bodyResp, err
}

// NodeHistory returns node history
func (a *Api) NodeHistory(nid int) ([]byte, error) {
	bodyResp, err := a.NewRequest(
		NilPayload,
		fmt.Sprintf("/node/%d/history", nid),
		requestTypeGet,
		DefaultService)

	return bodyResp, err
}

// NodeHostInterfaces return host interfaces on node
func (a *Api) NodeHostInterfaces(nid int) (*types.NodeHostIFaces, error) {
	bodyResp, err := a.NewRequest(
		NilPayload,
		fmt.Sprintf("/node/%d/host/ifaces", nid),
		requestTypeGet,
		DefaultService)

	var t *types.NodeHostIFaces
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// NodeHosts returns hosts on node
func (a *Api) NodeHosts(nid int) ([]byte, error) {
	bodyResp, err := a.NewRequest(
		NilPayload,
		fmt.Sprintf("/node/%d/hosts", nid),
		requestTypeGet,
		DefaultService)

	return bodyResp, err
}

// NodeLocalStorage returns node local storages
func (a *Api) NodeLocalStorage(nid int) ([]byte, error) {
	bodyResp, err := a.NewRequest(
		NilPayload,
		fmt.Sprintf("/node/%d/local_storage", nid),
		requestTypeGet,
		DefaultService)

	return bodyResp, err
}

// NodeMetrics returns node metrics
func (a *Api) NodeMetrics(nid int) ([]byte, error) {
	bodyResp, err := a.NewRequest(
		NilPayload,
		fmt.Sprintf("/node/%d/metrics", nid),
		requestTypeGet,
		DefaultService)

	return bodyResp, err
}

// NodeNetworkInterfaces return node interfaces
func (a *Api) NodeNetworkInterfaces(nid int) (*types.NodeNetworkIFaces, error) {
	bodyResp, err := a.NewRequest(
		NilPayload,
		fmt.Sprintf("/node/%d/network/iface", nid),
		requestTypeGet,
		DefaultService)

	var t *types.NodeNetworkIFaces
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// NodeStorageConnectionParams returns node storage connection params
func (a *Api) NodeStorageConnectionParams(nid int, sid int) ([]byte, error) {
	bodyResp, err := a.NewRequest(
		NilPayload,
		fmt.Sprintf("/node/%d/storage/%d/connect_params", nid, sid),
		requestTypeGet,
		DefaultService)

	return bodyResp, err
}

// NodeNew create new node
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

// NodeUpdate update node params
func (a *Api) NodeUpdate(nid int, params *types.NodeUpdateRequest) (*types.ID, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/node/%d", nid),
		requestTypePost,
		DefaultService)
	var t *types.ID
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// NodeLibvirtCertUpdate update node's libvirt certificate
func (a *Api) NodeLibvirtCertUpdate(nid int) (*types.Task, error) {
	bodyResp, err := a.NewRequest(
		NilPayload,
		fmt.Sprintf("/node/%d/cert", nid),
		requestTypePost,
		DefaultService)
	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// NodeFilesUpdate update files on node
func (a *Api) NodeFilesUpdate(nid int, params *types.NodeFilesUpdateRequest) (*types.ID, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/node/%d/files", nid),
		requestTypePost,
		DefaultService)
	var t *types.ID
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// NodeHAEnable enable node HA
func (a *Api) NodeHAEnable(nid int, params interface{}) (*types.ID, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/node/%d/ha/enable", nid),
		requestTypePost,
		DefaultService)
	var t *types.ID
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// NodeHAStateUpdate update Node HA state
func (a *Api) NodeHAStateUpdate(nid int, params *types.State) (*types.ID, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/node/%d/ha/state", nid),
		requestTypePost,
		DefaultService)
	var t *types.ID
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// NodeHostInterfaceUpdate update host interface on node
func (a *Api) NodeHostInterfaceUpdate(nid int, params *types.NodeHostIFaces) (*types.Task, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/node/%d/host/ifaces", nid),
		requestTypePost,
		DefaultService)
	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// NodeHostStatesUpdate update hosy state on node
func (a *Api) NodeHostStatesUpdate(nid int, params interface{}) (*types.Task, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/node/%d/host/states", nid),
		requestTypePost,
		DefaultService)
	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// NodeImportedConfigure import  node configuraton
func (a *Api) NodeImportedConfigure(nid int, params *types.NetworkAutosetupDisabled) (*types.Task, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/node/%d/import/configure", nid),
		requestTypePost,
		DefaultService)
	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// NodeInterfaceNew add new interface to node
func (a *Api) NodeInterfaceNew(nid int, params *types.NodeNetworkInterface) (*types.Task, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/node/%d/network/iface", nid),
		requestTypePost,
		DefaultService)
	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// NodeInterfaceUpdate update interface params on node
func (a *Api) NodeInterfaceUpdate(nid int, iface string, params *types.NodeNetworkInterface) (*types.Task, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/node/%d/network/iface/%s", nid, iface),
		requestTypePost,
		DefaultService)
	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// NodeInterfaceRevert revert interface changes on node
func (a *Api) NodeInterfaceRevert(nid int, iface string, params interface{}) (*types.Task, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/node/%d/network/iface/%s/revert", nid, iface),
		requestTypePost,
		DefaultService)
	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// NodeInterfacesUpdate update  interface params on node
func (a *Api) NodeInterfacesUpdate(nid int, params *types.NodeNetworkInterfaces) (*types.Task, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/node/%d/network/ifaces", nid),
		requestTypePost,
		DefaultService)
	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// NodeNetworkLock lock node network
func (a *Api) NodeNetworkLock(nid int, params interface{}) (*types.Task, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/node/%d/network/lock", nid),
		requestTypePost,
		DefaultService)
	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// NodeNetworkRevert revert network changes
func (a *Api) NodeNetworkRevert(nid int, params interface{}) (*types.Task, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/node/%d/network/revert", nid),
		requestTypePost,
		DefaultService)
	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// NodeNetworkSave save node network
func (a *Api) NodeNetworkSave(nid int, params interface{}) (*types.Task, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/node/%d/network/save", nid),
		requestTypePost,
		DefaultService)
	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// NodeNetworkUnlock unlock node network
func (a *Api) NodeNetworkUnlock(nid int, params interface{}) (*types.Task, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/node/%d/network/unlock", nid),
		requestTypePost,
		DefaultService)
	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// NodeProblemsNew create new problem
func (a *Api) NodeProblemsNew(nid int, params *types.Problems) (*types.Task, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/node/%d/problem", nid),
		requestTypePost,
		DefaultService)
	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// NodeRelocate start relocate node
func (a *Api) NodeRelocate(nid int, params interface{}) (*types.Task, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/node/%d/relocate", nid),
		requestTypePost,
		DefaultService)
	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// NodeStorageCheck check  storage on node
func (a *Api) NodeStorageCheck(cid int, sid int) (*types.Task, error) {
	bodyResp, err := a.NewRequest(
		NilPayload,
		fmt.Sprintf("/node/%d/storage/%d/check", cid, sid),
		requestTypePost,
		DefaultService)

	var r *types.Task
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// NodeHDDOverselling set node hdd  overselling
func (a *Api) NodeHDDOverselling(
	nid int, params *types.HDDOversellingRequest) (*types.HDDOversellingResponse, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/node/%d/problem", nid),
		requestTypePost,
		DefaultService)
	var t *types.HDDOversellingResponse
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// NodeDelete delete node
func (a *Api) NodeDelete(nsid int) (*types.Task, error) {
	bodyResp, err := a.NewRequest(
		NilPayload,
		fmt.Sprintf("/node/%d", nsid),
		requestTypeDelete,
		DefaultService)

	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// NodeNetworkInterfaceDelete delete node interface
func (a *Api) NodeNetworkInterfaceDelete(nsid int, iface string) (*types.Task, error) {
	bodyResp, err := a.NewRequest(
		NilPayload,
		fmt.Sprintf("/node/%d/network/iface/%s", nsid, iface),
		requestTypeDelete,
		DefaultService)

	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}
