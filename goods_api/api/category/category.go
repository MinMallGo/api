package category

import (
	"api/goods_api/api"
	"api/goods_api/forms"
	"api/goods_api/global"
	proto "api/goods_api/proto/gen"
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/emptypb"
	"net/http"
	"strconv"
)

type Category struct {
	ID       int32       `json:"ID"`
	Name     string      `json:"Name"`
	Level    int32       `json:"Level"`
	Pid      int32       `json:"ParentCategoryID"`
	IsTab    bool        `json:"IsTab"`
	Children []*Category `json:"children"`
}

func List(c *gin.Context) {
	categories, err := global.GoodsSrv.Category.GetAllCategories(context.Background(), &emptypb.Empty{})
	if err != nil {
		api.HandleGrpcErr(c, err)
		return
	}

	node := TidyCategory(categories.Data)

	c.JSON(http.StatusOK, node)
	return
}

func TidyCategory(p []*proto.CategoryInfoResponse) []*Category {
	var flatList []Category
	for _, data := range p {
		flatList = append(flatList, Category{
			ID:    data.ID,
			Name:  data.Name,
			Level: data.Level,
			Pid:   data.ParentCategoryID,
			IsTab: data.IsTab,
		})
	}

	// 2. 用 map[id] 节点保存，方便查找父节点
	m := make(map[int32]*Category)
	for i, v := range flatList {
		m[v.ID] = &flatList[i]
	}

	var roots []*Category
	for i, category := range flatList {
		node, ok := m[category.Pid]
		// 如果他的父级存在，则
		if ok {
			if node.Children == nil {
				node.Children = make([]*Category, 0)
			}
			// 将他挂在父级下面
			node.Children = append(node.Children, &flatList[i])
		} else {
			// 如果父级不存在说明它既是父级
			roots = append(roots, m[category.ID])
		}
	}

	return roots
}

func Detail(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	sub, err := global.GoodsSrv.Category.GetSubCategory(context.Background(), &proto.CategoryInfoRequest{
		ParentCategoryID: int32(id),
	})
	if err != nil {
		api.HandleGrpcErr(c, err)
		return
	}
	node := TidyCategory(sub.Data)
	c.JSON(http.StatusOK, node)
	return
}

func Create(c *gin.Context) {
	// 解析参数
	params := &forms.CategoryCreate{}
	if err := c.ShouldBind(params); err != nil {
		api.HandleValidatorErr(c, err)
		return
	}

	category, err := global.GoodsSrv.Category.CreateCategory(context.Background(), &proto.CreateCategoryInfo{
		Name:             params.Name,
		ParentCategoryID: int32(params.Pid),
		Level:            int32(params.Level),
		IsTab:            params.IsTab,
	})
	if err != nil {
		api.HandleGrpcErr(c, err)
		return
	}

	c.JSON(http.StatusOK, category)
	return
}

func Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	_, err = global.GoodsSrv.Category.DeleteCategory(context.Background(), &proto.DeleteCategoryInfo{
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
	param := &forms.CategoryCreate{}
	if err := c.ShouldBind(param); err != nil {
		api.HandleValidatorErr(c, err)
		return
	}
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	_, err = global.GoodsSrv.Category.UpdateCategory(context.Background(), &proto.UpdateCategoryInfo{
		ID:    int32(id),
		Name:  param.Name,
		IsTab: param.IsTab,
	})
	if err != nil {
		api.HandleGrpcErr(c, err)
		return
	}

	c.Status(http.StatusOK)
	return
}
