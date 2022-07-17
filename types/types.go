package types

type Error struct {
	Code int    `json:"code,omitempty"`
	Msg  string `json:"msg,omitempty"`
}

// Auth

type Password struct {
	Password string `json:"password,omitempty"`
}

type InviteUser struct {
	Password        string          `json:"password,omitempty"`
	Lang            string          `json:"lang,omitempty"`
	AdditionalProp1 AdditionalProp1 `json:"additionalProp1,omitempty"`
}

type AdditionalProp1 struct {
}

type TokenLifetime struct {
	ExpiresAt   string `json:"expires_at,omitempty"`
	Description string `json:"description,omitempty"`
}

//Backups

type BackupList struct {
	LastNotify int                 `json:"last_notify,omitempty"`
	List       []BackupListElement `json:"list,omitempty"`

	// Returns on API related errors
	Error Error `json:"error,omitempty"`
}

type BackupListElement struct {
	ID               int            `json:"id,omitempty"`
	Name             string         `json:"name,omitempty"`
	Account          BaseAccount    `json:"account,omitempty"`
	Disk             Disk           `json:"disk,omitempty"`
	ActualSizeMib    int            `json:"actual_size_mib,omitempty"`
	EstimatedSizeMib int            `json:"estimated_size_mib,omitempty"`
	Cluster          Cluster        `json:"cluster,omitempty"`
	Host             int            `json:"host,omitempty"`
	Schedule         int            `json:"schedule,omitempty"`
	DateCreate       string         `json:"date_create,omitempty"`
	AvailableUntil   string         `json:"available_until,omitempty"`
	Comment          string         `json:"comment,omitempty"`
	BackupLocation   BackupLocation `json:"backup_location,omitempty"`
}

type BaseAccount struct {
	ID    int    `json:"id,omitempty"`
	Email string `json:"email,omitempty"`
}

type Disk struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Cluster struct {
	ID               int    `json:"id,omitempty"`
	DatacenterType   string `json:"datacenter_type,omitempty"`
	ImageStoragePath string `json:"image_storage_path,omitempty"`
	Name             string `json:"name,omitempty"`
}

type BackupLocation struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
}

type Backup struct {
	ID             int            `json:"id,omitempty"`
	Name           string         `json:"name,omitempty"`
	Account        string         `json:"account,omitempty"`
	ForAll         bool           `json:"for_all,omitempty"`
	Type           string         `json:"type,omitempty"`
	Comment        string         `json:"comment,omitempty"`
	State          string         `json:"state,omitempty"`
	ExpandPart     string         `json:"expand_part,omitempty"`
	BackupLocation BackupLocation `json:"backup_location,omitempty"`

	// Returns on API related errors
	Error Error `json:"error,omitempty"`
}

type ChangeBackup struct {
	Name    string `json:"name,omitempty"`
	Comment string `json:"comment,omitempty"`
}

type Task struct {
	ID   int `json:"id,omitempty"`
	Task int `json:"task,omitempty"`

	// Returns on API related errors
	Error Error `json:"error,omitempty"`
}
type RelocateTask struct {
	ID           int `json:"id,omitempty"`
	Task         int `json:"task,omitempty"`
	RelocateTask int `json:"relocate_task,omitempty"`

	// Returns on API related errors
	Error Error `json:"error,omitempty"`
}

type Node struct {
	Node int `json:"node,omitempty"`
}

type Move struct {
	Source      int `json:"source,omitempty"`
	Destination int `json:"destination,omitempty"`
}

type RelocateBackup struct {
	Destination int `json:"destination,omitempty"`
}

type State struct {
	State string `json:"state,omitempty"`
}

type DiskBackup struct {
	Name            string `json:"name,omitempty"`
	Comment         string `json:"comment,omitempty"`
	BackupLocations []int  `json:"backup_locations,omitempty"`
	Schedule        int    `json:"schedule,omitempty"`
}

type VMBackupParams struct {
	Name            string `json:"name,omitempty"`
	Comment         string `json:"comment,omitempty"`
	BackupLocations []int  `json:"backup_locations,omitempty"`
	Schedule        int    `json:"schedule,omitempty"`
}

type BackupByHostIDResponse struct {
	LastNotify int      `json:"last_notify,omitempty"`
	List       []Backup `json:"list,omitempty"`

	// Returns on API related errors
	Error Error `json:"error,omitempty"`
}

type BackupLocationResponse struct {
	LastNotify int                     `json:"last_notify,omitempty"`
	List       []BackupLocationElement `json:"list,omitempty"`

	// Returns on API related errors
	Error Error `json:"error,omitempty"`
}

type BackupLocationElement struct {
	ID               int              `json:"id,omitempty"`
	Name             string           `json:"name,omitempty"`
	Comment          string           `json:"comment,omitempty"`
	State            string           `json:"state,omitempty"`
	QuotaMib         int              `json:"quota_mib,omitempty"`
	Type             string           `json:"type,omitempty"`
	ConnectionParams ConnectionParams `json:"connection_params,omitempty"`
	Clusters         []Cluster        `json:"clusters,omitempty"`
}

type ConnectionParams struct {
}

type ID struct {
	ID int `json:"id,omitempty"`

	// Returns on API related errors
	Error Error `json:"error,omitempty"`
}

type BackupLocationParams struct {
	Name                string           `json:"name,omitempty"`
	Comment             string           `json:"comment,omitempty"`
	QuotaMib            int              `json:"quota_mib,omitempty"`
	Type                string           `json:"type,omitempty"`
	ConnectionParams    ConnectionParams `json:"connection_params,omitempty"`
	Clusters            []int            `json:"clusters,omitempty"`
	Schedules           []int            `json:"schedules,omitempty"`
	SkipConnectionCheck bool             `json:"skip_connection_check,omitempty"`
}

// Cluster

type ClusterListResponse struct {
	LastNotify int              `json:"last_notify,omitempty"`
	List       []ClusterElement `json:"list,omitempty"`

	// Returns on API related errors
	Error Error `json:"error,omitempty"`
}

type ClusterElement struct {
	VirtualizationType     string           `json:"virtualization_type,omitempty"`
	OsStoragePath          string           `json:"os_storage_path,omitempty"`
	ImageStoragePath       string           `json:"image_storage_path,omitempty"`
	CPUNumber              CPUNumber        `json:"cpu_number,omitempty"`
	HostCount              int              `json:"host_count,omitempty"`
	AccountHostCount       int              `json:"account_host_count,omitempty"`
	HostPerNodeLimit       int              `json:"host_per_node_limit,omitempty"`
	Name                   string           `json:"name,omitempty"`
	Comment                string           `json:"comment,omitempty"`
	RAMMib                 RAMMib           `json:"ram_mib,omitempty"`
	StorageMib             StorageMib       `json:"storage_mib,omitempty"`
	TimeZone               string           `json:"time_zone,omitempty"`
	State                  string           `json:"state,omitempty"`
	DataCenter             string           `json:"data_center,omitempty"`
	HostDistributionPolicy string           `json:"host_distribution_policy,omitempty"`
	HostFilter             []HostFilter     `json:"host_filter,omitempty"`
	NodeCount              int              `json:"node_count,omitempty"`
	ID                     int              `json:"id,omitempty"`
	Overselling            int              `json:"overselling,omitempty"`
	BackupLocations        []BackupLocation `json:"backup_locations,omitempty"`
	DomainTemplate         string           `json:"domain_template,omitempty"`
	DatacenterParams       DatacenterParams `json:"datacenter_params,omitempty"`
	Storages               []Storage        `json:"storages,omitempty"`
	Storage                Storage          `json:"storage,omitempty"`
	DNSServers             []string         `json:"dns_servers,omitempty"`
	DcNetworks             []int            `json:"dc_networks,omitempty"`
	NodeNetwork            NodeNetwork      `json:"node_network,omitempty"`
	VxlanMode              string           `json:"vxlan_mode,omitempty"`
	VxlanSettings          string           `json:"vxlan_settings,omitempty"`
	Vxlans                 []Vxlan          `json:"vxlans,omitempty"`
	DatacenterType         string           `json:"datacenter_type,omitempty"`
}

type Vxlan struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Tag  int    `json:"tag,omitempty"`
}

type NodeNetwork struct {
	Timeout int    `json:"timeout,omitempty"`
	Gateway string `json:"gateway,omitempty"`
}

type CPUNumber struct {
	Total int `json:"total,omitempty"`
	Used  int `json:"used,omitempty"`
}

type RAMMib struct {
	Total int `json:"total,omitempty"`
	Used  int `json:"used,omitempty"`
}

type StorageMib struct {
	Total int `json:"total,omitempty"`
	Used  int `json:"used,omitempty"`
}

type HostFilter struct {
	Entity     string `json:"entity,omitempty"`
	Expression string `json:"expression,omitempty"`
}

type DatacenterParams struct {
	Gateway        string       `json:"gateway,omitempty"`
	Mac            string       `json:"mac,omitempty"`
	BgpCommunity   string       `json:"bgp_community,omitempty"`
	BgpAs          int          `json:"bgp_as,omitempty"`
	BgpSessions    []BgpSession `json:"bgp_sessions,omitempty"`
	BgpCommunityV6 string       `json:"bgp_community_v6,omitempty"`
	BgpAsV6        int          `json:"bgp_as_v6,omitempty"`
	BgpSessionsV6  []BgpSession `json:"bgp_sessions_v6,omitempty"`
}

type BgpSession struct {
	IP      string `json:"ip,omitempty"`
	BgpAs   int    `json:"bgp_as,omitempty"`
	Comment string `json:"comment,omitempty"`
}

type Storage struct {
	ID       int      `json:"id,omitempty"`
	Name     string   `json:"name,omitempty"`
	Type     string   `json:"type,omitempty"`
	Enabled  bool     `json:"enabled,omitempty"`
	IsMain   bool     `json:"is_main,omitempty"`
	VirtPool VirtPool `json:"virt_pool,omitempty"`
	Tags     []int    `json:"tags,omitempty"`
}

type VirtPool struct {
	VMStoragePath   string `json:"vm_storage_path,omitempty"`
	VolumeGroup     string `json:"volume_group,omitempty"`
	ZfsPoolName     string `json:"zfs_pool_name,omitempty"`
	RbdPoolName     string `json:"rbd_pool_name,omitempty"`
	RbdUser         string `json:"rbd_user,omitempty"`
	PlacementGroups int    `json:"placement_groups,omitempty"`
	NetworkDiskPath string `json:"network_disk_path,omitempty"`
}

