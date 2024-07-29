package db

import (
	"context"
	"database/sql"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestStore_TransferTx(t *testing.T) {
	store := NewStore(testDB)

	a1 := createRandomAccount(t)
	a2 := createRandomAccount(t)

	n := 5
	amount := int64(100)

	errs := make(chan error)
	results := make(chan TransferTxResult)

	for i := 0; i < n; i++ {
		go func() {
			ctx := context.Background()
			result, err := store.TransferTx(ctx, TransferTxParams{
				FromAccountID: sql.NullInt64{
					Int64: a1.ID,
					Valid: true,
				},
				ToAccountID: sql.NullInt64{
					Int64: a2.ID,
					Valid: true,
				},
				Amount: amount,
			})

			errs <- err
			results <- result
		}()
	}

	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		r := <-results
		require.NotEmpty(t, r)

		transfer := r.Transfer
		require.NotEmpty(t, transfer)

		_, err = store.GetTransfer(context.Background(), transfer.ID)
		require.NoError(t, err)

		require.Equal(t, r.Transfer.FromAccountID.Int64, a1.ID)

	}
}
