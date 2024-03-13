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

type RegionDiskAsyncPrimaryDiskInitParameters struct {

	// Primary disk for asynchronous disk replication.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.RegionDisk
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	Disk *string `json:"disk,omitempty" tf:"disk,omitempty"`

	// Reference to a RegionDisk in compute to populate disk.
	// +kubebuilder:validation:Optional
	DiskRef *v1.Reference `json:"diskRef,omitempty" tf:"-"`

	// Selector for a RegionDisk in compute to populate disk.
	// +kubebuilder:validation:Optional
	DiskSelector *v1.Selector `json:"diskSelector,omitempty" tf:"-"`
}

type RegionDiskAsyncPrimaryDiskObservation struct {

	// Primary disk for asynchronous disk replication.
	Disk *string `json:"disk,omitempty" tf:"disk,omitempty"`
}

type RegionDiskAsyncPrimaryDiskParameters struct {

	// Primary disk for asynchronous disk replication.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.RegionDisk
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	Disk *string `json:"disk,omitempty" tf:"disk,omitempty"`

	// Reference to a RegionDisk in compute to populate disk.
	// +kubebuilder:validation:Optional
	DiskRef *v1.Reference `json:"diskRef,omitempty" tf:"-"`

	// Selector for a RegionDisk in compute to populate disk.
	// +kubebuilder:validation:Optional
	DiskSelector *v1.Selector `json:"diskSelector,omitempty" tf:"-"`
}

type RegionDiskDiskEncryptionKeyInitParameters struct {

	// The name of the encryption key that is stored in Google Cloud KMS.
	KMSKeyName *string `json:"kmsKeyName,omitempty" tf:"kms_key_name,omitempty"`
}

type RegionDiskDiskEncryptionKeyObservation struct {

	// The name of the encryption key that is stored in Google Cloud KMS.
	KMSKeyName *string `json:"kmsKeyName,omitempty" tf:"kms_key_name,omitempty"`

	// (Output)
	// The RFC 4648 base64 encoded SHA-256 hash of the customer-supplied
	// encryption key that protects this resource.
	Sha256 *string `json:"sha256,omitempty" tf:"sha256,omitempty"`
}

type RegionDiskDiskEncryptionKeyParameters struct {

	// The name of the encryption key that is stored in Google Cloud KMS.
	// +kubebuilder:validation:Optional
	KMSKeyName *string `json:"kmsKeyName,omitempty" tf:"kms_key_name,omitempty"`

	// Specifies a 256-bit customer-supplied encryption key, encoded in
	// RFC 4648 base64 to either encrypt or decrypt this resource.
	// Note: This property is sensitive and will not be displayed in the plan.
	// +kubebuilder:validation:Optional
	RawKeySecretRef *v1.SecretKeySelector `json:"rawKeySecretRef,omitempty" tf:"-"`
}

