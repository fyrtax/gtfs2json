// hard-coded files in structure map
// TripUpdate_A : [type:'protobuf', filename:'TripUpdates_T.pb', mod_time: 0, vehicle: 'bus'] ...
// check if file is in map, else throw error and then 404
// if exists then request with Last-Modified header to server

// files:
// TripUpdates_A.pb
// TripUpdates_T.pb
// ServiceAlerts_A.pb
// ServiceAlerts_T.pb
// VehiclePositions_A.pb
// VehiclePositions_T.pb
// GTFS_KRK_A.zip
// GTFS_KRK_T.zip

package main

import (
	"io"
	"net/http"
	"strings"
)

// pass it to the interface
const GTFS_SERVER = "https://gtfs.ztp.krakow.pl/"

type GtfsProxyStore interface {
	UpdateFile(name string, time int64)
	GetFileDetails(name string) GtfsFile
	SaveFile(name string, file io.Writer) error
	GetFile(name string) (io.Reader, error)
}

type GtfsProxyServer struct {
	store GtfsProxyStore
	http.Handler
}

type GtfsFile struct {
	Type     string
	Filename string
	ModTime  int64
	Vehicle  string
}

type GtfsServerFiles map[string]GtfsFile

const jsonContentType = "application/json"

func NewProxyServer(store GtfsProxyStore) (*GtfsProxyServer, error) {
	p := new(GtfsProxyServer)

	p.store = store

	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(p.playersHandler))

	p.Handler = router

	return p, nil
}

func (p *GtfsProxyServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", jsonContentType)
	//json.NewEncoder(w).Encode(p.store.GetLeague())
}

func (p *GtfsProxyServer) playersHandler(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	switch r.Method {
	case http.MethodPost:
		p.processWin(w, player)
	case http.MethodGet:
		p.showScore(w, player)
	}
}

func (p *GtfsProxyServer) showScore(w http.ResponseWriter, player string) {
	//score := p.store.GetPlayerScore(player)
	//
	//if score == 0 {
	//	w.WriteHeader(http.StatusNotFound)
	//}
	//
	//fmt.Fprint(w, score)
}

func (p *GtfsProxyServer) processWin(w http.ResponseWriter, player string) {
	//p.store.RecordWin(player)
	//w.WriteHeader(http.StatusAccepted)
}
