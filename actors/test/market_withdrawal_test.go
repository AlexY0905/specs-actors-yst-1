package test_test

import (
	"context"
	"testing"

	"github.com/filecoin-project/go-state-types/big"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/AlexY0905/specs-actors-yst-1/actors/builtin"
	"github.com/AlexY0905/specs-actors-yst-1/actors/builtin/market"
	vm "github.com/AlexY0905/specs-actors-yst-1/support/vm"
)

func TestMarketWithdraw(t *testing.T) {
	ctx := context.Background()
	v := vm.NewVMWithSingletons(ctx, t)
	initialBalance := big.Mul(big.NewInt(6), big.NewInt(1e18))
	addrs := vm.CreateAccounts(ctx, t, v, 1, initialBalance, 93837778)
	caller := addrs[0]

	// add market collateral for clients and miner
	collateral := big.Mul(big.NewInt(3), vm.FIL)
	vm.ApplyOk(t, v, caller, builtin.StorageMarketActorAddr, collateral, builtin.MethodsMarket.AddBalance, &caller)

	a, found, err := v.GetActor(caller)
	require.NoError(t, err)
	require.True(t, found)
	assert.Equal(t, big.Sub(initialBalance, collateral), a.Balance)

	// withdraw collateral from market
	params := &market.WithdrawBalanceParams{
		ProviderOrClientAddress: caller,
		Amount:                  collateral,
	}
	vm.ApplyOk(t, v, caller, builtin.StorageMarketActorAddr, big.Zero(), builtin.MethodsMarket.WithdrawBalance, params)

	a, found, err = v.GetActor(caller)
	require.NoError(t, err)
	require.True(t, found)
	assert.Equal(t, initialBalance, a.Balance)
}
