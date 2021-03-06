package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"

	dnTypes "github.com/dfinance/dnode/helpers/types"
	"github.com/dfinance/dnode/x/orders"
)

// MatcherResult stores matcher results.
type MatcherResult struct {
	// MarketID
	MarketID dnTypes.ID
	// Total number of active bid orders
	BidOrdersCount int
	// Total number of active ask orders
	AskOrdersCount int
	// PQCurve crossing point data
	ClearanceState ClearanceState
	// Sum of matched bid orders volume
	MatchedBidVolume sdk.Dec
	// Sum of matched ask orders volume
	MatchedAskVolume sdk.Dec
	// Fully / partially filled orders with some meta
	OrderFills orders.OrderFills
}

func (r MatcherResult) ShortString() string {
	b := strings.Builder{}
	b.WriteString(fmt.Sprintf("MatcherResult for marketID %q:\n", r.MarketID.String()))
	b.WriteString(fmt.Sprintf("  Bid/Ask orders count:   %d/%d\n", r.BidOrdersCount, r.AskOrdersCount))
	b.WriteString(fmt.Sprintf("  ClearanceState.Price:   %s\n", r.ClearanceState.Price.String()))
	b.WriteString(fmt.Sprintf("  ClearanceState.ProRata: %s\n", r.ClearanceState.ProRata.String()))
	b.WriteString(fmt.Sprintf("  OrderFillsCount:        %d", len(r.OrderFills)))

	return b.String()
}

// Strings returns multi-line text object representation.
func (r MatcherResult) String() string {
	b := strings.Builder{}
	b.WriteString("MatcherResult:\n")
	b.WriteString(fmt.Sprintf("  MarketID:         %s\n", r.MarketID.String()))
	b.WriteString(fmt.Sprintf("  BidOrdersCount:   %d\n", r.BidOrdersCount))
	b.WriteString(fmt.Sprintf("  AskOrdersCount:   %d\n", r.AskOrdersCount))
	b.WriteString(fmt.Sprintf("  MatchedBidVolume: %s\n", r.MatchedBidVolume.String()))
	b.WriteString(fmt.Sprintf("  MatchedAskVolume: %s\n", r.MatchedAskVolume.String()))
	b.WriteString(r.ClearanceState.String())
	b.WriteString("OrderFills:\n")
	b.WriteString(r.OrderFills.String())

	return b.String()
}

// MatcherResult slice type.
type MatcherResults []MatcherResult
