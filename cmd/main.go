package main

import (
	"fmt"
	"net/http"
	"os"

	"flag"

	"github.com/knative-sample/weather-service/pkg/api"
	"github.com/knative-sample/weather-service/pkg/tablestore"
	"github.com/knative-sample/weather-service/pkg/utils/logs"
)

func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "ok")
}
func main() {
	flag.Parse()
	logs.InitLogs()
	defer logs.FlushLogs()
	tableClient := tablestore.InitClient()
	weatherApi := api.Api{
		TableClient: tableClient,
	}
	http.HandleFunc("/health", ping)
	http.HandleFunc("/api/weather/query", weatherApi.QueryWeather)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)

}
