/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type BfdObservation struct {

	// The minimum interval, in milliseconds, between BFD control packets
	// received from the peer router. The actual value is negotiated
	// between the two routers and is equal to the greater of this value
	// and the transmit interval of the other router. If set, this value
	// must be between 1000 and 30000.
	MinReceiveInterval *float64 `json:"minReceiveInterval,omitempty" tf:"min_receive_interval,omitempty"`

	// The minimum interval, in milliseconds, between BFD control packets
	// transmitted to the peer router. The actual value is negotiated
	// between the two routers and is equal to the greater of this value
	// and the corresponding receive interval of the other router. If set,
	// this value must be between 1000 and 30000.
	MinTransmitInterval *float64 `json:"minTransmitInterval,omitempty" tf:"min_transmit_interval,omitempty"`

	// The number of consecutive BFD packets that must be missed before
	// BFD declares that a peer is unavailable. If set, the value must
	// be a value between 5 and 16.
	Multiplier *float64 `json:"multiplier,omitempty" tf:"multiplier,omitempty"`

	// The BFD session initialization mode for this BGP peer.
	// If set to ACTIVE, the Cloud Router will initiate the BFD session
	// for this BGP peer. If set to PASSIVE, the Cloud Router will wait
	// for the peer router to initiate the BFD session for this BGP peer.
	// If set to DISABLED, BFD is disabled for this BGP peer.
	// Possible values are: ACTIVE, DISABLED, PASSIVE.
	SessionInitializationMode *string `json:"sessionInitializationMode,omitempty" tf:"session_initialization_mode,omitempty"`
}

type BfdParameters struct {

	// The minimum interval, in milliseconds, between BFD control packets
	// received from the peer router. The actual value is negotiated
	// between the two routers and is equal to the greater of this value
	// and the transmit interval of the other router. If set, this value
	// must be between 1000 and 30000.
	// +kubebuilder:validation:Optional
	MinReceiveInterval *float64 `json:"minReceiveInterval,omitempty" tf:"min_receive_interval,omitempty"`

	// The minimum interval, in milliseconds, between BFD control packets
	// transmitted to the peer router. The actual value is negotiated
	// between the two routers and is equal to the greater of this value
	// and the corresponding receive interval of the other router. If set,
	// this value must be between 1000 and 30000.
	// +kubebuilder:validation:Optional
	MinTransmitInterval *float64 `json:"minTransmitInterval,omitempty" tf:"min_transmit_interval,omitempty"`

	// The number of consecutive BFD packets that must be missed before
	// BFD declares that a peer is unavailable. If set, the value must
	// be a value between 5 and 16.
	// +kubebuilder:validation:Optional
	Multiplier *float64 `json:"multiplier,omitempty" tf:"multiplier,omitempty"`

	// The BFD session initialization mode for this BGP peer.
	// If set to ACTIVE, the Cloud Router will initiate the BFD session
	// for this BGP peer. If set to PASSIVE, the Cloud Router will wait
	// for the peer router to initiate the BFD session for this BGP peer.
	// If set to DISABLED, BFD is disabled for this BGP peer.
	// Possible values are: ACTIVE, DISABLED, PASSIVE.
	// +kubebuilder:validation:Required
	SessionInitializationMode *string `json:"sessionInitializationMode" tf:"session_initialization_mode,omitempty"`
}

type RouterPeerAdvertisedIPRangesObservation struct {

	// User-specified description for the IP range.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The IP range to advertise. The value must be a
	// CIDR-formatted string.
	Range *string `json:"range,omitempty" tf:"range,omitempty"`
}

type RouterPeerAdvertisedIPRangesParameters struct {

	// User-specified description for the IP range.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The IP range to advertise. The value must be a
	// CIDR-formatted string.
	// +kubebuilder:validation:Required
	Range *string `json:"range" tf:"range,omitempty"`
}

