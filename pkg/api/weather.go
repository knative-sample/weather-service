package api

import (
	"encoding/json"
	"net/http"

	"fmt"

	"github.com/golang/glog"
	"github.com/knative-sample/weather-service/pkg/tablestore"
)

type Api struct {
	TableClient *tablestore.TableClient
}
type QueryResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (api *Api) QueryWeather(w http.ResponseWriter, r *http.Request) {
	res := &QueryResponse{Code: http.StatusOK}
	defer ResponseJson(res, w)
	vars := r.URL.Query()
	cityCode, ok := vars["cityCode"]
	if !ok {
		errMsg := "cityCode param a does not exist"
		glog.Errorf(errMsg)
		res.Code = http.StatusBadRequest
		res.Message = errMsg
		return
	}
	date, ok := vars["date"]
	if !ok {
		errMsg := "date param a does not exist"
		glog.Errorf(errMsg)
		res.Code = http.StatusBadRequest
		res.Message = errMsg
		return
	}
	data, err := api.TableClient.Query(cityCode[0], date[0])
	if err != nil {
		errMsg := fmt.Sprintf("Query error: %s", err.Error())
		glog.Errorf(errMsg)
		res.Code = http.StatusBadRequest
		res.Message = errMsg
		return
	}
	res.Data = data
}

func ResponseJson(qr *QueryResponse, w http.ResponseWriter) {
	rs, err := json.Marshal(qr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	glog.Infof("response: %s", rs)
	w.Header().Set("Content-Type", "application/json")
	w.Write(rs)
}
