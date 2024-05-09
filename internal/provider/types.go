package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
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
	Description types.String       `tfsdk:"description"`
	Name        types.String       `tfsdk:"name"`
	Srcid       types.String       `tfsdk:"srcid"`
	Modifydate  types.Int64        `tfsdk:"modifydate"`
	Managed     types.Bool         `tfsdk:"managed"`
	Scheduleoff types.Bool         `tfsdk:"scheduleoff"`
	Sla         *planResourceModel `tfsdk:"sla"`
	Cluster     *ClusterRest       `tfsdk:"cluster"`
	Membercount types.Int64        `tfsdk:"membercount"`
	// Orglist     []OrganizationRest `tfsdk:"orglist"`
	ID       types.String `tfsdk:"id"`
	Href     types.String `tfsdk:"href"`
	Syncdate types.Int64  `tfsdk:"syncdate"`
	Stale    types.Bool   `tfsdk:"stale"`
}

// ###########################################
// #########     backupdr_host    ############
// ###########################################

type vcenterHostRest struct {
	Alternateip []types.String `tfsdk:"alternateip"`
	// Appliance        *ClusterRestRef `tfsdk:"appliance"`
	Autoupgrade      types.String `tfsdk:"autoupgrade"`
	CertRevoked      types.Bool   `tfsdk:"cert_revoked"`
	Clusterid        types.String `tfsdk:"clusterid"`
	Dbauthentication types.Bool   `tfsdk:"dbauthentication"`
	Diskpref         types.String `tfsdk:"diskpref"`
	Friendlypath     types.String `tfsdk:"friendlypath"`
	Hasagent         types.Bool   `tfsdk:"hasagent"`
	Hostname         types.String `tfsdk:"hostname"`
	Hosttype         types.String `tfsdk:"hosttype"`
	Href             types.String `tfsdk:"href"`
	Hypervisoragent  *agentRest   `tfsdk:"hypervisoragent"`
	ID               types.String `tfsdk:"id"`
	Ipaddress        types.String `tfsdk:"ipaddress"`
	IsClusterNode    types.Bool   `tfsdk:"isclusternode"`
	IsShadowHost     types.Bool   `tfsdk:"isshadowhost"`
	Isclusterhost    types.Bool   `tfsdk:"isclusterhost"`
	Isesxhost        types.Bool   `tfsdk:"isesxhost"`
	Isproxyhost      types.Bool   `tfsdk:"isproxyhost"`
	Isvcenterhost    types.Bool   `tfsdk:"isvcenterhost"`
	Isvm             types.Bool   `tfsdk:"isvm"`
	Maxjobs          types.Int64  `tfsdk:"maxjobs"`
	Modifydate       types.Int64  `tfsdk:"modifydate"`
	Multiregion      types.String `tfsdk:"multiregion"`
	Name             types.String `tfsdk:"name"`
	Originalhostid   types.String `tfsdk:"originalhostid"`
	OstypeSpecial    types.String `tfsdk:"ostype_special"`
	PkiState         types.String `tfsdk:"pki_state"`
	Sourcecluster    types.String `tfsdk:"sourcecluster"`
	// Sources          []vcenterHostRestRef `tfsdk:"sources"`
	ApplianceClusterID types.String `tfsdk:"appliance_clusterid"`
	Srcid              types.String `tfsdk:"srcid"`
	Svcname            types.String `tfsdk:"svcname"`
	Transport          types.String `tfsdk:"transport"`
	Uniquename         types.String `tfsdk:"uniquename"`
	Zone               types.String `tfsdk:"zone"`
}

type vcenterHostRestRef struct {
	ID        types.String `tfsdk:"id"`
	Clusterid types.String `tfsdk:"clusterid"`
	Href      types.String `tfsdk:"href"`
}

type agentRest struct {
	Agenttype       types.String `tfsdk:"agenttype"`
	Hasalternatekey types.Bool   `tfsdk:"hasalternatekey"`
	Haspassword     types.Bool   `tfsdk:"haspassword"`
	Password        types.String `tfsdk:"password"`
	Username        types.String `tfsdk:"username"`
}

