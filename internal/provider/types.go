package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	backupdr "github.com/umeshkumhar/backupdr-client"
)

type ApplicationResourceModel struct {
	// Immutable        types.Bool               `tfsdk:"immutable"`
	Description types.String `tfsdk:"description"`
	// Sources          []ApplicationRest  `tfsdk:"sources"`
	Name types.String `tfsdk:"name"`
	// Host             *HostRest          `tfsdk:"host"`
	// Srcid            types.String             `tfsdk:"srcid"`
	// Uniquename       types.String             `tfsdk:"uniquename"`
	Appname types.String `tfsdk:"appname"`
	// Isvm             types.Bool               `tfsdk:"isvm"`
	// Managed          types.Bool               `tfsdk:"managed"`
	// Scheduleoff      types.Bool               `tfsdk:"scheduleoff"`
	Apptype types.String `tfsdk:"apptype"`
	// Originalappid    types.String             `tfsdk:"originalappid"`
	// Pathname         types.String             `tfsdk:"pathname"`
	// Username         types.String             `tfsdk:"username"`
	// Backup           []BackupRest       `tfsdk:"backup"`
	// Isorphan         types.Bool               `tfsdk:"isorphan"`
	// Appclass         types.String             `tfsdk:"appclass"`
	// Sla              *SlaRest           `tfsdk:"sla"`
	// Cluster          *ClusterRest       `tfsdk:"cluster"`
	// Friendlypath     types.String             `tfsdk:"friendlypath"`
	// Sourcecluster    types.String             `tfsdk:"sourcecluster"`
	// Friendlytype     types.String             `tfsdk:"friendlytype"`
	// Volumes          []types.String           `tfsdk:"volumes"`
	// Protectable      types.String             `tfsdk:"protectable"`
	// Failoverstate    types.String             `tfsdk:"failoverstate"`
	// Auxinfo          types.String             `tfsdk:"auxinfo"`
	// Appversion       types.String             `tfsdk:"appversion"`
	// Networkname      types.String             `tfsdk:"networkname"`
	// Networkip        types.String             `tfsdk:"networkip"`
	// Ignore           types.Bool               `tfsdk:"ignore"`
	// Isclustered      types.Bool               `tfsdk:"isclustered"`
	// Frommount        types.Bool               `tfsdk:"frommount"`
	// Sensitivity      types.Int64              `tfsdk:"sensitivity"`
	// Mountedhosts     []HostRest         `tfsdk:"mountedhosts"`
	// AvailableSlp     []SlpRest          `tfsdk:"available_slp"`
	// Orglist          []OrganizationRest `tfsdk:"orglist"`
	// Isrestoring      types.Bool               `tfsdk:"isrestoring"`
	// Consistencygroup *ApplicationRest   `tfsdk:"consistencygroup"`
	// Logicalgroup     *LogicalGroupRest  `tfsdk:"logicalgroup"`
	// AppstateText     []types.String           `tfsdk:"appstate_text"`
	// Diskpools        []types.String           `tfsdk:"diskpools"`
	ID       types.String `tfsdk:"id"`
	Href     types.String `tfsdk:"href"`
	Syncdate types.Int64  `tfsdk:"syncdate"`
	Stale    types.Bool   `tfsdk:"stale"`
}

type LogicalGroupResourceModel struct {
	Description types.String      `tfsdk:"description"`
	Name        types.String      `tfsdk:"name"`
	Srcid       types.String      `tfsdk:"srcid"`
	Modifydate  types.Int64       `tfsdk:"modifydate"`
	Managed     types.Bool        `tfsdk:"managed"`
	Scheduleoff types.Bool        `tfsdk:"scheduleoff"`
	Sla         *SlaResourceModel `tfsdk:"sla"`
	Cluster     *ClusterRest      `tfsdk:"cluster"`
	Membercount types.Int64       `tfsdk:"membercount"`
	// Orglist     []OrganizationRest `tfsdk:"orglist"`
	ID       types.String `tfsdk:"id"`
	Href     types.String `tfsdk:"href"`
	Syncdate types.Int64  `tfsdk:"syncdate"`
	Stale    types.Bool   `tfsdk:"stale"`
}

