package vmmanagerapi

import (
	"encoding/json"
	"fmt"

	"github.com/gadost/go-vmmanager-api/types"
)

func (a *Api) BackupList() (*types.BackupList, error) {
	bodyResp, err := a.NewRequest([]byte(NilPayload), BackupUri, requestTypeGet, DefaultService)
	var b *types.BackupList
	json.Unmarshal(bodyResp, &b)
	return b, err
}

func (a *Api) BackupGet(id int) (*types.Backup, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		BackupUri+fmt.Sprintf("/%d", id),
		requestTypeGet,
		DefaultService)
	var b *types.Backup
	json.Unmarshal(bodyResp, &b)
	return b, err
}

func (a *Api) BackupUpdate(id int, name string, comment string) (*types.Task, error) {
	payload, _ := json.Marshal(&types.ChangeBackup{
		Name:    name,
		Comment: comment,
	})
	bodyResp, err := a.NewRequest(
		payload,
		BackupUri+fmt.Sprintf("/%d", id),
		requestTypePost,
		DefaultService)
	var i *types.Task
	json.Unmarshal(bodyResp, &i)
	return i, err
}

func (a *Api) BackupDelete(id int) (*types.Task, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		BackupUri+fmt.Sprintf("/%d", id),
		requestTypeDelete,
		DefaultService)
	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) BackupCopy(id int, node int) (*types.Task, error) {
	payload, _ := json.Marshal(&types.Node{
		Node: node,
	})
	bodyResp, err := a.NewRequest(
		payload,
		BackupUri+fmt.Sprintf("/%d/copy", id),
		requestTypePost,
		DefaultService)
	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) BackupMove(id int, src int, dst int) (*types.Task, error) {
	payload, _ := json.Marshal(&types.MoveBackup{
		Source:      src,
		Destination: dst,
	})
	bodyResp, err := a.NewRequest(
		payload,
		BackupUri+fmt.Sprintf("/%d/move", id),
		requestTypePost,
		DefaultService)
	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) BackupRelocate(id int, dst int) (*types.Task, error) {
	payload, _ := json.Marshal(&types.RelocateBackup{
		Destination: dst,
	})
	bodyResp, err := a.NewRequest(
		payload,
		BackupUri+fmt.Sprintf("/%d/relocate", id),
		requestTypePost,
		DefaultService)
	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) BackupStateUpdate(id int, state string) error {
	payload, _ := json.Marshal(&types.State{
		State: state,
	})
	_, err := a.NewRequest(
		payload,
		BackupUri+fmt.Sprintf("/%d/state", id),
		requestTypePost,
		DefaultService)

	return err
}

func (a *Api) BackupDiskNew(did int, dparams *types.DiskBackup) (*types.RelocateTask, error) {
	payload, _ := json.Marshal(dparams)
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/disk/%d/backup", did),
		requestTypePost,
		DefaultService)
	var t *types.RelocateTask
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) BackupByDiskGet(did int) (*types.BackupList, error) {

	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/disk/%d/backup", did),
		requestTypeGet,
		DefaultService)
	var b *types.BackupList
	json.Unmarshal(bodyResp, &b)
	return b, err
}

func (a *Api) BackupDiskDelete(did int, id int) (*types.Task, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/disk/%d/backup/%d", did, id),
		requestTypeDelete,
		DefaultService)
	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) BackupVMNew(hid int, params *types.VMBackupParams) (*types.RelocateTask, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/host/%d/backup", hid),
		requestTypePost,
		DefaultService)
	var t *types.RelocateTask
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) BackupHostGet(hid int) (*types.BackupByHostIDResponse, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/host/%d/backup", hid),
		requestTypeGet,
		DefaultService)
	var b *types.BackupByHostIDResponse
	json.Unmarshal(bodyResp, &b)
	return b, err
}

func (a *Api) BackupOutdatedGet() (*types.BackupList, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		"/outdated/disk/backup",
		requestTypeGet,
		DefaultService)
	var b *types.BackupList
	json.Unmarshal(bodyResp, &b)
	return b, err
}

func (a *Api) BackupLocationList() (*types.BackupLocationResponse, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		"/backup_location",
		requestTypeGet,
		DefaultService)
	var b *types.BackupLocationResponse
	json.Unmarshal(bodyResp, &b)
	return b, err
}

func (a *Api) BackupLocationNew(params *types.BackupLocationParams) (*types.Task, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		"/backup_location",
		requestTypePost,
		DefaultService)
	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) BackupLocationUpdate(lid int, params *types.BackupLocationParams) (*types.Task, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/backup_location/%d", lid),
		requestTypePost,
		DefaultService)
	var t *types.Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) BackupLocationDelete(lid int) (*types.DeletedResponse, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/backup_location/%d", lid),
		requestTypeDelete,
		DefaultService)
	var d *types.DeletedResponse
	json.Unmarshal(bodyResp, &d)
	return d, err
}
