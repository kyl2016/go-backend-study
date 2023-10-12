package main

import (
	"encoding/json"
	"net/http"
)

func main() {

}

type Model interface {
	ID() string
}
type DataProvider[MODEL Model] interface {
	FindByID(id string) (MODEL, error)
	List() ([]MODEL, error)
	Update(id string, model MODEL) error
	Insert(model MODEL) error
	Delete(id string) error
}

type HTTPHandler2[MODEL Model] struct {
	dataProvider DataProvider[MODEL]
}

type HTTPHandler[MODEL Model] struct {
	dataProvider    DataProvider[MODEL]
	InsertValidator func(new MODEL) error
	UpdateValidator func(old MODEL, new MODEL) error
}

func (h HTTPHandler[MODEL]) FindByID(rw http.ResponseWriter, req *http.Request) {
	// validate request here
	id := "" // extract id here
	model, err := h.dataProvider.FindByID(id)
	if err != nil {
		// error handling here
		return
	}
	err = json.NewEncoder(rw).Encode(model)
	if err != nil {
		// error handling here
		return
	}
}
