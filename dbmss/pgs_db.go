package dbmss

import (
	"errors"
	"fmt"

	"github.com/ghazalrafiei/BenchDBMSs/object"
	"github.com/jinzhu/gorm"
)

type Pgs_connection struct {
	dbs *gorm.DB
}

func (c *Pgs_connection) Connect(adr string) error {
	db, err := gorm.Open("postgres", adr)

	if err != nil {
		return err
	}
	c.dbs = db
	return nil
}

func (c *Pgs_connection) Create() error {
	if c.dbs.HasTable(&object.Object{}) {
		c.dbs.DropTable(&object.Object{})
		c.dbs.DropTable(&object.Key{})
		c.dbs.DropTable(&object.Meta{})
	}

	c.dbs.CreateTable(&object.Object{})
	c.dbs.CreateTable(&object.Key{})
	c.dbs.CreateTable(&object.Meta{})

	return nil
}
func (c *Pgs_connection) Set(obj *object.Object) error {
	c.dbs.Create(&obj)
	res := c.dbs.NewRecord(obj)
	if res {
		return errors.New("Blank Primary")
	}
	return nil
}

func (c *Pgs_connection) Delete(k object.Key) error {
	//TO DO
	c.dbs.Create(&k)
	c.dbs.Where("key = ?", k).Delete(&object.Object{}) //error: unsupported type main.key, a struct
	return nil
}

func (c *Pgs_connection) Get(*object.Key) (*object.Object, error) {
	//TO DO
	return nil, nil
}
func (c *Pgs_connection) Find(k *object.Key) ([]*object.Object, error) {
	//TO DO
	n := c.dbs.Where("key = ?", k).Find(&object.Object{})
	fmt.Println(n)
	return nil, nil
}
