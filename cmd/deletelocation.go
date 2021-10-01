package cmd

import (
	"locationhistoryserver/db"
	"net/http"
)

// DeleteLocationHandler for deleting the location
func DeleteLocationHandler(rw http.ResponseWriter, req *http.Request, service *db.Service) {
	OrderId := GetOrderId(req.URL.Path)

	if err := service.DeleteFromDB(OrderId); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
}