type ClusterResponse struct {
	VirtualizationType     string           `json:"virtualization_type,omitempty"`
	OsStoragePath          string           `json:"os_storage_path,omitempty"`
	ImageStoragePath       string           `json:"image_storage_path,omitempty"`
	CPUNumber              CPUNumber        `json:"cpu_number,omitempty"`
	HostCount              int              `json:"host_count,omitempty"`
	AccountHostCount       int              `json:"account_host_count,omitempty"`
	HostPerNodeLimit       int              `json:"host_per_node_limit,omitempty"`
	Name                   string           `json:"name,omitempty"`
	Comment                string           `json:"comment,omitempty"`
	RAMMib                 RAMMib           `json:"ram_mib,omitempty"`
	StorageMib             StorageMib       `json:"storage_mib,omitempty"`
	TimeZone               string           `json:"time_zone,omitempty"`
	State                  string           `json:"state,omitempty"`
	DataCenter             string           `json:"data_center,omitempty"`
	HostDistributionPolicy string           `json:"host_distribution_policy,omitempty"`
	HostFilter             []HostFilter     `json:"host_filter,omitempty"`
	NodeCount              int              `json:"node_count,omitempty"`
	ID                     int              `json:"id,omitempty"`
	Overselling            int              `json:"overselling,omitempty"`
	BackupLocations        []BackupLocation `json:"backup_locations,omitempty"`
	DomainTemplate         string           `json:"domain_template,omitempty"`
	DatacenterParams       DatacenterParams `json:"datacenter_params,omitempty"`
	Storages               []Storage        `json:"storages,omitempty"`
	Storage                Storage          `json:"storage,omitempty"`
	DNSServers             []string         `json:"dns_servers,omitempty"`
	DcNetworks             []int            `json:"dc_networks,omitempty"`
	NodeNetwork            NodeNetwork      `json:"node_network,omitempty"`
	VxlanMode              string           `json:"vxlan_mode,omitempty"`
	VxlanSettings          string           `json:"vxlan_settings,omitempty"`
	Vxlans                 []Vxlan          `json:"vxlans,omitempty"`

	// Returns on API related errors
	Error Error `json:"error,omitempty"`
}

type CreateClusterRequest struct {
	VirtualizationType     string           `json:"virtualization_type,omitempty"`
	DatacenterType         string           `json:"datacenter_type,omitempty"`
	Name                   string           `json:"name,omitempty"`
	Comment                string           `json:"comment,omitempty"`
	IsoEnabled             bool             `json:"iso_enabled,omitempty"`
	ManageDiskEnabled      bool             `json:"manage_disk_enabled,omitempty"`
	TimeZone               string           `json:"time_zone,omitempty"`
	Interfaces             []Interface      `json:"interfaces,omitempty"`
	Os                     []int            `json:"os,omitempty"`
	BackupLocations        []int            `json:"backup_locations,omitempty"`
	OsStoragePath          string           `json:"os_storage_path,omitempty"`
	ImageStoragePath       string           `json:"image_storage_path,omitempty"`
	Overselling            int              `json:"overselling,omitempty"`
	HostDistributionPolicy string           `json:"host_distribution_policy,omitempty"`
	HostFilter             []HostFilter     `json:"host_filter,omitempty"`
	DomainTemplate         string           `json:"domain_template,omitempty"`
	DatacenterParams       DatacenterParams `json:"datacenter_params,omitempty"`
	DomainChangeAllowed    bool             `json:"domain_change_allowed,omitempty"`
	ProxyEnabled           bool             `json:"proxy_enabled,omitempty"`
	HostPerNodeLimit       int              `json:"host_per_node_limit,omitempty"`
	DNSServers             []string         `json:"dns_servers,omitempty"`
	NodeNetwork            NodeNetwork      `json:"node_network,omitempty"`
	VxlanMode              string           `json:"vxlan_mode,omitempty"`
	VxlanSettings          VxlanSettings    `json:"vxlan_settings,omitempty"`
}

type VxlanSettings struct {
	BgpCommunity   string       `json:"bgp_community,omitempty"`
	BgpAs          int          `json:"bgp_as,omitempty"`
	BgpSessions    []BgpSession `json:"bgp_sessions,omitempty"`
	BgpCommunityV6 string       `json:"bgp_community_v6,omitempty"`
	BgpAsV6        int          `json:"bgp_as_v6,omitempty"`
	BgpSessionsV6  []BgpSession `json:"bgp_sessions_v6,omitempty"`
}

type Interface struct {
	Interface int   `json:"interface,omitempty"`
	Ippool    []int `json:"ippool,omitempty"`
}

type Tasks struct {
	ID    int   `json:"id,omitempty"`
	Tasks []int `json:"tasks,omitempty"`
}

type DeletedResponse struct {
	Deleted []ID `json:"deleted,omitempty"`

	// Returns on API related errors
	Error Error `json:"error,omitempty"`
}

type DCNetworksResponse struct {
	DcNetworks []int `json:"dc_networks,omitempty"`

	// Returns on API related errors
	Error Error `json:"error,omitempty"`
}

type QemuVersion struct {
	QemuVersion string `json:"qemu_version,omitempty"`
}

type IPPoolRequest struct {
	Interfaces []Interface `json:"interfaces,omitempty"`
}

type Name struct {
	Name string `json:"name,omitempty"`
}

type SSHKeysRequest struct {
	SSHKeys []int `json:"ssh_keys,omitempty"`
}

type AttachStorageToClusterRequest struct {
	VirtPool       VirtPool `json:"virt_pool,omitempty"`
	IsMain         bool     `json:"is_main,omitempty"`
	HddOverselling int      `json:"hdd_overselling,omitempty"`
	IgnoreChecks   bool     `json:"ignore_checks,omitempty"`
}

type StoragesTasks struct {
	ID            int   `json:"id,omitempty"`
	Task          int   `json:"task,omitempty"`
	StoragesTasks []int `json:"storages_tasks,omitempty"`

	// Returns on API related errors
	Error Error `json:"error,omitempty"`
}

type StorageTask struct {
	ID            int `json:"id,omitempty"`
	Task          int `json:"task,omitempty"`
	StoragesTasks int `json:"storages_tasks,omitempty"`

	// Returns on API related errors
	Error Error `json:"error,omitempty"`
}

type HDDOversellingRequest struct {
	Value float32 `json:"value,omitempty"`
}

type HDDOversellingResponse struct {
	Value          float32 `json:"value,omitempty"`
	HddOverselling float32 `json:"hdd_overselling,omitempty"`

	// Returns on API related errors
	Error Error `json:"error,omitempty"`
}

type VXLanResponse struct {
	LastNotify int            `json:"last_notify,omitempty"`
	List       []VxlanElement `json:"list,omitempty"`

	// Returns on API related errors
	Error Error `json:"error,omitempty"`
}

type VxlanElement struct {
	Account BaseAccount `json:"account,omitempty"`
	ID      int         `json:"id,omitempty"`
	Name    string      `json:"name,omitempty"`
	Tag     int         `json:"tag,omitempty"`
	Hosts   []Host      `json:"hosts,omitempty"`
	Ranges  []string    `json:"ranges,omitempty"`
}

type Host struct {
	ID   int      `json:"id,omitempty"`
	Name string   `json:"name,omitempty"`
	Node NodeHost `json:"node,omitempty"`
}

type NodeHost struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type NodesVXLanResponse struct {
	LastNotify int          `json:"last_notify,omitempty"`
	List       []NodesVxlan `json:"list,omitempty"`

	// Returns on API related errors
	Error Error `json:"error,omitempty"`
}

type HostVxlan struct {
	ID     int      `json:"id,omitempty"`
	Name   string   `json:"name,omitempty"`
	Ranges []string `json:"ranges,omitempty"`
}

type NodesVxlan struct {
	Node   NodeHost            `json:"node,omitempty"`
	Vxlans []NodesVxlanElement `json:"vxlans,omitempty"`
}

type NodesVxlanElement struct {
	ID    int         `json:"id,omitempty"`
	Name  string      `json:"name,omitempty"`
	Tag   int         `json:"tag,omitempty"`
	Hosts []HostVxlan `json:"hosts,omitempty"`
}

type UsersVXLanResponse struct {
	LastNotify int                `json:"last_notify,omitempty"`
	List       []UserVxlanElement `json:"list,omitempty"`

	// Returns on API related errors
	Error Error `json:"error,omitempty"`
}

type UserVxlanElement struct {
	Account BaseAccount `json:"account,omitempty"`
	Hosts   []UserHost  `json:"hosts,omitempty"`
}

type UserHost struct {
	ID     int         `json:"id,omitempty"`
	Name   string      `json:"name,omitempty"`
	Node   NodeHost    `json:"node,omitempty"`
	Vxlans []UserVxlan `json:"vxlans,omitempty"`
}

type UserVxlan struct {
	ID     int      `json:"id,omitempty"`
	Name   string   `json:"name,omitempty"`
	Tag    int      `json:"tag,omitempty"`
	Ranges []string `json:"ranges,omitempty"`
}

type SSHKeyResponse struct {
	LastNotify int      `json:"last_notify,omitempty"`
	List       []SSHKey `json:"list,omitempty"`

	// Returns on API related errors
	Error Error `json:"error,omitempty"`
}

type SSHKey struct {
	ID       int    `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	SSHKey   string `json:"ssh_key,omitempty"`
	Clusters []int  `json:"clusters,omitempty"`
}

type SSHKeyRequest struct {
	Name     string `json:"name,omitempty"`
	SSHKey   string `json:"ssh_key,omitempty"`
	Clusters []int  `json:"clusters,omitempty"`
}

// Disk

type DiskListResponse struct {
	LastNotify int           `json:"last_notify,omitempty"`
	List       []DiskElement `json:"list,omitempty"`

	// Returns on API related errors
	Error Error `json:"error,omitempty"`
}

type DiskElement struct {
	ID         int         `json:"id,omitempty"`
	Bus        string      `json:"bus,omitempty"`
	Name       string      `json:"name,omitempty"`
	ExpandPart string      `json:"expand_part,omitempty"`
	TargetDev  string      `json:"target_dev,omitempty"`
	SizeMib    int         `json:"size_mib,omitempty"`
	SizeMibNew int         `json:"size_mib_new,omitempty"`
	Account    BaseAccount `json:"account,omitempty"`
	Storage    StorageID   `json:"storage,omitempty"`
	Tags       []int       `json:"tags,omitempty"`
	Host       HostParams  `json:"host,omitempty"`
	Node       NodeHost    `json:"node,omitempty"`
	Cluster    ClusterID   `json:"cluster,omitempty"`
}

type ClusterID struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
}

type HostParams struct {
	ID        int    `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	BootOrder int    `json:"boot_order,omitempty"`
	IsMain    bool   `json:"is_main,omitempty"`
	Bus       string `json:"bus,omitempty"`
}

type StorageID struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type NewDiskRequest struct {
	Name       string `json:"name,omitempty"`
	SizeMib    int    `json:"size_mib,omitempty"`
	Node       int    `json:"node,omitempty"`
	Account    int    `json:"account,omitempty"`
	Storage    int    `json:"storage,omitempty"`
	ExpandPart string `json:"expand_part,omitempty"`
}

