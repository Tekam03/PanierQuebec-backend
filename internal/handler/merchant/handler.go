package merchant

import (
	"context"

	"connectrpc.com/connect"
	repoMerchant "github.com/tekam03/panierquebec-backend/internal/repository/merchant"
	"github.com/tekam03/panierquebec-backend/internal/service/merchant"

	productsv1 "github.com/tekam03/panierquebec-backend/gen/products/v1"
	// "github.com/tekam03/panierquebec-backend/gen/stores/v1/productsv1connect"
)

type MerchantHandler struct {
	// productsv1connect.UnimplementedMerchantServiceHandler
	service merchant.Service
}

func NewMerchantHandler(s merchant.Service) *MerchantHandler {
	return &MerchantHandler{service: s}
}

func (h *MerchantHandler) CreateMerchant(
	ctx context.Context,
	req *connect.Request[productsv1.CreateMerchantRequest],
) (*connect.Response[productsv1.StoreMerchant], error) {
	// Map proto to model
	m := &repoMerchant.MerchantCreate{
		Name: req.Msg.Name,
		Url:  req.Msg.Url,
	}

	// Call service
	createdMerchant, err := h.service.Create(ctx, m)
	if err != nil {
		return nil, err
	}

	// Respond with new ID
	return connect.NewResponse(&productsv1.StoreMerchant{
		Id:   createdMerchant.ID,
		Name: createdMerchant.Name,
		Url:  createdMerchant.Url,
	}), nil
}

func (h *MerchantHandler) GetAllMerchants(
	ctx context.Context,
	req *connect.Request[productsv1.GetAllMerchantsRequest],
) (*connect.Response[productsv1.GetAllMerchantsResponse], error) {
	// Call service to get all merchants
	merchants, err := h.service.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	// Map model to proto response
	var protoMerchants []*productsv1.StoreMerchant
	for _, m := range merchants {
		protoMerchants = append(protoMerchants, &productsv1.StoreMerchant{
			Id:   m.ID,
			Name: m.Name,
			Url:  m.Url,
		})
	}

	return connect.NewResponse(&productsv1.GetAllMerchantsResponse{
		Merchants: protoMerchants,
	}), nil
}

func (h *MerchantHandler) GetMerchantByID(
	ctx context.Context,
	req *connect.Request[productsv1.GetMerchantByIDRequest],
) (*connect.Response[productsv1.StoreMerchant], error) {
	// Call service to get merchant by ID
	m, err := h.service.GetByID(ctx, req.Msg.Id)
	if err != nil {
		return nil, err
	}

	// Map model to proto response
	protoMerchant := &productsv1.StoreMerchant{
		Id:   m.ID,
		Name: m.Name,
		Url:  m.Url,
	}

	return connect.NewResponse(&productsv1.StoreMerchant{
		Id:   protoMerchant.Id,
		Name: protoMerchant.Name,
		Url:  protoMerchant.Url,
	}), nil
}

func (h *MerchantHandler) DeleteMerchant(
	ctx context.Context,
	req *connect.Request[productsv1.DeleteMerchantRequest],
) (*connect.Response[productsv1.DeleteMerchantResponse], error) {
	// Call service to delete merchant by ID
	err := h.service.Delete(ctx, req.Msg.Id)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&productsv1.DeleteMerchantResponse{}), nil
}

func (h *MerchantHandler) UpdateMerchant(
	ctx context.Context,
	req *connect.Request[productsv1.UpdateMerchantRequest],
) (*connect.Response[productsv1.StoreMerchant], error) {
	// Map proto to model
	u := &repoMerchant.MerchantPatch{
		Name: req.Msg.Name,
		Url:  req.Msg.Url,
	}
	// Call service to update merchant
	m, err := h.service.Patch(ctx, req.Msg.Id, u)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&productsv1.StoreMerchant{
		Id:   m.ID,
		Name: m.Name,
		Url:  m.Url,
	}), nil
}
