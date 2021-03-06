// Code generated by protoc-gen-go. DO NOT EDIT.
// source: schema.tl.handshake_service.proto

package mtproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// /////////////////////////////////////////////////////////////////////////////
// req_pq#60469778 nonce:int128 = ResPQ;
type TLReqPq struct {
	Nonce []byte `protobuf:"bytes,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
}

func (m *TLReqPq) Reset()                    { *m = TLReqPq{} }
func (m *TLReqPq) String() string            { return proto.CompactTextString(m) }
func (*TLReqPq) ProtoMessage()               {}
func (*TLReqPq) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *TLReqPq) GetNonce() []byte {
	if m != nil {
		return m.Nonce
	}
	return nil
}

// /////////////////////////////////////////////////////////////////////////////
// req_DH_params#d712e4be nonce:int128 server_nonce:int128 p:string q:string public_key_fingerprint:long encrypted_data:string = Server_DH_Params;
type TLReq_DHParams struct {
	Nonce                []byte `protobuf:"bytes,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
	ServerNonce          []byte `protobuf:"bytes,2,opt,name=server_nonce,json=serverNonce,proto3" json:"server_nonce,omitempty"`
	P                    string `protobuf:"bytes,3,opt,name=p" json:"p,omitempty"`
	Q                    string `protobuf:"bytes,4,opt,name=q" json:"q,omitempty"`
	PublicKeyFingerprint int64  `protobuf:"varint,5,opt,name=public_key_fingerprint,json=publicKeyFingerprint" json:"public_key_fingerprint,omitempty"`
	EncryptedData        string `protobuf:"bytes,6,opt,name=encrypted_data,json=encryptedData" json:"encrypted_data,omitempty"`
}

func (m *TLReq_DHParams) Reset()                    { *m = TLReq_DHParams{} }
func (m *TLReq_DHParams) String() string            { return proto.CompactTextString(m) }
func (*TLReq_DHParams) ProtoMessage()               {}
func (*TLReq_DHParams) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *TLReq_DHParams) GetNonce() []byte {
	if m != nil {
		return m.Nonce
	}
	return nil
}

func (m *TLReq_DHParams) GetServerNonce() []byte {
	if m != nil {
		return m.ServerNonce
	}
	return nil
}

func (m *TLReq_DHParams) GetP() string {
	if m != nil {
		return m.P
	}
	return ""
}

func (m *TLReq_DHParams) GetQ() string {
	if m != nil {
		return m.Q
	}
	return ""
}

func (m *TLReq_DHParams) GetPublicKeyFingerprint() int64 {
	if m != nil {
		return m.PublicKeyFingerprint
	}
	return 0
}

func (m *TLReq_DHParams) GetEncryptedData() string {
	if m != nil {
		return m.EncryptedData
	}
	return ""
}

