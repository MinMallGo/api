package cart

import (
	"api/order_api/api"
	"api/order_api/forms"
	"api/order_api/global"
	proto "api/order_api/proto/gen"
	"api/order_api/structure"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func UserInfo(c *gin.Context) (*structure.MyClaims, error) {
	claims, ok := c.Get(global.JWTUser)
	if !ok {
		return nil, errors.New("请先登录后操作")
	}
	myClaim, ok := claims.(*structure.MyClaims)
	if !ok {
		return nil, errors.New("请登录后操作")
	}

	return myClaim, nil
}

func Create(c *gin.Context) {
	param := &forms.CartCreate{}
	if err := c.ShouldBind(&param); err != nil {
		api.HandleValidatorErr(c, err)
		return
	}
	crv := global.CrossSrv
	user, err := UserInfo(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		return
	}
	// 获取用户id
	// 需要商品id，数量
	ctx := context.Background()
	// 先检查商品是否存在
	goodsInfo, err := crv.Goods.GetGoodsDetail(ctx, &proto.GoodsInfoRequest{Id: int32(param.GoodsId)})
	if err != nil {
		api.HandleGrpcErr(c, err)
		return
	}

	stocks, err := crv.Inventory.GetStock(ctx, &proto.GetInfo{
		GoodsId: int32(param.GoodsId),
	})

	if err != nil {
		api.HandleGrpcErr(c, err)
		return
	}

	if stocks.Stock < int32(param.Nums) {
		c.JSON(http.StatusBadRequest, "库存不足")
		return
	}
	goodsImg := ""
	if len(goodsInfo.ImageUrl) > 0 {
		goodsImg = goodsInfo.ImageUrl[0]
	}
	// 调用购物车服务添加到购物车
	_, err = crv.CartSrv.AddGoods(ctx, &proto.AddGoodsReq{
		GoodsId:  int32(param.GoodsId),
		GoodsNum: int32(param.Nums),
		GoodsImg: goodsImg,
		UserId:   int32(user.ID),
	})
	if err != nil {
		api.HandleGrpcErr(c, err)
		return
	}
	c.JSON(http.StatusOK, "success")
	return
}

func Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, "参数错误")
		return
	}

	user, err := UserInfo(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		return
	}
	ctx := context.Background()
	_, err = global.CrossSrv.CartSrv.RemoveGoods(ctx, &proto.RemoveGoodsReq{
		GoodsId: []int32{int32(id)},
		UserId:  int32(user.ID),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, "移除商品失败")
		return
	}
	c.JSON(http.StatusOK, "success")
	return
}

func Update(c *gin.Context) {
	crv := global.CrossSrv
	user, err := UserInfo(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		return
	}
	// 获取用户id
	// 需要商品id，数量
	param := &forms.CartCreate{}
	if err := c.ShouldBind(&param); err != nil {
		c.JSON(http.StatusBadRequest, "参数异常")
		return
	}
	ctx := context.Background()

	// 调用购物车服务添加到购物车
	_, err = crv.CartSrv.UpdateGoodsNum(ctx, &proto.UpdateNumReq{
		GoodsId:  int32(param.GoodsId),
		GoodsNum: int32(param.Nums),
		UserId:   int32(user.ID),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, "更新购物车商品数量失败")
		return
	}
	c.JSON(http.StatusOK, "success")
	return
}

func List(c *gin.Context) {
	user, err := UserInfo(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		return
	}
	log.Println(user.ID)
	detail, err := global.CrossSrv.CartSrv.GetCartList(context.Background(), &proto.GetCartListReq{
		UserId: int32(user.ID),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, "获取购物车信息失败")
		return
	}
	resp := make([]structure.CartListResp, 0, len(detail.Data))
	for _, data := range detail.Data {
		resp = append(resp, structure.CartListResp{
			UserID:   int(data.UserID),
			GoodsID:  int(data.GoodsID),
			GoodsImg: data.GoodsImg,
			Nums:     int(data.GoodsID),
			Checked:  data.Checked,
		})
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"total": detail.Total,
		"data":  resp,
	})
	return
}

func SelectGoods(c *gin.Context) {
	user, err := UserInfo(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		return
	}
	ctx := context.Background()
	param := forms.CartSelect{}
	if err = c.ShouldBind(&param); err != nil {
		api.HandleValidatorErr(c, err)
		return
	}

	_, err = global.CrossSrv.CartSrv.SelectGoods(ctx, &proto.SelectGoodsReq{
		GoodsId: param.GoodsIds,
		UserId:  int32(user.ID),
	})
	if err != nil {
		api.HandleGrpcErr(c, err)
		return
	}

	c.JSON(http.StatusOK, "success")
	return
}
