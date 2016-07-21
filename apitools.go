// Copyright 2016 James Blizzard. All rights reserved.
//
// API helper functions

package apitools

import (
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
    "github.com/jimblizz/logger"
)

type ApiResponse struct {
    ResponseCode int    `json:"code"`
    Data interface{}    `json:"data"`
}

type Api struct {
    Routes map[string]string `json:"routes"`
    Log *logger.Logger
}

func New(log *logger.Logger) *Api {
    var a Api
    a.Log = log
    a.Routes = make(map[string]string)
    return &a
}

func (a Api) RegisterRoute (r *mux.Router, f func(http.ResponseWriter, *http.Request), method string, path string, description string) {
    r.HandleFunc(path, f).Methods(method)
    a.Routes[path] = description
}

func (a Api) SendOk(w http.ResponseWriter, results interface{}) () {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)

    err := json.NewEncoder(w).Encode(ApiResponse{200, results})
    if err != nil {
        a.Log.Error("Error sending response", err)
    }
}

func (a Api) SendBad(w http.ResponseWriter, message string) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusBadRequest)

    err := json.NewEncoder(w).Encode(ApiResponse{400, message})
    if err != nil {
        a.Log.Error("Error sending response", err)
    }
}

func (a Api) SendNotFound(w http.ResponseWriter, message string) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusNotFound)

    err := json.NewEncoder(w).Encode(ApiResponse{404, message})
    if err != nil {
        a.Log.Error("Error sending response", err)
    }
}

func (a Api) SendError(w http.ResponseWriter, message string) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusInternalServerError)

    err := json.NewEncoder(w).Encode(ApiResponse{500, message})
    if err != nil {
        a.Log.Error("Error sending response", err)
    }
}
