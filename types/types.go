package types

// Auth

type Password struct {
	Password string `json:"password"`
}

type InviteUser struct {
	Password        string          `json:"password"`
	Lang            string          `json:"lang"`
	AdditionalProp1 AdditionalProp1 `json:"additionalProp1"`
}

type AdditionalProp1 struct {
}

type TokenLifetime struct {
	ExpiresAt   string `json:"expires_at"`
	Description string `json:"description"`
}

//Backups

type BackupList struct {
	LastNotify int                 `json:"last_notify"`
	List       []BackupListElement `json:"list"`
}

type BackupListElement struct {
	ID               int            `json:"id"`
	Name             string         `json:"name"`
	Account          BaseAccount    `json:"account"`
	Disk             Disk           `json:"disk"`
	ActualSizeMib    int            `json:"actual_size_mib"`
	EstimatedSizeMib int            `json:"estimated_size_mib"`
	Cluster          Cluster        `json:"cluster"`
	Host             int            `json:"host"`
	Schedule         int            `json:"schedule"`
	DateCreate       string         `json:"date_create"`
	AvailableUntil   string         `json:"available_until"`
	Comment          string         `json:"comment"`
	BackupLocation   BackupLocation `json:"backup_location"`
}

type BaseAccount struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

type Disk struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Cluster struct {
	ID               int    `json:"id"`
	DatacenterType   string `json:"datacenter_type"`
	ImageStoragePath string `json:"image_storage_path"`
	Name             string `json:"name"`
}

type BackupLocation struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type Backup struct {
	ID             int            `json:"id"`
	Name           string         `json:"name"`
	Account        string         `json:"account"`
	ForAll         bool           `json:"for_all"`
	Type           string         `json:"type"`
	Comment        string         `json:"comment"`
	State          string         `json:"state"`
	ExpandPart     string         `json:"expand_part"`
	BackupLocation BackupLocation `json:"backup_location"`
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

type Move struct {
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
	LastNotify int      `json:"last_notify"`
	List       []Backup `json:"list"`
}

type BackupLocationResponse struct {
	LastNotify int                     `json:"last_notify"`
	List       []BackupLocationElement `json:"list"`
}

type BackupLocationElement struct {
	ID               int              `json:"id"`
	Name             string           `json:"name"`
	Comment          string           `json:"comment"`
	State            string           `json:"state"`
	QuotaMib         int              `json:"quota_mib"`
	Type             string           `json:"type"`
	ConnectionParams ConnectionParams `json:"connection_params"`
	Clusters         []Cluster        `json:"clusters"`
}

type ConnectionParams struct {
}

type ID struct {
	ID int `json:"id"`
}

type BackupLocationParams struct {
	Name                string           `json:"name"`
	Comment             string           `json:"comment"`
	QuotaMib            int              `json:"quota_mib"`
	Type                string           `json:"type"`
	ConnectionParams    ConnectionParams `json:"connection_params"`
	Clusters            []int            `json:"clusters"`
	Schedules           []int            `json:"schedules"`
	SkipConnectionCheck bool             `json:"skip_connection_check"`
}

// Cluster

type ClusterListResponse struct {
	LastNotify int              `json:"last_notify"`
	List       []ClusterElement `json:"list"`
}

type ClusterElement struct {
	VirtualizationType     string           `json:"virtualization_type"`
	OsStoragePath          string           `json:"os_storage_path,omitempty"`
	ImageStoragePath       string           `json:"image_storage_path,omitempty"`
	CPUNumber              CPUNumber        `json:"cpu_number,omitempty"`
	HostCount              int              `json:"host_count"`
	AccountHostCount       int              `json:"account_host_count,omitempty"`
	HostPerNodeLimit       int              `json:"host_per_node_limit,omitempty"`
	Name                   string           `json:"name"`
	Comment                string           `json:"comment,omitempty"`
	RAMMib                 RAMMib           `json:"ram_mib,omitempty"`
	StorageMib             StorageMib       `json:"storage_mib,omitempty"`
	TimeZone               string           `json:"time_zone,omitempty"`
	State                  string           `json:"state,omitempty"`
	DataCenter             string           `json:"data_center,omitempty"`
	HostDistributionPolicy string           `json:"host_distribution_policy,omitempty"`
	HostFilter             []HostFilter     `json:"host_filter,omitempty"`
	NodeCount              int              `json:"node_count,omitempty"`
	ID                     int              `json:"id"`
	Overselling            int              `json:"overselling,omitempty"`
	BackupLocations        []BackupLocation `json:"backup_locations,omitempty"`
	DomainTemplate         string           `json:"domain_template,omitempty"`
	DatacenterParams       DatacenterParams `json:"datacenter_params,omitempty"`
	Storages               []Storage        `json:"storages,omitempty"`
	Storage                Storage          `json:"storage,omitempty"`
	DNSServers             []string         `json:"dns_servers,omitempty"`
	DcNetworks             []int            `json:"dc_networks,omitempty"`
	NodeNetwork            NodeNetwork      `json:"node_network,omitempty"`
	VxlanMode              string           `json:"vxlan_mode"`
	VxlanSettings          string           `json:"vxlan_settings"`
	Vxlans                 []Vxlan          `json:"vxlans"`
	DatacenterType         string           `json:"datacenter_type,omitempty"`
}

type Vxlan struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Tag  int    `json:"tag"`
}

