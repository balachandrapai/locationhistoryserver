package cmd

import (
	"encoding/json"
	"fmt"
	"locationhistoryserver/db"
	"locationhistoryserver/schema"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

// RetrieveLocationHandler for retrieving the location
func RetrieveLocationHandler(rw http.ResponseWriter, req *http.Request, service *db.Service) {
	orderId := GetOrderId(req.URL.Path)

	maxHistory := req.URL.Query().Get("max")
	max := 0
	if maxHistory != "" {
		max, _= strconv.Atoi(maxHistory)
	}

	data := ReadFromDB(service, orderId, max)

	// Convert it to JSON response
	response, err := json.Marshal(data)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, err.Error())
		return
	}
	rw.WriteHeader(http.StatusOK)
	rw.Write(response)
}

func ReadFromDB(service *db.Service, orderId string, max int) schema.Response {
	data := service.ReadFromDB(orderId)

	response := schema.Response{}
	response.OrderId = data.OrderId

	var reversed []schema.History

	flag := 1
	// reverse order
	// and append into new slice
	for i := range data.History {
		n := data.History[len(data.History)-1-i]
		reversed = append(reversed, schema.History{Lat: n.Lat, Lng: n.Lng})
		if flag == max {
			break
		}
		flag +=1
	}

	response.History = reversed

	return response
}


func GetOrderId(path string) string {
	var re = regexp.MustCompile(`/(.\w+)`)
	s := re.FindAllString(path, -1)[1]
	return strings.TrimLeft(s, "/")
}