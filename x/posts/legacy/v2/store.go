package v2

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/desmos-labs/desmos/v4/x/posts/types"
	subspacestypes "github.com/desmos-labs/desmos/v4/x/subspaces/types"
)

type SubspacesKeeper interface {
	IterateSubspaces(ctx sdk.Context, fn func(subspaces subspacestypes.Subspace) (stop bool))
}

// MigrateStore performs in-place store migrations from v1 to v2
// The things done here are:
// 1. setting up the next post id key for existing subspaces
// 2. setting up the module parameters
func MigrateStore(ctx sdk.Context, storeKey sdk.StoreKey, paramsSubspace paramstypes.Subspace, sk SubspacesKeeper) error {
	store := ctx.KVStore(storeKey)

	// Set the next post id for all the subspaces
	sk.IterateSubspaces(ctx, func(subspaces subspacestypes.Subspace) (stop bool) {
		store.Set(types.NextPostIDStoreKey(subspaces.ID), types.GetPostIDBytes(1))
		return false
	})

	// Set the module params
	params := types.DefaultParams()
	paramsSubspace.SetParamSet(ctx, &params)

	return nil
}
