package handler

import (
	"context"
	userProto "goRookie_project/server/proto/user"
)

type UserServer struct {
}

func (u *UserServer) Login(ctx context.Context, req *userProto.LoginReq) (*userProto.LoginRsp, error) {
	//TODO implement me
	panic("implement me")
}
