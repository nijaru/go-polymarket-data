package polymarketdata

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDataClient(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.URL.Path {
		case "/positions":
			json.NewEncoder(w).Encode([]Position{
				{Asset: "0x123", Title: "Test Market", Size: "100"},
			})
		case "/trades":
			json.NewEncoder(w).Encode([]DataTrade{
				{Side: "BUY", Size: "50", Price: "0.65"},
			})
		case "/activity":
			json.NewEncoder(w).Encode([]Activity{
				{Type: "TRADE", Size: "10", USDCSize: "6.5"},
			})
		case "/holders":
			json.NewEncoder(w).Encode([]MetaHolder{
				{Token: "0x123", Holders: []Holder{{ProxyWallet: "0xabc", Amount: "100"}}},
			})
		case "/value":
			json.NewEncoder(w).Encode(struct {
				Value string `json:"value"`
			}{Value: "1000.50"})
		case "/v1/leaderboard":
			json.NewEncoder(w).Encode([]TraderLeaderboardEntry{
				{Rank: 1, Username: "trader1", Volume: "1000", PNL: "500"},
			})
		case "/v1/builders/leaderboard":
			json.NewEncoder(w).Encode([]BuilderLeaderboardEntry{
				{Rank: 1, Builder: "0xabc", Volume: "5000", ActiveUsers: 10},
			})
		case "/v1/builders/volume":
			json.NewEncoder(w).Encode([]BuilderVolumeEntry{
				{Builder: "0xabc", Volume: "1000", ActiveUsers: 5},
			})
		case "/traded":
			json.NewEncoder(w).Encode(struct {
				User   string `json:"user"`
				Traded int    `json:"traded"`
			}{User: "0x123", Traded: 42})
		case "/oi":
			json.NewEncoder(w).Encode([]OpenInterest{
				{Market: "0xabc", Value: "10000"},
			})
		case "/live-volume":
			json.NewEncoder(w).Encode(LiveVolume{
				Total:   "50000",
				Markets: []MarketVolume{{Market: "0xabc", Value: "25000"}},
			})
		case "/closed-positions":
			json.NewEncoder(w).Encode([]ClosedPosition{
				{Asset: "0x123", Title: "Closed", RealizedPNL: "50"},
			})
		}
	}))
	defer server.Close()

	client := New(Config{Host: server.URL})
	ctx := context.Background()

	t.Run("GetPositions", func(t *testing.T) {
		items, err := client.GetPositions(ctx, PositionParams{User: "0x123"})
		if err != nil {
			t.Fatalf("GetPositions: %v", err)
		}
		if len(items) != 1 || items[0].Asset != "0x123" {
			t.Errorf("unexpected: %+v", items)
		}
	})

	t.Run("GetTrades", func(t *testing.T) {
		items, err := client.GetTrades(ctx, TradeParams{User: "0x123"})
		if err != nil {
			t.Fatalf("GetTrades: %v", err)
		}
		if len(items) != 1 || items[0].Side != "BUY" {
			t.Errorf("unexpected: %+v", items)
		}
	})

	t.Run("GetActivity", func(t *testing.T) {
		items, err := client.GetActivity(ctx, ActivityParams{User: "0x123"})
		if err != nil {
			t.Fatalf("GetActivity: %v", err)
		}
		if len(items) != 1 || items[0].Type != "TRADE" {
			t.Errorf("unexpected: %+v", items)
		}
	})

	t.Run("GetHolders", func(t *testing.T) {
		items, err := client.GetHolders(ctx, HoldersParams{Markets: []string{"0x123"}})
		if err != nil {
			t.Fatalf("GetHolders: %v", err)
		}
		if len(items) != 1 || items[0].Token != "0x123" {
			t.Errorf("unexpected: %+v", items)
		}
	})

	t.Run("GetTotalValue", func(t *testing.T) {
		val, err := client.GetTotalValue(ctx, "0x123", nil)
		if err != nil {
			t.Fatalf("GetTotalValue: %v", err)
		}
		if val != "1000.50" {
			t.Errorf("unexpected value: %s", val)
		}
	})

	t.Run("GetLeaderboard", func(t *testing.T) {
		items, err := client.GetLeaderboard(ctx, LeaderboardParams{Category: "OVERALL"})
		if err != nil {
			t.Fatalf("GetLeaderboard: %v", err)
		}
		if len(items) != 1 || items[0].Username != "trader1" {
			t.Errorf("unexpected: %+v", items)
		}
	})

	t.Run("GetBuilderLeaderboard", func(t *testing.T) {
		items, err := client.GetBuilderLeaderboard(ctx, BuilderLeaderboardParams{})
		if err != nil {
			t.Fatalf("GetBuilderLeaderboard: %v", err)
		}
		if len(items) != 1 || items[0].Builder != "0xabc" {
			t.Errorf("unexpected: %+v", items)
		}
	})

	t.Run("GetBuilderVolume", func(t *testing.T) {
		items, err := client.GetBuilderVolume(ctx, BuilderVolumeParams{})
		if err != nil {
			t.Fatalf("GetBuilderVolume: %v", err)
		}
		if len(items) != 1 || items[0].Volume != "1000" {
			t.Errorf("unexpected: %+v", items)
		}
	})

	t.Run("GetTradedCount", func(t *testing.T) {
		count, err := client.GetTradedCount(ctx, "0x123")
		if err != nil {
			t.Fatalf("GetTradedCount: %v", err)
		}
		if count != 42 {
			t.Errorf("unexpected count: %d", count)
		}
	})

	t.Run("GetOpenInterest", func(t *testing.T) {
		items, err := client.GetOpenInterest(ctx, OpenInterestParams{Markets: []string{"0xabc"}})
		if err != nil {
			t.Fatalf("GetOpenInterest: %v", err)
		}
		if len(items) != 1 || items[0].Value != "10000" {
			t.Errorf("unexpected: %+v", items)
		}
	})

	t.Run("GetLiveVolume", func(t *testing.T) {
		item, err := client.GetLiveVolume(ctx, 12345)
		if err != nil {
			t.Fatalf("GetLiveVolume: %v", err)
		}
		if item.Total != "50000" {
			t.Errorf("unexpected total: %s", item.Total)
		}
	})

	t.Run("GetClosedPositions", func(t *testing.T) {
		items, err := client.GetClosedPositions(ctx, ClosedPositionParams{User: "0x123"})
		if err != nil {
			t.Fatalf("GetClosedPositions: %v", err)
		}
		if len(items) != 1 || items[0].RealizedPNL != "50" {
			t.Errorf("unexpected: %+v", items)
		}
	})

	t.Run("APIError", func(t *testing.T) {
		errServ := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("bad request"))
		}))
		defer errServ.Close()

		badClient := New(Config{Host: errServ.URL})
		_, err := badClient.GetPositions(ctx, PositionParams{User: "0x123"})
		if err == nil {
			t.Fatal("expected error")
		}
		apiErr, ok := err.(*APIError)
		if !ok {
			t.Fatalf("expected APIError, got %T", err)
		}
		if apiErr.StatusCode != http.StatusBadRequest {
			t.Errorf("unexpected status: %d", apiErr.StatusCode)
		}
	})
}