type NodeNetwork struct {
	Timeout int    `json:"timeout"`
	Gateway string `json:"gateway"`
}

type CPUNumber struct {
	Total int `json:"total"`
	Used  int `json:"used"`
}

type RAMMib struct {
	Total int `json:"total"`
	Used  int `json:"used"`
}

type StorageMib struct {
	Total int `json:"total"`
	Used  int `json:"used"`
}

type HostFilter struct {
	Entity     string `json:"entity"`
	Expression string `json:"expression"`
}

type DatacenterParams struct {
	Gateway        string       `json:"gateway"`
	Mac            string       `json:"mac"`
	BgpCommunity   string       `json:"bgp_community"`
	BgpAs          int          `json:"bgp_as"`
	BgpSessions    []BgpSession `json:"bgp_sessions"`
	BgpCommunityV6 string       `json:"bgp_community_v6"`
	BgpAsV6        int          `json:"bgp_as_v6"`
	BgpSessionsV6  []BgpSession `json:"bgp_sessions_v6"`
}

type BgpSession struct {
	IP      string `json:"ip"`
	BgpAs   int    `json:"bgp_as"`
	Comment string `json:"comment"`
}

type Storage struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Type     string   `json:"type"`
	Enabled  bool     `json:"enabled"`
	IsMain   bool     `json:"is_main"`
	VirtPool VirtPool `json:"virt_pool"`
	Tags     []int    `json:"tags"`
}

type VirtPool struct {
	VMStoragePath   string `json:"vm_storage_path"`
	VolumeGroup     string `json:"volume_group"`
	ZfsPoolName     string `json:"zfs_pool_name"`
	RbdPoolName     string `json:"rbd_pool_name"`
	RbdUser         string `json:"rbd_user"`
	PlacementGroups int    `json:"placement_groups"`
	NetworkDiskPath string `json:"network_disk_path"`
}

type ClusterResponse struct {
	VirtualizationType     string           `json:"virtualization_type"`
	OsStoragePath          string           `json:"os_storage_path"`
	ImageStoragePath       string           `json:"image_storage_path"`
	CPUNumber              CPUNumber        `json:"cpu_number"`
	HostCount              int              `json:"host_count"`
	AccountHostCount       int              `json:"account_host_count"`
	HostPerNodeLimit       int              `json:"host_per_node_limit"`
	Name                   string           `json:"name"`
	Comment                string           `json:"comment"`
	RAMMib                 RAMMib           `json:"ram_mib"`
	StorageMib             StorageMib       `json:"storage_mib"`
	TimeZone               string           `json:"time_zone"`
	State                  string           `json:"state"`
	DataCenter             string           `json:"data_center"`
	HostDistributionPolicy string           `json:"host_distribution_policy"`
	HostFilter             []HostFilter     `json:"host_filter"`
	NodeCount              int              `json:"node_count"`
	ID                     int              `json:"id"`
	Overselling            int              `json:"overselling"`
	BackupLocations        []BackupLocation `json:"backup_locations"`
	DomainTemplate         string           `json:"domain_template"`
	DatacenterParams       DatacenterParams `json:"datacenter_params"`
	Storages               []Storage        `json:"storages"`
	Storage                Storage          `json:"storage"`
	DNSServers             []string         `json:"dns_servers"`
	DcNetworks             []int            `json:"dc_networks"`
	NodeNetwork            NodeNetwork      `json:"node_network"`
	VxlanMode              string           `json:"vxlan_mode"`
	VxlanSettings          string           `json:"vxlan_settings"`
	Vxlans                 []Vxlan          `json:"vxlans"`
}

