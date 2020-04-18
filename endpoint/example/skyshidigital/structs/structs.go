package structs

import "github.com/jinzhu/gorm"

/*Person : field person*/
type Person struct {
	gorm.Model
	FirstName string
	LastName  string
}
