// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/extensions/filters/network/thrift_proxy/v2alpha1/route.proto

package v2

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/lyft/protoc-gen-validate/validate"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// [#comment:next free field: 3]
type RouteConfiguration struct {
	// The name of the route configuration. Reserved for future use in asynchronous route discovery.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The list of routes that will be matched, in order, against incoming requests. The first route
	// that matches will be used.
	Routes               []Route  `protobuf:"bytes,2,rep,name=routes" json:"routes"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RouteConfiguration) Reset()         { *m = RouteConfiguration{} }
func (m *RouteConfiguration) String() string { return proto.CompactTextString(m) }
func (*RouteConfiguration) ProtoMessage()    {}
func (*RouteConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_route_9d807bfa1b820735, []int{0}
}
func (m *RouteConfiguration) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RouteConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RouteConfiguration.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *RouteConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteConfiguration.Merge(dst, src)
}
func (m *RouteConfiguration) XXX_Size() int {
	return m.Size()
}
func (m *RouteConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_RouteConfiguration proto.InternalMessageInfo

func (m *RouteConfiguration) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RouteConfiguration) GetRoutes() []Route {
	if m != nil {
		return m.Routes
	}
	return nil
}

// [#comment:next free field: 3]
type Route struct {
	// Route matching prarameters.
	Match RouteMatch `protobuf:"bytes,1,opt,name=match" json:"match"`
	// Route request to some upstream cluster.
	Route                RouteAction `protobuf:"bytes,2,opt,name=route" json:"route"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Route) Reset()         { *m = Route{} }
func (m *Route) String() string { return proto.CompactTextString(m) }
func (*Route) ProtoMessage()    {}
func (*Route) Descriptor() ([]byte, []int) {
	return fileDescriptor_route_9d807bfa1b820735, []int{1}
}
func (m *Route) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Route) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Route.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Route) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Route.Merge(dst, src)
}
func (m *Route) XXX_Size() int {
	return m.Size()
}
func (m *Route) XXX_DiscardUnknown() {
	xxx_messageInfo_Route.DiscardUnknown(m)
}

var xxx_messageInfo_Route proto.InternalMessageInfo

func (m *Route) GetMatch() RouteMatch {
	if m != nil {
		return m.Match
	}
	return RouteMatch{}
}

func (m *Route) GetRoute() RouteAction {
	if m != nil {
		return m.Route
	}
	return RouteAction{}
}

// [#comment:next free field: 2]
type RouteMatch struct {
	// If specified, the route must exactly match the request method name. As a special case, an
	// empty string matches any request method name.
	Method               string   `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RouteMatch) Reset()         { *m = RouteMatch{} }
func (m *RouteMatch) String() string { return proto.CompactTextString(m) }
func (*RouteMatch) ProtoMessage()    {}
func (*RouteMatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_route_9d807bfa1b820735, []int{2}
}
func (m *RouteMatch) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RouteMatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RouteMatch.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *RouteMatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteMatch.Merge(dst, src)
}
func (m *RouteMatch) XXX_Size() int {
	return m.Size()
}
func (m *RouteMatch) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteMatch.DiscardUnknown(m)
}

var xxx_messageInfo_RouteMatch proto.InternalMessageInfo

func (m *RouteMatch) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

// [#comment:next free field: 2]
type RouteAction struct {
	// Indicates the upstream cluster to which the request should be routed.
	Cluster              string   `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RouteAction) Reset()         { *m = RouteAction{} }
func (m *RouteAction) String() string { return proto.CompactTextString(m) }
func (*RouteAction) ProtoMessage()    {}
func (*RouteAction) Descriptor() ([]byte, []int) {
	return fileDescriptor_route_9d807bfa1b820735, []int{3}
}
func (m *RouteAction) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RouteAction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RouteAction.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *RouteAction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteAction.Merge(dst, src)
}
func (m *RouteAction) XXX_Size() int {
	return m.Size()
}
func (m *RouteAction) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteAction.DiscardUnknown(m)
}

