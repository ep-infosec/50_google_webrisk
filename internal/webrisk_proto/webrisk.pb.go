// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: webrisk.proto

/*
Package google_cloud_webrisk_v1 is a generated protocol buffer package.

It is generated from these files:
	webrisk.proto

It has these top-level messages:
	ComputeThreatListDiffRequest
	ComputeThreatListDiffResponse
	SearchUrisRequest
	SearchUrisResponse
	SearchHashesRequest
	SearchHashesResponse
	ThreatEntryAdditions
	ThreatEntryRemovals
	RawIndices
	RawHashes
	RiceDeltaEncoding
*/
package google_cloud_webrisk_v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The type of threat. This maps dirrectly to the threat list a threat may
// belong to.
type ThreatType int32

const (
	// Unknown.
	ThreatType_THREAT_TYPE_UNSPECIFIED ThreatType = 0
	// Malware targeting any platform.
	ThreatType_MALWARE ThreatType = 1
	// Social engineering targeting any platform.
	ThreatType_SOCIAL_ENGINEERING ThreatType = 2
	// Unwanted software targeting any platform.
	ThreatType_UNWANTED_SOFTWARE ThreatType = 3
)

var ThreatType_name = map[int32]string{
	0: "THREAT_TYPE_UNSPECIFIED",
	1: "MALWARE",
	2: "SOCIAL_ENGINEERING",
	3: "UNWANTED_SOFTWARE",
}
var ThreatType_value = map[string]int32{
	"THREAT_TYPE_UNSPECIFIED": 0,
	"MALWARE":                 1,
	"SOCIAL_ENGINEERING":      2,
	"UNWANTED_SOFTWARE":       3,
}

func (x ThreatType) String() string {
	return proto.EnumName(ThreatType_name, int32(x))
}
func (ThreatType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// The ways in which threat entry sets can be compressed.
type CompressionType int32

const (
	// Unknown.
	CompressionType_COMPRESSION_TYPE_UNSPECIFIED CompressionType = 0
	// Raw, uncompressed data.
	CompressionType_RAW CompressionType = 1
	// Rice-Golomb encoded data.
	CompressionType_RICE CompressionType = 2
)

var CompressionType_name = map[int32]string{
	0: "COMPRESSION_TYPE_UNSPECIFIED",
	1: "RAW",
	2: "RICE",
}
var CompressionType_value = map[string]int32{
	"COMPRESSION_TYPE_UNSPECIFIED": 0,
	"RAW":                          1,
	"RICE":                         2,
}

func (x CompressionType) String() string {
	return proto.EnumName(CompressionType_name, int32(x))
}
func (CompressionType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// The type of response sent to the client.
type ComputeThreatListDiffResponse_ResponseType int32

const (
	// Unknown.
	ComputeThreatListDiffResponse_RESPONSE_TYPE_UNSPECIFIED ComputeThreatListDiffResponse_ResponseType = 0
	// Partial updates are applied to the client's existing local database.
	ComputeThreatListDiffResponse_DIFF ComputeThreatListDiffResponse_ResponseType = 1
	// Full updates resets the client's entire local database. This means
	// that either the client had no state, was seriously out-of-date,
	// or the client is believed to be corrupt.
	ComputeThreatListDiffResponse_RESET ComputeThreatListDiffResponse_ResponseType = 2
)

var ComputeThreatListDiffResponse_ResponseType_name = map[int32]string{
	0: "RESPONSE_TYPE_UNSPECIFIED",
	1: "DIFF",
	2: "RESET",
}
var ComputeThreatListDiffResponse_ResponseType_value = map[string]int32{
	"RESPONSE_TYPE_UNSPECIFIED": 0,
	"DIFF":                      1,
	"RESET":                     2,
}

func (x ComputeThreatListDiffResponse_ResponseType) String() string {
	return proto.EnumName(ComputeThreatListDiffResponse_ResponseType_name, int32(x))
}
func (ComputeThreatListDiffResponse_ResponseType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{1, 0}
}

// Describes an API diff request.
type ComputeThreatListDiffRequest struct {
	// Required. The ThreatList to update.
	ThreatType ThreatType `protobuf:"varint,1,opt,name=threat_type,json=threatType,enum=google.cloud.webrisk.v1.ThreatType" json:"threat_type,omitempty"`
	// The current version token of the client for the requested list (the
	// client version that was received from the last successful diff).
	VersionToken []byte `protobuf:"bytes,2,opt,name=version_token,json=versionToken,proto3" json:"version_token,omitempty"`
	// The constraints associated with this request.
	Constraints *ComputeThreatListDiffRequest_Constraints `protobuf:"bytes,3,opt,name=constraints" json:"constraints,omitempty"`
}

func (m *ComputeThreatListDiffRequest) Reset()                    { *m = ComputeThreatListDiffRequest{} }
func (m *ComputeThreatListDiffRequest) String() string            { return proto.CompactTextString(m) }
func (*ComputeThreatListDiffRequest) ProtoMessage()               {}
func (*ComputeThreatListDiffRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ComputeThreatListDiffRequest) GetThreatType() ThreatType {
	if m != nil {
		return m.ThreatType
	}
	return ThreatType_THREAT_TYPE_UNSPECIFIED
}

func (m *ComputeThreatListDiffRequest) GetVersionToken() []byte {
	if m != nil {
		return m.VersionToken
	}
	return nil
}

func (m *ComputeThreatListDiffRequest) GetConstraints() *ComputeThreatListDiffRequest_Constraints {
	if m != nil {
		return m.Constraints
	}
	return nil
}

// The constraints for this diff.
type ComputeThreatListDiffRequest_Constraints struct {
	// The maximum size in number of entries. The diff will not contain more
	// entries than this value.  This should be a power of 2 between 2**10 and
	// 2**20.  If zero, no diff size limit is set.
	MaxDiffEntries int32 `protobuf:"varint,1,opt,name=max_diff_entries,json=maxDiffEntries" json:"max_diff_entries,omitempty"`
	// Sets the maximum number of entries that the client is willing to have
	// in the local database. This should be a power of 2 between 2**10 and
	// 2**20. If zero, no database size limit is set.
	MaxDatabaseEntries int32 `protobuf:"varint,2,opt,name=max_database_entries,json=maxDatabaseEntries" json:"max_database_entries,omitempty"`
	// The compression types supported by the client.
	SupportedCompressions []CompressionType `protobuf:"varint,3,rep,packed,name=supported_compressions,json=supportedCompressions,enum=google.cloud.webrisk.v1.CompressionType" json:"supported_compressions,omitempty"`
}

func (m *ComputeThreatListDiffRequest_Constraints) Reset() {
	*m = ComputeThreatListDiffRequest_Constraints{}
}
func (m *ComputeThreatListDiffRequest_Constraints) String() string { return proto.CompactTextString(m) }
func (*ComputeThreatListDiffRequest_Constraints) ProtoMessage()    {}
func (*ComputeThreatListDiffRequest_Constraints) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 0}
}

