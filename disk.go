package vmmanagerapi

import (
	"encoding/json"
	"fmt"

	"github.com/gadost/go-vmmanager-api/types"
)

// Disks get disks info
func (a *Api) Disks() (*types.DiskListResponse, error) {
	bodyResp, err := a.NewRequest(NilPayload, DiskUri, requestTypeGet, DefaultService)

	var d *types.DiskListResponse
	json.Unmarshal(bodyResp, &d)
	return d, err
}

// DiskNew create new disk
func (a *Api) DiskNew(params *types.NewDiskRequest) (*types.Task, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(payload, DiskUri, requestTypePost, DefaultService)

	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// DiskUpdate update disk params
func (a *Api) DiskUpdate(did int, params *types.UpdateDiskRequest) (*types.Task, error) {
	uri := fmt.Sprintf("%s/%d", DiskUri, did)
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(payload, uri, requestTypePost, DefaultService)

	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// DiskDelete delete disk
func (a *Api) DiskDelete(did int) (*types.Task, error) {
	uri := fmt.Sprintf("%s/%d", DiskUri, did)
	bodyResp, err := a.NewRequest(NilPayload, uri, requestTypeDelete, DefaultService)

	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// DiskMigrate migrate disk
func (a *Api) DiskMigrate(did int, params *types.MigrateDiskRequest) (*types.Task, error) {
	uri := fmt.Sprintf("%s/%d/migrate", DiskUri, did)
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(payload, uri, requestTypePost, DefaultService)

	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// DiskLocalStorageForMigrationGet No example outputs in official documentation
func (a *Api) DiskLocalStorageForMigration(did int) ([]byte, error) {
	uri := fmt.Sprintf("%s/%d/migrate/local", DiskUri, did)
	bodyResp, err := a.NewRequest(NilPayload, uri, requestTypeGet, DefaultService)

	return bodyResp, err
}

// DiskNetworkForMigrationGet No example outputs in official documentation
func (a *Api) DiskNetworkForMigration(did int) ([]byte, error) {
	uri := fmt.Sprintf("%s/%d/migrate/network", DiskUri, did)
	bodyResp, err := a.NewRequest(NilPayload, uri, requestTypeGet, DefaultService)

	return bodyResp, err
}

// DiskResore restore disk
func (a *Api) DiskResore(did int, bid int) (*types.Task, error) {
	payload, _ := json.Marshal(&types.BackupIDRequest{Backup: bid})
	uri := fmt.Sprintf("%s/%d/restore", DiskUri, did)
	bodyResp, err := a.NewRequest(payload, uri, requestTypePost, DefaultService)

	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// HostDiskBootOrderUpdate change host's disk boot order
func (a *Api) HostDiskBootOrderUpdate(hid int, params *types.UpdateBootOrderRequest) (*types.Task, error) {
	payload, _ := json.Marshal(params)
	uri := fmt.Sprintf("%s/%d/disk_order", HostUri, hid)
	bodyResp, err := a.NewRequest(payload, uri, requestTypePost, DefaultService)

	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// HostDiskLink link disk to host
func (a *Api) HostDiskLink(hid int, did int, params *types.DiskParams) (*types.Task, error) {
	payload, _ := json.Marshal(params)
	uri := fmt.Sprintf("%s/%d/disk/%d", HostUri, hid, did)
	bodyResp, err := a.NewRequest(payload, uri, requestTypePost, DefaultService)

	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// HostDiskUnlink unlink disk from host
func (a *Api) HostDiskUnlink(hid int, did int) (*types.Task, error) {
	uri := fmt.Sprintf("%s/%d/disk/%d", HostUri, hid, did)
	bodyResp, err := a.NewRequest(NilPayload, uri, requestTypeDelete, DefaultService)

	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}
