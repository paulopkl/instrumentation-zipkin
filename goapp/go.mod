module github.com/paulopkl/otel-go

go 1.22

require (
	github.com/gorilla/mux v1.8.1
	go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux v0.51.0
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.51.0
	go.opentelemetry.io/otel v1.26.0
	go.opentelemetry.io/otel/exporters/zipkin v1.26.0
	go.opentelemetry.io/otel/sdk v1.26.0
	go.opentelemetry.io/otel/trace v1.26.0
)

require (
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/go-logr/logr v1.4.1 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/openzipkin/zipkin-go v0.4.2 // indirect
	go.opentelemetry.io/otel/metric v1.26.0 // indirect
	golang.org/x/sys v0.19.0 // indirect
)
