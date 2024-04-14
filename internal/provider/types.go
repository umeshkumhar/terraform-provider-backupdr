package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	backupdr "github.com/umeshkumhar/backupdr-client"
)

type sltRestModel struct {
	ID          types.String `tfsdk:"id"`
	Href        types.String `tfsdk:"href"`
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
	Immutable   types.Bool   `tfsdk:"immutable"`
	OptionHref  types.String `tfsdk:"option_href"`
	PolicyHref  types.String `tfsdk:"policy_href"`
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
	Exclusiontype     types.String `tfsdk:"exclusiontype"`
	Iscontinuous      types.Bool   `tfsdk:"iscontinuous"`
	Starttime         types.String `tfsdk:"starttime"`
	Endtime           types.String `tfsdk:"endtime"`
	Targetvault       types.Int64  `tfsdk:"targetvault"`
	Sourcevault       types.Int64  `tfsdk:"sourcevault"`
	Selection         types.String `tfsdk:"selection"`
	Scheduletype      types.String `tfsdk:"scheduletype"`
	Scheduling        types.String `tfsdk:"scheduling"`
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

type VaultPropsRest struct {
	Bucket      types.String `tfsdk:"bucket"`
	Compression types.Bool   `tfsdk:"compression"`
	Region      types.String `tfsdk:"region"`
	ID          types.String `tfsdk:"id"`
	Href        types.String `tfsdk:"href"`
	Syncdate    types.Int64  `tfsdk:"syncdate"`
	Stale       types.Bool   `tfsdk:"stale"`
}