type CreateClusterRequest struct {
	VirtualizationType     string           `json:"virtualization_type"`
	DatacenterType         string           `json:"datacenter_type"`
	Name                   string           `json:"name"`
	Comment                string           `json:"comment"`
	IsoEnabled             bool             `json:"iso_enabled"`
	ManageDiskEnabled      bool             `json:"manage_disk_enabled"`
	TimeZone               string           `json:"time_zone"`
	Interfaces             []Interface      `json:"interfaces"`
	Os                     []int            `json:"os"`
	BackupLocations        []int            `json:"backup_locations"`
	OsStoragePath          string           `json:"os_storage_path"`
	ImageStoragePath       string           `json:"image_storage_path"`
	Overselling            int              `json:"overselling"`
	HostDistributionPolicy string           `json:"host_distribution_policy"`
	HostFilter             []HostFilter     `json:"host_filter"`
	DomainTemplate         string           `json:"domain_template"`
	DatacenterParams       DatacenterParams `json:"datacenter_params"`
	DomainChangeAllowed    bool             `json:"domain_change_allowed"`
	ProxyEnabled           bool             `json:"proxy_enabled"`
	HostPerNodeLimit       int              `json:"host_per_node_limit"`
	DNSServers             []string         `json:"dns_servers"`
	NodeNetwork            NodeNetwork      `json:"node_network"`
	VxlanMode              string           `json:"vxlan_mode"`
	VxlanSettings          VxlanSettings    `json:"vxlan_settings"`
}

type VxlanSettings struct {
	BgpCommunity   string       `json:"bgp_community"`
	BgpAs          int          `json:"bgp_as"`
	BgpSessions    []BgpSession `json:"bgp_sessions"`
	BgpCommunityV6 string       `json:"bgp_community_v6"`
	BgpAsV6        int          `json:"bgp_as_v6"`
	BgpSessionsV6  []BgpSession `json:"bgp_sessions_v6"`
}

type Interface struct {
	Interface int   `json:"interface"`
	Ippool    []int `json:"ippool"`
}

type Tasks struct {
	ID    int   `json:"id"`
	Tasks []int `json:"tasks"`
}

type DeletedResponse struct {
	Deleted []ID `json:"deleted"`
}

type DCNetworksResponse struct {
	DcNetworks []int `json:"dc_networks"`
}

type QemuVersion struct {
	QemuVersion string `json:"qemu_version"`
}

type IPPoolRequest struct {
	Interfaces []Interface `json:"interfaces"`
}

type Name struct {
	Name string `json:"name"`
}

type SSHKeysRequest struct {
	SSHKeys []int `json:"ssh_keys"`
}

type AttachStorageToClusterRequest struct {
	VirtPool       VirtPool `json:"virt_pool"`
	IsMain         bool     `json:"is_main"`
	HddOverselling int      `json:"hdd_overselling"`
	IgnoreChecks   bool     `json:"ignore_checks"`
}

type StoragesTasks struct {
	ID            int   `json:"id"`
	Task          int   `json:"task"`
	StoragesTasks []int `json:"storages_tasks"`
}

type HDDOversellingRequest struct {
	Value float32 `json:"value"`
}

type HDDOversellingResponse struct {
	Value float32 `json:"value"`
}

type VXLanResponse struct {
	LastNotify int            `json:"last_notify"`
	List       []VxlanElement `json:"list"`
}

type VxlanElement struct {
	Account BaseAccount `json:"account"`
	ID      int         `json:"id"`
	Name    string      `json:"name"`
	Tag     int         `json:"tag"`
	Hosts   []Host      `json:"hosts"`
	Ranges  []string    `json:"ranges"`
}

type Host struct {
	ID   int      `json:"id"`
	Name string   `json:"name"`
	Node NodeHost `json:"node"`
}

type NodeHost struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type NodesVXLanResponse struct {
	LastNotify int          `json:"last_notify"`
	List       []NodesVxlan `json:"list"`
}

type HostVxlan struct {
	ID     int      `json:"id"`
	Name   string   `json:"name"`
	Ranges []string `json:"ranges"`
}

type NodesVxlan struct {
	Node   NodeHost            `json:"node"`
	Vxlans []NodesVxlanElement `json:"vxlans"`
}

type NodesVxlanElement struct {
	ID    int         `json:"id"`
	Name  string      `json:"name"`
	Tag   int         `json:"tag"`
	Hosts []HostVxlan `json:"hosts"`
}

type UsersVXLanResponse struct {
	LastNotify int                `json:"last_notify"`
	List       []UserVxlanElement `json:"list"`
}

