package router

import (
	"MyWebProject/api"
	"MyWebProject/view"
	"net/http"
)

func Router() {
	/*http.HandleFunc("/", view.HTML.Index)
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))
	http.HandleFunc("/login", view.HTML.Login)
	http.HandleFunc("/api/v1/login", api.API.Login)
	http.HandleFunc("/p/", view.HTML.Detail)
	http.HandleFunc("/c/", view.HTML.Category)
	http.HandleFunc("/pigeonhole", view.HTML.Pigeonhole)
	http.HandleFunc("/api/v1/post/search", api.API.SearchPost)
	http.HandleFunc("/writing", view.HTML.Writing)
	http.HandleFunc("/api/v1/post", api.API.SaveAndUpdatePost)
	http.HandleFunc("/api/v1/post/", api.API.GetPost)*/
	http.HandleFunc("/", view.HTML.Index)
	http.HandleFunc("/c/", view.HTML.Category)
	http.HandleFunc("/login", view.HTML.Login)
	//http://127.0.0.1:8081/p/7.html
	http.HandleFunc("/p/", view.HTML.Detail)
	http.HandleFunc("/writing", view.HTML.Writing)
	http.HandleFunc("/pigeonhole", view.HTML.Pigeonhole)
	http.HandleFunc("/api/v1/post", api.API.SaveAndUpdatePost)
	http.HandleFunc("/api/v1/post/", api.API.GetPost)
	http.HandleFunc("/api/v1/post/search", api.API.SearchPost)
	//http.HandleFunc("/api/v1/qiniu/token", api.API.QiniuToken)
	http.HandleFunc("/api/v1/login", api.API.Login)
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))
}
