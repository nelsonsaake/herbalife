package src

type ProductType struct {
	Id          int16
	Description string
}

func GetProductType(id int16) (productType ProductType, err error) {
	productType = ProductType{Id: id}
	err = Db.QueryRow("select Description from ProductType where id=$1", id).Scan(&productType.Description)
	return
}
