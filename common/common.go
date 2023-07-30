package common

import (
	"MyWebProject/config"
	"MyWebProject/models"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

var Template models.HtmlTemplate

func LoadTemplate() error {
	var err error
	Template, err = models.InitTemplate(config.Cfg.System.CurrentDir + "/template/")
	if err != nil {
		return err
	}
	return nil

}
func GetRequestJsonParam(r *http.Request) map[string]interface{} {
	var params map[string]interface{}
	body, _ := io.ReadAll(r.Body)
	_ = json.Unmarshal(body, &params)
	return params
}

func Error(w http.ResponseWriter, err error) {
	var result models.Result
	result.Code = -999
	result.Error = err.Error()
	resultJson, err := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(resultJson)
	if err != nil {
		log.Println(err)
	}

}

func Success(w http.ResponseWriter, data interface{}) {
	var result models.Result
	result.Code = 200
	result.Error = ""
	result.Data = data
	resultJson, err := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(resultJson)
	if err != nil {
		log.Println(err)
	}

}
