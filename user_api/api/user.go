package api

import (
	"api/user_api/forms"
	"api/user_api/global"
	"api/user_api/global/response"
	"api/user_api/middleware"
	"api/user_api/proto"
	"api/user_api/structure"
	"api/user_api/utils"
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
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

func FilterService(name string) (map[string]*api.AgentService, error) {
	config := api.DefaultConfig()
	config.Address = fmt.Sprintf("%s:%d", global.Cfg.Consul.Host, global.Cfg.Consul.Port)
	client, err := api.NewClient(config)
	if err != nil {
		panic(err)
	}
	services, err := client.Agent().ServicesWithFilter(fmt.Sprintf(`ID == "%s"`, name))
	return services, err
}

func GetUserList(ctx *gin.Context) {
	pageInfo := &forms.UserListForm{}
	if err := ctx.ShouldBind(pageInfo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  err.Error(),
		})
		return
	}

	list, err := global.UserSrv.GetUserList(context.Background(), &proto.PaginateInfo{
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
		return
	}
	// 验证图形验证码
	if !utils.VerifyCaptcha(param.CaptchaID, param.Captcha) {
		zap.L().Info("[PasswordLogin] 登录验证码验证失败")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "验证码错误",
		})
		//return
	}

	/*
		1. 通过手机号获取用户信息
		2. 通过用户信息比对的密码查询grpc中的密码
		3. 对比密码并登录
	*/
	user, err := global.UserSrv.GetUserByMobile(context.Background(), &proto.MobileRequest{Mobile: param.Mobile})
	if err != nil {
		HandleGrpcErr(ctx, err)
		return
	}

	login(ctx, param.Password, user)

}

func login(ctx *gin.Context, password string, user *proto.UserInfoResponse) {
	// 对比密码
	check, err := global.UserSrv.CheckPassword(context.Background(), &proto.CheckPasswordRequest{
		Password:        password,
		EncryptPassword: user.Password,
	})
	if err != nil {
		HandleGrpcErr(ctx, err)
		return
	}

	if !check.Success {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "请输入正确的账号或者密码",
		})
		return
	}

	j := middleware.NewJwt()
	token, err := j.CreateToken(structure.MyClaims{
		ID:              int(user.Id),
		NickName:        user.NickName,
		AuthorizationId: int(user.Role),
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "shop_api",
			Subject:   "login",
			Audience:  nil,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  nil,
			ID:        fmt.Sprintf("%d", user.Id),
		},
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "登录失败",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":       0,
		"token":      token,
		"token_type": "Bearer",
		"expires_at": jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
	})
	return
}

func UserCreate(ctx *gin.Context) {
	param := &forms.UserCreateForm{}
	if err := ctx.ShouldBind(param); err != nil {
		HandleValidatorErr(ctx, err)
		return
	}

	user, err := global.UserSrv.CreateUser(context.Background(), &proto.CreateUserRequest{
		NickName: param.Nickname,
		Password: param.Password,
		Mobile:   param.Mobile,
	})
	if err != nil {
		HandleGrpcErr(ctx, err)
		return
	}

	// 调用一下登录
	login(ctx, param.Password, user)
}
