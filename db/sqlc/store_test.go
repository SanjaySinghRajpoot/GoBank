package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {

	store := NewStore(testDB)

	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	// run n concurrent transfer
	n := 5
	amount := int64(10)

	errs := make(chan error)
	results := make(chan TransferTxResult)

	for i := 0; i < n; i++ {
		go func() { // go tranfer routine
			result, err := store.TransferTx(context.Background(), TranferTxParams{
				FromAccountID: account1.ID,
				ToAccountID:   account2.ID,
				Amount:        amount,
			})

			errs <- err
			results <- result
		}()
	}

	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		tranfer := result.Transfer
		require.NotEmpty(t, tranfer)
		require.Equal(t, account1.ID, tranfer.FromAccountID)
		require.Equal(t, account2.ID, tranfer.ToAccountID)
		require.NotZero(t, tranfer.ID)
		require.NotZero(t, tranfer.CreatedAt)
	}

}
