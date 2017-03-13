// Code generated by protoc-gen-go.
// source: export/export.proto
// DO NOT EDIT!

package export

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import ketchup_models "github.com/octavore/ketchup/proto/ketchup/models"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Export struct {
	Pages            []*ketchup_models.Page  `protobuf:"bytes,1,rep,name=pages" json:"pages,omitempty"`
	Routes           []*ketchup_models.Route `protobuf:"bytes,2,rep,name=routes" json:"routes,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *Export) Reset()                    { *m = Export{} }
func (m *Export) String() string            { return proto.CompactTextString(m) }
func (*Export) ProtoMessage()               {}
func (*Export) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *Export) GetPages() []*ketchup_models.Page {
	if m != nil {
		return m.Pages
	}
	return nil
}

func (m *Export) GetRoutes() []*ketchup_models.Route {
	if m != nil {
		return m.Routes
	}
	return nil
}

func init() {
	proto.RegisterType((*Export)(nil), "ketchup.models.export.Export")
}
func (m *Export) SetPages(v []*ketchup_models.Page) {
	m.Pages = v
}

func (m *Export) SetRoutes(v []*ketchup_models.Route) {
	m.Routes = v
}

func init() { proto.RegisterFile("export/export.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 128 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x4e, 0xad, 0x28, 0xc8,
	0x2f, 0x2a, 0xd1, 0x87, 0x50, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xa2, 0xd9, 0xa9, 0x25,
	0xc9, 0x19, 0xa5, 0x05, 0x7a, 0xb9, 0xf9, 0x29, 0xa9, 0x39, 0xc5, 0x7a, 0x10, 0x49, 0x29, 0xae,
	0x82, 0xc4, 0xf4, 0x54, 0x88, 0x12, 0x29, 0xee, 0xa2, 0xfc, 0xd2, 0x12, 0x28, 0x47, 0x29, 0x99,
	0x8b, 0xcd, 0x15, 0xac, 0x44, 0x48, 0x8b, 0x8b, 0x15, 0xa4, 0xa8, 0x58, 0x82, 0x51, 0x81, 0x59,
	0x83, 0xdb, 0x48, 0x44, 0x0f, 0xcd, 0xa4, 0x80, 0xc4, 0xf4, 0xd4, 0x20, 0x88, 0x12, 0x21, 0x5d,
	0x2e, 0x36, 0xb0, 0x21, 0xc5, 0x12, 0x4c, 0x60, 0xc5, 0xa2, 0xe8, 0x8a, 0x83, 0x40, 0xb2, 0x41,
	0x50, 0x45, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xed, 0x53, 0xdd, 0x4f, 0xaa, 0x00, 0x00, 0x00,
}
