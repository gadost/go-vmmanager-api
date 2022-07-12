package vmmanagerapi

import (
	"encoding/json"
	"fmt"
)

type BackupList struct {
	LastNotify int `json:"last_notify"`
	List       []struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		Account struct {
			ID    int    `json:"id"`
			Email string `json:"email"`
		} `json:"account"`
		Disk struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"disk"`
		ActualSizeMib    int `json:"actual_size_mib"`
		EstimatedSizeMib int `json:"estimated_size_mib"`
		Cluster          struct {
			ID               int    `json:"id"`
			DatacenterType   string `json:"datacenter_type"`
			ImageStoragePath string `json:"image_storage_path"`
			Name             string `json:"name"`
		} `json:"cluster"`
		Host           int    `json:"host"`
		Schedule       int    `json:"schedule"`
		DateCreate     string `json:"date_create"`
		AvailableUntil string `json:"available_until"`
		Comment        string `json:"comment"`
		BackupLocation struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
			Type string `json:"type"`
		} `json:"backup_location"`
	} `json:"list"`
}

type Backup struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Account        string `json:"account"`
	ForAll         bool   `json:"for_all"`
	Type           string `json:"type"`
	Comment        string `json:"comment"`
	State          string `json:"state"`
	ExpandPart     string `json:"expand_part"`
	BackupLocation struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"backup_location"`
}

type ChangeBackup struct {
	Name    string `json:"name"`
	Comment string `json:"comment"`
}

type Task struct {
	ID   int `json:"id"`
	Task int `json:"task"`
}
type RelocateTask struct {
	ID           int `json:"id"`
	Task         int `json:"task"`
	RelocateTask int `json:"relocate_task"`
}

type Node struct {
	Node int `json:"node"`
}

type MoveBackup struct {
	Source      int `json:"source"`
	Destination int `json:"destination"`
}

type RelocateBackup struct {
	Destination int `json:"destination"`
}

type State struct {
	State string `json:"state"`
}

type DiskBackup struct {
	Name            string `json:"name"`
	Comment         string `json:"comment"`
	BackupLocations []int  `json:"backup_locations"`
	Schedule        int    `json:"schedule"`
}

type VMBackupParams struct {
	Name            string `json:"name"`
	Comment         string `json:"comment"`
	BackupLocations []int  `json:"backup_locations"`
	Schedule        int    `json:"schedule"`
}

type BackupByHostIDResponse struct {
	LastNotify int `json:"last_notify"`
	List       []struct {
		ID             int    `json:"id"`
		Name           string `json:"name"`
		Account        string `json:"account"`
		ForAll         bool   `json:"for_all"`
		Type           string `json:"type"`
		Comment        string `json:"comment"`
		State          string `json:"state"`
		ExpandPart     string `json:"expand_part"`
		BackupLocation struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
			Type string `json:"type"`
		} `json:"backup_location"`
	} `json:"list"`
}