type UpdateDiskRequest struct {
	Name       string `json:"name,omitempty"`
	Account    int    `json:"account,omitempty"`
	SizeMib    int    `json:"size_mib,omitempty"`
	ExpandPart string `json:"expand_part,omitempty"`
	Pool       int    `json:"pool,omitempty"`
	Defer      Action `json:"defer,omitempty"`
}

type Action struct {
	Action string `json:"action,omitempty"`
}

type MigrateDiskRequest struct {
	Storage int  `json:"storage,omitempty"`
	Node    int  `json:"node,omitempty"`
	Plain   bool `json:"plain,omitempty"`
}

type BackupIDRequest struct {
	Backup int `json:"backup,omitempty"`
}

type UpdateBootOrderRequest struct {
	Disks []DiskBootOrder `json:"disks,omitempty"`
}

type DiskBootOrder struct {
	ID        int `json:"id,omitempty"`
	BootOrder int `json:"boot_order,omitempty"`
}

type DiskParams struct {
	IsMain bool   `json:"is_main,omitempty"`
	Bus    string `json:"bus,omitempty"`
}

// Host

type IP4 struct {
	IP string `json:"ip,omitempty"`
}

type HostInterface struct {
	HostInterface string `json:"host_interface,omitempty"`
	Mac           string `json:"mac,omitempty"`
	NodeInterface int    `json:"node_interface,omitempty"`
}

type HostNode struct {
	ID     int    `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	IPAddr string `json:"ip_addr,omitempty"`
}

type HostDisk struct {
	ID         int `json:"id,omitempty"`
	DiskMib    int `json:"disk_mib,omitempty"`
	DiskMibNew int `json:"disk_mib_new,omitempty"`
}

type FirewallRule struct {
	Action    string   `json:"action,omitempty"`
	Direction string   `json:"direction,omitempty"`
	Protocols []string `json:"protocols,omitempty"`
	Portstart int      `json:"portstart,omitempty"`
	Portend   int      `json:"portend,omitempty"`
}

type HostsResponse struct {
	LastNotify int           `json:"last_notify,omitempty"`
	List       []HostElement `json:"list,omitempty"`

	// Returns on API related errors
	Error Error `json:"error,omitempty"`
}

type HostElement struct {
	ExpandPart                string          `json:"expand_part,omitempty"`
	ID                        int             `json:"id,omitempty"`
	Name                      string          `json:"name,omitempty"`
	IP4                       []IP4           `json:"ip4,omitempty"`
	Interfaces                []HostInterface `json:"interfaces,omitempty"`
	Node                      HostNode        `json:"node,omitempty"`
	Cluster                   Cluster         `json:"cluster,omitempty"`
	State                     string          `json:"state,omitempty"`
	Domain                    string          `json:"domain,omitempty"`
	Account                   BaseAccount     `json:"account,omitempty"`
	Comment                   string          `json:"comment,omitempty"`
	Disk                      HostDisk        `json:"disk,omitempty"`
	DiskCount                 int             `json:"disk_count,omitempty"`
	CPUNumber                 int             `json:"cpu_number,omitempty"`
	CPUNumberNew              int             `json:"cpu_number_new,omitempty"`
	RAMMib                    int             `json:"ram_mib,omitempty"`
	RAMMibNew                 int             `json:"ram_mib_new,omitempty"`
	NetBandwidthMbitps        int             `json:"net_bandwidth_mbitps,omitempty"`
	NetBandwidthMbitpsChanged bool            `json:"net_bandwidth_mbitps_changed,omitempty"`
	IPAutomation              string          `json:"ip_automation,omitempty"`
	NetIsSynced               bool            `json:"net_is_synced,omitempty"`
	Tags                      []string        `json:"tags,omitempty"`
	OsName                    string          `json:"os_name,omitempty"`
	OsGroup                   string          `json:"os_group,omitempty"`
	Uptime                    int             `json:"uptime,omitempty"`
	RescueMode                bool            `json:"rescue_mode,omitempty"`
	IsoMounted                bool            `json:"iso_mounted,omitempty"`
	CPUMode                   string          `json:"cpu_mode,omitempty"`
	Nesting                   bool            `json:"nesting,omitempty"`
	CPUCustomModel            string          `json:"cpu_custom_model,omitempty"`
	CPUWeight                 int             `json:"cpu_weight,omitempty"`
	IoWeight                  int             `json:"io_weight,omitempty"`
	IoReadMbitps              int             `json:"io_read_mbitps,omitempty"`
	IoWriteMbitps             int             `json:"io_write_mbitps,omitempty"`
	IoReadIops                int             `json:"io_read_iops,omitempty"`
	IoWriteIops               int             `json:"io_write_iops,omitempty"`
	NetInMbitps               int             `json:"net_in_mbitps,omitempty"`
	NetOutMbitps              int             `json:"net_out_mbitps,omitempty"`
	NetWeight                 int             `json:"net_weight,omitempty"`
	AntiSpoofing              bool            `json:"anti_spoofing,omitempty"`
	Disabled                  bool            `json:"disabled,omitempty"`
	TCPConnectionsIn          int             `json:"tcp_connections_in,omitempty"`
	TCPConnectionsOut         int             `json:"tcp_connections_out,omitempty"`
	ProcessNumber             int             `json:"process_number,omitempty"`
	Vxlan                     Vxlan           `json:"vxlan,omitempty"`
	FirewallRules             []FirewallRule  `json:"firewall_rules,omitempty"`
	HasNonameIface            bool            `json:"has_noname_iface,omitempty"`
	VM5Restrictions           VM5Restrictions `json:"vm5_restrictions,omitempty"`
}

type VM5Restrictions struct {
	NetIfaceCount      bool `json:"net_iface_count,omitempty"`
	NatOrExtra         bool `json:"nat_or_extra,omitempty"`
	Ipv6               bool `json:"ipv6,omitempty"`
	UnsupportedStorage bool `json:"unsupported_storage,omitempty"`
	Iso                bool `json:"iso,omitempty"`
	Snapshot           bool `json:"snapshot,omitempty"`
}

type RecipeElement struct {
	Recipe       int            `json:"recipe,omitempty"`
	RecipeParams []RecipeParams `json:"recipe_params,omitempty"`
}

type RecipeParams struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type DiskHostRequest struct {
	Name       string `json:"name,omitempty"`
	SizeMib    int    `json:"size_mib,omitempty"`
	BootOrder  int    `json:"boot_order,omitempty"`
	ExpandPart string `json:"expand_part,omitempty"`
	Storage    int    `json:"storage,omitempty"`
	Tags       []int  `json:"tags,omitempty"`
}

type IPAddr struct {
	Name              string `json:"name,omitempty"`
	IPPool            int    `json:"ip_pool,omitempty"`
	IPNetwork         int    `json:"ip_network,omitempty"`
	WithoutAllocation bool   `json:"without_allocation,omitempty"`
}

type CustomInterface struct {
}

type HostVxlanRequest struct {
	ID         int `json:"id,omitempty"`
	Ipv4Number int `json:"ipv4_number,omitempty"`
	Ipnet      int `json:"ipnet,omitempty"`
}

type NewHostRequest struct {
	Name                string             `json:"name,omitempty"`
	Cluster             int                `json:"cluster,omitempty"`
	Node                int                `json:"node,omitempty"`
	Storage             int                `json:"storage,omitempty"`
	Account             int                `json:"account,omitempty"`
	Domain              string             `json:"domain,omitempty"`
	Preset              int                `json:"preset,omitempty"`
	Os                  int                `json:"os,omitempty"`
	Image               int                `json:"image,omitempty"`
	ExpandPart          string             `json:"expand_part,omitempty"`
	RecipeList          []RecipeElement    `json:"recipe_list,omitempty"`
	Recipe              int                `json:"recipe,omitempty"`
	RecipeParams        []RecipeParams     `json:"recipe_params,omitempty"`
	IgnoreRecipeFilters bool               `json:"ignore_recipe_filters,omitempty"`
	Password            string             `json:"password,omitempty"`
	RAMMib              int                `json:"ram_mib,omitempty"`
	HddMib              int                `json:"hdd_mib,omitempty"`
	Disks               []DiskHostRequest  `json:"disks,omitempty"`
	CPUNumber           int                `json:"cpu_number,omitempty"`
	NetBandwidthMbitps  int                `json:"net_bandwidth_mbitps,omitempty"`
	IPAddr              IPAddr             `json:"ip_addr,omitempty"`
	Ipv6Enabled         bool               `json:"ipv6_enabled,omitempty"`
	Ipv6Pool            []int              `json:"ipv6_pool,omitempty"`
	Ipv6Prefix          int                `json:"ipv6_prefix,omitempty"`
	Ipv4Pool            []int              `json:"ipv4_pool,omitempty"`
	Ipv4Number          int                `json:"ipv4_number,omitempty"`
	Comment             string             `json:"comment,omitempty"`
	Interfaces          []Interface        `json:"interfaces,omitempty"`
	CustomInterfaces    []CustomInterface  `json:"custom_interfaces,omitempty"`
	CPUMode             string             `json:"cpu_mode,omitempty"`
	Nesting             bool               `json:"nesting,omitempty"`
	CPUCustomModel      string             `json:"cpu_custom_model,omitempty"`
	CPUWeight           int                `json:"cpu_weight,omitempty"`
	IoWeight            int                `json:"io_weight,omitempty"`
	IoReadMbitps        int                `json:"io_read_mbitps,omitempty"`
	IoWriteMbitps       int                `json:"io_write_mbitps,omitempty"`
	IoReadIops          int                `json:"io_read_iops,omitempty"`
	IoWriteIops         int                `json:"io_write_iops,omitempty"`
	NetInMbitps         int                `json:"net_in_mbitps,omitempty"`
	NetOutMbitps        int                `json:"net_out_mbitps,omitempty"`
	NetWeight           int                `json:"net_weight,omitempty"`
	AntiSpoofing        bool               `json:"anti_spoofing,omitempty"`
	TCPConnectionsIn    int                `json:"tcp_connections_in,omitempty"`
	TCPConnectionsOut   int                `json:"tcp_connections_out,omitempty"`
	ProcessNumber       int                `json:"process_number,omitempty"`
	FirewallRules       []FirewallRule     `json:"firewall_rules,omitempty"`
	SendEmailMode       string             `json:"send_email_mode,omitempty"`
	Vxlan               []HostVxlanRequest `json:"vxlan,omitempty"`
}

type RecipeTask struct {
	ID             int   `json:"id,omitempty"`
	Task           int   `json:"task,omitempty"`
	RecipeTaskList []int `json:"recipe_task_list,omitempty"`
	RecipeTask     int   `json:"recipe_task,omitempty"`

	// Returns on API related errors
	Error Error `json:"error,omitempty"`
}

type HostResponse struct {
	ID           int    `json:"id,omitempty"`
	InternalName string `json:"internal_name,omitempty"`
	Node         int    `json:"node,omitempty"`
	Name         string `json:"name,omitempty"`
	State        string `json:"state,omitempty"`
	CPUNumber    int    `json:"cpu_number,omitempty"`
	RAMMib       int    `json:"ram_mib,omitempty"`
	ExpandPart   string `json:"expand_part,omitempty"`
	NetIsSynced  bool   `json:"net_is_synced,omitempty"`
	IPAutomation string `json:"ip_automation,omitempty"`

	// Returns on API related errors
	Error Error `json:"error,omitempty"`
}

type Property struct {
}

type HostUpdateValues struct {
	Password         string   `json:"password,omitempty"`
	VncPassword      string   `json:"vnc_password,omitempty"`
	State            string   `json:"state,omitempty"`
	StateUpdatedDate string   `json:"state_updated_date,omitempty"`
	GuestAgent       bool     `json:"guest_agent,omitempty"`
	VncPort          int      `json:"vnc_port,omitempty"`
	Property         Property `json:"property,omitempty"`
	StartDate        int      `json:"start_date,omitempty"`
}
type HostUpdateRequest struct {
	Values            HostUpdateValues `json:"values,omitempty"`
	Comment           string           `json:"comment,omitempty"`
	Name              string           `json:"name,omitempty"`
	HaRestoreOnFail   bool             `json:"ha_restore_on_fail,omitempty"`
	HaRestorePriority int              `json:"ha_restore_priority,omitempty"`
	ExpandPart        string           `json:"expand_part,omitempty"`
}

type Account struct {
	Account int `json:"account,omitempty"`
}

type ImageSize struct {
	ImageGib int `json:"image_gib,omitempty"`

	// Returns on API related errors
	Error Error `json:"error,omitempty"`
}

type HistoryElement struct {
	Name       string      `json:"name,omitempty"`
	DateCreate string      `json:"date_create,omitempty"`
	State      string      `json:"state,omitempty"`
	Params     interface{} `json:"params,omitempty"`
	User       string      `json:"user,omitempty"`
}

type HostHistoryResponse struct {
	List []HistoryElement `json:"list,omitempty"`

	// Returns on API related errors
	Error Error `json:"error,omitempty"`
}

type IFaceElement struct {
	ID     int      `json:"id,omitempty"`
	Name   string   `json:"name,omitempty"`
	Bridge string   `json:"bridge,omitempty"`
	IP     []string `json:"ip,omitempty"`
	Mac    string   `json:"mac,omitempty"`
	Model  string   `json:"model,omitempty"`
	Vxlan  int      `json:"vxlan,omitempty"`
}

type IFace struct {
	LastNotify int            `json:"last_notify,omitempty"`
	List       []IFaceElement `json:"list,omitempty"`

	// Returns on API related errors
	Error Error `json:"error,omitempty"`
}

type IFaceParams struct {
	Name          string `json:"name,omitempty"`
	Mac           string `json:"mac,omitempty"`
	Bridge        string `json:"bridge,omitempty"`
	IsMainNetwork bool   `json:"is_main_network,omitempty"`
	Ippool        int    `json:"ippool,omitempty"`
	IPCount       int    `json:"ip_count,omitempty"`
}

type IFaceModel struct {
	Model string `json:"model,omitempty"`
}

type AddIPToHostRequest struct {
	IPAddr      IPAddr             `json:"ip_addr,omitempty"`
	Ipv6Pool    []int              `json:"ipv6_pool,omitempty"`
	Ipv6Prefix  int                `json:"ipv6_prefix,omitempty"`
	Ipv6Enabled bool               `json:"ipv6_enabled,omitempty"`
	Ipv4Pool    []int              `json:"ipv4_pool,omitempty"`
	Ipv4Number  int                `json:"ipv4_number,omitempty"`
	Interfaces  []interface{}      `json:"interfaces,omitempty"`
	Vxlan       []HostVxlanRequest `json:"vxlan,omitempty"`
}

type IPAutomationType struct {
	IPAutomationType string `json:"ip_automation_type,omitempty"`
}

type IPV4Host struct {
	ID        int    `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Interface int    `json:"interface,omitempty"`
}

