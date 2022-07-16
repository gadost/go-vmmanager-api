package vmmanagerapi

import (
	"encoding/json"
	"fmt"

	"github.com/gadost/go-vmmanager-api/types"
)

// Hosts get hosts info
func (a *Api) Hosts(where *ParamsQuery) (*types.HostsResponse, error) {
	uri := fmt.Sprintf("/host%s", where)
	bodyResp, err := a.NewRequest(NilPayload, uri, requestTypeGet, DefaultService)

	var h *types.HostsResponse
	json.Unmarshal(bodyResp, &h)
	return h, err
}

// HostNew create new host
func (a *Api) HostNew(hid int, params *types.NewHostRequest) (*types.RecipeTask, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(payload, HostUri, requestTypePost, DefaultService)

	var t *types.RecipeTask
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// HostByID returns host params
func (a *Api) HostByID(hid int) (*types.HostResponse, error) {
	uri := fmt.Sprintf("/host/%d", hid)
	bodyResp, err := a.NewRequest(NilPayload, uri, requestTypeGet, DefaultService)

	var h *types.HostResponse
	json.Unmarshal(bodyResp, &h)
	return h, err
}

// HostUpdate update host params
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

// HostDelete delete host
func (a *Api) HostDelete(hid int) (*types.Task, error) {
	bodyResp, err := a.NewRequest(
		NilPayload,
		fmt.Sprintf("/host/%d", hid),
		requestTypeDelete,
		DefaultService)
	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// HostChangeOwner change host owner
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

// HostImageSizeLimit set image size limit
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

// HostClone clone host
func (a *Api) HostClone(hid int) (*types.Task, error) {
	bodyResp, err := a.NewRequest(
		NilPayload,
		fmt.Sprintf("/host/%d/clone", hid),
		requestTypePost,
		DefaultService)
	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// HostRepairGuestAgent repair guest agent
func (a *Api) HostRepairGuestAgent(hid int) (*types.Task, error) {
	bodyResp, err := a.NewRequest(
		NilPayload,
		fmt.Sprintf("/host/%d/guest_agent", hid),
		requestTypePost,
		DefaultService)
	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// HostHistory returns host history
func (a *Api) HostHistory(hid int) (*types.HostHistoryResponse, error) {
	bodyResp, err := a.NewRequest(
		NilPayload,
		fmt.Sprintf("/host/%d/history", hid),
		requestTypeGet,
		DefaultService)
	var h *types.HostHistoryResponse
	json.Unmarshal(bodyResp, &h)
	return h, err
}

// HostNewLogEntry create new log entry
func (a *Api) HostNewLogEntry(hid int, payload []byte) ([]byte, error) {
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/host/%d/histroy", hid),
		requestTypePost,
		DefaultService)

	return bodyResp, err
}

// HostInteface get host interfaces
func (a *Api) HostInteface(hid int) (*types.IFace, error) {
	bodyResp, err := a.NewRequest(
		NilPayload,
		fmt.Sprintf("/host/%d/iface", hid),
		requestTypeGet,
		DefaultService)
	var h *types.IFace
	json.Unmarshal(bodyResp, &h)
	return h, err
}

// HostIntefaceNew create new interface
func (a *Api) HostIntefaceNew(hid int, params *types.IFaceParams) ([]byte, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/host/%d/iface", hid),
		requestTypePost,
		DefaultService)

	return bodyResp, err
}

// HostIntefaceUpdate update host interface params
func (a *Api) HostIntefaceUpdate(hid int, iface string, params *types.IFaceModel) (*types.Task, error) {
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

// HostIntefaceDelete delete host interface
func (a *Api) HostIntefaceDelete(hid int, iface string) (*types.Task, error) {
	bodyResp, err := a.NewRequest(
		NilPayload,
		fmt.Sprintf("/host/%d/iface/%s", hid, iface),
		requestTypeDelete,
		DefaultService)

	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// HostIPNew add ip to host
func (a *Api) HostIPNew(hid int, params *types.AddIPToHostRequest) (*types.Task, error) {
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

// HostIPAutomationUpdate update IP automation on host
func (a *Api) HostIPAutomationUpdate(hid int, params *types.IPAutomationType) (*types.Task, error) {
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

// HostIPV4 return IPV4 host info
func (a *Api) HostIPV4(hid int) (*types.IPV4Info, error) {
	bodyResp, err := a.NewRequest(
		NilPayload,
		fmt.Sprintf("/host/%d/ipv4", hid),
		requestTypeGet,
		DefaultService)
	var h *types.IPV4Info
	json.Unmarshal(bodyResp, &h)
	return h, err
}

// HostIPV6 return IPV6 host info
func (a *Api) HostIPV6(hid int) (*types.IPV6Info, error) {
	bodyResp, err := a.NewRequest(
		NilPayload,
		fmt.Sprintf("/host/%d/ipv6", hid),
		requestTypeGet,
		DefaultService)
	var h *types.IPV6Info
	json.Unmarshal(bodyResp, &h)
	return h, err
}

// HostISODisconnectWithoutReinstall disconnect ISO without reinstall
func (a *Api) HostISODisconnectWithoutReinstall(hid int) (*types.Task, error) {
	bodyResp, err := a.NewRequest(
		NilPayload,
		fmt.Sprintf("/host/%d/iso/cancel", hid),
		requestTypePost,
		DefaultService)
	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// HostISODisconnectWithReinstall disconnect ISO with reinstall
func (a *Api) HostISODisconnectWithReinstall(hid int) (*types.Task, error) {
	bodyResp, err := a.NewRequest(
		NilPayload,
		fmt.Sprintf("/host/%d/iso/finish", hid),
		requestTypePost,
		DefaultService)
	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// HostISOConnect connect ISO to host
func (a *Api) HostISOConnect(hid int, params *types.ISOTags) (*types.Task, error) {
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

// HostLXDConsole return LXD console socket
func (a *Api) HostLXDConsole(hid int) (*types.Socket, error) {

	bodyResp, err := a.NewRequest(
		NilPayload,
		fmt.Sprintf("/host/%d/lxd/console", hid),
		requestTypePost,
		DefaultService)

	var t *types.Socket
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// HostMetrics returns host metrics
func (a *Api) HostMetrics(hid int) ([]byte, error) {
	bodyResp, err := a.NewRequest(
		NilPayload,
		fmt.Sprintf("/host/%d/metrics", hid),
		requestTypeGet,
		DefaultService)

	return bodyResp, err
}

// HostMigrateForm returns migrate form
func (a *Api) HostMigrateForm(hid int) (*types.HostMigrateResponse, error) {
	bodyResp, err := a.NewRequest(
		NilPayload,
		fmt.Sprintf("/host/%d/migrate", hid),
		requestTypeGet,
		DefaultService)
	var h *types.HostMigrateResponse
	json.Unmarshal(bodyResp, &h)
	return h, err
}

// HostMigrate start host migration
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

// HostPasswordUpdate update host password
func (a *Api) HostPasswordUpdate(hid int, params *types.Password) (*types.Task, error) {
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

// HostPTRNew add PTR record
func (a *Api) HostPTRNew(hid int, params *types.PTRUpdateRequest) (*types.Task, error) {
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

// HostPTR return PTRs
func (a *Api) HostPTR(hid int) (*types.PTRResponse, error) {
	bodyResp, err := a.NewRequest(
		NilPayload,
		fmt.Sprintf("/host/%d/ptr", hid),
		requestTypeGet,
		DefaultService)
	var h *types.PTRResponse
	json.Unmarshal(bodyResp, &h)
	return h, err
}

// HostPTRUpdate update host PTR record
func (a *Api) HostPTRUpdate(hid int, domain string, params *types.Domain) (*types.Task, error) {
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

// HostPTRDelete delete PTR record
func (a *Api) HostPTRDelete(hid int, domain string) (*types.Task, error) {
	bodyResp, err := a.NewRequest(
		NilPayload,
		fmt.Sprintf("/host/%d/ptr/%s", hid, domain),
		requestTypeDelete,
		DefaultService)

	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// HostReinstall reinstall host
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

// HostRelocate relocate host
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

// HostRescueMode start rescue mode
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

// HostResourceUpdate update host resource
func (a *Api) HostResourceUpdate(hid int, params *types.HostResourceUpdateRequest) (*types.Task, error) {
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

// HostRestart restart host
func (a *Api) HostRestart(hid int) (*types.Task, error) {
	bodyResp, err := a.NewRequest(
		NilPayload,
		fmt.Sprintf("/host/%d/restart", hid),
		requestTypePost,
		DefaultService)

	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// HostRestore restore host
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

// HostRecipeRun run recipe on host
func (a *Api) HostRecipeRun(hid int, params *types.Recipe) (*types.RecipeTask, error) {
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

// HostStart host start
func (a *Api) HostStart(hid int) (*types.Task, error) {
	bodyResp, err := a.NewRequest(
		NilPayload,
		fmt.Sprintf("/host/%d/start", hid),
		requestTypePost,
		DefaultService)

	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// HostStop stop host
func (a *Api) HostStop(hid int) (*types.Task, error) {
	bodyResp, err := a.NewRequest(
		NilPayload,
		fmt.Sprintf("/host/%d/stop", hid),
		requestTypePost,
		DefaultService)

	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// HostHAMetadataSync sync host metadata
func (a *Api) HostHAMetadataSync(hid int) (*types.Task, error) {
	bodyResp, err := a.NewRequest(
		NilPayload,
		fmt.Sprintf("/host/%d/sync_ha_metadata", hid),
		requestTypePost,
		DefaultService)

	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// HostVNCSettings returns host VNC settings
func (a *Api) HostVNCSettings(hid int) ([]byte, error) {
	bodyResp, err := a.NewRequest(
		NilPayload,
		fmt.Sprintf("/host/%d/vnc_settings", hid),
		requestTypeGet,
		DefaultService)

	return bodyResp, err
}

// HostVNCSettingsUpdate update host VNC settings
func (a *Api) HostVNCSettingsUpdate(hid int, params *types.Password) (*types.Task, error) {
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

// HostVNCPortsUpdate update host VNC ports
func (a *Api) HostVNCPortsUpdate(hid int, params *types.VNCPortUpdateRequest) (*types.Task, error) {
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

// IPs returns IPs
func (a *Api) IPs(where *ParamsQuery) (*types.IPList, error) {
	bodyResp, err := a.NewRequest(
		NilPayload,
		fmt.Sprintf("/ip%s", where),
		requestTypeGet,
		DefaultService)

	var h *types.IPList
	json.Unmarshal(bodyResp, &h)
	return h, err
}

// IPUpdate update IP params
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

// IPDelete delete IP
func (a *Api) IPDelete(iid int) (*types.Task, error) {
	bodyResp, err := a.NewRequest(
		NilPayload,
		fmt.Sprintf("/ip/%d", iid),
		requestTypeDelete,
		DefaultService)

	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// IPPTRNew add PTR to IP
func (a *Api) IPPTRNew(hid int, params *types.Domain) (*types.Task, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/ip/%d/ptr", hid),
		requestTypePost,
		DefaultService)

	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// HostSchedules returns host schedules
func (a *Api) HostSchedules() (*types.ScheduleListResponse, error) {
	bodyResp, err := a.NewRequest(
		NilPayload,
		"/schedule_list",
		requestTypeGet,
		DefaultService)

	var t *types.ScheduleListResponse
	json.Unmarshal(bodyResp, &t)
	return t, err
}
