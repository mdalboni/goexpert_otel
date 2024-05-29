package services

import (
	"github.com/mdalboni/goexpert_3/internals"
	"go.opentelemetry.io/otel/trace"
)

type BaseHttpService struct {
	Client internals.HTTPClient
	Tracer trace.Tracer
}
