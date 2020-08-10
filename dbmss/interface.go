package dbmss

import (
	"github.com/ghazalrafiei/BenchDBMSs/object"
)

type Dbms interface {
	Connect(string) error
	Create() error

	Set(*object.Object) error
	Delete(object.Key) error
	Get(*object.Key) (*object.Object, error)
	Find(*object.Key) ([]*object.Object, error)
}
