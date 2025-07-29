package brand

import (
	"context"
	"github.com/gin-gonic/gin"
	"goods_api/api"
	"goods_api/forms"
	"goods_api/global"
	proto "goods_api/proto/gen"
	"net/http"
	"strconv"
)

func List(c *gin.Context) {
	idStr := c.DefaultQuery("id", "")
	pagesStr := c.DefaultQuery("pages", "1")
	pageSizeStr := c.DefaultQuery("page_size", "10")
	id, _ := strconv.Atoi(idStr)
	pages, _ := strconv.Atoi(pagesStr)
	pageSize, _ := strconv.Atoi(pageSizeStr)

	list, err := global.GoodsSrv.Brand.BrandList(context.Background(), &proto.BrandInfoRequest{
		ID:       int32(id),
		Page:     int32(pages),
		PageSize: int32(pageSize),
	})
	if err != nil {
		api.HandleGrpcErr(c, err)
		return
	}
	c.JSON(http.StatusOK, list)
	return
}

func Create(c *gin.Context) {
	param := &forms.BrandCreate{}
	if err := c.ShouldBind(param); err != nil {
		api.HandleValidatorErr(c, err)
		return
	}

	brand, err := global.GoodsSrv.Brand.CreateBrand(context.Background(), &proto.CreateBrandInfo{
		Name: param.Name,
		Logo: param.Logo,
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
	_, err = global.GoodsSrv.Brand.DeleteBrand(context.Background(), &proto.DeleteBrandInfo{
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
	param := &forms.BrandCreate{}
	if err := c.ShouldBind(param); err != nil {
		api.HandleValidatorErr(c, err)
		return
	}

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, "参数异常")
		return
	}

	_, err = global.GoodsSrv.Brand.UpdateBrand(context.Background(), &proto.UpdateBrandInfo{
		ID:   int32(id),
		Name: param.Name,
		Logo: param.Logo,
	})
	if err != nil {
		api.HandleGrpcErr(c, err)
		return
	}

	c.Status(http.StatusOK)
	return
}
