package modules

import "github.com/shopspring/decimal"

type Product struct {
	Model
	Code  int64
	Price decimal.Decimal
}
