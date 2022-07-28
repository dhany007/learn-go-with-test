package fundamentalstest

import (
	"reflect"
	"testing"
)

func Reduce[A any](collection []A, accumulator func(A, A) A, initalValue A) A {
	result := initalValue
	for _, v := range collection {
		result = accumulator(result, v)
	}

	return result
}

// SumGeneric calculates the total from a slices of numbers
func SumGeneric(numbers []int) int {
	add := func(acc, x int) int {
		return acc + x
	}

	return Reduce(numbers, add, 0)
}

// SumAllTails calculates the sum of all but the first number given a collection of slices
func SumAllTailsGeneric(numbers ...[]int) []int {
	sumTail := func(acc, x []int) []int {
		if len(x) == 0 {
			return append(acc, 0)
		} else {
			tail := x[1:]
			return append(acc, SumGeneric(tail))
		}
	}

	return Reduce(numbers, sumTail, []int{})
}

func TestSumGeneric(t *testing.T) {
	testCases := []struct {
		desc     string
		numbers  []int
		expected int
	}{
		{
			desc:     "sum numbers 1",
			numbers:  []int{1, 2, 3},
			expected: 6,
		},
		{
			desc:     "sum numbers 2",
			numbers:  []int{5, 2, 3, 4},
			expected: 14,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result := SumGeneric(tC.numbers)

			if tC.expected != result {
				t.Error("expected:", tC.expected, "got:", result)
			}
		})
	}
}

func TestSumTaillAllGeneric(t *testing.T) {
	testCases := []struct {
		desc     string
		numbers1 []int
		numbers2 []int
		expected []int
	}{
		{
			desc:     "sums slices 1",
			numbers1: []int{3, 2, 3},
			numbers2: []int{3, 5, 6},
			expected: []int{5, 11},
		},
		{
			desc:     "sums slices 2",
			numbers1: []int{1, 3},
			numbers2: []int{3, 5, 6},
			expected: []int{3, 11},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result := SumAllTailsGeneric(tC.numbers1, tC.numbers2)

			if !reflect.DeepEqual(result, tC.expected) {
				t.Error("got", result, "want", tC.expected)
			}
		})
	}
}

func TestReduce(t *testing.T) {
	t.Run("multiplication of all elements", func(t *testing.T) {
		multipy := func(x, y int) int {
			return x * y
		}
		AssertEqual(t, Reduce([]int{1, 2, 3, 4}, multipy, 1), 24)
	})

	t.Run("concatnate strings", func(t *testing.T) {
		concatenate := func(x, y string) string {
			return x + y
		}
		AssertEqual(t, Reduce([]string{"a", "b", "c"}, concatenate, ""), "abc")
	})
}

func Reduce2[A, B any](collection []A, accumulator func(B, A) B, initalValue B) B {
	result := initalValue
	for _, v := range collection {
		result = accumulator(result, v)
	}

	return result
}

func TestBadBank(t *testing.T) {
	var (
		riya  = Account{Name: "Riya", Balance: 100}
		chris = Account{Name: "Chris", Balance: 75}
		adil  = Account{Name: "Adil", Balance: 200}

		transactions = []Transaction{
			NewTransaction(chris, riya, 100),
			NewTransaction(adil, chris, 25),
		}
	)

	newBalanceFor := func(account Account) float64 {
		return NewBalanceFor(account, transactions).Balance
	}

	AssertEqual(t, newBalanceFor(riya), 200)
	AssertEqual(t, newBalanceFor(chris), 0)
	AssertEqual(t, newBalanceFor(adil), 175)

}

type Transaction struct {
	From string
	To   string
	Sum  float64
}

func NewTransaction(from, to Account, sum float64) Transaction {
	return Transaction{From: from.Name, To: to.Name, Sum: sum}
}

type Account struct {
	Name    string
	Balance float64
}

func NewBalanceFor(account Account, transactions []Transaction) Account {
	return Reduce2(
		transactions,
		applyTransactions,
		account,
	)
}

func BalanceFor(transactions []Transaction, name string) float64 {
	adjustBalance := func(currentBalance float64, t Transaction) float64 {
		if t.From == name {
			return currentBalance - t.Sum
		}
		if t.To == name {
			return currentBalance + t.Sum
		}
		return currentBalance
	}

	return Reduce2(transactions, adjustBalance, 0.0)
}

func applyTransactions(a Account, transaction Transaction) Account {
	if a.Name == transaction.From {
		a.Balance -= transaction.Sum
	}

	if a.Name == transaction.To {
		a.Balance += transaction.Sum
	}

	return a
}
