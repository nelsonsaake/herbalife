package main

func GetProduct(id int16) (product Product, err error) {
	product = Product{}
	err = Db.QueryRow("select Id, SKU, ProductName, VolumePoints, RetailPrice, EarnBase, ProductTypeId from Product where id=?", id).Scan(&product.Id, &product.SKU, &product.ProductName, &product.VolumePoints, &product.RetailPrice, &product.EarnBase, &product.ProductType.Id)
	if err != nil {
		panic(err)
	}
	product.ProductType, err = GetProductType(product.ProductType.Id)
	return
}

func GetAllProducts() (products []Product, err error) {
	products = []Product{}

	rows, err := Db.Query("select Id, SKU, ProductName, VolumePoints, RetailPrice, EarnBase, ProductTypeId from Product")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		product := Product{}
		err = rows.Scan(&product.Id, &product.SKU, &product.ProductName, &product.VolumePoints, &product.RetailPrice, &product.EarnBase, &product.ProductType.Id)
		if err != nil {
			panic(err)
		}

		product.ProductType, err = GetProductType(product.ProductType.Id)
		if err != nil {
			panic(err)
		}

		products = append(products, product)
	}

	return
}

func GetAllProductWith(search string) (products []Product, err error) {
	products = []Product{}

	rows, err := Db.Query("select Id, SKU, ProductName, VolumePoints, RetailPrice, EarnBase, ProductTypeId from Product where ProductName Like '%" + search + "%'")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		product := Product{}
		err = rows.Scan(&product.Id, &product.SKU, &product.ProductName, &product.VolumePoints, &product.RetailPrice, &product.EarnBase, &product.ProductType.Id)
		if err != nil {
			panic(err)
		}

		product.ProductType, err = GetProductType(product.ProductType.Id)
		if err != nil {
			panic(err)
		}

		products = append(products, product)
	}

	return
}
