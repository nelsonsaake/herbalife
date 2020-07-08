package src

type Package struct {
	Description string
	Products    []Product
}

func GetPackage(ids []int) (pkg Package) {
	products := []Product{}
	for _, id := range ids {
		product, err := GetProduct(id)
		if err != nil {
			panic(err)
		}
		products = append(products, product)
	}

	pkg.Products = products
	return
}

func (pkg *Package) TotalSellingPrice() (total float64) {
	total = 0
	for _, product := range pkg.Products {
		total += product.SuggestedSellingPrice()
	}
	total = precesion(TwoDP, total)
	return
}

func (pkg *Package) Len() int {
	return len(pkg.Products)
}
