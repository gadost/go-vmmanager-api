package vmmanagerapi

import (
	"encoding/json"
	"fmt"

	"github.com/gadost/go-vmmanager-api/types"
)

func (a *Api) Clusters() (*types.ClusterListResponse, error) {
	uri := ClusterUri
	bodyResp, err := a.NewRequest(NilPayload, uri, requestTypeGet, DefaultService)

	var c *types.ClusterListResponse
	json.Unmarshal(bodyResp, &c)

	return c, err
}

func (a *Api) ClusterNew(params *types.CreateClusterRequest) (*types.Task, error) {
	payload, _ := json.Marshal(params)
	uri := ClusterUri
	bodyResp, err := a.NewRequest(payload, uri, requestTypePost, DefaultService)

	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) Cluster(cid int) (*types.ClusterResponse, error) {
	uri := fmt.Sprintf("%s/%d", ClusterUri, cid)
	bodyResp, err := a.NewRequest(NilPayload, uri, requestTypeGet, DefaultService)

	var c *types.ClusterResponse
	json.Unmarshal(bodyResp, &c)
	return c, err
}

func (a *Api) ClusterUpdate(cid int, params *types.CreateClusterRequest) (*types.Tasks, error) {
	payload, _ := json.Marshal(params)
	uri := fmt.Sprintf("%s/%d", ClusterUri, cid)
	bodyResp, err := a.NewRequest(payload, uri, requestTypePost, DefaultService)
	var t *types.Tasks
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) ClusterDelete(cid int) (*types.DeletedResponse, error) {
	uri := fmt.Sprintf("%s/%d", ClusterUri, cid)
	bodyResp, err := a.NewRequest(NilPayload, uri, requestTypeDelete, DefaultService)

	var d *types.DeletedResponse
	json.Unmarshal(bodyResp, &d)
	return d, err
}

func (a *Api) ClusterConnectDCNetwork(cid int) (*types.DCNetworksResponse, error) {
	uri := fmt.Sprintf("%s/%d/dc_network", ClusterUri, cid)
	bodyResp, err := a.NewRequest(NilPayload, uri, requestTypePost, DefaultService)

	var d *types.DCNetworksResponse
	json.Unmarshal(bodyResp, &d)
	return d, err
}

func (a *Api) ClusterFixFRR(cid int) ([]byte, error) {
	uri := fmt.Sprintf("%s/%d/fix_frr", ClusterUri, cid)
	bodyResp, err := a.NewRequest(NilPayload, uri, requestTypePost, DefaultService)

	return bodyResp, err
}

func (a *Api) ClusterHaAgentUpdate(cid int) (*types.Task, error) {
	uri := fmt.Sprintf("%s/%d/ha_agent_update", ClusterUri, cid)
	bodyResp, err := a.NewRequest(NilPayload, uri, requestTypePost, DefaultService)

	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) ClusterInternalEditUpdate(cid int, params *types.QemuVersion) (*types.Task, error) {
	payload, _ := json.Marshal(params)
	uri := fmt.Sprintf("%s/%d/internal_edit", ClusterUri, cid)
	bodyResp, err := a.NewRequest(
		payload,
		uri,
		requestTypePost,
		DefaultService)
	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) ClusterConnectIPPool(cid int, params *types.IPPoolRequest) ([]byte, error) {
	payload, _ := json.Marshal(params)
	uri := fmt.Sprintf("%s/%d/ippool", ClusterUri, cid)
	bodyResp, err := a.NewRequest(payload, uri, requestTypePost, DefaultService)

	return bodyResp, err
}

func (a *Api) ClusterLocalStorage(cid int) ([]byte, error) {
	uri := fmt.Sprintf("%s/%d/local_storage", ClusterUri, cid)
	bodyResp, err := a.NewRequest(NilPayload, uri, requestTypeGet, DefaultService)

	return bodyResp, err
}

func (a *Api) ClusterSettingsUpdate(cid int, params *types.Name) ([]byte, error) {
	payload, _ := json.Marshal(params)
	uri := fmt.Sprintf("%s/%d/settings", ClusterUri, cid)
	bodyResp, err := a.NewRequest(payload, uri, requestTypePost, DefaultService)

	return bodyResp, err
}

func (a *Api) ClusterSSHKeysUpdate(cid int, params *types.SSHKeysRequest) ([]byte, error) {
	payload, _ := json.Marshal(params)
	uri := fmt.Sprintf("%s/%d/ssh_key", ClusterUri, cid)
	bodyResp, err := a.NewRequest(payload, uri, requestTypePost, DefaultService)

	return bodyResp, err
}

