package vmmanagerapi

import (
	"encoding/json"
	"fmt"

	"github.com/gadost/go-vmmanager-api/types"
)

// BackupList returns backup list as *types.BackupList
func (a *Api) BackupList() (*types.BackupList, error) {
	uri := BackupUri
	bodyResp, err := a.NewRequest(NilPayload, uri, requestTypeGet, DefaultService)

	var b *types.BackupList
	json.Unmarshal(bodyResp, &b)
	return b, err
}

// Backup return backup as *types.Backup details by ID
func (a *Api) Backup(id int) (*types.Backup, error) {
	uri := fmt.Sprintf("%s/%d", BackupUri, id)
	bodyResp, err := a.NewRequest(NilPayload, uri, requestTypeGet, DefaultService)

	var b *types.Backup
	json.Unmarshal(bodyResp, &b)
	return b, err
}

// BackupUpdate update backup name / comment
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

// BackupDelete delete backup
func (a *Api) BackupDelete(id int) (*types.Task, error) {
	uri := fmt.Sprintf("%s/%d", BackupUri, id)
	bodyResp, err := a.NewRequest(NilPayload, uri, requestTypeDelete, DefaultService)

	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// BackupCopy copy backup
func (a *Api) BackupCopy(id int, node int) (*types.Task, error) {
	payload, _ := json.Marshal(&types.Node{Node: node})
	uri := fmt.Sprintf("%s/%d/copy", BackupUri, id)
	bodyResp, err := a.NewRequest(payload, uri, requestTypePost, DefaultService)

	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// BackupMove move backup
func (a *Api) BackupMove(id int, src int, dst int) (*types.Task, error) {
	payload, _ := json.Marshal(&types.Move{Source: src, Destination: dst})
	uri := fmt.Sprintf("%s/%d/move", BackupUri, id)
	bodyResp, err := a.NewRequest(payload, uri, requestTypePost, DefaultService)

	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// BackupRelocate relocate backup
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

// BackupStateUpdate update backup state
func (a *Api) BackupStateUpdate(id int, state string) ([]byte, error) {
	payload, _ := json.Marshal(&types.State{
		State: state,
	})
	uri := fmt.Sprintf("%s/%d/state", BackupUri, id)
	bodyResp, err := a.NewRequest(payload, uri, requestTypePost, DefaultService)

	return bodyResp, err
}

// DiskBackupNew create new disk backup
func (a *Api) DiskBackupNew(did int, dparams *types.DiskBackup) (*types.RelocateTask, error) {
	payload, _ := json.Marshal(dparams)
	uri := fmt.Sprintf("%s/%d/backup", DiskUri, did)
	bodyResp, err := a.NewRequest(payload, uri, requestTypePost, DefaultService)

	var t *types.RelocateTask
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// DiskBackup returns disk backup params
func (a *Api) DiskBackup(did int) (*types.BackupList, error) {
	uri := fmt.Sprintf("%s/%d/backup", DiskUri, did)
	bodyResp, err := a.NewRequest(NilPayload, uri, requestTypeGet, DefaultService)

	var b *types.BackupList
	json.Unmarshal(bodyResp, &b)
	return b, err
}

// DiskBackupDelete delete disk backup
func (a *Api) DiskBackupDelete(did int, id int) (*types.Task, error) {
	uri := fmt.Sprintf("%s/%d/backup/%d", DiskUri, did, id)
	bodyResp, err := a.NewRequest(NilPayload, uri, requestTypeDelete, DefaultService)

	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// HostBackupNew create new host backup
func (a *Api) HostBackupNew(hid int, params *types.VMBackupParams) (*types.RelocateTask, error) {
	payload, _ := json.Marshal(params)
	uri := fmt.Sprintf("%s/%d/backup", HostUri, hid)
	bodyResp, err := a.NewRequest(payload, uri, requestTypePost, DefaultService)

	var t *types.RelocateTask
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// HostBackup get host backup params
func (a *Api) HostBackup(hid int) (*types.BackupByHostIDResponse, error) {
	uri := fmt.Sprintf("%s/%d/backup", HostUri, hid)
	bodyResp, err := a.NewRequest(NilPayload, uri, requestTypeGet, DefaultService)

	var b *types.BackupByHostIDResponse
	json.Unmarshal(bodyResp, &b)
	return b, err
}

// OutdatedDiskBackup returns outdated disk backups
func (a *Api) OutdatedDiskBackup() (*types.BackupList, error) {
	uri := "/outdated/disk/backup"
	bodyResp, err := a.NewRequest(NilPayload, uri, requestTypeGet, DefaultService)

	var b *types.BackupList
	json.Unmarshal(bodyResp, &b)
	return b, err
}

// BackupLocations returns backup locations
func (a *Api) BackupLocations() (*types.BackupLocationResponse, error) {
	uri := "/backup_location"
	bodyResp, err := a.NewRequest(NilPayload, uri, requestTypeGet, DefaultService)

	var b *types.BackupLocationResponse
	json.Unmarshal(bodyResp, &b)
	return b, err
}

// BackupLocationNew create new backup location
func (a *Api) BackupLocationNew(params *types.BackupLocationParams) (*types.Task, error) {
	payload, _ := json.Marshal(params)
	uri := "/backup_location"
	bodyResp, err := a.NewRequest(payload, uri, requestTypePost, DefaultService)

	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// BackupLocationUpdateг update backup location
func (a *Api) BackupLocationUpdate(lid int, params *types.BackupLocationParams) (*types.Task, error) {
	payload, _ := json.Marshal(params)
	uri := fmt.Sprintf("/backup_location/%d", lid)
	bodyResp, err := a.NewRequest(payload, uri, requestTypePost, DefaultService)

	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

// BackupLocationDelete delete backup location
func (a *Api) BackupLocationDelete(lid int) (*types.DeletedResponse, error) {
	uri := fmt.Sprintf("/backup_location/%d", lid)
	bodyResp, err := a.NewRequest(NilPayload, uri, requestTypeDelete, DefaultService)

	var d *types.DeletedResponse
	json.Unmarshal(bodyResp, &d)
	return d, err
}
