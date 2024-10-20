package handler

import (
	"context"

	"ms/server/service"
	"ms/server/types"

	"google.golang.org/grpc"
)


type MsgHandler struct {
	crudService service.CrudService
	types.UnimplementedMsgServer
}

func NewMsgHandler(grpc *grpc.Server, cs service.CrudService) {
	msgHandler := &MsgHandler{
		crudService: cs,
	}

	types.RegisterMsgServer(grpc, msgHandler)
}

func (h *MsgHandler) Create(ctx context.Context, req *types.MsgCreate) (*types.MsgCreateResponse, error) {
	err := h.crudService.Create(ctx, req)

	if err != nil {
		return &types.MsgCreateResponse{Status: false}, err
	}

	res := &types.MsgCreateResponse{Status: true}

	return res, nil
}

func (h *MsgHandler) Read(ctx context.Context, req *types.MsgRead) (*types.MsgReadResponse, error) {
	content, err := h.crudService.Read(ctx, req)

	if err != nil {
		return &types.MsgReadResponse{}, err
	}

	res := &types.MsgReadResponse{Content: content}

	return res, nil
}

func (h *MsgHandler) Update(ctx context.Context, req *types.MsgUpdate) (*types.MsgUpdateResponse, error) {
	err := h.crudService.Update(ctx, req)

	if err != nil {
		return &types.MsgUpdateResponse{Status: false}, err
	}

	res := &types.MsgUpdateResponse{Status: true}

	return res, nil
}

func (h *MsgHandler) Delete(ctx context.Context, req *types.MsgDelete) (*types.MsgDeleteResponse, error) {
	err := h.crudService.Delete(ctx, req)

	if err != nil {
		return &types.MsgDeleteResponse{Status: false}, err
	}

	res := &types.MsgDeleteResponse{Status: true}

	return res, nil
}
