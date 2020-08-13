package dbmss

import (
	"github.com/ghazalrafiei/BenchDBMSs/object"
)

type Dbms interface {
	Connect(string) error
	Create() error

	Set(*object.Object) error
	Find(uint) error
	Delete(uint) error
}