type UserVxlanElement struct {
	Account BaseAccount `json:"account"`
	Hosts   []UserHost  `json:"hosts"`
}

type UserHost struct {
	ID     int         `json:"id"`
	Name   string      `json:"name"`
	Node   NodeHost    `json:"node"`
	Vxlans []UserVxlan `json:"vxlans"`
}

type UserVxlan struct {
	ID     int      `json:"id"`
	Name   string   `json:"name"`
	Tag    int      `json:"tag"`
	Ranges []string `json:"ranges"`
}

type SSHKeyResponse struct {
	LastNotify int      `json:"last_notify"`
	List       []SSHKey `json:"list"`
}

type SSHKey struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	SSHKey   string `json:"ssh_key"`
	Clusters []int  `json:"clusters"`
}

type SSHKeyRequest struct {
	Name     string `json:"name"`
	SSHKey   string `json:"ssh_key"`
	Clusters []int  `json:"clusters"`
}

// Disk

type DiskListResponse struct {
	LastNotify int           `json:"last_notify"`
	List       []DiskElement `json:"list"`
}

type DiskElement struct {
	ID         int         `json:"id"`
	Bus        string      `json:"bus"`
	Name       string      `json:"name"`
	ExpandPart string      `json:"expand_part"`
	TargetDev  string      `json:"target_dev"`
	SizeMib    int         `json:"size_mib"`
	SizeMibNew int         `json:"size_mib_new"`
	Account    BaseAccount `json:"account"`
	Storage    StorageID   `json:"storage"`
	Tags       []int       `json:"tags"`
	Host       HostParams  `json:"host"`
	Node       NodeHost    `json:"node"`
	Cluster    ClusterID   `json:"cluster"`
}

type ClusterID struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type HostParams struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	BootOrder int    `json:"boot_order"`
	IsMain    bool   `json:"is_main"`
	Bus       string `json:"bus"`
}

type StorageID struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type NewDiskRequest struct {
	Name       string `json:"name"`
	SizeMib    int    `json:"size_mib"`
	Node       int    `json:"node"`
	Account    int    `json:"account"`
	Storage    int    `json:"storage"`
	ExpandPart string `json:"expand_part"`
}

type UpdateDiskRequest struct {
	Name       string `json:"name"`
	Account    int    `json:"account"`
	SizeMib    int    `json:"size_mib"`
	ExpandPart string `json:"expand_part"`
	Pool       int    `json:"pool"`
	Defer      Action `json:"defer"`
}

type Action struct {
	Action string `json:"action"`
}

type MigrateDiskRequest struct {
	Storage int  `json:"storage"`
	Node    int  `json:"node"`
	Plain   bool `json:"plain"`
}

type BackupIDRequest struct {
	Backup int `json:"backup"`
}

type UpdateBootOrderRequest struct {
	Disks []DiskBootOrder `json:"disks"`
}

type DiskBootOrder struct {
	ID        int `json:"id"`
	BootOrder int `json:"boot_order"`
}

type DiskParams struct {
	IsMain bool   `json:"is_main"`
	Bus    string `json:"bus"`
}

// Host

type IP4 struct {
	IP string `json:"ip"`
}

type HostInterface struct {
	HostInterface string `json:"host_interface"`
	Mac           string `json:"mac"`
	NodeInterface int    `json:"node_interface"`
}

type HostNode struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	IPAddr string `json:"ip_addr"`
}

type HostDisk struct {
	ID         int `json:"id"`
	DiskMib    int `json:"disk_mib"`
	DiskMibNew int `json:"disk_mib_new"`
}

type FirewallRule struct {
	Action    string   `json:"action"`
	Direction string   `json:"direction"`
	Protocols []string `json:"protocols"`
	Portstart int      `json:"portstart"`
	Portend   int      `json:"portend"`
}

type HostsResponse struct {
	LastNotify int           `json:"last_notify"`
	List       []HostElement `json:"list"`
}

