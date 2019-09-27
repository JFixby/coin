package coinamount

import "fmt"

const AtomsPerCoin = 1e8

type CoinsAmount struct {
	AtomsValue int64
}

func (a CoinsAmount) String() string {
	return fmt.Sprintf("%v coins", a.ToCoins())
}

func CoinsAmountFromFloat(coinsFloat float64) CoinsAmount {
	return CoinsAmount{int64(coinsFloat * AtomsPerCoin)}
}

func (a *CoinsAmount) ToCoins() float64 {
	return float64(a.AtomsValue) / AtomsPerCoin
}

func (a *CoinsAmount) ToAtoms() int64 {
	return a.AtomsValue
}

func (a *CoinsAmount) Copy() CoinsAmount {
	return CoinsAmount{a.AtomsValue}
}
