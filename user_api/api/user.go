package api

import (
	"api/user_api/forms"
	"api/user_api/global"
	"api/user_api/global/response"
	"api/user_api/proto"
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"log"
	"net/http"
	"strings"
	"time"
)

func HandleGrpcErr(c *gin.Context, err error) {
	sts, ok := status.FromError(err)
	if ok {
		switch sts.Code() {
		case codes.Unavailable:
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": sts.Code(),
				"msg":  sts.Message(),
			})
		case codes.NotFound:
			c.JSON(http.StatusNotFound, gin.H{
				"code": sts.Code(),
				"msg":  sts.Message(),
			})
		case codes.AlreadyExists:
			c.JSON(http.StatusConflict, gin.H{
				"code": sts.Code(),
				"msg":  sts.Message(),
			})
		case codes.InvalidArgument:
			c.JSON(http.StatusBadRequest, gin.H{
				"code": sts.Code(),
				"msg":  sts.Message(),
			})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": 0,
				"msg":  sts.Message(),
			})
		}
	}
}

func removeTopStruct(fields map[string]string) map[string]string {
	rsp := map[string]string{}
	for field, err := range fields {
		rsp[field[strings.Index(field, ".")+1:]] = err
	}
	return rsp
}

func HandleValidatorErr(c *gin.Context, err error) {
	var errs validator.ValidationErrors
	ok := errors.As(err, &errs)
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}
	for _, fieldError := range errs {
		log.Println(fieldError.Field())

	}
	c.JSON(http.StatusOK, gin.H{
		"msg": removeTopStruct(errs.Translate(global.Trans)),
	})
	return
}

func GetUserList(ctx *gin.Context) {
	/*
		1. 连接grpc服务
		2. 调用
		3. 处理返回值
		4. 返回
	*/
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", global.Cfg.UserServer.Host, global.Cfg.UserServer.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		HandleGrpcErr(ctx, err)
	}
	defer conn.Close()
	pageInfo := &forms.UserListForm{}
	if err := ctx.ShouldBind(pageInfo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  err.Error(),
		})
		return
	}

	userClient := proto.NewUserClient(conn)
	list, err := userClient.GetUserList(context.Background(), &proto.PaginateInfo{
		Page: uint32(pageInfo.Page),
		Size: uint32(pageInfo.Size),
	})
	if err != nil {
		HandleGrpcErr(ctx, err)
	}

	res := make([]response.UserResponse, 0, len(list.Data))
	for _, user := range list.Data {
		res = append(res, response.UserResponse{
			Id:       user.Id,
			Mobile:   user.Mobile,
			NickName: user.NickName,
			Birthday: response.JsonTime(time.Unix(int64(user.Birthday), 0)),
			Gender:   user.Gender,
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":  0,
		"data":  res,
		"total": list.Total,
	})
}

func PasswordLogin(ctx *gin.Context) {
	param := &forms.PasswordLoginForm{}
	if err := ctx.ShouldBind(param); err != nil {
		HandleValidatorErr(ctx, err)
	}
}
