package vmmanagerapi

import (
	"encoding/json"
	"fmt"

	"github.com/gadost/go-vmmanager-api/types"
)

// VXLans return VX Lans
func (a *Api) VXLans() (*types.VXLans, error) {
	bodyResp, err := a.NewRequest(NilPayload, "/vxlan", requestTypeGet, DefaultService)

	var r *types.VXLans
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// VXLan returns VX Lan params
func (a *Api) VXLan(vid int) (*types.VXLan, error) {
	uri := fmt.Sprintf("/vxlan/%d", vid)
	bodyResp, err := a.NewRequest(NilPayload, uri, requestTypeGet, DefaultService)

	var r *types.VXLan
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// VXLanIPs returns VX lan IPs
func (a *Api) VXLanIPs(vid int) (*types.VXLanIPs, error) {
	uri := fmt.Sprintf("/vxlan/%d/ip", vid)
	bodyResp, err := a.NewRequest(NilPayload, uri, requestTypeGet, DefaultService)

	var r *types.VXLanIPs
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// VXLanNew create new VX lan
func (a *Api) VXLanNew(params *types.VXLanNew) (*types.ID, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(payload, "/vxlan", requestTypePost, DefaultService)

	var r *types.ID
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// VXLanUpdate update VX lan params
func (a *Api) VXLanUpdate(vid int, params *types.VXLanNew) (*types.ID, error) {
	payload, _ := json.Marshal(params)
	uri := fmt.Sprintf("/vxlan/%d", vid)
	bodyResp, err := a.NewRequest(payload, uri, requestTypePost, DefaultService)

	var r *types.ID
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// VXLanIPNetNew add IP net to VX lan
func (a *Api) VXLanIPNetNew(vid int, params *types.VXLanIPNet) (*types.ID, error) {
	payload, _ := json.Marshal(params)
	uri := fmt.Sprintf("/vxlan/%d/ipnet", vid)
	bodyResp, err := a.NewRequest(payload, uri, requestTypePost, DefaultService)

	var r *types.ID
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// VXLanIPNetUpdate update VX lan IPnet params
func (a *Api) VXLanIPNetUpdate(vid int, ipnid int, params *types.VXLanIPNet) (*types.ID, error) {
	payload, _ := json.Marshal(params)
	uri := fmt.Sprintf("/vxlan/%d/ipnet/%d", vid, ipnid)
	bodyResp, err := a.NewRequest(payload, uri, requestTypePost, DefaultService)

	var r *types.ID
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// VXLanIPNetDelete delete VX lan IP net
func (a *Api) VXLanIPNetDelete(vid int, ipnid int) (*types.DeletedResponse, error) {
	uri := fmt.Sprintf("/vxlan/%d/ipnet/%d", vid, ipnid)
	bodyResp, err := a.NewRequest(NilPayload, uri, requestTypeDelete, DefaultService)

	var r *types.DeletedResponse
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// VXLanDelete delete VX lan
func (a *Api) VXLanDelete(vid int) (*types.DeletedResponse, error) {
	uri := fmt.Sprintf("/vxlan/%d", vid)
	bodyResp, err := a.NewRequest(NilPayload, uri, requestTypeDelete, DefaultService)

	var r *types.DeletedResponse
	json.Unmarshal(bodyResp, &r)
	return r, err
}
