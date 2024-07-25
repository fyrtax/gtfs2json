package main

import (
	_ "embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"strings"
	"time"
)

//go:embed templates/list_files.gohtml
var listFilesTemplate string

type GtfsProxyStore interface {
	UpdateFile(name string, time int64)
	GetFileDetails(name string) GtfsFile
	SaveFile(name string, file io.Writer) error
	GetFile(name string) (io.Reader, error)
	GetFullFileList() GtfsServerFiles
}

type GtfsProxyServer struct {
	store    GtfsProxyStore
	template *template.Template
	http.Handler
}

type GtfsFile struct {
	Type          string
	Filename      string
	LastChecked   time.Time
	ModTime       time.Time
	Vehicle       string
	FormattedTime string
}

type GtfsServerFiles map[string]GtfsFile

func NewProxyServer(store GtfsProxyStore) (*GtfsProxyServer, error) {
	p := new(GtfsProxyServer)

	p.store = store

	// Load the template
	tmpl, err := template.New("list_files").Parse(listFilesTemplate)
	if err != nil {
		return nil, fmt.Errorf("error loading template: %v", err)
	}
	p.template = tmpl

	router := http.NewServeMux()

	router.Handle("/", http.HandlerFunc(p.indexHandler))
	router.Handle("/ping", http.HandlerFunc(p.pingHandler))

	p.Handler = router

	return p, nil
}

func (p *GtfsProxyServer) pingHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := fmt.Fprint(w, "pong")
	if err != nil {
		fmt.Println("error writing response", err)
	}
}

func (p *GtfsProxyServer) indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)

	if p.handleGtfsRequest(w, r) {
		return
	}

	err := p.template.ExecuteTemplate(w, "list_files", p.store.GetFullFileList())
	if err != nil {
		fmt.Println("error executing template", err)
	}
}

func (p *GtfsProxyServer) handleGtfsRequest(w http.ResponseWriter, r *http.Request) bool {
	fileRequested := strings.TrimPrefix(r.URL.Path, "/")

	// write the requested file to the response
	file, err := p.store.GetFile(fileRequested)

	if err != nil {
		// file not found
		return false
	} else {
		// set header last modified
		fileDetails := p.store.GetFileDetails(fileRequested)
		w.Header().Set("Last-Modified", fileDetails.ModTime.Format(time.RFC1123))

		_, err := io.Copy(w, file)
		if err != nil {
			http.Error(w, "error copying file to response", http.StatusInternalServerError)
		}
		return true
	}
}
