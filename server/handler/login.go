package handler

import (
	"context"
	loginProto "goRookie_project/server/proto/login"
)

type LoginServer struct {
}

func (l LoginServer) Login(ctx context.Context, req *loginProto.LoginReq) (*loginProto.LoginRsp, error) {
	//TODO implement me
	panic("implement me")
}
