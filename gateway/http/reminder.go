package http

import (
	"errors"
	"net/http"
	"stu_Assistant/idl/pb"
    "stu_Assistant/gateway/rpc"

	"github.com/gin-gonic/gin"
)

func UpdateReminderHandler(ctx *gin.Context) {
	var req pb.ReminderRequest
	if err:= ctx.Bind(&req); err!= nil {
		ctx.JSON(http.StatusBadRequest, RespError(ctx, err, "绑定参数失败"))	
    	return
	}
	uid, exists := ctx.Get("user_id")
    if !exists {
		err := errors.New("user_id not found in context")
        ctx.JSON(http.StatusBadRequest, RespError(ctx, err, "JWT中获取用户信息失败"))
		return
	} 
	user_id, ok := uid.(uint)
    if !ok {
		err := errors.New("user_id not uint")
        ctx.JSON(http.StatusBadRequest, RespError(ctx, err, "user_id类型转换失败"))
		return
	}
	req.Uid = uint64(user_id)
    reminderRes, err := rpc.ReminderUpdate(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, RespError(ctx, err, "ReminderUpdate RPC 调度失败"))
		return
	}
	ctx.JSON(http.StatusOK, RespSuccess(ctx, reminderRes))	
}

func GetReminderHandler(ctx *gin.Context) {
	var taskReq pb.GetTaskRequest
	if err := ctx.Bind(&taskReq); err != nil {
		ctx.JSON(http.StatusBadRequest, RespError(ctx, err, "绑定参数失败"))
		return
	}	
	uid, exists := ctx.Get("user_id")
    if !exists {
		err := errors.New("user_id not found in context")
        ctx.JSON(http.StatusBadRequest, RespError(ctx, err, "JWT中获取用户信息失败"))
		return
	} 
        user_id,ok := uid.(uint)
    if !ok {
		err := errors.New("user_id not uint")
        ctx.JSON(http.StatusBadRequest, RespError(ctx, err, "user_id类型转换失败"))
		return
	}
	taskReq.Uid = uint64(user_id)
	reminderRes, err := rpc.ReminderGet(ctx, &taskReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, RespError(ctx, err, "ReminderGet RPC 调度失败"))
		return
	}
	ctx.JSON(http.StatusOK, RespSuccess(ctx, reminderRes))
}