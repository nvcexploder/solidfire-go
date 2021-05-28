package types

type AddAccountRequest struct {
	Username        string      `json:"username"`
	InitiatorSecret CHAPSecret  `json:"initiatorSecret,omitempty"`
	TargetSecret    CHAPSecret  `json:"targetSecret,omitempty"`
	Attributes      interface{} `json:"attributes,omitempty"`
}

type AddClusterAdminRequest struct {
	Username   string      `json:"username"`
	Password   string      `json:"password"`
	Access     []string    `json:"access"`
	AcceptEula bool        `json:"acceptEula,omitempty"`
	Attributes interface{} `json:"attributes,omitempty"`
}

type AddDrivesRequest struct {
	Drives             []NewDrive `json:"drives"`
	ForceDuringUpgrade bool       `json:"forceDuringUpgrade,omitempty"`
}

type AddInitiatorsToVolumeAccessGroupRequest struct {
	VolumeAccessGroupID int64    `json:"volumeAccessGroupID"`
	Initiators          []string `json:"initiators"`
}

type AddLdapClusterAdminRequest struct {
	Username   string      `json:"username"`
	Access     []string    `json:"access"`
	AcceptEula bool        `json:"acceptEula,omitempty"`
	Attributes interface{} `json:"attributes,omitempty"`
}

type AddNodesRequest struct {
	PendingNodes []int64 `json:"pendingNodes"`
	AutoInstall  bool    `json:"autoInstall,omitempty"`
}

type AddVirtualNetworkRequest struct {
	VirtualNetworkTag int64          `json:"virtualNetworkTag"`
	Name              string         `json:"name"`
	AddressBlocks     []AddressBlock `json:"addressBlocks"`
	Netmask           string         `json:"netmask"`
	Svip              string         `json:"svip"`
	Gateway           string         `json:"gateway,omitempty"`
	Namespace         bool           `json:"namespace,omitempty"`
	Attributes        interface{}    `json:"attributes,omitempty"`
}

type AddVolumesToVolumeAccessGroupRequest struct {
	VolumeAccessGroupID int64   `json:"volumeAccessGroupID"`
	Volumes             []int64 `json:"volumes"`
}

type CancelCloneRequest struct {
	CloneID int64 `json:"cloneID"`
}

type CancelGroupCloneRequest struct {
	GroupCloneID int64 `json:"groupCloneID"`
}

type ClearClusterFaultsRequest struct {
	FaultTypes string `json:"faultTypes,omitempty"`
}

type CloneMultipleVolumesRequest struct {
	Volumes         []CloneMultipleVolumeParams `json:"volumes"`
	Access          string                      `json:"access,omitempty"`
	GroupSnapshotID int64                       `json:"groupSnapshotID,omitempty"`
	NewAccountID    int64                       `json:"newAccountID,omitempty"`
}

type CloneVolumeRequest struct {
	VolumeID     int64       `json:"volumeID"`
	Name         string      `json:"name"`
	NewAccountID int64       `json:"newAccountID,omitempty"`
	NewSize      int64       `json:"newSize,omitempty"`
	Access       string      `json:"access,omitempty"`
	SnapshotID   int64       `json:"snapshotID,omitempty"`
	Attributes   interface{} `json:"attributes,omitempty"`
	Enable512e   bool        `json:"enable512e,omitempty"`
}

type CompleteClusterPairingRequest struct {
	ClusterPairingKey string `json:"clusterPairingKey"`
}

type CompleteVolumePairingRequest struct {
	VolumePairingKey string `json:"volumePairingKey"`
	VolumeID         int64  `json:"volumeID"`
}

type CopyVolumeRequest struct {
	VolumeID    int64 `json:"volumeID"`
	DstVolumeID int64 `json:"dstVolumeID"`
	SnapshotID  int64 `json:"snapshotID,omitempty"`
}