// vCenter Host
type vcenterHostRest struct {
	// Applicationagent *AgentRest   `tfsdk:"applicationagent"`
	// Arrays               []ArrayRest          `tfsdk:"arrays"`
	// Chappassword     types.String `tfsdk:"chappassword"`
	// Chapusername     types.String `tfsdk:"chapusername"`
	// Clearchap        types.Bool   `tfsdk:"clearchap"`
	// Cloudcredential      *CloudCredentialRest `tfsdk:"cloudcredential"`
	// Connect2actip          types.String `tfsdk:"connect2actip"`
	// Connectorversion     types.String   `tfsdk:"connectorversion"`
	// Dataip           types.String   `tfsdk:"dataip"`
	// Datastorelist          []DataStoreRest      `tfsdk:"datastorelist"`
	// Description      types.String   `tfsdk:"description"`
	// Hmcconnector     *AgentRest   `tfsdk:"hmcconnector"`
	// Immutable   types.Bool     `tfsdk:"immutable"`
	// Iogrp                types.String   `tfsdk:"iogrp"`
	// IogrpCount           types.String   `tfsdk:"iogrp_count"`
	// IsAutoDiscoveryEnabled types.Bool   `tfsdk:"is_auto_discovery_enabled"`
	// IscsiName            []types.String `tfsdk:"iscsi_name"`
	// Machinetype            types.String `tfsdk:"machinetype"`
	// Mask        types.String   `tfsdk:"mask"`
	// Nasdconfig           *NasdConfigRest      `tfsdk:"nasdconfig"`
	// Nfsoption              *NfsOptionsRest      `tfsdk:"nfsoption"`
	// NodeList []HostRest `tfsdk:"nodeList"`
	// NodeLoggedInCount    types.String   `tfsdk:"node_logged_in_count"`
	// Orglist                []OrganizationRest   `tfsdk:"orglist"`
	// Osrelease        types.String   `tfsdk:"osrelease"`
	// Ostype           types.String   `tfsdk:"ostype"`
	// Osversion        types.String   `tfsdk:"osversion"`
	// PkiErrors   []types.String `tfsdk:"pki_errors"`
	// PortCount            types.String   `tfsdk:"port_count"`
	// ReconciliationFields []types.String `tfsdk:"reconciliation_fields"`
	// ReconciliationNeeded types.Bool     `tfsdk:"reconciliation_needed"`
	// Source                 []SourceRest         `tfsdk:"source"`
	// Stale            types.Bool   `tfsdk:"stale"`
	// State            types.String   `tfsdk:"state"`
	// Status           types.String   `tfsdk:"status"`
	// Syncdate         types.Int64  `tfsdk:"syncdate"`
	// Systemdetail           types.String `tfsdk:"systemdetail"`
	// Udsagent         *AgentRest   `tfsdk:"udsagent"`
	// VcenterHost *HostRest `tfsdk:"vcenterHost"`
	// Vcenterhostid    types.String   `tfsdk:"vcenterhostid"`
	// Vmtype           types.String   `tfsdk:"vmtype"`
	// WWPN                 []types.String `tfsdk:"WWPN"`
	// Agents      []types.Object `tfsdk:"agents"`
	Alternateip []types.String `tfsdk:"alternateip"`
	// Appliance        *ClusterRestRef `tfsdk:"appliance"`
	Autoupgrade      types.String `tfsdk:"autoupgrade"`
	CertRevoked      types.Bool   `tfsdk:"cert_revoked"`
	Clusterid        types.String `tfsdk:"clusterid"`
	Dbauthentication types.Bool   `tfsdk:"dbauthentication"`
	Diskpref         types.String `tfsdk:"diskpref"`
	// Esxlist          []vcenterHostRestRef `tfsdk:"esxlist"`
	Friendlypath    types.String         `tfsdk:"friendlypath"`
	Hasagent        types.Bool           `tfsdk:"hasagent"`
	Hostname        types.String         `tfsdk:"hostname"`
	Hosttype        types.String         `tfsdk:"hosttype"`
	Href            types.String         `tfsdk:"href"`
	Hypervisoragent *AgentRest           `tfsdk:"hypervisoragent"`
	ID              types.String         `tfsdk:"id"`
	Ipaddress       types.String         `tfsdk:"ipaddress"`
	IsClusterNode   types.Bool           `tfsdk:"isclusternode"`
	IsShadowHost    types.Bool           `tfsdk:"isshadowhost"`
	Isclusterhost   types.Bool           `tfsdk:"isclusterhost"`
	Isesxhost       types.Bool           `tfsdk:"isesxhost"`
	Isproxyhost     types.Bool           `tfsdk:"isproxyhost"`
	Isvcenterhost   types.Bool           `tfsdk:"isvcenterhost"`
	Isvm            types.Bool           `tfsdk:"isvm"`
	Maxjobs         types.Int64          `tfsdk:"maxjobs"`
	Modifydate      types.Int64          `tfsdk:"modifydate"`
	Multiregion     types.String         `tfsdk:"multiregion"`
	Name            types.String         `tfsdk:"name"`
	Originalhostid  types.String         `tfsdk:"originalhostid"`
	OstypeSpecial   types.String         `tfsdk:"ostype_special"`
	PkiState        types.String         `tfsdk:"pki_state"`
	Sourcecluster   types.String         `tfsdk:"sourcecluster"`
	Sources         []vcenterHostRestRef `tfsdk:"sources"`
	Srcid           types.String         `tfsdk:"srcid"`
	Svcname         types.String         `tfsdk:"svcname"`
	Transport       types.String         `tfsdk:"transport"`
	Uniquename      types.String         `tfsdk:"uniquename"`
	Zone            types.String         `tfsdk:"zone"`
}

