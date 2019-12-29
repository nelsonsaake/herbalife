package main

import "math"

// types
type Product struct {
	Id           int16
	SKU          string
	ProductName  string
	VolumePoints float64
	EarnBase     float64
	RetailPrice  float64
	ProductType  ProductType
}

type MemberType int

const (
	Distributor       MemberType = 1
	SeniorConsultant             = 2
	SuccessBuilder               = 3
	QualifiedProducer            = 4
	Supervisor                   = 5
)

//
var TwoDP int

func init() {
	TwoDP = 2
}

//
func precesion(n int, val float64) (res float64) {
	exp := math.Pow10(n)
	res = math.Round(val*exp) / exp
	return
}

func Percent(percentage float64, num float64) (res float64) {
	percentage /= 100
	res = percentage * num
	res = precesion(TwoDP, res)
	return
}

// special rates
func (product *Product) Percent(percentage float64) (res float64) {
	return Percent(percentage, product.RetailPrice)
}

func (product *Product) Percent1() (res float64) {
	percent := 1.00
	return product.Percent(percent)
}

func (product *Product) Percent094() (res float64) {
	percent := 0.94
	return product.Percent(percent)
}

func (product *Product) Percent3() (res float64) {
	percent := 3.00
	num := product.RetailPrice + product.Percent1() + product.Percent094()
	return Percent(percent, num)
}

func (product *Product) SuggestedSellingPrice() (res float64) {
	res = product.RetailPrice + product.Percent1() + product.Percent3() + product.Percent094()
	res = precesion(TwoDP, res)
	return
}

// special price
func PriceRate(member MemberType) float64 {
	switch member {
	case Distributor:
		return 25
	case SeniorConsultant:
		return 35
	case SuccessBuilder:
	case QualifiedProducer:
		return 45
	case Supervisor:
		return 50
	}
	return 0
}

func (product *Product) MemberPrice(member MemberType) (res float64) {
	res = product.RetailPrice + product.Percent1() + product.Percent3() + product.Percent094()
	res = precesion(TwoDP, res)
	rate := PriceRate(member)
	res = (product.SuggestedSellingPrice() - Percent(rate, product.EarnBase)) * 1.03
	res = precesion(TwoDP, res)
	return
}

func (product *Product) DistributorPrice() (res float64) {
	return product.MemberPrice(Distributor)
}

func (product *Product) SeniorConsultantPrice() (res float64) {
	return product.MemberPrice(SeniorConsultant)
}

func (product *Product) SuccessBuilderPrice() (res float64) {
	return product.MemberPrice(SuccessBuilder)
}

func (product *Product) SupervisorPrice() (res float64) {
	return product.MemberPrice(Supervisor)
}

func (product *Product) QualifiedProducerPrice() (res float64) {
	return product.MemberPrice(QualifiedProducer)
}


