package api

import (
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type APIServer struct {
	RaftHandler  RaftStruct
	StoreHandler StoreStruct
}

func NewAPIServer() *APIServer {
	return API
}

func (s *APIServer) StartAPI() {
	log.Print("Starting API...")

	router := mux.NewRouter()

	// store endpoints
	router.HandleFunc("/store", s.storeStore).Methods("POST")
	router.HandleFunc("/store/{key}", s.storeGet).Methods("GET")
	router.HandleFunc("/store/{key}", s.storeDelete).Methods("DELETE")

	// raft endpoints
	router.HandleFunc("/join", s.raftJoin).Methods("POST")
	router.HandleFunc("/disconnect", s.raftDisconnect).Methods("POST")
	router.HandleFunc("/stats", s.raftStats).Methods("GET")
}

func (s *APIServer) storeStore(w http.ResponseWriter, req *http.Request) {
	s.StoreHandler.Store()
}

func (s *APIServer) storeGet(w http.ResponseWriter, req *http.Request) {
	s.StoreHandler.Get()
}

func (s *APIServer) storeDelete(w http.ResponseWriter, req *http.Request) {
	s.StoreHandler.Delete()
}

func (s *APIServer) raftJoin(w http.ResponseWriter, req *http.Request) {
	s.RaftHandler.Join()
}

func (s *APIServer) raftDisconnect(w http.ResponseWriter, req *http.Request) {
	s.RaftHandler.Disconnect()
}

func (s *APIServer) raftStats(w http.ResponseWriter, req *http.Request) {
	s.RaftHandler.Stats()
}
