package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"stu_Assistant/gateway/gwconfig"
	"stu_Assistant/gateway/rpc"
	"stu_Assistant/idl/pb"
)

// UserRegisterHandler 用户注册
func UserRegisterHandler(ctx *gin.Context) {
	var req pb.UserRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, RespError(ctx, err, "UserRegister Bind 绑定参数失败"))
		return
	}
	userResp, err := rpc.UserRegister(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, RespError(ctx, err, "UserRegister RPC 调用失败"))
		return
	}
	ctx.JSON(http.StatusOK, RespSuccess(ctx, userResp))
}

// UserLoginHandler 用户登录
func UserLoginHandler(ctx *gin.Context) {
	var req pb.UserRequest
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, RespError(ctx, err, "UserLogin Bind 绑定参数失败"))
		return
	}
	userResp, err := rpc.UserLogin(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, RespError(ctx, err, "UserLogin RPC 调用失败"))
		return
	}
	token, err := gwconfig.GenerateToken(uint(userResp.UserDetail.Id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, RespError(ctx, err, "GenerateToken 失败"))
		return
	}
	res := &gwconfig.TokenData{
		User:  userResp,
		Token: token,
	}
	ctx.JSON(http.StatusOK, RespSuccess(ctx, res))
}