type vcenterHostRestRef struct {
	ID        types.String `tfsdk:"id"`
	Clusterid types.String `tfsdk:"clusterid"`
	Href      types.String `tfsdk:"href"`
}

type VmDiscoveryRest struct {
	Cluster          string   `json:"cluster,omitempty"`
	Addvms           bool     `json:"addvms,omitempty"`
	VcenterHostId    string   `json:"vcenterHostId,omitempty"`
	Esxcluster       string   `json:"esxcluster,omitempty"`
	Addall           bool     `json:"addall,omitempty"`
	Vms              []string `json:"vms,omitempty"`
	Dc               string   `json:"dc,omitempty"`
	Discoverclusters bool     `json:"discoverclusters,omitempty"`
	Discovervms      bool     `json:"discovervms,omitempty"`
	Org              string   `json:"org,omitempty"`
	// Unique ID for this object
	Id string `json:"id,omitempty"`
	// URL to access this object
	Href string `json:"href,omitempty"`
	// When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources.
	Syncdate int64 `json:"syncdate,omitempty"`
	// Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources.
	Stale bool `json:"stale,omitempty"`
}

type AgentRest struct {
	// Alternatekey    types.String `tfsdk:"alternatekey"`
	// Port            types.Int64  `tfsdk:"port"`
	// SharedSecret    types.String `tfsdk:"shared_secret"`
	// AgentVersion    types.String `tfsdk:"agent_version"`
	Agenttype       types.String `tfsdk:"agenttype"`
	Hasalternatekey types.Bool   `tfsdk:"hasalternatekey"`
	Haspassword     types.Bool   `tfsdk:"haspassword"`
	// Href            types.String `tfsdk:"href"`
	// ID              types.String `tfsdk:"id"`
	Password types.String `tfsdk:"password"`
	// Stale           types.Bool   `tfsdk:"stale"`
	// Syncdate        types.Int64  `tfsdk:"syncdate"`
	Username types.String `tfsdk:"username"`
}

