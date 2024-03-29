package main

import (
	"context"
	"fmt"
	"time"

	"leoseiji.com/multithreading/dto"
	"leoseiji.com/multithreading/repository"
)

type addressAPI interface {
	GetResponse(context.Context, string) (interface{}, error)
}

func getBrasilAddressAPI(cepParameter string) *dto.BrasilResponse {
	response, err := repository.GetBrasilAPIResponse(context.Background(), cepParameter)
	if err != nil {
		fmt.Println("Erro ao obter resposta Brasil API: ", err)
		return nil
	}
	return response
}

func getViaCepAddressAPI(cepParameter string) *dto.ViaCepResponse {
	response, err := repository.GetCEPViaResponse(context.Background(), cepParameter)
	if err != nil {
		fmt.Println("Erro ao obter resposta Via Cep: ", err)
		return nil
	}
	return response
}

func main() {
	cepParameter := "01153000"

	c1 := make(chan *dto.BrasilResponse)
	c2 := make(chan *dto.ViaCepResponse)

	go func() {
		c1 <- getBrasilAddressAPI(cepParameter)
	}()

	go func() {
		c2 <- getViaCepAddressAPI(cepParameter)
	}()

	select {
	case responseBrasilAPI := <-c1:
		fmt.Println("API mais rápida: Brasil API")
		fmt.Printf("API Resposta: %+v\n", responseBrasilAPI)
	case responseViaAPI := <-c2:
		fmt.Println("API mais rápida: Via API")
		fmt.Printf("API Resposta: %+v\n", responseViaAPI)
	case <-time.After(1 * time.Second):
		fmt.Println("Erro: timeout")
	}
}
