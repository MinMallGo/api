package goods

import (
	"context"
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/gin-gonic/gin"
	"goods_api/api"
	"goods_api/forms"
	"goods_api/global"
	proto "goods_api/proto/gen"
	"log"
	"net/http"
	"strconv"
)

func List(c *gin.Context) {
	// 从get请求里面获取参数，然后解析参数并发送数据
	// 先拿到等待构造的参数
	priceMinStr := c.DefaultQuery("price_min", "0")
	priceMaxStr := c.DefaultQuery("price_max", "0")
	isHotStr := c.DefaultQuery("is_hot", "false")
	isNewStr := c.DefaultQuery("is_new", "false")
	isTabStr := c.DefaultQuery("is_tab", "false")
	topCategoryStr := c.DefaultQuery("top_category", "0")
	pagesStr := c.DefaultQuery("pages", "1")
	pageSizeStr := c.DefaultQuery("page_size", "10")
	keyWord := c.DefaultQuery("key_word", "")
	brandStr := c.DefaultQuery("brand", "0")

	// 转换
	priceMin, _ := strconv.Atoi(priceMinStr)
	priceMax, _ := strconv.Atoi(priceMaxStr)
	isHot, _ := strconv.ParseBool(isHotStr)
	isNew, _ := strconv.ParseBool(isNewStr)
	isTab, _ := strconv.ParseBool(isTabStr)
	topCategory, _ := strconv.Atoi(topCategoryStr)
	pages, _ := strconv.Atoi(pagesStr)
	pageSize, _ := strconv.Atoi(pageSizeStr)
	brand, _ := strconv.Atoi(brandStr)

	req := &proto.GoodsFilterRequest{
		PriceMin:    int32(priceMin),
		PriceMax:    int32(priceMax),
		IsHot:       isHot,
		IsNew:       isNew,
		IsTab:       isTab,
		TopCategory: int32(topCategory),
		Pages:       int32(pages),
		PageSize:    int32(pageSize),
		KeyWord:     keyWord,
		Brand:       int32(brand),
	}
	e, b := sentinel.Entry("goods-list")
	if b != nil {
		c.JSON(http.StatusTooManyRequests, gin.H{
			"msg": "请求过于频繁，请稍后再试",
		})
		return
	}
	list, err := global.GoodsSrv.Goods.GoodsList(context.Background(), req)
	e.Exit()
	if err != nil {
		api.HandleGrpcErr(c, err)
		return
	}
	// TODO 定义返回信息然后自己去改
	c.JSON(http.StatusOK, list)
	return
}

func Create(c *gin.Context) {
	params := &forms.GoodsCreate{}
	if err := c.ShouldBind(params); err != nil {
		api.HandleValidatorErr(c, err)
		return
	}

	// 整理然后发送请求
	goods, err := global.GoodsSrv.Goods.CreateGoods(context.Background(), &proto.CreateGoodsInfo{
		CategoryId:      params.CategoryID,
		BrandId:         params.BrandID,
		ShipFree:        params.ShipFree,
		Stock:           params.Stock,
		Name:            params.Name,
		GoodsSn:         params.GoodsSn,
		MarketPrice:     params.MarketPrice,
		ShopPrice:       params.ShopPrice,
		GoodsBrief:      params.GoodsBrief,
		ImageUrl:        params.ImageUrl,
		Description:     params.Description,
		GoodsFrontImage: params.GoodsFrontImage,
	})
	// TODO 库存服务  分布式事务的一致性
	if err != nil {
		api.HandleGrpcErr(c, err)
		return
	}
	c.JSON(http.StatusOK, goods)
	return
}

func Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, "参数错误")
		return
	}

	_, err = global.GoodsSrv.Goods.DeleteGoods(context.Background(), &proto.DeleteGoodsInfo{
		Id: int32(id),
	})
	if err != nil {
		api.HandleGrpcErr(c, err)
		return
	}

	c.Status(http.StatusOK)
	return
}

func UpdateStatus(c *gin.Context) {
	params := &forms.UpdateStatus{}
	if err := c.ShouldBind(params); err != nil {
		api.HandleValidatorErr(c, err)
		return
	}

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, "参数错误")
		return
	}

	_, err = global.GoodsSrv.Goods.UpdateGoods(context.Background(), &proto.UpdateGoodsInfo{
		Id:     int32(id),
		IsNew:  params.IsNew,
		OnSale: params.IsOnSale,
		IsHot:  params.IsHot,
	})
	if err != nil {
		api.HandleGrpcErr(c, err)
		return
	}
	c.Status(http.StatusOK)
	return
}

func Update(c *gin.Context) {
	params := &forms.GoodsCreate{}
	if err := c.ShouldBind(params); err != nil {
		api.HandleValidatorErr(c, err)
		return
	}
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, "参数异常")
		return
	}
	_, err = global.GoodsSrv.Goods.UpdateGoods(context.Background(), &proto.UpdateGoodsInfo{
		Id:              int32(id),
		CategoryId:      params.CategoryID,
		BrandId:         params.BrandID,
		ShipFree:        params.ShipFree,
		Stock:           params.Stock,
		Name:            params.Name,
		GoodsSn:         params.GoodsSn,
		MarketPrice:     params.MarketPrice,
		ShopPrice:       params.ShopPrice,
		GoodsBrief:      params.GoodsBrief,
		ImageUrl:        params.ImageUrl,
		Description:     params.Description,
		GoodsFrontImage: params.GoodsFrontImage,
	})
	if err != nil {
		api.HandleGrpcErr(c, err)
		return
	}
	c.Status(http.StatusOK)
	return
}

func Stock(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, "参数异常")
		return
	}
	log.Println(id)
	// TODO 查询库存
	return
}
