package controllers

import (
	"encoding/json"
	"engine_app/controllers/filters"
	"net/http"
)

type BaseCrudController struct {
	AppInjection *AppInjection
}

func (bc *BaseCrudController) Delete(obj interface{}, request *http.Request, writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
	var deleteIntent filters.IdsFilter
	var err = json.NewDecoder(request.Body).Decode(&deleteIntent)
	if err != nil {
		writer.WriteHeader(500)
		writer.Write([]byte(err.Error()))
	} else {
		err = bc.AppInjection.Provider.DeleteStatement(obj, deleteIntent.Ids)
		if err != nil {
			writer.WriteHeader(500)
			writer.Write([]byte(err.Error()))
		} else {
			writer.WriteHeader(200)
		}
	}
}

func (bc *BaseCrudController) GetFilteredBy(
	obj interface{},
	request *http.Request,
	writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")

	field, _ := request.URL.Query()["field"]
	args, _ := request.URL.Query()["val"]
	filter := filters.FilterBy{
		Field: field[0],
		Args:  args,
	}

	if err := bc.AppInjection.Provider.QueryListStatement(obj, filter); err != nil {
		writer.WriteHeader(500)
		writer.Write([]byte(err.Error()))
	} else {
		writer.WriteHeader(200)
		json.NewEncoder(writer).Encode(obj)
	}
}

func (bc *BaseCrudController) GetAll(
	obj interface{},
	_ *http.Request,
	writer http.ResponseWriter) {
	if err := bc.AppInjection.Provider.QueryListStatementAll(obj); err != nil {
		writer.WriteHeader(500)
		writer.Write([]byte(err.Error()))
	} else {
		writer.WriteHeader(200)
		json.NewEncoder(writer).Encode(obj)
	}
}

func (bc *BaseCrudController) Put(
	obj interface{},
	_ *http.Request,
	writer http.ResponseWriter) {
	if err := bc.AppInjection.Provider.UpdateStatement(obj); err != nil {
		writer.WriteHeader(500)
		writer.Write([]byte(err.Error()))
	} else {
		writer.WriteHeader(200)
		json.NewEncoder(writer).Encode(obj)
	}
}

func (bc *BaseCrudController) Add(
	obj interface{},
	_ *http.Request,
	writer http.ResponseWriter) {
	if err := bc.AppInjection.Provider.AddStatement(obj); err != nil {
		writer.WriteHeader(500)
		writer.Write([]byte(err.Error()))
	} else {
		writer.WriteHeader(200)
		json.NewEncoder(writer).Encode(obj)
	}
}

func Decode(request *http.Request, obj interface{}, writer http.ResponseWriter) {
	err := json.NewDecoder(request.Body).Decode(obj)
	if err != nil {
		writer.Write([]byte("Decode error"))
	}
}
