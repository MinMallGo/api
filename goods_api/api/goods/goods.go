package goods

import (
	"api/goods_api/global"
	"errors"
	"log"
	"net/http"
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

}
