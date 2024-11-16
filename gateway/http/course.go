package http

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	"stu_Assistant/gateway/rpc"
	"stu_Assistant/idl/pb"
)

func ListCourseHandler(ctx *gin.Context) {
	var req pb.CourseRequest
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
	req.Week = cast.ToUint32(ctx.Param("week"))
	// 调用服务端的函数
	courseResp, err := rpc.CourseList(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, RespError(ctx, err, "courseResp RPC 调用失败"))
		return
	}
	ctx.JSON(http.StatusOK, RespSuccess(ctx, courseResp))
}

func CreateCourseHandler(ctx *gin.Context) {
	var req pb.CourseRequest
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
	CourseRes, err := rpc.CourseCreate(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, RespError(ctx, err, "CourseList RPC 调度失败"))
		return
	}
	ctx.JSON(http.StatusOK, RespSuccess(ctx, CourseRes))
}

func UpdateCourseHandler(ctx *gin.Context) {
	var req pb.CourseRequest
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
	CourseRes, err := rpc.CourseUpdate(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, RespError(ctx, err, "CourseUpdate RPC 调度失败"))
		return
	}
	ctx.JSON(http.StatusOK, RespSuccess(ctx, CourseRes))
}

func DeleteCourseHandler(ctx *gin.Context) {
	var req pb.CourseRequest
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
	CourseRes, err := rpc.CourseDelete(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, RespError(ctx, err, "CourseDelete RPC 调度失败"))
		return
	}
	ctx.JSON(http.StatusOK, RespSuccess(ctx, CourseRes))
}