func (a *Api) ClusterStorageAttach(
	cid int, sid int, params *types.AttachStorageToClusterRequest) (
	*types.StoragesTasks, error) {
	payload, _ := json.Marshal(params)
	uri := fmt.Sprintf("%s/%d/storage/%d", ClusterUri, cid, sid)
	bodyResp, err := a.NewRequest(payload, uri, requestTypePost, DefaultService)

	var r *types.StoragesTasks
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) ClusterStorageDetach(cid int, sid int) (*types.DeletedResponse, error) {
	uri := fmt.Sprintf("%s/%d/storage/%d", ClusterUri, cid, sid)
	bodyResp, err := a.NewRequest(NilPayload, uri, requestTypeDelete, DefaultService)

	var d *types.DeletedResponse
	json.Unmarshal(bodyResp, &d)
	return d, err
}

func (a *Api) ClusterStorageCephConnect(cid int, sid int) (*types.StoragesTasks, error) {
	uri := fmt.Sprintf("%s/%d/storage/%d/ceph_connect", ClusterUri, cid, sid)
	bodyResp, err := a.NewRequest(NilPayload, uri, requestTypePost, DefaultService)

	var r *types.StoragesTasks
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) ClusterStorageCheck(cid int, sid int) (*types.Task, error) {
	uri := fmt.Sprintf("%s/%d/storage/%d/check", ClusterUri, cid, sid)
	bodyResp, err := a.NewRequest(NilPayload, uri, requestTypePost, DefaultService)

	var r *types.Task
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) ClusterStorageDisable(cid int, sid int) (*types.Task, error) {
	uri := fmt.Sprintf("%s/%d/storage/%d/disable", ClusterUri, cid, sid)
	bodyResp, err := a.NewRequest(NilPayload, uri, requestTypePost, DefaultService)

	var r *types.Task
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) ClusterStorageEnable(cid int, sid int) (*types.Task, error) {
	uri := fmt.Sprintf("%s/%d/storage/%d/enable", ClusterUri, cid, sid)
	bodyResp, err := a.NewRequest(NilPayload, uri, requestTypePost, DefaultService)

	var r *types.Task
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) ClusteHDDOversellingUpdate(
	cid int, sid int, params *types.HDDOversellingRequest) (
	*types.HDDOversellingResponse, error) {
	payload, _ := json.Marshal(params)
	uri := fmt.Sprintf("%s/%d/storage/%d/hdd_overselling", ClusterUri, cid, sid)
	bodyResp, err := a.NewRequest(payload, uri, requestTypePost, DefaultService)

	var r *types.HDDOversellingResponse
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) ClusterStorageMakeMain(cid int, sid int) (*types.Task, error) {
	uri := fmt.Sprintf("%s/%d/storage/%d/make_main", ClusterUri, cid, sid)
	bodyResp, err := a.NewRequest(NilPayload, uri, requestTypePost, DefaultService)

	var r *types.Task
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) ClusterVXLan(cid int) (*types.VXLanResponse, error) {
	uri := fmt.Sprintf("%s/%d/vxlan", ClusterUri, cid)
	bodyResp, err := a.NewRequest(NilPayload, uri, requestTypeGet, DefaultService)

	var r *types.VXLanResponse
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) ClusterNodesVXLan(cid int) (*types.NodesVXLanResponse, error) {
	uri := fmt.Sprintf("%s/%d/vxlan/node", ClusterUri, cid)
	bodyResp, err := a.NewRequest(NilPayload, uri, requestTypeGet, DefaultService)

	var r *types.NodesVXLanResponse
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) ClusterUsersVXLan(cid int) (*types.UsersVXLanResponse, error) {
	uri := fmt.Sprintf("%s/%d/vxlan/user", ClusterUri, cid)
	bodyResp, err := a.NewRequest(NilPayload, uri, requestTypeGet, DefaultService)

	var r *types.UsersVXLanResponse
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) SSHKeys() (*types.SSHKeyResponse, error) {
	bodyResp, err := a.NewRequest(NilPayload, SSHUri, requestTypeGet, DefaultService)

	var r *types.SSHKeyResponse
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) SSHKeyNew(params *types.SSHKeyRequest) (*types.Task, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(payload, SSHUri, requestTypePost, DefaultService)

	var r *types.Task
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) SSHKeyDelete(kid int) (*types.Task, error) {
	uri := fmt.Sprintf("%s/%d", SSHUri, kid)
	bodyResp, err := a.NewRequest(NilPayload, uri, requestTypeDelete, DefaultService)

	var r *types.Task
	json.Unmarshal(bodyResp, &r)
	return r, err
}
