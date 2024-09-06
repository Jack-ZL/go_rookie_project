package handler

import (
	"context"
	adminProto "goRookie_project/server/proto/admin"
)

type AdminServer struct {
}

func (a *AdminServer) AdminLogin(ctx context.Context, req *adminProto.AdminLoginReq) (*adminProto.AdminLoginRsp, error) {
	return &adminProto.AdminLoginRsp{
		Code: 200,
		Msg:  "success",
		Data: &adminProto.Admin{
			Id:     1,
			Name:   "asdf",
			Mobile: "qwere456867",
			Email:  "asdf@123",
			Token:  "7a86sdyvaerf89ads887f",
		},
	}, nil
}