// ###########################################
// #########     backupdr_plan    ############
// ###########################################

type planResourceModel struct {
	Description types.String              `tfsdk:"description"`
	Application *ApplicationResourceModel `tfsdk:"application"`
	Slt         *templateResourceRefModel `tfsdk:"slt"`
	// Options          []AdvancedOptionRest      `tfsdk:"options"`
	Modifydate       types.Int64              `tfsdk:"modifydate"`
	Scheduleoff      types.String             `tfsdk:"scheduleoff"`
	Slp              *profileResourceRefModel `tfsdk:"slp"`
	Logexpirationoff types.Bool               `tfsdk:"logexpirationoff"`
	Dedupasyncoff    types.String             `tfsdk:"dedupasyncoff"`
	Expirationoff    types.String             `tfsdk:"expirationoff"`
	// Group            *backupdr.LogicalGroupRest `tfsdk:"group"`
	ID       types.String `tfsdk:"id"`
	Href     types.String `tfsdk:"href"`
	Syncdate types.Int64  `tfsdk:"syncdate"`
	Stale    types.Bool   `tfsdk:"stale"`
}

type planResourceRefModel struct {
	ID       types.String `tfsdk:"id"`
	Href     types.String `tfsdk:"href"`
	Syncdate types.Int64  `tfsdk:"syncdate"`
	Stale    types.Bool   `tfsdk:"stale"`
}

// ###########################################
// #########   backupdr_template  ############
// ###########################################

type templateResourceModel struct {
	ID          types.String      `tfsdk:"id"`
	Href        types.String      `tfsdk:"href"`
	Name        types.String      `tfsdk:"name"`
	Description types.String      `tfsdk:"description"`
	OptionHref  types.String      `tfsdk:"option_href"`
	PolicyHref  types.String      `tfsdk:"policy_href"`
	Sourcename  types.String      `tfsdk:"sourcename"`
	Override    types.String      `tfsdk:"override"`
	Policies    []policyRestModel `tfsdk:"policies"`
	// Options        []backupdr.AdvancedOptionRest `tfsdk:"options"`
	Managedbyagm   types.Bool  `tfsdk:"managedbyagm"`
	Usedbycloudapp types.Bool  `tfsdk:"usedbycloudapp"`
	Syncdate       types.Int64 `tfsdk:"syncdate"`
	Stale          types.Bool  `tfsdk:"stale"`
}

type templateResourceRefModel struct {
	ID         types.String `tfsdk:"id"`
	Href       types.String `tfsdk:"href"`
	Name       types.String `tfsdk:"name"`
	Sourcename types.String `tfsdk:"sourcename"`
	Override   types.String `tfsdk:"override"`
	Stale      types.Bool   `tfsdk:"stale"`
}

