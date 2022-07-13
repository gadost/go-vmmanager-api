package vmmanagerapi

import (
	"encoding/json"
	"fmt"

	"github.com/gadost/go-vmmanager-api/types"
)

func (a *Api) DiskList() (*types.DiskListResponse, error) {
	bodyResp, err := a.NewRequest([]byte(NilPayload), "/disk", requestTypeGet, DefaultService)
	var d *types.DiskListResponse
	json.Unmarshal(bodyResp, &d)
	return d, err
}

func (a *Api) DiskNew(params *types.NewDiskRequest) (*types.Task, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		"/disk",
		requestTypePost,
		DefaultService)
	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) DiskUpdate(did int, params *types.UpdateDiskRequest) (*types.Task, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/disk/%d", did),
		requestTypePost,
		DefaultService)
	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) DiskDelete(did int) (*types.Task, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/disk/%d", did),
		requestTypeDelete,
		DefaultService)
	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) DiskMigrate(did int, params *types.MigrateDiskRequest) (*types.Task, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/disk/%d/migrate", did),
		requestTypePost,
		DefaultService)
	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// DiskLocalStorageForMigrationGet No example outputs in official documentation
func (a *Api) DiskLocalStorageForMigrationGet(did int) ([]byte, error) {
	bodyResp, err := a.NewRequest([]byte(NilPayload),
		fmt.Sprintf("/disk/%d/migrate/local", did),
		requestTypeGet,
		DefaultService)

	return bodyResp, err
}

// DiskNetworkForMigrationGet No example outputs in official documentation
func (a *Api) DiskNetworkForMigrationGet(did int) ([]byte, error) {
	bodyResp, err := a.NewRequest([]byte(NilPayload),
		fmt.Sprintf("/disk/%d/migrate/network", did),
		requestTypeGet,
		DefaultService)

	return bodyResp, err
}

func (a *Api) DiskResore(did int, bid int) (*types.Task, error) {
	payload, _ := json.Marshal(&types.BackupIDRequest{
		Backup: bid,
	})
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/disk/%d/restore", did),
		requestTypePost,
		DefaultService)
	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) DiskBootOrderUpdate(hid int, params *types.UpdateBootOrderRequest) (*types.Task, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/host/%d/disk_order", hid),
		requestTypePost,
		DefaultService)
	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) DiskLink(hid int, did int, params *types.DiskParams) (*types.Task, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/host/%d/disk/%d", hid, did),
		requestTypePost,
		DefaultService)
	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) DiskUnlink(hid int, did int) (*types.Task, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/host/%d/disk/%d", hid, did),
		requestTypeDelete,
		DefaultService)
	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}
