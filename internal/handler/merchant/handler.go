package merchant

import (
	"context"
	"connectrpc.com/connect"
	"github.com/tekam03/panierquebec-backend/internal/model"
	"github.com/tekam03/panierquebec-backend/internal/service/merchant"

	storesv1 "github.com/tekam03/panierquebec-backend/gen/stores/v1"
	// "github.com/tekam03/panierquebec-backend/gen/stores/v1/storesv1connect"
)

type MerchantHandler struct {
	// storesv1connect.UnimplementedMerchantServiceHandler
	service merchant.Service
}

func NewMerchantHandler(s merchant.Service) *MerchantHandler {
	return &MerchantHandler{service: s}
}

func (h *MerchantHandler) CreateMerchant(
	ctx context.Context,
	req *connect.Request[storesv1.CreateMerchantRequest],
) (*connect.Response[storesv1.CreateMerchantResponse], error) {
	// Map proto to model
	m := &model.StoreMerchant{
		Name: req.Msg.Name,
		Url:  req.Msg.Url,
	}

	// Call service
	err := h.service.Create(ctx, m)
	if err != nil {
		return nil, err
	}

	// Respond with new ID
	return connect.NewResponse(&storesv1.CreateMerchantResponse{
		Id: m.ID,
	}), nil
}

func (h *MerchantHandler) GetAllMerchants(
	ctx context.Context,
	req *connect.Request[storesv1.GetAllMerchantsRequest],
) (*connect.Response[storesv1.GetAllMerchantsResponse], error) {
	// Call service to get all merchants
	merchants, err := h.service.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	// Map model to proto response
	var protoMerchants []*storesv1.StoreMerchant
	for _, m := range merchants {
		protoMerchants = append(protoMerchants, &storesv1.StoreMerchant{
			Id:   m.ID,
			Name: m.Name,
			Url:  m.Url,
		})
	}

	return connect.NewResponse(&storesv1.GetAllMerchantsResponse{
		Merchants: protoMerchants,
	}), nil
}

func (h *MerchantHandler) GetMerchantByID(
	ctx context.Context,
	req *connect.Request[storesv1.GetMerchantByIDRequest],
) (*connect.Response[storesv1.GetMerchantByIDResponse], error) {
	// Call service to get merchant by ID
	m, err := h.service.GetByID(ctx, int(req.Msg.Id))
	if err != nil {
		return nil, err
	}

	// Map model to proto response
	protoMerchant := &storesv1.StoreMerchant{
		Id:   m.ID,
		Name: m.Name,
		Url:  m.Url,
	}

	return connect.NewResponse(&storesv1.GetMerchantByIDResponse{
		Merchant: protoMerchant,
	}), nil
}

func (h *MerchantHandler) DeleteMerchant(
	ctx context.Context,
	req *connect.Request[storesv1.DeleteMerchantRequest],
) (*connect.Response[storesv1.DeleteMerchantResponse], error) {
	// Call service to delete merchant by ID
	err := h.service.Delete(ctx, int(req.Msg.Id))
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&storesv1.DeleteMerchantResponse{}), nil
}

func (h *MerchantHandler) UpdateMerchant(
	ctx context.Context,
	req *connect.Request[storesv1.UpdateMerchantRequest],
) (*connect.Response[storesv1.UpdateMerchantResponse], error) {
	// Map proto to model
	u := &model.UpdateStoreMerchant{
		Name: req.Msg.Merchant.Name,
		Url:  req.Msg.Merchant.Url,
	}
	// Call service to update merchant
	m, err := h.service.Update(ctx, int(req.Msg.Id), u)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&storesv1.UpdateMerchantResponse{
		Merchant: &storesv1.StoreMerchant{
			Id:   m.ID,
			Name: m.Name,
			Url:  m.Url,
		},
	}), nil
}
