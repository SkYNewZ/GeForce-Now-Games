package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const gameListURL = "https://static.nvidiagrid.net/supported-public-game-list/locales/gfnpc-en-US.json"

// Game on GeForce Now available games list
type Game struct {
	ID               int      `json:"id"`
	Title            string   `json:"title"`
	SortName         string   `json:"sortName"`
	IsFullyOptimized bool     `json:"isFullyOptimized"`
	SteamURL         string   `json:"steamUrl"`
	Store            string   `json:"store"`
	Publisher        string   `json:"publisher"`
	Genres           []string `json:"genres"`
	Status           string   `json:"status"`
}

// ListGames return each games available on GeForce Now
func ListGames(_ context.Context) ([]*Game, error) {
	c := &http.Client{Timeout: time.Second * 10}
	resp, err := c.Get(gameListURL)
	if err != nil {
		return nil, fmt.Errorf("error requesting game list: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("recevied non 200 status code")
	}

	if resp.Body == http.NoBody {
		return nil, fmt.Errorf("empty body received")
	}

	var games []*Game
	if err := json.NewDecoder(resp.Body).Decode(&games); err != nil {
		return nil, fmt.Errorf("unable to read body: %w", err)
	}

	return games, nil
}
