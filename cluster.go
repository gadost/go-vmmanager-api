package vmmanagerapi

import (
	"encoding/json"
	"fmt"

	"github.com/gadost/go-vmmanager-api/types"
)

func (a *Api) ClusterList() (*types.ClusterListResponse, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		"/cluster",
		requestTypeGet,
		DefaultService)

	var c *types.ClusterListResponse
	json.Unmarshal(bodyResp, &c)

	return c, err
}

func (a *Api) ClusterNew(params *types.CreateClusterRequest) (*types.Task, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		"/cluster",
		requestTypePost,
		DefaultService)
	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) ClusterGet(cid int) (*types.ClusterResponse, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/cluster/%d", cid),
		requestTypeGet,
		DefaultService)
	var c *types.ClusterResponse
	json.Unmarshal(bodyResp, &c)
	return c, err
}

func (a *Api) ClusterUpdate(cid int, params *types.CreateClusterRequest) (*types.Tasks, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/cluster/%d", cid),
		requestTypePost,
		DefaultService)
	var t *types.Tasks
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) ClusterDelete(cid int) (*types.DeletedResponse, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/cluster/%d", cid),
		requestTypeDelete,
		DefaultService)
	var d *types.DeletedResponse
	json.Unmarshal(bodyResp, &d)
	return d, err
}

func (a *Api) ClusterConnectDCNetwork(cid int) (*types.DCNetworksResponse, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/cluster/%d/dc_network", cid),
		requestTypePost,
		DefaultService)
	var d *types.DCNetworksResponse
	json.Unmarshal(bodyResp, &d)
	return d, err
}

func (a *Api) ClusterFixFRR(cid int) error {
	_, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/cluster/%d/fix_frr", cid),
		requestTypePost,
		DefaultService)

	return err
}

func (a *Api) ClusterHaAgentUpdate(cid int) (*types.Task, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/cluster/%d/ha_agent_update", cid),
		requestTypePost,
		DefaultService)
	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) ClusterInternalEditUpdate(cid int, params *types.QemuVersion) (*types.Task, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/cluster/%d/internal_edit", cid),
		requestTypePost,
		DefaultService)
	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) ClusterConnectIPPool(cid int, params *types.IPPoolRequest) error {
	payload, _ := json.Marshal(params)
	_, err := a.NewRequest(
		payload,
		fmt.Sprintf("/cluster/%d/ippool", cid),
		requestTypePost,
		DefaultService)

	return err
}

func (a *Api) ClusterLocalStorageGet(cid int) error {
	_, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/cluster/%d/local_storage", cid),
		requestTypeGet,
		DefaultService)

	return err
}

func (a *Api) ClusterSettingsUpdate(cid int, params *types.Name) error {
	payload, _ := json.Marshal(params)
	_, err := a.NewRequest(
		payload,
		fmt.Sprintf("/cluster/%d/settings", cid),
		requestTypePost,
		DefaultService)

	return err
}

func (a *Api) ClusterSSHKeysUpdate(cid int, params *types.SSHKeysRequest) error {
	payload, _ := json.Marshal(params)
	_, err := a.NewRequest(
		payload,
		fmt.Sprintf("/cluster/%d/ssh_key", cid),
		requestTypePost,
		DefaultService)

	return err
}

func (a *Api) ClusterStorageAttach(
	cid int, sid int, params *types.AttachStorageToClusterRequest) (
	*types.StoragesTasks, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/cluster/%d/storage/%d", cid, sid),
		requestTypePost,
		DefaultService)

	var r *types.StoragesTasks
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) ClusterStorageDetach(cid int, sid int) (*types.DeletedResponse, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/cluster/%d/storage/%d", cid, sid),
		requestTypeDelete,
		DefaultService)
	var d *types.DeletedResponse
	json.Unmarshal(bodyResp, &d)
	return d, err
}

func (a *Api) ClusterStorageCephConnect(
	cid int, sid int) (
	*types.StoragesTasks, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/cluster/%d/storage/%d/ceph_connect", cid, sid),
		requestTypePost,
		DefaultService)

	var r *types.StoragesTasks
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) ClusterStorageCheck(cid int, sid int) (*types.Task, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/cluster/%d/storage/%d/check", cid, sid),
		requestTypePost,
		DefaultService)

	var r *types.Task
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) ClusterStorageDisable(cid int, sid int) (*types.Task, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/cluster/%d/storage/%d/disable", cid, sid),
		requestTypePost,
		DefaultService)

	var r *types.Task
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) ClusterStorageEnable(cid int, sid int) (*types.Task, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/cluster/%d/storage/%d/enable", cid, sid),
		requestTypePost,
		DefaultService)

	var r *types.Task
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) ClusteHDDOversellingUpdate(
	cid int, sid int, params *types.HDDOversellingRequest) (
	*types.HDDOversellingResponse, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/cluster/%d/storage/%d/hdd_overselling", cid, sid),
		requestTypePost,
		DefaultService)

	var r *types.HDDOversellingResponse
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) ClusterStorageMakeMain(cid int, sid int) (*types.Task, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/cluster/%d/storage/%d/make_main", cid, sid),
		requestTypePost,
		DefaultService)

	var r *types.Task
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) ClusterVXLanGet(cid int) (*types.VXLanResponse, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/cluster/%d/vxlan", cid),
		requestTypeGet,
		DefaultService)

	var r *types.VXLanResponse
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) ClusterNodesVXLanGet(cid int) (*types.NodesVXLanResponse, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/cluster/%d/vxlan/node", cid),
		requestTypeGet,
		DefaultService)

	var r *types.NodesVXLanResponse
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) ClusterUsersVXLanGet(cid int) (*types.UsersVXLanResponse, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/cluster/%d/vxlan/user", cid),
		requestTypeGet,
		DefaultService)

	var r *types.UsersVXLanResponse
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) ClusterSSHKeyList() (*types.SSHKeyResponse, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		"/ssh_key",
		requestTypeGet,
		DefaultService)

	var r *types.SSHKeyResponse
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) ClusterAddSSHKey(params *types.SSHKeyRequest) (*types.Task, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		"/ssh_key",
		requestTypePost,
		DefaultService)

	var r *types.Task
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) ClusterDeleteSSHKey(kid int) (*types.Task, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/ssh_key/%d", kid),
		requestTypeDelete,
		DefaultService)

	var r *types.Task
	json.Unmarshal(bodyResp, &r)
	return r, err
}
