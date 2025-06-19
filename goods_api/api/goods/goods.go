package goods

import (
	"api/goods_api/forms"
	"api/goods_api/global"
	proto "api/goods_api/proto/gen"
	"context"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

	list, err := global.GoodsSrv.Goods.GoodsList(context.Background(), req)
	log.Println(list, err)
	if err != nil {
		HandleGrpcErr(c, err)
		return
	}
	// TODO 定义返回信息然后自己去改
	c.JSON(http.StatusOK, list)
	return
}

func Create(c *gin.Context) {
	params := &forms.GoodsCreate{}
	if err := c.ShouldBind(params); err != nil {
		HandleGrpcErr(c, err)
	}

	// 整理然后发送请求
	goods, err := global.GoodsSrv.Goods.CreateGods(context.Background(), &proto.CreateGoodsInfo{
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
		HandleGrpcErr(c, err)
		return
	}
	c.JSON(http.StatusOK, goods)
	return
}
