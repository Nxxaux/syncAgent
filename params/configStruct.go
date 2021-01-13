package params

import "time"

//DeviceIDLength for DeviceID
//const DeviceIDLength = 32

//DeviceID in DeviceConfiguration
//type DeviceID [DeviceIDLength]byte

//GUIConfiguration for GUI
type GUIConfiguration struct {
	Enabled                   bool   `json:"enabled" xml:"enabled,attr" default:"true"`
	RawAddress                string `json:"address" xml:"address" default:"127.0.0.1:8384"`
	RawUnixSocketPermissions  string `json:"unixSocketPermissions" xml:"unixSocketPermissions,omitempty"`
	User                      string `json:"user" xml:"user,omitempty"`
	Password                  string `json:"password" xml:"password,omitempty"`
	AuthMode                  string `json:"authMode" xml:"authMode,omitempty"`
	RawUseTLS                 bool   `json:"useTLS" xml:"tls,attr"`
	APIKey                    string `json:"apiKey" xml:"apikey,omitempty"`
	InsecureAdminAccess       bool   `json:"insecureAdminAccess" xml:"insecureAdminAccess,omitempty"`
	Theme                     string `json:"theme" xml:"theme" default:"default"`
	Debugging                 bool   `json:"debugging" xml:"debugging,attr"`
	InsecureSkipHostCheck     bool   `json:"insecureSkipHostcheck" xml:"insecureSkipHostcheck,omitempty"`
	InsecureAllowFrameLoading bool   `json:"insecureAllowFrameLoading" xml:"insecureAllowFrameLoading,omitempty"`
}

//ObservedFolder in DeviceConfiguration
type ObservedFolder struct {
	Time  time.Time `json:"time" xml:"time,attr"`
	ID    string    `json:"id" xml:"id,attr"`
	Label string    `json:"label" xml:"label,attr"`
}

//DeviceConfiguration here
type DeviceConfiguration struct {
	DeviceID                 string           `json:"deviceID" xml:"id,attr"`
	Name                     string           `json:"name" xml:"name,attr,omitempty"`
	Addresses                []string         `json:"addresses" xml:"address,omitempty" default:"dynamic"`
	Compression              string           `json:"compression" xml:"compression,attr"`
	CertName                 string           `json:"certName" xml:"certName,attr,omitempty"`
	Introducer               bool             `json:"introducer" xml:"introducer,attr"`
	SkipIntroductionRemovals bool             `json:"skipIntroductionRemovals" xml:"skipIntroductionRemovals,attr"`
	IntroducedBy             string           `json:"introducedBy" xml:"introducedBy,attr"`
	Paused                   bool             `json:"paused" xml:"paused"`
	AllowedNetworks          []string         `json:"allowedNetworks" xml:"allowedNetwork,omitempty"`
	AutoAcceptFolders        bool             `json:"autoAcceptFolders" xml:"autoAcceptFolders"`
	MaxSendKbps              int              `json:"maxSendKbps" xml:"maxSendKbps"`
	MaxRecvKbps              int              `json:"maxRecvKbps" xml:"maxRecvKbps"`
	IgnoredFolders           []ObservedFolder `json:"ignoredFolders" xml:"ignoredFolder"`
	DeprecatedPendingFolders []ObservedFolder `json:"-" xml:"pendingFolder,omitempty"` // Deprecated: Do not use.
	MaxRequestKiB            int              `json:"maxRequestKiB" xml:"maxRequestKiB"`
	Untrusted                bool             `json:"untrusted" xml:"untrusted"`
	RemoteGUIPort            int              `json:"remoteGUIPort" xml:"remoteGUIPort"`
}

//Size for folder
type Size struct {
	Value float64 `json:"value" xml:",chardata"`
	Unit  string  `json:"unit" xml:"unit,attr"`
}

//FolderDeviceConfiguration in FolderConfiguration
type FolderDeviceConfiguration struct {
	DeviceID           string `json:"deviceID"`     //web controller given
	IntroducedBy       string `json:"introducedBy"` //web controller given
	EncryptionPassword string `json:"encryptionPassword"`
}

