package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/gocolly/colly/v2"
)

const (
	// Separators
	steam = " (Steam)"
	epic  = " (Epic Games Store)"

	// Methods
	scrap = "scrap"
	list  = "list"
)

func main() {
	search := flag.String("search", "", "Search if given game is available.")
	method := flag.String("method", scrap, fmt.Sprintf("Search for game with [%s] or [%s] method.", scrap, list))
	flag.Parse()

	if *search == "" {
		flag.PrintDefaults()
		exit(fmt.Errorf("\nmissing required -search flag"), 2)
	}

	var f func(query string) error
	switch *method {
	case scrap:
		f = scrapAndSearch
	case list:
		f = listAndSearch
	default:
		flag.PrintDefaults()
		exit(fmt.Errorf("\nplease specify a valid search method [%s] or [%s]", scrap, list), 2)
	}

	exit(f(*search), 1)
}

func listAndSearch(query string) error {
	games, err := ListGames(context.Background())
	if err != nil {
		return err
	}

	for _, game := range games {
		if strings.Contains(strings.ToLower(game.Title), strings.ToLower(query)) {
			fmt.Printf("🎮 \t%s: Optimized? %t\tStore: %s\tStatus: %s\t%s\n", game.Title, game.IsFullyOptimized, game.Store, game.Status, game.SteamURL)
		}
	}

	return nil
}

func scrapAndSearch(query string) error {
	c := colly.NewCollector(colly.AllowedDomains("www.nvidia.com"))

	var htmlNewGamesListRaw string
	c.OnHTML(`div.text-center.tab-text-center div.body-text.description.color-body-copy p`, func(element *colly.HTMLElement) {
		// We have some paragraphs, have to parse them with CSS
		if element.Attr("style") == "text-align: center;" && htmlNewGamesListRaw == "" {
			htmlNewGamesListRaw = element.Text
		}
	})

	if err := c.Visit("https://www.nvidia.com/en-eu/geforce-now/games/"); err != nil {
		return err
	}

	s := strings.ReplaceAll(htmlNewGamesListRaw, steam, "<CUT_HERE>")
	s = strings.ReplaceAll(s, epic, "<CUT_HERE>")
	for _, game := range strings.Split(s, "<CUT_HERE>") {
		if strings.Contains(strings.ToLower(game), strings.ToLower(query)) {
			fmt.Printf("🎮 \t%s\n", game)
		}
	}

	return nil
}

func exit(err error, code int) {
	if err == nil {
		return
	}

	_, _ = fmt.Fprint(os.Stderr, err)
	os.Exit(code)
}
