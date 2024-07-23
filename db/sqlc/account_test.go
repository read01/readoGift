package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"readoGift/util"
	"testing"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	a1 := createRandomAccount(t)
	a2, err := testQueries.GetAccount(context.Background(), a1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, a2)

	require.Equal(t, a1.ID, a2.ID)
	require.Equal(t, a1.Owner, a2.Owner)
	require.Equal(t, a1.Balance, a2.Balance)
	require.Equal(t, a1.Currency, a2.Currency)

}

func TestQueries_UpdateAccount(t *testing.T) {
	a1 := createRandomAccount(t)
	arg := UpdateAccountParams{
		ID:      a1.ID,
		Balance: util.RandomMoney(),
	}
	err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
}

func TestQueries_DeleteAccount(t *testing.T) {
	a1 := createRandomAccount(t)
	a2, err := testQueries.DeleteAccount(context.Background(), a1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, a2)

	require.Equal(t, a1.ID, a2.ID)
	require.Equal(t, a1.Owner, a2.Owner)
	require.Equal(t, a1.Balance, a2.Balance)
	require.Equal(t, a1.Currency, a2.Currency)
}

func TestQueries_ListAccount(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	arg := ListAccountParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccount(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

}
