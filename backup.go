package vmmanagerapi

import (
	"encoding/json"
	"fmt"

	"github.com/gadost/go-vmmanager-api/types"
)

func (a *Api) BackupList() (*types.BackupList, error) {
	uri := BackupUri
	bodyResp, err := a.NewRequest(NilPayload, uri, requestTypeGet, DefaultService)

	var b *types.BackupList
	json.Unmarshal(bodyResp, &b)
	return b, err
}

func (a *Api) BackupGet(id int) (*types.Backup, error) {
	uri := fmt.Sprintf("%s/%d", BackupUri, id)
	bodyResp, err := a.NewRequest(NilPayload, uri, requestTypeGet, DefaultService)

	var b *types.Backup
	json.Unmarshal(bodyResp, &b)
	return b, err
}

func (a *Api) BackupUpdate(id int, name string, comment string) (*types.Task, error) {
	payload, _ := json.Marshal(&types.ChangeBackup{
		Name:    name,
		Comment: comment,
	})
	uri := fmt.Sprintf("%s/%d", BackupUri, id)
	bodyResp, err := a.NewRequest(payload, uri, requestTypePost, DefaultService)

	var i *types.Task
	json.Unmarshal(bodyResp, &i)
	return i, err
}

func (a *Api) BackupDelete(id int) (*types.Task, error) {
	uri := fmt.Sprintf("%s/%d", BackupUri, id)
	bodyResp, err := a.NewRequest(NilPayload, uri, requestTypeDelete, DefaultService)

	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) BackupCopy(id int, node int) (*types.Task, error) {
	payload, _ := json.Marshal(&types.Node{Node: node})
	uri := fmt.Sprintf("%s/%d/copy", BackupUri, id)
	bodyResp, err := a.NewRequest(payload, uri, requestTypePost, DefaultService)

	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) BackupMove(id int, src int, dst int) (*types.Task, error) {
	payload, _ := json.Marshal(&types.Move{Source: src, Destination: dst})
	uri := fmt.Sprintf("%s/%d/move", BackupUri, id)
	bodyResp, err := a.NewRequest(payload, uri, requestTypePost, DefaultService)

	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) BackupRelocate(id int, dst int) (*types.Task, error) {
	payload, _ := json.Marshal(&types.RelocateBackup{
		Destination: dst,
	})
	uri := fmt.Sprintf("%s/%d/relocate", BackupUri, id)
	bodyResp, err := a.NewRequest(payload, uri, requestTypePost, DefaultService)

	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) BackupStateUpdate(id int, state string) ([]byte, error) {
	payload, _ := json.Marshal(&types.State{
		State: state,
	})
	uri := fmt.Sprintf("%s/%d/state", BackupUri, id)
	bodyResp, err := a.NewRequest(payload, uri, requestTypePost, DefaultService)

	return bodyResp, err
}

func (a *Api) DiskBackupNew(did int, dparams *types.DiskBackup) (*types.RelocateTask, error) {
	payload, _ := json.Marshal(dparams)
	uri := fmt.Sprintf("%s/%d/backup", DiskUri, did)
	bodyResp, err := a.NewRequest(payload, uri, requestTypePost, DefaultService)

	var t *types.RelocateTask
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) DiskBackupGet(did int) (*types.BackupList, error) {
	uri := fmt.Sprintf("%s/%d/backup", DiskUri, did)
	bodyResp, err := a.NewRequest(NilPayload, uri, requestTypeGet, DefaultService)

	var b *types.BackupList
	json.Unmarshal(bodyResp, &b)
	return b, err
}

func (a *Api) BackupDiskDelete(did int, id int) (*types.Task, error) {
	uri := fmt.Sprintf("%s/%d/backup/%d", DiskUri, did, id)
	bodyResp, err := a.NewRequest(NilPayload, uri, requestTypeDelete, DefaultService)

	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) HostBackupNew(hid int, params *types.VMBackupParams) (*types.RelocateTask, error) {
	payload, _ := json.Marshal(params)
	uri := fmt.Sprintf("%s/%d/backup", HostUri, hid)
	bodyResp, err := a.NewRequest(payload, uri, requestTypePost, DefaultService)

	var t *types.RelocateTask
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) HostBackupGet(hid int) (*types.BackupByHostIDResponse, error) {
	uri := fmt.Sprintf("%s/%d/backup", HostUri, hid)
	bodyResp, err := a.NewRequest(NilPayload, uri, requestTypeGet, DefaultService)

	var b *types.BackupByHostIDResponse
	json.Unmarshal(bodyResp, &b)
	return b, err
}

func (a *Api) OutdatedDiskBackup() (*types.BackupList, error) {
	uri := "/outdated/disk/backup"
	bodyResp, err := a.NewRequest(NilPayload, uri, requestTypeGet, DefaultService)

	var b *types.BackupList
	json.Unmarshal(bodyResp, &b)
	return b, err
}

func (a *Api) BackupLocations() (*types.BackupLocationResponse, error) {
	uri := "/backup_location"
	bodyResp, err := a.NewRequest(NilPayload, uri, requestTypeGet, DefaultService)

	var b *types.BackupLocationResponse
	json.Unmarshal(bodyResp, &b)
	return b, err
}

func (a *Api) BackupLocationNew(params *types.BackupLocationParams) (*types.Task, error) {
	payload, _ := json.Marshal(params)
	uri := "/backup_location"
	bodyResp, err := a.NewRequest(payload, uri, requestTypePost, DefaultService)

	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) BackupLocationUpdate(lid int, params *types.BackupLocationParams) (*types.Task, error) {
	payload, _ := json.Marshal(params)
	uri := fmt.Sprintf("/backup_location/%d", lid)
	bodyResp, err := a.NewRequest(payload, uri, requestTypePost, DefaultService)

	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) BackupLocationDelete(lid int) (*types.DeletedResponse, error) {
	uri := fmt.Sprintf("/backup_location/%d", lid)
	bodyResp, err := a.NewRequest(NilPayload, uri, requestTypeDelete, DefaultService)

	var d *types.DeletedResponse
	json.Unmarshal(bodyResp, &d)
	return d, err
}
