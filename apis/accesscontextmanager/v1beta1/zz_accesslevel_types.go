// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AccessLevelInitParameters struct {

	// A set of predefined conditions for the access level and a combining function.
	// Structure is documented below.
	Basic []BasicInitParameters `json:"basic,omitempty" tf:"basic,omitempty"`

	// Custom access level conditions are set using the Cloud Common Expression Language to represent the necessary conditions for the level to apply to a request.
	// See CEL spec at: https://github.com/google/cel-spec.
	// Structure is documented below.
	Custom []CustomInitParameters `json:"custom,omitempty" tf:"custom,omitempty"`

	// Description of the AccessLevel and its use. Does not affect behavior.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Resource name for the Access Level. The short_name component must begin
	// with a letter and only include alphanumeric and '_'.
	// Format: accessPolicies/{policy_id}/accessLevels/{short_name}
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The AccessPolicy this AccessLevel lives in.
	// Format: accessPolicies/{policy_id}
	Parent *string `json:"parent,omitempty" tf:"parent,omitempty"`

	// Human readable title. Must be unique within the Policy.
	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type AccessLevelObservation struct {

	// A set of predefined conditions for the access level and a combining function.
	// Structure is documented below.
	Basic []BasicObservation `json:"basic,omitempty" tf:"basic,omitempty"`

	// Custom access level conditions are set using the Cloud Common Expression Language to represent the necessary conditions for the level to apply to a request.
	// See CEL spec at: https://github.com/google/cel-spec.
	// Structure is documented below.
	Custom []CustomObservation `json:"custom,omitempty" tf:"custom,omitempty"`

	// Description of the AccessLevel and its use. Does not affect behavior.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// an identifier for the resource with format {{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Resource name for the Access Level. The short_name component must begin
	// with a letter and only include alphanumeric and '_'.
	// Format: accessPolicies/{policy_id}/accessLevels/{short_name}
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The AccessPolicy this AccessLevel lives in.
	// Format: accessPolicies/{policy_id}
	Parent *string `json:"parent,omitempty" tf:"parent,omitempty"`

	// Human readable title. Must be unique within the Policy.
	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type AccessLevelParameters struct {

	// A set of predefined conditions for the access level and a combining function.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Basic []BasicParameters `json:"basic,omitempty" tf:"basic,omitempty"`

	// Custom access level conditions are set using the Cloud Common Expression Language to represent the necessary conditions for the level to apply to a request.
	// See CEL spec at: https://github.com/google/cel-spec.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Custom []CustomParameters `json:"custom,omitempty" tf:"custom,omitempty"`

	// Description of the AccessLevel and its use. Does not affect behavior.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Resource name for the Access Level. The short_name component must begin
	// with a letter and only include alphanumeric and '_'.
	// Format: accessPolicies/{policy_id}/accessLevels/{short_name}
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The AccessPolicy this AccessLevel lives in.
	// Format: accessPolicies/{policy_id}
	// +kubebuilder:validation:Optional
	Parent *string `json:"parent,omitempty" tf:"parent,omitempty"`

	// Human readable title. Must be unique within the Policy.
	// +kubebuilder:validation:Optional
	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type BasicInitParameters struct {

	// How the conditions list should be combined to determine if a request
	// is granted this AccessLevel. If AND is used, each Condition in
	// conditions must be satisfied for the AccessLevel to be applied. If
	// OR is used, at least one Condition in conditions must be satisfied
	// for the AccessLevel to be applied.
	// Default value is AND.
	// Possible values are: AND, OR.
	CombiningFunction *string `json:"combiningFunction,omitempty" tf:"combining_function,omitempty"`

	// A set of requirements for the AccessLevel to be granted.
	// Structure is documented below.
	Conditions []ConditionsInitParameters `json:"conditions,omitempty" tf:"conditions,omitempty"`
}

type BasicObservation struct {

	// How the conditions list should be combined to determine if a request
	// is granted this AccessLevel. If AND is used, each Condition in
	// conditions must be satisfied for the AccessLevel to be applied. If
	// OR is used, at least one Condition in conditions must be satisfied
	// for the AccessLevel to be applied.
	// Default value is AND.
	// Possible values are: AND, OR.
	CombiningFunction *string `json:"combiningFunction,omitempty" tf:"combining_function,omitempty"`

	// A set of requirements for the AccessLevel to be granted.
	// Structure is documented below.
	Conditions []ConditionsObservation `json:"conditions,omitempty" tf:"conditions,omitempty"`
}

type BasicParameters struct {

	// How the conditions list should be combined to determine if a request
	// is granted this AccessLevel. If AND is used, each Condition in
	// conditions must be satisfied for the AccessLevel to be applied. If
	// OR is used, at least one Condition in conditions must be satisfied
	// for the AccessLevel to be applied.
	// Default value is AND.
	// Possible values are: AND, OR.
	// +kubebuilder:validation:Optional
	CombiningFunction *string `json:"combiningFunction,omitempty" tf:"combining_function,omitempty"`

	// A set of requirements for the AccessLevel to be granted.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Conditions []ConditionsParameters `json:"conditions" tf:"conditions,omitempty"`
}

type ConditionsInitParameters struct {

	// Device specific restrictions, all restrictions must hold for
	// the Condition to be true. If not specified, all devices are
	// allowed.
	// Structure is documented below.
	DevicePolicy []DevicePolicyInitParameters `json:"devicePolicy,omitempty" tf:"device_policy,omitempty"`

	// A list of CIDR block IP subnetwork specification. May be IPv4
	// or IPv6.
	// Note that for a CIDR IP address block, the specified IP address
	// portion must be properly truncated (i.e. all the host bits must
	// be zero) or the input is considered malformed. For example,
	// "192.0.2.0/24" is accepted but "192.0.2.1/24" is not. Similarly,
	// for IPv6, "2001:db8::/32" is accepted whereas "2001:db8::1/32"
	// is not. The originating IP of a request must be in one of the
	// listed subnets in order for this Condition to be true.
	// If empty, all IP addresses are allowed.
	IPSubnetworks []*string `json:"ipSubnetworks,omitempty" tf:"ip_subnetworks,omitempty"`

	// An allowed list of members (users, service accounts).
	// Using groups is not supported yet.
	// The signed-in user originating the request must be a part of one
	// of the provided members. If not specified, a request may come
	// from any user (logged in/not logged in, not present in any
	// groups, etc.).
	// Formats: user:{emailid}, serviceAccount:{emailid}
	Members []*string `json:"members,omitempty" tf:"members,omitempty"`

	// Whether to negate the Condition. If true, the Condition becomes
	// a NAND over its non-empty fields, each field must be false for
	// the Condition overall to be satisfied. Defaults to false.
	Negate *bool `json:"negate,omitempty" tf:"negate,omitempty"`

	// The request must originate from one of the provided
	// countries/regions.
	// Format: A valid ISO 3166-1 alpha-2 code.
	Regions []*string `json:"regions,omitempty" tf:"regions,omitempty"`

	// A list of other access levels defined in the same Policy,
	// referenced by resource name. Referencing an AccessLevel which
	// does not exist is an error. All access levels listed must be
	// granted for the Condition to be true.
	// Format: accessPolicies/{policy_id}/accessLevels/{short_name}
	RequiredAccessLevels []*string `json:"requiredAccessLevels,omitempty" tf:"required_access_levels,omitempty"`

	// The request must originate from one of the provided VPC networks in Google Cloud. Cannot specify this field together with ip_subnetworks.
	// Structure is documented below.
	VPCNetworkSources []VPCNetworkSourcesInitParameters `json:"vpcNetworkSources,omitempty" tf:"vpc_network_sources,omitempty"`
}

type ConditionsObservation struct {

	// Device specific restrictions, all restrictions must hold for
	// the Condition to be true. If not specified, all devices are
	// allowed.
	// Structure is documented below.
	DevicePolicy []DevicePolicyObservation `json:"devicePolicy,omitempty" tf:"device_policy,omitempty"`

	// A list of CIDR block IP subnetwork specification. May be IPv4
	// or IPv6.
	// Note that for a CIDR IP address block, the specified IP address
	// portion must be properly truncated (i.e. all the host bits must
	// be zero) or the input is considered malformed. For example,
	// "192.0.2.0/24" is accepted but "192.0.2.1/24" is not. Similarly,
	// for IPv6, "2001:db8::/32" is accepted whereas "2001:db8::1/32"
	// is not. The originating IP of a request must be in one of the
	// listed subnets in order for this Condition to be true.
	// If empty, all IP addresses are allowed.
	IPSubnetworks []*string `json:"ipSubnetworks,omitempty" tf:"ip_subnetworks,omitempty"`

	// An allowed list of members (users, service accounts).
	// Using groups is not supported yet.
	// The signed-in user originating the request must be a part of one
	// of the provided members. If not specified, a request may come
	// from any user (logged in/not logged in, not present in any
	// groups, etc.).
	// Formats: user:{emailid}, serviceAccount:{emailid}
	Members []*string `json:"members,omitempty" tf:"members,omitempty"`

	// Whether to negate the Condition. If true, the Condition becomes
	// a NAND over its non-empty fields, each field must be false for
	// the Condition overall to be satisfied. Defaults to false.
	Negate *bool `json:"negate,omitempty" tf:"negate,omitempty"`

	// The request must originate from one of the provided
	// countries/regions.
	// Format: A valid ISO 3166-1 alpha-2 code.
	Regions []*string `json:"regions,omitempty" tf:"regions,omitempty"`

	// A list of other access levels defined in the same Policy,
	// referenced by resource name. Referencing an AccessLevel which
	// does not exist is an error. All access levels listed must be
	// granted for the Condition to be true.
	// Format: accessPolicies/{policy_id}/accessLevels/{short_name}
	RequiredAccessLevels []*string `json:"requiredAccessLevels,omitempty" tf:"required_access_levels,omitempty"`

	// The request must originate from one of the provided VPC networks in Google Cloud. Cannot specify this field together with ip_subnetworks.
	// Structure is documented below.
	VPCNetworkSources []VPCNetworkSourcesObservation `json:"vpcNetworkSources,omitempty" tf:"vpc_network_sources,omitempty"`
}

type ConditionsParameters struct {

	// Device specific restrictions, all restrictions must hold for
	// the Condition to be true. If not specified, all devices are
	// allowed.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	DevicePolicy []DevicePolicyParameters `json:"devicePolicy,omitempty" tf:"device_policy,omitempty"`

	// A list of CIDR block IP subnetwork specification. May be IPv4
	// or IPv6.
	// Note that for a CIDR IP address block, the specified IP address
	// portion must be properly truncated (i.e. all the host bits must
	// be zero) or the input is considered malformed. For example,
	// "192.0.2.0/24" is accepted but "192.0.2.1/24" is not. Similarly,
	// for IPv6, "2001:db8::/32" is accepted whereas "2001:db8::1/32"
	// is not. The originating IP of a request must be in one of the
	// listed subnets in order for this Condition to be true.
	// If empty, all IP addresses are allowed.
	// +kubebuilder:validation:Optional
	IPSubnetworks []*string `json:"ipSubnetworks,omitempty" tf:"ip_subnetworks,omitempty"`

	// An allowed list of members (users, service accounts).
	// Using groups is not supported yet.
	// The signed-in user originating the request must be a part of one
	// of the provided members. If not specified, a request may come
	// from any user (logged in/not logged in, not present in any
	// groups, etc.).
	// Formats: user:{emailid}, serviceAccount:{emailid}
	// +kubebuilder:validation:Optional
	Members []*string `json:"members,omitempty" tf:"members,omitempty"`

	// Whether to negate the Condition. If true, the Condition becomes
	// a NAND over its non-empty fields, each field must be false for
	// the Condition overall to be satisfied. Defaults to false.
	// +kubebuilder:validation:Optional
	Negate *bool `json:"negate,omitempty" tf:"negate,omitempty"`

	// The request must originate from one of the provided
	// countries/regions.
	// Format: A valid ISO 3166-1 alpha-2 code.
	// +kubebuilder:validation:Optional
	Regions []*string `json:"regions,omitempty" tf:"regions,omitempty"`

	// A list of other access levels defined in the same Policy,
	// referenced by resource name. Referencing an AccessLevel which
	// does not exist is an error. All access levels listed must be
	// granted for the Condition to be true.
	// Format: accessPolicies/{policy_id}/accessLevels/{short_name}
	// +kubebuilder:validation:Optional
	RequiredAccessLevels []*string `json:"requiredAccessLevels,omitempty" tf:"required_access_levels,omitempty"`

	// The request must originate from one of the provided VPC networks in Google Cloud. Cannot specify this field together with ip_subnetworks.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	VPCNetworkSources []VPCNetworkSourcesParameters `json:"vpcNetworkSources,omitempty" tf:"vpc_network_sources,omitempty"`
}

type CustomInitParameters struct {

	// Represents a textual expression in the Common Expression Language (CEL) syntax. CEL is a C-like expression language.
	// This page details the objects and attributes that are used to the build the CEL expressions for
	// custom access levels - https://cloud.google.com/access-context-manager/docs/custom-access-level-spec.
	// Structure is documented below.
	Expr []ExprInitParameters `json:"expr,omitempty" tf:"expr,omitempty"`
}

type CustomObservation struct {

	// Represents a textual expression in the Common Expression Language (CEL) syntax. CEL is a C-like expression language.
	// This page details the objects and attributes that are used to the build the CEL expressions for
	// custom access levels - https://cloud.google.com/access-context-manager/docs/custom-access-level-spec.
	// Structure is documented below.
	Expr []ExprObservation `json:"expr,omitempty" tf:"expr,omitempty"`
}

type CustomParameters struct {

	// Represents a textual expression in the Common Expression Language (CEL) syntax. CEL is a C-like expression language.
	// This page details the objects and attributes that are used to the build the CEL expressions for
	// custom access levels - https://cloud.google.com/access-context-manager/docs/custom-access-level-spec.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Expr []ExprParameters `json:"expr" tf:"expr,omitempty"`
}

type DevicePolicyInitParameters struct {

	// A list of allowed device management levels.
	// An empty list allows all management levels.
	// Each value may be one of: MANAGEMENT_UNSPECIFIED, NONE, BASIC, COMPLETE.
	AllowedDeviceManagementLevels []*string `json:"allowedDeviceManagementLevels,omitempty" tf:"allowed_device_management_levels,omitempty"`

	// A list of allowed encryptions statuses.
	// An empty list allows all statuses.
	// Each value may be one of: ENCRYPTION_UNSPECIFIED, ENCRYPTION_UNSUPPORTED, UNENCRYPTED, ENCRYPTED.
	AllowedEncryptionStatuses []*string `json:"allowedEncryptionStatuses,omitempty" tf:"allowed_encryption_statuses,omitempty"`

	// A list of allowed OS versions.
	// An empty list allows all types and all versions.
	// Structure is documented below.
	OsConstraints []OsConstraintsInitParameters `json:"osConstraints,omitempty" tf:"os_constraints,omitempty"`

	// Whether the device needs to be approved by the customer admin.
	RequireAdminApproval *bool `json:"requireAdminApproval,omitempty" tf:"require_admin_approval,omitempty"`

	// Whether the device needs to be corp owned.
	RequireCorpOwned *bool `json:"requireCorpOwned,omitempty" tf:"require_corp_owned,omitempty"`

	// Whether or not screenlock is required for the DevicePolicy
	// to be true. Defaults to false.
	RequireScreenLock *bool `json:"requireScreenLock,omitempty" tf:"require_screen_lock,omitempty"`
}

type DevicePolicyObservation struct {

	// A list of allowed device management levels.
	// An empty list allows all management levels.
	// Each value may be one of: MANAGEMENT_UNSPECIFIED, NONE, BASIC, COMPLETE.
	AllowedDeviceManagementLevels []*string `json:"allowedDeviceManagementLevels,omitempty" tf:"allowed_device_management_levels,omitempty"`

	// A list of allowed encryptions statuses.
	// An empty list allows all statuses.
	// Each value may be one of: ENCRYPTION_UNSPECIFIED, ENCRYPTION_UNSUPPORTED, UNENCRYPTED, ENCRYPTED.
	AllowedEncryptionStatuses []*string `json:"allowedEncryptionStatuses,omitempty" tf:"allowed_encryption_statuses,omitempty"`

	// A list of allowed OS versions.
	// An empty list allows all types and all versions.
	// Structure is documented below.
	OsConstraints []OsConstraintsObservation `json:"osConstraints,omitempty" tf:"os_constraints,omitempty"`

	// Whether the device needs to be approved by the customer admin.
	RequireAdminApproval *bool `json:"requireAdminApproval,omitempty" tf:"require_admin_approval,omitempty"`

	// Whether the device needs to be corp owned.
	RequireCorpOwned *bool `json:"requireCorpOwned,omitempty" tf:"require_corp_owned,omitempty"`

	// Whether or not screenlock is required for the DevicePolicy
	// to be true. Defaults to false.
	RequireScreenLock *bool `json:"requireScreenLock,omitempty" tf:"require_screen_lock,omitempty"`
}

type DevicePolicyParameters struct {

	// A list of allowed device management levels.
	// An empty list allows all management levels.
	// Each value may be one of: MANAGEMENT_UNSPECIFIED, NONE, BASIC, COMPLETE.
	// +kubebuilder:validation:Optional
	AllowedDeviceManagementLevels []*string `json:"allowedDeviceManagementLevels,omitempty" tf:"allowed_device_management_levels,omitempty"`

	// A list of allowed encryptions statuses.
	// An empty list allows all statuses.
	// Each value may be one of: ENCRYPTION_UNSPECIFIED, ENCRYPTION_UNSUPPORTED, UNENCRYPTED, ENCRYPTED.
	// +kubebuilder:validation:Optional
	AllowedEncryptionStatuses []*string `json:"allowedEncryptionStatuses,omitempty" tf:"allowed_encryption_statuses,omitempty"`

	// A list of allowed OS versions.
	// An empty list allows all types and all versions.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	OsConstraints []OsConstraintsParameters `json:"osConstraints,omitempty" tf:"os_constraints,omitempty"`

	// Whether the device needs to be approved by the customer admin.
	// +kubebuilder:validation:Optional
	RequireAdminApproval *bool `json:"requireAdminApproval,omitempty" tf:"require_admin_approval,omitempty"`

	// Whether the device needs to be corp owned.
	// +kubebuilder:validation:Optional
	RequireCorpOwned *bool `json:"requireCorpOwned,omitempty" tf:"require_corp_owned,omitempty"`

	// Whether or not screenlock is required for the DevicePolicy
	// to be true. Defaults to false.
	// +kubebuilder:validation:Optional
	RequireScreenLock *bool `json:"requireScreenLock,omitempty" tf:"require_screen_lock,omitempty"`
}

type ExprInitParameters struct {

	// Description of the expression
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Textual representation of an expression in Common Expression Language syntax.
	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`

	// String indicating the location of the expression for error reporting, e.g. a file name and a position in the file
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Title for the expression, i.e. a short string describing its purpose.
	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type ExprObservation struct {

	// Description of the expression
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Textual representation of an expression in Common Expression Language syntax.
	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`

	// String indicating the location of the expression for error reporting, e.g. a file name and a position in the file
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Title for the expression, i.e. a short string describing its purpose.
	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type ExprParameters struct {

	// Description of the expression
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Textual representation of an expression in Common Expression Language syntax.
	// +kubebuilder:validation:Optional
	Expression *string `json:"expression" tf:"expression,omitempty"`

	// String indicating the location of the expression for error reporting, e.g. a file name and a position in the file
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Title for the expression, i.e. a short string describing its purpose.
	// +kubebuilder:validation:Optional
	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type OsConstraintsInitParameters struct {

	// The minimum allowed OS version. If not set, any version
	// of this OS satisfies the constraint.
	// Format: "major.minor.patch" such as "10.5.301", "9.2.1".
	MinimumVersion *string `json:"minimumVersion,omitempty" tf:"minimum_version,omitempty"`

	// The operating system type of the device.
	// Possible values are: OS_UNSPECIFIED, DESKTOP_MAC, DESKTOP_WINDOWS, DESKTOP_LINUX, DESKTOP_CHROME_OS, ANDROID, IOS.
	OsType *string `json:"osType,omitempty" tf:"os_type,omitempty"`

	// If you specify DESKTOP_CHROME_OS for osType, you can optionally include requireVerifiedChromeOs to require Chrome Verified Access.
	RequireVerifiedChromeOs *bool `json:"requireVerifiedChromeOs,omitempty" tf:"require_verified_chrome_os,omitempty"`
}

type OsConstraintsObservation struct {

	// The minimum allowed OS version. If not set, any version
	// of this OS satisfies the constraint.
	// Format: "major.minor.patch" such as "10.5.301", "9.2.1".
	MinimumVersion *string `json:"minimumVersion,omitempty" tf:"minimum_version,omitempty"`

	// The operating system type of the device.
	// Possible values are: OS_UNSPECIFIED, DESKTOP_MAC, DESKTOP_WINDOWS, DESKTOP_LINUX, DESKTOP_CHROME_OS, ANDROID, IOS.
	OsType *string `json:"osType,omitempty" tf:"os_type,omitempty"`

	// If you specify DESKTOP_CHROME_OS for osType, you can optionally include requireVerifiedChromeOs to require Chrome Verified Access.
	RequireVerifiedChromeOs *bool `json:"requireVerifiedChromeOs,omitempty" tf:"require_verified_chrome_os,omitempty"`
}

type OsConstraintsParameters struct {

	// The minimum allowed OS version. If not set, any version
	// of this OS satisfies the constraint.
	// Format: "major.minor.patch" such as "10.5.301", "9.2.1".
	// +kubebuilder:validation:Optional
	MinimumVersion *string `json:"minimumVersion,omitempty" tf:"minimum_version,omitempty"`

	// The operating system type of the device.
	// Possible values are: OS_UNSPECIFIED, DESKTOP_MAC, DESKTOP_WINDOWS, DESKTOP_LINUX, DESKTOP_CHROME_OS, ANDROID, IOS.
	// +kubebuilder:validation:Optional
	OsType *string `json:"osType" tf:"os_type,omitempty"`

	// If you specify DESKTOP_CHROME_OS for osType, you can optionally include requireVerifiedChromeOs to require Chrome Verified Access.
	// +kubebuilder:validation:Optional
	RequireVerifiedChromeOs *bool `json:"requireVerifiedChromeOs,omitempty" tf:"require_verified_chrome_os,omitempty"`
}

type VPCNetworkSourcesInitParameters struct {

	// Sub networks within a VPC network.
	// Structure is documented below.
	VPCSubnetwork []VPCSubnetworkInitParameters `json:"vpcSubnetwork,omitempty" tf:"vpc_subnetwork,omitempty"`
}

type VPCNetworkSourcesObservation struct {

	// Sub networks within a VPC network.
	// Structure is documented below.
	VPCSubnetwork []VPCSubnetworkObservation `json:"vpcSubnetwork,omitempty" tf:"vpc_subnetwork,omitempty"`
}

type VPCNetworkSourcesParameters struct {

	// Sub networks within a VPC network.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	VPCSubnetwork []VPCSubnetworkParameters `json:"vpcSubnetwork,omitempty" tf:"vpc_subnetwork,omitempty"`
}

type VPCSubnetworkInitParameters struct {

	// Required. Network name to be allowed by this Access Level. Networks of foreign organizations requires compute.network.get permission to be granted to caller.
	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	// CIDR block IP subnetwork specification. Must be IPv4.
	VPCIPSubnetworks []*string `json:"vpcIpSubnetworks,omitempty" tf:"vpc_ip_subnetworks,omitempty"`
}

type VPCSubnetworkObservation struct {

	// Required. Network name to be allowed by this Access Level. Networks of foreign organizations requires compute.network.get permission to be granted to caller.
	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	// CIDR block IP subnetwork specification. Must be IPv4.
	VPCIPSubnetworks []*string `json:"vpcIpSubnetworks,omitempty" tf:"vpc_ip_subnetworks,omitempty"`
}

type VPCSubnetworkParameters struct {

	// Required. Network name to be allowed by this Access Level. Networks of foreign organizations requires compute.network.get permission to be granted to caller.
	// +kubebuilder:validation:Optional
	Network *string `json:"network" tf:"network,omitempty"`

	// CIDR block IP subnetwork specification. Must be IPv4.
	// +kubebuilder:validation:Optional
	VPCIPSubnetworks []*string `json:"vpcIpSubnetworks,omitempty" tf:"vpc_ip_subnetworks,omitempty"`
}

// AccessLevelSpec defines the desired state of AccessLevel
type AccessLevelSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AccessLevelParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider AccessLevelInitParameters `json:"initProvider,omitempty"`
}

// AccessLevelStatus defines the observed state of AccessLevel.
type AccessLevelStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AccessLevelObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// AccessLevel is the Schema for the AccessLevels API. An AccessLevel is a label that can be applied to requests to GCP services, along with a list of requirements necessary for the label to be applied.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type AccessLevel struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.parent) || (has(self.initProvider) && has(self.initProvider.parent))",message="spec.forProvider.parent is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.title) || (has(self.initProvider) && has(self.initProvider.title))",message="spec.forProvider.title is a required parameter"
	Spec   AccessLevelSpec   `json:"spec"`
	Status AccessLevelStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AccessLevelList contains a list of AccessLevels
type AccessLevelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AccessLevel `json:"items"`
}

// Repository type metadata.
var (
	AccessLevel_Kind             = "AccessLevel"
	AccessLevel_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AccessLevel_Kind}.String()
	AccessLevel_KindAPIVersion   = AccessLevel_Kind + "." + CRDGroupVersion.String()
	AccessLevel_GroupVersionKind = CRDGroupVersion.WithKind(AccessLevel_Kind)
)

func init() {
	SchemeBuilder.Register(&AccessLevel{}, &AccessLevelList{})
}
