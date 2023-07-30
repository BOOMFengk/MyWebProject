package main

import (
	"MyWebProject/common"
	"MyWebProject/router"
	"log"
	"net/http"
)

func init() {
	common.LoadTemplate()

}

func main() {

	server := http.Server{
		Addr: "127.0.0.1:8082",
	}
	router.Router()

	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
