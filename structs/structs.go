package structs

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Product_Name string
	Product_Category string
}