package vmmanagerapi

import (
	"encoding/json"
	"fmt"

	"github.com/gadost/go-vmmanager-api/types"
)

func (a *Api) HostGet(where *ParamsQuery) (*types.HostsResponse, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/host%s", where),
		requestTypeGet,
		DefaultService)
	var h *types.HostsResponse
	json.Unmarshal(bodyResp, &h)
	return h, err
}

func (a *Api) HostNew(hid int, params *types.NewHostRequest) (*types.RecipeTask, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		"/host",
		requestTypePost,
		DefaultService)
	var t *types.RecipeTask
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) HostGetByID(hid int) (*types.HostResponse, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/host/%d", hid),
		requestTypeGet,
		DefaultService)
	var h *types.HostResponse
	json.Unmarshal(bodyResp, &h)
	return h, err
}

func (a *Api) HostUpdate(hid int, params *types.HostUpdateRequest) (*types.Task, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/host/%d", hid),
		requestTypePost,
		DefaultService)
	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) HostDelete(hid int) (*types.Task, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/host/%d", hid),
		requestTypeDelete,
		DefaultService)
	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) HostChangeOwner(hid int, params *types.Account) (*types.Task, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/host/%d/account", hid),
		requestTypePost,
		DefaultService)
	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) HostImageSizeLimit(hid int, params *types.ImageSize) (*types.ImageSize, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/host/%d/bill_option/image_gib", hid),
		requestTypePost,
		DefaultService)
	var i *types.ImageSize
	json.Unmarshal(bodyResp, &i)
	return i, err
}

func (a *Api) HostClone(hid int) (*types.Task, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/host/%d/clone", hid),
		requestTypePost,
		DefaultService)
	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) HostRepairGuestAgent(hid int) (*types.Task, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/host/%d/guest_agent", hid),
		requestTypePost,
		DefaultService)
	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) HostHistory(hid int) (*types.HostHistoryResponse, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/host/%d/history", hid),
		requestTypeGet,
		DefaultService)
	var h *types.HostHistoryResponse
	json.Unmarshal(bodyResp, &h)
	return h, err
}

func (a *Api) HostNewLogEntry(hid int, payload []byte) ([]byte, error) {
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/host/%d/histroy", hid),
		requestTypePost,
		DefaultService)

	return bodyResp, err
}

func (a *Api) HostIFace(hid int) (*types.IFace, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/host/%d/iface", hid),
		requestTypeGet,
		DefaultService)
	var h *types.IFace
	json.Unmarshal(bodyResp, &h)
	return h, err
}

func (a *Api) HostNewIFace(hid int, params *types.IFaceParams) ([]byte, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/host/%d/iface", hid),
		requestTypePost,
		DefaultService)

	return bodyResp, err
}

func (a *Api) HostUpdateIFace(hid int, iface string, params *types.IFaceModel) (*types.Task, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/host/%d/iface/%s", hid, iface),
		requestTypePost,
		DefaultService)

	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) HostDeleteIFace(hid int, iface string) (*types.Task, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/host/%d/iface/%s", hid, iface),
		requestTypeDelete,
		DefaultService)

	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) HostAddIP(hid int, params *types.AddIPToHostRequest) (*types.Task, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/host/%d/ip", hid),
		requestTypePost,
		DefaultService)

	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) HostUpdateIPAutomation(hid int, params *types.IPAutomationType) (*types.Task, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/host/%d/ip_automation", hid),
		requestTypePost,
		DefaultService)

	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) HostIPV4Info(hid int) (*types.IPV4Info, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/host/%d/ipv4", hid),
		requestTypeGet,
		DefaultService)
	var h *types.IPV4Info
	json.Unmarshal(bodyResp, &h)
	return h, err
}

func (a *Api) HostIPV6Info(hid int) (*types.IPV6Info, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/host/%d/ipv6", hid),
		requestTypeGet,
		DefaultService)
	var h *types.IPV6Info
	json.Unmarshal(bodyResp, &h)
	return h, err
}

func (a *Api) HostDisconnectISOWithoutReinstall(hid int) (*types.Task, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/host/%d/iso/cancel", hid),
		requestTypePost,
		DefaultService)
	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) HostDisconnectISOWithReinstall(hid int) (*types.Task, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/host/%d/iso/finish", hid),
		requestTypePost,
		DefaultService)
	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) HostConnectISO(hid int, params *types.ISOTags) (*types.Task, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/host/%d/iso/mount", hid),
		requestTypePost,
		DefaultService)

	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) HostLXDConsole(hid int) (*types.Socket, error) {

	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/host/%d/lxd/console", hid),
		requestTypePost,
		DefaultService)

	var t *types.Socket
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) HostMetrics(hid int) ([]byte, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/host/%d/metrics", hid),
		requestTypeGet,
		DefaultService)

	return bodyResp, err
}

