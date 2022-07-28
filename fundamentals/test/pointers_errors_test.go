package fundamentalstest

import (
	"errors"
	"testing"
)

var ErrInsufficientFunds = errors.New("withdraw insufficienct funds")

type Bitcoin float32

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if w.balance < amount {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}

func TestWallet(t *testing.T) {
	testCases := []struct {
		desc            string
		amountDeposit   Bitcoin
		withdraw        Bitcoin
		expectedBalance Bitcoin
	}{
		{
			desc:            "success deposit 10",
			amountDeposit:   Bitcoin(10),
			expectedBalance: Bitcoin(10),
		},
		{
			desc:            "success withdraw 10",
			amountDeposit:   Bitcoin(10),
			withdraw:        Bitcoin(10),
			expectedBalance: Bitcoin(0),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			wallet := Wallet{}
			wallet.Deposit(tC.amountDeposit)
			err := wallet.Withdraw(tC.withdraw)

			if err != nil {
				t.Error(err.Error())
			}

			balance := wallet.Balance()

			if balance != tC.expectedBalance {
				t.Error("expected:", tC.expectedBalance, "got:", balance)
			}
		})
	}
}