type CreateBackupTargetRequest struct {
	Name       string      `json:"name"`
	Attributes interface{} `json:"attributes,omitempty"`
}

type CreateClusterRequest struct {
	AcceptEula bool        `json:"acceptEula,omitempty"`
	Mvip       string      `json:"mvip"`
	Svip       string      `json:"svip"`
	RepCount   int64       `json:"repCount"`
	Username   string      `json:"username"`
	Password   string      `json:"password"`
	Nodes      []string    `json:"nodes"`
	Attributes interface{} `json:"attributes,omitempty"`
}

type CreateGroupSnapshotRequest struct {
	Volumes                 []int64     `json:"volumes"`
	Name                    string      `json:"name,omitempty"`
	EnableRemoteReplication bool        `json:"enableRemoteReplication,omitempty"`
	Retention               string      `json:"retention,omitempty"`
	Attributes              interface{} `json:"attributes,omitempty"`
}

type CreateInitiatorsRequest struct {
	Initiators []CreateInitiator `json:"initiators"`
}

type CreateScheduleRequest struct {
	Schedule Schedule `json:"schedule"`
}

type CreateSnapshotRequest struct {
	Attributes              interface{} `json:"attributes,omitempty"`
	EnableRemoteReplication bool        `json:"enableRemoteReplication,omitempty"`
	EnsureSerialCreate      bool        `json:"enableSerialCreation,omitempty"`
	ExpirationTime          string      `json:"expirationTime,omitempty"`
	Name                    string      `json:"name,omitempty"`
	Retention               string      `json:"retention,omitempty"`
	SnapMirrorLabel         string      `json:"snapMirrorLabel,omitempty"`
	SnapshotID              int64       `json:"snapshotID,omitempty"`
	VolumeID                int64       `json:"volumeID"`
}

type CreateStorageContainerRequest struct {
	Name            string `json:"name"`
	InitiatorSecret string `json:"initiatorSecret,omitempty"`
	TargetSecret    string `json:"targetSecret,omitempty"`
	AccountID       int64  `json:"accountID,omitempty"`
}

type CreateSupportBundleRequest struct {
	BundleName string `json:"bundleName,omitempty"`
	ExtraArgs  string `json:"extraArgs,omitempty"`
	TimeoutSec int64  `json:"timeoutSec,omitempty"`
}

type CreateVolumeAccessGroupRequest struct {
	Name               string      `json:"name"`
	Initiators         []string    `json:"initiators,omitempty"`
	Volumes            []int64     `json:"volumes,omitempty"`
	VirtualNetworkID   []int64     `json:"virtualNetworkID,omitempty"`
	VirtualNetworkTags []int64     `json:"virtualNetworkTags,omitempty"`
	Attributes         interface{} `json:"attributes,omitempty"`
}

type CreateVolumeRequest struct {
	Name                        string      `json:"name"`
	AccountID                   int64       `json:"accountID"`
	TotalSize                   int64       `json:"totalSize"`
	Enable512e                  bool        `json:"enable512e"`
	Qos                         QoS         `json:"qos,omitempty"`
	QosPolicyID                 int64       `json:"qosPolicyID,omitempty"`
	Access                      string      `json:"access,omitempty"`
	AssociateWithQoSPolicy      bool        `json:"associateWithQosPolicy,omitempty"`
	EnableSnapMirrorReplication bool        `json:"enableSnapMirrorReplication,omitempty"`
	FifoSize                    int64       `json:"fifoSize,omitempty"`
	MinFifoSize                 int64       `json:"minFifoSize,omitempty"`
	Attributes                  interface{} `json:"attributes,omitempty"`
}

type DeleteGroupSnapshotRequest struct {
	GroupSnapshotID int64 `json:"groupSnapshotID"`
	SaveMembers     bool  `json:"saveMembers"`
}

type DeleteInitiatorsRequest struct {
	Initiators []int64 `json:"initiators"`
}