var xxx_messageInfo_RouteAction proto.InternalMessageInfo

func (m *RouteAction) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func init() {
	proto.RegisterType((*RouteConfiguration)(nil), "envoy.extensions.filters.network.thrift_proxy.v2alpha1.RouteConfiguration")
	proto.RegisterType((*Route)(nil), "envoy.extensions.filters.network.thrift_proxy.v2alpha1.Route")
	proto.RegisterType((*RouteMatch)(nil), "envoy.extensions.filters.network.thrift_proxy.v2alpha1.RouteMatch")
	proto.RegisterType((*RouteAction)(nil), "envoy.extensions.filters.network.thrift_proxy.v2alpha1.RouteAction")
}
func (m *RouteConfiguration) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RouteConfiguration) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRoute(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.Routes) > 0 {
		for _, msg := range m.Routes {
			dAtA[i] = 0x12
			i++
			i = encodeVarintRoute(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Route) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Route) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintRoute(dAtA, i, uint64(m.Match.Size()))
	n1, err := m.Match.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	dAtA[i] = 0x12
	i++
	i = encodeVarintRoute(dAtA, i, uint64(m.Route.Size()))
	n2, err := m.Route.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *RouteMatch) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RouteMatch) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Method) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRoute(dAtA, i, uint64(len(m.Method)))
		i += copy(dAtA[i:], m.Method)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *RouteAction) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RouteAction) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Cluster) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRoute(dAtA, i, uint64(len(m.Cluster)))
		i += copy(dAtA[i:], m.Cluster)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintRoute(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *RouteConfiguration) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovRoute(uint64(l))
	}
	if len(m.Routes) > 0 {
		for _, e := range m.Routes {
			l = e.Size()
			n += 1 + l + sovRoute(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Route) Size() (n int) {
	var l int
	_ = l
	l = m.Match.Size()
	n += 1 + l + sovRoute(uint64(l))
	l = m.Route.Size()
	n += 1 + l + sovRoute(uint64(l))
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *RouteMatch) Size() (n int) {
	var l int
	_ = l
	l = len(m.Method)
	if l > 0 {
		n += 1 + l + sovRoute(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *RouteAction) Size() (n int) {
	var l int
	_ = l
	l = len(m.Cluster)
	if l > 0 {
		n += 1 + l + sovRoute(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovRoute(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozRoute(x uint64) (n int) {
	return sovRoute(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RouteConfiguration) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRoute
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RouteConfiguration: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RouteConfiguration: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoute
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRoute
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Routes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoute
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRoute
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Routes = append(m.Routes, Route{})
			if err := m.Routes[len(m.Routes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRoute(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRoute
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Route) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRoute
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Route: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Route: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Match", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoute
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRoute
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Match.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Route", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoute
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRoute
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Route.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRoute(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRoute
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RouteMatch) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRoute
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RouteMatch: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RouteMatch: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Method", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoute
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRoute
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Method = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRoute(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRoute
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RouteAction) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRoute
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RouteAction: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RouteAction: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cluster", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoute
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRoute
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Cluster = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRoute(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRoute
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipRoute(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRoute
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
					return 0, ErrIntOverflowRoute
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRoute
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthRoute
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRoute
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipRoute(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthRoute = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRoute   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("envoy/extensions/filters/network/thrift_proxy/v2alpha1/route.proto", fileDescriptor_route_9d807bfa1b820735)
}

var fileDescriptor_route_9d807bfa1b820735 = []byte{
	// 341 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x91, 0x31, 0x4f, 0x32, 0x31,
	0x18, 0xc7, 0xdf, 0x1e, 0x1c, 0x6f, 0x28, 0x89, 0x21, 0x8d, 0x51, 0xc2, 0x80, 0xe4, 0x74, 0x60,
	0x6a, 0xe3, 0x99, 0xb8, 0x39, 0x78, 0xcc, 0x2e, 0x37, 0xea, 0x60, 0x2a, 0x14, 0xae, 0xf1, 0x68,
	0x49, 0xef, 0xe1, 0x84, 0x0f, 0xe0, 0xe2, 0x47, 0x72, 0x72, 0x64, 0xf4, 0x03, 0x18, 0x63, 0xd8,
	0xfc, 0x16, 0xa6, 0xbd, 0xbb, 0x40, 0xe2, 0x26, 0xdb, 0xbf, 0x4f, 0x9b, 0xdf, 0xef, 0x9f, 0xa7,
	0x38, 0x12, 0x2a, 0xd7, 0x2b, 0x26, 0x96, 0x20, 0x54, 0x26, 0xb5, 0xca, 0xd8, 0x44, 0xa6, 0x20,
	0x4c, 0xc6, 0x94, 0x80, 0x27, 0x6d, 0x1e, 0x19, 0x24, 0x46, 0x4e, 0xe0, 0x7e, 0x6e, 0xf4, 0x72,
	0xc5, 0xf2, 0x90, 0xa7, 0xf3, 0x84, 0x9f, 0x33, 0xa3, 0x17, 0x20, 0xe8, 0xdc, 0x68, 0xd0, 0xe4,
	0xd2, 0x31, 0xe8, 0x96, 0x41, 0x4b, 0x06, 0x2d, 0x19, 0x74, 0x97, 0x41, 0x2b, 0x46, 0xf7, 0x38,
	0xe7, 0xa9, 0x1c, 0x73, 0x10, 0xac, 0x0a, 0x05, 0xb0, 0x7b, 0x38, 0xd5, 0x53, 0xed, 0x22, 0xb3,
	0xa9, 0x98, 0x06, 0xcf, 0x08, 0x93, 0xd8, 0x6a, 0x87, 0x5a, 0x4d, 0xe4, 0x74, 0x61, 0x38, 0x48,
	0xad, 0x08, 0xc1, 0x75, 0xc5, 0x67, 0xa2, 0x83, 0xfa, 0x68, 0xd0, 0x8c, 0x5d, 0x26, 0x77, 0xb8,
	0xe1, 0x0a, 0x66, 0x1d, 0xaf, 0x5f, 0x1b, 0xb4, 0xc2, 0x2b, 0xfa, 0xb7, 0x8a, 0xd4, 0xf9, 0xa2,
	0xfa, 0xfa, 0xf3, 0xe4, 0x5f, 0x5c, 0x22, 0x83, 0x0f, 0x84, 0x7d, 0x37, 0x27, 0x09, 0xf6, 0x67,
	0x1c, 0x46, 0x89, 0x73, 0xb7, 0xc2, 0x68, 0x2f, 0xcb, 0x8d, 0x25, 0x45, 0x07, 0x56, 0xf5, 0xfa,
	0xfd, 0x56, 0xf3, 0x5f, 0x90, 0xd7, 0x46, 0x71, 0x21, 0x20, 0x12, 0xfb, 0xce, 0xde, 0xf1, 0x9c,
	0x69, 0xb8, 0x97, 0xe9, 0x7a, 0x64, 0x17, 0xf7, 0x5b, 0xe5, 0x0c, 0xc1, 0x19, 0xc6, 0xdb, 0x3e,
	0xe4, 0x08, 0x37, 0x66, 0x02, 0x12, 0x3d, 0x2e, 0xf7, 0x5b, 0x9e, 0x82, 0x10, 0xb7, 0x76, 0x58,
	0xe4, 0x14, 0xff, 0x1f, 0xa5, 0x8b, 0x0c, 0x84, 0x29, 0xde, 0x45, 0x4d, 0x0b, 0xae, 0x1b, 0xaf,
	0x8f, 0xe2, 0xea, 0x26, 0x6a, 0xaf, 0x37, 0x3d, 0xf4, 0xbe, 0xe9, 0xa1, 0xaf, 0x4d, 0x0f, 0xdd,
	0x7a, 0x79, 0xf8, 0xd0, 0x70, 0x3f, 0x7b, 0xf1, 0x13, 0x00, 0x00, 0xff, 0xff, 0xb1, 0x44, 0x91,
	0xc7, 0x86, 0x02, 0x00, 0x00,
}
