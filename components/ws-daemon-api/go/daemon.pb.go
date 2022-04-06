// Copyright (c) 2022 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License-AGPL.txt in the project root for license information.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.0
// source: daemon.proto

package api

import (
	api "github.com/gitpod-io/gitpod/content-service/api"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// WorkspaceContentState describes the availability and reliability of the workspace content
type WorkspaceContentState int32

const (
	// NONE means that there currently is no workspace content and no work is underway to change that.
	WorkspaceContentState_NONE WorkspaceContentState = 0
	// SETTING_UP indicates that the workspace content is currently being produced/checked out/unarchived and is
	// very likely to change. In this state one must not modify or rely on the workspace content.
	WorkspaceContentState_SETTING_UP WorkspaceContentState = 1
	// AVAILABLE indicates that the workspace content is fully present and ready for use.
	WorkspaceContentState_AVAILABLE WorkspaceContentState = 2
	// WRAPPING_UP means that the workspace is being torn down, i.e. a final backup is being produced and the content
	// is deleted locally. In this state one must not modify or rely on the workspace content.
	WorkspaceContentState_WRAPPING_UP WorkspaceContentState = 3
)

// Enum value maps for WorkspaceContentState.
var (
	WorkspaceContentState_name = map[int32]string{
		0: "NONE",
		1: "SETTING_UP",
		2: "AVAILABLE",
		3: "WRAPPING_UP",
	}
	WorkspaceContentState_value = map[string]int32{
		"NONE":        0,
		"SETTING_UP":  1,
		"AVAILABLE":   2,
		"WRAPPING_UP": 3,
	}
)

func (x WorkspaceContentState) Enum() *WorkspaceContentState {
	p := new(WorkspaceContentState)
	*p = x
	return p
}

func (x WorkspaceContentState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WorkspaceContentState) Descriptor() protoreflect.EnumDescriptor {
	return file_daemon_proto_enumTypes[0].Descriptor()
}

func (WorkspaceContentState) Type() protoreflect.EnumType {
	return &file_daemon_proto_enumTypes[0]
}

func (x WorkspaceContentState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WorkspaceContentState.Descriptor instead.
func (WorkspaceContentState) EnumDescriptor() ([]byte, []int) {
	return file_daemon_proto_rawDescGZIP(), []int{0}
}

// InitWorkspaceRequest intialises a new workspace folder in the working area
type InitWorkspaceRequest struct {
	state         protoimpl.MessageState  `json:"state,omitempty"`
	sizeCache     protoimpl.SizeCache     `json:"sizeCache,omitempty"`
	unknownFields protoimpl.UnknownFields `json:"unknownFields,omitempty"`

	// ID is a unique identifier of this workspace. No other workspace with the same name must exist in the realm of this daemon
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Metadata is data associated with this workspace that's required for other parts of Gitpod to function
	Metadata *WorkspaceMetadata `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// Initializer specifies how the workspace is to be initialized
	Initializer *api.WorkspaceInitializer `protobuf:"bytes,3,opt,name=initializer,proto3" json:"initializer,omitempty"`
	// full_workspace_backup means we ignore the initializer and wait for a workspace pod with the given instance ID to
	// appear at our local containerd.
	FullWorkspaceBackup bool `protobuf:"varint,4,opt,name=full_workspace_backup,json=fullWorkspaceBackup,proto3" json:"fullWorkspaceBackup,omitempty"`
	// content_manifest describes the layers that comprise the workspace image content.
	// This manifest is not used to actually download content, but to produce a new manifest for snapshots and backups.
	// This field is ignored if full_workspace_backup is false.
	ContentManifest []byte `protobuf:"bytes,5,opt,name=content_manifest,json=contentManifest,proto3" json:"contentManifest,omitempty"`
	// remote_storage_disabled disables any support for remote storage operations, specifically backups and snapshots.
	// When any such operation is attempted, a FAILED_PRECONDITION error will be the result.
	RemoteStorageDisabled bool `protobuf:"varint,7,opt,name=remote_storage_disabled,json=remoteStorageDisabled,proto3" json:"remoteStorageDisabled,omitempty"`
	// storage_quota_bytes enforces a storage quate for the workspace if set to a value != 0
	StorageQuotaBytes int64 `protobuf:"varint,8,opt,name=storage_quota_bytes,json=storageQuotaBytes,proto3" json:"storageQuotaBytes,omitempty"`
	// persistent_volume_claim means that we use PVC instead of HostPath to mount /workspace folder and content-init
	// happens inside workspacekit instead of in ws-daemon. We also use k8s Snapshots to store\restore workspace content
	// instead of GCS\tar.
	PersistentVolumeClaim bool `protobuf:"varint,9,opt,name=persistent_volume_claim,json=persistentVolumeClaim,proto3" json:"persistentVolumeClaim,omitempty"`
}

func (x *InitWorkspaceRequest) Reset() {
	*x = InitWorkspaceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_daemon_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitWorkspaceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitWorkspaceRequest) ProtoMessage() {}

func (x *InitWorkspaceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_daemon_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitWorkspaceRequest.ProtoReflect.Descriptor instead.
func (*InitWorkspaceRequest) Descriptor() ([]byte, []int) {
	return file_daemon_proto_rawDescGZIP(), []int{0}
}

func (x *InitWorkspaceRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *InitWorkspaceRequest) GetMetadata() *WorkspaceMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *InitWorkspaceRequest) GetInitializer() *api.WorkspaceInitializer {
	if x != nil {
		return x.Initializer
	}
	return nil
}

func (x *InitWorkspaceRequest) GetFullWorkspaceBackup() bool {
	if x != nil {
		return x.FullWorkspaceBackup
	}
	return false
}

func (x *InitWorkspaceRequest) GetContentManifest() []byte {
	if x != nil {
		return x.ContentManifest
	}
	return nil
}

func (x *InitWorkspaceRequest) GetRemoteStorageDisabled() bool {
	if x != nil {
		return x.RemoteStorageDisabled
	}
	return false
}

func (x *InitWorkspaceRequest) GetStorageQuotaBytes() int64 {
	if x != nil {
		return x.StorageQuotaBytes
	}
	return 0
}

func (x *InitWorkspaceRequest) GetPersistentVolumeClaim() bool {
	if x != nil {
		return x.PersistentVolumeClaim
	}
	return false
}

// WorkspaceMetadata is data associated with a workspace that's required for other parts of the system to function
type WorkspaceMetadata struct {
	state         protoimpl.MessageState  `json:"state,omitempty"`
	sizeCache     protoimpl.SizeCache     `json:"sizeCache,omitempty"`
	unknownFields protoimpl.UnknownFields `json:"unknownFields,omitempty"`

	// owner is the ID of the Gitpod user to whom we'll bill this workspace and who we consider responsible for its content
	Owner string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	// meta_id is the workspace ID of this currently running workspace instance on the "meta pool" side
	MetaId string `protobuf:"bytes,2,opt,name=meta_id,json=metaId,proto3" json:"metaId,omitempty"`
}

func (x *WorkspaceMetadata) Reset() {
	*x = WorkspaceMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_daemon_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkspaceMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkspaceMetadata) ProtoMessage() {}

func (x *WorkspaceMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_daemon_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkspaceMetadata.ProtoReflect.Descriptor instead.
func (*WorkspaceMetadata) Descriptor() ([]byte, []int) {
	return file_daemon_proto_rawDescGZIP(), []int{1}
}

func (x *WorkspaceMetadata) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *WorkspaceMetadata) GetMetaId() string {
	if x != nil {
		return x.MetaId
	}
	return ""
}

type InitWorkspaceResponse struct {
	state         protoimpl.MessageState  `json:"state,omitempty"`
	sizeCache     protoimpl.SizeCache     `json:"sizeCache,omitempty"`
	unknownFields protoimpl.UnknownFields `json:"unknownFields,omitempty"`
}

func (x *InitWorkspaceResponse) Reset() {
	*x = InitWorkspaceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_daemon_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitWorkspaceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitWorkspaceResponse) ProtoMessage() {}

func (x *InitWorkspaceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_daemon_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitWorkspaceResponse.ProtoReflect.Descriptor instead.
func (*InitWorkspaceResponse) Descriptor() ([]byte, []int) {
	return file_daemon_proto_rawDescGZIP(), []int{2}
}

// WaitForInitRequest waits for a workspace to be initialized
type WaitForInitRequest struct {
	state         protoimpl.MessageState  `json:"state,omitempty"`
	sizeCache     protoimpl.SizeCache     `json:"sizeCache,omitempty"`
	unknownFields protoimpl.UnknownFields `json:"unknownFields,omitempty"`

	// ID is a unique identifier of the workspace
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *WaitForInitRequest) Reset() {
	*x = WaitForInitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_daemon_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WaitForInitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WaitForInitRequest) ProtoMessage() {}

func (x *WaitForInitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_daemon_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WaitForInitRequest.ProtoReflect.Descriptor instead.
func (*WaitForInitRequest) Descriptor() ([]byte, []int) {
	return file_daemon_proto_rawDescGZIP(), []int{3}
}

func (x *WaitForInitRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type WaitForInitResponse struct {
	state         protoimpl.MessageState  `json:"state,omitempty"`
	sizeCache     protoimpl.SizeCache     `json:"sizeCache,omitempty"`
	unknownFields protoimpl.UnknownFields `json:"unknownFields,omitempty"`
}

func (x *WaitForInitResponse) Reset() {
	*x = WaitForInitResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_daemon_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WaitForInitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WaitForInitResponse) ProtoMessage() {}

func (x *WaitForInitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_daemon_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WaitForInitResponse.ProtoReflect.Descriptor instead.
func (*WaitForInitResponse) Descriptor() ([]byte, []int) {
	return file_daemon_proto_rawDescGZIP(), []int{4}
}

// TakeSnapshotRequest creates a backup/snapshot of a workspace
type TakeSnapshotRequest struct {
	state         protoimpl.MessageState  `json:"state,omitempty"`
	sizeCache     protoimpl.SizeCache     `json:"sizeCache,omitempty"`
	unknownFields protoimpl.UnknownFields `json:"unknownFields,omitempty"`

	// ID is the identifier of the workspace of which we want to create a snapshot of
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// return_immediately means we're not waiting until the snapshot is done but return immediately after starting it
	ReturnImmediately bool `protobuf:"varint,2,opt,name=return_immediately,json=returnImmediately,proto3" json:"returnImmediately,omitempty"`
}

func (x *TakeSnapshotRequest) Reset() {
	*x = TakeSnapshotRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_daemon_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TakeSnapshotRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TakeSnapshotRequest) ProtoMessage() {}

func (x *TakeSnapshotRequest) ProtoReflect() protoreflect.Message {
	mi := &file_daemon_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TakeSnapshotRequest.ProtoReflect.Descriptor instead.
func (*TakeSnapshotRequest) Descriptor() ([]byte, []int) {
	return file_daemon_proto_rawDescGZIP(), []int{5}
}

func (x *TakeSnapshotRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TakeSnapshotRequest) GetReturnImmediately() bool {
	if x != nil {
		return x.ReturnImmediately
	}
	return false
}

type TakeSnapshotResponse struct {
	state         protoimpl.MessageState  `json:"state,omitempty"`
	sizeCache     protoimpl.SizeCache     `json:"sizeCache,omitempty"`
	unknownFields protoimpl.UnknownFields `json:"unknownFields,omitempty"`

	// url is the name of the resulting snapshot
	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *TakeSnapshotResponse) Reset() {
	*x = TakeSnapshotResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_daemon_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TakeSnapshotResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TakeSnapshotResponse) ProtoMessage() {}

func (x *TakeSnapshotResponse) ProtoReflect() protoreflect.Message {
	mi := &file_daemon_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TakeSnapshotResponse.ProtoReflect.Descriptor instead.
func (*TakeSnapshotResponse) Descriptor() ([]byte, []int) {
	return file_daemon_proto_rawDescGZIP(), []int{6}
}

func (x *TakeSnapshotResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type DisposeWorkspaceRequest struct {
	state         protoimpl.MessageState  `json:"state,omitempty"`
	sizeCache     protoimpl.SizeCache     `json:"sizeCache,omitempty"`
	unknownFields protoimpl.UnknownFields `json:"unknownFields,omitempty"`

	// ID is a unique identifier of the workspace to dispose of
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Backup triggers a final backup prior to disposal
	Backup bool `protobuf:"varint,2,opt,name=backup,proto3" json:"backup,omitempty"`
	// backup_logs triggers the upload of terminal logs
	BackupLogs bool `protobuf:"varint,3,opt,name=backup_logs,json=backupLogs,proto3" json:"backupLogs,omitempty"`
}

func (x *DisposeWorkspaceRequest) Reset() {
	*x = DisposeWorkspaceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_daemon_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisposeWorkspaceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisposeWorkspaceRequest) ProtoMessage() {}

func (x *DisposeWorkspaceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_daemon_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisposeWorkspaceRequest.ProtoReflect.Descriptor instead.
func (*DisposeWorkspaceRequest) Descriptor() ([]byte, []int) {
	return file_daemon_proto_rawDescGZIP(), []int{7}
}

func (x *DisposeWorkspaceRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DisposeWorkspaceRequest) GetBackup() bool {
	if x != nil {
		return x.Backup
	}
	return false
}

func (x *DisposeWorkspaceRequest) GetBackupLogs() bool {
	if x != nil {
		return x.BackupLogs
	}
	return false
}

type DisposeWorkspaceResponse struct {
	state         protoimpl.MessageState  `json:"state,omitempty"`
	sizeCache     protoimpl.SizeCache     `json:"sizeCache,omitempty"`
	unknownFields protoimpl.UnknownFields `json:"unknownFields,omitempty"`

	// git_status is the current state of the Git repo in this workspace prior to disposal.
	// If the workspace has no Git repo at its checkout location, this is nil.
	GitStatus *api.GitStatus `protobuf:"bytes,1,opt,name=git_status,json=gitStatus,proto3" json:"gitStatus,omitempty"`
}

func (x *DisposeWorkspaceResponse) Reset() {
	*x = DisposeWorkspaceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_daemon_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisposeWorkspaceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisposeWorkspaceResponse) ProtoMessage() {}

func (x *DisposeWorkspaceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_daemon_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisposeWorkspaceResponse.ProtoReflect.Descriptor instead.
func (*DisposeWorkspaceResponse) Descriptor() ([]byte, []int) {
	return file_daemon_proto_rawDescGZIP(), []int{8}
}

func (x *DisposeWorkspaceResponse) GetGitStatus() *api.GitStatus {
	if x != nil {
		return x.GitStatus
	}
	return nil
}

// BackupWorkspaceRequest creates a backup of a workspace
// TODO(rl): do we need an optional bool 'backup_logs' to capture the logs as well?
type BackupWorkspaceRequest struct {
	state         protoimpl.MessageState  `json:"state,omitempty"`
	sizeCache     protoimpl.SizeCache     `json:"sizeCache,omitempty"`
	unknownFields protoimpl.UnknownFields `json:"unknownFields,omitempty"`

	// ID is the identifier of the workspace of which we want to create a backup of
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *BackupWorkspaceRequest) Reset() {
	*x = BackupWorkspaceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_daemon_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BackupWorkspaceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BackupWorkspaceRequest) ProtoMessage() {}

func (x *BackupWorkspaceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_daemon_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BackupWorkspaceRequest.ProtoReflect.Descriptor instead.
func (*BackupWorkspaceRequest) Descriptor() ([]byte, []int) {
	return file_daemon_proto_rawDescGZIP(), []int{9}
}

func (x *BackupWorkspaceRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type BackupWorkspaceResponse struct {
	state         protoimpl.MessageState  `json:"state,omitempty"`
	sizeCache     protoimpl.SizeCache     `json:"sizeCache,omitempty"`
	unknownFields protoimpl.UnknownFields `json:"unknownFields,omitempty"`

	// url is the name of the resulting backup
	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *BackupWorkspaceResponse) Reset() {
	*x = BackupWorkspaceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_daemon_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BackupWorkspaceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BackupWorkspaceResponse) ProtoMessage() {}

func (x *BackupWorkspaceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_daemon_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BackupWorkspaceResponse.ProtoReflect.Descriptor instead.
func (*BackupWorkspaceResponse) Descriptor() ([]byte, []int) {
	return file_daemon_proto_rawDescGZIP(), []int{10}
}

func (x *BackupWorkspaceResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

var File_daemon_proto protoreflect.FileDescriptor

var file_daemon_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08,
	0x77, 0x73, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x1a, 0x25, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6e,
	0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xac, 0x03, 0x0a, 0x14, 0x49, 0x6e, 0x69, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x37, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x77, 0x73, 0x64,
	0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x46, 0x0a, 0x0b, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x72, 0x52, 0x0b, 0x69, 0x6e,
	0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x72, 0x12, 0x32, 0x0a, 0x15, 0x66, 0x75, 0x6c,
	0x6c, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x62, 0x61, 0x63, 0x6b,
	0x75, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x66, 0x75, 0x6c, 0x6c, 0x57, 0x6f,
	0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x12, 0x29, 0x0a,
	0x10, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x17, 0x72, 0x65, 0x6d, 0x6f,
	0x74, 0x65, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x64, 0x69, 0x73, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x15, 0x72, 0x65, 0x6d, 0x6f, 0x74,
	0x65, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x12, 0x2e, 0x0a, 0x13, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x71, 0x75, 0x6f, 0x74,
	0x61, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x42, 0x79, 0x74, 0x65, 0x73,
	0x12, 0x36, 0x0a, 0x17, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x76,
	0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x5f, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x15, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x74, 0x56, 0x6f, 0x6c,
	0x75, 0x6d, 0x65, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x4a, 0x04, 0x08, 0x06, 0x10, 0x07, 0x22, 0x42,
	0x0a, 0x11, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x6d, 0x65, 0x74,
	0x61, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x61,
	0x49, 0x64, 0x22, 0x17, 0x0a, 0x15, 0x49, 0x6e, 0x69, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x24, 0x0a, 0x12, 0x57,
	0x61, 0x69, 0x74, 0x46, 0x6f, 0x72, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x15, 0x0a, 0x13, 0x57, 0x61, 0x69, 0x74, 0x46, 0x6f, 0x72, 0x49, 0x6e, 0x69, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x54, 0x0a, 0x13, 0x54, 0x61, 0x6b, 0x65,
	0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x2d, 0x0a, 0x12, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x5f, 0x69, 0x6d, 0x6d, 0x65, 0x64, 0x69,
	0x61, 0x74, 0x65, 0x6c, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x72, 0x65, 0x74,
	0x75, 0x72, 0x6e, 0x49, 0x6d, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x22, 0x28,
	0x0a, 0x14, 0x54, 0x61, 0x6b, 0x65, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x62, 0x0a, 0x17, 0x44, 0x69, 0x73, 0x70,
	0x6f, 0x73, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x06, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x12, 0x1f, 0x0a, 0x0b, 0x62,
	0x61, 0x63, 0x6b, 0x75, 0x70, 0x5f, 0x6c, 0x6f, 0x67, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0a, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x4c, 0x6f, 0x67, 0x73, 0x22, 0x54, 0x0a, 0x18,
	0x44, 0x69, 0x73, 0x70, 0x6f, 0x73, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x0a, 0x67, 0x69, 0x74, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x69,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x09, 0x67, 0x69, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x22, 0x28, 0x0a, 0x16, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x57, 0x6f, 0x72, 0x6b,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2b, 0x0a, 0x17,
	0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x2a, 0x51, 0x0a, 0x15, 0x57, 0x6f, 0x72,
	0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a,
	0x53, 0x45, 0x54, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x55, 0x50, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09,
	0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x57,
	0x52, 0x41, 0x50, 0x50, 0x49, 0x4e, 0x47, 0x5f, 0x55, 0x50, 0x10, 0x03, 0x32, 0xc3, 0x03, 0x0a,
	0x17, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x52, 0x0a, 0x0d, 0x49, 0x6e, 0x69, 0x74,
	0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x1e, 0x2e, 0x77, 0x73, 0x64, 0x61,
	0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x77, 0x73, 0x64, 0x61,
	0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x0b,
	0x57, 0x61, 0x69, 0x74, 0x46, 0x6f, 0x72, 0x49, 0x6e, 0x69, 0x74, 0x12, 0x1c, 0x2e, 0x77, 0x73,
	0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x57, 0x61, 0x69, 0x74, 0x46, 0x6f, 0x72, 0x49, 0x6e,
	0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x77, 0x73, 0x64, 0x61,
	0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x57, 0x61, 0x69, 0x74, 0x46, 0x6f, 0x72, 0x49, 0x6e, 0x69, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x0c, 0x54, 0x61,
	0x6b, 0x65, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x12, 0x1d, 0x2e, 0x77, 0x73, 0x64,
	0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x54, 0x61, 0x6b, 0x65, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68,
	0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x77, 0x73, 0x64, 0x61,
	0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x54, 0x61, 0x6b, 0x65, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5b, 0x0a, 0x10, 0x44,
	0x69, 0x73, 0x70, 0x6f, 0x73, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12,
	0x21, 0x2e, 0x77, 0x73, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x44, 0x69, 0x73, 0x70, 0x6f,
	0x73, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x22, 0x2e, 0x77, 0x73, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x44, 0x69,
	0x73, 0x70, 0x6f, 0x73, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x58, 0x0a, 0x0f, 0x42, 0x61, 0x63, 0x6b,
	0x75, 0x70, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x20, 0x2e, 0x77, 0x73,
	0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x57, 0x6f, 0x72,
	0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e,
	0x77, 0x73, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x57,
	0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x67, 0x69, 0x74, 0x70, 0x6f, 0x64, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x69, 0x74, 0x70, 0x6f,
	0x64, 0x2f, 0x77, 0x73, 0x2d, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_daemon_proto_rawDescOnce sync.Once
	file_daemon_proto_rawDescData = file_daemon_proto_rawDesc
)

func file_daemon_proto_rawDescGZIP() []byte {
	file_daemon_proto_rawDescOnce.Do(func() {
		file_daemon_proto_rawDescData = protoimpl.X.CompressGZIP(file_daemon_proto_rawDescData)
	})
	return file_daemon_proto_rawDescData
}

var file_daemon_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_daemon_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_daemon_proto_goTypes = []interface{}{
	(WorkspaceContentState)(0),       // 0: wsdaemon.WorkspaceContentState
	(*InitWorkspaceRequest)(nil),     // 1: wsdaemon.InitWorkspaceRequest
	(*WorkspaceMetadata)(nil),        // 2: wsdaemon.WorkspaceMetadata
	(*InitWorkspaceResponse)(nil),    // 3: wsdaemon.InitWorkspaceResponse
	(*WaitForInitRequest)(nil),       // 4: wsdaemon.WaitForInitRequest
	(*WaitForInitResponse)(nil),      // 5: wsdaemon.WaitForInitResponse
	(*TakeSnapshotRequest)(nil),      // 6: wsdaemon.TakeSnapshotRequest
	(*TakeSnapshotResponse)(nil),     // 7: wsdaemon.TakeSnapshotResponse
	(*DisposeWorkspaceRequest)(nil),  // 8: wsdaemon.DisposeWorkspaceRequest
	(*DisposeWorkspaceResponse)(nil), // 9: wsdaemon.DisposeWorkspaceResponse
	(*BackupWorkspaceRequest)(nil),   // 10: wsdaemon.BackupWorkspaceRequest
	(*BackupWorkspaceResponse)(nil),  // 11: wsdaemon.BackupWorkspaceResponse
	(*api.WorkspaceInitializer)(nil), // 12: contentservice.WorkspaceInitializer
	(*api.GitStatus)(nil),            // 13: contentservice.GitStatus
}
var file_daemon_proto_depIdxs = []int32{
	2,  // 0: wsdaemon.InitWorkspaceRequest.metadata:type_name -> wsdaemon.WorkspaceMetadata
	12, // 1: wsdaemon.InitWorkspaceRequest.initializer:type_name -> contentservice.WorkspaceInitializer
	13, // 2: wsdaemon.DisposeWorkspaceResponse.git_status:type_name -> contentservice.GitStatus
	1,  // 3: wsdaemon.WorkspaceContentService.InitWorkspace:input_type -> wsdaemon.InitWorkspaceRequest
	4,  // 4: wsdaemon.WorkspaceContentService.WaitForInit:input_type -> wsdaemon.WaitForInitRequest
	6,  // 5: wsdaemon.WorkspaceContentService.TakeSnapshot:input_type -> wsdaemon.TakeSnapshotRequest
	8,  // 6: wsdaemon.WorkspaceContentService.DisposeWorkspace:input_type -> wsdaemon.DisposeWorkspaceRequest
	10, // 7: wsdaemon.WorkspaceContentService.BackupWorkspace:input_type -> wsdaemon.BackupWorkspaceRequest
	3,  // 8: wsdaemon.WorkspaceContentService.InitWorkspace:output_type -> wsdaemon.InitWorkspaceResponse
	5,  // 9: wsdaemon.WorkspaceContentService.WaitForInit:output_type -> wsdaemon.WaitForInitResponse
	7,  // 10: wsdaemon.WorkspaceContentService.TakeSnapshot:output_type -> wsdaemon.TakeSnapshotResponse
	9,  // 11: wsdaemon.WorkspaceContentService.DisposeWorkspace:output_type -> wsdaemon.DisposeWorkspaceResponse
	11, // 12: wsdaemon.WorkspaceContentService.BackupWorkspace:output_type -> wsdaemon.BackupWorkspaceResponse
	8,  // [8:13] is the sub-list for method output_type
	3,  // [3:8] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_daemon_proto_init() }
func file_daemon_proto_init() {
	if File_daemon_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_daemon_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitWorkspaceRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_daemon_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkspaceMetadata); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_daemon_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitWorkspaceResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_daemon_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WaitForInitRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_daemon_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WaitForInitResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_daemon_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TakeSnapshotRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_daemon_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TakeSnapshotResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_daemon_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DisposeWorkspaceRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_daemon_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DisposeWorkspaceResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_daemon_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BackupWorkspaceRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_daemon_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BackupWorkspaceResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_daemon_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_daemon_proto_goTypes,
		DependencyIndexes: file_daemon_proto_depIdxs,
		EnumInfos:         file_daemon_proto_enumTypes,
		MessageInfos:      file_daemon_proto_msgTypes,
	}.Build()
	File_daemon_proto = out.File
	file_daemon_proto_rawDesc = nil
	file_daemon_proto_goTypes = nil
	file_daemon_proto_depIdxs = nil
}