func (m *ComputeThreatListDiffRequest_Constraints) GetMaxDiffEntries() int32 {
	if m != nil {
		return m.MaxDiffEntries
	}
	return 0
}

func (m *ComputeThreatListDiffRequest_Constraints) GetMaxDatabaseEntries() int32 {
	if m != nil {
		return m.MaxDatabaseEntries
	}
	return 0
}

func (m *ComputeThreatListDiffRequest_Constraints) GetSupportedCompressions() []CompressionType {
	if m != nil {
		return m.SupportedCompressions
	}
	return nil
}

type ComputeThreatListDiffResponse struct {
	// The type of response. This may indicate that an action is required by the
	// client when the response is received.
	ResponseType ComputeThreatListDiffResponse_ResponseType `protobuf:"varint,4,opt,name=response_type,json=responseType,enum=google.cloud.webrisk.v1.ComputeThreatListDiffResponse_ResponseType" json:"response_type,omitempty"`
	// A set of entries to add to a local threat type's list.
	Additions *ThreatEntryAdditions `protobuf:"bytes,5,opt,name=additions" json:"additions,omitempty"`
	// A set of entries to remove from a local threat type's list.
	// This field may be empty.
	Removals *ThreatEntryRemovals `protobuf:"bytes,6,opt,name=removals" json:"removals,omitempty"`
	// The new opaque client version token.
	NewVersionToken []byte `protobuf:"bytes,7,opt,name=new_version_token,json=newVersionToken,proto3" json:"new_version_token,omitempty"`
	// The expected SHA256 hash of the client state; that is, of the sorted list
	// of all hashes present in the database after applying the provided diff.
	// If the client state doesn't match the expected state, the client must
	// disregard this diff and retry later.
	Checksum *ComputeThreatListDiffResponse_Checksum `protobuf:"bytes,8,opt,name=checksum" json:"checksum,omitempty"`
	// The soonest the client should wait before issuing any diff
	// request. Querying sooner is unlikely to produce a meaningful diff.
	// Waiting longer is acceptable considering the use case.
	// If this field is not set clients may update as soon as they want.
	RecommendedNextDiff *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=recommended_next_diff,json=recommendedNextDiff" json:"recommended_next_diff,omitempty"`
}

func (m *ComputeThreatListDiffResponse) Reset()                    { *m = ComputeThreatListDiffResponse{} }
func (m *ComputeThreatListDiffResponse) String() string            { return proto.CompactTextString(m) }
func (*ComputeThreatListDiffResponse) ProtoMessage()               {}
func (*ComputeThreatListDiffResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ComputeThreatListDiffResponse) GetResponseType() ComputeThreatListDiffResponse_ResponseType {
	if m != nil {
		return m.ResponseType
	}
	return ComputeThreatListDiffResponse_RESPONSE_TYPE_UNSPECIFIED
}

func (m *ComputeThreatListDiffResponse) GetAdditions() *ThreatEntryAdditions {
	if m != nil {
		return m.Additions
	}
	return nil
}

func (m *ComputeThreatListDiffResponse) GetRemovals() *ThreatEntryRemovals {
	if m != nil {
		return m.Removals
	}
	return nil
}

func (m *ComputeThreatListDiffResponse) GetNewVersionToken() []byte {
	if m != nil {
		return m.NewVersionToken
	}
	return nil
}

func (m *ComputeThreatListDiffResponse) GetChecksum() *ComputeThreatListDiffResponse_Checksum {
	if m != nil {
		return m.Checksum
	}
	return nil
}

func (m *ComputeThreatListDiffResponse) GetRecommendedNextDiff() *google_protobuf.Timestamp {
	if m != nil {
		return m.RecommendedNextDiff
	}
	return nil
}

