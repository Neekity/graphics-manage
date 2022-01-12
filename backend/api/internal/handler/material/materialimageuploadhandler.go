package handler

import (
	"crypto/md5"
	"fmt"
	"go-project/graphics-manage/backend/common/helper"
	"io"
	"net/http"
	"os"
	"path"
	"strconv"

	"go-project/graphics-manage/backend/api/internal/logic/material"
	"go-project/graphics-manage/backend/api/internal/svc"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func MaterialImageUploadHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		f, fh, err := r.FormFile("graphics_img")
		if err != nil {
			httpx.Error(w, err)
			return
		}
		channelId, err := strconv.Atoi(r.FormValue("channel_id"))
		if err != nil {
			httpx.Error(w, err)
			return
		}
		if channelId == 0 {
			httpx.OkJson(w, helper.ApiError("请选择频道后上传！", nil))
			return
		}
		defer f.Close()
		// 创建保存文件
		data := []byte(fh.Filename)
		urlByte := md5.Sum(data)
		fileSuffix := path.Ext(fh.Filename)
		url := fmt.Sprintf("/%d-%x%s", channelId, urlByte, fileSuffix)
		destFile, err := os.Create(ctx.Config.RootPath + "/resource/image/" + url)
		if err != nil {
			httpx.Error(w, err)
			return
		}
		defer destFile.Close()

		// 读取表单文件，写入保存文件
		_, err = io.Copy(destFile, f)
		if err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewMaterialImageUploadLogic(r.Context(), ctx)

		resp, err := l.MaterialImageUpload(channelId, fh.Filename, url)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
