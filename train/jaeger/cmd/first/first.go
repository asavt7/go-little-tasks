package main

import (
	"github.com/asavt7/go-little-tasks/train/jaeger/internal"
	"io"
	"log"
)

func main() {

	closer, err := internal.InitTracer("first service")
	if err != nil {
		log.Fatal(err)
	}

	defer func(closet io.Closer) {
		err := closer.Close()
		if err != nil {

		}
	}(closer)

	err = internal.RunServer(":9875", "http://localhost:9876/withTracing")
	if err != nil {
		log.Fatal(err)
	}
}
