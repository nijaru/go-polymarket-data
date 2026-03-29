# go-polymarket-data

Go client for the [Polymarket Data API](https://data-api.polymarket.com).

## Install

```bash
go get github.com/nijaru/go-polymarket-data
```

## Usage

```go
package main

import (
	"context"
	"fmt"
	"log"

	polymarketdata "github.com/nijaru/go-polymarket-data/polymarketdata"
)

func main() {
	client := polymarketdata.New(polymarketdata.Config{})
	ctx := context.Background()

	// Leaderboard
	entries, err := client.GetLeaderboard(ctx, polymarketdata.LeaderboardParams{
		TimePeriod: "WEEK",
		SortBy:     "PNL",
		Limit:      5,
	})
	if err != nil {
		log.Fatal(err)
	}
	for _, e := range entries {
		fmt.Printf("#%d %s: $%s PNL\n", e.Rank, e.Username, e.PNL)
	}

	// Iterate positions with auto-pagination
	for pos, err := range client.IterPositions(ctx, polymarketdata.PositionParams{
		User: "0x1234...",
	}) {
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s: %s tokens @ avg $%s\n", pos.Title, pos.Size, pos.AvgPrice)
	}
}
```

## API Coverage

| Method | Description |
|--------|-------------|
| `GetPositions` / `IterPositions` | Open positions for a user |
| `GetClosedPositions` / `IterClosedPositions` | Historical closed positions |
| `GetTotalValue` | Total USDC value of a user's positions |
| `GetTrades` / `IterTrades` | Trade history |
| `GetActivity` / `IterActivity` | On-chain activity logs |
| `GetHolders` | Top token holders for markets |
| `GetTradedCount` | Unique markets traded by a user |
| `GetOpenInterest` | Open interest per market |
| `GetLiveVolume` | Live trading volume per event |
| `GetLeaderboard` / `IterLeaderboard` | Trader rankings |
| `GetBuilderLeaderboard` / `IterBuilderLeaderboard` | Builder rankings |
| `GetBuilderVolume` | Daily builder volume data |

## License

MIT
