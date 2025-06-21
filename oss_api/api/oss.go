package api

import (
	"api/oss_api/global"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/qiniu/go-sdk/v7/storagev2/credentials"
	"github.com/qiniu/go-sdk/v7/storagev2/uploader"
	"github.com/qiniu/go-sdk/v7/storagev2/uptoken"
	"net/http"
	"path"
	"path/filepath"
	"time"
)

var basePath = "static/"

type OSSCustomResp struct {
	Key    string
	Hash   string
	Bucket string
	Name   string
	Fsize  int // 文件大小
}

type FileInfo struct {
	path string
	name string
	key  string
}

// TODO 文件上传时要判断。比如说支持什么类型，以及文件的大小

func Upload(c *gin.Context) {
	// Source
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, "get form err: %s", err.Error())
		return
	}

	filename := filepath.Base(file.Filename)
	filepath1 := path.Join(basePath, filename)
	if err = c.SaveUploadedFile(file, filepath1); err != nil {
		c.String(http.StatusBadRequest, "upload file err: %s", err.Error())
		return
	}

	info := &FileInfo{
		path: filepath1, // == 这里需要获取realpath pwd + file_path?
		name: filename,
		key:  uuid.New().String(),
	}

	str, err := SaveToOss(info)
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("save to oss err: %s", err.Error()))
		return
	}

	c.String(http.StatusOK, "File %s uploaded to oss successfully. click to view %s.", file.Filename, str)
}

func SaveToOss(info *FileInfo) (string, error) {
	dest := ""
	cnf := global.Cfg.Oss
	putPolicy, err := uptoken.NewPutPolicy(cnf.Bucket, time.Now().Add(1*time.Hour))
	if err != nil {
		return dest, err
	}
	putPolicy.SetReturnBody(`{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket)","name":"$(x:name)"}`)
	ret := &OSSCustomResp{}
	uploadManager := uploader.NewUploadManager(&uploader.UploadManagerOptions{})
	err = uploadManager.UploadFile(context.Background(), info.path, &uploader.ObjectOptions{
		UpToken:    uptoken.NewSigner(putPolicy, credentials.NewCredentials(cnf.AccessKey, cnf.SecretKey)),
		ObjectName: &info.key,
		//CustomVars: map[string]string{
		//	"name": "test upload file",
		//},
		FileName: info.key,
	}, &ret)
	if err != nil {
		return dest, err
	}

	dest = fmt.Sprintf("%s/%s", cnf.Endpoint, ret.Key)
	return dest, nil
}
