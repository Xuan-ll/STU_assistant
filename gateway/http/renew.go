package http

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"stu_Assistant/gateway/gwconfig"
)


func RenewExpireTimeHandler(ctx *gin.Context) {
    uid, exists := ctx.Get("user_id")
	if !exists {
		err := errors.New("user_id not found in context at RenewExpireTimeHandler")
		ctx.JSON(http.StatusBadRequest, RespError(ctx, err, "JWT中获取用户信息失败"))
		return	
	}
	user_id, ok := uid.(uint)
    if !ok {
		err := errors.New("user_id not uint at RenewExpireTimeHandler")
        ctx.JSON(http.StatusBadRequest, RespError(ctx, err, "user_id类型转换失败"))
		return
	}	
    token, err := gwconfig.GenerateToken(user_id)
	if err!= nil {
		ctx.JSON(http.StatusInternalServerError, RespError(ctx, err, "GenerateToken fail at RenewExpireTimeHandler"))
		return
	}	
	ctx.JSON(http.StatusOK, token)
}