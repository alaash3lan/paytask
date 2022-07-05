package account

import (
	"context"
	"errors"
)

type Transaction struct {
	From   string  `json:"from"`
	To     string  `json:"to"`
	Amount float64 `json:"amount"`
}

//transfer
func (t *Transaction) Transfer(c context.Context) error {
	accounts := c.Value("db").(map[string]*Account)
	if t.Amount == 0 {
		return errors.New("cannot transfer 0 balance")
	}
	if accounts[t.From].Balance < t.Amount {
		return errors.New("failed to transfer because you don't have enough balance for this transaction")
	}
	accounts[t.From].Balance -= t.Amount
	accounts[t.To].Balance += t.Amount
	return nil
}
