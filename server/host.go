package server

import (
	"fmt"
	"net/http"
)

type Server struct {
	UrlAddress string
	HostServer *http.Server
}

func Start() {
	it := &Server{}
	it.bindtHttp("http", "localhost", "8889")
}

func (it *Server) bindtHttp(proto string, address string, port string) {
	it.UrlAddress = proto + "://" + address + ":" + port
	url := address + ":" + port
	mux := http.NewServeMux()
	mux.HandleFunc("/", it.router)
	it.HostServer = &http.Server{
		Addr:    url,
		Handler: mux,
	}
	if err := it.HostServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		it.HostServer.Addr = "0.0.0.0:" + port
		if err := it.HostServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Bind http error: %s\n", fmt.Sprint(err))
		}
	}
}
