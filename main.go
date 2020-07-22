package main

import (
	"fmt"
	"net/http"
	"pan/handler"
)

func main()  {
	http.HandleFunc("/file/upload", handler.UploadHandler)
	http.HandleFunc("/file/upload/success", handler.UploadSuccessHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Failed to start server ,err:%s", err.Error())
	}
}
