/*
Copyright 2021 Adevinta
*/

// Code generated by goagen v1.4.3, DO NOT EDIT.
//
// API "vulcan-results": Application Controllers
//
// Command:
// $ goagen
// --design=github.com/adevinta/vulcan-results/design
// --out=/Users/manel.montilla/develop/vulcan-results
// --version=v1.4.3

package app

import (
	"context"
	"github.com/goadesign/goa"
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Encoder.Register(goa.NewGobEncoder, "application/gob", "application/x-gob")
	service.Encoder.Register(goa.NewXMLEncoder, "application/xml")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// ResultsController is the controller interface for the Results actions.
type ResultsController interface {
	goa.Muxer
	GetLog(*GetLogResultsContext) error
	GetReport(*GetReportResultsContext) error
	Raw(*RawResultsContext) error
	Report(*ReportResultsContext) error
}

// MountResultsController "mounts" a Results resource controller on the given service.
func MountResultsController(service *goa.Service, ctrl ResultsController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewGetLogResultsContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.GetLog(rctx)
	}
	service.Mux.Handle("GET", "/v1/logs/:date/:scan/:check", ctrl.MuxHandler("getLog", h, nil))
	service.LogInfo("mount", "ctrl", "Results", "action", "GetLog", "route", "GET /v1/logs/:date/:scan/:check")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewGetReportResultsContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.GetReport(rctx)
	}
	service.Mux.Handle("GET", "/v1/reports/:date/:scan/:check", ctrl.MuxHandler("getReport", h, nil))
	service.LogInfo("mount", "ctrl", "Results", "action", "GetReport", "route", "GET /v1/reports/:date/:scan/:check")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewRawResultsContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*RawPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Raw(rctx)
	}
	service.Mux.Handle("POST", "/v1/raw", ctrl.MuxHandler("raw", h, unmarshalRawResultsPayload))
	service.LogInfo("mount", "ctrl", "Results", "action", "Raw", "route", "POST /v1/raw")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewReportResultsContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*ReportPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Report(rctx)
	}
	service.Mux.Handle("POST", "/v1/report", ctrl.MuxHandler("report", h, unmarshalReportResultsPayload))
	service.LogInfo("mount", "ctrl", "Results", "action", "Report", "route", "POST /v1/report")
}

// unmarshalRawResultsPayload unmarshals the request body into the context request data Payload field.
func unmarshalRawResultsPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &rawPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalReportResultsPayload unmarshals the request body into the context request data Payload field.
func unmarshalReportResultsPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &reportPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// HealthcheckController is the controller interface for the Healthcheck actions.
type HealthcheckController interface {
	goa.Muxer
	Show(*ShowHealthcheckContext) error
}

// MountHealthcheckController "mounts" a Healthcheck resource controller on the given service.
func MountHealthcheckController(service *goa.Service, ctrl HealthcheckController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowHealthcheckContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	service.Mux.Handle("GET", "/healthcheck", ctrl.MuxHandler("show", h, nil))
	service.LogInfo("mount", "ctrl", "Healthcheck", "action", "Show", "route", "GET /healthcheck")
}
