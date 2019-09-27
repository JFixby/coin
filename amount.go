package coin

import "fmt"

const AtomsPerCoin = 1e8

type Amount struct {
	AtomsValue int64
}

func (a Amount) String() string {
	return fmt.Sprintf("%v coins", a.ToCoins())
}

func CoinsAmountFromFloat(coinsFloat float64) Amount {
	return Amount{int64(coinsFloat * AtomsPerCoin)}
}

func (a *Amount) ToCoins() float64 {
	return float64(a.AtomsValue) / AtomsPerCoin
}

func (a *Amount) ToAtoms() int64 {
	return a.AtomsValue
}

func (a *Amount) Copy() Amount {
	return Amount{a.AtomsValue}
}