type IPV4InfoElement struct {
	ID               int      `json:"id,omitempty"`
	IPAddr           string   `json:"ip_addr,omitempty"`
	Domain           string   `json:"domain,omitempty"`
	Gateway          string   `json:"gateway,omitempty"`
	Mask             string   `json:"mask,omitempty"`
	State            string   `json:"state,omitempty"`
	Family           int      `json:"family,omitempty"`
	Ippool           int      `json:"ippool,omitempty"`
	Network          int      `json:"network,omitempty"`
	Host             IPV4Host `json:"host,omitempty"`
	ClusterInterface int      `json:"cluster_interface,omitempty"`
	Vxlan            Vxlan    `json:"vxlan,omitempty"`
}

type IPV4Info struct {
	LastNotify int               `json:"last_notify,omitempty"`
	List       []IPV4InfoElement `json:"list,omitempty"`
	Size       int               `json:"size,omitempty"`

	// Returns on API related errors
	Error Error `json:"error,omitempty"`
}

type IPV6Info struct {
	LastNotify     int  `json:"last_notify,omitempty"`
	Ipv6Enabled    bool `json:"ipv6_enabled,omitempty"`
	Ipv6PtrEnabled bool `json:"ipv6_ptr_enabled,omitempty"`

	// Returns on API related errors
	Error Error `json:"error,omitempty"`
}

type ISOTags struct {
	URL  string   `json:"url,omitempty"`
	Tags []string `json:"tags,omitempty"`

	// Returns on API related errors
	Error Error `json:"error,omitempty"`
}

type Socket struct {
	Socket string `json:"socket,omitempty"`

	// Returns on API related errors
	Error Error `json:"error,omitempty"`
}

type Metadata struct {
	Metadata interface{} `json:"metadata,omitempty"`
}

type HostMigrationStorage struct {
	ID           int    `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	Type         string `json:"type,omitempty"`
	AvailableMib int    `json:"available_mib,omitempty"`
}

type HostMigrateResponse struct {
	Nodes []NodeElement `json:"nodes,omitempty"`

	// Returns on API related errors
	Error Error `json:"error,omitempty"`
}

type DiskMigration struct {
	ID         int `json:"id,omitempty"`
	DstStorage int `json:"dst_storage,omitempty"`
}

type HostMigrateRequest struct {
	Plain bool            `json:"plain,omitempty"`
	Node  int             `json:"node,omitempty"`
	Disks []DiskMigration `json:"disks,omitempty"`
}

type PTRUpdateRequest struct {
	Name   string `json:"name,omitempty"`
	Domain string `json:"domain,omitempty"`
}

type PTRIP struct {
	ID     int    `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Domain string `json:"domain,omitempty"`
	Status string `json:"status,omitempty"`
}

type PTRElement struct {
	Family  int     `json:"family,omitempty"`
	Gateway string  `json:"gateway,omitempty"`
	IP      []PTRIP `json:"ip,omitempty"`
	IPCount int     `json:"ip_count,omitempty"`
	Network string  `json:"network,omitempty"`
	Prefix  int     `json:"prefix,omitempty"`
}

type PTRResponse struct {
	List []PTRElement `json:"list,omitempty"`

	// Returns on API related errors
	Error Error `json:"error,omitempty"`
}

type Domain struct {
	Domain string `json:"domain,omitempty"`
}

type HostRecipe struct {
	Recipe       int                `json:"recipe,omitempty"`
	RecipeParams []HostRecipeParams `json:"recipe_params,omitempty"`
}

type HostRecipeParams struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type HostReinstallRequest struct {
	Os                  int                `json:"os,omitempty"`
	Image               int                `json:"image,omitempty"`
	RecipeList          []HostRecipe       `json:"recipe_list,omitempty"`
	Recipe              int                `json:"recipe,omitempty"`
	RecipeParams        []HostRecipeParams `json:"recipe_params,omitempty"`
	IgnoreRecipeFilters bool               `json:"ignore_recipe_filters,omitempty"`
	SendEmailMode       string             `json:"send_email_mode,omitempty"`
	Password            string             `json:"password,omitempty"`
	Disk                int                `json:"disk,omitempty"`
}

type RescueMode struct {
	RescueMode bool `json:"rescue_mode,omitempty"`
}

type HostResourceUpdateRequest struct {
	CPUNumber          int            `json:"cpu_number,omitempty"`
	RAMMib             int            `json:"ram_mib,omitempty"`
	NetBandwidthMbitps int            `json:"net_bandwidth_mbitps,omitempty"`
	CPUMode            string         `json:"cpu_mode,omitempty"`
	Nesting            bool           `json:"nesting,omitempty"`
	CPUCustomModel     string         `json:"cpu_custom_model,omitempty"`
	CPUWeight          int            `json:"cpu_weight,omitempty"`
	IoWeight           int            `json:"io_weight,omitempty"`
	IoReadMbitps       int            `json:"io_read_mbitps,omitempty"`
	IoWriteMbitps      int            `json:"io_write_mbitps,omitempty"`
	IoReadIops         int            `json:"io_read_iops,omitempty"`
	IoWriteIops        int            `json:"io_write_iops,omitempty"`
	NetInMbitps        int            `json:"net_in_mbitps,omitempty"`
	NetOutMbitps       int            `json:"net_out_mbitps,omitempty"`
	NetWeight          int            `json:"net_weight,omitempty"`
	AntiSpoofing       bool           `json:"anti_spoofing,omitempty"`
	TCPConnectionsIn   int            `json:"tcp_connections_in,omitempty"`
	TCPConnectionsOut  int            `json:"tcp_connections_out,omitempty"`
	ProcessNumber      int            `json:"process_number,omitempty"`
	FirewallRules      []FirewallRule `json:"firewall_rules,omitempty"`
	Defer              Action         `json:"defer,omitempty"`
}

type HostBackup struct {
	Backup int `json:"backup,omitempty"`
}

type Recipient struct {
	Lang  string `json:"lang,omitempty"`
	Email string `json:"email,omitempty"`
}

type Recipe struct {
	Recipe              int            `json:"recipe,omitempty"`
	Body                string         `json:"body,omitempty"`
	RecipeParams        []RecipeParams `json:"recipe_params,omitempty"`
	SendEmail           bool           `json:"send_email,omitempty"`
	IgnoreRecipeFilters bool           `json:"ignore_recipe_filters,omitempty"`
	Recipients          []Recipient    `json:"recipients,omitempty"`
}

type VNCPort struct {
	Domain  string `json:"domain,omitempty"`
	VncPort int    `json:"vnc_port,omitempty"`
	VncPass string `json:"vnc_pass,omitempty"`
}

type VNCPortUpdateRequest struct {
	List []VNCPort `json:"list,omitempty"`
}

