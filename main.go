package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/execute/:time", executeGetHandler)
	router.POST("/execute", executePostHandler)

	server := http.Server{
		Addr:    "127.0.0.1:8000",
		Handler: router,
	}
	server.ListenAndServe()
}

func executeGetHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	sleepTime, err := strconv.Atoi(p.ByName("time"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	time.Sleep(time.Duration(sleepTime) * time.Millisecond)

	fmt.Fprintf(w, "%d", time.Now().Unix())
}

func executePostHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	sleepTime, err := strconv.Atoi(r.FormValue("duration"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	time.Sleep(time.Duration(sleepTime) * time.Millisecond)

	fmt.Fprintf(w, "%d", time.Now().Unix())
}