type HostElement struct {
	ExpandPart                string          `json:"expand_part"`
	ID                        int             `json:"id"`
	Name                      string          `json:"name"`
	IP4                       []IP4           `json:"ip4"`
	Interfaces                []HostInterface `json:"interfaces"`
	Node                      HostNode        `json:"node"`
	Cluster                   Cluster         `json:"cluster"`
	State                     string          `json:"state"`
	Domain                    string          `json:"domain"`
	Account                   BaseAccount     `json:"account"`
	Comment                   string          `json:"comment"`
	Disk                      HostDisk        `json:"disk"`
	DiskCount                 int             `json:"disk_count"`
	CPUNumber                 int             `json:"cpu_number"`
	CPUNumberNew              int             `json:"cpu_number_new"`
	RAMMib                    int             `json:"ram_mib"`
	RAMMibNew                 int             `json:"ram_mib_new"`
	NetBandwidthMbitps        int             `json:"net_bandwidth_mbitps"`
	NetBandwidthMbitpsChanged bool            `json:"net_bandwidth_mbitps_changed"`
	IPAutomation              string          `json:"ip_automation"`
	NetIsSynced               bool            `json:"net_is_synced"`
	Tags                      []string        `json:"tags"`
	OsName                    string          `json:"os_name"`
	OsGroup                   string          `json:"os_group"`
	Uptime                    int             `json:"uptime"`
	RescueMode                bool            `json:"rescue_mode"`
	IsoMounted                bool            `json:"iso_mounted"`
	CPUMode                   string          `json:"cpu_mode"`
	Nesting                   bool            `json:"nesting"`
	CPUCustomModel            string          `json:"cpu_custom_model"`
	CPUWeight                 int             `json:"cpu_weight"`
	IoWeight                  int             `json:"io_weight"`
	IoReadMbitps              int             `json:"io_read_mbitps"`
	IoWriteMbitps             int             `json:"io_write_mbitps"`
	IoReadIops                int             `json:"io_read_iops"`
	IoWriteIops               int             `json:"io_write_iops"`
	NetInMbitps               int             `json:"net_in_mbitps"`
	NetOutMbitps              int             `json:"net_out_mbitps"`
	NetWeight                 int             `json:"net_weight"`
	AntiSpoofing              bool            `json:"anti_spoofing"`
	Disabled                  bool            `json:"disabled"`
	TCPConnectionsIn          int             `json:"tcp_connections_in"`
	TCPConnectionsOut         int             `json:"tcp_connections_out"`
	ProcessNumber             int             `json:"process_number"`
	Vxlan                     Vxlan           `json:"vxlan"`
	FirewallRules             []FirewallRule  `json:"firewall_rules"`
	HasNonameIface            bool            `json:"has_noname_iface"`
	VM5Restrictions           VM5Restrictions `json:"vm5_restrictions"`
}

type VM5Restrictions struct {
	NetIfaceCount      bool `json:"net_iface_count"`
	NatOrExtra         bool `json:"nat_or_extra"`
	Ipv6               bool `json:"ipv6"`
	UnsupportedStorage bool `json:"unsupported_storage"`
	Iso                bool `json:"iso"`
	Snapshot           bool `json:"snapshot"`
}

type RecipeElement struct {
	Recipe       int            `json:"recipe"`
	RecipeParams []RecipeParams `json:"recipe_params"`
}

type RecipeParams struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type DiskHostRequest struct {
	Name       string `json:"name"`
	SizeMib    int    `json:"size_mib"`
	BootOrder  int    `json:"boot_order"`
	ExpandPart string `json:"expand_part"`
	Storage    int    `json:"storage"`
	Tags       []int  `json:"tags"`
}

type IPAddr struct {
	Name              string `json:"name"`
	IPPool            int    `json:"ip_pool"`
	IPNetwork         int    `json:"ip_network"`
	WithoutAllocation bool   `json:"without_allocation"`
}

type CustomInterface struct {
}

type HostVxlanRequest struct {
	ID         int `json:"id"`
	Ipv4Number int `json:"ipv4_number"`
	Ipnet      int `json:"ipnet"`
}

