package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpcpractise/weather/api"
	"math/rand"
	"net"
	"time"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	api.RegisterWeatherServiceServer(srv, &myWeatherService{})
	fmt.Println("starting server...")
	panic(srv.Serve(lis))
}

type myWeatherService struct {
	api.UnsafeWeatherServiceServer
}

func (m *myWeatherService) ListCities(ctx context.Context, req *api.ListCitiesRequest) (*api.ListCitiesResponse, error) {
	return &api.ListCitiesResponse{
		Items: []*api.CityEntry{
			&api.CityEntry{CityCode: "ank", CityName: "Ankara"},
			&api.CityEntry{CityCode: "ist", CityName: "İstanbul"},
			&api.CityEntry{CityCode: "izm", CityName: "İzmir"},
		},
	}, nil
}
func (m *myWeatherService) QueryWeather(req *api.WeatherRequest, resp api.WeatherService_QueryWeatherServer) error {
	for {
		err := resp.Send(&api.WeatherResponse{Temperature: rand.Float32()*10 + 9})

		if err != nil {
			break
		}
		time.Sleep(time.Second * 2)

	}
	return nil
}
