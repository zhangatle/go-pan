package handler

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func UploadHandler(w http.ResponseWriter, r *http.Request)  {
	if r.Method == "GET" {
		// 返回上传页面
		data,err := ioutil.ReadFile("./static/view/index.html")
		if err != nil {
			io.WriteString(w, "internal server error")
		}
		io.WriteString(w, string(data))
	} else if r.Method == "POST" {
		// 接收文件上传流及存储到本地目录
		file, head, err := r.FormFile("file")
		if err != nil {
			fmt.Printf("Failed get file err:%s\n", err.Error())
			return
		}
		defer file.Close()
		newFile, err := os.Create("/tmp/" + head.Filename)
		if err != nil {
			fmt.Printf("Failed to create file, err:%s\n", err.Error())
			return
		}
		defer newFile.Close()
		_, err = io.Copy(newFile, file)
		if err != nil {
			fmt.Printf("Failed to save data to file, err:%s\n", err.Error())
			return
		}
		http.Redirect(w, r, "/file/upload/success", http.StatusFound)
	}
}

func UploadSuccessHandler(w http.ResponseWriter, r *http.Request)  {
	io.WriteString(w, "upload success")
}
