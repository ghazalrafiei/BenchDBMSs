package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/ghazalrafiei/BenchDBMSs/dbmss"
	"github.com/ghazalrafiei/BenchDBMSs/object"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var bench_size int = 1500

func BenchSetting(db dbmss.Dbms) (time.Duration, error) {
	st := time.Now()
	for i := 0; i < bench_size; i++ {
		index := rand.Intn(5)
		obj := object.Object{
			Key: object.Key{
				Type:      object.Types[index],
				Name:      object.Names[index],
				Namespace: object.Namespaces[index]},
			Value: object.Values[index],
			Meta: object.Meta{CreationTime: time.Now(),
				Owner:  object.Owners[index],
				Method: object.Methods[index]}}

		err := db.Set(&obj)
		if err != nil {
			return 0, err
		}
	}
	ed := time.Now()
	return ed.Sub(st), nil
}

func BenchDeleting(db dbmss.Dbms) (time.Duration, error) {
	st := time.Now()
	//TODO: BODY

	ed := time.Now()
	return ed.Sub(st), nil
}
func BenchFinding(db dbmss.Dbms) (time.Duration, error) {
	st := time.Now()
	//TODO: BODY

	ed := time.Now()
	return ed.Sub(st), nil
}
func BenchGetting(db dbmss.Dbms) (time.Duration, error) {
	st := time.Now()
	//TODO: BODY

	ed := time.Now()
	return ed.Sub(st), nil
}

func Bench(dbs dbmss.Dbms, address string, name string) error {
	err := dbs.Connect(address)
	if err != nil {
		return err
	}

	err = dbs.Create()
	if err != nil {
		return err
	}

	var db_result result
	db_result.name = name

	setting_duration, _ := BenchSetting(dbs)
	deletion_duration, _ := BenchDeleting(dbs)
	finding_duration, _ := BenchFinding(dbs)
	getting_duration, _ := BenchGetting(dbs)

	db_result.setting = setting_duration
	db_result.deletion = deletion_duration
	db_result.finding = finding_duration
	db_result.getting = getting_duration

	fmt.Println(db_result)
	return nil
}

func main() {
	var pstgo dbmss.Dbms
	pstgo = &dbmss.Pgs_connection{}
	Bench(pstgo, "host=localhost port=5432 user=postgres dbname=pst-go password=paSs", "PostgreSQL")
}

type result struct {
	name string

	deletion time.Duration
	setting  time.Duration
	finding  time.Duration
	getting  time.Duration
}

func (r result) String() string {
	return fmt.Sprintf("Results for Comparing Few DBMSs for %d Random Tests:\n\n-%s:\n\tSetting: %v\n\tDeleting: %v\n\tFinding: %v\n\tGetting: %v", bench_size, r.name, r.setting, r.deletion, r.finding, r.getting)
}
