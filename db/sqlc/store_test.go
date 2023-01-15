package db

// func TestTransferTx(t *testing.T) {

// 	store := NewStore(testDB)

// 	account1 := createRandomAccount(t)
// 	account2 := createRandomAccount(t)

// 	// run n concurrent transfer
// 	n := 5
// 	amount := int64(10)

// 	errs := make(chan error)
// 	// results := make(chan TransferTxResult)

// 	for i := 0; i < n; i++ {

// 		fromAccountID := account1.ID
// 		toAccountID := account2.ID

// 		if i%2 == 1 {
// 			fromAccountID = account2.ID
// 			toAccountID = account1.ID
// 		}

// 		txName := fmt.Sprintf("tx -- %d", i+1)
// 		go func() { // go tranfer routine
// 			ctx := context.WithValue(context.Background(), txKey, txName)
// 			_, err := store.TransferTx(ctx, TranferTxParams{
// 				FromAccountID: fromAccountID,
// 				ToAccountID:   toAccountID,
// 				Amount:        amount,
// 			})

// 			errs <- err
// 		}()
// 	}

// 	for i := 0; i < n; i++ {
// 		err := <-errs
// 		require.NoError(t, err)
// 	}

// 	// check the final updated balance
// 	updatedAccount1, err := store.GetAccount(context.Background(), account1.ID)
// 	require.NoError(t, err)

// 	updatedAccount2, err := store.GetAccount(context.Background(), account2.ID)
// 	require.NoError(t, err)

// 	fmt.Println(">> after:", updatedAccount1.Balance, updatedAccount2.Balance)

// 	require.Equal(t, account1.Balance-int64(n)*amount, updatedAccount1.Balance)
// 	require.Equal(t, account2.Balance+int64(n)*amount, updatedAccount2.Balance)
// }
