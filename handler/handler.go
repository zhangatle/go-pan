package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"pan/meta"
	"pan/util"
	"time"
)

func UploadHandler(w http.ResponseWriter, r *http.Request)  {
	if r.Method == "GET" {
		// 返回上传页面
		data,err := ioutil.ReadFile("./static/view/index.html")
		if err != nil {
			_, _ = io.WriteString(w, "internal server error")
		}
		_, _ = io.WriteString(w, string(data))
	} else if r.Method == "POST" {
		// 接收文件上传流及存储到本地目录
		file, head, err := r.FormFile("file")
		if err != nil {
			fmt.Printf("Failed get file err:%s\n", err.Error())
			return
		}
		defer file.Close()

		fileMeta := meta.FileMeta{
			FileName: head.Filename,
			FileLocation: "/tmp/" + head.Filename,
			UploadAt: time.Now().Format("2006-01-02 15:03:04"),
		}

		newFile, err := os.Create(fileMeta.FileLocation)
		if err != nil {
			fmt.Printf("Failed to create file, err:%s\n", err.Error())
			return
		}
		defer newFile.Close()
		fileMeta.FileSize, err = io.Copy(newFile, file)
		if err != nil {
			fmt.Printf("Failed to save data to file, err:%s\n", err.Error())
			return
		}

		_, _ = newFile.Seek(0, 0)
		fileMeta.FileSha1 = util.FileSha1(newFile)
		meta.UpdateFileMeta(fileMeta)

		http.Redirect(w, r, "/file/upload/success", http.StatusFound)
	}
}

func UploadSuccessHandler(w http.ResponseWriter, r *http.Request)  {
	_, _ = io.WriteString(w, "upload success")
}

func GetFileMetaHandler(w http.ResponseWriter, r *http.Request)  {
	_ = r.ParseForm()
	fileHash := r.Form["file_hash"][0]
	fMeta := meta.GetFileMeta(fileHash)
	data,err := json.Marshal(fMeta)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, _ = w.Write(data)
}
