package dbmss

import (
	"errors"

	"github.com/ghazalrafiei/BenchDBMSs/object"
	"github.com/jinzhu/gorm"
)

type Pgs_connection_tool struct {
	dbs *gorm.DB
}

func (c *Pgs_connection_tool) Connect(adr string) error {

	db, err := gorm.Open("postgres", adr)

	if err != nil {
		return err
	}

	c.dbs = db

	return nil
}

func (c *Pgs_connection_tool) Create() error {

	if c.dbs.HasTable(&object.Object{}) {
		c.dbs.DropTable(&object.Object{})
	}

	c.dbs.CreateTable(&object.Object{})

	return nil
}

func (c *Pgs_connection_tool) Set(obj *object.Object) error {

	c.dbs.Create(&obj)
	res := c.dbs.NewRecord(obj)

	if res {
		return errors.New("Blank Primary")
	}

	return nil
}

func (c *Pgs_connection_tool) Find(id uint) error {

	var obj object.Object
	c.dbs.Find(&obj, id)

	return nil
}

func (c *Pgs_connection_tool) Delete(id uint) error {

	var obj object.Object
	c.dbs.Delete(&obj, id)

	return nil
}
