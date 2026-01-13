package external_product

import (
	"context"
	"fmt"
	"strconv"

	"connectrpc.com/connect"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/tekam03/panierquebec-backend/internal/model"
	"github.com/tekam03/panierquebec-backend/internal/service/external_product"

	productsv1 "github.com/tekam03/panierquebec-backend/gen/products/v1"
)

type ExternalProductHandler struct {
	// productsv1connect.UnimplementedMerchantServiceHandler
	service external_product.Service
}

func NewExternalProductHandler(s external_product.Service) *ExternalProductHandler {
	return &ExternalProductHandler{service: s}
}

func (h *ExternalProductHandler) CreateExternalProduct(
	ctx context.Context,
	req *connect.Request[productsv1.CreateExternalProductRequest],
) (*connect.Response[productsv1.CreateExternalProductResponse], error) {
	// Map proto to model
	store_product := &model.StoreProduct{}
	if req.Msg.StoreProductId != nil {
		store_product.ID = *req.Msg.StoreProductId
	}

	ep := &model.ExternalProduct{
		Source:       req.Msg.Source,
		ExternalId:   req.Msg.ExternalId,
		Name:         req.Msg.Name,
		Description:  req.Msg.Description,
		Brand:        req.Msg.Brand,
		StoreProduct: store_product,
	}

	// Call service
	err := h.service.Create(ctx, ep)
	if err != nil {
		return nil, err
	}

	// Respond with new ID
	return connect.NewResponse(&productsv1.CreateExternalProductResponse{
		Id: ep.ID,
	}), nil
}

func (h *ExternalProductHandler) GetAllExternalProducts(
	ctx context.Context,
	req *connect.Request[productsv1.GetAllExternalProductsRequest],
) (*connect.Response[productsv1.GetAllExternalProductsResponse], error) {
	// Call service to get all external products
	external_products, err := h.service.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	// Map model to proto response
	var protoExternalProducts []*productsv1.ExternalProduct
	for _, ep := range external_products {
		protoExternalProduct, err := mapModelToProto(ep)
		if err != nil {
			return nil, err
		}

		protoExternalProducts = append(protoExternalProducts, protoExternalProduct)
	}

	return connect.NewResponse(&productsv1.GetAllExternalProductsResponse{
		ExternalProducts: protoExternalProducts,
	}), nil
}

func (h *ExternalProductHandler) DeleteExternalProduct(
	ctx context.Context,
	req *connect.Request[productsv1.DeleteExternalProductRequest],
) (*connect.Response[productsv1.DeleteExternalProductResponse], error) {
	// Call service to delete External Product by ID
	err := h.service.Delete(ctx, int(req.Msg.Id))
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&productsv1.DeleteExternalProductResponse{}), nil
}

func mapModelToProto(ep *model.ExternalProduct) (*productsv1.ExternalProduct, error) {

	var protoStoreProduct *productsv1.StoreProduct
	if ep.StoreProduct != nil {
		// TODO: Redo this part to use proper decimal handling
		var price string
		var lastUpdatedTime *timestamppb.Timestamp
		if ep.StoreProduct.Price != nil {
			price = fmt.Sprintf("%.2f", *ep.StoreProduct.Price)
		}
		lastUpdatedTime = timestamppb.New(ep.StoreProduct.LastUpdated)

		protoStoreProduct = &productsv1.StoreProduct{
			Id:                 ep.StoreProduct.ID,
			StoreId:            ep.StoreProduct.StoreId,
			SpecificProductsId: ep.StoreProduct.SpecificProductId,
			Price:              &price,
			LastUpdated:        lastUpdatedTime,
		}
	}

	protoExternalProduct := &productsv1.ExternalProduct{
		Id:           ep.ID,
		Source:       ep.Source,
		ExternalId:   ep.ExternalId,
		Name:         ep.Name,
		Description:  ep.Description,
		Brand:        ep.Brand,
		ScrapedAt:    timestamppb.New(ep.ScrapedAt),
		StoreProduct: protoStoreProduct,
	}
	return protoExternalProduct, nil
}

func mapProtoToModel(proto *productsv1.ExternalProduct) (*model.ExternalProduct, error) {

	var price *float64
	if proto.StoreProduct.Price != nil {
		priceString := *proto.StoreProduct.Price
		priceValue, err := strconv.ParseFloat(priceString, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid price format: %v", err)
		}
		price = &priceValue
	}
	modelStoreProduct := &model.StoreProduct{
		ID:                proto.StoreProduct.Id,
		StoreId:           proto.StoreProduct.StoreId,
		SpecificProductId: proto.StoreProduct.SpecificProductsId,
		Price:             price,
		LastUpdated:       proto.StoreProduct.LastUpdated.AsTime(),
	}

	return &model.ExternalProduct{
		ID:           proto.Id,
		Source:       proto.Source,
		ExternalId:   proto.ExternalId,
		Name:         proto.Name,
		Description:  proto.Description,
		Brand:        proto.Brand,
		ScrapedAt:    proto.ScrapedAt.AsTime(),
		StoreProduct: modelStoreProduct,
	}, nil
}