type IP struct {
	ID               int    `json:"id,omitempty"`
	IPAddr           string `json:"ip_addr,omitempty"`
	Domain           string `json:"domain,omitempty"`
	Gateway          string `json:"gateway,omitempty"`
	Mask             string `json:"mask,omitempty"`
	State            string `json:"state,omitempty"`
	Family           int    `json:"family,omitempty"`
	Ippool           int    `json:"ippool,omitempty"`
	Network          int    `json:"network,omitempty"`
	Host             IPHost `json:"host,omitempty"`
	ClusterInterface int    `json:"cluster_interface,omitempty"`
}

type IPHost struct {
	ID        int    `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Interface int    `json:"interface,omitempty"`
}

type IPList struct {
	LastNotify int  `json:"last_notify,omitempty"`
	List       []IP `json:"list,omitempty"`
	Size       int  `json:"size,omitempty"`

	// Returns on API related errors
	Error Error `json:"error,omitempty"`
}

type UpdateIPResponseValues struct {
	Domain  string `json:"domain,omitempty"`
	Gateway string `json:"gateway,omitempty"`
	IPAddr  string `json:"ip_addr,omitempty"`
	Netmask string `json:"netmask,omitempty"`
}
type UpdateIPResponse struct {
	Values UpdateIPResponseValues `json:"values,omitempty"`
}

type ScheduleElement struct {
	Name           string      `json:"name,omitempty"`
	Handler        string      `json:"handler,omitempty"`
	Service        string      `json:"service,omitempty"`
	Method         string      `json:"method,omitempty"`
	InstanceID     int         `json:"instance_id,omitempty"`
	URLQueryParams interface{} `json:"url_query_params,omitempty"`
	PostParams     interface{} `json:"post_params,omitempty"`
	CronExpression string      `json:"cron_expression,omitempty"`
}

type ScheduleListResponse struct {
	SheduleList []ScheduleElement `json:"shedule_list,omitempty"`

	// Returns on API related errors
	Error Error `json:"error,omitempty"`
}

// Images

type Image struct {
	Name    string `json:"name,omitempty"`
	Account int    `json:"account,omitempty"`
	Comment string `json:"comment,omitempty"`
	ForAll  bool   `json:"for_all,omitempty"`
}

type ImagesListResponse struct {
	LastNotify int                         `json:"last_notify,omitempty"`
	List       []ImagesListResponseElement `json:"list,omitempty"`

	// Returns on API related errors
	Error Error `json:"error,omitempty"`
}

type ImagesListResponseElement struct {
	ID      int                      `json:"id,omitempty"`
	Name    string                   `json:"name,omitempty"`
	Account string                   `json:"account,omitempty"`
	ForAll  bool                     `json:"for_all,omitempty"`
	Type    string                   `json:"type,omitempty"`
	Comment string                   `json:"comment,omitempty"`
	State   string                   `json:"state,omitempty"`
	Nodes   []ImagesListResponseNode `json:"nodes,omitempty"`
}

type ImagesListResponseNode struct {
	ID      int    `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	IPAddr  string `json:"ip_addr,omitempty"`
	SSHPort int    `json:"ssh_port,omitempty"`
}

type ImageResponse struct {
	ID    int    `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	State string `json:"state,omitempty"`

	// Returns on API related errors
	Error Error `json:"error,omitempty"`
}

type ImageState struct {
	State string `json:"state,omitempty"`
}

// Import

type ImportHistoryResponse struct {
	List []ImportHistoryResponseElement `json:"list,omitempty"`

	// Returns on API related errors
	Error Error `json:"error,omitempty"`
}

type ImportHistoryResponseElement struct {
	Data   string `json:"data,omitempty"`
	Date   string `json:"date,omitempty"`
	TaskID int    `json:"task_id,omitempty"`
	Type   string `json:"type,omitempty"`
}

type ImportHistoryRequest struct {
	Data   string `json:"data,omitempty"`
	TaskID int    `json:"task_id,omitempty"`
	State  string `json:"state,omitempty"`
	Type   string `json:"type,omitempty"`
	URL    string `json:"url,omitempty"`
}

type ImportHistoryResult struct {
	State  string `json:"state,omitempty"`
	Result Result `json:"result,omitempty"`
}

type Result struct {
	AdditionalProp1 interface{} `json:"additionalProp1,omitempty"`
}

type ImportedFrom struct {
	IP       string `json:"ip,omitempty"`
	Type     string `json:"type,omitempty"`
	SSHPort  int    `json:"ssh_port,omitempty"`
	Password string `json:"password,omitempty"`
}

type User struct {
	ID              int         `json:"id,omitempty"`
	Name            string      `json:"name,omitempty"`
	Level           string      `json:"level,omitempty"`
	Enabled         bool        `json:"enabled,omitempty"`
	AdditionalProp1 interface{} `json:"additionalProp1,omitempty"`
}

type Event struct {
	AdditionalProp1 interface{} `json:"additionalProp1,omitempty"`
}

type Network struct {
	AdditionalProp1 interface{} `json:"additionalProp1,omitempty"`
}

type ImportStorage struct {
	Name            string      `json:"name,omitempty"`
	PoolType        string      `json:"pool_type,omitempty"`
	TgtPath         string      `json:"tgt_path,omitempty"`
	AdditionalProp1 interface{} `json:"additionalProp1,omitempty"`
}

type Setting struct {
	AdditionalProp1 interface{} `json:"additionalProp1,omitempty"`
}

type ImportClusterParams struct {
	ImportHistory    int             `json:"import_history,omitempty"`
	ClusterName      string          `json:"cluster_name,omitempty"`
	OsStoragePath    string          `json:"os_storage_path,omitempty"`
	ImageStoragePath string          `json:"image_storage_path,omitempty"`
	RAMOverselling   int             `json:"ram_overselling,omitempty"`
	DNSServers       []string        `json:"dns_servers,omitempty"`
	Timezone         string          `json:"timezone,omitempty"`
	ImportedFrom     ImportedFrom    `json:"imported_from,omitempty"`
	User             []User          `json:"user,omitempty"`
	Event            Event           `json:"event,omitempty"`
	Network          []Network       `json:"network,omitempty"`
	Storage          []ImportStorage `json:"storage,omitempty"`
	Settings         []Setting       `json:"settings,omitempty"`
	Node             []ImportNode    `json:"node,omitempty"`
	AdditionalProp1  interface{}     `json:"additionalProp1,omitempty"`
}

type ImportNode struct {
	Name            string                 `json:"name,omitempty"`
	IP              string                 `json:"ip,omitempty"`
	SSHPort         int                    `json:"ssh_port,omitempty"`
	Password        string                 `json:"password,omitempty"`
	Status          int                    `json:"status,omitempty"`
	Fake            bool                   `json:"fake,omitempty"`
	PhysicalCpucore int                    `json:"physical_cpucore,omitempty"`
	AllocationRule  []ImportAllocationRule `json:"allocation_rule,omitempty"`
	Host            []ImportHost           `json:"host,omitempty"`
	AdditionalProp1 interface{}            `json:"additionalProp1,omitempty"`
}

type ImportAllocationRule struct {
	Priority        int         `json:"priority,omitempty"`
	RuleType        int         `json:"rule_type,omitempty"`
	RuleOper        int         `json:"rule_oper,omitempty"`
	Value           string      `json:"value,omitempty"`
	Action          int         `json:"action,omitempty"`
	Stop            bool        `json:"stop,omitempty"`
	AdditionalProp1 interface{} `json:"additionalProp1,omitempty"`
}

type ImportHost struct {
	Name            string            `json:"name,omitempty"`
	CPU             int               `json:"cpu,omitempty"`
	RAM             int               `json:"ram,omitempty"`
	Fake            bool              `json:"fake,omitempty"`
	Interface       []ImportInterface `json:"interface,omitempty"`
	FirewallRules   []FirewallRule    `json:"firewall_rules,omitempty"`
	AdditionalProp1 interface{}       `json:"additionalProp1,omitempty"`
}

type ImportInterface struct {
	ID              int         `json:"id,omitempty"`
	VMID            int         `json:"vm_id,omitempty"`
	NetID           int         `json:"net_id,omitempty"`
	Mac             string      `json:"mac,omitempty"`
	NicModel        string      `json:"nic_model,omitempty"`
	AdditionalProp1 interface{} `json:"additionalProp1,omitempty"`
}

// License

type License struct {
	Key string `json:"key,omitempty"`
}

type Clusters struct {
	Clusters []int `json:"clusters,omitempty"`
}

type IPNetUpdate struct {
	Gateway     string `json:"gateway,omitempty"`
	NeedReserve bool   `json:"need_reserve,omitempty"`
	Name        string `json:"name,omitempty"`
	Vlan        string `json:"vlan,omitempty"`
	Note        string `json:"note,omitempty"`
}

type IPPoolConnectRequest struct {
	Clusters []IPPoolCluster `json:"clusters,omitempty"`
}

type IPPoolCluster struct {
	Name      string `json:"name,omitempty"`
	ID        int    `json:"id,omitempty"`
	Enabled   bool   `json:"enabled,omitempty"`
	Interface int    `json:"interface,omitempty"`
}

type IPPoolRangeRequest struct {
	Name string `json:"name,omitempty"`
}

type IPMgr5MigrateRequest struct {
	Dbhost     string `json:"dbhost,omitempty"`
	Dbname     string `json:"dbname,omitempty"`
	Dbuser     string `json:"dbuser,omitempty"`
	Dbpassword string `json:"dbpassword,omitempty"`
}

type HetznerIPs struct {
	HetznerIP []string `json:"hetzner_ip,omitempty"`
}

type NodeIPUpdateRequest struct {
	IP      string `json:"ip,omitempty"`
	SSHPort int    `json:"ssh_port,omitempty"`
}

type IPPoolResponse struct {
	LastNotify int      `json:"last_notify,omitempty"`
	List       []IPPool `json:"list,omitempty"`

	// Returns on API related errors
	Error Error `json:"error,omitempty"`
}

type IPPool struct {
	ID       int         `json:"id,omitempty"`
	Name     string      `json:"name,omitempty"`
	Note     string      `json:"note,omitempty"`
	TotalIP  string      `json:"total_ip,omitempty"`
	UsingIP  string      `json:"using_ip,omitempty"`
	Clusters []ClusterID `json:"clusters,omitempty"`
}

type IPPoolClusterListResponse struct {
	LastNotify int             `json:"last_notify,omitempty"`
	List       []IPPoolCluster `json:"list,omitempty"`

	// Returns on API related errors
	Error Error `json:"error,omitempty"`
}

type NodeHetznerIPResponse struct {
	LastNotify int                    `json:"last_notify,omitempty"`
	List       []NodeHetznerIPElement `json:"list,omitempty"`
	Size       int                    `json:"size,omitempty"`

	// Returns on API related errors
	Error Error `json:"error,omitempty"`
}

type NodeHetznerIPElement struct {
	Host   Host          `json:"host,omitempty"`
	IP     HetznerIP     `json:"ip,omitempty"`
	Subnet HetznerSubnet `json:"subnet,omitempty"`
}

type HetznerIP struct {
	ID     int    `json:"id,omitempty"`
	IPAddr string `json:"ip_addr,omitempty"`
	State  string `json:"state,omitempty"`
}

type HetznerSubnet struct {
	ID     int    `json:"id,omitempty"`
	Subnet string `json:"subnet,omitempty"`
	UsedIP int    `json:"used_ip,omitempty"`
}

// Nodes
type NodeCluster struct {
	DataCenter string `json:"data_center,omitempty"`
	ID         int    `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	TimeZone   string `json:"time_zone,omitempty"`
}
type NodeInterface struct {
	NodeInterface string `json:"node_interface,omitempty"`
	NetworkSwitch string `json:"network_switch,omitempty"`
	ID            int    `json:"id,omitempty"`
}
type NodesResponse struct {
	LastNotify int           `json:"last_notify,omitempty"`
	List       []NodeElement `json:"list,omitempty"`

	// Returns on API related errors
	Error Error `json:"error,omitempty"`
}

