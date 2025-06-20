package banner

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
	idStr := c.DefaultQuery("id", "")
	pagesStr := c.DefaultQuery("pages", "1")
	pageSizeStr := c.DefaultQuery("page_size", "10")
	id, _ := strconv.Atoi(idStr)
	pages, _ := strconv.Atoi(pagesStr)
	pageSize, _ := strconv.Atoi(pageSizeStr)

	list, err := global.GoodsSrv.Banner.BannerList(context.Background(), &proto.BannerInfoRequest{
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
	param := &forms.BannerCreate{}
	if err := c.ShouldBind(param); err != nil {
		api.HandleValidatorErr(c, err)
		return
	}

	banner, err := global.GoodsSrv.Banner.CreateBanner(context.Background(), &proto.CreateBannerInfo{
		Image: param.Image,
		Url:   param.Url,
		Index: int32(param.Index),
	})
	if err != nil {
		api.HandleGrpcErr(c, err)
		return
	}
	c.JSON(http.StatusOK, banner)
	return
}

func Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, "参数异常")
		return
	}
	_, err = global.GoodsSrv.Banner.DeleteBanner(context.Background(), &proto.DeleteBannerInfo{
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
	param := &forms.BannerCreate{}
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

	_, err = global.GoodsSrv.Banner.UpdateBanner(context.Background(), &proto.UpdateBannerInfo{
		ID:    int32(id),
		Image: param.Image,
		Url:   param.Url,
		Index: int32(param.Index),
	})
	if err != nil {
		api.HandleGrpcErr(c, err)
		return
	}

	c.Status(http.StatusOK)
	return
}