// SLA
type SlaResourceModel struct {
	// Immutable   types.Bool                `tfsdk:"immutable"`
	Description types.String              `tfsdk:"description"`
	Application *ApplicationResourceModel `tfsdk:"application"`
	Slt         *SltResourceRefModel      `tfsdk:"slt"`
	// Options          []AdvancedOptionRest      `tfsdk:"options"`
	Modifydate       types.Int64          `tfsdk:"modifydate"`
	Scheduleoff      types.String         `tfsdk:"scheduleoff"`
	Slp              *SlpResourceRefModel `tfsdk:"slp"`
	Logexpirationoff types.Bool           `tfsdk:"logexpirationoff"`
	Dedupasyncoff    types.String         `tfsdk:"dedupasyncoff"`
	Expirationoff    types.String         `tfsdk:"expirationoff"`
	// Group            *backupdr.LogicalGroupRest `tfsdk:"group"`
	ID       types.String `tfsdk:"id"`
	Href     types.String `tfsdk:"href"`
	Syncdate types.Int64  `tfsdk:"syncdate"`
	Stale    types.Bool   `tfsdk:"stale"`
}

type SlaResourceRefModel struct {
	ID       types.String `tfsdk:"id"`
	Href     types.String `tfsdk:"href"`
	Syncdate types.Int64  `tfsdk:"syncdate"`
	Stale    types.Bool   `tfsdk:"stale"`
}

// SLT
type SltResourceModel struct {
	ID          types.String `tfsdk:"id"`
	Href        types.String `tfsdk:"href"`
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
	// Immutable   types.Bool   `tfsdk:"immutable"`
	OptionHref types.String `tfsdk:"option_href"`
	PolicyHref types.String `tfsdk:"policy_href"`
	// Source         []backupdr.SourceRest         `tfsdk:"source"`
	Sourcename types.String      `tfsdk:"sourcename"`
	Override   types.String      `tfsdk:"override"`
	Policies   []policyRestModel `tfsdk:"policies"`
	// Options        []backupdr.AdvancedOptionRest `tfsdk:"options"`
	// Orglist        []backupdr.OrganizationRest   `tfsdk:"orglist"`
	Managedbyagm   types.Bool  `tfsdk:"managedbyagm"`
	Usedbycloudapp types.Bool  `tfsdk:"usedbycloudapp"`
	Syncdate       types.Int64 `tfsdk:"syncdate"`
	Stale          types.Bool  `tfsdk:"stale"`
}

type SltResourceRefModel struct {
	ID         types.String `tfsdk:"id"`
	Href       types.String `tfsdk:"href"`
	Name       types.String `tfsdk:"name"`
	Sourcename types.String `tfsdk:"sourcename"`
	Override   types.String `tfsdk:"override"`
	Stale      types.Bool   `tfsdk:"stale"`
}