type NodeElement struct {
	ID                  int             `json:"id,omitempty"`
	Name                string          `json:"name,omitempty"`
	Cluster             NodeCluster     `json:"cluster,omitempty"`
	IP                  string          `json:"ip,omitempty"`
	Ipv6                string          `json:"ipv6,omitempty"`
	State               string          `json:"state,omitempty"`
	StorageMib          StorageMib      `json:"storage_mib,omitempty"`
	RAMMib              RAMMib          `json:"ram_mib,omitempty"`
	CPUNumber           int             `json:"cpu_number,omitempty"`
	Comment             string          `json:"comment,omitempty"`
	Interfaces          []NodeInterface `json:"interfaces,omitempty"`
	Overselling         int             `json:"overselling,omitempty"`
	HostCount           int             `json:"host_count,omitempty"`
	HostLimit           int             `json:"host_limit,omitempty"`
	HostFilter          []HostFilter    `json:"host_filter,omitempty"`
	HostCreationBlocked bool            `json:"host_creation_blocked,omitempty"`
}

type NodeScriptResponse struct {
	LastNotify int                 `json:"last_notify,omitempty"`
	List       []NodeScriptElement `json:"list,omitempty"`

	// Returns on API related errors
	Error Error `json:"error,omitempty"`
}

type UpdatedAt struct {
	Date    string      `json:"date,omitempty"`
	Account BaseAccount `json:"account,omitempty"`
}

type NodeScriptElement struct {
	ID          int         `json:"id,omitempty"`
	Name        string      `json:"name,omitempty"`
	Script      string      `json:"script,omitempty"`
	Description string      `json:"description,omitempty"`
	Type        string      `json:"type,omitempty"`
	Priority    int         `json:"priority,omitempty"`
	Account     BaseAccount `json:"account,omitempty"`
	Autorun     []int       `json:"autorun,omitempty"`
	UpdatedAt   UpdatedAt   `json:"updated_at,omitempty"`
}

type NodeReducedClusterListResponse struct {
	LastNotify int                         `json:"last_notify,omitempty"`
	List       []NodeReducedClusterElement `json:"list,omitempty"`

	// Returns on API related errors
	Error Error `json:"error,omitempty"`
}

type NodeReducedClusterElement struct {
	ID           int                       `json:"id,omitempty"`
	Cluster      ClusterVirtualizationType `json:"cluster,omitempty"`
	IP           string                    `json:"ip,omitempty"`
	LibvirtError string                    `json:"libvirt_error,omitempty"`
}

type ClusterVirtualizationType struct {
	VirtualizationType string `json:"virtualization_type,omitempty"`
}

type NodeResponse struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Cluster     int    `json:"cluster,omitempty"`
	State       string `json:"state,omitempty"`
	IPAddr      string `json:"ip_addr,omitempty"`
	QemuVersion string `json:"qemu_version,omitempty"`

	// Returns on API related errors
	Error Error `json:"error,omitempty"`
}

type NodeFileListResponse struct {
	LastNotify int                   `json:"last_notify,omitempty"`
	List       []NodeFileListElement `json:"list,omitempty"`
	Size       int                   `json:"size,omitempty"`

	// Returns on API related errors
	Error Error `json:"error,omitempty"`
}

type NodeFileListElement struct {
	EntityID int    `json:"entity_id,omitempty"`
	Name     string `json:"name,omitempty"`
	Node     int    `json:"node,omitempty"`
	Path     string `json:"path,omitempty"`
	SizeMib  int    `json:"size_mib,omitempty"`
	Type     string `json:"type,omitempty"`
}

type NodeHostIFaces struct {
	List []NodeHostIFace `json:"list,omitempty"`

	// Returns on API related errors
	Error Error `json:"error,omitempty"`
}

type NodeHostIFace struct {
	InternalName string  `json:"internal_name,omitempty"`
	NetIsSynced  bool    `json:"net_is_synced,omitempty"`
	Ifaces       []Iface `json:"ifaces,omitempty"`
	Date         string  `json:"date,omitempty"`
}

type Iface struct {
	Name string   `json:"name,omitempty"`
	Mac  string   `json:"mac,omitempty"`
	IP   []NodeIP `json:"ip,omitempty"`
}

type NodeIP struct {
	IPFamily int    `json:"ip_family,omitempty"`
	Name     string `json:"name,omitempty"`
}

type IPParams struct {
	V4 string `json:"v4,omitempty"`
	V6 string `json:"v6,omitempty"`
}

type NodeNetworkIFaces struct {
	LastNotify int                `json:"last_notify,omitempty"`
	List       []NodeNetworkIFace `json:"list,omitempty"`

	// Returns on API related errors
	Error Error `json:"error,omitempty"`
}

type NodeNetworkIFace struct {
	Name          string                     `json:"name,omitempty"`
	IPParams      IPParams                   `json:"ip_params,omitempty"`
	Type          string                     `json:"type,omitempty"`
	State         string                     `json:"state,omitempty"`
	Used          bool                       `json:"used,omitempty"`
	HostCount     int                        `json:"host_count,omitempty"`
	Vlan          int                        `json:"vlan,omitempty"`
	IsDefault     bool                       `json:"is_default,omitempty"`
	IsMainNetwork bool                       `json:"is_main_network,omitempty"`
	Slaves        []string                   `json:"slaves,omitempty"`
	Mode          string                     `json:"mode,omitempty"`
	OldValues     OldValuesNodeNetworkIFaces `json:"old_values,omitempty"`
}

type OldValuesNodeNetworkIFaces struct {
	Name          string   `json:"name,omitempty"`
	IPParams      IPParams `json:"ip_params,omitempty"`
	Type          string   `json:"type,omitempty"`
	State         string   `json:"state,omitempty"`
	Used          bool     `json:"used,omitempty"`
	HostCount     int      `json:"host_count,omitempty"`
	Vlan          int      `json:"vlan,omitempty"`
	IsDefault     bool     `json:"is_default,omitempty"`
	IsMainNetwork bool     `json:"is_main_network,omitempty"`
	Slaves        []string `json:"slaves,omitempty"`
	Mode          string   `json:"mode,omitempty"`
}

type HostStorage struct {
	Path string `json:"path,omitempty"`
	Tags []int  `json:"tags,omitempty"`
}

type NodeInterfaceRequest struct {
	NodeInterface    string `json:"node_interface,omitempty"`
	ClusterInterface int    `json:"cluster_interface,omitempty"`
}
type NodeNewRequest struct {
	Name                     string                 `json:"name,omitempty"`
	Cluster                  int                    `json:"cluster,omitempty"`
	IP                       string                 `json:"ip,omitempty"`
	Ipv6                     string                 `json:"ipv6,omitempty"`
	SSHPort                  int                    `json:"ssh_port,omitempty"`
	HostMax                  int                    `json:"host_max,omitempty"`
	Password                 string                 `json:"password,omitempty"`
	HostStorage              HostStorage            `json:"host_storage,omitempty"`
	ImageStoragePath         string                 `json:"image_storage_path,omitempty"`
	Comment                  string                 `json:"comment,omitempty"`
	Overselling              int                    `json:"overselling,omitempty"`
	LvmVgRename              bool                   `json:"lvm_vg_rename,omitempty"`
	HetznerIP                []string               `json:"hetzner_ip,omitempty"`
	Interfaces               []NodeInterfaceRequest `json:"interfaces,omitempty"`
	HostLimit                int                    `json:"host_limit,omitempty"`
	HostFilter               []HostFilter           `json:"host_filter,omitempty"`
	NetworkAutosetupDisabled bool                   `json:"network_autosetup_disabled,omitempty"`
}

type NodeScriptRequest struct {
	Name        string `json:"name,omitempty"`
	Script      string `json:"script,omitempty"`
	Description string `json:"description,omitempty"`
	Type        string `json:"type,omitempty"`
	Priority    int    `json:"priority,omitempty"`
	Autorun     []int  `json:"autorun,omitempty"`
}

type NodeUpdateRequest struct {
	HostCreationBlocked bool         `json:"host_creation_blocked,omitempty"`
	OpenVswitch         bool         `json:"open_vswitch,omitempty"`
	Comment             string       `json:"comment,omitempty"`
	Name                string       `json:"name,omitempty"`
	Overselling         int          `json:"overselling,omitempty"`
	Ipv6                string       `json:"ipv6,omitempty"`
	HostLimit           int          `json:"host_limit,omitempty"`
	HostFilter          []HostFilter `json:"host_filter,omitempty"`
	Values              interface{}  `json:"values,omitempty"`
}

type NodeFilesUpdateRequest struct {
	List []NodeFileUpdateRequest `json:"list,omitempty"`
}

type NodeFileUpdateRequest struct {
	Size int    `json:"size,omitempty"`
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
	Path string `json:"path,omitempty"`
}

type NetworkAutosetupDisabled struct {
	NetworkAutosetupDisabled bool `json:"network_autosetup_disabled,omitempty"`
}

type NodeNetworkInterface struct {
	Name          string   `json:"name,omitempty"`
	Type          string   `json:"type,omitempty"`
	IPParams      IPParams `json:"ip_params,omitempty"`
	Vlan          int      `json:"vlan,omitempty"`
	Slaves        []string `json:"slaves,omitempty"`
	TakeSlaveIP   bool     `json:"take_slave_ip,omitempty"`
	Mode          string   `json:"mode,omitempty"`
	IsDefault     bool     `json:"is_default,omitempty"`
	IsMainNetwork bool     `json:"is_main_network,omitempty"`
}

type NodeNetworkInterfaces struct {
	List []NodeNetworkInterface `json:"list,omitempty"`
}

type Problems struct {
	Problems []string `json:"problems,omitempty"`
}

type Script struct {
	Script int `json:"script,omitempty"`
}

type Scripts struct {
	Script []int `json:"script,omitempty"`
}

