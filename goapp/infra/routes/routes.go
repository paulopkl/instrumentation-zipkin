package routes

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel/baggage"
	"go.opentelemetry.io/otel/trace"
)

func GenerateHomeHandler(tracer trace.Tracer) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		ctx := baggage.ContextWithoutBaggage(request.Context())

		// Rotina 1 - Process File
		ctx, processFile := tracer.Start(ctx, "process-file")
		time.Sleep(time.Millisecond * 100)
		processFile.End()
		// Fim Rotina 1

		// Rotina 2 - Fazer Request HTTP
		ctx, httpCall := tracer.Start(ctx, "request-remote-json")
		client := http.Client{Transport: otelhttp.NewTransport(http.DefaultTransport)}
		req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:3000/", nil)
		if err != nil {
			log.Fatal(err)
		}
		res, err := client.Do(req) // chamo a requisição

		if err != nil {
			log.Fatalln(err)
		}
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatalln(err)
		}

		time.Sleep(time.Millisecond * 300) // Atrasa a requisição em 300ms
		httpCall.End()
		// Fim Rotina 2

		// Rotina 3 - Exibir resultado
		ctx, renderContent := tracer.Start(ctx, "render-content")
		time.Sleep(time.Millisecond * 100)
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte(body))
		renderContent.End()
		// Fim Rotina 3
	}
}

// func HomeHandler(writer http.ResponseWriter, request *http.Request) {
// 	ctx := baggage.ContextWithoutBaggage(request.Context())

// 	// rotina 1 - Process File
// 	ctx, processFile := tracer.Start(ctx, "process-file")
// 	time.Sleep(time.Millisecond * 100)
// 	processFile.End()

// 	// rotina 2 - Fazer Request HTTP
// 	ctx, httpCall := tracer.Start(ctx, "request-remote-json")
// 	client := http.Client{Transport: otelhttp.NewTransport(http.DefaultTransport)}
// 	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:3000/", nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	res, err := client.Do(req) // chamo a requisição

// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	body, err := ioutil.ReadAll(res.Body)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	time.Sleep(time.Millisecond * 300)
// 	httpCall.End()

// 	// rotina 3 - Exibir resultado
// 	ctx, renderContent := tracer.Start(ctx, "render-content")
// 	time.Sleep(time.Millisecond * 100)
// 	writer.WriteHeader(http.StatusOK)
// 	writer.Write([]byte(body))
// 	renderContent.End()
// }
