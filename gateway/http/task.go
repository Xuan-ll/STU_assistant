package http

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	"stu_Assistant/gateway/rpc"
	"stu_Assistant/idl/pb"
)

func ListTaskHandler(ctx *gin.Context) {
	var taskReq pb.TaskRequest
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
	// 调用服务端的函数
	taskResp, err := rpc.TaskList(ctx, &taskReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, RespError(ctx, err, "taskResp RPC 调用失败"))
		return
	}
	ctx.JSON(http.StatusOK, RespSuccess(ctx, taskResp))
}

func CreateTaskHandler(ctx *gin.Context) {
	var req pb.TaskRequest
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, RespError(ctx, err, "绑定参数失败"))
		return
	}
	// user, err := ctl.GetUserInfo(ctx.Request.Context())
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
	taskRes, err := rpc.TaskCreate(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, RespError(ctx, err, "TaskList RPC 调度失败"))
		return
	}
	ctx.JSON(http.StatusOK, RespSuccess(ctx, taskRes))
}

func GetTaskHandler(ctx *gin.Context) {
	var req pb.TaskRequest
	if err := ctx.Bind(&req); err != nil {
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
	req.Id = cast.ToUint64(ctx.Param("id"))
	taskRes, err := rpc.TaskGet(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, RespError(ctx, err, "TaskList RPC 调度失败"))
		return
	}
	ctx.JSON(http.StatusOK, RespSuccess(ctx, taskRes))
}

func UpdateTaskHandler(ctx *gin.Context) {
	var req pb.TaskRequest
	if err := ctx.Bind(&req); err != nil {
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
	req.Id = cast.ToUint64(ctx.Param("id"))
	taskRes, err := rpc.TaskUpdate(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, RespError(ctx, err, "TaskUpdate RPC 调度失败"))
		return
	}
	ctx.JSON(http.StatusOK, RespSuccess(ctx, taskRes))
}

func DeleteTaskHandler(ctx *gin.Context) {
	var req pb.TaskRequest
	if err := ctx.Bind(&req); err != nil {
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
		err := errors.New("user_id not uint64")
        ctx.JSON(http.StatusBadRequest, RespError(ctx, err, "user_id类型转换失败"))
		return
	}
	req.Uid = uint64(user_id)
	req.Id = cast.ToUint64(ctx.Param("id"))
        taskRes, err := rpc.TaskDelete(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, RespError(ctx, err, "TaskDelete RPC 调度失败"))
		return
	}
	ctx.JSON(http.StatusOK, RespSuccess(ctx, taskRes))
}
