# GeForce-Now-Games

Search if given game name is available on GeForce Now by scraping website or using a JSON static file

## Usage

```
Usage of geforcenow-notifier:
  -method string
    	Search for game with [scrap] or [list] method. (default "scrap")
  -search string
    	Search if given game is available.
```

```bash
# Search "moto" using list method (in https://static.nvidiagrid.net/supported-public-game-list/locales/gfnpc-en-US.json)
➜ go run . -search moto -method list
🎮 	MotoGP™21: Optimized? false	Store: Steam	Status: MAINTENANCE	https://store.steampowered.com/app/1447000
🎮 	MotoGP™19: Optimized? false	Store: Steam	Status: AVAILABLE	https://store.steampowered.com/app/984780
🎮 	MotoGP™ 20: Optimized? false	Store: Steam	Status: AVAILABLE	https://store.steampowered.com/app/1161490
🎮 	MXGP 2019 - The Official Motocross Videogame: Optimized? false	Store: Steam	Status: AVAILABLE	https://store.steampowered.com/app/1018160
🎮 	MXGP 2020 - The Official Motocross Videogame: Optimized? false	Store: Steam	Status: AVAILABLE	https://store.steampowered.com/app/1259800

# Search "moto" usinng scrap method (in https://www.nvidia.com/en-eu/geforce-now/games/)
➜ go run . -search moto -method scrap
🎮 	 MotoGP21
```