type NewHostRequest struct {
	Name                string             `json:"name"`
	Cluster             int                `json:"cluster"`
	Node                int                `json:"node"`
	Storage             int                `json:"storage"`
	Account             int                `json:"account"`
	Domain              string             `json:"domain"`
	Preset              int                `json:"preset"`
	Os                  int                `json:"os"`
	Image               int                `json:"image"`
	ExpandPart          string             `json:"expand_part"`
	RecipeList          []RecipeElement    `json:"recipe_list"`
	Recipe              int                `json:"recipe"`
	RecipeParams        []RecipeParams     `json:"recipe_params"`
	IgnoreRecipeFilters bool               `json:"ignore_recipe_filters"`
	Password            string             `json:"password"`
	RAMMib              int                `json:"ram_mib"`
	HddMib              int                `json:"hdd_mib"`
	Disks               []DiskHostRequest  `json:"disks"`
	CPUNumber           int                `json:"cpu_number"`
	NetBandwidthMbitps  int                `json:"net_bandwidth_mbitps"`
	IPAddr              IPAddr             `json:"ip_addr"`
	Ipv6Enabled         bool               `json:"ipv6_enabled"`
	Ipv6Pool            []int              `json:"ipv6_pool"`
	Ipv6Prefix          int                `json:"ipv6_prefix"`
	Ipv4Pool            []int              `json:"ipv4_pool"`
	Ipv4Number          int                `json:"ipv4_number"`
	Comment             string             `json:"comment"`
	Interfaces          []Interface        `json:"interfaces"`
	CustomInterfaces    []CustomInterface  `json:"custom_interfaces"`
	CPUMode             string             `json:"cpu_mode"`
	Nesting             bool               `json:"nesting"`
	CPUCustomModel      string             `json:"cpu_custom_model"`
	CPUWeight           int                `json:"cpu_weight"`
	IoWeight            int                `json:"io_weight"`
	IoReadMbitps        int                `json:"io_read_mbitps"`
	IoWriteMbitps       int                `json:"io_write_mbitps"`
	IoReadIops          int                `json:"io_read_iops"`
	IoWriteIops         int                `json:"io_write_iops"`
	NetInMbitps         int                `json:"net_in_mbitps"`
	NetOutMbitps        int                `json:"net_out_mbitps"`
	NetWeight           int                `json:"net_weight"`
	AntiSpoofing        bool               `json:"anti_spoofing"`
	TCPConnectionsIn    int                `json:"tcp_connections_in"`
	TCPConnectionsOut   int                `json:"tcp_connections_out"`
	ProcessNumber       int                `json:"process_number"`
	FirewallRules       []FirewallRule     `json:"firewall_rules"`
	SendEmailMode       string             `json:"send_email_mode"`
	Vxlan               []HostVxlanRequest `json:"vxlan"`
}

type RecipeTask struct {
	ID             int   `json:"id"`
	Task           int   `json:"task"`
	RecipeTaskList []int `json:"recipe_task_list"`
	RecipeTask     int   `json:"recipe_task"`
}

type HostResponse struct {
	ID           int    `json:"id"`
	InternalName string `json:"internal_name"`
	Node         int    `json:"node"`
	Name         string `json:"name"`
	State        string `json:"state"`
	CPUNumber    int    `json:"cpu_number"`
	RAMMib       int    `json:"ram_mib"`
	ExpandPart   string `json:"expand_part"`
	NetIsSynced  bool   `json:"net_is_synced"`
	IPAutomation string `json:"ip_automation"`
}

type Property struct {
}

type HostUpdateValues struct {
	Password         string   `json:"password"`
	VncPassword      string   `json:"vnc_password"`
	State            string   `json:"state"`
	StateUpdatedDate string   `json:"state_updated_date"`
	GuestAgent       bool     `json:"guest_agent"`
	VncPort          int      `json:"vnc_port"`
	Property         Property `json:"property"`
	StartDate        int      `json:"start_date"`
}
type HostUpdateRequest struct {
	Values            HostUpdateValues `json:"values"`
	Comment           string           `json:"comment"`
	Name              string           `json:"name"`
	HaRestoreOnFail   bool             `json:"ha_restore_on_fail"`
	HaRestorePriority int              `json:"ha_restore_priority"`
	ExpandPart        string           `json:"expand_part"`
}

type Account struct {
	Account int `json:"account"`
}

type ImageSize struct {
	ImageGib int `json:"image_gib"`
}

type HistoryElement struct {
	Name       string      `json:"name"`
	DateCreate string      `json:"date_create"`
	State      string      `json:"state"`
	Params     interface{} `json:"params"`
	User       string      `json:"user"`
}

type HostHistoryResponse struct {
	List []HistoryElement `json:"list"`
}

type IFaceElement struct {
	ID     int      `json:"id"`
	Name   string   `json:"name"`
	Bridge string   `json:"bridge"`
	IP     []string `json:"ip"`
	Mac    string   `json:"mac"`
	Model  string   `json:"model"`
	Vxlan  int      `json:"vxlan"`
}

type IFace struct {
	LastNotify int            `json:"last_notify"`
	List       []IFaceElement `json:"list"`
}

type IFaceParams struct {
	Name          string `json:"name"`
	Mac           string `json:"mac"`
	Bridge        string `json:"bridge"`
	IsMainNetwork bool   `json:"is_main_network"`
	Ippool        int    `json:"ippool"`
	IPCount       int    `json:"ip_count"`
}

