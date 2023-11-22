// Package traefik_header_class_marker create header from other header's value
package traefik_header_class_marker

import (
	"context"
	"fmt"
	"net/http"
	"text/template"
)

// Config the plugin configuration.
type Config struct {
	SourceHeader            string              `json:"sourceHeader"`
	DestinationHeaderPrefix string              `json:"destinationHeaderPrefix"`
	Classes                 map[string][]string `json:"classes,omitempty"`
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{
		SourceHeader:            "x-jwt-preferred_username",
		DestinationHeaderPrefix: "x-throttling-class-",
		Classes:                 make(map[string][]string),
	}
}

// ClassMarker a ClassMarker plugin.
type ClassMarker struct {
	next                    http.Handler
	sourceHeader            string
	destinationHeaderPrefix string
	classes                 map[string][]string
	name                    string
	template                *template.Template
}

// New created a new ClassMarker plugin.
func New(_ context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	if len(config.Classes) == 0 {
		return nil, fmt.Errorf("classes cannot be empty")
	}

	return &ClassMarker{
		sourceHeader:            config.SourceHeader,
		destinationHeaderPrefix: config.DestinationHeaderPrefix,
		classes:                 config.Classes,
		next:                    next,
		name:                    name,
		template:                template.New("marker").Delims("[[", "]]"),
	}, nil
}

func (a *ClassMarker) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	for _, value := range req.Header.Values(a.sourceHeader) {
		for class, subjects := range a.classes {
            for _, subject := range subjects {
                if subject == value {
                    headerName := a.destinationHeaderPrefix + class
                    req.Header.Set(headerName, value)
                    break
                }
            }
		}
	}

	a.next.ServeHTTP(rw, req)
}
