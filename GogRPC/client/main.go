package main

import (
	"client/services"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	creds := insecure.NewCredentials()
	cc, err := grpc.Dial("localhost:5000", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	calculatorClient := services.NewCalculatorClient(cc)
	calculatorService := services.NewCalculatorService(calculatorClient)

	// err = calculatorService.Hello("sun")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = calculatorService.Fibonacci(20)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = calculatorService.Average(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	// if err != nil {
	// 	log.Fatal(err)
	// }s

	// err = calculatorService.Sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	err = calculatorService.Sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	if err != nil {
		log.Fatal(err)
	}
}