type OSClusterParams struct {
	ID        int    `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

type OSList struct {
	LastNotify int        `json:"last_notify,omitempty"`
	List       []OSParams `json:"list,omitempty"`

	// Returns on API related errors
	Error Error `json:"error,omitempty"`
}

type OSParams struct {
	Name         string                   `json:"name,omitempty"`
	ID           int                      `json:"id,omitempty"`
	Clusters     []OSClusterParams        `json:"clusters,omitempty"`
	Repository   string                   `json:"repository,omitempty"`
	UpdatedAt    string                   `json:"updated_at,omitempty"`
	Tags         []string                 `json:"tags,omitempty"`
	State        string                   `json:"state,omitempty"`
	Adminonly    bool                     `json:"adminonly,omitempty"`
	Nodes        []ImagesListResponseNode `json:"nodes,omitempty"`
	OsGroup      string                   `json:"os_group,omitempty"`
	KmsSupported bool                     `json:"kms_supported,omitempty"`
	KmsIP        string                   `json:"kms_ip,omitempty"`
	KmsPort      string                   `json:"kms_port,omitempty"`
	IsLxdImage   bool                     `json:"is_lxd_image,omitempty"`
}

type OSUpdateParams struct {
	Adminonly  bool   `json:"adminonly,omitempty"`
	Clusters   []int  `json:"clusters,omitempty"`
	ProductKey string `json:"product_key,omitempty"`
	KmsIP      string `json:"kms_ip,omitempty"`
	KmsPort    string `json:"kms_port,omitempty"`
}

type OSSaveRequest struct {
	RepositoryID   int      `json:"repository_id,omitempty"`
	Os             []string `json:"os,omitempty"`
	RepositoryType string   `json:"repository_type,omitempty"`
}

type PlatformBackupSchedule struct {
	LastNotify int                        `json:"last_notify,omitempty"`
	List       []PlatformBackuListElement `json:"list,omitempty"`

	// Returns on API related errors
	Error Error `json:"error,omitempty"`
}

type PlatformBackupElement struct {
	ID           int    `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	Downloadable bool   `json:"downloadable,omitempty"`
	Location     string `json:"location,omitempty"`
	Date         string `json:"date,omitempty"`
	Size         int    `json:"size,omitempty"`
}

type PlatformBackuListElement struct {
	ID               int                     `json:"id,omitempty"`
	ScheduleType     string                  `json:"schedule_type,omitempty"`
	CronExpression   string                  `json:"cron_expression,omitempty"`
	Name             string                  `json:"name,omitempty"`
	Enabled          bool                    `json:"enabled,omitempty"`
	NextRun          string                  `json:"next_run,omitempty"`
	Comment          string                  `json:"comment,omitempty"`
	StorageType      string                  `json:"storage_type,omitempty"`
	ConnectionParams interface{}             `json:"connection_params,omitempty"`
	BackupList       []PlatformBackupElement `json:"backup_list,omitempty"`
}

type PlatformBackupScheduleRequest struct {
	ScheduleType     string `json:"schedule_type,omitempty"`
	CronExpression   string `json:"cron_expression,omitempty"`
	Name             string `json:"name,omitempty"`
	Comment          string `json:"comment,omitempty"`
	StorageType      string `json:"storage_type,omitempty"`
	ConnectionParams struct {
	} `json:"connection_params,omitempty"`
}

// presets

type PresetListResponse struct {
	LastNotify int                 `json:"last_notify,omitempty"`
	List       []PresetListElement `json:"list,omitempty"`

	// Returns on API related errors
	Error Error `json:"error,omitempty"`
}

type PresetListElement struct {
	ID                 int               `json:"id,omitempty"`
	VirtualizationType string            `json:"virtualization_type,omitempty"`
	Name               string            `json:"name,omitempty"`
	Comment            string            `json:"comment,omitempty"`
	RAMMib             int               `json:"ram_mib,omitempty"`
	HddMib             int               `json:"hdd_mib,omitempty"`
	Disks              []DiskHostRequest `json:"disks,omitempty"`
	CPUNumber          int               `json:"cpu_number,omitempty"`
	CPUMode            string            `json:"cpu_mode,omitempty"`
	Nesting            bool              `json:"nesting,omitempty"`
	CPUCustomModel     string            `json:"cpu_custom_model,omitempty"`
	CPUWeight          int               `json:"cpu_weight,omitempty"`
	IoWeight           int               `json:"io_weight,omitempty"`
	IoReadMbitps       int               `json:"io_read_mbitps,omitempty"`
	IoWriteMbitps      int               `json:"io_write_mbitps,omitempty"`
	IoReadIops         int               `json:"io_read_iops,omitempty"`
	IoWriteIops        int               `json:"io_write_iops,omitempty"`
	NetInMbitps        int               `json:"net_in_mbitps,omitempty"`
	NetOutMbitps       int               `json:"net_out_mbitps,omitempty"`
	NetWeight          int               `json:"net_weight,omitempty"`
	AntiSpoofing       bool              `json:"anti_spoofing,omitempty"`
	TCPConnectionsIn   int               `json:"tcp_connections_in,omitempty"`
	TCPConnectionsOut  int               `json:"tcp_connections_out,omitempty"`
	ProcessNumber      int               `json:"process_number,omitempty"`
	FirewallRules      []FirewallRule    `json:"firewall_rules,omitempty"`
}

type MigrateFormResponse struct {
	Nodes []struct {
		ID     int    `json:"id,omitempty"`
		Name   string `json:"name,omitempty"`
		RAMMib struct {
			Unused int `json:"unused,omitempty"`
		} `json:"ram_mib,omitempty"`
		Storage struct {
			ID           int    `json:"id,omitempty"`
			Name         string `json:"name,omitempty"`
			Type         string `json:"type,omitempty"`
			AvailableMib int    `json:"available_mib,omitempty"`
		} `json:"storage,omitempty"`
		Storages []struct {
			ID           int    `json:"id,omitempty"`
			Name         string `json:"name,omitempty"`
			Type         string `json:"type,omitempty"`
			AvailableMib int    `json:"available_mib,omitempty"`
		} `json:"storages,omitempty"`
		Priority            int    `json:"priority,omitempty"`
		State               string `json:"state,omitempty"`
		Suitable            bool   `json:"suitable,omitempty"`
		HostCreationBlocked bool   `json:"host_creation_blocked,omitempty"`
		HostCount           int    `json:"host_count,omitempty"`
		HostLimit           int    `json:"host_limit,omitempty"`
	} `json:"nodes,omitempty"`

	// Returns on API related errors
	Error Error `json:"error,omitempty"`
}

// recipes

type Recipes struct {
	LastNotify int          `json:"last_notify,omitempty"`
	List       []RecipeResp `json:"list,omitempty"`

	// Returns on API related errors
	Error Error `json:"error,omitempty"`
}

type RecipeResp struct {
	ID          int      `json:"id,omitempty"`
	Name        string   `json:"name,omitempty"`
	Repository  string   `json:"repository,omitempty"`
	UpdatedAt   string   `json:"updated_at,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	HasEmail    bool     `json:"has_email,omitempty"`
	Access      string   `json:"access,omitempty"`
	ParamsCount int      `json:"params_count,omitempty"`
}

type RecipeNew struct {
	Name         string            `json:"name,omitempty"`
	Tags         []string          `json:"tags,omitempty"`
	Description  string            `json:"description,omitempty"`
	ScriptFilter []ScriptFilter    `json:"script_filter,omitempty"`
	Script       string            `json:"script,omitempty"`
	ForAll       bool              `json:"for_all,omitempty"`
	UserVisible  bool              `json:"user_visible,omitempty"`
	Account      int               `json:"account,omitempty"`
	EmailContent []EmailContent    `json:"email_content,omitempty"`
	Type         string            `json:"type,omitempty"`
	Params       []RecipeNewParams `json:"params,omitempty"`
}

type ScriptFilter struct {
	Entity     string `json:"entity,omitempty"`
	Expression string `json:"expression,omitempty"`
}

type EmailContent struct {
	Lang    string `json:"lang,omitempty"`
	Content string `json:"content,omitempty"`
	Subject string `json:"subject,omitempty"`
}

type RecipeNewParams struct {
	Name         string         `json:"name,omitempty"`
	Description  string         `json:"description,omitempty"`
	Required     bool           `json:"required,omitempty"`
	Type         string         `json:"type,omitempty"`
	SelectValues []SelectValues `json:"select_values,omitempty"`
}

type SelectValues struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type RecipeResponse struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type RecipesSaveRequest struct {
	RepositoryID int         `json:"repository_id,omitempty"`
	URL          string      `json:"url,omitempty"`
	Recipes      interface{} `json:"recipes,omitempty"`
}

type ScriptVariablesResponse struct {
	LastNotify int              `json:"last_notify,omitempty"`
	List       []ScriptVariable `json:"list,omitempty"`

	// Returns on API related errors
	Error Error `json:"error,omitempty"`
}

type ScriptVariable struct {
	Name    string `json:"name,omitempty"`
	Value   string `json:"value,omitempty"`
	Comment string `json:"comment,omitempty"`
	Hidden  bool   `json:"hidden,omitempty"`
	Enabled bool   `json:"enabled,omitempty"`
	Builtin bool   `json:"builtin,omitempty"`
	Account int    `json:"account,omitempty"`
	ID      int    `json:"id,omitempty"`
	Scope   string `json:"scope,omitempty"`

	// Returns on API related errors
	Error Error `json:"error,omitempty"`
}

type Repositories struct {
	LastNotify int          `json:"last_notify,omitempty"`
	List       []Repository `json:"list,omitempty"`

	// Returns on API related errors
	Error Error `json:"error,omitempty"`
}

type Repository struct {
	Name  string `json:"name,omitempty"`
	URL   string `json:"url,omitempty"`
	ID    int    `json:"id,omitempty"`
	State string `json:"state,omitempty"`
	Type  string `json:"type,omitempty"`
}

type RepositoryName struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`

	// Returns on API related errors
	Error Error `json:"error,omitempty"`
}

type RepositoryNew struct {
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}

// tags

type Tags struct {
	LastNotify int   `json:"last_notify,omitempty"`
	List       []Tag `json:"list,omitempty"`

	// Returns on API related errors
	Error Error `json:"error,omitempty"`
}

type Tag struct {
	ID       int          `json:"id,omitempty"`
	Name     string       `json:"name,omitempty"`
	Note     string       `json:"note,omitempty"`
	Immortal bool         `json:"immortal,omitempty"`
	Storages []TagStorage `json:"storages,omitempty"`
}

type TagStorage struct {
	ID      int          `json:"id,omitempty"`
	Name    string       `json:"name,omitempty"`
	Type    string       `json:"type,omitempty"`
	Cluster []TagCluster `json:"cluster,omitempty"`
}
type TagCluster struct {
	ID                 int    `json:"id,omitempty"`
	Name               string `json:"name,omitempty"`
	VirtualizationType string `json:"virtualization_type,omitempty"`
}

