package servs

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go-test/config"
	"go-test/dto"
	"net"
	"net/http"
)

var httpClient = &http.Client{
	Timeout: config.HttpTimeout,
}

func FetchPokemon(ctx context.Context, nameOrID string) (*dto.Pokemon, int, error) {
	url := fmt.Sprintf("%s/%s", config.BaseURL, nameOrID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	req.Header.Set("Accept", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		var netErr net.Error
		if errors.As(err, &netErr) && netErr.Timeout() {
			return nil, http.StatusGatewayTimeout, fmt.Errorf("timeout chamando PokéAPI")
		}
		return nil, http.StatusBadGateway, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, http.StatusNotFound, fmt.Errorf("pokémon não encontrado")
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, http.StatusBadGateway, fmt.Errorf("pokéapi retornou status %d", resp.StatusCode)
	}

	var p dto.Pokemon
	if err := json.NewDecoder(resp.Body).Decode(&p); err != nil {
		return nil, http.StatusBadGateway, fmt.Errorf("falha ao decodificar resposta: %w", err)
	}
	return &p, http.StatusOK, nil
}