type IFaceModel struct {
	Model string `json:"model"`
}

type AddIPToHostRequest struct {
	IPAddr      IPAddr             `json:"ip_addr"`
	Ipv6Pool    []int              `json:"ipv6_pool"`
	Ipv6Prefix  int                `json:"ipv6_prefix"`
	Ipv6Enabled bool               `json:"ipv6_enabled"`
	Ipv4Pool    []int              `json:"ipv4_pool"`
	Ipv4Number  int                `json:"ipv4_number"`
	Interfaces  []interface{}      `json:"interfaces"`
	Vxlan       []HostVxlanRequest `json:"vxlan"`
}

type IPAutomationType struct {
	IPAutomationType string `json:"ip_automation_type"`
}

type IPV4Host struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Interface int    `json:"interface"`
}

type IPV4InfoElement struct {
	ID               int      `json:"id"`
	IPAddr           string   `json:"ip_addr"`
	Domain           string   `json:"domain"`
	Gateway          string   `json:"gateway"`
	Mask             string   `json:"mask"`
	State            string   `json:"state"`
	Family           int      `json:"family"`
	Ippool           int      `json:"ippool"`
	Network          int      `json:"network"`
	Host             IPV4Host `json:"host"`
	ClusterInterface int      `json:"cluster_interface"`
	Vxlan            Vxlan    `json:"vxlan"`
}

type IPV4Info struct {
	LastNotify int               `json:"last_notify"`
	List       []IPV4InfoElement `json:"list"`
	Size       int               `json:"size"`
}

type IPV6Info struct {
	LastNotify     int  `json:"last_notify"`
	Ipv6Enabled    bool `json:"ipv6_enabled"`
	Ipv6PtrEnabled bool `json:"ipv6_ptr_enabled"`
}

type ISOTags struct {
	URL  string   `json:"url"`
	Tags []string `json:"tags"`
}

type Socket struct {
	Socket string `json:"socket"`
}

type Metadata struct {
	Metadata interface{} `json:"metadata"`
}

type HostMigrationStorage struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	AvailableMib int    `json:"available_mib"`
}

type NodeElement struct {
	ID                  int                    `json:"id"`
	Name                string                 `json:"name"`
	RAMMib              RAMMib                 `json:"ram_mib"`
	Storage             HostMigrationStorage   `json:"storage"`
	Storages            []HostMigrationStorage `json:"storages"`
	Priority            int                    `json:"priority"`
	State               string                 `json:"state"`
	Suitable            bool                   `json:"suitable"`
	HostCreationBlocked bool                   `json:"host_creation_blocked"`
	HostCount           int                    `json:"host_count"`
	HostLimit           int                    `json:"host_limit"`
}

type HostMigrateResponse struct {
	Nodes []NodeElement `json:"nodes"`
}

type DiskMigration struct {
	ID         int `json:"id"`
	DstStorage int `json:"dst_storage"`
}

type HostMigrateRequest struct {
	Plain bool            `json:"plain"`
	Node  int             `json:"node"`
	Disks []DiskMigration `json:"disks"`
}

type PTRUpdateRequest struct {
	Name   string `json:"name"`
	Domain string `json:"domain"`
}

type PTRIP struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Domain string `json:"domain"`
	Status string `json:"status"`
}

type PTRElement struct {
	Family  int     `json:"family"`
	Gateway string  `json:"gateway"`
	IP      []PTRIP `json:"ip"`
	IPCount int     `json:"ip_count"`
	Network string  `json:"network"`
	Prefix  int     `json:"prefix"`
}

type PTRResponse struct {
	List []PTRElement `json:"list"`
}

type Domain struct {
	Domain string `json:"domain"`
}

type HostRecipe struct {
	Recipe       int                `json:"recipe"`
	RecipeParams []HostRecipeParams `json:"recipe_params"`
}

type HostRecipeParams struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type HostReinstallRequest struct {
	Os                  int                `json:"os"`
	Image               int                `json:"image"`
	RecipeList          []HostRecipe       `json:"recipe_list"`
	Recipe              int                `json:"recipe"`
	RecipeParams        []HostRecipeParams `json:"recipe_params"`
	IgnoreRecipeFilters bool               `json:"ignore_recipe_filters"`
	SendEmailMode       string             `json:"send_email_mode"`
	Password            string             `json:"password"`
	Disk                int                `json:"disk"`
}

type RescueMode struct {
	RescueMode bool `json:"rescue_mode"`
}

