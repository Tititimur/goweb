package server

import (
	"net/http"

	"github.com/massarakhsh/lik"
	"github.com/massarakhsh/lik/stack"
	"github.com/tititimur/goweb/ui"
)

func (it *Server) router(w http.ResponseWriter, r *http.Request) {
	if lik.RegExCompare(r.RequestURI, "\\.(css|js|gif|png|ico|jpg)") {
		ResponseFile(w, r, "."+r.RequestURI)
	} else {
		it.routerUi(w, r)
	}
}

func (it *Server) routerUi(w http.ResponseWriter, r *http.Request) {
	if do := stack.BuildRequest(r); do != nil {
		if page := ui.GenPage(do); page != nil {
			ResponseHtml(w, page)
		}
	}
}
