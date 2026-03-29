package polymarketdata

import "time"

type Position struct {
	ProxyWallet        string `json:"proxyWallet"`
	Asset              string `json:"asset"`
	ConditionID        string `json:"conditionId"`
	Size               string `json:"size"`
	AvgPrice           string `json:"avgPrice"`
	InitialValue       string `json:"initialValue"`
	CurrentValue       string `json:"currentValue"`
	CashPNL            string `json:"cashPnl"`
	PercentPNL         string `json:"percentPnl"`
	TotalBought        string `json:"totalBought"`
	RealizedPNL        string `json:"realizedPnl"`
	PercentRealizedPNL string `json:"percentRealizedPnl"`
	CurPrice           string `json:"curPrice"`
	Redeemable         bool   `json:"redeemable"`
	Mergeable          bool   `json:"mergeable"`
	Title              string `json:"title"`
	Slug               string `json:"slug"`
	Icon               string `json:"icon"`
	EventSlug          string `json:"eventSlug"`
	EventID            string `json:"eventId,omitzero"`
	Outcome            string `json:"outcome"`
	OutcomeIndex       int    `json:"outcomeIndex"`
	OppositeOutcome    string `json:"oppositeOutcome"`
	OppositeAsset      string `json:"oppositeAsset"`
	EndDate            string `json:"endDate,omitzero"`
	NegativeRisk       bool   `json:"negativeRisk"`
}

type ClosedPosition struct {
	ProxyWallet     string `json:"proxyWallet"`
	Asset           string `json:"asset"`
	ConditionID     string `json:"conditionId"`
	AvgPrice        string `json:"avgPrice"`
	TotalBought     string `json:"totalBought"`
	RealizedPNL     string `json:"realizedPnl"`
	CurPrice        string `json:"curPrice"`
	Timestamp       int64  `json:"timestamp"`
	Title           string `json:"title"`
	Slug            string `json:"slug"`
	Icon            string `json:"icon"`
	EventSlug       string `json:"eventSlug"`
	Outcome         string `json:"outcome"`
	OutcomeIndex    int    `json:"outcomeIndex"`
	OppositeOutcome string `json:"oppositeOutcome"`
	OppositeAsset   string `json:"oppositeAsset"`
	EndDate         string `json:"endDate"`
}

type DataTrade struct {
	ProxyWallet           string `json:"proxyWallet"`
	Side                  string `json:"side"`
	Asset                 string `json:"asset"`
	ConditionID           string `json:"conditionId"`
	Size                  string `json:"size"`
	Price                 string `json:"price"`
	Timestamp             int64  `json:"timestamp"`
	Title                 string `json:"title"`
	Slug                  string `json:"slug"`
	Icon                  string `json:"icon"`
	EventSlug             string `json:"eventSlug"`
	Outcome               string `json:"outcome"`
	OutcomeIndex          int    `json:"outcomeIndex"`
	Name                  string `json:"name,omitzero"`
	Pseudonym             string `json:"pseudonym,omitzero"`
	Bio                   string `json:"bio,omitzero"`
	ProfileImage          string `json:"profileImage,omitzero"`
	ProfileImageOptimized string `json:"profileImageOptimized,omitzero"`
	TransactionHash       string `json:"transactionHash"`
}

type Activity struct {
	ProxyWallet           string `json:"proxyWallet"`
	Timestamp             int64  `json:"timestamp"`
	ConditionID           string `json:"conditionId,omitzero"`
	Type                  string `json:"type"`
	Size                  string `json:"size"`
	USDCSize              string `json:"usdcSize"`
	TransactionHash       string `json:"transactionHash"`
	Price                 string `json:"price,omitzero"`
	Asset                 string `json:"asset,omitzero"`
	Side                  string `json:"side,omitzero"`
	OutcomeIndex          *int   `json:"outcomeIndex,omitzero"`
	Title                 string `json:"title,omitzero"`
	Slug                  string `json:"slug,omitzero"`
	Icon                  string `json:"icon,omitzero"`
	EventSlug             string `json:"eventSlug,omitzero"`
	Outcome               string `json:"outcome,omitzero"`
	Name                  string `json:"name,omitzero"`
	Pseudonym             string `json:"pseudonym,omitzero"`
	Bio                   string `json:"bio,omitzero"`
	ProfileImage          string `json:"profileImage,omitzero"`
	ProfileImageOptimized string `json:"profileImageOptimized,omitzero"`
}

type Holder struct {
	ProxyWallet           string `json:"proxyWallet"`
	Bio                   string `json:"bio,omitzero"`
	Asset                 string `json:"asset"`
	Pseudonym             string `json:"pseudonym,omitzero"`
	Amount                string `json:"amount"`
	DisplayUsernamePublic *bool  `json:"displayUsernamePublic,omitzero"`
	OutcomeIndex          int    `json:"outcomeIndex"`
	Name                  string `json:"name,omitzero"`
	ProfileImage          string `json:"profileImage,omitzero"`
	ProfileImageOptimized string `json:"profileImageOptimized,omitzero"`
	Verified              *bool  `json:"verified,omitzero"`
}

