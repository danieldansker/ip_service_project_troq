package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/danieldansker/ip_service_project_troq/accessors"
	"github.com/danieldansker/ip_service_project_troq/app/services"
	"github.com/danieldansker/ip_service_project_troq/model"
)

const IP = "ip"

func GetIpData(DB accessors.DBReader, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ipAddress := r.URL.Query().Get(IP)
	ipData := services.IpDataReciverService(DB, ipAddress)
	json.NewEncoder(w).Encode(&ipData)
}

func ReturnError(w http.ResponseWriter, err int) {
	responseErr := new(model.JSONError)
	responseErr.Err = strconv.Itoa(err)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseErr)
}
