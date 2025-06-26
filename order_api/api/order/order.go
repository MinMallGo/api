package order

import (
	"api/order_api/api"
	"api/order_api/forms"
	"api/order_api/global"
	proto "api/order_api/proto/gen"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Create(c *gin.Context) {
	param := &forms.OrderCreate{}
	if err := c.ShouldBind(&param); err != nil {
		api.HandleValidatorErr(c, err)
		return
	}
	crv := global.CrossSrv
	user, err := api.UserInfo(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		return
	}
	// 获取用户id
	// 需要商品id，数量

	order, err := crv.OrderSrv.CreateOrder(context.Background(), &proto.CreateOrderReq{
		UserID:          int32(user.ID),
		Address:         param.Address,
		RecipientName:   param.Name,
		RecipientMobile: param.Mobile,
		Message:         param.Message,
	})
	if err != nil {
		api.HandleGrpcErr(c, err)
		return
	}
	c.JSON(http.StatusOK, order)
	return
}

func Detail(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, "参数错误")
		return
	}

	user, err := api.UserInfo(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		return
	}
	// 获取用户id
	// 需要商品id，数量
	userID := 0
	if user.AuthorizationId == 1 {
		userID = user.ID
	}
	detail, err := global.CrossSrv.OrderSrv.GetListDetail(context.Background(), &proto.OrderDetailReq{
		OrderId: int32(id),
		//OrderSn: "", // 应该不需要通过订单号来获取账单详细吧
		UserId: int32(userID),
	})
	if err != nil {
		api.HandleGrpcErr(c, err)
		return
	}
	c.JSON(http.StatusOK, detail)
	return
}

func List(c *gin.Context) {
	param := &forms.OrderList{}
	if err := c.ShouldBind(&param); err != nil {
		api.HandleValidatorErr(c, err)
		return
	}

	user, err := api.UserInfo(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		return
	}
	// 获取用户id
	// 需要商品id，数量
	userID := 0
	if user.AuthorizationId == 1 {
		userID = user.ID
	}

	list, err := global.CrossSrv.OrderSrv.GetList(context.Background(), &proto.OrderListReq{
		UserId:   int32(userID),
		Page:     int32(param.Page),
		PageSize: int32(param.Size),
	})
	if err != nil {
		api.HandleGrpcErr(c, err)
		return
	}

	c.JSON(http.StatusOK, list)
	return

}
