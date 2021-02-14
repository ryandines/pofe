package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Claim is gay
type Claim struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	Proof   string         `json:"proof" yaml:"proof"`
}
