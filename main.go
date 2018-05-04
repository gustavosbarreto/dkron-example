package main

import (
	"fmt"
	"net/http"

	"github.com/gustavosbarreto/dkron-go"
	"github.com/gustavosbarreto/dkron-http-executor/pkg/executorhttp"
)

func main() {
	scheduler, err := dkron.NewClient("localhost:8080")
	if err != nil {
		panic(err)
	}

	job, err := scheduler.Jobs.Add(dkron.Job{
		Name:     "651082e8-f2d2-4752-9018-0319949a3b4d",
		Schedule: "@every 10s",
		Command:  "dkron-http-executor",
		Payload: executorhttp.RequestPayload(executorhttp.Request{
			Method: http.MethodGet,
			URL:    "http://localhost:8888/post",
			Body:   []byte(`{ "name": "gustavo" }`),
		}),
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(job)
}
