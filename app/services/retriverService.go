package services

import (
	"github.com/danieldansker/ip_service_project_troq/accessors"
	"github.com/danieldansker/ip_service_project_troq/model"
)

func IpDataReciverService(DB accessors.DBReader, ipAddress string) model.IpData {
	return recordTransformer(accessors.ReadData(DB, ipAddress))
}

func recordTransformer(record accessors.IpDataRecord) model.IpData {
	data := new(model.IpData)
	data.Country = record.Country
	data.City = record.City
	return *data
}