type DeleteSnapshotRequest struct {
	SnapshotID int64 `json:"snapshotID"`
}

type DeleteStorageContainersRequest struct {
	StorageContainerIDs []string `json:"storageContainerIDs"`
}

type DeleteVolumeAccessGroupRequest struct {
	VolumeAccessGroupID int64 `json:"volumeAccessGroupID"`
}

type DeleteVolumeRequest struct {
	VolumeID int64 `json:"volumeID"`
}

type DeleteVolumesRequest struct {
	AccountIDs           []int64 `json:"accountIDs,omitempty"`
	VolumeAccessGroupIDs []int64 `json:"volumeAccessGroupIDs,omitempty"`
	VolumeIDs            []int64 `json:"volumeIDs,omitempty"`
}

type EnableFeatureRequest struct {
	Feature string `json:"feature"`
}

type EnableLdapAuthenticationRequest struct {
	AuthType                string   `json:"authType,omitempty"`
	GroupSearchBaseDN       string   `json:"groupSearchBaseDN,omitempty"`
	GroupSearchCustomFilter string   `json:"groupSearchCustomFilter,omitempty"`
	GroupSearchType         string   `json:"groupSearchType,omitempty"`
	SearchBindDN            string   `json:"searchBindDN,omitempty"`
	SearchBindPassword      string   `json:"searchBindPassword,omitempty"`
	ServerURIs              []string `json:"serverURIs"`
	UserDNTemplate          string   `json:"userDNTemplate,omitempty"`
	UserSearchBaseDN        string   `json:"userSearchBaseDN,omitempty"`
	UserSearchFilter        string   `json:"userSearchFilter,omitempty"`
}

type EnableSnmpRequest struct {
	SnmpV3Enabled bool `json:"snmpV3Enabled"`
}

type GetAccountByIDRequest struct {
	AccountID int64 `json:"accountID"`
}

type GetAccountByNameRequest struct {
	Username string `json:"username"`
}

type GetAccountEfficiencyRequest struct {
	AccountID int64 `json:"accountID"`
}

type GetAsyncResultRequest struct {
	AsyncHandle int64 `json:"asyncHandle"`
	KeepResult  bool  `json:"keepResult,omitempty"`
}

type GetBackupTargetRequest struct {
	BackupTargetID int64 `json:"backupTargetID"`
}

type GetClusterHardwareInfoRequest struct {
	Type string `json:"type,omitempty"`
}

type GetClusterStateRequest struct {
	Force bool `json:"force"`
}

type GetDriveHardwareInfoRequest struct {
	DriveID int64 `json:"driveID"`
}

type GetDriveStatsRequest struct {
	DriveID int64 `json:"driveID"`
}

type GetFeatureStatusRequest struct {
	Feature string `json:"feature,omitempty"`
}

type GetIpmiConfigRequest struct {
	ChassisType string `json:"chassisType,omitempty"`
}

type GetNodeHardwareInfoRequest struct {
	NodeID int64 `json:"nodeID"`
}

type GetNodeStatsRequest struct {
	NodeID int64 `json:"nodeID"`
}

type GetNvramInfoRequest struct {
	Force bool `json:"force,omitempty"`
}

type GetScheduleRequest struct {
	ScheduleID int64 `json:"scheduleID"`
}

type GetStorageContainerEfficiencyRequest struct {
	StorageContainerID string `json:"storageContainerID"`
}

type GetVolumeAccessGroupEfficiencyRequest struct {
	VolumeAccessGroupID int64 `json:"volumeAccessGroupID"`
}

type GetVolumeAccessGroupLunAssignmentsRequest struct {
	VolumeAccessGroupID int64 `json:"volumeAccessGroupID"`
}

type GetVolumeEfficiencyRequest struct {
	VolumeID int64 `json:"volumeID"`
}

type GetVolumeStatsRequest struct {
	VolumeID int64 `json:"volumeID"`
}

