package db

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/tidwall/buntdb"
	"locationhistoryserver/schema"
	"time"
)

type Service struct {
	db *buntdb.DB
}

func NewService() *Service {
	// Open a file that does not persist to disk.
	db, _ := buntdb.Open(":memory:")

	return &Service{db: db}
}

// ReadFromDB returns the Response for the orderId
func (s *Service) ReadFromDB(orderId string) schema.Payload {
	response := schema.Payload{}
	err := s.db.View(func(tx *buntdb.Tx) error {
		val, err := tx.Get(orderId)
		if err != nil {
			return err
		}
		err = json.Unmarshal([]byte(val), &response)

		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return schema.Payload{}
	}

	return response
}

// WriteToDB writes the data to the in memory database
func (s *Service) WriteToDB(orderId string, r schema.Request) error {
	response := s.ReadFromDB(orderId)

	if response.OrderId == "" {
		response.OrderId = orderId
	}

	response.History = append(
		response.History,
		schema.PayloadDataHistory{Lng: r.Lng, Lat: r.Lat, CreatedOn: time.Now()},
	)

	data, err := json.Marshal(response)
	if err != nil {
		errors.New("error marshalling json")
	}

	return s.db.Update(func(tx *buntdb.Tx) error {
		_, _, err := tx.Set(orderId, string(data), nil)
		return err
	})
}

// DeleteFromDB deletes the data from the in memory database
func (s *Service) DeleteFromDB(orderId string) error {
	err := s.db.Update(func(tx *buntdb.Tx) error {
		_, err := tx.Delete(orderId)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		fmt.Println(err)
		return errors.New("error deleting the data")
	}
	return nil
}
