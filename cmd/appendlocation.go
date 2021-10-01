package cmd

import (
	"encoding/json"
	"io/ioutil"
	"locationhistoryserver/db"
	"locationhistoryserver/schema"
	"net/http"
	"strings"
)

// AppendLocationHandler for appending the location
func AppendLocationHandler(rw http.ResponseWriter, req *http.Request, dbService *db.Service) {
	// Get the JSON body
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	if len(strings.Split(req.URL.String(), "/")) == 3 {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	r := schema.Request{}

	// Convert to Request struct
	if json.Unmarshal(body, &r) != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	orderId := strings.Split(req.URL.String(), "/")[2]

	writeToDB(dbService, orderId, r)

	rw.WriteHeader(http.StatusOK)
}

func writeToDB(dbService *db.Service, orderId string, r schema.Request) error {
	return dbService.WriteToDB(orderId, r)
}

