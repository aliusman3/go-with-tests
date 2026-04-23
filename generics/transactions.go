package generics

type Transaction struct {
	From string
	To   string
	Sum  float64
}

type Account struct {
	Name    string
	Balance float64
}

func NewTransaction(from, to Account, sum float64) Transaction {
	return Transaction{
		From: from.Name,
		To:   to.Name,
		Sum:  sum,
	}
}

func NewBalanceFor(account Account, transactions []Transaction) Account {
	return Reduce(transactions, applyTransactions, account)
}

func applyTransactions(account Account, t Transaction) Account {
	if t.From == account.Name {
		account.Balance -= t.Sum
	}
	if t.To == account.Name {
		account.Balance += t.Sum
	}
	return account
}

func BalanceFor(transactions []Transaction, person string) float64 {
	adjustBalance := func(currentBalance float64, t Transaction) float64 {
		if t.From == person {
			return currentBalance - t.Sum
		}
		if t.To == person {
			return currentBalance + t.Sum
		}
		return currentBalance
	}
	return Reduce(transactions, adjustBalance, 0)
}