type HostResourceUpdateRequest struct {
	CPUNumber          int            `json:"cpu_number"`
	RAMMib             int            `json:"ram_mib"`
	NetBandwidthMbitps int            `json:"net_bandwidth_mbitps"`
	CPUMode            string         `json:"cpu_mode"`
	Nesting            bool           `json:"nesting"`
	CPUCustomModel     string         `json:"cpu_custom_model"`
	CPUWeight          int            `json:"cpu_weight"`
	IoWeight           int            `json:"io_weight"`
	IoReadMbitps       int            `json:"io_read_mbitps"`
	IoWriteMbitps      int            `json:"io_write_mbitps"`
	IoReadIops         int            `json:"io_read_iops"`
	IoWriteIops        int            `json:"io_write_iops"`
	NetInMbitps        int            `json:"net_in_mbitps"`
	NetOutMbitps       int            `json:"net_out_mbitps"`
	NetWeight          int            `json:"net_weight"`
	AntiSpoofing       bool           `json:"anti_spoofing"`
	TCPConnectionsIn   int            `json:"tcp_connections_in"`
	TCPConnectionsOut  int            `json:"tcp_connections_out"`
	ProcessNumber      int            `json:"process_number"`
	FirewallRules      []FirewallRule `json:"firewall_rules"`
	Defer              Action         `json:"defer"`
}

type HostBackup struct {
	Backup int `json:"backup"`
}

type Recipient struct {
	Lang  string `json:"lang"`
	Email string `json:"email"`
}

type Recipe struct {
	Recipe              int            `json:"recipe"`
	Body                string         `json:"body"`
	RecipeParams        []RecipeParams `json:"recipe_params"`
	SendEmail           bool           `json:"send_email"`
	IgnoreRecipeFilters bool           `json:"ignore_recipe_filters"`
	Recipients          []Recipient    `json:"recipients"`
}

type VNCPort struct {
	Domain  string `json:"domain"`
	VncPort int    `json:"vnc_port"`
	VncPass string `json:"vnc_pass"`
}

type VNCPortUpdateRequest struct {
	List []VNCPort `json:"list"`
}

type IP struct {
	ID               int    `json:"id"`
	IPAddr           string `json:"ip_addr"`
	Domain           string `json:"domain"`
	Gateway          string `json:"gateway"`
	Mask             string `json:"mask"`
	State            string `json:"state"`
	Family           int    `json:"family"`
	Ippool           int    `json:"ippool"`
	Network          int    `json:"network"`
	Host             IPHost `json:"host"`
	ClusterInterface int    `json:"cluster_interface"`
}

type IPHost struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Interface int    `json:"interface"`
}

type IPList struct {
	LastNotify int  `json:"last_notify"`
	List       []IP `json:"list"`
	Size       int  `json:"size"`
}

type UpdateIPResponseValues struct {
	Domain  string `json:"domain"`
	Gateway string `json:"gateway"`
	IPAddr  string `json:"ip_addr"`
	Netmask string `json:"netmask"`
}
type UpdateIPResponse struct {
	Values UpdateIPResponseValues `json:"values"`
}

type ScheduleElement struct {
	Name           string      `json:"name"`
	Handler        string      `json:"handler"`
	Service        string      `json:"service"`
	Method         string      `json:"method"`
	InstanceID     int         `json:"instance_id"`
	URLQueryParams interface{} `json:"url_query_params"`
	PostParams     interface{} `json:"post_params"`
	CronExpression string      `json:"cron_expression"`
}

type ScheduleListResponse struct {
	SheduleList []ScheduleElement `json:"shedule_list"`
}

// Images

type Image struct {
	Name    string `json:"name"`
	Account int    `json:"account"`
	Comment string `json:"comment"`
	ForAll  bool   `json:"for_all"`
}

type ImagesListResponse struct {
	LastNotify int                         `json:"last_notify"`
	List       []ImagesListResponseElement `json:"list"`
}

type ImagesListResponseElement struct {
	ID      int                      `json:"id"`
	Name    string                   `json:"name"`
	Account string                   `json:"account"`
	ForAll  bool                     `json:"for_all"`
	Type    string                   `json:"type"`
	Comment string                   `json:"comment"`
	State   string                   `json:"state"`
	Nodes   []ImagesListResponseNode `json:"nodes"`
}

type ImagesListResponseNode struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	IPAddr  string `json:"ip_addr"`
	SSHPort int    `json:"ssh_port"`
}

type ImageResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	State string `json:"state"`
}

type ImageState struct {
	State string `json:"state"`
}
