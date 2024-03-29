package repository

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"path"
	"time"

	"leoseiji.com/multithreading/dto"
)

func GetCEPViaResponse(ctx context.Context, cep string) (*dto.ViaCepResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	cepURL, err := url.Parse("http://viacep.com.br/ws/")
	if err != nil {
		log.Println("error to parse url via cep api", err)
		return nil, err
	}

	cepURL.Path = path.Join(cepURL.Path, url.PathEscape(cep), "json")

	url := cepURL.String()
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		log.Println("get via cep api error", err)
		return nil, err
	}

	res, resErr := http.DefaultClient.Do(req)
	if condition := resErr != nil; condition {
		log.Println("response error", resErr)
		return nil, resErr
	}
	defer res.Body.Close()
	body, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		log.Println("read all error", readErr)
		return nil, readErr
	}

	var viaCepResponse dto.ViaCepResponse
	err = json.Unmarshal(body, &viaCepResponse)
	if err != nil {
		log.Println("unmarshall error", err)
		return nil, err
	}
	return &viaCepResponse, nil
}