type policyRestModel struct {
	Description   types.String `tfsdk:"description"`
	Name          types.String `tfsdk:"name"`
	Priority      types.String `tfsdk:"priority"`
	Rpom          types.String `tfsdk:"rpom"`
	Rpo           types.String `tfsdk:"rpo"`
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

// type complianceSettingsRestModel struct {
// 	Policy               *backupdr.PolicyRest `tfsdk:"policy"`
// 	WarnThresholdType    types.String         `tfsdk:"warn_threshold_type"`
// 	WarnThresholdCustom  types.Int64          `tfsdk:"warn_threshold_custom"`
// 	ErrorThresholdType   types.String         `tfsdk:"error_threshold_type"`
// 	ErrorThresholdCustom types.Int64          `tfsdk:"error_threshold_custom"`
// 	ID                   types.String         `tfsdk:"id"`
// 	Href                 types.String         `tfsdk:"href"`
// 	Syncdate             types.Int64          `tfsdk:"syncdate"`
// 	Stale                types.Bool           `tfsdk:"stale"`
// }

// ###########################################
// #########   backupdr_profile   ############
// ###########################################

type profileResourceModel struct {
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
	Dedupasyncnode types.String                  `tfsdk:"dedupasyncnode"`
	Vaultpool      *profileDiskPoolResourceModel `tfsdk:"vaultpool"`
	Vaultpool2     *profileDiskPoolResourceModel `tfsdk:"vaultpool2"`
	Vaultpool3     *profileDiskPoolResourceModel `tfsdk:"vaultpool3"`
	Vaultpool4     *profileDiskPoolResourceModel `tfsdk:"vaultpool4"`
	Createdate     types.Int64                   `tfsdk:"createdate"`
	Localnode      types.String                  `tfsdk:"localnode"`
	// Orglist         []OrganizationRest   `tfsdk:"orglist"`
	// CloudCredential *CloudCredentialRest `tfsdk:"cloudCredential"`
	ID       types.String `tfsdk:"id"`
	Href     types.String `tfsdk:"href"`
	Syncdate types.Int64  `tfsdk:"syncdate"`
	Stale    types.Bool   `tfsdk:"stale"`
}

type profileResourceRefModel struct {
	Name     types.String `tfsdk:"name"`
	Cid      types.String `tfsdk:"cid"`
	ID       types.String `tfsdk:"id"`
	Href     types.String `tfsdk:"href"`
	Syncdate types.Int64  `tfsdk:"syncdate"`
	Stale    types.Bool   `tfsdk:"stale"`
}

// DiskPoolResourceModel represent diskpool object
type profileDiskPoolResourceModel struct {
	Name types.String `tfsdk:"name"`
	ID   types.String `tfsdk:"id"`
	Href types.String `tfsdk:"href"`
}

// ###########################################
// #########  backupdr_diskpool   ############
// ###########################################

type diskPoolResourceModel struct {
	Name                types.String        `tfsdk:"name"`
	Pooltype            types.String        `tfsdk:"pooltype"`
	Cluster             *ClusterRest        `tfsdk:"cluster"`
	ApplianceClusterID  types.String        `tfsdk:"appliance_clusterid"`
	Properties          []keyValueRestModel `tfsdk:"properties"`
	Vaultprops          *vaultPropsRest     `tfsdk:"vaultprops"`
	Usedefaultsa        types.Bool          `tfsdk:"usedefaultsa"`
	Immutable           types.Bool          `tfsdk:"immutable"`
	Metadataonly        types.Bool          `tfsdk:"metadataonly"`
	State               types.String        `tfsdk:"state"`
	Srcid               types.String        `tfsdk:"srcid"`
	Status              types.String        `tfsdk:"status"`
	Mdiskgrp            types.String        `tfsdk:"mdiskgrp"`
	Modifydate          types.Int64         `tfsdk:"modifydate"`
	Warnpct             types.Int64         `tfsdk:"warnpct"`
	Safepct             types.Int64         `tfsdk:"safepct"`
	Udsuid              types.Int64         `tfsdk:"udsuid"`
	FreeMb              types.Int64         `tfsdk:"free_mb"`
	UsageMb             types.Int64         `tfsdk:"usage_mb"`
	CapacityMb          types.Int64         `tfsdk:"capacity_mb"`
	Pct                 types.Float64       `tfsdk:"pct"`
	Pooltypedisplayname types.String        `tfsdk:"pooltypedisplayname"`
	ID                  types.String        `tfsdk:"id"`
	Href                types.String        `tfsdk:"href"`
	Syncdate            types.Int64         `tfsdk:"syncdate"`
	Stale               types.Bool          `tfsdk:"stale"`
}

type keyValueRestModel struct {
	Value types.String `tfsdk:"value"`
	Key   types.String `tfsdk:"key"`
}

type vaultPropsRest struct {
	Bucket      types.String `tfsdk:"bucket"`
	Compression types.Bool   `tfsdk:"compression"`
	Region      types.String `tfsdk:"region"`
	ID          types.String `tfsdk:"id"`
	Href        types.String `tfsdk:"href"`
	Syncdate    types.Int64  `tfsdk:"syncdate"`
	Stale       types.Bool   `tfsdk:"stale"`
}

// ###########################################
// ########  backupdr_appliance   ############
// ###########################################

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