type RouterPeerObservation struct {

	// User-specified flag to indicate which mode to use for advertisement.
	// Valid values of this enum field are: DEFAULT, CUSTOM
	// Default value is DEFAULT.
	// Possible values are: DEFAULT, CUSTOM.
	AdvertiseMode *string `json:"advertiseMode,omitempty" tf:"advertise_mode,omitempty"`

	// User-specified list of prefix groups to advertise in custom
	// mode, which can take one of the following options:
	AdvertisedGroups []*string `json:"advertisedGroups,omitempty" tf:"advertised_groups,omitempty"`

	// User-specified list of individual IP ranges to advertise in
	// custom mode. This field can only be populated if advertiseMode
	// is CUSTOM and is advertised to all peers of the router. These IP
	// ranges will be advertised in addition to any specified groups.
	// Leave this field blank to advertise no custom IP ranges.
	// Structure is documented below.
	AdvertisedIPRanges []RouterPeerAdvertisedIPRangesObservation `json:"advertisedIpRanges,omitempty" tf:"advertised_ip_ranges,omitempty"`

	// The priority of routes advertised to this BGP peer.
	// Where there is more than one matching route of maximum
	// length, the routes with the lowest priority value win.
	AdvertisedRoutePriority *float64 `json:"advertisedRoutePriority,omitempty" tf:"advertised_route_priority,omitempty"`

	// BFD configuration for the BGP peering.
	// Structure is documented below.
	Bfd []BfdObservation `json:"bfd,omitempty" tf:"bfd,omitempty"`

	// The status of the BGP peer connection. If set to false, any active session
	// with the peer is terminated and all associated routing information is removed.
	// If set to true, the peer connection can be established with routing information.
	// The default is true.
	Enable *bool `json:"enable,omitempty" tf:"enable,omitempty"`

	// Enable IPv6 traffic over BGP Peer. If not specified, it is disabled by default.
	EnableIPv6 *bool `json:"enableIpv6,omitempty" tf:"enable_ipv6,omitempty"`

	// an identifier for the resource with format projects/{{project}}/regions/{{region}}/routers/{{router}}/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// IP address of the interface inside Google Cloud Platform.
	// Only IPv4 is supported.
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// IPv6 address of the interface inside Google Cloud Platform.
	// The address must be in the range 2600:2d00:0:2::/64 or 2600:2d00:0:3::/64.
	// If you do not specify the next hop addresses, Google Cloud automatically
	// assigns unused addresses from the 2600:2d00:0:2::/64 or 2600:2d00:0:3::/64 range for you.
	IPv6NexthopAddress *string `json:"ipv6NexthopAddress,omitempty" tf:"ipv6_nexthop_address,omitempty"`

	// Name of the interface the BGP peer is associated with.
	Interface *string `json:"interface,omitempty" tf:"interface,omitempty"`

	// The resource that configures and manages this BGP peer.
	ManagementType *string `json:"managementType,omitempty" tf:"management_type,omitempty"`

	// Peer BGP Autonomous System Number (ASN).
	// Each BGP interface may use a different value.
	PeerAsn *float64 `json:"peerAsn,omitempty" tf:"peer_asn,omitempty"`

	// IP address of the BGP interface outside Google Cloud Platform.
	// Only IPv4 is supported.
	PeerIPAddress *string `json:"peerIpAddress,omitempty" tf:"peer_ip_address,omitempty"`

	// IPv6 address of the BGP interface outside Google Cloud Platform.
	// The address must be in the range 2600:2d00:0:2::/64 or 2600:2d00:0:3::/64.
	// If you do not specify the next hop addresses, Google Cloud automatically
	// assigns unused addresses from the 2600:2d00:0:2::/64 or 2600:2d00:0:3::/64 range for you.
	PeerIPv6NexthopAddress *string `json:"peerIpv6NexthopAddress,omitempty" tf:"peer_ipv6_nexthop_address,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Region where the router and BgpPeer reside.
	// If it is not provided, the provider region is used.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The name of the Cloud Router in which this BgpPeer will be configured.
	Router *string `json:"router,omitempty" tf:"router,omitempty"`

	// The URI of the VM instance that is used as third-party router appliances
	// such as Next Gen Firewalls, Virtual Routers, or Router Appliances.
	// The VM instance must be located in zones contained in the same region as
	// this Cloud Router. The VM instance is the peer side of the BGP session.
	RouterApplianceInstance *string `json:"routerApplianceInstance,omitempty" tf:"router_appliance_instance,omitempty"`
}

type RouterPeerParameters struct {

	// User-specified flag to indicate which mode to use for advertisement.
	// Valid values of this enum field are: DEFAULT, CUSTOM
	// Default value is DEFAULT.
	// Possible values are: DEFAULT, CUSTOM.
	// +kubebuilder:validation:Optional
	AdvertiseMode *string `json:"advertiseMode,omitempty" tf:"advertise_mode,omitempty"`

	// User-specified list of prefix groups to advertise in custom
	// mode, which can take one of the following options:
	// +kubebuilder:validation:Optional
	AdvertisedGroups []*string `json:"advertisedGroups,omitempty" tf:"advertised_groups,omitempty"`

	// User-specified list of individual IP ranges to advertise in
	// custom mode. This field can only be populated if advertiseMode
	// is CUSTOM and is advertised to all peers of the router. These IP
	// ranges will be advertised in addition to any specified groups.
	// Leave this field blank to advertise no custom IP ranges.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	AdvertisedIPRanges []RouterPeerAdvertisedIPRangesParameters `json:"advertisedIpRanges,omitempty" tf:"advertised_ip_ranges,omitempty"`

	// The priority of routes advertised to this BGP peer.
	// Where there is more than one matching route of maximum
	// length, the routes with the lowest priority value win.
	// +kubebuilder:validation:Optional
	AdvertisedRoutePriority *float64 `json:"advertisedRoutePriority,omitempty" tf:"advertised_route_priority,omitempty"`

	// BFD configuration for the BGP peering.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Bfd []BfdParameters `json:"bfd,omitempty" tf:"bfd,omitempty"`

	// The status of the BGP peer connection. If set to false, any active session
	// with the peer is terminated and all associated routing information is removed.
	// If set to true, the peer connection can be established with routing information.
	// The default is true.
	// +kubebuilder:validation:Optional
	Enable *bool `json:"enable,omitempty" tf:"enable,omitempty"`

	// Enable IPv6 traffic over BGP Peer. If not specified, it is disabled by default.
	// +kubebuilder:validation:Optional
	EnableIPv6 *bool `json:"enableIpv6,omitempty" tf:"enable_ipv6,omitempty"`

	// IP address of the interface inside Google Cloud Platform.
	// Only IPv4 is supported.
	// +kubebuilder:validation:Optional
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// IPv6 address of the interface inside Google Cloud Platform.
	// The address must be in the range 2600:2d00:0:2::/64 or 2600:2d00:0:3::/64.
	// If you do not specify the next hop addresses, Google Cloud automatically
	// assigns unused addresses from the 2600:2d00:0:2::/64 or 2600:2d00:0:3::/64 range for you.
	// +kubebuilder:validation:Optional
	IPv6NexthopAddress *string `json:"ipv6NexthopAddress,omitempty" tf:"ipv6_nexthop_address,omitempty"`

	// Name of the interface the BGP peer is associated with.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.RouterInterface
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("name",false)
	// +kubebuilder:validation:Optional
	Interface *string `json:"interface,omitempty" tf:"interface,omitempty"`

	// Reference to a RouterInterface in compute to populate interface.
	// +kubebuilder:validation:Optional
	InterfaceRef *v1.Reference `json:"interfaceRef,omitempty" tf:"-"`

	// Selector for a RouterInterface in compute to populate interface.
	// +kubebuilder:validation:Optional
	InterfaceSelector *v1.Selector `json:"interfaceSelector,omitempty" tf:"-"`

	// Peer BGP Autonomous System Number (ASN).
	// Each BGP interface may use a different value.
	// +kubebuilder:validation:Optional
	PeerAsn *float64 `json:"peerAsn,omitempty" tf:"peer_asn,omitempty"`

	// IP address of the BGP interface outside Google Cloud Platform.
	// Only IPv4 is supported.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.Address
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("address",false)
	// +kubebuilder:validation:Optional
	PeerIPAddress *string `json:"peerIpAddress,omitempty" tf:"peer_ip_address,omitempty"`

	// Reference to a Address in compute to populate peerIpAddress.
	// +kubebuilder:validation:Optional
	PeerIPAddressRef *v1.Reference `json:"peerIpAddressRef,omitempty" tf:"-"`

	// Selector for a Address in compute to populate peerIpAddress.
	// +kubebuilder:validation:Optional
	PeerIPAddressSelector *v1.Selector `json:"peerIpAddressSelector,omitempty" tf:"-"`

	// IPv6 address of the BGP interface outside Google Cloud Platform.
	// The address must be in the range 2600:2d00:0:2::/64 or 2600:2d00:0:3::/64.
	// If you do not specify the next hop addresses, Google Cloud automatically
	// assigns unused addresses from the 2600:2d00:0:2::/64 or 2600:2d00:0:3::/64 range for you.
	// +kubebuilder:validation:Optional
	PeerIPv6NexthopAddress *string `json:"peerIpv6NexthopAddress,omitempty" tf:"peer_ipv6_nexthop_address,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Region where the router and BgpPeer reside.
	// If it is not provided, the provider region is used.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.Router
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("region",false)
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Reference to a Router in compute to populate region.
	// +kubebuilder:validation:Optional
	RegionRef *v1.Reference `json:"regionRef,omitempty" tf:"-"`

	// Selector for a Router in compute to populate region.
	// +kubebuilder:validation:Optional
	RegionSelector *v1.Selector `json:"regionSelector,omitempty" tf:"-"`

	// The name of the Cloud Router in which this BgpPeer will be configured.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.Router
	// +kubebuilder:validation:Optional
	Router *string `json:"router,omitempty" tf:"router,omitempty"`

	// The URI of the VM instance that is used as third-party router appliances
	// such as Next Gen Firewalls, Virtual Routers, or Router Appliances.
	// The VM instance must be located in zones contained in the same region as
	// this Cloud Router. The VM instance is the peer side of the BGP session.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.Instance
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("self_link",true)
	// +kubebuilder:validation:Optional
	RouterApplianceInstance *string `json:"routerApplianceInstance,omitempty" tf:"router_appliance_instance,omitempty"`

	// Reference to a Instance in compute to populate routerApplianceInstance.
	// +kubebuilder:validation:Optional
	RouterApplianceInstanceRef *v1.Reference `json:"routerApplianceInstanceRef,omitempty" tf:"-"`

	// Selector for a Instance in compute to populate routerApplianceInstance.
	// +kubebuilder:validation:Optional
	RouterApplianceInstanceSelector *v1.Selector `json:"routerApplianceInstanceSelector,omitempty" tf:"-"`

	// Reference to a Router in compute to populate router.
	// +kubebuilder:validation:Optional
	RouterRef *v1.Reference `json:"routerRef,omitempty" tf:"-"`

	// Selector for a Router in compute to populate router.
	// +kubebuilder:validation:Optional
	RouterSelector *v1.Selector `json:"routerSelector,omitempty" tf:"-"`
}

// RouterPeerSpec defines the desired state of RouterPeer
type RouterPeerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RouterPeerParameters `json:"forProvider"`
}

// RouterPeerStatus defines the observed state of RouterPeer.
type RouterPeerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RouterPeerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RouterPeer is the Schema for the RouterPeers API. BGP information that must be configured into the routing stack to establish BGP peering.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type RouterPeer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.peerAsn)",message="peerAsn is a required parameter"
	Spec   RouterPeerSpec   `json:"spec"`
	Status RouterPeerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RouterPeerList contains a list of RouterPeers
type RouterPeerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RouterPeer `json:"items"`
}

// Repository type metadata.
var (
	RouterPeer_Kind             = "RouterPeer"
	RouterPeer_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RouterPeer_Kind}.String()
	RouterPeer_KindAPIVersion   = RouterPeer_Kind + "." + CRDGroupVersion.String()
	RouterPeer_GroupVersionKind = CRDGroupVersion.WithKind(RouterPeer_Kind)
)

func init() {
	SchemeBuilder.Register(&RouterPeer{}, &RouterPeerList{})
}