// The expected state of a client's local database.
type ComputeThreatListDiffResponse_Checksum struct {
	// The SHA256 hash of the client state; that is, of the sorted list of all
	// hashes present in the database.
	Sha256 []byte `protobuf:"bytes,1,opt,name=sha256,proto3" json:"sha256,omitempty"`
}

func (m *ComputeThreatListDiffResponse_Checksum) Reset() {
	*m = ComputeThreatListDiffResponse_Checksum{}
}
func (m *ComputeThreatListDiffResponse_Checksum) String() string { return proto.CompactTextString(m) }
func (*ComputeThreatListDiffResponse_Checksum) ProtoMessage()    {}
func (*ComputeThreatListDiffResponse_Checksum) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{1, 0}
}

func (m *ComputeThreatListDiffResponse_Checksum) GetSha256() []byte {
	if m != nil {
		return m.Sha256
	}
	return nil
}

// Request to check URI entries against threatLists.
type SearchUrisRequest struct {
	// The URI to be checked for matches.
	Uri string `protobuf:"bytes,1,opt,name=uri" json:"uri,omitempty"`
	// Required. The ThreatLists to search in.
	ThreatTypes []ThreatType `protobuf:"varint,2,rep,packed,name=threat_types,json=threatTypes,enum=google.cloud.webrisk.v1.ThreatType" json:"threat_types,omitempty"`
}

func (m *SearchUrisRequest) Reset()                    { *m = SearchUrisRequest{} }
func (m *SearchUrisRequest) String() string            { return proto.CompactTextString(m) }
func (*SearchUrisRequest) ProtoMessage()               {}
func (*SearchUrisRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *SearchUrisRequest) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

func (m *SearchUrisRequest) GetThreatTypes() []ThreatType {
	if m != nil {
		return m.ThreatTypes
	}
	return nil
}

type SearchUrisResponse struct {
	// The threat list matches. This may be empty if the URI is on no list.
	Threat *SearchUrisResponse_ThreatUri `protobuf:"bytes,1,opt,name=threat" json:"threat,omitempty"`
}

func (m *SearchUrisResponse) Reset()                    { *m = SearchUrisResponse{} }
func (m *SearchUrisResponse) String() string            { return proto.CompactTextString(m) }
func (*SearchUrisResponse) ProtoMessage()               {}
func (*SearchUrisResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *SearchUrisResponse) GetThreat() *SearchUrisResponse_ThreatUri {
	if m != nil {
		return m.Threat
	}
	return nil
}

// Contains threat information on a matching uri.
type SearchUrisResponse_ThreatUri struct {
	// The ThreatList this threat belongs to.
	ThreatTypes []ThreatType `protobuf:"varint,1,rep,packed,name=threat_types,json=threatTypes,enum=google.cloud.webrisk.v1.ThreatType" json:"threat_types,omitempty"`
	// The cache lifetime for the returned match. Clients must not cache this
	// response past this timestamp to avoid false positives.
	ExpireTime *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=expire_time,json=expireTime" json:"expire_time,omitempty"`
}

func (m *SearchUrisResponse_ThreatUri) Reset()                    { *m = SearchUrisResponse_ThreatUri{} }
func (m *SearchUrisResponse_ThreatUri) String() string            { return proto.CompactTextString(m) }
func (*SearchUrisResponse_ThreatUri) ProtoMessage()               {}
func (*SearchUrisResponse_ThreatUri) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 0} }

func (m *SearchUrisResponse_ThreatUri) GetThreatTypes() []ThreatType {
	if m != nil {
		return m.ThreatTypes
	}
	return nil
}

func (m *SearchUrisResponse_ThreatUri) GetExpireTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.ExpireTime
	}
	return nil
}

// Request to return full hashes matched by the provided hash prefixes.
type SearchHashesRequest struct {
	// A hash prefix, consisting of the most significant 4-32 bytes of a SHA256
	// hash. For JSON requests, this field is base64-encoded.
	HashPrefix []byte `protobuf:"bytes,1,opt,name=hash_prefix,json=hashPrefix,proto3" json:"hash_prefix,omitempty"`
	// Required. The ThreatLists to search in.
	ThreatTypes []ThreatType `protobuf:"varint,2,rep,packed,name=threat_types,json=threatTypes,enum=google.cloud.webrisk.v1.ThreatType" json:"threat_types,omitempty"`
}

func (m *SearchHashesRequest) Reset()                    { *m = SearchHashesRequest{} }
func (m *SearchHashesRequest) String() string            { return proto.CompactTextString(m) }
func (*SearchHashesRequest) ProtoMessage()               {}
func (*SearchHashesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *SearchHashesRequest) GetHashPrefix() []byte {
	if m != nil {
		return m.HashPrefix
	}
	return nil
}

func (m *SearchHashesRequest) GetThreatTypes() []ThreatType {
	if m != nil {
		return m.ThreatTypes
	}
	return nil
}

type SearchHashesResponse struct {
	// The full hashes that matched the requested prefixes.
	// The hash will be populated in the key.
	Threats []*SearchHashesResponse_ThreatHash `protobuf:"bytes,1,rep,name=threats" json:"threats,omitempty"`
	// For requested entities that did not match the threat list, how long to
	// cache the response until.
	NegativeExpireTime *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=negative_expire_time,json=negativeExpireTime" json:"negative_expire_time,omitempty"`
}

