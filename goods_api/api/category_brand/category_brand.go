package catagory_brand

import (
	"api/goods_api/api"
	"api/goods_api/forms"
	"api/goods_api/global"
	proto "api/goods_api/proto/gen"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func List(c *gin.Context) {
	CategoryIdStr := c.Query("pages")
	BrandIdStr := c.Query("page_size")
	CategoryId, _ := strconv.Atoi(CategoryIdStr)
	BrandId, _ := strconv.Atoi(BrandIdStr)

	list, err := global.GoodsSrv.CategoryBrand.CategoryBrandList(context.Background(), &proto.CategoryBrandInfoRequest{
		CategoryId: int32(CategoryId),
		BrandId:    int32(BrandId),
	})
	if err != nil {
		api.HandleGrpcErr(c, err)
		return
	}
	c.JSON(http.StatusOK, list)
	return
}

func Detail(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, "参数异常")
		return
	}
	brand, err := global.GoodsSrv.CategoryBrand.GetCategoryBrand(context.Background(), &proto.CategoryBrandRequest{
		Id: int32(id),
	})
	if err != nil {
		api.HandleGrpcErr(c, err)
		return
	}
	c.JSON(http.StatusOK, brand)
	return
}

func Create(c *gin.Context) {
	params := forms.CreateCategoryBrand{}
	if err := c.ShouldBind(&params); err != nil {
		api.HandleValidatorErr(c, err)
		return
	}

	brand, err := global.GoodsSrv.CategoryBrand.CreateCategoryBrand(context.Background(), &proto.CreateCategoryBrandInfo{
		CategoryId: int32(params.CategoryId),
		BrandId:    int32(params.BrandId),
	})
	if err != nil {
		api.HandleGrpcErr(c, err)
		return
	}
	c.JSON(http.StatusOK, brand)
	return
}

func Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, "参数异常")
		return
	}
	_, err = global.GoodsSrv.CategoryBrand.DeleteCategoryBrand(context.Background(), &proto.DeleteCategoryBrandInfo{
		Id: int32(id),
	})

	if err != nil {
		api.HandleGrpcErr(c, err)
		return
	}
	c.Status(http.StatusOK)
	return
}

func Update(c *gin.Context) {
	params := forms.CreateCategoryBrand{}
	if err := c.ShouldBind(&params); err != nil {
		api.HandleValidatorErr(c, err)
		return
	}

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, "参数异常")
		return
	}

	_, err = global.GoodsSrv.CategoryBrand.UpdateCategoryBrand(context.Background(), &proto.UpdateCategoryBrandInfo{
		Id:         int32(id),
		CategoryId: int32(params.CategoryId),
		BrandId:    int32(params.BrandId),
	})

	if err != nil {
		api.HandleGrpcErr(c, err)
		return
	}
	c.Status(http.StatusOK)
	return
}
