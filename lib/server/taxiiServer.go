// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package server

import (
	"encoding/json"
	"html/template"
	"net/http"
	"strings"

	"github.com/freetaxii/freetaxii-server/lib/headers"
	"github.com/freetaxii/libstix2/defs"
	"github.com/gologme/log"
)

/*
DiscoveryHandler - This method will handle all Discovery requests
*/
func (s *ServerHandlerType) DiscoveryHandler(w http.ResponseWriter, r *http.Request) {
	log.Infoln("INFO: Found Discovery request from", r.RemoteAddr, "at", r.RequestURI)
	s.baseHandler(w, r)
}

/*
APIRootHandler - This method will handle all API Root requests
*/
func (s *ServerHandlerType) APIRootHandler(w http.ResponseWriter, r *http.Request) {
	log.Infoln("INFO: Found API Root request from", r.RemoteAddr, "at", r.RequestURI)
	s.baseHandler(w, r)
}

/*
CollectionsHandler - This method will handle all Collections requests
*/
func (s *ServerHandlerType) CollectionsHandler(w http.ResponseWriter, r *http.Request) {
	log.Infoln("INFO: Found Collections request from", r.RemoteAddr, "at", r.RequestURI)
	s.baseHandler(w, r)
}

/*
CollectionHandler - This method will handle all Collection requests
*/
func (s *ServerHandlerType) CollectionHandler(w http.ResponseWriter, r *http.Request) {
	log.Infoln("INFO: Found Collection request from", r.RemoteAddr, "at", r.RequestURI)
	s.baseHandler(w, r)
}

/*
baseHandler - This method handles all requests for the following TAXII
media type responses: Discovery, API-Root, Collections, and Collection
*/
func (s *ServerHandlerType) baseHandler(w http.ResponseWriter, r *http.Request) {
	var httpHeaderAccept string
	var taxiiHeader headers.HttpHeaderType

	// If trace is enabled in the logger, than decode the HTTP Request to the log
	if log.GetLevel("trace") {
		taxiiHeader.DebugHttpRequest(r)
	}

	// --------------------------------------------------
	// Encode outgoing response message
	// --------------------------------------------------

	httpHeaderAccept = r.Header.Get("Accept")

	// Setup JSON stream encoder
	j := json.NewEncoder(w)

	// Set header for TLS
	w.Header().Add("Strict-Transport-Security", "max-age=86400; includeSubDomains")

	if strings.Contains(httpHeaderAccept, defs.TAXII_MEDIA_TYPE) {
		w.Header().Set("Content-Type", defs.CONTENT_TYPE_TAXII)
		w.WriteHeader(http.StatusOK)
		j.Encode(s.Resource)

	} else if strings.Contains(httpHeaderAccept, "application/json") {
		w.Header().Set("Content-Type", defs.CONTENT_TYPE_JSON)
		w.WriteHeader(http.StatusOK)
		j.SetIndent("", "    ")
		j.Encode(s.Resource)

	} else if s.HTMLEnabled == true && strings.Contains(httpHeaderAccept, "text/html") {
		w.Header().Set("Content-Type", defs.CONTENT_TYPE_HTML)
		w.WriteHeader(http.StatusOK)

		// ----------------------------------------------------------------------
		// Setup HTML Template
		// ----------------------------------------------------------------------
		htmlTemplateResource := template.Must(template.ParseFiles(s.HTMLTemplate))
		htmlTemplateResource.Execute(w, s)

	} else {
		w.WriteHeader(http.StatusUnsupportedMediaType)
	}
}
