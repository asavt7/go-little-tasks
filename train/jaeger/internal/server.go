package internal

import (
	"fmt"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
	"time"

	"io"
	"log"
	"net/http"
)

func InitTracer(service string) (io.Closer, error) {
	cfg := &config.Configuration{ServiceName: service, Sampler: &config.SamplerConfig{Type: "const", Param: 1},
		Reporter: &config.ReporterConfig{LogSpans: true}}
	//tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))

	tracer, err := cfg.InitGlobalTracer(service)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return tracer, err
}

func RunServer(addr string, remoteService string) error {
	client := http.Client{
		Timeout: 10 * time.Second,
	}

	log.Println("Starting server " + addr)
	http.HandleFunc("/", defaultHandler(addr))

	http.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		span := opentracing.GlobalTracer().StartSpan("test method second span")
		defer span.Finish()

		newRequest, err := http.NewRequest("GET", remoteService, nil)
		if err != nil {
			log.Fatal(err)
		}
		// inject
		carrier := opentracing.HTTPHeadersCarrier(newRequest.Header)
		err = opentracing.GlobalTracer().Inject(
			span.Context(),
			opentracing.HTTPHeaders,
			carrier)
		if err != nil {
			log.Fatal(err)
		}

		//do req
		res, err := client.Do(newRequest)
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()
		_, err = io.Copy(writer, res.Body)
		if err != nil {
			return
		}
	})

	http.HandleFunc("/withTracing", func(writer http.ResponseWriter, request *http.Request) {

		carrier := opentracing.HTTPHeadersCarrier(request.Header)
		clientContext, err := opentracing.GlobalTracer().Extract(opentracing.HTTPHeaders, carrier)
		if err != nil {
			log.Fatal(err)
		}

		span := opentracing.GlobalTracer().StartSpan("Span in second service", opentracing.ChildOf(clientContext))
		defer span.Finish()

		fmt.Fprintf(writer, "Hello from %s as second service", addr)
	})

	return http.ListenAndServe(addr, nil)
}

func defaultHandler(addr string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		span := opentracing.GlobalTracer().StartSpan("first span")
		defer span.Finish()
		fmt.Fprintf(w, "Hello from %s", addr)
	}
}