type InvokeSFApiRequest struct {
	Method     string      `json:"method"`
	Parameters interface{} `json:"parameters,omitempty"`
}

type ListAccountsRequest struct {
	StartAccountID           int64 `json:"startAccountID,omitempty"`
	Limit                    int64 `json:"limit,omitempty"`
	IncludeStorageContainers bool  `json:"includeStorageContainers,omitempty"`
}

type ListActivePairedVolumesRequest struct {
	StartVolumeID int64 `json:"startVolumeID,omitempty"`
	Limit         int64 `json:"limit,omitempty"`
}

type ListActiveVolumesRequest struct {
	StartVolumeID         int64 `json:"startVolumeID,omitempty"`
	Limit                 int64 `json:"limit,omitempty"`
	IncludeVirtualVolumes bool  `json:"includeVirtualVolumes,omitempty"`
}

type ListAsyncResultsRequest struct {
	AsyncResultTypes []string `json:"asyncResultTypes,omitempty"`
}

type ListClusterFaultsRequest struct {
	BestPractices bool   `json:"bestPractices,omitempty"`
	FaultTypes    string `json:"faultTypes,omitempty"`
}

type ListDeletedVolumesRequest struct {
	IncludeVirtualVolumes bool `json:"includeVirtualVolumes,omitempty"`
}

type ListDriveHardwareRequest struct {
	Force bool `json:"force"`
}

type ListDriveStatsRequest struct {
	Drives []int64 `json:"drives,omitempty"`
}

type ListEventsRequest struct {
	MaxEvents    int64 `json:"maxEvents,omitempty"`
	StartEventID int64 `json:"startEventID,omitempty"`
	EndEventID   int64 `json:"endEventID,omitempty"`
}

type ListGroupSnapshotsRequest struct {
	Volumes         []int64 `json:"volumes,omitempty"`
	GroupSnapshotID int64   `json:"groupSnapshotID,omitempty"`
}

type ListInitiatorsRequest struct {
	StartInitiatorID int64   `json:"startInitiatorID,omitempty"`
	Limit            int64   `json:"limit,omitempty"`
	Initiators       []int64 `json:"initiators,omitempty"`
}

type ListProtocolEndpointsRequest struct {
	ProtocolEndpointIDs []string `json:"protocolEndpointIDs,omitempty"`
}

type ListSnapshotsRequest struct {
	VolumeID   int64 `json:"volumeID,omitempty"`
	SnapshotID int64 `json:"snapshotID,omitempty"`
}

type ListStorageContainersRequest struct {
	StorageContainerIDs []string `json:"storageContainerIDs,omitempty"`
}

type ListVirtualNetworksRequest struct {
	VirtualNetworkID   int64   `json:"virtualNetworkID,omitempty"`
	VirtualNetworkTag  int64   `json:"virtualNetworkTag,omitempty"`
	VirtualNetworkIDs  []int64 `json:"virtualNetworkIDs,omitempty"`
	VirtualNetworkTags []int64 `json:"virtualNetworkTags,omitempty"`
}

type ListVirtualVolumeBindingsRequest struct {
	VirtualVolumeBindingIDs []int64 `json:"virtualVolumeBindingIDs,omitempty"`
}

type ListVirtualVolumeHostsRequest struct {
	VirtualVolumeHostIDs []string `json:"virtualVolumeHostIDs,omitempty"`
}

type ListVirtualVolumeTasksRequest struct {
	VirtualVolumeTaskIDs []string `json:"virtualVolumeTaskIDs,omitempty"`
}

type ListVirtualVolumesRequest struct {
	Details              bool     `json:"details,omitempty"`
	Limit                int64    `json:"limit,omitempty"`
	Recursive            bool     `json:"recursive,omitempty"`
	StartVirtualVolumeID string   `json:"startVirtualVolumeID,omitempty"`
	VirtualVolumeIDs     []string `json:"virtualVolumeIDs,omitempty"`
}