type MetaHolder struct {
	Token   string   `json:"token"`
	Holders []Holder `json:"holders"`
}

type OpenInterest struct {
	Market string `json:"market"`
	Value  string `json:"value"`
}

type MarketVolume struct {
	Market string `json:"market"`
	Value  string `json:"value"`
}

type LiveVolume struct {
	Total   string         `json:"total"`
	Markets []MarketVolume `json:"markets"`
}

type BuilderLeaderboardEntry struct {
	Rank        int    `json:"rank,string"`
	Builder     string `json:"builder"`
	Volume      string `json:"volume"`
	ActiveUsers int    `json:"activeUsers"`
	Verified    bool   `json:"verified"`
	BuilderLogo string `json:"builderLogo,omitzero"`
}

type BuilderVolumeEntry struct {
	Timestamp   time.Time `json:"dt"`
	Builder     string    `json:"builder"`
	BuilderLogo string    `json:"builderLogo,omitzero"`
	Verified    bool      `json:"verified"`
	Volume      string    `json:"volume"`
	ActiveUsers int       `json:"activeUsers"`
	Rank        int       `json:"rank,string"`
}

type TraderLeaderboardEntry struct {
	Rank         int    `json:"rank,string"`
	ProxyWallet  string `json:"proxyWallet"`
	Username     string `json:"userName,omitzero"`
	Volume       string `json:"vol"`
	PNL          string `json:"pnl"`
	ProfileImage string `json:"profileImage,omitzero"`
	XUsername    string `json:"xUsername,omitzero"`
	Verified     bool   `json:"verifiedBadge,omitzero"`
}

type PositionParams struct {
	User          string   `url:"user"`
	Markets       []string `url:"market,omitzero"`
	EventIDs      []string `url:"eventID,omitzero"`
	SizeThreshold string   `url:"sizeThreshold,omitzero"`
	Redeemable    *bool    `url:"redeemable,omitzero"`
	Mergeable     *bool    `url:"mergeable,omitzero"`
	Limit         int      `url:"limit,omitzero"`
	Offset        int      `url:"offset,omitzero"`
	SortBy        string   `url:"sortBy,omitzero"`
	SortDirection string   `url:"sortDirection,omitzero"`
	Title         string   `url:"title,omitzero"`
}

type ClosedPositionParams struct {
	User          string   `url:"user"`
	Markets       []string `url:"market,omitzero"`
	EventIDs      []string `url:"eventID,omitzero"`
	Title         string   `url:"title,omitzero"`
	Limit         int      `url:"limit,omitzero"`
	Offset        int      `url:"offset,omitzero"`
	SortBy        string   `url:"sortBy,omitzero"`
	SortDirection string   `url:"sortDirection,omitzero"`
}

type TradeParams struct {
	User      string   `url:"user,omitzero"`
	Markets   []string `url:"market,omitzero"`
	EventIDs  []string `url:"eventID,omitzero"`
	Limit     int      `url:"limit,omitzero"`
	Offset    int      `url:"offset,omitzero"`
	TakerOnly *bool    `url:"takerOnly,omitzero"`
	Side      string   `url:"side,omitzero"`
}

type ActivityParams struct {
	User          string   `url:"user"`
	Markets       []string `url:"market,omitzero"`
	EventIDs      []string `url:"eventID,omitzero"`
	ActivityTypes []string `url:"type,omitzero"`
	Limit         int      `url:"limit,omitzero"`
	Offset        int      `url:"offset,omitzero"`
	Start         int64    `url:"start,omitzero"`
	End           int64    `url:"end,omitzero"`
	SortBy        string   `url:"sortBy,omitzero"`
	SortDirection string   `url:"sortDirection,omitzero"`
	Side          string   `url:"side,omitzero"`
}

type HoldersParams struct {
	Markets    []string `url:"market"`
	Limit      int      `url:"limit,omitzero"`
	MinBalance int      `url:"minBalance,omitzero"`
}

type OpenInterestParams struct {
	Markets []string `url:"market,omitzero"`
}

type LeaderboardParams struct {
	Category   string `url:"category,omitzero"`
	TimePeriod string `url:"timePeriod,omitzero"`
	SortBy     string `url:"orderBy,omitzero"`
	Limit      int    `url:"limit,omitzero"`
	Offset     int    `url:"offset,omitzero"`
	User       string `url:"user,omitzero"`
	UserName   string `url:"userName,omitzero"`
}

type BuilderLeaderboardParams struct {
	TimePeriod string `url:"timePeriod,omitzero"`
	Limit      int    `url:"limit,omitzero"`
	Offset     int    `url:"offset,omitzero"`
}

type BuilderVolumeParams struct {
	TimePeriod string `url:"timePeriod,omitzero"`
}