// VersioningConfiguration is used in the code and for JSON serialization
type VersioningConfiguration struct {
	Type             string            `json:"type" `
	Params           map[string]string `json:"params"  protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	CleanupIntervalS int               `json:"cleanupIntervalS" default:"3600"`
}

//FolderConfiguration for forlder
type FolderConfiguration struct {
	ID                      string                      `json:"id" xml:"id,attr"`
	Label                   string                      `json:"label" xml:"label,attr" restart:"false"`
	FilesystemType          string                      `json:"filesystemType" xml:"filesystemType"`
	Path                    string                      `json:"path" xml:"path,attr"`
	Type                    string                      `json:"type" xml:"type,attr"`
	Devices                 []FolderDeviceConfiguration `json:"devices" xml:"device"`
	RescanIntervalS         int                         `json:"rescanIntervalS" xml:"rescanIntervalS,attr" default:"3600"`
	FSWatcherEnabled        bool                        `json:"fsWatcherEnabled" xml:"fsWatcherEnabled,attr" default:"true"`
	FSWatcherDelayS         int                         `json:"fsWatcherDelayS" xml:"fsWatcherDelayS,attr" default:"10"`
	IgnorePerms             bool                        `json:"ignorePerms" xml:"ignorePerms,attr"`
	AutoNormalize           bool                        `json:"autoNormalize" xml:"autoNormalize,attr" default:"true"`
	MinDiskFree             Size                        `json:"minDiskFree" xml:"minDiskFree"`
	Versioning              VersioningConfiguration     `json:"versioning" xml:"versioning"`
	Copiers                 int                         `json:"copiers" xml:"copiers"`
	PullerMaxPendingKiB     int                         `json:"pullerMaxPendingKiB" xml:"pullerMaxPendingKiB"`
	Hashers                 int                         `json:"hashers" xml:"hashers"`
	Order                   string                      `json:"order" xml:"order"`
	IgnoreDelete            bool                        `json:"ignoreDelete" xml:"ignoreDelete"`
	ScanProgressIntervalS   int                         `json:"scanProgressIntervalS" xml:"scanProgressIntervalS"`
	PullerPauseS            int                         `json:"pullerPauseS" xml:"pullerPauseS"`
	MaxConflicts            int                         `json:"maxConflicts" xml:"maxConflicts" default:"-1"`
	DisableSparseFiles      bool                        `json:"disableSparseFiles" xml:"disableSparseFiles"`
	DisableTempIndexes      bool                        `json:"disableTempIndexes" xml:"disableTempIndexes"`
	Paused                  bool                        `json:"paused" xml:"paused"`
	WeakHashThresholdPct    int                         `json:"weakHashThresholdPct" xml:"weakHashThresholdPct"`
	MarkerName              string                      `json:"markerName" xml:"markerName"`
	CopyOwnershipFromParent bool                        `json:"copyOwnershipFromParent" xml:"copyOwnershipFromParent"`
	RawModTimeWindowS       int                         `json:"modTimeWindowS" xml:"modTimeWindowS"`
	MaxConcurrentWrites     int                         `json:"maxConcurrentWrites" xml:"maxConcurrentWrites" default:"2"`
	DisableFsync            bool                        `json:"disableFsync" xml:"disableFsync"`
	BlockPullOrder          string                      `json:"blockPullOrder" xml:"blockPullOrder"`
	CopyRangeMethod         string                      `json:"copyRangeMethod" xml:"copyRangeMethod" default:"standard"`
	CaseSensitiveFS         bool                        `json:"caseSensitiveFS" xml:"caseSensitiveFS"`
	JunctionsAsDirs         bool                        `json:"junctionsAsDirs" xml:"junctionsAsDirs"`
	// Legacy deprecated
	DeprecatedReadOnly       bool    `json:"-" xml:"ro,attr,omitempty"`        // Deprecated: Do not use.
	DeprecatedMinDiskFreePct float64 `json:"-" xml:"minDiskFreePct,omitempty"` // Deprecated: Do not use.
	DeprecatedPullers        int     `json:"-" xml:"pullers,omitempty"`        // Deprecated: Do not use.
}

//LDAPConfiguration here
type LDAPConfiguration struct {
	Address            string `json:"address" xml:"address,omitempty"`
	BindDN             string `json:"bindDN" xml:"bindDN,omitempty"`
	Transport          string `json:"transport" xml:"transport,omitempty"`
	InsecureSkipVerify bool   `json:"insecureSkipVerify" xml:"insecureSkipVerify,omitempty" default:"false"`
	SearchBaseDN       string `json:"searchBaseDN" xml:"searchBaseDN,omitempty"`
	SearchFilter       string `json:"searchFilter" xml:"searchFilter,omitempty"`
}

//OptionsConfiguration for Configuration
type OptionsConfiguration struct {
	RawListenAddresses      []string `json:"listenAddresses" xml:"listenAddress" default:"default"`
	RawGlobalAnnServers     []string `json:"globalAnnounceServers" xml:"globalAnnounceServer" default:"default"`
	GlobalAnnEnabled        bool     `json:"globalAnnounceEnabled" xml:"globalAnnounceEnabled" default:"true"`
	LocalAnnEnabled         bool     `json:"localAnnounceEnabled" xml:"localAnnounceEnabled" default:"true"`
	LocalAnnPort            int      `json:"localAnnouncePort" xml:"localAnnouncePort" default:"21027"`
	LocalAnnMCAddr          string   `json:"localAnnounceMCAddr" xml:"localAnnounceMCAddr" default:"[ff12::8384]:21027"`
	MaxSendKbps             int      `json:"maxSendKbps" xml:"maxSendKbps"`
	MaxRecvKbps             int      `json:"maxRecvKbps" xml:"maxRecvKbps"`
	ReconnectIntervalS      int      `json:"reconnectionIntervalS" xml:"reconnectionIntervalS" default:"60"`
	RelaysEnabled           bool     `json:"relaysEnabled" xml:"relaysEnabled" default:"true"`
	RelayReconnectIntervalM int      `json:"relayReconnectIntervalM" xml:"relayReconnectIntervalM" default:"10"`
	StartBrowser            bool     `json:"startBrowser" xml:"startBrowser" default:"true"`
	NATEnabled              bool     `json:"natEnabled" xml:"natEnabled" default:"true"`
	NATLeaseM               int      `json:"natLeaseMinutes" xml:"natLeaseMinutes" default:"60"`
	NATRenewalM             int      `json:"natRenewalMinutes" xml:"natRenewalMinutes" default:"30"`
	NATTimeoutS             int      `json:"natTimeoutSeconds" xml:"natTimeoutSeconds" default:"10"`
	URAccepted              int      `json:"urAccepted" xml:"urAccepted"`
	URSeen                  int      `json:"urSeen" xml:"urSeen"`
	URUniqueID              string   `json:"urUniqueId" xml:"urUniqueID"`
	URURL                   string   `json:"urURL" xml:"urURL" default:"https://data.syncthing.net/newdata"`
	URPostInsecurely        bool     `json:"urPostInsecurely" xml:"urPostInsecurely" default:"false"`
	URInitialDelayS         int      `json:"urInitialDelayS" xml:"urInitialDelayS" default:"1800"`
	RestartOnWakeup         bool     `json:"restartOnWakeup" xml:"restartOnWakeup" default:"true"`
	AutoUpgradeIntervalH    int      `json:"autoUpgradeIntervalH" xml:"autoUpgradeIntervalH" default:"12"`
	UpgradeToPreReleases    bool     `json:"upgradeToPreReleases" xml:"upgradeToPreReleases"`
	KeepTemporariesH        int      `json:"keepTemporariesH" xml:"keepTemporariesH" default:"24"`
	CacheIgnoredFiles       bool     `json:"cacheIgnoredFiles" xml:"cacheIgnoredFiles" default:"false"`
	ProgressUpdateIntervalS int      `json:"progressUpdateIntervalS" xml:"progressUpdateIntervalS" default:"5"`
	LimitBandwidthInLan     bool     `json:"limitBandwidthInLan" xml:"limitBandwidthInLan" default:"false"`
	MinHomeDiskFree         Size     `json:"minHomeDiskFree" xml:"minHomeDiskFree" default:"1 %"`
	ReleasesURL             string   `json:"releasesURL" xml:"releasesURL" default:"https://upgrades.syncthing.net/meta.json"`
	AlwaysLocalNets         []string `json:"alwaysLocalNets" xml:"alwaysLocalNet"`
	OverwriteRemoteDevNames bool     `json:"overwriteRemoteDeviceNamesOnConnect" xml:"overwriteRemoteDeviceNamesOnConnect" default:"false"`
	TempIndexMinBlocks      int      `json:"tempIndexMinBlocks" xml:"tempIndexMinBlocks" default:"10"`
	UnackedNotificationIDs  []string `json:"unackedNotificationIDs" xml:"unackedNotificationID"`
	TrafficClass            int      `json:"trafficClass" xml:"trafficClass"`
	DefaultFolderPath       string   `json:"defaultFolderPath" xml:"defaultFolderPath" default:"~"`
	SetLowPriority          bool     `json:"setLowPriority" xml:"setLowPriority" default:"true"`
	RawMaxFolderConcurrency int      `json:"maxFolderConcurrency" xml:"maxFolderConcurrency"`
	CRURL                   string   `json:"crURL" xml:"crashReportingURL" default:"https://crash.syncthing.net/newcrash"`
	CREnabled               bool     `json:"crashReportingEnabled" xml:"crashReportingEnabled" default:"true"`
	StunKeepaliveStartS     int      `json:"stunKeepaliveStartS" xml:"stunKeepaliveStartS" default:"180"`
	StunKeepaliveMinS       int      `json:"stunKeepaliveMinS" xml:"stunKeepaliveMinS" default:"20"`
	RawStunServers          []string `json:"stunServers" xml:"stunServer" default:"default"`
	DatabaseTuning          string   `json:"databaseTuning" xml:"databaseTuning" restart:"true"`
	RawMaxCIRequestKiB      int      `json:"maxConcurrentIncomingRequestKiB" xml:"maxConcurrentIncomingRequestKiB"`
	AnnounceLANAddresses    bool     `json:"announceLANAddresses" xml:"announceLANAddresses" default:"true"`
	SendFullIndexOnUpgrade  bool     `json:"sendFullIndexOnUpgrade" xml:"sendFullIndexOnUpgrade"`
	FeatureFlags            []string `json:"featureFlags" xml:"featureFlag"`
	// Legacy deprecated
	DeprecatedUPnPEnabled        bool     `json:"-" xml:"upnpEnabled,omitempty"`        // Deprecated: Do not use.
	DeprecatedUPnPLeaseM         int      `json:"-" xml:"upnpLeaseMinutes,omitempty"`   // Deprecated: Do not use.
	DeprecatedUPnPRenewalM       int      `json:"-" xml:"upnpRenewalMinutes,omitempty"` // Deprecated: Do not use.
	DeprecatedUPnPTimeoutS       int      `json:"-" xml:"upnpTimeoutSeconds,omitempty"` // Deprecated: Do not use.
	DeprecatedRelayServers       []string `json:"-" xml:"relayServer,omitempty"`        // Deprecated: Do not use.
	DeprecatedMinHomeDiskFreePct float64  `json:"-" xml:"minHomeDiskFreePct,omitempty"` // Deprecated: Do not use.
	DeprecatedMaxConcurrentScans int      `json:"-" xml:"maxConcurrentScans,omitempty"` // Deprecated: Do not use.
}

//ObservedDevice in Configuration
type ObservedDevice struct {
	Time    time.Time `json:"time" xml:"time,attr"`
	ID      string    `json:"deviceID" xml:"id,attr"`
	Name    string    `json:"name" xml:"name,attr"`
	Address string    `json:"address" xml:"address,attr"`
}

//Configuration of syncthing
type Configuration struct {
	Version                  int                   `json:"version" xml:"version,attr"`
	Folders                  []FolderConfiguration `json:"folders" xml:"folder"`
	Devices                  []DeviceConfiguration `json:"devices" xml:"device"`
	GUI                      GUIConfiguration      `json:"gui" xml:"gui"`
	LDAP                     LDAPConfiguration     `json:"ldap" xml:"ldap"`
	Options                  OptionsConfiguration  `json:"options" xml:"options"`
	IgnoredDevices           []ObservedDevice      `json:"remoteIgnoredDevices" xml:"remoteIgnoredDevice"`
	DeprecatedPendingDevices []ObservedDevice      `json:"-" xml:"pendingDevice,omitempty"` // Deprecated: Do not use.
}
