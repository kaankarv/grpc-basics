package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpcpractise/weather/api"
	"io"
)

func main() {

	addr := "localhost:8080"
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := api.NewWeatherServiceClient(conn)
	ctx := context.Background()

	resp, err := client.ListCities(ctx, &api.ListCitiesRequest{})
	if err != nil {
		panic(err)
	}
	fmt.Println("cities:")
	for _, city := range resp.Items {
		fmt.Printf("\t%s: %s \n", city.GetCityCode(), city.CityName)
	}

	stream, err := client.QueryWeather(ctx, &api.WeatherRequest{
		CityCode: "ank",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("Weather in Ankara:")
	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)

		}

		go fmt.Printf("\t temperature: %1.f\n", msg.GetTemperature())

		if msg.GetTemperature() > 25 {
			fmt.Printf(" Ankara is too hot ----->")
		} else if msg.GetTemperature() > 15 {
			fmt.Printf(" Ankara's weather is normal ----->")
		} else if msg.GetTemperature() > 10 {
			fmt.Printf(" Ankara is cold ----->")
		} else {
			fmt.Printf(" Ankara is too cold ----->")
		}

	}
	fmt.Println("server stopped sending")

}
