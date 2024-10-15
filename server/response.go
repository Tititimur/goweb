package server

import (
	"fmt"
	"image"
	"image/jpeg"
	"net/http"

	"github.com/massarakhsh/lik"
	"github.com/massarakhsh/lik/likdom"
)

func ResponseJson(w http.ResponseWriter, content lik.Seter, opts ...interface{}) {
	present := content != nil && content.Count() > 0
	options := lik.BuildSet(opts...)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Cache-control", "private, max-age=0, no-cache")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	rc := int(options.GetInt("rc"))
	if rc != 0 {
	} else if present && content.GetBool("success") {
		rc = 200
	} else {
		rc = 404
	}
	w.WriteHeader(rc)
	if !present {
	} else if options.GetBool("format") {
		serial := content.Format("")
		fmt.Fprint(w, serial)
	} else {
		serial := content.Serialize()
		fmt.Fprint(w, serial)
	}
}

func ResponseHtml(w http.ResponseWriter, content likdom.Domer, opts ...interface{}) {
	present := content != nil && content.GetDataCount() > 0
	options := lik.BuildSet(opts...)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Cache-control", "private, max-age=0, no-cache")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	rc := int(options.GetInt("rc"))
	if rc != 0 {
	} else if present {
		rc = 200
	} else {
		rc = 404
	}
	w.WriteHeader(rc)
	if !present {
	} else if options.GetBool("format") {
		serial := content.ToString()
		fmt.Fprint(w, serial)
	} else {
		serial := content.ToString()
		fmt.Fprint(w, serial)
	}
}

func ResponseFile(w http.ResponseWriter, r *http.Request, file string) bool {
	ok := false
	if match := lik.RegExParse(file, "^(.*?)\\?"); match != nil {
		file = match[1]
	}
	http.ServeFile(w, r, file)
	return ok
}

func ResponseText(w http.ResponseWriter, text string, opts ...interface{}) {
	options := lik.BuildSet(opts...)
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Cache-control", "private, max-age=0, no-cache")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	rc := int(options.GetInt("rc"))
	if rc == 0 {
		if text != "" {
			rc = 200
		} else {
			rc = 404
		}
	}
	w.WriteHeader(rc)
	w.Write([]byte(text))
}

func ResponsePng(w http.ResponseWriter, img image.Image, opts ...interface{}) {
	options := lik.BuildSet(opts...)
	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Cache-control", "private, max-age=0, no-cache")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	rc := int(options.GetInt("rc"))
	if rc == 0 {
		if img != nil {
			rc = 200
		} else {
			rc = 404
		}
	}
	w.WriteHeader(rc)
	if img != nil {
		options := jpeg.Options{Quality: 100}
		jpeg.Encode(w, img, &options)
	}
}