type ListVolumeAccessGroupsRequest struct {
	StartVolumeAccessGroupID int64   `json:"startVolumeAccessGroupID,omitempty"`
	Limit                    int64   `json:"limit,omitempty"`
	VolumeAccessGroups       []int64 `json:"volumeAccessGroups,omitempty"`
}

type ListVolumeStatsByAccountRequest struct {
	Accounts              []int64 `json:"accounts,omitempty"`
	IncludeVirtualVolumes bool    `json:"includeVirtualVolumes,omitempty"`
}

type ListVolumeStatsByVirtualVolumeRequest struct {
	VirtualVolumeIDs []string `json:"virtualVolumeIDs,omitempty"`
}

type ListVolumeStatsByVolumeAccessGroupRequest struct {
	VolumeAccessGroups    []int64 `json:"volumeAccessGroups,omitempty"`
	IncludeVirtualVolumes bool    `json:"includeVirtualVolumes,omitempty"`
}

type ListVolumeStatsByVolumeRequest struct {
	IncludeVirtualVolumes bool `json:"includeVirtualVolumes,omitempty"`
}

type ListVolumeStatsRequest struct {
	VolumeIDs []int64 `json:"volumeIDs,omitempty"`
}

type ListVolumesForAccountRequest struct {
	AccountID             int64 `json:"accountID"`
	StartVolumeID         int64 `json:"startVolumeID,omitempty"`
	Limit                 int64 `json:"limit,omitempty"`
	IncludeVirtualVolumes bool  `json:"includeVirtualVolumes,omitempty"`
}

type ListVolumesRequest struct {
	StartVolumeID         int64   `json:"startVolumeID,omitempty"`
	Limit                 int64   `json:"limit,omitempty"`
	VolumeStatus          string  `json:"volumeStatus,omitempty"`
	Accounts              []int64 `json:"accounts,omitempty"`
	IsPaired              bool    `json:"isPaired,omitempty"`
	VolumeIDs             []int64 `json:"volumeIDs,omitempty"`
	VolumeName            string  `json:"volumeName,omitempty"`
	IncludeVirtualVolumes bool    `json:"includeVirtualVolumes,omitempty"`
}

type ModifyAccountRequest struct {
	AccountID       int64       `json:"accountID"`
	Username        string      `json:"username,omitempty"`
	Status          string      `json:"status,omitempty"`
	InitiatorSecret CHAPSecret  `json:"initiatorSecret,omitempty"`
	TargetSecret    CHAPSecret  `json:"targetSecret,omitempty"`
	Attributes      interface{} `json:"attributes,omitempty"`
}

type ModifyBackupTargetRequest struct {
	BackupTargetID int64       `json:"backupTargetID"`
	Name           string      `json:"name,omitempty"`
	Attributes     interface{} `json:"attributes,omitempty"`
}

type ModifyClusterAdminRequest struct {
	ClusterAdminID int64       `json:"clusterAdminID"`
	Password       string      `json:"password,omitempty"`
	Access         []string    `json:"access,omitempty"`
	Attributes     interface{} `json:"attributes,omitempty"`
}

type ModifyClusterFullThresholdRequest struct {
	Stage2AwareThreshold           int64 `json:"stage2AwareThreshold,omitempty"`
	Stage3BlockThresholdPercent    int64 `json:"stage3BlockThresholdPercent,omitempty"`
	MaxMetadataOverProvisionFactor int64 `json:"maxMetadataOverProvisionFactor,omitempty"`
}

type ModifyGroupSnapshotRequest struct {
	GroupSnapshotID         int64  `json:"groupSnapshotID"`
	ExpirationTime          string `json:"expirationTime,omitempty"`
	EnableRemoteReplication bool   `json:"enableRemoteReplication,omitempty"`
}

type ModifyInitiatorsRequest struct {
	Initiators []ModifyInitiator `json:"initiators"`
}

