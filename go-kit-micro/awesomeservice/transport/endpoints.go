package transport

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	. "github.com/halapastefan/go-kit-micro/awesomeservice"
)

type Endpoints struct {
	Uppercase endpoint.Endpoint
	Count     endpoint.Endpoint
}

func createEndpoints(svc StringService) Endpoints {
	return Endpoints{
		Uppercase: makeUppercaseEndpoint(svc),
		Count:     makeCountEndpoint(svc),
	}
}

func makeUppercaseEndpoint(svc StringService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (response interface{}, err error) {

		req := request.(UppercaseRequest)
		result, err := svc.Uppercase(req.S)

		if err != nil {
			return uppercaseResponse{result, err.Error()}, nil
		}

		return uppercaseResponse{V: result, Err: ""}, nil
	}
}

func makeCountEndpoint(svc StringService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (response interface{}, err error) {

		req := request.(CountRequest)

		result := svc.Count(req.S)

		return countResponse{V: result}, nil
	}
}
