package hipstershop

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CartItem struct {
	ProductId            string   `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Quantity             int32    `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CartItem) Reset()         { *m = CartItem{} }
func (m *CartItem) String() string { return proto.CompactTextString(m) }
func (*CartItem) ProtoMessage()    {}
func (*CartItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca53982754088a9d, []int{0}
}

func (m *CartItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CartItem.Unmarshal(m, b)
}
func (m *CartItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CartItem.Marshal(b, m, deterministic)
}
func (m *CartItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CartItem.Merge(m, src)
}
func (m *CartItem) XXX_Size() int {
	return xxx_messageInfo_CartItem.Size(m)
}
func (m *CartItem) XXX_DiscardUnknown() {
	xxx_messageInfo_CartItem.DiscardUnknown(m)
}

var xxx_messageInfo_CartItem proto.InternalMessageInfo

func (m *CartItem) GetProductId() string {
	if m != nil {
		return m.ProductId
	}
	return ""
}

func (m *CartItem) GetQuantity() int32 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

type AddItemRequest struct {
	UserId               string    `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Item                 *CartItem `protobuf:"bytes,2,opt,name=item,proto3" json:"item,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *AddItemRequest) Reset()         { *m = AddItemRequest{} }
func (m *AddItemRequest) String() string { return proto.CompactTextString(m) }
func (*AddItemRequest) ProtoMessage()    {}
func (*AddItemRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca53982754088a9d, []int{1}
}

func (m *AddItemRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddItemRequest.Unmarshal(m, b)
}
func (m *AddItemRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddItemRequest.Marshal(b, m, deterministic)
}
func (m *AddItemRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddItemRequest.Merge(m, src)
}
func (m *AddItemRequest) XXX_Size() int {
	return xxx_messageInfo_AddItemRequest.Size(m)
}
func (m *AddItemRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddItemRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddItemRequest proto.InternalMessageInfo

func (m *AddItemRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *AddItemRequest) GetItem() *CartItem {
	if m != nil {
		return m.Item
	}
	return nil
}

type EmptyCartRequest struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmptyCartRequest) Reset()         { *m = EmptyCartRequest{} }
func (m *EmptyCartRequest) String() string { return proto.CompactTextString(m) }
func (*EmptyCartRequest) ProtoMessage()    {}
func (*EmptyCartRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca53982754088a9d, []int{2}
}

func (m *EmptyCartRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmptyCartRequest.Unmarshal(m, b)
}
func (m *EmptyCartRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmptyCartRequest.Marshal(b, m, deterministic)
}
func (m *EmptyCartRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmptyCartRequest.Merge(m, src)
}
func (m *EmptyCartRequest) XXX_Size() int {
	return xxx_messageInfo_EmptyCartRequest.Size(m)
}
func (m *EmptyCartRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EmptyCartRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EmptyCartRequest proto.InternalMessageInfo

func (m *EmptyCartRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type GetCartRequest struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCartRequest) Reset()         { *m = GetCartRequest{} }
func (m *GetCartRequest) String() string { return proto.CompactTextString(m) }
func (*GetCartRequest) ProtoMessage()    {}
func (*GetCartRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca53982754088a9d, []int{3}
}

func (m *GetCartRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCartRequest.Unmarshal(m, b)
}
func (m *GetCartRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCartRequest.Marshal(b, m, deterministic)
}
func (m *GetCartRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCartRequest.Merge(m, src)
}
func (m *GetCartRequest) XXX_Size() int {
	return xxx_messageInfo_GetCartRequest.Size(m)
}
func (m *GetCartRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCartRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCartRequest proto.InternalMessageInfo

func (m *GetCartRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type Cart struct {
	UserId               string      `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Items                []*CartItem `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Cart) Reset()         { *m = Cart{} }
func (m *Cart) String() string { return proto.CompactTextString(m) }
func (*Cart) ProtoMessage()    {}
func (*Cart) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca53982754088a9d, []int{4}
}

func (m *Cart) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cart.Unmarshal(m, b)
}
func (m *Cart) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cart.Marshal(b, m, deterministic)
}
func (m *Cart) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cart.Merge(m, src)
}
func (m *Cart) XXX_Size() int {
	return xxx_messageInfo_Cart.Size(m)
}
func (m *Cart) XXX_DiscardUnknown() {
	xxx_messageInfo_Cart.DiscardUnknown(m)
}

var xxx_messageInfo_Cart proto.InternalMessageInfo

