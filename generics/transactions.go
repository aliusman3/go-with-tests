package generics

type Transaction struct {
	From string
	To   string
	Sum  float64
}

func BalanceFor(transactions []Transaction, person string) float64 {
	var balance float64
	for _, t := range transactions {
		if t.From == person {
			balance -= t.Sum
		}
		if t.To == person {
			balance += t.Sum
		}
	}
	return balance
}
