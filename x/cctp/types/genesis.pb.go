// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: circle/cctp/v1/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// GenesisState defines the cctp module's genesis state.
type GenesisState struct {
	Owner                             string                             `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	AttesterManager                   string                             `protobuf:"bytes,3,opt,name=attester_manager,json=attesterManager,proto3" json:"attester_manager,omitempty"`
	Pauser                            string                             `protobuf:"bytes,4,opt,name=pauser,proto3" json:"pauser,omitempty"`
	TokenController                   string                             `protobuf:"bytes,5,opt,name=token_controller,json=tokenController,proto3" json:"token_controller,omitempty"`
	AttesterList                      []Attester                         `protobuf:"bytes,6,rep,name=attester_list,json=attesterList,proto3" json:"attester_list"`
	PerMessageBurnLimitList           []PerMessageBurnLimit              `protobuf:"bytes,7,rep,name=per_message_burn_limit_list,json=perMessageBurnLimitList,proto3" json:"per_message_burn_limit_list"`
	BurningAndMintingPaused           *BurningAndMintingPaused           `protobuf:"bytes,8,opt,name=burning_and_minting_paused,json=burningAndMintingPaused,proto3" json:"burning_and_minting_paused,omitempty"`
	SendingAndReceivingMessagesPaused *SendingAndReceivingMessagesPaused `protobuf:"bytes,9,opt,name=sending_and_receiving_messages_paused,json=sendingAndReceivingMessagesPaused,proto3" json:"sending_and_receiving_messages_paused,omitempty"`
	MaxMessageBodySize                *MaxMessageBodySize                `protobuf:"bytes,10,opt,name=max_message_body_size,json=maxMessageBodySize,proto3" json:"max_message_body_size,omitempty"`
	NextAvailableNonce                *Nonce                             `protobuf:"bytes,11,opt,name=next_available_nonce,json=nextAvailableNonce,proto3" json:"next_available_nonce,omitempty"`
	SignatureThreshold                *SignatureThreshold                `protobuf:"bytes,12,opt,name=signature_threshold,json=signatureThreshold,proto3" json:"signature_threshold,omitempty"`
	TokenPairList                     []TokenPair                        `protobuf:"bytes,13,rep,name=token_pair_list,json=tokenPairList,proto3" json:"token_pair_list"`
	UsedNoncesList                    []Nonce                            `protobuf:"bytes,14,rep,name=used_nonces_list,json=usedNoncesList,proto3" json:"used_nonces_list"`
	TokenMessengerList                []RemoteTokenMessenger             `protobuf:"bytes,15,rep,name=token_messenger_list,json=tokenMessengerList,proto3" json:"token_messenger_list"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_2053ebb2404c1e41, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *GenesisState) GetAttesterManager() string {
	if m != nil {
		return m.AttesterManager
	}
	return ""
}

func (m *GenesisState) GetPauser() string {
	if m != nil {
		return m.Pauser
	}
	return ""
}

func (m *GenesisState) GetTokenController() string {
	if m != nil {
		return m.TokenController
	}
	return ""
}

func (m *GenesisState) GetAttesterList() []Attester {
	if m != nil {
		return m.AttesterList
	}
	return nil
}

func (m *GenesisState) GetPerMessageBurnLimitList() []PerMessageBurnLimit {
	if m != nil {
		return m.PerMessageBurnLimitList
	}
	return nil
}

func (m *GenesisState) GetBurningAndMintingPaused() *BurningAndMintingPaused {
	if m != nil {
		return m.BurningAndMintingPaused
	}
	return nil
}

func (m *GenesisState) GetSendingAndReceivingMessagesPaused() *SendingAndReceivingMessagesPaused {
	if m != nil {
		return m.SendingAndReceivingMessagesPaused
	}
	return nil
}

func (m *GenesisState) GetMaxMessageBodySize() *MaxMessageBodySize {
	if m != nil {
		return m.MaxMessageBodySize
	}
	return nil
}

func (m *GenesisState) GetNextAvailableNonce() *Nonce {
	if m != nil {
		return m.NextAvailableNonce
	}
	return nil
}

func (m *GenesisState) GetSignatureThreshold() *SignatureThreshold {
	if m != nil {
		return m.SignatureThreshold
	}
	return nil
}

func (m *GenesisState) GetTokenPairList() []TokenPair {
	if m != nil {
		return m.TokenPairList
	}
	return nil
}

func (m *GenesisState) GetUsedNoncesList() []Nonce {
	if m != nil {
		return m.UsedNoncesList
	}
	return nil
}

func (m *GenesisState) GetTokenMessengerList() []RemoteTokenMessenger {
	if m != nil {
		return m.TokenMessengerList
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "circle.cctp.v1.GenesisState")
}

