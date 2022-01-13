package app

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/danieldansker/ip_service_project_troq/accessors"
	"github.com/danieldansker/ip_service_project_troq/app/handlers"
	"github.com/danieldansker/ip_service_project_troq/utils"
	"github.com/gorilla/mux"
)

const PATH = "/v1/find-coutry"
const PORT = ":80"

type App struct {
	Router *mux.Router
	DB     accessors.DBReader
}

var limiter = new(utils.RateLimiter)

type Handler func(DB accessors.DBReader, w http.ResponseWriter, r *http.Request)

func (a *App) InitApp(config utils.Configuration) {
	a.Router = mux.NewRouter()
	a.DB = getActiveDB(config.Db)
	intVar, err := strconv.ParseUint(string(config.RateLimiter), 10, 64)
	if err != nil {
		log.Fatal("Fatal Error in Configuration shutting down service")
	}
	atomic.AddUint64(&limiter.Tokens, intVar)
	go sendTick(intVar)
	setRouts(a)
}

func (a *App) handleRequest(handler Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			handlers.ReturnError(w, http.StatusTooManyRequests)
			return
		}
		handler(a.DB, w, r)
	}
}

func (a *App) RunApp() {
	fmt.Println("Running server on port " + PORT)
	log.Fatal(http.ListenAndServe(PORT, a.Router))
}

func setRouts(a *App) {
	a.Router.HandleFunc(PATH, a.handleRequest(handlers.GetIpData)).Methods("GET")
}

func sendTick(initialTokens uint64) {
	rate := time.Tick(time.Second)
	for range rate {
		limiter.ResetTokens(initialTokens)
	}
}

func getActiveDB(db string) accessors.DBReader {
	return new(accessors.DBReaderMockData)
}