type RegionDiskGuestOsFeaturesInitParameters struct {

	// The type of supported feature. Read Enabling guest operating system features to see a list of available options.
	// Possible values are: MULTI_IP_SUBNET, SECURE_BOOT, SEV_CAPABLE, UEFI_COMPATIBLE, VIRTIO_SCSI_MULTIQUEUE, WINDOWS, GVNIC, SEV_LIVE_MIGRATABLE, SEV_SNP_CAPABLE, SUSPEND_RESUME_COMPATIBLE, TDX_CAPABLE.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type RegionDiskGuestOsFeaturesObservation struct {

	// The type of supported feature. Read Enabling guest operating system features to see a list of available options.
	// Possible values are: MULTI_IP_SUBNET, SECURE_BOOT, SEV_CAPABLE, UEFI_COMPATIBLE, VIRTIO_SCSI_MULTIQUEUE, WINDOWS, GVNIC, SEV_LIVE_MIGRATABLE, SEV_SNP_CAPABLE, SUSPEND_RESUME_COMPATIBLE, TDX_CAPABLE.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type RegionDiskGuestOsFeaturesParameters struct {

	// The type of supported feature. Read Enabling guest operating system features to see a list of available options.
	// Possible values are: MULTI_IP_SUBNET, SECURE_BOOT, SEV_CAPABLE, UEFI_COMPATIBLE, VIRTIO_SCSI_MULTIQUEUE, WINDOWS, GVNIC, SEV_LIVE_MIGRATABLE, SEV_SNP_CAPABLE, SUSPEND_RESUME_COMPATIBLE, TDX_CAPABLE.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type RegionDiskInitParameters struct {

	// A nested object resource
	// Structure is documented below.
	AsyncPrimaryDisk []RegionDiskAsyncPrimaryDiskInitParameters `json:"asyncPrimaryDisk,omitempty" tf:"async_primary_disk,omitempty"`

	// An optional description of this resource. Provide this property when
	// you create the resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Encrypts the disk using a customer-supplied encryption key.
	// After you encrypt a disk with a customer-supplied key, you must
	// provide the same key if you use the disk later (e.g. to create a disk
	// snapshot or an image, or to attach the disk to a virtual machine).
	// Customer-supplied encryption keys do not protect access to metadata of
	// the disk.
	// If you do not provide an encryption key when creating the disk, then
	// the disk will be encrypted using an automatically generated key and
	// you do not need to provide a key to use the disk later.
	// Structure is documented below.
	DiskEncryptionKey []RegionDiskDiskEncryptionKeyInitParameters `json:"diskEncryptionKey,omitempty" tf:"disk_encryption_key,omitempty"`

	// A list of features to enable on the guest operating system.
	// Applicable only for bootable disks.
	// Structure is documented below.
	GuestOsFeatures []RegionDiskGuestOsFeaturesInitParameters `json:"guestOsFeatures,omitempty" tf:"guest_os_features,omitempty"`

	// Labels to apply to this disk.  A list of key->value pairs.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Any applicable license URI.
	Licenses []*string `json:"licenses,omitempty" tf:"licenses,omitempty"`

	// Physical block size of the persistent disk, in bytes. If not present
	// in a request, a default value is used. Currently supported sizes
	// are 4096 and 16384, other sizes may be added in the future.
	// If an unsupported value is requested, the error message will list
	// the supported values for the caller's project.
	PhysicalBlockSizeBytes *float64 `json:"physicalBlockSizeBytes,omitempty" tf:"physical_block_size_bytes,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// URLs of the zones where the disk should be replicated to.
	ReplicaZones []*string `json:"replicaZones,omitempty" tf:"replica_zones,omitempty"`

	// Size of the persistent disk, specified in GB. You can specify this
	// field when creating a persistent disk using the sourceImage or
	// sourceSnapshot parameter, or specify it alone to create an empty
	// persistent disk.
	// If you specify this field along with sourceImage or sourceSnapshot,
	// the value of sizeGb must not be less than the size of the sourceImage
	// or the size of the snapshot.
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// The source snapshot used to create this disk. You can provide this as
	// a partial or full URL to the resource. For example, the following are
	// valid values:
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.Snapshot
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	Snapshot *string `json:"snapshot,omitempty" tf:"snapshot,omitempty"`

	// Reference to a Snapshot in compute to populate snapshot.
	// +kubebuilder:validation:Optional
	SnapshotRef *v1.Reference `json:"snapshotRef,omitempty" tf:"-"`

	// Selector for a Snapshot in compute to populate snapshot.
	// +kubebuilder:validation:Optional
	SnapshotSelector *v1.Selector `json:"snapshotSelector,omitempty" tf:"-"`

	// The source disk used to create this disk. You can provide this as a partial or full URL to the resource.
	// For example, the following are valid values:
	SourceDisk *string `json:"sourceDisk,omitempty" tf:"source_disk,omitempty"`

	// The customer-supplied encryption key of the source snapshot. Required
	// if the source snapshot is protected by a customer-supplied encryption
	// key.
	// Structure is documented below.
	SourceSnapshotEncryptionKey []RegionDiskSourceSnapshotEncryptionKeyInitParameters `json:"sourceSnapshotEncryptionKey,omitempty" tf:"source_snapshot_encryption_key,omitempty"`

	// URL of the disk type resource describing which disk type to use to
	// create the disk. Provide this when creating the disk.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type RegionDiskObservation struct {

	// A nested object resource
	// Structure is documented below.
	AsyncPrimaryDisk []RegionDiskAsyncPrimaryDiskObservation `json:"asyncPrimaryDisk,omitempty" tf:"async_primary_disk,omitempty"`

	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `json:"creationTimestamp,omitempty" tf:"creation_timestamp,omitempty"`

	// An optional description of this resource. Provide this property when
	// you create the resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Encrypts the disk using a customer-supplied encryption key.
	// After you encrypt a disk with a customer-supplied key, you must
	// provide the same key if you use the disk later (e.g. to create a disk
	// snapshot or an image, or to attach the disk to a virtual machine).
	// Customer-supplied encryption keys do not protect access to metadata of
	// the disk.
	// If you do not provide an encryption key when creating the disk, then
	// the disk will be encrypted using an automatically generated key and
	// you do not need to provide a key to use the disk later.
	// Structure is documented below.
	DiskEncryptionKey []RegionDiskDiskEncryptionKeyObservation `json:"diskEncryptionKey,omitempty" tf:"disk_encryption_key,omitempty"`

	// for all of the labels present on the resource.
	// +mapType=granular
	EffectiveLabels map[string]*string `json:"effectiveLabels,omitempty" tf:"effective_labels,omitempty"`

	// A list of features to enable on the guest operating system.
	// Applicable only for bootable disks.
	// Structure is documented below.
	GuestOsFeatures []RegionDiskGuestOsFeaturesObservation `json:"guestOsFeatures,omitempty" tf:"guest_os_features,omitempty"`

	// an identifier for the resource with format projects/{{project}}/regions/{{region}}/disks/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The fingerprint used for optimistic locking of this resource.  Used
	// internally during updates.
	LabelFingerprint *string `json:"labelFingerprint,omitempty" tf:"label_fingerprint,omitempty"`

	// Labels to apply to this disk.  A list of key->value pairs.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Last attach timestamp in RFC3339 text format.
	LastAttachTimestamp *string `json:"lastAttachTimestamp,omitempty" tf:"last_attach_timestamp,omitempty"`

	// Last detach timestamp in RFC3339 text format.
	LastDetachTimestamp *string `json:"lastDetachTimestamp,omitempty" tf:"last_detach_timestamp,omitempty"`

	// Any applicable license URI.
	Licenses []*string `json:"licenses,omitempty" tf:"licenses,omitempty"`

	// Physical block size of the persistent disk, in bytes. If not present
	// in a request, a default value is used. Currently supported sizes
	// are 4096 and 16384, other sizes may be added in the future.
	// If an unsupported value is requested, the error message will list
	// the supported values for the caller's project.
	PhysicalBlockSizeBytes *float64 `json:"physicalBlockSizeBytes,omitempty" tf:"physical_block_size_bytes,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// A reference to the region where the disk resides.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// URLs of the zones where the disk should be replicated to.
	ReplicaZones []*string `json:"replicaZones,omitempty" tf:"replica_zones,omitempty"`

	// The URI of the created resource.
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`

	// Size of the persistent disk, specified in GB. You can specify this
	// field when creating a persistent disk using the sourceImage or
	// sourceSnapshot parameter, or specify it alone to create an empty
	// persistent disk.
	// If you specify this field along with sourceImage or sourceSnapshot,
	// the value of sizeGb must not be less than the size of the sourceImage
	// or the size of the snapshot.
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// The source snapshot used to create this disk. You can provide this as
	// a partial or full URL to the resource. For example, the following are
	// valid values:
	Snapshot *string `json:"snapshot,omitempty" tf:"snapshot,omitempty"`

	// The source disk used to create this disk. You can provide this as a partial or full URL to the resource.
	// For example, the following are valid values:
	SourceDisk *string `json:"sourceDisk,omitempty" tf:"source_disk,omitempty"`

	// The ID value of the disk used to create this image. This value may
	// be used to determine whether the image was taken from the current
	// or a previous instance of a given disk name.
	SourceDiskID *string `json:"sourceDiskId,omitempty" tf:"source_disk_id,omitempty"`

	// The customer-supplied encryption key of the source snapshot. Required
	// if the source snapshot is protected by a customer-supplied encryption
	// key.
	// Structure is documented below.
	SourceSnapshotEncryptionKey []RegionDiskSourceSnapshotEncryptionKeyObservation `json:"sourceSnapshotEncryptionKey,omitempty" tf:"source_snapshot_encryption_key,omitempty"`

	// The unique ID of the snapshot used to create this disk. This value
	// identifies the exact snapshot that was used to create this persistent
	// disk. For example, if you created the persistent disk from a snapshot
	// that was later deleted and recreated under the same name, the source
	// snapshot ID would identify the exact version of the snapshot that was
	// used.
	SourceSnapshotID *string `json:"sourceSnapshotId,omitempty" tf:"source_snapshot_id,omitempty"`

	// The combination of labels configured directly on the resource
	// and default labels configured on the provider.
	// +mapType=granular
	TerraformLabels map[string]*string `json:"terraformLabels,omitempty" tf:"terraform_labels,omitempty"`

	// URL of the disk type resource describing which disk type to use to
	// create the disk. Provide this when creating the disk.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Links to the users of the disk (attached instances) in form:
	// project/zones/zone/instances/instance
	Users []*string `json:"users,omitempty" tf:"users,omitempty"`
}

type RegionDiskParameters struct {

	// A nested object resource
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	AsyncPrimaryDisk []RegionDiskAsyncPrimaryDiskParameters `json:"asyncPrimaryDisk,omitempty" tf:"async_primary_disk,omitempty"`

	// An optional description of this resource. Provide this property when
	// you create the resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Encrypts the disk using a customer-supplied encryption key.
	// After you encrypt a disk with a customer-supplied key, you must
	// provide the same key if you use the disk later (e.g. to create a disk
	// snapshot or an image, or to attach the disk to a virtual machine).
	// Customer-supplied encryption keys do not protect access to metadata of
	// the disk.
	// If you do not provide an encryption key when creating the disk, then
	// the disk will be encrypted using an automatically generated key and
	// you do not need to provide a key to use the disk later.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	DiskEncryptionKey []RegionDiskDiskEncryptionKeyParameters `json:"diskEncryptionKey,omitempty" tf:"disk_encryption_key,omitempty"`

	// A list of features to enable on the guest operating system.
	// Applicable only for bootable disks.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	GuestOsFeatures []RegionDiskGuestOsFeaturesParameters `json:"guestOsFeatures,omitempty" tf:"guest_os_features,omitempty"`

	// Labels to apply to this disk.  A list of key->value pairs.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Any applicable license URI.
	// +kubebuilder:validation:Optional
	Licenses []*string `json:"licenses,omitempty" tf:"licenses,omitempty"`

	// Physical block size of the persistent disk, in bytes. If not present
	// in a request, a default value is used. Currently supported sizes
	// are 4096 and 16384, other sizes may be added in the future.
	// If an unsupported value is requested, the error message will list
	// the supported values for the caller's project.
	// +kubebuilder:validation:Optional
	PhysicalBlockSizeBytes *float64 `json:"physicalBlockSizeBytes,omitempty" tf:"physical_block_size_bytes,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// A reference to the region where the disk resides.
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"region,omitempty"`

	// URLs of the zones where the disk should be replicated to.
	// +kubebuilder:validation:Optional
	ReplicaZones []*string `json:"replicaZones,omitempty" tf:"replica_zones,omitempty"`

	// Size of the persistent disk, specified in GB. You can specify this
	// field when creating a persistent disk using the sourceImage or
	// sourceSnapshot parameter, or specify it alone to create an empty
	// persistent disk.
	// If you specify this field along with sourceImage or sourceSnapshot,
	// the value of sizeGb must not be less than the size of the sourceImage
	// or the size of the snapshot.
	// +kubebuilder:validation:Optional
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// The source snapshot used to create this disk. You can provide this as
	// a partial or full URL to the resource. For example, the following are
	// valid values:
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.Snapshot
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	Snapshot *string `json:"snapshot,omitempty" tf:"snapshot,omitempty"`

	// Reference to a Snapshot in compute to populate snapshot.
	// +kubebuilder:validation:Optional
	SnapshotRef *v1.Reference `json:"snapshotRef,omitempty" tf:"-"`

	// Selector for a Snapshot in compute to populate snapshot.
	// +kubebuilder:validation:Optional
	SnapshotSelector *v1.Selector `json:"snapshotSelector,omitempty" tf:"-"`

	// The source disk used to create this disk. You can provide this as a partial or full URL to the resource.
	// For example, the following are valid values:
	// +kubebuilder:validation:Optional
	SourceDisk *string `json:"sourceDisk,omitempty" tf:"source_disk,omitempty"`

	// The customer-supplied encryption key of the source snapshot. Required
	// if the source snapshot is protected by a customer-supplied encryption
	// key.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	SourceSnapshotEncryptionKey []RegionDiskSourceSnapshotEncryptionKeyParameters `json:"sourceSnapshotEncryptionKey,omitempty" tf:"source_snapshot_encryption_key,omitempty"`

	// URL of the disk type resource describing which disk type to use to
	// create the disk. Provide this when creating the disk.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type RegionDiskSourceSnapshotEncryptionKeyInitParameters struct {

	// Specifies a 256-bit customer-supplied encryption key, encoded in
	// RFC 4648 base64 to either encrypt or decrypt this resource.
	RawKey *string `json:"rawKey,omitempty" tf:"raw_key,omitempty"`
}

type RegionDiskSourceSnapshotEncryptionKeyObservation struct {

	// Specifies a 256-bit customer-supplied encryption key, encoded in
	// RFC 4648 base64 to either encrypt or decrypt this resource.
	RawKey *string `json:"rawKey,omitempty" tf:"raw_key,omitempty"`

	// (Output)
	// The RFC 4648 base64 encoded SHA-256 hash of the customer-supplied
	// encryption key that protects this resource.
	Sha256 *string `json:"sha256,omitempty" tf:"sha256,omitempty"`
}

type RegionDiskSourceSnapshotEncryptionKeyParameters struct {

	// Specifies a 256-bit customer-supplied encryption key, encoded in
	// RFC 4648 base64 to either encrypt or decrypt this resource.
	// +kubebuilder:validation:Optional
	RawKey *string `json:"rawKey,omitempty" tf:"raw_key,omitempty"`
}

// RegionDiskSpec defines the desired state of RegionDisk
type RegionDiskSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RegionDiskParameters `json:"forProvider"`
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
	InitProvider RegionDiskInitParameters `json:"initProvider,omitempty"`
}

// RegionDiskStatus defines the observed state of RegionDisk.
type RegionDiskStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RegionDiskObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// RegionDisk is the Schema for the RegionDisks API. Persistent disks are durable storage devices that function similarly to the physical disks in a desktop or a server.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type RegionDisk struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.replicaZones) || (has(self.initProvider) && has(self.initProvider.replicaZones))",message="spec.forProvider.replicaZones is a required parameter"
	Spec   RegionDiskSpec   `json:"spec"`
	Status RegionDiskStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RegionDiskList contains a list of RegionDisks
type RegionDiskList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RegionDisk `json:"items"`
}

// Repository type metadata.
var (
	RegionDisk_Kind             = "RegionDisk"
	RegionDisk_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RegionDisk_Kind}.String()
	RegionDisk_KindAPIVersion   = RegionDisk_Kind + "." + CRDGroupVersion.String()
	RegionDisk_GroupVersionKind = CRDGroupVersion.WithKind(RegionDisk_Kind)
)

func init() {
	SchemeBuilder.Register(&RegionDisk{}, &RegionDiskList{})
}