type policyRestModel struct {
	// Source             []backupdr.SourceRest            `tfsdk:"source"`
	Description types.String `tfsdk:"description"`
	Name        types.String `tfsdk:"name"`
	Priority    types.String `tfsdk:"priority"`
	// Slt         *sltRestModel `tfsdk:"slt"`
	// Endtime types.String `tfsdk:"endtime"`
	Rpom types.String `tfsdk:"rpom"`
	Rpo  types.String `tfsdk:"rpo"`
	// Predecessor        *backupdr.PolicyRest             `tfsdk:"predecessor"`
	Exclusiontype types.String `tfsdk:"exclusiontype"`
	Iscontinuous  types.Bool   `tfsdk:"iscontinuous"`
	Starttime     types.String `tfsdk:"starttime"`
	Endtime       types.String `tfsdk:"endtime"`
	Targetvault   types.Int64  `tfsdk:"targetvault"`
	Sourcevault   types.Int64  `tfsdk:"sourcevault"`
	Selection     types.String `tfsdk:"selection"`
	Scheduletype  types.String `tfsdk:"scheduletype"`
	// Scheduling        types.String `tfsdk:"scheduling"`
	Exclusion         types.String `tfsdk:"exclusion"`
	Reptype           types.String `tfsdk:"reptype"`
	Retention         types.String `tfsdk:"retention"`
	Retentionm        types.String `tfsdk:"retentionm"`
	Encrypt           types.String `tfsdk:"encrypt"`
	Repeatinterval    types.String `tfsdk:"repeatinterval"`
	Exclusioninterval types.String `tfsdk:"exclusioninterval"`
	Remoteretention   types.Int64  `tfsdk:"remoteretention"`
	// Compliancesettings complianceSettingsRestModel `tfsdk:"compliancesettings"`
	// Options            []backupdr.AdvancedOptionRest `tfsdk:"options"`
	PolicyType   types.String `tfsdk:"policytype"`
	Truncatelog  types.String `tfsdk:"truncatelog"`
	Verifychoice types.String `tfsdk:"verifychoice"`
	Op           types.String `tfsdk:"op"`
	Verification types.Bool   `tfsdk:"verification"`
	ID           types.String `tfsdk:"id"`
	Href         types.String `tfsdk:"href"`
	Syncdate     types.Int64  `tfsdk:"syncdate"`
	Stale        types.Bool   `tfsdk:"stale"`
}

type complianceSettingsRestModel struct {
	Policy               *backupdr.PolicyRest `tfsdk:"policy"`
	WarnThresholdType    types.String         `tfsdk:"warn_threshold_type"`
	WarnThresholdCustom  types.Int64          `tfsdk:"warn_threshold_custom"`
	ErrorThresholdType   types.String         `tfsdk:"error_threshold_type"`
	ErrorThresholdCustom types.Int64          `tfsdk:"error_threshold_custom"`
	ID                   types.String         `tfsdk:"id"`
	Href                 types.String         `tfsdk:"href"`
	Syncdate             types.Int64          `tfsdk:"syncdate"`
	Stale                types.Bool           `tfsdk:"stale"`
}

// SLP
// SlpResourceModel represent TF SLA profile object
type SlpResourceModel struct {
	Description     types.String `tfsdk:"description"`
	Name            types.String `tfsdk:"name"`
	Srcid           types.String `tfsdk:"srcid"`
	Clusterid       types.String `tfsdk:"clusterid"`
	Modifydate      types.Int64  `tfsdk:"modifydate"`
	Cid             types.String `tfsdk:"cid"`
	Performancepool types.String `tfsdk:"performancepool"`
	//** Primarystorage  types.String           `tfsdk:"primarystorage"`
	Remotenode types.String `tfsdk:"remotenode"`
	// **
	Dedupasyncnode types.String              `tfsdk:"dedupasyncnode"`
	Vaultpool      *slpDiskPoolResourceModel `tfsdk:"vaultpool"`
	Vaultpool2     *slpDiskPoolResourceModel `tfsdk:"vaultpool2"`
	Vaultpool3     *slpDiskPoolResourceModel `tfsdk:"vaultpool3"`
	Vaultpool4     *slpDiskPoolResourceModel `tfsdk:"vaultpool4"`
	Createdate     types.Int64               `tfsdk:"createdate"`
	Localnode      types.String              `tfsdk:"localnode"`
	// Orglist         []OrganizationRest   `tfsdk:"orglist"`
	// CloudCredential *CloudCredentialRest `tfsdk:"cloudCredential"`
	ID       types.String `tfsdk:"id"`
	Href     types.String `tfsdk:"href"`
	Syncdate types.Int64  `tfsdk:"syncdate"`
	Stale    types.Bool   `tfsdk:"stale"`
}

type SlpResourceRefModel struct {
	Name     types.String `tfsdk:"name"`
	Cid      types.String `tfsdk:"cid"`
	ID       types.String `tfsdk:"id"`
	Href     types.String `tfsdk:"href"`
	Syncdate types.Int64  `tfsdk:"syncdate"`
	Stale    types.Bool   `tfsdk:"stale"`
}