func (m *Cart) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *Cart) GetItems() []*CartItem {
	if m != nil {
		return m.Items
	}
	return nil
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca53982754088a9d, []int{5}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type ListRecommendationsRequest struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ProductIds           []string `protobuf:"bytes,2,rep,name=product_ids,json=productIds,proto3" json:"product_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRecommendationsRequest) Reset()         { *m = ListRecommendationsRequest{} }
func (m *ListRecommendationsRequest) String() string { return proto.CompactTextString(m) }
func (*ListRecommendationsRequest) ProtoMessage()    {}
func (*ListRecommendationsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca53982754088a9d, []int{6}
}

func (m *ListRecommendationsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRecommendationsRequest.Unmarshal(m, b)
}
func (m *ListRecommendationsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRecommendationsRequest.Marshal(b, m, deterministic)
}
func (m *ListRecommendationsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRecommendationsRequest.Merge(m, src)
}
func (m *ListRecommendationsRequest) XXX_Size() int {
	return xxx_messageInfo_ListRecommendationsRequest.Size(m)
}
func (m *ListRecommendationsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRecommendationsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRecommendationsRequest proto.InternalMessageInfo

func (m *ListRecommendationsRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *ListRecommendationsRequest) GetProductIds() []string {
	if m != nil {
		return m.ProductIds
	}
	return nil
}

type ListRecommendationsResponse struct {
	ProductIds           []string `protobuf:"bytes,1,rep,name=product_ids,json=productIds,proto3" json:"product_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRecommendationsResponse) Reset()         { *m = ListRecommendationsResponse{} }
func (m *ListRecommendationsResponse) String() string { return proto.CompactTextString(m) }
func (*ListRecommendationsResponse) ProtoMessage()    {}
func (*ListRecommendationsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca53982754088a9d, []int{7}
}

func (m *ListRecommendationsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRecommendationsResponse.Unmarshal(m, b)
}
func (m *ListRecommendationsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRecommendationsResponse.Marshal(b, m, deterministic)
}
func (m *ListRecommendationsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRecommendationsResponse.Merge(m, src)
}
func (m *ListRecommendationsResponse) XXX_Size() int {
	return xxx_messageInfo_ListRecommendationsResponse.Size(m)
}
func (m *ListRecommendationsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRecommendationsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListRecommendationsResponse proto.InternalMessageInfo

func (m *ListRecommendationsResponse) GetProductIds() []string {
	if m != nil {
		return m.ProductIds
	}
	return nil
}

type Product struct {
	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Picture     string `protobuf:"bytes,4,opt,name=picture,proto3" json:"picture,omitempty"`
	PriceUsd    *Money `protobuf:"bytes,5,opt,name=price_usd,json=priceUsd,proto3" json:"price_usd,omitempty"`
	// Categories such as "clothing" or "kitchen" that can be used to look up
	// other related products.
	Categories           []string `protobuf:"bytes,6,rep,name=categories,proto3" json:"categories,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Product) Reset()         { *m = Product{} }
func (m *Product) String() string { return proto.CompactTextString(m) }
func (*Product) ProtoMessage()    {}
func (*Product) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca53982754088a9d, []int{8}
}

func (m *Product) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Product.Unmarshal(m, b)
}
func (m *Product) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Product.Marshal(b, m, deterministic)
}
func (m *Product) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Product.Merge(m, src)
}
func (m *Product) XXX_Size() int {
	return xxx_messageInfo_Product.Size(m)
}
func (m *Product) XXX_DiscardUnknown() {
	xxx_messageInfo_Product.DiscardUnknown(m)
}

var xxx_messageInfo_Product proto.InternalMessageInfo

func (m *Product) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Product) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Product) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Product) GetPicture() string {
	if m != nil {
		return m.Picture
	}
	return ""
}

func (m *Product) GetPriceUsd() *Money {
	if m != nil {
		return m.PriceUsd
	}
	return nil
}

func (m *Product) GetCategories() []string {
	if m != nil {
		return m.Categories
	}
	return nil
}

type ListProductsResponse struct {
	Products             []*Product `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ListProductsResponse) Reset()         { *m = ListProductsResponse{} }
func (m *ListProductsResponse) String() string { return proto.CompactTextString(m) }
func (*ListProductsResponse) ProtoMessage()    {}
func (*ListProductsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca53982754088a9d, []int{9}
}

func (m *ListProductsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListProductsResponse.Unmarshal(m, b)
}
func (m *ListProductsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListProductsResponse.Marshal(b, m, deterministic)
}
func (m *ListProductsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListProductsResponse.Merge(m, src)
}
func (m *ListProductsResponse) XXX_Size() int {
	return xxx_messageInfo_ListProductsResponse.Size(m)
}
func (m *ListProductsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListProductsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListProductsResponse proto.InternalMessageInfo

func (m *ListProductsResponse) GetProducts() []*Product {
	if m != nil {
		return m.Products
	}
	return nil
}

type GetProductRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetProductRequest) Reset()         { *m = GetProductRequest{} }
func (m *GetProductRequest) String() string { return proto.CompactTextString(m) }
func (*GetProductRequest) ProtoMessage()    {}
func (*GetProductRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca53982754088a9d, []int{10}
}

func (m *GetProductRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProductRequest.Unmarshal(m, b)
}
func (m *GetProductRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProductRequest.Marshal(b, m, deterministic)
}
func (m *GetProductRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProductRequest.Merge(m, src)
}
func (m *GetProductRequest) XXX_Size() int {
	return xxx_messageInfo_GetProductRequest.Size(m)
}
func (m *GetProductRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProductRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetProductRequest proto.InternalMessageInfo

func (m *GetProductRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type SearchProductsRequest struct {
	Query                string   `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchProductsRequest) Reset()         { *m = SearchProductsRequest{} }
func (m *SearchProductsRequest) String() string { return proto.CompactTextString(m) }
func (*SearchProductsRequest) ProtoMessage()    {}
func (*SearchProductsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca53982754088a9d, []int{11}
}

func (m *SearchProductsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchProductsRequest.Unmarshal(m, b)
}
func (m *SearchProductsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchProductsRequest.Marshal(b, m, deterministic)
}
func (m *SearchProductsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchProductsRequest.Merge(m, src)
}
func (m *SearchProductsRequest) XXX_Size() int {
	return xxx_messageInfo_SearchProductsRequest.Size(m)
}
func (m *SearchProductsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchProductsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchProductsRequest proto.InternalMessageInfo

func (m *SearchProductsRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

type SearchProductsResponse struct {
	Results              []*Product `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *SearchProductsResponse) Reset()         { *m = SearchProductsResponse{} }
func (m *SearchProductsResponse) String() string { return proto.CompactTextString(m) }
func (*SearchProductsResponse) ProtoMessage()    {}
func (*SearchProductsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca53982754088a9d, []int{12}
}

func (m *SearchProductsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchProductsResponse.Unmarshal(m, b)
}
func (m *SearchProductsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchProductsResponse.Marshal(b, m, deterministic)
}
func (m *SearchProductsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchProductsResponse.Merge(m, src)
}
func (m *SearchProductsResponse) XXX_Size() int {
	return xxx_messageInfo_SearchProductsResponse.Size(m)
}
func (m *SearchProductsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchProductsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SearchProductsResponse proto.InternalMessageInfo

func (m *SearchProductsResponse) GetResults() []*Product {
	if m != nil {
		return m.Results
	}
	return nil
}

type GetQuoteRequest struct {
	Address              *Address    `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Items                []*CartItem `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GetQuoteRequest) Reset()         { *m = GetQuoteRequest{} }
func (m *GetQuoteRequest) String() string { return proto.CompactTextString(m) }
func (*GetQuoteRequest) ProtoMessage()    {}
func (*GetQuoteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca53982754088a9d, []int{13}
}

func (m *GetQuoteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetQuoteRequest.Unmarshal(m, b)
}
func (m *GetQuoteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetQuoteRequest.Marshal(b, m, deterministic)
}
func (m *GetQuoteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetQuoteRequest.Merge(m, src)
}
func (m *GetQuoteRequest) XXX_Size() int {
	return xxx_messageInfo_GetQuoteRequest.Size(m)
}
func (m *GetQuoteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetQuoteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetQuoteRequest proto.InternalMessageInfo

func (m *GetQuoteRequest) GetAddress() *Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *GetQuoteRequest) GetItems() []*CartItem {
	if m != nil {
		return m.Items
	}
	return nil
}

type GetQuoteResponse struct {
	CostUsd              *Money   `protobuf:"bytes,1,opt,name=cost_usd,json=costUsd,proto3" json:"cost_usd,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetQuoteResponse) Reset()         { *m = GetQuoteResponse{} }
func (m *GetQuoteResponse) String() string { return proto.CompactTextString(m) }
func (*GetQuoteResponse) ProtoMessage()    {}
func (*GetQuoteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca53982754088a9d, []int{14}
}

func (m *GetQuoteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetQuoteResponse.Unmarshal(m, b)
}
func (m *GetQuoteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetQuoteResponse.Marshal(b, m, deterministic)
}
func (m *GetQuoteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetQuoteResponse.Merge(m, src)
}
func (m *GetQuoteResponse) XXX_Size() int {
	return xxx_messageInfo_GetQuoteResponse.Size(m)
}
func (m *GetQuoteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetQuoteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetQuoteResponse proto.InternalMessageInfo

func (m *GetQuoteResponse) GetCostUsd() *Money {
	if m != nil {
		return m.CostUsd
	}
	return nil
}

type ShipOrderRequest struct {
	Address              *Address    `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Items                []*CartItem `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ShipOrderRequest) Reset()         { *m = ShipOrderRequest{} }
func (m *ShipOrderRequest) String() string { return proto.CompactTextString(m) }
func (*ShipOrderRequest) ProtoMessage()    {}
func (*ShipOrderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca53982754088a9d, []int{15}
}

func (m *ShipOrderRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShipOrderRequest.Unmarshal(m, b)
}
func (m *ShipOrderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShipOrderRequest.Marshal(b, m, deterministic)
}
func (m *ShipOrderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShipOrderRequest.Merge(m, src)
}
func (m *ShipOrderRequest) XXX_Size() int {
	return xxx_messageInfo_ShipOrderRequest.Size(m)
}
func (m *ShipOrderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ShipOrderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ShipOrderRequest proto.InternalMessageInfo

func (m *ShipOrderRequest) GetAddress() *Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *ShipOrderRequest) GetItems() []*CartItem {
	if m != nil {
		return m.Items
	}
	return nil
}

type ShipOrderResponse struct {
	TrackingId           string   `protobuf:"bytes,1,opt,name=tracking_id,json=trackingId,proto3" json:"tracking_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShipOrderResponse) Reset()         { *m = ShipOrderResponse{} }
func (m *ShipOrderResponse) String() string { return proto.CompactTextString(m) }
func (*ShipOrderResponse) ProtoMessage()    {}
func (*ShipOrderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca53982754088a9d, []int{16}
}

func (m *ShipOrderResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShipOrderResponse.Unmarshal(m, b)
}
func (m *ShipOrderResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShipOrderResponse.Marshal(b, m, deterministic)
}
func (m *ShipOrderResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShipOrderResponse.Merge(m, src)
}
func (m *ShipOrderResponse) XXX_Size() int {
	return xxx_messageInfo_ShipOrderResponse.Size(m)
}
func (m *ShipOrderResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ShipOrderResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ShipOrderResponse proto.InternalMessageInfo

func (m *ShipOrderResponse) GetTrackingId() string {
	if m != nil {
		return m.TrackingId
	}
	return ""
}

type Address struct {
	StreetAddress        string   `protobuf:"bytes,1,opt,name=street_address,json=streetAddress,proto3" json:"street_address,omitempty"`
	City                 string   `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty"`
	State                string   `protobuf:"bytes,3,opt,name=state,proto3" json:"state,omitempty"`
	Country              string   `protobuf:"bytes,4,opt,name=country,proto3" json:"country,omitempty"`
	ZipCode              int32    `protobuf:"varint,5,opt,name=zip_code,json=zipCode,proto3" json:"zip_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Address) Reset()         { *m = Address{} }
func (m *Address) String() string { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()    {}
func (*Address) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca53982754088a9d, []int{17}
}

func (m *Address) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Address.Unmarshal(m, b)
}
func (m *Address) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Address.Marshal(b, m, deterministic)
}
func (m *Address) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Address.Merge(m, src)
}
func (m *Address) XXX_Size() int {
	return xxx_messageInfo_Address.Size(m)
}
func (m *Address) XXX_DiscardUnknown() {
	xxx_messageInfo_Address.DiscardUnknown(m)
}

var xxx_messageInfo_Address proto.InternalMessageInfo

func (m *Address) GetStreetAddress() string {
	if m != nil {
		return m.StreetAddress
	}
	return ""
}

func (m *Address) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *Address) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *Address) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *Address) GetZipCode() int32 {
	if m != nil {
		return m.ZipCode
	}
	return 0
}

// Represents an amount of money with its currency type.
type Money struct {
	// The 3-letter currency code defined in ISO 4217.
	CurrencyCode string `protobuf:"bytes,1,opt,name=currency_code,json=currencyCode,proto3" json:"currency_code,omitempty"`
	// The whole units of the amount.
	// For example if `currencyCode` is `"USD"`, then 1 unit is one US dollar.
	Units int64 `protobuf:"varint,2,opt,name=units,proto3" json:"units,omitempty"`
	// Number of nano (10^-9) units of the amount.
	// The value must be between -999,999,999 and +999,999,999 inclusive.
	// If `units` is positive, `nanos` must be positive or zero.
	// If `units` is zero, `nanos` can be positive, zero, or negative.
	// If `units` is negative, `nanos` must be negative or zero.
	// For example $-1.75 is represented as `units`=-1 and `nanos`=-750,000,000.
	Nanos                int32    `protobuf:"varint,3,opt,name=nanos,proto3" json:"nanos,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Money) Reset()         { *m = Money{} }
func (m *Money) String() string { return proto.CompactTextString(m) }
func (*Money) ProtoMessage()    {}
func (*Money) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca53982754088a9d, []int{18}
}

func (m *Money) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Money.Unmarshal(m, b)
}
func (m *Money) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Money.Marshal(b, m, deterministic)
}
func (m *Money) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Money.Merge(m, src)
}
func (m *Money) XXX_Size() int {
	return xxx_messageInfo_Money.Size(m)
}
func (m *Money) XXX_DiscardUnknown() {
	xxx_messageInfo_Money.DiscardUnknown(m)
}

var xxx_messageInfo_Money proto.InternalMessageInfo

func (m *Money) GetCurrencyCode() string {
	if m != nil {
		return m.CurrencyCode
	}
	return ""
}

func (m *Money) GetUnits() int64 {
	if m != nil {
		return m.Units
	}
	return 0
}

func (m *Money) GetNanos() int32 {
	if m != nil {
		return m.Nanos
	}
	return 0
}

type GetSupportedCurrenciesResponse struct {
	// The 3-letter currency code defined in ISO 4217.
	CurrencyCodes        []string `protobuf:"bytes,1,rep,name=currency_codes,json=currencyCodes,proto3" json:"currency_codes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetSupportedCurrenciesResponse) Reset()         { *m = GetSupportedCurrenciesResponse{} }
func (m *GetSupportedCurrenciesResponse) String() string { return proto.CompactTextString(m) }
func (*GetSupportedCurrenciesResponse) ProtoMessage()    {}
func (*GetSupportedCurrenciesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca53982754088a9d, []int{19}
}

func (m *GetSupportedCurrenciesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSupportedCurrenciesResponse.Unmarshal(m, b)
}
func (m *GetSupportedCurrenciesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSupportedCurrenciesResponse.Marshal(b, m, deterministic)
}
func (m *GetSupportedCurrenciesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSupportedCurrenciesResponse.Merge(m, src)
}
func (m *GetSupportedCurrenciesResponse) XXX_Size() int {
	return xxx_messageInfo_GetSupportedCurrenciesResponse.Size(m)
}
func (m *GetSupportedCurrenciesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSupportedCurrenciesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetSupportedCurrenciesResponse proto.InternalMessageInfo

func (m *GetSupportedCurrenciesResponse) GetCurrencyCodes() []string {
	if m != nil {
		return m.CurrencyCodes
	}
	return nil
}

type CurrencyConversionRequest struct {
	From *Money `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	// The 3-letter currency code defined in ISO 4217.
	ToCode               string   `protobuf:"bytes,2,opt,name=to_code,json=toCode,proto3" json:"to_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CurrencyConversionRequest) Reset()         { *m = CurrencyConversionRequest{} }
func (m *CurrencyConversionRequest) String() string { return proto.CompactTextString(m) }
func (*CurrencyConversionRequest) ProtoMessage()    {}
func (*CurrencyConversionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca53982754088a9d, []int{20}
}

func (m *CurrencyConversionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CurrencyConversionRequest.Unmarshal(m, b)
}
func (m *CurrencyConversionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CurrencyConversionRequest.Marshal(b, m, deterministic)
}
func (m *CurrencyConversionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CurrencyConversionRequest.Merge(m, src)
}
func (m *CurrencyConversionRequest) XXX_Size() int {
	return xxx_messageInfo_CurrencyConversionRequest.Size(m)
}
func (m *CurrencyConversionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CurrencyConversionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CurrencyConversionRequest proto.InternalMessageInfo

func (m *CurrencyConversionRequest) GetFrom() *Money {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *CurrencyConversionRequest) GetToCode() string {
	if m != nil {
		return m.ToCode
	}
	return ""
}

type CreditCardInfo struct {
	CreditCardNumber          string   `protobuf:"bytes,1,opt,name=credit_card_number,json=creditCardNumber,proto3" json:"credit_card_number,omitempty"`
	CreditCardCvv             int32    `protobuf:"varint,2,opt,name=credit_card_cvv,json=creditCardCvv,proto3" json:"credit_card_cvv,omitempty"`
	CreditCardExpirationYear  int32    `protobuf:"varint,3,opt,name=credit_card_expiration_year,json=creditCardExpirationYear,proto3" json:"credit_card_expiration_year,omitempty"`
	CreditCardExpirationMonth int32    `protobuf:"varint,4,opt,name=credit_card_expiration_month,json=creditCardExpirationMonth,proto3" json:"credit_card_expiration_month,omitempty"`
	XXX_NoUnkeyedLiteral      struct{} `json:"-"`
	XXX_unrecognized          []byte   `json:"-"`
	XXX_sizecache             int32    `json:"-"`
}

func (m *CreditCardInfo) Reset()         { *m = CreditCardInfo{} }
func (m *CreditCardInfo) String() string { return proto.CompactTextString(m) }
func (*CreditCardInfo) ProtoMessage()    {}
func (*CreditCardInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca53982754088a9d, []int{21}
}

func (m *CreditCardInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreditCardInfo.Unmarshal(m, b)
}
func (m *CreditCardInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreditCardInfo.Marshal(b, m, deterministic)
}
func (m *CreditCardInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreditCardInfo.Merge(m, src)
}
func (m *CreditCardInfo) XXX_Size() int {
	return xxx_messageInfo_CreditCardInfo.Size(m)
}
func (m *CreditCardInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_CreditCardInfo.DiscardUnknown(m)
}

var xxx_messageInfo_CreditCardInfo proto.InternalMessageInfo

func (m *CreditCardInfo) GetCreditCardNumber() string {
	if m != nil {
		return m.CreditCardNumber
	}
	return ""
}

func (m *CreditCardInfo) GetCreditCardCvv() int32 {
	if m != nil {
		return m.CreditCardCvv
	}
	return 0
}

func (m *CreditCardInfo) GetCreditCardExpirationYear() int32 {
	if m != nil {
		return m.CreditCardExpirationYear
	}
	return 0
}

func (m *CreditCardInfo) GetCreditCardExpirationMonth() int32 {
	if m != nil {
		return m.CreditCardExpirationMonth
	}
	return 0
}

type ChargeRequest struct {
	Amount               *Money          `protobuf:"bytes,1,opt,name=amount,proto3" json:"amount,omitempty"`
	CreditCard           *CreditCardInfo `protobuf:"bytes,2,opt,name=credit_card,json=creditCard,proto3" json:"credit_card,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ChargeRequest) Reset()         { *m = ChargeRequest{} }
func (m *ChargeRequest) String() string { return proto.CompactTextString(m) }
func (*ChargeRequest) ProtoMessage()    {}
func (*ChargeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca53982754088a9d, []int{22}
}

func (m *ChargeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChargeRequest.Unmarshal(m, b)
}
func (m *ChargeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChargeRequest.Marshal(b, m, deterministic)
}
func (m *ChargeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChargeRequest.Merge(m, src)
}
func (m *ChargeRequest) XXX_Size() int {
	return xxx_messageInfo_ChargeRequest.Size(m)
}
func (m *ChargeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ChargeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ChargeRequest proto.InternalMessageInfo

func (m *ChargeRequest) GetAmount() *Money {
	if m != nil {
		return m.Amount
	}
	return nil
}

func (m *ChargeRequest) GetCreditCard() *CreditCardInfo {
	if m != nil {
		return m.CreditCard
	}
	return nil
}

type ChargeResponse struct {
	TransactionId        string   `protobuf:"bytes,1,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChargeResponse) Reset()         { *m = ChargeResponse{} }
func (m *ChargeResponse) String() string { return proto.CompactTextString(m) }
func (*ChargeResponse) ProtoMessage()    {}
func (*ChargeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca53982754088a9d, []int{23}
}

func (m *ChargeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChargeResponse.Unmarshal(m, b)
}
func (m *ChargeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChargeResponse.Marshal(b, m, deterministic)
}
func (m *ChargeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChargeResponse.Merge(m, src)
}
func (m *ChargeResponse) XXX_Size() int {
	return xxx_messageInfo_ChargeResponse.Size(m)
}
func (m *ChargeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ChargeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ChargeResponse proto.InternalMessageInfo

func (m *ChargeResponse) GetTransactionId() string {
	if m != nil {
		return m.TransactionId
	}
	return ""
}

type OrderItem struct {
	Item                 *CartItem `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
	Cost                 *Money    `protobuf:"bytes,2,opt,name=cost,proto3" json:"cost,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *OrderItem) Reset()         { *m = OrderItem{} }
func (m *OrderItem) String() string { return proto.CompactTextString(m) }
func (*OrderItem) ProtoMessage()    {}
func (*OrderItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca53982754088a9d, []int{24}
}

func (m *OrderItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderItem.Unmarshal(m, b)
}
func (m *OrderItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderItem.Marshal(b, m, deterministic)
}
func (m *OrderItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderItem.Merge(m, src)
}
func (m *OrderItem) XXX_Size() int {
	return xxx_messageInfo_OrderItem.Size(m)
}
func (m *OrderItem) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderItem.DiscardUnknown(m)
}

var xxx_messageInfo_OrderItem proto.InternalMessageInfo

func (m *OrderItem) GetItem() *CartItem {
	if m != nil {
		return m.Item
	}
	return nil
}

func (m *OrderItem) GetCost() *Money {
	if m != nil {
		return m.Cost
	}
	return nil
}

type OrderResult struct {
	OrderId              string       `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	ShippingTrackingId   string       `protobuf:"bytes,2,opt,name=shipping_tracking_id,json=shippingTrackingId,proto3" json:"shipping_tracking_id,omitempty"`
	ShippingCost         *Money       `protobuf:"bytes,3,opt,name=shipping_cost,json=shippingCost,proto3" json:"shipping_cost,omitempty"`
	ShippingAddress      *Address     `protobuf:"bytes,4,opt,name=shipping_address,json=shippingAddress,proto3" json:"shipping_address,omitempty"`
	Items                []*OrderItem `protobuf:"bytes,5,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *OrderResult) Reset()         { *m = OrderResult{} }
func (m *OrderResult) String() string { return proto.CompactTextString(m) }
func (*OrderResult) ProtoMessage()    {}
func (*OrderResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca53982754088a9d, []int{25}
}

func (m *OrderResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderResult.Unmarshal(m, b)
}
func (m *OrderResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderResult.Marshal(b, m, deterministic)
}
func (m *OrderResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderResult.Merge(m, src)
}
func (m *OrderResult) XXX_Size() int {
	return xxx_messageInfo_OrderResult.Size(m)
}
func (m *OrderResult) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderResult.DiscardUnknown(m)
}

var xxx_messageInfo_OrderResult proto.InternalMessageInfo

func (m *OrderResult) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

func (m *OrderResult) GetShippingTrackingId() string {
	if m != nil {
		return m.ShippingTrackingId
	}
	return ""
}

func (m *OrderResult) GetShippingCost() *Money {
	if m != nil {
		return m.ShippingCost
	}
	return nil
}

func (m *OrderResult) GetShippingAddress() *Address {
	if m != nil {
		return m.ShippingAddress
	}
	return nil
}

func (m *OrderResult) GetItems() []*OrderItem {
	if m != nil {
		return m.Items
	}
	return nil
}

type SendOrderConfirmationRequest struct {
	Email                string       `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Order                *OrderResult `protobuf:"bytes,2,opt,name=order,proto3" json:"order,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *SendOrderConfirmationRequest) Reset()         { *m = SendOrderConfirmationRequest{} }
func (m *SendOrderConfirmationRequest) String() string { return proto.CompactTextString(m) }
func (*SendOrderConfirmationRequest) ProtoMessage()    {}
func (*SendOrderConfirmationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca53982754088a9d, []int{26}
}

func (m *SendOrderConfirmationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendOrderConfirmationRequest.Unmarshal(m, b)
}
func (m *SendOrderConfirmationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendOrderConfirmationRequest.Marshal(b, m, deterministic)
}
func (m *SendOrderConfirmationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendOrderConfirmationRequest.Merge(m, src)
}
func (m *SendOrderConfirmationRequest) XXX_Size() int {
	return xxx_messageInfo_SendOrderConfirmationRequest.Size(m)
}
func (m *SendOrderConfirmationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SendOrderConfirmationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SendOrderConfirmationRequest proto.InternalMessageInfo

func (m *SendOrderConfirmationRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *SendOrderConfirmationRequest) GetOrder() *OrderResult {
	if m != nil {
		return m.Order
	}
	return nil
}

type PlaceOrderRequest struct {
	UserId               string          `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserCurrency         string          `protobuf:"bytes,2,opt,name=user_currency,json=userCurrency,proto3" json:"user_currency,omitempty"`
	Address              *Address        `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Email                string          `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	CreditCard           *CreditCardInfo `protobuf:"bytes,6,opt,name=credit_card,json=creditCard,proto3" json:"credit_card,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *PlaceOrderRequest) Reset()         { *m = PlaceOrderRequest{} }
func (m *PlaceOrderRequest) String() string { return proto.CompactTextString(m) }
func (*PlaceOrderRequest) ProtoMessage()    {}
func (*PlaceOrderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca53982754088a9d, []int{27}
}

func (m *PlaceOrderRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlaceOrderRequest.Unmarshal(m, b)
}
func (m *PlaceOrderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlaceOrderRequest.Marshal(b, m, deterministic)
}
func (m *PlaceOrderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlaceOrderRequest.Merge(m, src)
}
func (m *PlaceOrderRequest) XXX_Size() int {
	return xxx_messageInfo_PlaceOrderRequest.Size(m)
}
func (m *PlaceOrderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PlaceOrderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PlaceOrderRequest proto.InternalMessageInfo

func (m *PlaceOrderRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *PlaceOrderRequest) GetUserCurrency() string {
	if m != nil {
		return m.UserCurrency
	}
	return ""
}

func (m *PlaceOrderRequest) GetAddress() *Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *PlaceOrderRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *PlaceOrderRequest) GetCreditCard() *CreditCardInfo {
	if m != nil {
		return m.CreditCard
	}
	return nil
}

type PlaceOrderResponse struct {
	Order                *OrderResult `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *PlaceOrderResponse) Reset()         { *m = PlaceOrderResponse{} }
func (m *PlaceOrderResponse) String() string { return proto.CompactTextString(m) }
func (*PlaceOrderResponse) ProtoMessage()    {}
func (*PlaceOrderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca53982754088a9d, []int{28}
}

func (m *PlaceOrderResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlaceOrderResponse.Unmarshal(m, b)
}
func (m *PlaceOrderResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlaceOrderResponse.Marshal(b, m, deterministic)
}
func (m *PlaceOrderResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlaceOrderResponse.Merge(m, src)
}
func (m *PlaceOrderResponse) XXX_Size() int {
	return xxx_messageInfo_PlaceOrderResponse.Size(m)
}
func (m *PlaceOrderResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PlaceOrderResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PlaceOrderResponse proto.InternalMessageInfo

func (m *PlaceOrderResponse) GetOrder() *OrderResult {
	if m != nil {
		return m.Order
	}
	return nil
}

type AdRequest struct {
	// List of important key words from the current page describing the context.
	ContextKeys          []string `protobuf:"bytes,1,rep,name=context_keys,json=contextKeys,proto3" json:"context_keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdRequest) Reset()         { *m = AdRequest{} }
func (m *AdRequest) String() string { return proto.CompactTextString(m) }
func (*AdRequest) ProtoMessage()    {}
func (*AdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca53982754088a9d, []int{29}
}

func (m *AdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdRequest.Unmarshal(m, b)
}
func (m *AdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdRequest.Marshal(b, m, deterministic)
}
func (m *AdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdRequest.Merge(m, src)
}
func (m *AdRequest) XXX_Size() int {
	return xxx_messageInfo_AdRequest.Size(m)
}
func (m *AdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AdRequest proto.InternalMessageInfo

func (m *AdRequest) GetContextKeys() []string {
	if m != nil {
		return m.ContextKeys
	}
	return nil
}

type AdResponse struct {
	Ads                  []*Ad    `protobuf:"bytes,1,rep,name=ads,proto3" json:"ads,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdResponse) Reset()         { *m = AdResponse{} }
func (m *AdResponse) String() string { return proto.CompactTextString(m) }
func (*AdResponse) ProtoMessage()    {}
func (*AdResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca53982754088a9d, []int{30}
}

func (m *AdResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdResponse.Unmarshal(m, b)
}
func (m *AdResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdResponse.Marshal(b, m, deterministic)
}
func (m *AdResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdResponse.Merge(m, src)
}
func (m *AdResponse) XXX_Size() int {
	return xxx_messageInfo_AdResponse.Size(m)
}
func (m *AdResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AdResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AdResponse proto.InternalMessageInfo

func (m *AdResponse) GetAds() []*Ad {
	if m != nil {
		return m.Ads
	}
	return nil
}

type Ad struct {
	// url to redirect to when an ad is clicked.
	RedirectUrl string `protobuf:"bytes,1,opt,name=redirect_url,json=redirectUrl,proto3" json:"redirect_url,omitempty"`
	// short advertisement text to display.
	Text                 string   `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ad) Reset()         { *m = Ad{} }
func (m *Ad) String() string { return proto.CompactTextString(m) }
func (*Ad) ProtoMessage()    {}
func (*Ad) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca53982754088a9d, []int{31}
}

func (m *Ad) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ad.Unmarshal(m, b)
}
func (m *Ad) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ad.Marshal(b, m, deterministic)
}
func (m *Ad) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ad.Merge(m, src)
}
func (m *Ad) XXX_Size() int {
	return xxx_messageInfo_Ad.Size(m)
}
func (m *Ad) XXX_DiscardUnknown() {
	xxx_messageInfo_Ad.DiscardUnknown(m)
}

var xxx_messageInfo_Ad proto.InternalMessageInfo

func (m *Ad) GetRedirectUrl() string {
	if m != nil {
		return m.RedirectUrl
	}
	return ""
}

func (m *Ad) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func init() {
	proto.RegisterType((*CartItem)(nil), "hipstershop.CartItem")
	proto.RegisterType((*AddItemRequest)(nil), "hipstershop.AddItemRequest")
	proto.RegisterType((*EmptyCartRequest)(nil), "hipstershop.EmptyCartRequest")
	proto.RegisterType((*GetCartRequest)(nil), "hipstershop.GetCartRequest")
	proto.RegisterType((*Cart)(nil), "hipstershop.Cart")
	proto.RegisterType((*Empty)(nil), "hipstershop.Empty")
	proto.RegisterType((*ListRecommendationsRequest)(nil), "hipstershop.ListRecommendationsRequest")
	proto.RegisterType((*ListRecommendationsResponse)(nil), "hipstershop.ListRecommendationsResponse")
	proto.RegisterType((*Product)(nil), "hipstershop.Product")
	proto.RegisterType((*ListProductsResponse)(nil), "hipstershop.ListProductsResponse")
	proto.RegisterType((*GetProductRequest)(nil), "hipstershop.GetProductRequest")
	proto.RegisterType((*SearchProductsRequest)(nil), "hipstershop.SearchProductsRequest")
	proto.RegisterType((*SearchProductsResponse)(nil), "hipstershop.SearchProductsResponse")
	proto.RegisterType((*GetQuoteRequest)(nil), "hipstershop.GetQuoteRequest")
	proto.RegisterType((*GetQuoteResponse)(nil), "hipstershop.GetQuoteResponse")
	proto.RegisterType((*ShipOrderRequest)(nil), "hipstershop.ShipOrderRequest")
	proto.RegisterType((*ShipOrderResponse)(nil), "hipstershop.ShipOrderResponse")
	proto.RegisterType((*Address)(nil), "hipstershop.Address")
	proto.RegisterType((*Money)(nil), "hipstershop.Money")
	proto.RegisterType((*GetSupportedCurrenciesResponse)(nil), "hipstershop.GetSupportedCurrenciesResponse")
	proto.RegisterType((*CurrencyConversionRequest)(nil), "hipstershop.CurrencyConversionRequest")
	proto.RegisterType((*CreditCardInfo)(nil), "hipstershop.CreditCardInfo")
	proto.RegisterType((*ChargeRequest)(nil), "hipstershop.ChargeRequest")
	proto.RegisterType((*ChargeResponse)(nil), "hipstershop.ChargeResponse")
	proto.RegisterType((*OrderItem)(nil), "hipstershop.OrderItem")
	proto.RegisterType((*OrderResult)(nil), "hipstershop.OrderResult")
	proto.RegisterType((*SendOrderConfirmationRequest)(nil), "hipstershop.SendOrderConfirmationRequest")
	proto.RegisterType((*PlaceOrderRequest)(nil), "hipstershop.PlaceOrderRequest")
	proto.RegisterType((*PlaceOrderResponse)(nil), "hipstershop.PlaceOrderResponse")
	proto.RegisterType((*AdRequest)(nil), "hipstershop.AdRequest")
	proto.RegisterType((*AdResponse)(nil), "hipstershop.AdResponse")
	proto.RegisterType((*Ad)(nil), "hipstershop.Ad")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CartServiceClient is the client API for CartService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CartServiceClient interface {
	AddItem(ctx context.Context, in *AddItemRequest, opts ...grpc.CallOption) (*Empty, error)
	GetCart(ctx context.Context, in *GetCartRequest, opts ...grpc.CallOption) (*Cart, error)
	EmptyCart(ctx context.Context, in *EmptyCartRequest, opts ...grpc.CallOption) (*Empty, error)
}

type cartServiceClient struct {
	cc *grpc.ClientConn
}

func NewCartServiceClient(cc *grpc.ClientConn) CartServiceClient {
	return &cartServiceClient{cc}
}

func (c *cartServiceClient) AddItem(ctx context.Context, in *AddItemRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/hipstershop.CartService/AddItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) GetCart(ctx context.Context, in *GetCartRequest, opts ...grpc.CallOption) (*Cart, error) {
	out := new(Cart)
	err := c.cc.Invoke(ctx, "/hipstershop.CartService/GetCart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) EmptyCart(ctx context.Context, in *EmptyCartRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/hipstershop.CartService/EmptyCart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CartServiceServer is the server API for CartService service.
type CartServiceServer interface {
	AddItem(context.Context, *AddItemRequest) (*Empty, error)
	GetCart(context.Context, *GetCartRequest) (*Cart, error)
	EmptyCart(context.Context, *EmptyCartRequest) (*Empty, error)
}

func RegisterCartServiceServer(s *grpc.Server, srv CartServiceServer) {
	s.RegisterService(&_CartService_serviceDesc, srv)
}

func _CartService_AddItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).AddItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hipstershop.CartService/AddItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).AddItem(ctx, req.(*AddItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_GetCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).GetCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hipstershop.CartService/GetCart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).GetCart(ctx, req.(*GetCartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_EmptyCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyCartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).EmptyCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hipstershop.CartService/EmptyCart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).EmptyCart(ctx, req.(*EmptyCartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CartService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hipstershop.CartService",
	HandlerType: (*CartServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddItem",
			Handler:    _CartService_AddItem_Handler,
		},
		{
			MethodName: "GetCart",
			Handler:    _CartService_GetCart_Handler,
		},
		{
			MethodName: "EmptyCart",
			Handler:    _CartService_EmptyCart_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "demo.proto",
}

// RecommendationServiceClient is the client API for RecommendationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RecommendationServiceClient interface {
	ListRecommendations(ctx context.Context, in *ListRecommendationsRequest, opts ...grpc.CallOption) (*ListRecommendationsResponse, error)
}

type recommendationServiceClient struct {
	cc *grpc.ClientConn
}

func NewRecommendationServiceClient(cc *grpc.ClientConn) RecommendationServiceClient {
	return &recommendationServiceClient{cc}
}

func (c *recommendationServiceClient) ListRecommendations(ctx context.Context, in *ListRecommendationsRequest, opts ...grpc.CallOption) (*ListRecommendationsResponse, error) {
	out := new(ListRecommendationsResponse)
	err := c.cc.Invoke(ctx, "/hipstershop.RecommendationService/ListRecommendations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RecommendationServiceServer is the server API for RecommendationService service.
type RecommendationServiceServer interface {
	ListRecommendations(context.Context, *ListRecommendationsRequest) (*ListRecommendationsResponse, error)
}

func RegisterRecommendationServiceServer(s *grpc.Server, srv RecommendationServiceServer) {
	s.RegisterService(&_RecommendationService_serviceDesc, srv)
}

func _RecommendationService_ListRecommendations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRecommendationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecommendationServiceServer).ListRecommendations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hipstershop.RecommendationService/ListRecommendations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecommendationServiceServer).ListRecommendations(ctx, req.(*ListRecommendationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RecommendationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hipstershop.RecommendationService",
	HandlerType: (*RecommendationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListRecommendations",
			Handler:    _RecommendationService_ListRecommendations_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "demo.proto",
}

// ProductCatalogServiceClient is the client API for ProductCatalogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProductCatalogServiceClient interface {
	ListProducts(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListProductsResponse, error)
	GetProduct(ctx context.Context, in *GetProductRequest, opts ...grpc.CallOption) (*Product, error)
	SearchProducts(ctx context.Context, in *SearchProductsRequest, opts ...grpc.CallOption) (*SearchProductsResponse, error)
}

type productCatalogServiceClient struct {
	cc *grpc.ClientConn
}

func NewProductCatalogServiceClient(cc *grpc.ClientConn) ProductCatalogServiceClient {
	return &productCatalogServiceClient{cc}
}

func (c *productCatalogServiceClient) ListProducts(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListProductsResponse, error) {
	out := new(ListProductsResponse)
	err := c.cc.Invoke(ctx, "/hipstershop.ProductCatalogService/ListProducts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productCatalogServiceClient) GetProduct(ctx context.Context, in *GetProductRequest, opts ...grpc.CallOption) (*Product, error) {
	out := new(Product)
	err := c.cc.Invoke(ctx, "/hipstershop.ProductCatalogService/GetProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productCatalogServiceClient) SearchProducts(ctx context.Context, in *SearchProductsRequest, opts ...grpc.CallOption) (*SearchProductsResponse, error) {
	out := new(SearchProductsResponse)
	err := c.cc.Invoke(ctx, "/hipstershop.ProductCatalogService/SearchProducts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductCatalogServiceServer is the server API for ProductCatalogService service.
type ProductCatalogServiceServer interface {
	ListProducts(context.Context, *Empty) (*ListProductsResponse, error)
	GetProduct(context.Context, *GetProductRequest) (*Product, error)
	SearchProducts(context.Context, *SearchProductsRequest) (*SearchProductsResponse, error)
}

func RegisterProductCatalogServiceServer(s *grpc.Server, srv ProductCatalogServiceServer) {
	s.RegisterService(&_ProductCatalogService_serviceDesc, srv)
}

func _ProductCatalogService_ListProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductCatalogServiceServer).ListProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hipstershop.ProductCatalogService/ListProducts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductCatalogServiceServer).ListProducts(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductCatalogService_GetProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductCatalogServiceServer).GetProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hipstershop.ProductCatalogService/GetProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductCatalogServiceServer).GetProduct(ctx, req.(*GetProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductCatalogService_SearchProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchProductsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductCatalogServiceServer).SearchProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hipstershop.ProductCatalogService/SearchProducts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductCatalogServiceServer).SearchProducts(ctx, req.(*SearchProductsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProductCatalogService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hipstershop.ProductCatalogService",
	HandlerType: (*ProductCatalogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListProducts",
			Handler:    _ProductCatalogService_ListProducts_Handler,
		},
		{
			MethodName: "GetProduct",
			Handler:    _ProductCatalogService_GetProduct_Handler,
		},
		{
			MethodName: "SearchProducts",
			Handler:    _ProductCatalogService_SearchProducts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "demo.proto",
}

// ShippingServiceClient is the client API for ShippingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ShippingServiceClient interface {
	GetQuote(ctx context.Context, in *GetQuoteRequest, opts ...grpc.CallOption) (*GetQuoteResponse, error)
	ShipOrder(ctx context.Context, in *ShipOrderRequest, opts ...grpc.CallOption) (*ShipOrderResponse, error)
}

type shippingServiceClient struct {
	cc *grpc.ClientConn
}

func NewShippingServiceClient(cc *grpc.ClientConn) ShippingServiceClient {
	return &shippingServiceClient{cc}
}

func (c *shippingServiceClient) GetQuote(ctx context.Context, in *GetQuoteRequest, opts ...grpc.CallOption) (*GetQuoteResponse, error) {
	out := new(GetQuoteResponse)
	err := c.cc.Invoke(ctx, "/hipstershop.ShippingService/GetQuote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shippingServiceClient) ShipOrder(ctx context.Context, in *ShipOrderRequest, opts ...grpc.CallOption) (*ShipOrderResponse, error) {
	out := new(ShipOrderResponse)
	err := c.cc.Invoke(ctx, "/hipstershop.ShippingService/ShipOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShippingServiceServer is the server API for ShippingService service.
type ShippingServiceServer interface {
	GetQuote(context.Context, *GetQuoteRequest) (*GetQuoteResponse, error)
	ShipOrder(context.Context, *ShipOrderRequest) (*ShipOrderResponse, error)
}

func RegisterShippingServiceServer(s *grpc.Server, srv ShippingServiceServer) {
	s.RegisterService(&_ShippingService_serviceDesc, srv)
}

func _ShippingService_GetQuote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetQuoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShippingServiceServer).GetQuote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hipstershop.ShippingService/GetQuote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShippingServiceServer).GetQuote(ctx, req.(*GetQuoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShippingService_ShipOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShipOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShippingServiceServer).ShipOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hipstershop.ShippingService/ShipOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShippingServiceServer).ShipOrder(ctx, req.(*ShipOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ShippingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hipstershop.ShippingService",
	HandlerType: (*ShippingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetQuote",
			Handler:    _ShippingService_GetQuote_Handler,
		},
		{
			MethodName: "ShipOrder",
			Handler:    _ShippingService_ShipOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "demo.proto",
}

// CurrencyServiceClient is the client API for CurrencyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CurrencyServiceClient interface {
	GetSupportedCurrencies(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetSupportedCurrenciesResponse, error)
	Convert(ctx context.Context, in *CurrencyConversionRequest, opts ...grpc.CallOption) (*Money, error)
}

type currencyServiceClient struct {
	cc *grpc.ClientConn
}

func NewCurrencyServiceClient(cc *grpc.ClientConn) CurrencyServiceClient {
	return &currencyServiceClient{cc}
}

func (c *currencyServiceClient) GetSupportedCurrencies(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetSupportedCurrenciesResponse, error) {
	out := new(GetSupportedCurrenciesResponse)
	err := c.cc.Invoke(ctx, "/hipstershop.CurrencyService/GetSupportedCurrencies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *currencyServiceClient) Convert(ctx context.Context, in *CurrencyConversionRequest, opts ...grpc.CallOption) (*Money, error) {
	out := new(Money)
	err := c.cc.Invoke(ctx, "/hipstershop.CurrencyService/Convert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CurrencyServiceServer is the server API for CurrencyService service.
type CurrencyServiceServer interface {
	GetSupportedCurrencies(context.Context, *Empty) (*GetSupportedCurrenciesResponse, error)
	Convert(context.Context, *CurrencyConversionRequest) (*Money, error)
}

func RegisterCurrencyServiceServer(s *grpc.Server, srv CurrencyServiceServer) {
	s.RegisterService(&_CurrencyService_serviceDesc, srv)
}

func _CurrencyService_GetSupportedCurrencies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurrencyServiceServer).GetSupportedCurrencies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hipstershop.CurrencyService/GetSupportedCurrencies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurrencyServiceServer).GetSupportedCurrencies(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CurrencyService_Convert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CurrencyConversionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurrencyServiceServer).Convert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hipstershop.CurrencyService/Convert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurrencyServiceServer).Convert(ctx, req.(*CurrencyConversionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CurrencyService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hipstershop.CurrencyService",
	HandlerType: (*CurrencyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSupportedCurrencies",
			Handler:    _CurrencyService_GetSupportedCurrencies_Handler,
		},
		{
			MethodName: "Convert",
			Handler:    _CurrencyService_Convert_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "demo.proto",
}

// PaymentServiceClient is the client API for PaymentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PaymentServiceClient interface {
	Charge(ctx context.Context, in *ChargeRequest, opts ...grpc.CallOption) (*ChargeResponse, error)
}

type paymentServiceClient struct {
	cc *grpc.ClientConn
}

func NewPaymentServiceClient(cc *grpc.ClientConn) PaymentServiceClient {
	return &paymentServiceClient{cc}
}

func (c *paymentServiceClient) Charge(ctx context.Context, in *ChargeRequest, opts ...grpc.CallOption) (*ChargeResponse, error) {
	out := new(ChargeResponse)
	err := c.cc.Invoke(ctx, "/hipstershop.PaymentService/Charge", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentServiceServer is the server API for PaymentService service.
type PaymentServiceServer interface {
	Charge(context.Context, *ChargeRequest) (*ChargeResponse, error)
}

func RegisterPaymentServiceServer(s *grpc.Server, srv PaymentServiceServer) {
	s.RegisterService(&_PaymentService_serviceDesc, srv)
}

func _PaymentService_Charge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChargeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).Charge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hipstershop.PaymentService/Charge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).Charge(ctx, req.(*ChargeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PaymentService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hipstershop.PaymentService",
	HandlerType: (*PaymentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Charge",
			Handler:    _PaymentService_Charge_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "demo.proto",
}

// EmailServiceClient is the client API for EmailService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EmailServiceClient interface {
	SendOrderConfirmation(ctx context.Context, in *SendOrderConfirmationRequest, opts ...grpc.CallOption) (*Empty, error)
}

type emailServiceClient struct {
	cc *grpc.ClientConn
}

func NewEmailServiceClient(cc *grpc.ClientConn) EmailServiceClient {
	return &emailServiceClient{cc}
}

func (c *emailServiceClient) SendOrderConfirmation(ctx context.Context, in *SendOrderConfirmationRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/hipstershop.EmailService/SendOrderConfirmation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EmailServiceServer is the server API for EmailService service.
type EmailServiceServer interface {
	SendOrderConfirmation(context.Context, *SendOrderConfirmationRequest) (*Empty, error)
}

func RegisterEmailServiceServer(s *grpc.Server, srv EmailServiceServer) {
	s.RegisterService(&_EmailService_serviceDesc, srv)
}

func _EmailService_SendOrderConfirmation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendOrderConfirmationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailServiceServer).SendOrderConfirmation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hipstershop.EmailService/SendOrderConfirmation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailServiceServer).SendOrderConfirmation(ctx, req.(*SendOrderConfirmationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EmailService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hipstershop.EmailService",
	HandlerType: (*EmailServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendOrderConfirmation",
			Handler:    _EmailService_SendOrderConfirmation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "demo.proto",
}

// CheckoutServiceClient is the client API for CheckoutService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CheckoutServiceClient interface {
	PlaceOrder(ctx context.Context, in *PlaceOrderRequest, opts ...grpc.CallOption) (*PlaceOrderResponse, error)
}

type checkoutServiceClient struct {
	cc *grpc.ClientConn
}

func NewCheckoutServiceClient(cc *grpc.ClientConn) CheckoutServiceClient {
	return &checkoutServiceClient{cc}
}

func (c *checkoutServiceClient) PlaceOrder(ctx context.Context, in *PlaceOrderRequest, opts ...grpc.CallOption) (*PlaceOrderResponse, error) {
	out := new(PlaceOrderResponse)
	err := c.cc.Invoke(ctx, "/hipstershop.CheckoutService/PlaceOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CheckoutServiceServer is the server API for CheckoutService service.
type CheckoutServiceServer interface {
	PlaceOrder(context.Context, *PlaceOrderRequest) (*PlaceOrderResponse, error)
}

func RegisterCheckoutServiceServer(s *grpc.Server, srv CheckoutServiceServer) {
	s.RegisterService(&_CheckoutService_serviceDesc, srv)
}

func _CheckoutService_PlaceOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlaceOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckoutServiceServer).PlaceOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hipstershop.CheckoutService/PlaceOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckoutServiceServer).PlaceOrder(ctx, req.(*PlaceOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CheckoutService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hipstershop.CheckoutService",
	HandlerType: (*CheckoutServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PlaceOrder",
			Handler:    _CheckoutService_PlaceOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "demo.proto",
}

// AdServiceClient is the client API for AdService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AdServiceClient interface {
	GetAds(ctx context.Context, in *AdRequest, opts ...grpc.CallOption) (*AdResponse, error)
}

type adServiceClient struct {
	cc *grpc.ClientConn
}

func NewAdServiceClient(cc *grpc.ClientConn) AdServiceClient {
	return &adServiceClient{cc}
}

func (c *adServiceClient) GetAds(ctx context.Context, in *AdRequest, opts ...grpc.CallOption) (*AdResponse, error) {
	out := new(AdResponse)
	err := c.cc.Invoke(ctx, "/hipstershop.AdService/GetAds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdServiceServer is the server API for AdService service.
type AdServiceServer interface {
	GetAds(context.Context, *AdRequest) (*AdResponse, error)
}

func RegisterAdServiceServer(s *grpc.Server, srv AdServiceServer) {
	s.RegisterService(&_AdService_serviceDesc, srv)
}

func _AdService_GetAds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdServiceServer).GetAds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hipstershop.AdService/GetAds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdServiceServer).GetAds(ctx, req.(*AdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AdService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hipstershop.AdService",
	HandlerType: (*AdServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAds",
			Handler:    _AdService_GetAds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "demo.proto",
}

func init() { proto.RegisterFile("demo.proto", fileDescriptor_ca53982754088a9d) }

var fileDescriptor_ca53982754088a9d = []byte{
	// 1500 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x58, 0xef, 0x72, 0x13, 0xb7,
	0x16, 0xcf, 0x26, 0xb1, 0x1d, 0x1f, 0xc7, 0x4e, 0xa2, 0x9b, 0x04, 0xb3, 0x81, 0x10, 0x94, 0x81,
	0x0b, 0x17, 0x08, 0x4c, 0xee, 0x9d, 0xe1, 0x03, 0xdc, 0xd2, 0x8c, 0xc9, 0x18, 0x4f, 0xa1, 0xd0,
	0x0d, 0xe9, 0xd0, 0xa1, 0x53, 0xcf, 0xb2, 0x12, 0xf1, 0x96, 0xec, 0x6a, 0x91, 0xb4, 0x19, 0xcc,
	0xc7, 0xf6, 0x01, 0xfa, 0x1e, 0x7d, 0x81, 0xce, 0xf4, 0x11, 0xfa, 0xbd, 0xaf, 0xd0, 0xe7, 0xe8,
	0x48, 0xbb, 0xda, 0x7f, 0xb1, 0x13, 0xf8, 0xd2, 0x6f, 0xab, 0xa3, 0x9f, 0xce, 0xf9, 0xe9, 0xe8,
	0xfc, 0xb3, 0x01, 0x08, 0x0d, 0xd8, 0x4e, 0xc4, 0x99, 0x64, 0xa8, 0x35, 0xf2, 0x23, 0x21, 0x29,
	0x17, 0x23, 0x16, 0xe1, 0x7d, 0x58, 0xe8, 0xb9, 0x5c, 0x0e, 0x24, 0x0d, 0xd0, 0x65, 0x80, 0x88,
	0x33, 0x12, 0x7b, 0x72, 0xe8, 0x93, 0xae, 0xb5, 0x65, 0xdd, 0x68, 0x3a, 0xcd, 0x54, 0x32, 0x20,
	0xc8, 0x86, 0x85, 0xf7, 0xb1, 0x1b, 0x4a, 0x5f, 0x8e, 0xbb, 0xb3, 0x5b, 0xd6, 0x8d, 0x9a, 0x93,
	0xad, 0xf1, 0x4b, 0xe8, 0xec, 0x11, 0xa2, 0xb4, 0x38, 0xf4, 0x7d, 0x4c, 0x85, 0x44, 0x17, 0xa0,
	0x11, 0x0b, 0xca, 0x73, 0x4d, 0x75, 0xb5, 0x1c, 0x10, 0x74, 0x13, 0xe6, 0x7d, 0x49, 0x03, 0xad,
	0xa2, 0xb5, 0xbb, 0xb6, 0x53, 0x60, 0xb3, 0x63, 0xa8, 0x38, 0x1a, 0x82, 0x6f, 0xc1, 0xf2, 0x7e,
	0x10, 0xc9, 0xb1, 0x12, 0x9f, 0xa7, 0x17, 0xdf, 0x84, 0x4e, 0x9f, 0xca, 0x4f, 0x82, 0x3e, 0x85,
	0x79, 0x85, 0x9b, 0xce, 0xf1, 0x16, 0xd4, 0x14, 0x01, 0xd1, 0x9d, 0xdd, 0x9a, 0x9b, 0x4e, 0x32,
	0xc1, 0xe0, 0x06, 0xd4, 0x34, 0x4b, 0xfc, 0x2d, 0xd8, 0x4f, 0x7d, 0x21, 0x1d, 0xea, 0xb1, 0x20,
	0xa0, 0x21, 0x71, 0xa5, 0xcf, 0x42, 0x71, 0xae, 0x43, 0xae, 0x40, 0x2b, 0x77, 0x7b, 0x62, 0xb2,
	0xe9, 0x40, 0xe6, 0x77, 0x81, 0xbf, 0x80, 0x8d, 0x89, 0x7a, 0x45, 0xc4, 0x42, 0x41, 0xab, 0xe7,
	0xad, 0x53, 0xe7, 0x7f, 0xb7, 0xa0, 0xf1, 0x22, 0x59, 0xa2, 0x0e, 0xcc, 0x66, 0x04, 0x66, 0x7d,
	0x82, 0x10, 0xcc, 0x87, 0x6e, 0x40, 0xf5, 0x6b, 0x34, 0x1d, 0xfd, 0x8d, 0xb6, 0xa0, 0x45, 0xa8,
	0xf0, 0xb8, 0x1f, 0x29, 0x43, 0xdd, 0x39, 0xbd, 0x55, 0x14, 0xa1, 0x2e, 0x34, 0x22, 0xdf, 0x93,
	0x31, 0xa7, 0xdd, 0x79, 0xbd, 0x6b, 0x96, 0xe8, 0x2e, 0x34, 0x23, 0xee, 0x7b, 0x74, 0x18, 0x0b,
	0xd2, 0xad, 0xe9, 0x27, 0x46, 0x25, 0xef, 0x3d, 0x63, 0x21, 0x1d, 0x3b, 0x0b, 0x1a, 0x74, 0x28,
	0x08, 0xda, 0x04, 0xf0, 0x5c, 0x49, 0x8f, 0x18, 0xf7, 0xa9, 0xe8, 0xd6, 0x13, 0xf2, 0xb9, 0x04,
	0x3f, 0x81, 0x55, 0x75, 0xf9, 0x94, 0x7f, 0x7e, 0xeb, 0x7b, 0xb0, 0x90, 0x5e, 0x31, 0xb9, 0x72,
	0x6b, 0x77, 0xb5, 0x64, 0x27, 0x3d, 0xe0, 0x64, 0x28, 0xbc, 0x0d, 0x2b, 0x7d, 0x6a, 0x14, 0x99,
	0x57, 0xa9, 0xf8, 0x03, 0xdf, 0x81, 0xb5, 0x03, 0xea, 0x72, 0x6f, 0x94, 0x1b, 0x4c, 0x80, 0xab,
	0x50, 0x7b, 0x1f, 0x53, 0x3e, 0x4e, 0xb1, 0xc9, 0x02, 0x3f, 0x81, 0xf5, 0x2a, 0x3c, 0xe5, 0xb7,
	0x03, 0x0d, 0x4e, 0x45, 0x7c, 0x7c, 0x0e, 0x3d, 0x03, 0xc2, 0x21, 0x2c, 0xf5, 0xa9, 0xfc, 0x26,
	0x66, 0x92, 0x1a, 0x93, 0x3b, 0xd0, 0x70, 0x09, 0xe1, 0x54, 0x08, 0x6d, 0xb4, 0xaa, 0x62, 0x2f,
	0xd9, 0x73, 0x0c, 0xe8, 0xf3, 0xa2, 0x76, 0x0f, 0x96, 0x73, 0x7b, 0x29, 0xe7, 0x3b, 0xb0, 0xe0,
	0x31, 0x21, 0xf5, 0xdb, 0x59, 0x53, 0xdf, 0xae, 0xa1, 0x30, 0x87, 0x82, 0x60, 0x06, 0xcb, 0x07,
	0x23, 0x3f, 0x7a, 0xce, 0x09, 0xe5, 0xff, 0x08, 0xe7, 0xff, 0xc1, 0x4a, 0xc1, 0x60, 0x1e, 0xfe,
	0x92, 0xbb, 0xde, 0x3b, 0x3f, 0x3c, 0xca, 0x73, 0x0b, 0x8c, 0x68, 0x40, 0xf0, 0x2f, 0x16, 0x34,
	0x52, 0xbb, 0xe8, 0x1a, 0x74, 0x84, 0xe4, 0x94, 0xca, 0x61, 0x91, 0x65, 0xd3, 0x69, 0x27, 0x52,
	0x03, 0x43, 0x30, 0xef, 0x99, 0x32, 0xd7, 0x74, 0xf4, 0xb7, 0x0a, 0x00, 0x21, 0x5d, 0x49, 0xd3,
	0x7c, 0x48, 0x16, 0x2a, 0x13, 0x3c, 0x16, 0x87, 0x92, 0x8f, 0x4d, 0x26, 0xa4, 0x4b, 0x74, 0x11,
	0x16, 0x3e, 0xfa, 0xd1, 0xd0, 0x63, 0x84, 0xea, 0x44, 0xa8, 0x39, 0x8d, 0x8f, 0x7e, 0xd4, 0x63,
	0x84, 0xe2, 0x57, 0x50, 0xd3, 0xae, 0x44, 0xdb, 0xd0, 0xf6, 0x62, 0xce, 0x69, 0xe8, 0x8d, 0x13,
	0x60, 0xc2, 0x66, 0xd1, 0x08, 0x15, 0x5a, 0x19, 0x8e, 0x43, 0x5f, 0x0a, 0xcd, 0x66, 0xce, 0x49,
	0x16, 0x4a, 0x1a, 0xba, 0x21, 0x13, 0x9a, 0x4e, 0xcd, 0x49, 0x16, 0xb8, 0x0f, 0x9b, 0x7d, 0x2a,
	0x0f, 0xe2, 0x28, 0x62, 0x5c, 0x52, 0xd2, 0x4b, 0xf4, 0xf8, 0x34, 0x8f, 0xcb, 0x6b, 0xd0, 0x29,
	0x99, 0x34, 0x05, 0xa3, 0x5d, 0xb4, 0x29, 0xf0, 0xf7, 0x70, 0xb1, 0x97, 0x09, 0xc2, 0x13, 0xca,
	0x85, 0xcf, 0x42, 0xf3, 0xc8, 0xd7, 0x61, 0xfe, 0x2d, 0x67, 0xc1, 0x19, 0x31, 0xa2, 0xf7, 0x55,
	0xc9, 0x93, 0x2c, 0xb9, 0x58, 0xe2, 0xc9, 0xba, 0x64, 0xda, 0x01, 0x7f, 0x59, 0xd0, 0xe9, 0x71,
	0x4a, 0x7c, 0x55, 0xaf, 0xc9, 0x20, 0x7c, 0xcb, 0xd0, 0x6d, 0x40, 0x9e, 0x96, 0x0c, 0x3d, 0x97,
	0x93, 0x61, 0x18, 0x07, 0x6f, 0x28, 0x4f, 0xfd, 0xb1, 0xec, 0x65, 0xd8, 0xaf, 0xb5, 0x1c, 0x5d,
	0x87, 0xa5, 0x22, 0xda, 0x3b, 0x39, 0x49, 0x5b, 0x52, 0x3b, 0x87, 0xf6, 0x4e, 0x4e, 0xd0, 0xff,
	0x61, 0xa3, 0x88, 0xa3, 0x1f, 0x22, 0x9f, 0xeb, 0xf2, 0x39, 0x1c, 0x53, 0x97, 0xa7, 0xbe, 0xeb,
	0xe6, 0x67, 0xf6, 0x33, 0xc0, 0x77, 0xd4, 0xe5, 0xe8, 0x11, 0x5c, 0x9a, 0x72, 0x3c, 0x60, 0xa1,
	0x1c, 0xe9, 0x27, 0xaf, 0x39, 0x17, 0x27, 0x9d, 0x7f, 0xa6, 0x00, 0x78, 0x0c, 0xed, 0xde, 0xc8,
	0xe5, 0x47, 0x59, 0x4e, 0xff, 0x07, 0xea, 0x6e, 0xa0, 0x22, 0xe4, 0x0c, 0xe7, 0xa5, 0x08, 0xf4,
	0x10, 0x5a, 0x05, 0xeb, 0x69, 0xc3, 0xdc, 0x28, 0x67, 0x48, 0xc9, 0x89, 0x0e, 0xe4, 0x4c, 0xf0,
	0x7d, 0xe8, 0x18, 0xd3, 0xf9, 0xd3, 0x4b, 0xee, 0x86, 0xc2, 0xf5, 0xf4, 0x15, 0xb2, 0x64, 0x69,
	0x17, 0xa4, 0x03, 0x82, 0x7f, 0x80, 0xa6, 0xce, 0x30, 0x3d, 0x13, 0x98, 0x6e, 0x6d, 0x9d, 0xdb,
	0xad, 0x55, 0x54, 0xa8, 0xca, 0x90, 0xf2, 0x9c, 0x18, 0x15, 0x6a, 0x1f, 0xff, 0x34, 0x0b, 0x2d,
	0x93, 0xc2, 0xf1, 0xb1, 0x54, 0x89, 0xc2, 0xd4, 0x32, 0x27, 0xd4, 0xd0, 0xeb, 0x01, 0x41, 0xf7,
	0x60, 0x55, 0x8c, 0xfc, 0x28, 0x52, 0xb9, 0x5d, 0x4c, 0xf2, 0x24, 0x9a, 0x90, 0xd9, 0x7b, 0x99,
	0x25, 0x3b, 0xba, 0x0f, 0xed, 0xec, 0x84, 0x66, 0x33, 0x37, 0x95, 0xcd, 0xa2, 0x01, 0xf6, 0x98,
	0x90, 0xe8, 0x11, 0x2c, 0x67, 0x07, 0x4d, 0x6d, 0x98, 0x3f, 0xa3, 0x82, 0x2d, 0x19, 0xb4, 0xa9,
	0x19, 0xb7, 0x4d, 0x25, 0xab, 0xe9, 0x4a, 0xb6, 0x5e, 0x3a, 0x95, 0x39, 0xd4, 0x94, 0x32, 0x02,
	0x97, 0x0e, 0x68, 0x48, 0xb4, 0xbc, 0xc7, 0xc2, 0xb7, 0x3e, 0x0f, 0x74, 0xd8, 0x14, 0xda, 0x0d,
	0x0d, 0x5c, 0xff, 0xd8, 0xb4, 0x1b, 0xbd, 0x40, 0x3b, 0x50, 0xd3, 0xae, 0x49, 0x7d, 0xdc, 0x3d,
	0x6d, 0x23, 0xf1, 0xa9, 0x93, 0xc0, 0xf0, 0x9f, 0x16, 0xac, 0xbc, 0x38, 0x76, 0x3d, 0x5a, 0xaa,
	0xd1, 0x53, 0x27, 0x91, 0x6d, 0x68, 0xeb, 0x0d, 0x53, 0x0a, 0x52, 0x3f, 0x2f, 0x2a, 0xa1, 0xa9,
	0x06, 0xc5, 0x0a, 0x3f, 0xf7, 0x29, 0x15, 0x3e, 0xbb, 0x49, 0xad, 0x78, 0x93, 0x4a, 0x6c, 0xd7,
	0x3f, 0x2f, 0xb6, 0x1f, 0x03, 0x2a, 0x5e, 0x2b, 0x6b, 0xb9, 0xa9, 0x77, 0xac, 0x4f, 0xf3, 0xce,
	0x0e, 0x34, 0xf7, 0x88, 0x71, 0xca, 0x55, 0x58, 0xf4, 0x58, 0x28, 0xe9, 0x07, 0x39, 0x7c, 0x47,
	0xc7, 0xa6, 0x2a, 0xb6, 0x52, 0xd9, 0x57, 0x74, 0x2c, 0xf0, 0x5d, 0x00, 0x85, 0x4f, 0xad, 0x5d,
	0x85, 0x39, 0x97, 0x98, 0xe6, 0xbe, 0x54, 0xf1, 0x81, 0xa3, 0xf6, 0xf0, 0x03, 0x98, 0xdd, 0x23,
	0x4a, 0xb3, 0x62, 0xce, 0xa9, 0x27, 0x87, 0x31, 0x37, 0x2f, 0xda, 0x32, 0xb2, 0x43, 0x7e, 0xac,
	0xfa, 0x8d, 0xb2, 0x62, 0xfa, 0x8d, 0xfa, 0xde, 0xfd, 0xc3, 0x82, 0x96, 0xca, 0xb0, 0x03, 0xca,
	0x4f, 0x7c, 0x8f, 0xa2, 0x87, 0xba, 0x8b, 0xe9, 0xa4, 0xdc, 0xa8, 0x7a, 0xbc, 0x30, 0x78, 0xdb,
	0xe5, 0x50, 0x4f, 0x26, 0xd3, 0x19, 0xf4, 0x00, 0x1a, 0xe9, 0x74, 0x5c, 0x39, 0x5d, 0x9e, 0x99,
	0xed, 0x95, 0x53, 0x19, 0x8e, 0x67, 0xd0, 0x97, 0xd0, 0xcc, 0xe6, 0x70, 0x74, 0xf9, 0xb4, 0xfe,
	0xa2, 0x82, 0x89, 0xe6, 0x77, 0x7f, 0xb6, 0x60, 0xad, 0x3c, 0xbf, 0x9a, 0x6b, 0xfd, 0x08, 0xff,
	0x9a, 0x30, 0xdc, 0xa2, 0x7f, 0x97, 0xd4, 0x4c, 0x1f, 0xab, 0xed, 0x1b, 0xe7, 0x03, 0x93, 0x07,
	0x53, 0x2c, 0x66, 0x61, 0x2d, 0x1d, 0xbc, 0x7a, 0xae, 0x74, 0x8f, 0xd9, 0x91, 0x61, 0xd1, 0x87,
	0xc5, 0xe2, 0x94, 0x89, 0x26, 0xdc, 0xc2, 0xbe, 0x7a, 0xca, 0x52, 0x75, 0xe8, 0xc3, 0x33, 0xe8,
	0x31, 0x40, 0x3e, 0x64, 0xa2, 0xcd, 0xaa, 0xab, 0xcb, 0xd3, 0xa7, 0x3d, 0x71, 0x26, 0xc4, 0x33,
	0xe8, 0x35, 0x74, 0xca, 0x63, 0x25, 0xc2, 0x25, 0xe4, 0xc4, 0x11, 0xd5, 0xde, 0x3e, 0x13, 0x93,
	0x79, 0xe1, 0x57, 0x0b, 0x96, 0x0e, 0xd2, 0xe2, 0x65, 0xee, 0x3f, 0x80, 0x05, 0x33, 0x0d, 0xa2,
	0x4b, 0x55, 0xd2, 0xc5, 0xa1, 0xd4, 0xbe, 0x3c, 0x65, 0x37, 0xf3, 0xc0, 0x53, 0x68, 0x66, 0x43,
	0x5a, 0x25, 0x58, 0xaa, 0xd3, 0xa2, 0xbd, 0x39, 0x6d, 0x3b, 0x23, 0xfb, 0x9b, 0x05, 0x4b, 0xa6,
	0xf4, 0x18, 0xb2, 0xaf, 0x61, 0x7d, 0xf2, 0x90, 0x33, 0xf1, 0xd9, 0x6e, 0x55, 0x09, 0x9f, 0x31,
	0x1d, 0xe1, 0x19, 0xd4, 0x87, 0x46, 0x32, 0xf0, 0x48, 0x74, 0xbd, 0x9c, 0x0b, 0xd3, 0xc6, 0x21,
	0x7b, 0x42, 0x73, 0xc1, 0x33, 0xbb, 0x87, 0xd0, 0x79, 0xe1, 0x8e, 0x03, 0x1a, 0x66, 0x19, 0xdc,
	0x83, 0x7a, 0xd2, 0x91, 0x91, 0x5d, 0xd6, 0x5c, 0x9c, 0x10, 0xec, 0x8d, 0x89, 0x7b, 0x99, 0x43,
	0x46, 0xb0, 0xb8, 0xaf, 0x2a, 0xa8, 0x51, 0xfa, 0x4a, 0xfd, 0x60, 0x99, 0xd0, 0x48, 0xd0, 0xcd,
	0x4a, 0x34, 0x4c, 0x6f, 0x36, 0x53, 0x72, 0xf6, 0x0d, 0x2c, 0xf5, 0x46, 0xd4, 0x7b, 0xc7, 0xe2,
	0xec, 0x06, 0xcf, 0x01, 0xf2, 0xba, 0x5b, 0x89, 0xee, 0x53, 0x7d, 0xc6, 0xbe, 0x32, 0x75, 0x3f,
	0xbb, 0xcd, 0x13, 0x55, 0x82, 0x8d, 0xf6, 0x07, 0x50, 0xef, 0xab, 0x19, 0x5c, 0xa0, 0xf5, 0x6a,
	0x39, 0x4d, 0x35, 0x5e, 0x38, 0x25, 0x37, 0x9a, 0xde, 0xd4, 0xf5, 0x9f, 0x1b, 0xff, 0xfd, 0x3b,
	0x00, 0x00, 0xff, 0xff, 0xb2, 0xa0, 0x6e, 0x6c, 0xea, 0x10, 0x00, 0x00,
}
