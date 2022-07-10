package keeper

import (
	"github.com/marcoeg/tokenfactory/x/tokenfactory/types"
)

var _ types.QueryServer = Keeper{}
