package object

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Object struct {
	gorm.Model
	Type      string
	Name      string
	Namespace string
}

var ( //0-4
	Types      = []string{"type1", "type2", "type3", "type4", "type5"}
	Names      = []string{"name1", "name2", "name3", "name4", "name5"}
	Namespaces = []string{"namespace1", "namespace2", "namespace3", "namespace4", "namespace5"}
	Owners     = []string{"owner1", "owner2", "owner3", "owner4", "owner5"}
	Methods    = []string{"get", "send", "delete", "find", "watch"}
	Values     = []string{"value1", "value2", "value3", "value4", "value5"}
)
