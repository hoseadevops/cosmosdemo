package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/hoseadevops/cosmosdemo/testutil/keeper"
	"github.com/hoseadevops/cosmosdemo/x/cosmosdemo/keeper"
	"github.com/hoseadevops/cosmosdemo/x/cosmosdemo/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.CosmosdemoKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
