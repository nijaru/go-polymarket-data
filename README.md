# go-polymarket-data

[![Go Reference](https://pkg.go.dev/badge/github.com/nijaru/go-polymarket-data.svg)](https://pkg.go.dev/github.com/nijaru/go-polymarket-data)
[![CI](https://github.com/nijaru/go-polymarket-data/actions/workflows/ci.yml/badge.svg)](https://github.com/nijaru/go-polymarket-data/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/nijaru/go-polymarket-data)](https://goreportcard.com/report/github.com/nijaru/go-polymarket-data)

> [!WARNING]
> Unofficial, community-maintained SDK. Not extensively tested in production. Use at your own risk.

Go client for the [Polymarket Data API](https://data-api.polymarket.com). Targets the latest stable Go release.

## Features

- **Positions**: open and closed positions with full PNL data.
- **Trades**: trade history with pagination.
- **Activity**: on-chain activity logs.
- **Leaderboards**: trader and builder rankings.
- **Market data**: holders, open interest, live volume.
- **Iterators**: Go 1.23 range-over-function iterators for memory-efficient streaming on all list endpoints.

## Install

```bash
go get github.com/nijaru/go-polymarket-data
```

Requires **Go 1.26.1+**.

## Quickstart

```go
import "github.com/nijaru/go-polymarket-data"

client := polymarketdata.New(polymarketdata.Config{})
ctx := context.Background()

// Leaderboard
entries, err := client.GetLeaderboard(ctx, polymarketdata.LeaderboardParams{
	TimePeriod: "WEEK",
	SortBy:     "PNL",
	Limit:      5,
})

// Iterate positions with auto-pagination
for pos, err := range client.IterPositions(ctx, polymarketdata.PositionParams{
	User: "0x1234...",
}) {
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s: %s tokens @ avg $%s\n", pos.Title, pos.Size, pos.AvgPrice)
}
```

## Pagination Iterators

All list endpoints expose both a slice variant and a Go 1.23 range-over-function iterator:

```go
for pos, err := range client.IterPositions(ctx, polymarketdata.PositionParams{
	User: "0x1234...",
}) {
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s: %s tokens\n", pos.Title, pos.Size)
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

## Error Handling

API errors are returned as `*polymarketdata.APIError` and expose the HTTP status code and body:

```go
if apiErr, ok := err.(*polymarketdata.APIError); ok {
	fmt.Printf("HTTP %d: %s\n", apiErr.StatusCode, apiErr.Message)
}
```

## Contributing

Bug reports and PRs welcome. Please run `make check` before submitting.

## License

MIT
