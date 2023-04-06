package keeper

import (
	"github.com/hoseadevops/cosmosdemo/x/cosmosdemo/types"
)

var _ types.QueryServer = Keeper{}