// /////////////////////////////////////////////////////////////////////////////
// set_client_DH_params#f5045f1f nonce:int128 server_nonce:int128 encrypted_data:string = Set_client_DH_params_answer;
type TLSetClient_DHParams struct {
	Nonce         []byte `protobuf:"bytes,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
	ServerNonce   []byte `protobuf:"bytes,2,opt,name=server_nonce,json=serverNonce,proto3" json:"server_nonce,omitempty"`
	EncryptedData string `protobuf:"bytes,3,opt,name=encrypted_data,json=encryptedData" json:"encrypted_data,omitempty"`
}

func (m *TLSetClient_DHParams) Reset()                    { *m = TLSetClient_DHParams{} }
func (m *TLSetClient_DHParams) String() string            { return proto.CompactTextString(m) }
func (*TLSetClient_DHParams) ProtoMessage()               {}
func (*TLSetClient_DHParams) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

func (m *TLSetClient_DHParams) GetNonce() []byte {
	if m != nil {
		return m.Nonce
	}
	return nil
}

func (m *TLSetClient_DHParams) GetServerNonce() []byte {
	if m != nil {
		return m.ServerNonce
	}
	return nil
}

func (m *TLSetClient_DHParams) GetEncryptedData() string {
	if m != nil {
		return m.EncryptedData
	}
	return ""
}

// /////////////////////////////////////////////////////////////////////////////
// destroy_auth_key#d1435160 = DestroyAuthKeyRes;
type TLDestroyAuthKey struct {
}

func (m *TLDestroyAuthKey) Reset()                    { *m = TLDestroyAuthKey{} }
func (m *TLDestroyAuthKey) String() string            { return proto.CompactTextString(m) }
func (*TLDestroyAuthKey) ProtoMessage()               {}
func (*TLDestroyAuthKey) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{3} }

func init() {
	proto.RegisterType((*TLReqPq)(nil), "mtproto.TL_req_pq")
	proto.RegisterType((*TLReq_DHParams)(nil), "mtproto.TL_req_DH_params")
	proto.RegisterType((*TLSetClient_DHParams)(nil), "mtproto.TL_set_client_DH_params")
	proto.RegisterType((*TLDestroyAuthKey)(nil), "mtproto.TL_destroy_auth_key")
}

func init() { proto.RegisterFile("schema.tl.handshake_service.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 291 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x91, 0x41, 0x4b, 0xc3, 0x30,
	0x1c, 0xc5, 0xc9, 0xe6, 0x36, 0x16, 0xa7, 0x48, 0x9d, 0x5a, 0xf0, 0xb2, 0x15, 0x84, 0x9e, 0x72,
	0xd1, 0x4f, 0x30, 0x86, 0x0c, 0xac, 0x22, 0xa5, 0xf7, 0x3f, 0x69, 0xfa, 0x77, 0x0d, 0x6b, 0xd3,
	0x34, 0x49, 0x85, 0x7e, 0x3f, 0x3f, 0x98, 0xb4, 0x1d, 0x7a, 0x70, 0x47, 0x6f, 0xff, 0xf7, 0x7e,
	0x8f, 0x3c, 0x1e, 0xa1, 0x6b, 0x2b, 0x72, 0x2c, 0x39, 0x73, 0x05, 0xcb, 0xb9, 0xca, 0x6c, 0xce,
	0x0f, 0x08, 0x16, 0xcd, 0xa7, 0x14, 0xc8, 0xb4, 0xa9, 0x5c, 0xe5, 0xcd, 0x4a, 0xd7, 0x1f, 0xc1,
	0x9a, 0xce, 0x93, 0x08, 0x0c, 0xd6, 0xa0, 0x6b, 0x6f, 0x49, 0x27, 0xaa, 0x52, 0x02, 0x7d, 0xb2,
	0x22, 0xe1, 0x22, 0x1e, 0x44, 0xf0, 0x45, 0xe8, 0xd5, 0x31, 0xb3, 0xdd, 0x81, 0xe6, 0x86, 0x97,
	0xf6, 0x74, 0xd4, 0x5b, 0xd3, 0x45, 0xd7, 0x83, 0x06, 0x06, 0x38, 0xea, 0xe1, 0xf9, 0xe0, 0xbd,
	0xf5, 0x91, 0x05, 0x25, 0xda, 0x1f, 0xaf, 0x48, 0x38, 0x8f, 0x89, 0xee, 0x54, 0xed, 0x9f, 0x0d,
	0xaa, 0xf6, 0x9e, 0xe8, 0xad, 0x6e, 0xd2, 0x42, 0x0a, 0x38, 0x60, 0x0b, 0x1f, 0x52, 0xed, 0xd1,
	0x68, 0x23, 0x95, 0xf3, 0x27, 0x2b, 0x12, 0x8e, 0xe3, 0xe5, 0x40, 0x5f, 0xb0, 0x7d, 0xfe, 0x65,
	0xde, 0x03, 0xbd, 0x44, 0x25, 0x4c, 0xab, 0x1d, 0x66, 0x90, 0x71, 0xc7, 0xfd, 0x69, 0xff, 0xe0,
	0xc5, 0x8f, 0xbb, 0xe5, 0x8e, 0x07, 0x2d, 0xbd, 0x4b, 0x22, 0xb0, 0xe8, 0x40, 0x14, 0x12, 0x95,
	0xfb, 0x8f, 0x31, 0x7f, 0xab, 0xc7, 0xa7, 0xaa, 0x6f, 0xe8, 0x75, 0x12, 0x41, 0x86, 0xd6, 0x99,
	0xaa, 0x05, 0xde, 0xb8, 0xbc, 0x1b, 0xb8, 0x09, 0xe9, 0xbd, 0xa8, 0x4a, 0xa6, 0x30, 0x6d, 0x0a,
	0x2e, 0x4b, 0x86, 0x6a, 0x2f, 0x15, 0xb2, 0xe3, 0xd7, 0x6c, 0x66, 0xaf, 0xc9, 0x7b, 0x77, 0xec,
	0x46, 0xe9, 0xb4, 0x77, 0x1e, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x14, 0x59, 0xe1, 0xb2, 0xda,
	0x01, 0x00, 0x00,
}