func init() { proto.RegisterFile("circle/cctp/v1/genesis.proto", fileDescriptor_2053ebb2404c1e41) }

var fileDescriptor_2053ebb2404c1e41 = []byte{
	// 690 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xcd, 0x6e, 0x13, 0x3b,
	0x14, 0x4e, 0xfa, 0x77, 0x5b, 0xf7, 0x57, 0xbe, 0xe9, 0xed, 0xdc, 0xdc, 0x4b, 0x5a, 0x0a, 0x88,
	0x00, 0x22, 0xa3, 0x96, 0x1d, 0xbb, 0xa4, 0x42, 0x15, 0xa8, 0x41, 0xd5, 0xa4, 0x6c, 0x10, 0xd2,
	0xc8, 0x33, 0x73, 0x98, 0x58, 0x9d, 0xb1, 0x23, 0xdb, 0x49, 0xd3, 0x6e, 0x79, 0x01, 0x5e, 0x80,
	0xf7, 0xe9, 0xb2, 0x4b, 0x56, 0x08, 0xb5, 0x2f, 0x82, 0xc6, 0x1e, 0xb7, 0xaa, 0x9b, 0x0a, 0x76,
	0xe3, 0xef, 0x7c, 0xdf, 0x77, 0x7c, 0xc6, 0xe7, 0x1c, 0xf4, 0x7f, 0x4c, 0x45, 0x9c, 0x81, 0x1f,
	0xc7, 0x6a, 0xe0, 0x8f, 0x76, 0xfc, 0x14, 0x18, 0x48, 0x2a, 0x5b, 0x03, 0xc1, 0x15, 0xc7, 0x2b,
	0x26, 0xda, 0x2a, 0xa2, 0xad, 0xd1, 0x4e, 0xfd, 0x81, 0xc3, 0x26, 0x4a, 0x81, 0x54, 0x20, 0x0c,
	0xbd, 0xee, 0x3b, 0xe1, 0x68, 0x28, 0x18, 0x65, 0x69, 0x48, 0x58, 0x12, 0xe6, 0x94, 0xa9, 0xe2,
	0x7b, 0x40, 0x86, 0x12, 0x92, 0x52, 0xf0, 0xdc, 0x11, 0xe4, 0x64, 0x1c, 0xe6, 0x20, 0x25, 0x49,
	0x21, 0x8c, 0x78, 0x72, 0x1a, 0x4a, 0x7a, 0x06, 0x25, 0xb7, 0xee, 0x70, 0x19, 0x67, 0xb1, 0x8d,
	0xbd, 0x70, 0x62, 0x03, 0x10, 0x37, 0x3e, 0x43, 0xc1, 0xc2, 0x8c, 0xe6, 0x54, 0xdd, 0x43, 0x16,
	0x90, 0x73, 0x05, 0xa1, 0xe2, 0xc7, 0xc0, 0xb4, 0x0a, 0x58, 0x7a, 0x5d, 0xd2, 0x6b, 0x87, 0x2c,
	0x81, 0x25, 0xb6, 0x24, 0x01, 0x31, 0xd0, 0x51, 0x71, 0x2a, 0x73, 0xc9, 0xdb, 0xd5, 0x35, 0x5d,
	0x2d, 0x4d, 0x19, 0x51, 0x43, 0x01, 0xa1, 0xea, 0x0b, 0x90, 0x7d, 0x9e, 0x59, 0xe6, 0xa6, 0xc3,
	0x34, 0x77, 0x19, 0x10, 0x6a, 0xaf, 0x51, 0x4b, 0x79, 0xca, 0xf5, 0xa7, 0x5f, 0x7c, 0x19, 0x74,
	0xfb, 0xdb, 0x3c, 0x5a, 0xda, 0x37, 0x0f, 0xd6, 0x53, 0x44, 0x01, 0xae, 0xa1, 0x59, 0x7e, 0xc2,
	0x40, 0x78, 0x53, 0x5b, 0xd5, 0xe6, 0x42, 0x60, 0x0e, 0xf8, 0x19, 0x5a, 0xb3, 0x0f, 0x15, 0xe6,
	0x84, 0x91, 0x14, 0x84, 0x37, 0xad, 0x09, 0xab, 0x16, 0xef, 0x1a, 0x18, 0xff, 0x83, 0xe6, 0x74,
	0x09, 0xc2, 0x9b, 0xd1, 0x84, 0xf2, 0x54, 0x58, 0x98, 0x3b, 0xc5, 0x9c, 0x29, 0xc1, 0xb3, 0x0c,
	0x84, 0x37, 0x6b, 0x2c, 0x34, 0xbe, 0x77, 0x0d, 0xe3, 0x3d, 0xb4, 0x7c, 0x9d, 0x2d, 0xa3, 0x52,
	0x79, 0x73, 0x5b, 0xd3, 0xcd, 0xc5, 0x5d, 0xaf, 0x75, 0xbb, 0x97, 0x5a, 0xed, 0x92, 0xd4, 0x99,
	0x39, 0xff, 0xb1, 0x59, 0x09, 0x96, 0xac, 0xe8, 0x80, 0x4a, 0x85, 0x53, 0xf4, 0xdf, 0xe4, 0x37,
	0x34, 0x96, 0x7f, 0x69, 0xcb, 0x47, 0xae, 0xe5, 0x21, 0x88, 0xae, 0x51, 0x74, 0x86, 0x82, 0x1d,
	0x14, 0xfc, 0xd2, 0x7d, 0x63, 0x70, 0x37, 0xa4, 0x13, 0x25, 0xa8, 0x7e, 0x7f, 0x97, 0x7a, 0xf3,
	0x5b, 0xd5, 0xe6, 0xe2, 0xee, 0x53, 0x37, 0x4f, 0xc7, 0x28, 0xda, 0x2c, 0xe9, 0x1a, 0xfe, 0xa1,
	0xa6, 0x07, 0x1b, 0xd1, 0xe4, 0x00, 0xfe, 0x52, 0x45, 0x4f, 0xfe, 0xa8, 0x73, 0xbc, 0x05, 0x9d,
	0x71, 0xc7, 0xcd, 0xd8, 0x33, 0xe2, 0x36, 0x4b, 0x02, 0x2b, 0x2d, 0xcb, 0x91, 0x65, 0xee, 0x87,
	0xf2, 0x77, 0x14, 0xfc, 0x01, 0xad, 0x4f, 0x1c, 0x30, 0x0f, 0xe9, 0xa4, 0xdb, 0x6e, 0xd2, 0x2e,
	0x19, 0xdb, 0x7f, 0xc6, 0x93, 0xd3, 0x1e, 0x3d, 0x83, 0x00, 0xe7, 0x77, 0x30, 0xbc, 0x8f, 0x6a,
	0x0c, 0xc6, 0x2a, 0x24, 0x23, 0x42, 0x33, 0x12, 0x65, 0x10, 0xea, 0xd1, 0xf4, 0x16, 0xb5, 0xeb,
	0xba, 0xeb, 0xfa, 0xbe, 0x08, 0x06, 0xb8, 0x90, 0xb4, 0xad, 0x42, 0x63, 0xb8, 0x87, 0xfe, 0x9e,
	0x30, 0x22, 0xde, 0xd2, 0xe4, 0xdb, 0xf5, 0x2c, 0xf5, 0xc8, 0x32, 0x03, 0x2c, 0xef, 0x60, 0x78,
	0x1f, 0xad, 0xde, 0x4c, 0x93, 0xe9, 0x9e, 0x65, 0xdd, 0x3d, 0xff, 0xba, 0x86, 0x47, 0x05, 0xed,
	0x90, 0x50, 0xdb, 0x91, 0xcb, 0xca, 0x02, 0xba, 0x53, 0xde, 0xa0, 0xb5, 0xe2, 0x2f, 0x9a, 0xe2,
	0xa4, 0x71, 0x5a, 0xd1, 0x4e, 0x93, 0x4b, 0x2c, 0x5d, 0x56, 0x0a, 0x91, 0x06, 0xa4, 0xb6, 0xf9,
	0x84, 0x6a, 0xce, 0xa6, 0x31, 0x56, 0xab, 0xda, 0xea, 0xb1, 0x6b, 0x15, 0xe8, 0xe5, 0xa4, 0xaf,
	0xd6, 0xb5, 0x82, 0xd2, 0x19, 0xab, 0x5b, 0x68, 0xe1, 0xfe, 0x6e, 0x66, 0xbe, 0xba, 0x36, 0x55,
	0x4c, 0xad, 0x20, 0xb9, 0xec, 0xbc, 0x3d, 0xbf, 0x6c, 0x54, 0x2f, 0x2e, 0x1b, 0xd5, 0x9f, 0x97,
	0x8d, 0xea, 0xd7, 0xab, 0x46, 0xe5, 0xe2, 0xaa, 0x51, 0xf9, 0x7e, 0xd5, 0xa8, 0x7c, 0xf4, 0x53,
	0xaa, 0xfa, 0xc3, 0xa8, 0x15, 0xf3, 0xdc, 0x3f, 0xf9, 0x1c, 0x65, 0x3c, 0x3e, 0x8e, 0xfb, 0x84,
	0x32, 0x9f, 0xf1, 0x28, 0x83, 0x97, 0x7a, 0x0f, 0x8d, 0xcd, 0x3a, 0x52, 0xa7, 0x03, 0x90, 0xd1,
	0x9c, 0xde, 0x38, 0xaf, 0x7e, 0x05, 0x00, 0x00, 0xff, 0xff, 0x7d, 0xf5, 0x82, 0x75, 0x30, 0x06,
	0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TokenMessengerList) > 0 {
		for iNdEx := len(m.TokenMessengerList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TokenMessengerList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x7a
		}
	}
	if len(m.UsedNoncesList) > 0 {
		for iNdEx := len(m.UsedNoncesList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.UsedNoncesList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x72
		}
	}
	if len(m.TokenPairList) > 0 {
		for iNdEx := len(m.TokenPairList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TokenPairList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x6a
		}
	}
	if m.SignatureThreshold != nil {
		{
			size, err := m.SignatureThreshold.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x62
	}
	if m.NextAvailableNonce != nil {
		{
			size, err := m.NextAvailableNonce.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x5a
	}
	if m.MaxMessageBodySize != nil {
		{
			size, err := m.MaxMessageBodySize.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x52
	}
	if m.SendingAndReceivingMessagesPaused != nil {
		{
			size, err := m.SendingAndReceivingMessagesPaused.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x4a
	}
	if m.BurningAndMintingPaused != nil {
		{
			size, err := m.BurningAndMintingPaused.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x42
	}
	if len(m.PerMessageBurnLimitList) > 0 {
		for iNdEx := len(m.PerMessageBurnLimitList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PerMessageBurnLimitList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.AttesterList) > 0 {
		for iNdEx := len(m.AttesterList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AttesterList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.TokenController) > 0 {
		i -= len(m.TokenController)
		copy(dAtA[i:], m.TokenController)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.TokenController)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Pauser) > 0 {
		i -= len(m.Pauser)
		copy(dAtA[i:], m.Pauser)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Pauser)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.AttesterManager) > 0 {
		i -= len(m.AttesterManager)
		copy(dAtA[i:], m.AttesterManager)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.AttesterManager)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.AttesterManager)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.Pauser)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.TokenController)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.AttesterList) > 0 {
		for _, e := range m.AttesterList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.PerMessageBurnLimitList) > 0 {
		for _, e := range m.PerMessageBurnLimitList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.BurningAndMintingPaused != nil {
		l = m.BurningAndMintingPaused.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.SendingAndReceivingMessagesPaused != nil {
		l = m.SendingAndReceivingMessagesPaused.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.MaxMessageBodySize != nil {
		l = m.MaxMessageBodySize.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.NextAvailableNonce != nil {
		l = m.NextAvailableNonce.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.SignatureThreshold != nil {
		l = m.SignatureThreshold.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.TokenPairList) > 0 {
		for _, e := range m.TokenPairList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.UsedNoncesList) > 0 {
		for _, e := range m.UsedNoncesList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.TokenMessengerList) > 0 {
		for _, e := range m.TokenMessengerList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AttesterManager", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AttesterManager = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pauser", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pauser = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenController", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenController = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AttesterList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AttesterList = append(m.AttesterList, Attester{})
			if err := m.AttesterList[len(m.AttesterList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PerMessageBurnLimitList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PerMessageBurnLimitList = append(m.PerMessageBurnLimitList, PerMessageBurnLimit{})
			if err := m.PerMessageBurnLimitList[len(m.PerMessageBurnLimitList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BurningAndMintingPaused", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BurningAndMintingPaused == nil {
				m.BurningAndMintingPaused = &BurningAndMintingPaused{}
			}
			if err := m.BurningAndMintingPaused.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SendingAndReceivingMessagesPaused", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SendingAndReceivingMessagesPaused == nil {
				m.SendingAndReceivingMessagesPaused = &SendingAndReceivingMessagesPaused{}
			}
			if err := m.SendingAndReceivingMessagesPaused.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxMessageBodySize", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MaxMessageBodySize == nil {
				m.MaxMessageBodySize = &MaxMessageBodySize{}
			}
			if err := m.MaxMessageBodySize.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextAvailableNonce", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.NextAvailableNonce == nil {
				m.NextAvailableNonce = &Nonce{}
			}
			if err := m.NextAvailableNonce.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignatureThreshold", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SignatureThreshold == nil {
				m.SignatureThreshold = &SignatureThreshold{}
			}
			if err := m.SignatureThreshold.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenPairList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenPairList = append(m.TokenPairList, TokenPair{})
			if err := m.TokenPairList[len(m.TokenPairList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UsedNoncesList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UsedNoncesList = append(m.UsedNoncesList, Nonce{})
			if err := m.UsedNoncesList[len(m.UsedNoncesList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenMessengerList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenMessengerList = append(m.TokenMessengerList, RemoteTokenMessenger{})
			if err := m.TokenMessengerList[len(m.TokenMessengerList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