type BackupLocationResponse struct {
	LastNotify int `json:"last_notify"`
	List       []struct {
		ID               int    `json:"id"`
		Name             string `json:"name"`
		Comment          string `json:"comment"`
		State            string `json:"state"`
		QuotaMib         int    `json:"quota_mib"`
		Type             string `json:"type"`
		ConnectionParams struct {
		} `json:"connection_params"`
		Clusters []struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"clusters"`
	} `json:"list"`
}

type DeletedBackupResponse struct {
	Deleted []struct {
		ID int `json:"id"`
	} `json:"deleted"`
}

type BackupLocationParams struct {
	Name             string `json:"name"`
	Comment          string `json:"comment"`
	QuotaMib         int    `json:"quota_mib"`
	Type             string `json:"type"`
	ConnectionParams struct {
	} `json:"connection_params"`
	Clusters            []int `json:"clusters"`
	Schedules           []int `json:"schedules"`
	SkipConnectionCheck bool  `json:"skip_connection_check"`
}

func (a *Api) BackupList() (*BackupList, error) {
	bodyResp, err := a.NewRequest([]byte(NilPayload), BackupUri, requestTypeGet, DefaultService)
	var b *BackupList
	json.Unmarshal(bodyResp, &b)
	return b, err
}

func (a *Api) BackupGet(id int) (*Backup, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		BackupUri+fmt.Sprintf("/%d", id),
		requestTypeGet,
		DefaultService)
	var b *Backup
	json.Unmarshal(bodyResp, &b)
	return b, err
}

func (a *Api) BackupUpdate(id int, name string, comment string) (*Task, error) {
	payload, _ := json.Marshal(&ChangeBackup{
		Name:    name,
		Comment: comment,
	})
	bodyResp, err := a.NewRequest(
		payload,
		BackupUri+fmt.Sprintf("/%d", id),
		requestTypePost,
		DefaultService)
	var i *Task
	json.Unmarshal(bodyResp, &i)
	return i, err
}

func (a *Api) BackupDelete(id int) (*Task, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		BackupUri+fmt.Sprintf("/%d", id),
		requestTypeDelete,
		DefaultService)
	var t *Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) BackupCopy(id int, node int) (*Task, error) {
	payload, _ := json.Marshal(&Node{
		Node: node,
	})
	bodyResp, err := a.NewRequest(
		payload,
		BackupUri+fmt.Sprintf("/%d/copy", id),
		requestTypePost,
		DefaultService)
	var t *Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) BackupMove(id int, src int, dst int) (*Task, error) {
	payload, _ := json.Marshal(&MoveBackup{
		Source:      src,
		Destination: dst,
	})
	bodyResp, err := a.NewRequest(
		payload,
		BackupUri+fmt.Sprintf("/%d/move", id),
		requestTypePost,
		DefaultService)
	var t *Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) BackupRelocate(id int, dst int) (*Task, error) {
	payload, _ := json.Marshal(&RelocateBackup{
		Destination: dst,
	})
	bodyResp, err := a.NewRequest(
		payload,
		BackupUri+fmt.Sprintf("/%d/relocate", id),
		requestTypePost,
		DefaultService)
	var t *Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) BackupStateUpdate(id int, state string) error {
	payload, _ := json.Marshal(&State{
		State: state,
	})
	_, err := a.NewRequest(
		payload,
		BackupUri+fmt.Sprintf("/%d/state", id),
		requestTypePost,
		DefaultService)

	return err
}

func (a *Api) BackupDiskNew(did int, dparams *DiskBackup) (*RelocateTask, error) {
	payload, _ := json.Marshal(dparams)
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/disk/%d/backup", did),
		requestTypePost,
		DefaultService)
	var t *RelocateTask
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) BackupByDiskGet(did int) (*BackupList, error) {

	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/disk/%d/backup", did),
		requestTypeGet,
		DefaultService)
	var b *BackupList
	json.Unmarshal(bodyResp, &b)
	return b, err
}

func (a *Api) BackupDiskDelete(did int, id int) (*Task, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/disk/%d/backup/%d", did, id),
		requestTypeDelete,
		DefaultService)
	var t *Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) BackupVMNew(hid int, params *VMBackupParams) (*RelocateTask, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/host/%d/backup", hid),
		requestTypePost,
		DefaultService)
	var t *RelocateTask
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) BackupHostGet(hid int) (*BackupByHostIDResponse, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/host/%d/backup", hid),
		requestTypeGet,
		DefaultService)
	var b *BackupByHostIDResponse
	json.Unmarshal(bodyResp, &b)
	return b, err
}

func (a *Api) BackupOutdatedGet() (*BackupList, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		"/outdated/disk/backup",
		requestTypeGet,
		DefaultService)
	var b *BackupList
	json.Unmarshal(bodyResp, &b)
	return b, err
}

func (a *Api) BackupLocationList() (*BackupLocationResponse, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		"/backup_location",
		requestTypeGet,
		DefaultService)
	var b *BackupLocationResponse
	json.Unmarshal(bodyResp, &b)
	return b, err
}

func (a *Api) BackupLocationNew(params *BackupLocationParams) (*Task, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		"/backup_location",
		requestTypePost,
		DefaultService)
	var t *Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) BackupLocationUpdate(lid int, params *BackupLocationParams) (*Task, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(
		payload,
		fmt.Sprintf("/backup_location/%d", lid),
		requestTypePost,
		DefaultService)
	var t *Task
	json.Unmarshal(bodyResp, &t)
	return t, err
}

func (a *Api) BackupLocationDelete(lid int) (*DeletedBackupResponse, error) {
	bodyResp, err := a.NewRequest(
		[]byte(NilPayload),
		fmt.Sprintf("/backup_location/%d", lid),
		requestTypeDelete,
		DefaultService)
	var d *DeletedBackupResponse
	json.Unmarshal(bodyResp, &d)
	return d, err
}
