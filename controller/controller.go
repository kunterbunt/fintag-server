package controller

import (
    "net/http"
    "github.com/kunterbunt/fintag-server/model"
)

/**
* Handles server URL routing.
*/
type Controller interface {
    /** GET Processing. */
    Get(http.ResponseWriter, *http.Request)
    /** POST Processing. */
    Post(http.ResponseWriter, *http.Request)
    /** DELETE Processing. */
    Delete(http.ResponseWriter, *http.Request)
    /** PUT Processing. */
    Put(http.ResponseWriter, *http.Request)
}

/**
* Controller base class.
* Provides rudimentary functionality that should be overloaded by proper implementations.
*/
type ControllerBase struct {
    /** A Model is needed to work with the data. */
    model model.Model
}

func (controller *ControllerBase) Get(writer http.ResponseWriter, request *http.Request) {
    http.Error(writer, "GET: Method Not Allowed", http.StatusMethodNotAllowed)
}

func (controller *ControllerBase) Post(writer http.ResponseWriter, request *http.Request) {
    http.Error(writer, "POST: Method Not Allowed", http.StatusMethodNotAllowed)
}

func (controller *ControllerBase) Delete(writer http.ResponseWriter, request *http.Request) {
    http.Error(writer, "DELETE: Method Not Allowed", http.StatusMethodNotAllowed)
}

func (controller *ControllerBase) Put(writer http.ResponseWriter, request *http.Request) {
    http.Error(writer, "PUT: Method Not Allowed", http.StatusMethodNotAllowed)
}
