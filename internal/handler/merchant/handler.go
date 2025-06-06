package merchant

import (
	"context"
	"connectrpc.com/connect"
	"github.com/tekam03/panierquebec-backend/internal/model"
	"github.com/tekam03/panierquebec-backend/internal/service/merchant"

	storesv1 "github.com/tekam03/panierquebec-backend/gen/stores/v1"
	"github.com/tekam03/panierquebec-backend/gen/stores/v1/storesv1connect"
)

type MerchantHandler struct {
	storesv1connect.UnimplementedMerchantServiceHandler
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
