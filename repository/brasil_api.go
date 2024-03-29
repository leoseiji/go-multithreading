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

func GetBrasilAPIResponse(ctx context.Context, cep string) (*dto.BrasilResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	cepURL, err := url.Parse("https://brasilapi.com.br/api/cep/v1/")
	if err != nil {
		log.Println("error to parse url brasil api", err)
		return nil, err
	}

	cepURL.Path = path.Join(cepURL.Path, url.PathEscape(cep))

	url := cepURL.String()
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		log.Println("get brasil api error", err)
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

	var brasilResponse dto.BrasilResponse
	err = json.Unmarshal(body, &brasilResponse)
	if err != nil {
		log.Println("unmarshall error", err)
		return nil, err
	}
	return &brasilResponse, nil
}
