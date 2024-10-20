package db

import (
	"context"
	"simplebank/simplebank/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreatAccount(t *testing.T) {
	arg := CreateAccountParams{
		Username: util.RandomUsername(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, arg.Username, account.Username)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
}