func (m *SearchHashesResponse) Reset()                    { *m = SearchHashesResponse{} }
func (m *SearchHashesResponse) String() string            { return proto.CompactTextString(m) }
func (*SearchHashesResponse) ProtoMessage()               {}
func (*SearchHashesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *SearchHashesResponse) GetThreats() []*SearchHashesResponse_ThreatHash {
	if m != nil {
		return m.Threats
	}
	return nil
}

func (m *SearchHashesResponse) GetNegativeExpireTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.NegativeExpireTime
	}
	return nil
}

// Contains threat information on a matching hash.
type SearchHashesResponse_ThreatHash struct {
	// The ThreatList this threat belongs to.
	// This must contain at least one entry.
	ThreatTypes []ThreatType `protobuf:"varint,1,rep,packed,name=threat_types,json=threatTypes,enum=google.cloud.webrisk.v1.ThreatType" json:"threat_types,omitempty"`
	// A 32 byte SHA256 hash. This field is in binary format. For JSON
	// requests, hashes are base64-encoded.
	Hash []byte `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	// The cache lifetime for the returned match. Clients must not cache this
	// response past this timestamp to avoid false positives.
	ExpireTime *google_protobuf.Timestamp `protobuf:"bytes,3,opt,name=expire_time,json=expireTime" json:"expire_time,omitempty"`
}

func (m *SearchHashesResponse_ThreatHash) Reset()         { *m = SearchHashesResponse_ThreatHash{} }
func (m *SearchHashesResponse_ThreatHash) String() string { return proto.CompactTextString(m) }
func (*SearchHashesResponse_ThreatHash) ProtoMessage()    {}
func (*SearchHashesResponse_ThreatHash) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{5, 0}
}

func (m *SearchHashesResponse_ThreatHash) GetThreatTypes() []ThreatType {
	if m != nil {
		return m.ThreatTypes
	}
	return nil
}

func (m *SearchHashesResponse_ThreatHash) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *SearchHashesResponse_ThreatHash) GetExpireTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.ExpireTime
	}
	return nil
}

// Contains the set of entries to add to a local database.
// May contain a combination of compressed and raw data in a single response.
type ThreatEntryAdditions struct {
	// The raw SHA256-formatted entries.
	// Repeated to allow returning sets of hashes with different prefix sizes.
	RawHashes []*RawHashes `protobuf:"bytes,1,rep,name=raw_hashes,json=rawHashes" json:"raw_hashes,omitempty"`
	// The encoded 4-byte prefixes of SHA256-formatted entries, using a
	// Golomb-Rice encoding. The hashes are converted to uint32, sorted in
	// ascending order, then delta encoded and stored as encoded_data.
	RiceHashes *RiceDeltaEncoding `protobuf:"bytes,2,opt,name=rice_hashes,json=riceHashes" json:"rice_hashes,omitempty"`
}

func (m *ThreatEntryAdditions) Reset()                    { *m = ThreatEntryAdditions{} }
func (m *ThreatEntryAdditions) String() string            { return proto.CompactTextString(m) }
func (*ThreatEntryAdditions) ProtoMessage()               {}
func (*ThreatEntryAdditions) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ThreatEntryAdditions) GetRawHashes() []*RawHashes {
	if m != nil {
		return m.RawHashes
	}
	return nil
}

func (m *ThreatEntryAdditions) GetRiceHashes() *RiceDeltaEncoding {
	if m != nil {
		return m.RiceHashes
	}
	return nil
}

// Contains the set of entries to remove from a local database.
type ThreatEntryRemovals struct {
	// The raw removal indices for a local list.
	RawIndices *RawIndices `protobuf:"bytes,1,opt,name=raw_indices,json=rawIndices" json:"raw_indices,omitempty"`
	// The encoded local, lexicographically-sorted list indices, using a
	// Golomb-Rice encoding. Used for sending compressed removal indices. The
	// removal indices (uint32) are sorted in ascending order, then delta encoded
	// and stored as encoded_data.
	RiceIndices *RiceDeltaEncoding `protobuf:"bytes,2,opt,name=rice_indices,json=riceIndices" json:"rice_indices,omitempty"`
}

func (m *ThreatEntryRemovals) Reset()                    { *m = ThreatEntryRemovals{} }
func (m *ThreatEntryRemovals) String() string            { return proto.CompactTextString(m) }
func (*ThreatEntryRemovals) ProtoMessage()               {}
func (*ThreatEntryRemovals) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ThreatEntryRemovals) GetRawIndices() *RawIndices {
	if m != nil {
		return m.RawIndices
	}
	return nil
}

func (m *ThreatEntryRemovals) GetRiceIndices() *RiceDeltaEncoding {
	if m != nil {
		return m.RiceIndices
	}
	return nil
}

// A set of raw indices to remove from a local list.
type RawIndices struct {
	// The indices to remove from a lexicographically-sorted local list.
	Indices []int32 `protobuf:"varint,1,rep,packed,name=indices" json:"indices,omitempty"`
}

func (m *RawIndices) Reset()                    { *m = RawIndices{} }
func (m *RawIndices) String() string            { return proto.CompactTextString(m) }
func (*RawIndices) ProtoMessage()               {}
func (*RawIndices) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *RawIndices) GetIndices() []int32 {
	if m != nil {
		return m.Indices
	}
	return nil
}

// The uncompressed threat entries in hash format.
// Hashes can be anywhere from 4 to 32 bytes in size. A large majority are 4
// bytes, but some hashes are lengthened if they collide with the hash of a
// popular URI.
//
// Used for sending ThreatEntryAdditons to clients that do not support
// compression, or when sending non-4-byte hashes to clients that do support
// compression.
type RawHashes struct {
	// The number of bytes for each prefix encoded below.  This field can be
	// anywhere from 4 (shortest prefix) to 32 (full SHA256 hash).
	PrefixSize int32 `protobuf:"varint,1,opt,name=prefix_size,json=prefixSize" json:"prefix_size,omitempty"`
	// The hashes, in binary format, concatenated into one long string. Hashes are
	// sorted in lexicographic order. For JSON API users, hashes are
	// base64-encoded.
	RawHashes []byte `protobuf:"bytes,2,opt,name=raw_hashes,json=rawHashes,proto3" json:"raw_hashes,omitempty"`
}

func (m *RawHashes) Reset()                    { *m = RawHashes{} }
func (m *RawHashes) String() string            { return proto.CompactTextString(m) }
func (*RawHashes) ProtoMessage()               {}
func (*RawHashes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *RawHashes) GetPrefixSize() int32 {
	if m != nil {
		return m.PrefixSize
	}
	return 0
}

func (m *RawHashes) GetRawHashes() []byte {
	if m != nil {
		return m.RawHashes
	}
	return nil
}

// The Rice-Golomb encoded data. Used for sending compressed 4-byte hashes or
// compressed removal indices.
type RiceDeltaEncoding struct {
	// The offset of the first entry in the encoded data, or, if only a single
	// integer was encoded, that single integer's value. If the field is empty or
	// missing, assume zero.
	FirstValue int64 `protobuf:"varint,1,opt,name=first_value,json=firstValue" json:"first_value,omitempty"`
	// The Golomb-Rice parameter, which is a number between 2 and 28. This field
	// is missing (that is, zero) if `num_entries` is zero.
	RiceParameter int32 `protobuf:"varint,2,opt,name=rice_parameter,json=riceParameter" json:"rice_parameter,omitempty"`
	// The number of entries that are delta encoded in the encoded data. If only a
	// single integer was encoded, this will be zero and the single value will be
	// stored in `first_value`.
	EntryCount int32 `protobuf:"varint,3,opt,name=entry_count,json=entryCount" json:"entry_count,omitempty"`
	// The encoded deltas that are encoded using the Golomb-Rice coder.
	EncodedData []byte `protobuf:"bytes,4,opt,name=encoded_data,json=encodedData,proto3" json:"encoded_data,omitempty"`
}

func (m *RiceDeltaEncoding) Reset()                    { *m = RiceDeltaEncoding{} }
func (m *RiceDeltaEncoding) String() string            { return proto.CompactTextString(m) }
func (*RiceDeltaEncoding) ProtoMessage()               {}
func (*RiceDeltaEncoding) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *RiceDeltaEncoding) GetFirstValue() int64 {
	if m != nil {
		return m.FirstValue
	}
	return 0
}

func (m *RiceDeltaEncoding) GetRiceParameter() int32 {
	if m != nil {
		return m.RiceParameter
	}
	return 0
}

func (m *RiceDeltaEncoding) GetEntryCount() int32 {
	if m != nil {
		return m.EntryCount
	}
	return 0
}

func (m *RiceDeltaEncoding) GetEncodedData() []byte {
	if m != nil {
		return m.EncodedData
	}
	return nil
}

func init() {
	proto.RegisterType((*ComputeThreatListDiffRequest)(nil), "google.cloud.webrisk.v1.ComputeThreatListDiffRequest")
	proto.RegisterType((*ComputeThreatListDiffRequest_Constraints)(nil), "google.cloud.webrisk.v1.ComputeThreatListDiffRequest.Constraints")
	proto.RegisterType((*ComputeThreatListDiffResponse)(nil), "google.cloud.webrisk.v1.ComputeThreatListDiffResponse")
	proto.RegisterType((*ComputeThreatListDiffResponse_Checksum)(nil), "google.cloud.webrisk.v1.ComputeThreatListDiffResponse.Checksum")
	proto.RegisterType((*SearchUrisRequest)(nil), "google.cloud.webrisk.v1.SearchUrisRequest")
	proto.RegisterType((*SearchUrisResponse)(nil), "google.cloud.webrisk.v1.SearchUrisResponse")
	proto.RegisterType((*SearchUrisResponse_ThreatUri)(nil), "google.cloud.webrisk.v1.SearchUrisResponse.ThreatUri")
	proto.RegisterType((*SearchHashesRequest)(nil), "google.cloud.webrisk.v1.SearchHashesRequest")
	proto.RegisterType((*SearchHashesResponse)(nil), "google.cloud.webrisk.v1.SearchHashesResponse")
	proto.RegisterType((*SearchHashesResponse_ThreatHash)(nil), "google.cloud.webrisk.v1.SearchHashesResponse.ThreatHash")
	proto.RegisterType((*ThreatEntryAdditions)(nil), "google.cloud.webrisk.v1.ThreatEntryAdditions")
	proto.RegisterType((*ThreatEntryRemovals)(nil), "google.cloud.webrisk.v1.ThreatEntryRemovals")
	proto.RegisterType((*RawIndices)(nil), "google.cloud.webrisk.v1.RawIndices")
	proto.RegisterType((*RawHashes)(nil), "google.cloud.webrisk.v1.RawHashes")
	proto.RegisterType((*RiceDeltaEncoding)(nil), "google.cloud.webrisk.v1.RiceDeltaEncoding")
	proto.RegisterEnum("google.cloud.webrisk.v1.ThreatType", ThreatType_name, ThreatType_value)
	proto.RegisterEnum("google.cloud.webrisk.v1.CompressionType", CompressionType_name, CompressionType_value)
	proto.RegisterEnum("google.cloud.webrisk.v1.ComputeThreatListDiffResponse_ResponseType", ComputeThreatListDiffResponse_ResponseType_name, ComputeThreatListDiffResponse_ResponseType_value)
}

func init() { proto.RegisterFile("webrisk.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1095 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0xcf, 0x6e, 0x1b, 0xb7,
	0x13, 0xfe, 0xad, 0xe4, 0x7f, 0x9a, 0x95, 0x1d, 0x99, 0x71, 0x12, 0xfd, 0xd4, 0x04, 0x71, 0x37,
	0x68, 0x21, 0x18, 0xad, 0xd2, 0xaa, 0x48, 0x51, 0xa0, 0x87, 0x42, 0x95, 0x56, 0x89, 0x60, 0x5b,
	0x16, 0x28, 0x39, 0x46, 0xd1, 0xc3, 0x82, 0x5e, 0x51, 0x16, 0x61, 0xef, 0x9f, 0x92, 0x94, 0x25,
	0xe7, 0xd0, 0x67, 0xe8, 0x1b, 0xf4, 0xd4, 0x1e, 0xfa, 0x0a, 0x7d, 0x83, 0x3e, 0x45, 0xfb, 0x12,
	0x3d, 0x17, 0xe4, 0x72, 0x25, 0x39, 0xb6, 0x1c, 0x3b, 0xc8, 0x8d, 0xfb, 0x71, 0xe6, 0x9b, 0xe1,
	0x37, 0xc3, 0x59, 0xc2, 0xfa, 0x98, 0x1e, 0x73, 0x26, 0x4e, 0x2b, 0x31, 0x8f, 0x64, 0x84, 0x1e,
	0x9d, 0x44, 0xd1, 0xc9, 0x19, 0xad, 0xf8, 0x67, 0xd1, 0xa8, 0x5f, 0x49, 0xf7, 0xce, 0xbf, 0x2c,
	0x3d, 0x4d, 0x36, 0x9e, 0x6b, 0xb3, 0xe3, 0xd1, 0xe0, 0xb9, 0x64, 0x01, 0x15, 0x92, 0x04, 0x71,
	0xe2, 0xe9, 0xfc, 0x99, 0x85, 0xc7, 0xf5, 0x28, 0x88, 0x47, 0x92, 0xf6, 0x86, 0x9c, 0x12, 0xb9,
	0xc7, 0x84, 0x6c, 0xb0, 0xc1, 0x00, 0xd3, 0x9f, 0x46, 0x54, 0x48, 0xd4, 0x00, 0x5b, 0xea, 0x0d,
	0x4f, 0x5e, 0xc4, 0xb4, 0x68, 0x6d, 0x5b, 0xe5, 0x8d, 0xea, 0xb3, 0xca, 0x82, 0x80, 0x95, 0x84,
	0xa4, 0x77, 0x11, 0x53, 0x0c, 0x72, 0xba, 0x46, 0xcf, 0x60, 0xfd, 0x9c, 0x72, 0xc1, 0xa2, 0xd0,
	0x93, 0xd1, 0x29, 0x0d, 0x8b, 0x99, 0x6d, 0xab, 0x9c, 0xc7, 0x79, 0x03, 0xf6, 0x14, 0x86, 0x7c,
	0xb0, 0xfd, 0x28, 0x14, 0x92, 0x13, 0x16, 0x4a, 0x51, 0xcc, 0x6e, 0x5b, 0x65, 0xbb, 0x5a, 0x5b,
	0x18, 0xea, 0xa6, 0xb4, 0x2b, 0xf5, 0x19, 0x11, 0x9e, 0x67, 0x2d, 0xfd, 0x65, 0x81, 0x3d, 0xb7,
	0x89, 0xca, 0x50, 0x08, 0xc8, 0xc4, 0xeb, 0xb3, 0xc1, 0xc0, 0xa3, 0xa1, 0xe4, 0x8c, 0x0a, 0x7d,
	0xc8, 0x65, 0xbc, 0x11, 0x90, 0x89, 0xa2, 0x74, 0x13, 0x14, 0x7d, 0x01, 0x5b, 0xda, 0x92, 0x48,
	0x72, 0x4c, 0x04, 0x9d, 0x5a, 0x67, 0xb4, 0x35, 0x52, 0xd6, 0x66, 0x2b, 0xf5, 0xf0, 0xe0, 0xa1,
	0x18, 0xc5, 0x71, 0xc4, 0x25, 0xed, 0x7b, 0x7e, 0x14, 0xc4, 0x9c, 0x0a, 0x75, 0x5c, 0x75, 0xb6,
	0x6c, 0x79, 0xa3, 0x5a, 0xbe, 0xf1, 0x6c, 0xc6, 0x58, 0x6b, 0xf9, 0x60, 0xca, 0x33, 0xb7, 0x23,
	0x9c, 0x7f, 0x96, 0xe0, 0xc9, 0x02, 0x19, 0x44, 0x1c, 0x85, 0x82, 0xa2, 0x21, 0xac, 0x73, 0xb3,
	0x4e, 0x0a, 0xb8, 0xa4, 0x0b, 0x58, 0xbf, 0xab, 0xaa, 0x09, 0x45, 0x25, 0x5d, 0xe8, 0xa4, 0xf2,
	0x7c, 0xee, 0x0b, 0xed, 0x42, 0x8e, 0xf4, 0xfb, 0x4c, 0xea, 0xf3, 0x2d, 0xeb, 0xda, 0x7d, 0xfe,
	0x8e, 0x36, 0x51, 0x3a, 0x5d, 0xd4, 0x52, 0x27, 0x3c, 0xf3, 0x47, 0xaf, 0x60, 0x8d, 0xd3, 0x20,
	0x3a, 0x27, 0x67, 0xa2, 0xb8, 0xa2, 0xb9, 0x3e, 0xbb, 0x0d, 0x17, 0x36, 0x3e, 0x78, 0xea, 0x8d,
	0x76, 0x60, 0x33, 0xa4, 0x63, 0xef, 0x72, 0xf7, 0xad, 0xea, 0xee, 0xbb, 0x17, 0xd2, 0xf1, 0xeb,
	0xf9, 0x06, 0xfc, 0x11, 0xd6, 0xfc, 0x21, 0xf5, 0x4f, 0xc5, 0x28, 0x28, 0xae, 0xe9, 0xa8, 0xdf,
	0xbd, 0xa7, 0x4e, 0x75, 0x43, 0x83, 0xa7, 0x84, 0xa8, 0x0d, 0x0f, 0x38, 0xf5, 0xa3, 0x20, 0xa0,
	0x61, 0x9f, 0xf6, 0xbd, 0x90, 0x4e, 0xa4, 0xee, 0x3a, 0xdd, 0x3f, 0x76, 0xb5, 0x94, 0x46, 0x4a,
	0xaf, 0x6a, 0xa5, 0x97, 0x5e, 0x55, 0x7c, 0x7f, 0xce, 0xb1, 0x4d, 0x27, 0x3a, 0x54, 0xc9, 0x81,
	0xb5, 0x34, 0x0a, 0x7a, 0x08, 0x2b, 0x62, 0x48, 0xaa, 0x2f, 0xbe, 0xd6, 0xad, 0x9b, 0xc7, 0xe6,
	0xcb, 0xf9, 0x1e, 0xf2, 0xf3, 0x15, 0x43, 0x4f, 0xe0, 0xff, 0xd8, 0xed, 0x76, 0x0e, 0xda, 0x5d,
	0xd7, 0xeb, 0xfd, 0xd0, 0x71, 0xbd, 0xc3, 0x76, 0xb7, 0xe3, 0xd6, 0x5b, 0xcd, 0x96, 0xdb, 0x28,
	0xfc, 0x0f, 0xad, 0xc1, 0x52, 0xa3, 0xd5, 0x6c, 0x16, 0x2c, 0x94, 0x83, 0x65, 0xec, 0x76, 0xdd,
	0x5e, 0x21, 0xe3, 0x04, 0xb0, 0xd9, 0xa5, 0x84, 0xfb, 0xc3, 0x43, 0xce, 0x44, 0x3a, 0x15, 0x0a,
	0x90, 0x1d, 0x71, 0xa6, 0xa3, 0xe5, 0xb0, 0x5a, 0xa2, 0x26, 0xe4, 0xe7, 0xe6, 0x84, 0xba, 0x15,
	0xd9, 0xdb, 0x0e, 0x0a, 0x7b, 0x36, 0x28, 0x84, 0xf3, 0xaf, 0x05, 0x68, 0x3e, 0x9e, 0xe9, 0xe3,
	0x7d, 0x58, 0x49, 0xac, 0x74, 0x4c, 0xbb, 0xfa, 0x62, 0x21, 0xf1, 0x55, 0x67, 0x13, 0xeb, 0x90,
	0x33, 0x6c, 0x48, 0x4a, 0xbf, 0x58, 0x90, 0x9b, 0xa2, 0x57, 0x72, 0xb7, 0xde, 0x2f, 0x77, 0xf4,
	0x2d, 0xd8, 0x74, 0x12, 0x33, 0x4e, 0x3d, 0x35, 0x66, 0x6f, 0x51, 0x58, 0x48, 0xcc, 0x15, 0xe0,
	0xfc, 0x0c, 0xf7, 0x93, 0xd4, 0x5f, 0x11, 0x31, 0xa4, 0x53, 0xa5, 0x9f, 0x82, 0x3d, 0x24, 0x62,
	0xe8, 0xc5, 0x9c, 0x0e, 0xd8, 0xc4, 0xd4, 0x17, 0x14, 0xd4, 0xd1, 0xc8, 0x07, 0x13, 0xfe, 0xef,
	0x0c, 0x6c, 0x5d, 0x4e, 0xc0, 0x48, 0x8f, 0x61, 0x35, 0xb1, 0x4b, 0x84, 0xb1, 0xab, 0xdf, 0xbc,
	0x43, 0xfb, 0xcb, 0xfe, 0x26, 0xa0, 0x02, 0x71, 0x4a, 0x84, 0xf6, 0x60, 0x2b, 0xa4, 0x27, 0x44,
	0xb2, 0x73, 0xea, 0xdd, 0x4d, 0x32, 0x94, 0xfa, 0xb9, 0x53, 0xe9, 0x4a, 0xbf, 0x59, 0x00, 0xb3,
	0x28, 0x1f, 0xac, 0x9c, 0x08, 0x96, 0x94, 0xce, 0xe6, 0x5f, 0xa5, 0xd7, 0x6f, 0x97, 0x38, 0x7b,
	0xa7, 0x12, 0xff, 0x6e, 0xc1, 0xd6, 0x75, 0x93, 0x0f, 0xd5, 0x00, 0x38, 0x19, 0x7b, 0x43, 0x2d,
	0x9c, 0x51, 0xd9, 0x59, 0x98, 0x2f, 0x26, 0x63, 0x23, 0x71, 0x8e, 0xa7, 0x4b, 0xb4, 0x0b, 0x36,
	0x67, 0x3e, 0x4d, 0x39, 0x12, 0x21, 0x77, 0x16, 0x73, 0x30, 0x9f, 0x36, 0xe8, 0x99, 0x24, 0x6e,
	0xe8, 0x47, 0x7d, 0x16, 0x9e, 0x60, 0x50, 0xee, 0x09, 0x99, 0xf3, 0x87, 0x05, 0xf7, 0xaf, 0x19,
	0xab, 0xea, 0x31, 0xa0, 0xf2, 0x64, 0x61, 0x9f, 0xf9, 0xe6, 0x3f, 0x69, 0xdf, 0x20, 0x2c, 0x26,
	0xe3, 0x56, 0x62, 0x8a, 0xd5, 0xf9, 0xcc, 0x1a, 0xed, 0x43, 0x5e, 0xa7, 0x9a, 0xd2, 0xdc, 0x3d,
	0x57, 0x7d, 0x54, 0x43, 0xe7, 0x7c, 0x0a, 0x30, 0x0b, 0x84, 0x8a, 0xb0, 0x3a, 0x4b, 0x2f, 0x5b,
	0x5e, 0xc6, 0xe9, 0xa7, 0xb3, 0x0b, 0xb9, 0xa9, 0x72, 0xea, 0x5a, 0x25, 0x37, 0xca, 0x13, 0xec,
	0x0d, 0x35, 0x7f, 0x7c, 0x48, 0xa0, 0x2e, 0x7b, 0xa3, 0x46, 0xe5, 0x7c, 0x49, 0x92, 0x16, 0x98,
	0xc9, 0xed, 0xfc, 0x6a, 0xc1, 0xe6, 0x95, 0xbc, 0x14, 0xeb, 0x80, 0x71, 0x21, 0xbd, 0x73, 0x72,
	0x36, 0x4a, 0x58, 0xb3, 0x18, 0x34, 0xf4, 0x5a, 0x21, 0xe8, 0x13, 0xd8, 0xd0, 0x47, 0x8f, 0x09,
	0x27, 0x01, 0x95, 0x94, 0x9b, 0xd7, 0xc3, 0xba, 0x42, 0x3b, 0x29, 0xa8, 0x78, 0xd4, 0xeb, 0xe2,
	0xc2, 0xf3, 0xa3, 0x51, 0x28, 0x75, 0x97, 0x2d, 0x63, 0xd0, 0x50, 0x5d, 0x21, 0xe8, 0x63, 0xc8,
	0x53, 0x15, 0x94, 0xf6, 0xf5, 0x7b, 0x44, 0xff, 0xd5, 0xf3, 0xd8, 0x36, 0x98, 0x7a, 0x87, 0xec,
	0xd0, 0xf4, 0x4e, 0xe8, 0xc9, 0xff, 0x11, 0x3c, 0xea, 0xbd, 0xc2, 0x6e, 0xad, 0x77, 0xdd, 0xdc,
	0xb7, 0x61, 0x75, 0xbf, 0xb6, 0x77, 0x54, 0xc3, 0x6e, 0xc1, 0x42, 0x0f, 0x01, 0x75, 0x0f, 0xea,
	0xad, 0xda, 0x9e, 0xe7, 0xb6, 0x5f, 0xb6, 0xda, 0xae, 0x8b, 0x5b, 0xed, 0x97, 0x85, 0x0c, 0x7a,
	0x00, 0x9b, 0x87, 0xed, 0xa3, 0x5a, 0xbb, 0xe7, 0x36, 0xbc, 0xee, 0x41, 0xb3, 0xa7, 0xcd, 0xb3,
	0x3b, 0x4d, 0xb8, 0xf7, 0xd6, 0x63, 0x05, 0x6d, 0xc3, 0xe3, 0xfa, 0xc1, 0x7e, 0x07, 0xbb, 0xdd,
	0x6e, 0xeb, 0xa0, 0x7d, 0x5d, 0xc0, 0x55, 0xc8, 0xe2, 0xda, 0x51, 0xc1, 0x52, 0x7f, 0x1c, 0xdc,
	0xaa, 0xbb, 0x85, 0xcc, 0xf1, 0x8a, 0xbe, 0x3b, 0x5f, 0xfd, 0x17, 0x00, 0x00, 0xff, 0xff, 0xf9,
	0x0f, 0x45, 0xae, 0xda, 0x0a, 0x00, 0x00,
}