type ModifyScheduleRequest struct {
	Schedule Schedule `json:"schedule"`
}

type ModifySnapshotRequest struct {
	SnapshotID              int64  `json:"snapshotID"`
	ExpirationTime          string `json:"expirationTime,omitempty"`
	EnableRemoteReplication bool   `json:"enableRemoteReplication,omitempty"`
	Name                    string `json:"name,omitempty"`
	SnapMirrorLabel         string `json:"snapMirrorLabel,omitempty"`
}

type ModifyStorageContainerRequest struct {
	StorageContainerID string `json:"storageContainerID"`
	InitiatorSecret    string `json:"initiatorSecret,omitempty"`
	TargetSecret       string `json:"targetSecret,omitempty"`
}

type ModifyVirtualNetworkRequest struct {
	VirtualNetworkID  int64          `json:"virtualNetworkID,omitempty"`
	VirtualNetworkTag int64          `json:"virtualNetworkTag,omitempty"`
	Name              string         `json:"name,omitempty"`
	AddressBlocks     []AddressBlock `json:"addressBlocks,omitempty"`
	Netmask           string         `json:"netmask,omitempty"`
	Svip              string         `json:"svip,omitempty"`
	Gateway           string         `json:"gateway,omitempty"`
	Namespace         bool           `json:"namespace,omitempty"`
	Attributes        interface{}    `json:"attributes,omitempty"`
}

type ModifyVolumeAccessGroupLunAssignmentsRequest struct {
	VolumeAccessGroupID int64           `json:"volumeAccessGroupID"`
	LunAssignments      []LunAssignment `json:"lunAssignments"`
}

type ModifyVolumeAccessGroupRequest struct {
	VolumeAccessGroupID    int64       `json:"volumeAccessGroupID"`
	VirtualNetworkID       []int64     `json:"virtualNetworkID,omitempty"`
	VirtualNetworkTags     []int64     `json:"virtualNetworkTags,omitempty"`
	Name                   string      `json:"name,omitempty"`
	Initiators             []string    `json:"initiators,omitempty"`
	Volumes                []int64     `json:"volumes,omitempty"`
	DeleteOrphanInitiators bool        `json:"deleteOrphanInitiators,omitempty"`
	Attributes             interface{} `json:"attributes,omitempty"`
}

type ModifyVolumePairRequest struct {
	VolumeID     int64  `json:"volumeID"`
	PausedManual bool   `json:"pausedManual,omitempty"`
	Mode         string `json:"mode,omitempty"`
	PauseLimit   int64  `json:"pauseLimit,omitempty"`
}

type ModifyVolumeRequest struct {
	VolumeID                    int64       `json:"volumeID"`
	AccountID                   int64       `json:"accountID,omitempty"`
	Access                      string      `json:"access,omitempty"`
	Qos                         QoS         `json:"qos,omitempty"`
	QosPolicyID                 int64       `json:"qosPolicyID,omitempty"`
	FifoSize                    int64       `json:"fifoSize,omitempty"`
	MinFifoSize                 int64       `json:"minFifoSize,omitempty"`
	TotalSize                   int64       `json:"totalSize,omitempty"`
	Attributes                  interface{} `json:"attributes,omitempty"`
	AssociateWithQoSPolicy      bool        `json:"associateWithQosPolicy,omitempty"`
	EnableSnapMirrorReplication bool        `json:"enableSnapMirrorReplication,omitempty"`
	Mode                        string      `json:"mode,omitempty"`
}

type ModifyVolumesRequest struct {
	VolumeIDs  []int64     `json:"volumeIDs"`
	AccountID  int64       `json:"accountID,omitempty"`
	Access     string      `json:"access,omitempty"`
	Qos        QoS         `json:"qos,omitempty"`
	TotalSize  int64       `json:"totalSize,omitempty"`
	Attributes interface{} `json:"attributes,omitempty"`
}