// DiskPoolResourceModel represent diskpool object
type slpDiskPoolResourceModel struct {
	Name types.String `tfsdk:"name"`
	ID   types.String `tfsdk:"id"`
	Href types.String `tfsdk:"href"`
}

// DiskPoolResourceModel represent diskpool object
type DiskPoolResourceModel struct {
	Name       types.String        `tfsdk:"name"`
	Pooltype   types.String        `tfsdk:"pooltype"`
	Cluster    *ClusterRest        `tfsdk:"cluster"`
	Properties []KeyValueRestModel `tfsdk:"properties"`
	// Computed Attributes
	Vaultprops          *VaultPropsRest `tfsdk:"vaultprops"`
	Usedefaultsa        types.Bool      `tfsdk:"usedefaultsa"`
	Immutable           types.Bool      `tfsdk:"immutable"`
	Metadataonly        types.Bool      `tfsdk:"metadataonly"`
	State               types.String    `tfsdk:"state"`
	Srcid               types.String    `tfsdk:"srcid"`
	Status              types.String    `tfsdk:"status"`
	Mdiskgrp            types.String    `tfsdk:"mdiskgrp"`
	Modifydate          types.Int64     `tfsdk:"modifydate"`
	Warnpct             types.Int64     `tfsdk:"warnpct"`
	Safepct             types.Int64     `tfsdk:"safepct"`
	Udsuid              types.Int64     `tfsdk:"udsuid"`
	FreeMb              types.Int64     `tfsdk:"free_mb"`
	UsageMb             types.Int64     `tfsdk:"usage_mb"`
	CapacityMb          types.Int64     `tfsdk:"capacity_mb"`
	Pct                 types.Float64   `tfsdk:"pct"`
	Pooltypedisplayname types.String    `tfsdk:"pooltypedisplayname"`
	ID                  types.String    `tfsdk:"id"`
	Href                types.String    `tfsdk:"href"`
	Syncdate            types.Int64     `tfsdk:"syncdate"`
	Stale               types.Bool      `tfsdk:"stale"`
}

type KeyValueRestModel struct {
	Value types.String `tfsdk:"value"`
	Key   types.String `tfsdk:"key"`
}

type ClusterRest struct {
	Clusterid types.String `tfsdk:"clusterid"`
	// Computed Attributes
	Serviceaccount  types.String `tfsdk:"serviceaccount"`
	Zone            types.String `tfsdk:"zone"`
	Region          types.String `tfsdk:"region"`
	Projectid       types.String `tfsdk:"projectid"`
	Version         types.String `tfsdk:"version"`
	Name            types.String `tfsdk:"name"`
	Type            types.String `tfsdk:"type"`
	Ipaddress       types.String `tfsdk:"ipaddress"`
	Publicip        types.String `tfsdk:"publicip"`
	Secureconnect   types.Bool   `tfsdk:"secureconnect"`
	PkiBootstrapped types.Bool   `tfsdk:"pkibootstrapped"`
	Supportstatus   types.String `tfsdk:"supportstatus"`
	ID              types.String `tfsdk:"id"`
	Href            types.String `tfsdk:"href"`
	Syncdate        types.Int64  `tfsdk:"syncdate"`
	Stale           types.Bool   `tfsdk:"stale"`
}

type ClusterRestRef struct {
	Clusterid types.String `tfsdk:"clusterid"`
	// Computed Attributes
	ID   types.String `tfsdk:"id"`
	Href types.String `tfsdk:"href"`
}

type VaultPropsRest struct {
	Bucket      types.String `tfsdk:"bucket"`
	Compression types.Bool   `tfsdk:"compression"`
	Region      types.String `tfsdk:"region"`
	ID          types.String `tfsdk:"id"`
	Href        types.String `tfsdk:"href"`
	Syncdate    types.Int64  `tfsdk:"syncdate"`
	Stale       types.Bool   `tfsdk:"stale"`
}