func (a *Api) HostMigrateForm(hid int) (*types.HostMigrateResponse, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/host/%d/migrate", hid),
		requestTypeGet,
		DefaultService)
	var h *types.HostMigrateResponse
	json.Unmarshal(bodyResp, &h)
	return h, err
}

func (a *Api) HostMigrate(hid int, params *types.HostMigrateRequest) (*types.Task, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/host/%d/migrate", hid),
		requestTypePost,
		DefaultService)

	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) HostUpdatePassword(hid int, params *types.Password) (*types.Task, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/host/%d/password", hid),
		requestTypePost,
		DefaultService)

	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) HostNewPTR(hid int, params *types.PTRUpdateRequest) (*types.Task, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/host/%d/ptr", hid),
		requestTypePost,
		DefaultService)

	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) HostPTR(hid int) (*types.PTRResponse, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/host/%d/ptr", hid),
		requestTypeGet,
		DefaultService)
	var h *types.PTRResponse
	json.Unmarshal(bodyResp, &h)
	return h, err
}

func (a *Api) HostUpdatePTR(hid int, domain string, params *types.Domain) (*types.Task, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/host/%d/ptr/%s", hid, domain),
		requestTypePost,
		DefaultService)

	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) HostDeletePTR(hid int, domain string) (*types.Task, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/host/%d/ptr/%s", hid, domain),
		requestTypeDelete,
		DefaultService)

	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) HostReinstall(hid int, params *types.HostReinstallRequest) (*types.RecipeTask, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/host/%d/reinstall", hid),
		requestTypePost,
		DefaultService)

	var t *types.RecipeTask
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) HostRelocate(hid int, params *types.Node) (*types.Task, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/host/%d/relocate", hid),
		requestTypePost,
		DefaultService)

	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) HostRescueMode(hid int, params *types.RescueMode) (*types.Task, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/host/%d/rescue_mode", hid),
		requestTypePost,
		DefaultService)

	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) HostUpdateResource(hid int, params *types.HostResourceUpdateRequest) (*types.Task, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/host/%d/resource", hid),
		requestTypePost,
		DefaultService)

	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) HostRestart(hid int) (*types.Task, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/host/%d/restart", hid),
		requestTypePost,
		DefaultService)

	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) HostRestore(hid int, params *types.HostBackup) (*types.Task, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/host/%d/restore", hid),
		requestTypePost,
		DefaultService)

	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) HostRunRecipe(hid int, params *types.Recipe) (*types.RecipeTask, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/host/%d/runrecipe", hid),
		requestTypePost,
		DefaultService)

	var t *types.RecipeTask
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) HostStart(hid int) (*types.Task, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/host/%d/start", hid),
		requestTypePost,
		DefaultService)

	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) HostStop(hid int) (*types.Task, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/host/%d/stop", hid),
		requestTypePost,
		DefaultService)

	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) HostSyncHAMetadata(hid int) (*types.Task, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/host/%d/sync_ha_metadata", hid),
		requestTypePost,
		DefaultService)

	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) HostVNCSettings(hid int) ([]byte, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/host/%d/vnc_settings", hid),
		requestTypeGet,
		DefaultService)

	return bodyResp, err
}

func (a *Api) HostUpdateVNCSettings(hid int, params *types.Password) (*types.Task, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/host/%d/vnc_settings", hid),
		requestTypePost,
		DefaultService)

	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) HostUpdateVNCPorts(hid int, params *types.VNCPortUpdateRequest) (*types.Task, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/host/%d/vnc_ports", hid),
		requestTypePost,
		DefaultService)

	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) IPs(where *ParamsQuery) (*types.IPList, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/ip%s", where),
		requestTypeGet,
		DefaultService)

	var h *types.IPList
	json.Unmarshal(bodyResp, &h)
	return h, err
}

func (a *Api) IPUpdate(iid int, params *types.UpdateIPResponse) (*types.Task, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/ip/%d", iid),
		requestTypePost,
		DefaultService)

	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) IPDelete(iid int) (*types.Task, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/ip/%d", iid),
		requestTypeDelete,
		DefaultService)

	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) IPAddPTR(hid int, params *types.Domain) (*types.Task, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/ip/%d/ptr", hid),
		requestTypeDelete,
		DefaultService)

	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) HostScheduleList() (*types.ScheduleListResponse, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		"/schedule_list",
		requestTypeGet,
		DefaultService)

	var t *types.ScheduleListResponse
	json.Unmarshal(bodyResp, &t)
	return t, err
}