type PurgeDeletedVolumeRequest struct {
	VolumeID int64 `json:"volumeID"`
}

type PurgeDeletedVolumesRequest struct {
	VolumeIDs            []int64 `json:"volumeIDs,omitempty"`
	AccountIDs           []int64 `json:"accountIDs,omitempty"`
	VolumeAccessGroupIDs []int64 `json:"volumeAccessGroupIDs,omitempty"`
}

type RemoveAccountRequest struct {
	AccountID int64 `json:"accountID"`
}

type RemoveBackupTargetRequest struct {
	BackupTargetID int64 `json:"backupTargetID"`
}

type RemoveClusterAdminRequest struct {
	ClusterAdminID int64 `json:"clusterAdminID"`
}

type RemoveClusterPairRequest struct {
	ClusterPairID int64 `json:"clusterPairID"`
}

type RemoveDrivesRequest struct {
	Drives             []int64 `json:"drives"`
	ForceDuringUpgrade bool    `json:"forceDuringUpgrade,omitempty"`
}

type RemoveInitiatorsFromVolumeAccessGroupRequest struct {
	VolumeAccessGroupID    int64    `json:"volumeAccessGroupID"`
	Initiators             []string `json:"initiators"`
	DeleteOrphanInitiators bool     `json:"deleteOrphanInitiators,omitempty"`
}

type RemoveNodesRequest struct {
	Nodes []int64 `json:"nodes"`
}

type RemoveVirtualNetworkRequest struct {
	VirtualNetworkID  int64 `json:"virtualNetworkID,omitempty"`
	VirtualNetworkTag int64 `json:"virtualNetworkTag,omitempty"`
}

type RemoveVolumePairRequest struct {
	VolumeID int64 `json:"volumeID"`
}

type RemoveVolumesFromVolumeAccessGroupRequest struct {
	VolumeAccessGroupID int64   `json:"volumeAccessGroupID"`
	Volumes             []int64 `json:"volumes"`
}

type ResetDrivesRequest struct {
	Drives string `json:"drives"`
	Force  bool   `json:"force"`
}

type ResetNodeRequest struct {
	Build   string `json:"build"`
	Force   bool   `json:"force"`
	Options string `json:"options,omitempty"`
	Reboot  bool   `json:"reboot,omitempty"`
}

type RestartNetworkingRequest struct {
	Force bool `json:"force"`
}

type RestartServicesRequest struct {
	Force   bool   `json:"force"`
	Service string `json:"service,omitempty"`
	Action  string `json:"action,omitempty"`
}

type RestoreDeletedVolumeRequest struct {
	VolumeID int64 `json:"volumeID"`
}

type RollbackToGroupSnapshotRequest struct {
	GroupSnapshotID  int64       `json:"groupSnapshotID"`
	SaveCurrentState bool        `json:"saveCurrentState"`
	Name             string      `json:"name,omitempty"`
	Attributes       interface{} `json:"attributes,omitempty"`
}

type RollbackToSnapshotRequest struct {
	VolumeID         int64       `json:"volumeID"`
	SnapshotID       int64       `json:"snapshotID"`
	SaveCurrentState bool        `json:"saveCurrentState"`
	Name             string      `json:"name,omitempty"`
	Attributes       interface{} `json:"attributes,omitempty"`
}

type SecureEraseDrivesRequest struct {
	Drives []int64 `json:"drives"`
}

type SetClusterConfigRequest struct {
	Cluster ClusterConfig `json:"cluster"`
}

type SetConfigRequest struct {
	Config Config `json:"config"`
}

type SetDefaultQoSRequest struct {
	MinIOPS   int64 `json:"minIOPS,omitempty"`
	MaxIOPS   int64 `json:"maxIOPS,omitempty"`
	BurstIOPS int64 `json:"burstIOPS,omitempty"`
}

