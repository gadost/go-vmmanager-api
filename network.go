package vmmanagerapi

import (
	"encoding/json"
	"fmt"

	"github.com/gadost/go-vmmanager-api/types"
)

// DCNetworkConnect connect DC network to cluster
func (a *Api) DCNetworkConnect(dcnid int, params *types.Clusters) (*types.Task, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/dc_network/%d/cluster", dcnid),
		requestTypePost,
		DefaultService)

	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// IPNetUpdate update ipnet params
func (a *Api) IPNetUpdate(ipnetID int, params *types.IPNetUpdate) (*types.Task, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/ipnet/%d", ipnetID),
		requestTypePost,
		DefaultService)

	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// IPPoolClusterConnect connect IPPool to cluster
func (a *Api) IPPoolClusterConnect(ippoolID int, params *types.IPPoolConnectRequest) (*types.Task, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/ippool/%d/cluster", ippoolID),
		requestTypePost,
		DefaultService)

	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// IPPoolRangeClusterConnect  connect IPPool range to cluster
func (a *Api) IPPoolRangeClusterConnect(ippoolID int, params *types.IPPoolRangeRequest) (*types.Task, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/ippool/%d/range", ippoolID),
		requestTypePost,
		DefaultService)

	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// IPMgr5NetworkMigrate migrate ipmanager
func (a *Api) IPMgr5NetworkMigrate(params *types.IPMgr5MigrateRequest) (*types.Task, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		"/migrate_ipmgr",
		requestTypePost,
		DefaultService)

	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// NodeHetznerIPNew add hetzner IP
func (a *Api) NodeHetznerIPNew(nid int, params *types.HetznerIPs) (*types.Task, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/node/%d/hetzner_ip", nid),
		requestTypePost,
		DefaultService)

	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// NodeIPUpdate update Node IP
func (a *Api) NodeIPUpdate(nid int, params *types.NodeIPUpdateRequest) (*types.Task, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/node/%d/ip", nid),
		requestTypePost,
		DefaultService)

	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// UserspaceIPNetworkNew add IPnetwork to userspace
func (a *Api) UserspaceIPNetworkNew(usid int, params *types.IPNetUpdate) (*types.Task, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/userspace/%d/ipnet", usid),
		requestTypePost,
		DefaultService)

	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// IPPools returns IPPools params
func (a *Api) IPPools() (*types.IPPoolResponse, error) {
	bodyResp, err := a.NewRequest(
		NilPayload,
		"/ippool",
		requestTypeGet,
		DefaultService)

	var t *types.IPPoolResponse
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// IPPoolClusters return ippool clusters list
func (a *Api) IPPoolClusters(ipid int) (*types.IPPoolClusterListResponse, error) {
	bodyResp, err := a.NewRequest(
		NilPayload,
		fmt.Sprintf("/ippool/%d/cluster", ipid),
		requestTypeGet,
		DefaultService)

	var t *types.IPPoolClusterListResponse
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// NodeHetznerIPs returns nodes hatzner ips
func (a *Api) NodeHetznerIPs(nid int) (*types.NodeHetznerIPResponse, error) {
	bodyResp, err := a.NewRequest(
		NilPayload,
		fmt.Sprintf("/node/%d/hetzner_ip", nid),
		requestTypeGet,
		DefaultService)

	var t *types.NodeHetznerIPResponse
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// IPRanges return IP ranges
func (a *Api) IPRanges() ([]byte, error) {
	bodyResp, err := a.NewRequest(
		NilPayload,
		"/range",
		requestTypeGet,
		DefaultService)

	return bodyResp, err
}

// IPRange return IP range params
func (a *Api) IPRange(rid int) ([]byte, error) {
	bodyResp, err := a.NewRequest(
		NilPayload,
		fmt.Sprintf("/range/%d/ip", rid),
		requestTypeGet,
		DefaultService)

	return bodyResp, err
}

// HetznerIPDelete delete hetnzre IP
func (a *Api) HetznerIPDelete(hipid int) (*types.Task, error) {
	bodyResp, err := a.NewRequest(
		NilPayload,
		fmt.Sprintf("/hetzner_ip/%d", hipid),
		requestTypeDelete,
		DefaultService)

	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// HetznerSubNetDelete delete hetzner subnet
func (a *Api) HetznerSubNetDelete(hsid int) (*types.Task, error) {
	bodyResp, err := a.NewRequest(
		NilPayload,
		fmt.Sprintf("/hetzner_subnet/%d", hsid),
		requestTypeDelete,
		DefaultService)

	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// IPNetworkDelete delete ip network
func (a *Api) IPNetworkDelete(ipnid int) (*types.DeletedResponse, error) {
	bodyResp, err := a.NewRequest(
		NilPayload,
		fmt.Sprintf("/ipnet/%d", ipnid),
		requestTypeDelete,
		DefaultService)

	var t *types.DeletedResponse
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// IPPoolDelete delete ip pool
func (a *Api) IPPoolDelete(ippid int) (*types.DeletedResponse, error) {
	bodyResp, err := a.NewRequest(
		NilPayload,
		fmt.Sprintf("/ippool/%d", ippid),
		requestTypeDelete,
		DefaultService)

	var t *types.DeletedResponse
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// IPRangeDelete delete ip range
func (a *Api) IPRangeDelete(iprid int) (*types.DeletedResponse, error) {
	bodyResp, err := a.NewRequest(
		NilPayload,
		fmt.Sprintf("/range/%d", iprid),
		requestTypeDelete,
		DefaultService)

	var t *types.DeletedResponse
	json.Unmarshal(bodyResp, &t)
	return t, err
}
