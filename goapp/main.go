package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/paulopkl/otel-go/infra/opentel"
	routes "github.com/paulopkl/otel-go/infra/routes"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"

	"go.opentelemetry.io/otel/trace"
)

var tracer trace.Tracer

func main() {
	ot := opentel.NewOpenTel()
	ot.ServiceName = "GoApp"
	ot.ServiceVersion = "0.1"
	ot.ExporterEndpoint = "http://localhost:9411/api/v2/spans"
	tracer = ot.GetTracer()

	repositoryHomeHandler := routes.GenerateHomeHandler(tracer)

	router := mux.NewRouter()

	router.Use(otelmux.Middleware(ot.ServiceName))

	router.HandleFunc("/", repositoryHomeHandler)
	http.ListenAndServe(":8888", router)
}