type TagNew struct {
	Name string `json:"name,omitempty"`
	Note string `json:"note,omitempty"`
}

// schedule

type Schedules struct {
	LastNotify int        `json:"last_notify,omitempty"`
	List       []Schedule `json:"list,omitempty"`

	// Returns on API related errors
	Error Error `json:"error,omitempty"`
}

type Schedule struct {
	ID              int       `json:"id,omitempty"`
	Name            string    `json:"name,omitempty"`
	Comment         string    `json:"comment,omitempty"`
	BackupMode      string    `json:"backup_mode,omitempty"`
	Type            string    `json:"type,omitempty"`
	LimitCount      int       `json:"limit_count,omitempty"`
	LimitSizeMib    int       `json:"limit_size_mib,omitempty"`
	StartTime       StartTime `json:"start_time,omitempty"`
	BackupLocations []int     `json:"backup_locations,omitempty"`
	Clusters        []int     `json:"clusters,omitempty"`
	Disks           []int     `json:"disks,omitempty"`
	Nodes           []int     `json:"nodes,omitempty"`
	Hosts           []int     `json:"hosts,omitempty"`
	Presets         []int     `json:"presets,omitempty"`
	NextRun         string    `json:"next_run,omitempty"`
}

type StartTime struct {
	Day     int `json:"day,omitempty"`
	WeekDay int `json:"week_day,omitempty"`
	Hour    int `json:"hour,omitempty"`
	Minute  int `json:"minute,omitempty"`
}

type AffectedHostsCount struct {
	AffectedHostsCount int `json:"affected_hosts_count,omitempty"`

	// Returns on API related errors
	Error Error `json:"error,omitempty"`
}

type VXLans struct {
	LastNotify int     `json:"last_notify,omitempty"`
	List       []VXLan `json:"list,omitempty"`

	// Returns on API related errors
	Error Error `json:"error,omitempty"`
}

type VXLan struct {
	ID        int            `json:"id,omitempty"`
	Name      string         `json:"name,omitempty"`
	Comment   string         `json:"comment,omitempty"`
	Account   BaseAccount    `json:"account,omitempty"`
	Tag       int            `json:"tag,omitempty"`
	HostCount int            `json:"host_count,omitempty"`
	Clusters  []VXLanCluster `json:"clusters,omitempty"`
	Ipnets    []IPnet        `json:"ipnets,omitempty"`
	Ippool    int            `json:"ippool,omitempty"`

	// Returns on API related errors
	Error Error `json:"error,omitempty"`
}

type VXLanCluster struct {
	ID             int    `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	VxlanMode      string `json:"vxlan_mode,omitempty"`
	VxlanSettings  string `json:"vxlan_settings,omitempty"`
	VxlanHostCount int    `json:"vxlan_host_count,omitempty"`
}

type IPnet struct {
	ID             int    `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	Gateway        string `json:"gateway,omitempty"`
	VxlanHostCount int    `json:"vxlan_host_count,omitempty"`
}

type VXLanNew struct {
	Name     string          `json:"name,omitempty"`
	Comment  string          `json:"comment,omitempty"`
	Account  int             `json:"account,omitempty"`
	Ipnets   []VXLanNewIPnet `json:"ipnets,omitempty"`
	Clusters []int           `json:"clusters,omitempty"`
}

type VXLanNewIPnet struct {
	Name    string `json:"name,omitempty"`
	Gateway string `json:"gateway,omitempty"`
}

type VXLanIPs struct {
	LastNotify int       `json:"last_notify,omitempty"`
	List       []VXLanIP `json:"list,omitempty"`

	// Returns on API related errors
	Error Error `json:"error,omitempty"`
}

type VXLanIP struct {
	ID      int    `json:"id,omitempty"`
	IPAddr  string `json:"ip_addr,omitempty"`
	Gateway string `json:"gateway,omitempty"`
	Ipnet   string `json:"ipnet,omitempty"`
	Domain  string `json:"domain,omitempty"`
	State   string `json:"state,omitempty"`
	Comment string `json:"comment,omitempty"`
	Host    Host   `json:"host,omitempty"`
}

type VXLanIPNet struct {
	Name    string `json:"name,omitempty"`
	Gateway string `json:"gateway,omitempty"`
}

type TasksList struct {
	LastNotify int           `json:"last_notify,omitempty"`
	List       []TaskElement `json:"list,omitempty"`

	// Returns on API related errors
	Error Error `json:"error,omitempty"`
}

type TaskElement struct {
	ID          int    `json:"id,omitempty"`
	ConsulID    int    `json:"consul_id,omitempty"`
	DeferReason string `json:"defer_reason,omitempty"`
	Name        string `json:"name,omitempty"`
	Status      string `json:"status,omitempty"`

	// Returns on API related errors
	Error Error `json:"error,omitempty"`
}

type TaskInternalUpdate struct {
	Status string `json:"status,omitempty"`
	Output string `json:"output,omitempty"`
}

type TaskSyncRequest struct {
	Hosts []TaskHost `json:"hosts,omitempty"`
}
type TaskHost struct {
	Name   string      `json:"name,omitempty"`
	Events []TaskEvent `json:"events,omitempty"`
}

type TaskEvent struct {
	RelocateID int         `json:"relocate_id,omitempty"`
	Status     string      `json:"status,omitempty"`
	Params     interface{} `json:"params,omitempty"`
}

type LicenseInfo struct {
	HostLimit         int               `json:"host_limit,omitempty"`
	HostLimitExceeded bool              `json:"host_limit_exceeded,omitempty"`
	License           LicenseInfoParams `json:"license,omitempty"`
	NodeLimit         int               `json:"node_limit,omitempty"`
	NodeLimitExceeded bool              `json:"node_limit_exceeded,omitempty"`
	LastNotify        int               `json:"last_notify,omitempty"`

	// Returns on API related errors
	Error Error `json:"error,omitempty"`
}

type LicenseInfoParams struct {
	LicenseTTL   int    `json:"license_ttl,omitempty"`
	SignatureTTL int    `json:"signature_ttl,omitempty"`
	Valid        bool   `json:"valid,omitempty"`
	UserID       string `json:"user_id,omitempty"`
}

type LicenseSet struct {
	Signature string `json:"signature,omitempty"`
	Content   string `json:"content,omitempty"`
}

type Accounts struct {
	List []AccountElement `json:"list,omitempty"`

	// Returns on API related errors
	Error Error `json:"error,omitempty"`
}

type AccountElement struct {
	ID         int      `json:"id,omitempty"`
	Email      string   `json:"email,omitempty"`
	State      string   `json:"state,omitempty"`
	AuthSource string   `json:"auth_source,omitempty"`
	HostCount  int      `json:"host_count,omitempty"`
	CPUCount   int      `json:"cpu_count,omitempty"`
	RAMMib     int      `json:"ram_mib,omitempty"`
	HddMib     int      `json:"hdd_mib,omitempty"`
	Roles      []string `json:"roles,omitempty"`
	HostIds    []int    `json:"host_ids,omitempty"`
}

type AccountNew struct {
	Email       string `json:"email,omitempty"`
	FullName    string `json:"full_name,omitempty"`
	Password    string `json:"password,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	Role        string `json:"role,omitempty"`
}

type AccountSyncingParams struct {
	List   []AccountSyncingParansElement `json:"list,omitempty"`
	Delete []int                         `json:"delete,omitempty"`
}

type AccountSyncingParansElement struct {
	ID              int      `json:"id,omitempty"`
	Email           string   `json:"email,omitempty"`
	Roles           []string `json:"roles,omitempty"`
	State           string   `json:"state,omitempty"`
	AuthSource      string   `json:"auth_source,omitempty"`
	AdditionalProp1 struct {
	} `json:"additionalProp1,omitempty"`
}

type InviteUserParams struct {
	Email string `json:"email,omitempty"`
	Role  string `json:"role,omitempty"`
}

type UserParams struct {
	Roles       []string `json:"roles,omitempty"`
	FullName    string   `json:"full_name,omitempty"`
	Lang        string   `json:"lang,omitempty"`
	OldPassword string   `json:"old_password,omitempty"`
	Password    string   `json:"password,omitempty"`
}

type Storages struct {
	LastNotify int              `json:"last_notify,omitempty"`
	List       []StorageElement `json:"list,omitempty"`

	// Returns on API related errors
	Error Error `json:"error,omitempty"`
}

type StorageElement struct {
	ID            int           `json:"id,omitempty"`
	Name          string        `json:"name,omitempty"`
	Comment       string        `json:"comment,omitempty"`
	Type          string        `json:"type,omitempty"`
	StorageParams StorageParams `json:"storage_params,omitempty"`
	IsNetwork     bool          `json:"is_network,omitempty"`
	Clusters      []ClusterID   `json:"clusters,omitempty"`
	Tags          []Tag         `json:"tags,omitempty"`
}

type StorageParams struct {
	IPAddr   string `json:"ip_addr,omitempty"`
	SSHPort  int    `json:"ssh_port,omitempty"`
	Login    string `json:"login,omitempty"`
	Password string `json:"password,omitempty"`

	// Returns on API related errors
	Error Error `json:"error,omitempty"`
}

type StorageNew struct {
	Name          string        `json:"name,omitempty"`
	Comment       string        `json:"comment,omitempty"`
	Type          string        `json:"type,omitempty"`
	StorageParams StorageParams `json:"storage_params,omitempty"`
	IgnoreChecks  bool          `json:"ignore_checks,omitempty"`
	Tags          []int         `json:"tags,omitempty"`
}

type StorageUpdate struct {
	Name    string `json:"name,omitempty"`
	Comment string `json:"comment,omitempty"`
	Tags    []int  `json:"tags,omitempty"`
}

type UserLimits struct {
	HostCount       int `json:"host_count,omitempty"`
	ImageCount      int `json:"image_count,omitempty"`
	ImageCountTotal int `json:"image_count_total,omitempty"`
	CPUNumber       int `json:"cpu_number,omitempty"`
	CPUNumberTotal  int `json:"cpu_number_total,omitempty"`
	RAMMib          int `json:"ram_mib,omitempty"`
	RAMMibTotal     int `json:"ram_mib_total,omitempty"`
	HddMib          int `json:"hdd_mib,omitempty"`
	HddMibTotal     int `json:"hdd_mib_total,omitempty"`
	Ipv4Number      int `json:"ipv4_number,omitempty"`
	Ipv4NumberTotal int `json:"ipv4_number_total,omitempty"`
	VxlanCountTotal int `json:"vxlan_count_total,omitempty"`

	// Returns on API related errors
	Error Error `json:"error,omitempty"`
}

type KeyValue struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type Value struct {
	Value string `json:"value,omitempty"`
}

type TimeoutSeconds struct {
	TimeoutSeconds int `json:"timeout_seconds,omitempty"`
}
