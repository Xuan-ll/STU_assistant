package rpc

import (
	"context"
	"errors"

	"stu_Assistant/idl/pb"
	"stu_Assistant/gateway/gwconfig"
)

// UserLogin 用户登陆
func UserLogin(ctx context.Context, req *pb.UserRequest) (resp *pb.UserDetailResponse, err error) {
	resp, err = UserService.UserLogin(ctx, req)
	if err != nil {
		return
	}

	if resp.Code != gwconfig.SUCCESS {
		err = errors.New(gwconfig.GetMsg(int(resp.Code)))
		return
	}

	return
}

// UserRegister 用户注册
func UserRegister(ctx context.Context, req *pb.UserRequest) (resp *pb.UserDetailResponse, err error) {
	resp, err = UserService.UserRegister(ctx, req)
	if err != nil {
		return
	}

	return
}