type SetLoginSessionInfoRequest struct {
	Timeout string `json:"timeout"`
}

type SetNetworkConfigRequest struct {
	Network NetworkParams `json:"network"`
}

type SetNtpInfoRequest struct {
	Servers         []string `json:"servers"`
	Broadcastclient bool     `json:"broadcastclient,omitempty"`
}

type SetRemoteLoggingHostsRequest struct {
	RemoteHosts []LoggingServer `json:"remoteHosts"`
}

type SetSnmpACLRequest struct {
	Networks []SnmpNetwork   `json:"networks"`
	UsmUsers []SnmpV3UsmUser `json:"usmUsers"`
}

type SetSnmpInfoRequest struct {
	Networks      []SnmpNetwork   `json:"networks,omitempty"`
	Enabled       bool            `json:"enabled,omitempty"`
	SnmpV3Enabled bool            `json:"snmpV3Enabled,omitempty"`
	UsmUsers      []SnmpV3UsmUser `json:"usmUsers,omitempty"`
}

type SetSnmpTrapInfoRequest struct {
	TrapRecipients                   []SnmpTrapRecipient `json:"trapRecipients"`
	ClusterFaultTrapsEnabled         bool                `json:"clusterFaultTrapsEnabled"`
	ClusterFaultResolvedTrapsEnabled bool                `json:"clusterFaultResolvedTrapsEnabled"`
	ClusterEventTrapsEnabled         bool                `json:"clusterEventTrapsEnabled"`
}

type ShutdownRequest struct {
	Nodes  []int64 `json:"nodes"`
	Option string  `json:"option,omitempty"`
}

type StartBulkVolumeReadRequest struct {
	VolumeID         int64       `json:"volumeID"`
	Format           string      `json:"format"`
	SnapshotID       int64       `json:"snapshotID,omitempty"`
	Script           string      `json:"script,omitempty"`
	ScriptParameters interface{} `json:"scriptParameters,omitempty"`
	Attributes       interface{} `json:"attributes,omitempty"`
}

type StartBulkVolumeWriteRequest struct {
	VolumeID         int64       `json:"volumeID"`
	Format           string      `json:"format"`
	Script           string      `json:"script,omitempty"`
	ScriptParameters interface{} `json:"scriptParameters,omitempty"`
	Attributes       interface{} `json:"attributes,omitempty"`
}

type StartVolumePairingRequest struct {
	VolumeID int64  `json:"volumeID"`
	Mode     string `json:"mode,omitempty"`
}

type TestConnectEnsembleRequest struct {
	Ensemble string `json:"ensemble,omitempty"`
}

type TestConnectMvipRequest struct {
	Mvip string `json:"mvip,omitempty"`
}

type TestConnectSvipRequest struct {
	Svip string `json:"svip,omitempty"`
}

type TestDrivesRequest struct {
	Minutes int64 `json:"minutes,omitempty"`
}

type TestLdapAuthenticationRequest struct {
	Username          string            `json:"username"`
	Password          string            `json:"password"`
	LdapConfiguration LdapConfiguration `json:"ldapConfiguration,omitempty"`
}

type TestPingRequest struct {
	Attempts              int64  `json:"attempts,omitempty"`
	Hosts                 string `json:"hosts,omitempty"`
	TotalTimeoutSec       int64  `json:"totalTimeoutSec,omitempty"`
	PacketSize            int64  `json:"packetSize,omitempty"`
	PingTimeoutMsec       int64  `json:"pingTimeoutMsec,omitempty"`
	ProhibitFragmentation bool   `json:"prohibitFragmentation,omitempty"`
}

type UpdateBulkVolumeStatusRequest struct {
	Key             string      `json:"key"`
	Status          string      `json:"status"`
	PercentComplete string      `json:"percentComplete,omitempty"`
	Message         string      `json:"message,omitempty"`
	Attributes      interface{} `json:"attributes,omitempty"`
}
