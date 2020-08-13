package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/ghazalrafiei/BenchDBMSs/dbmss"
	"github.com/ghazalrafiei/BenchDBMSs/object"
)

var bench_size int = 1000

func BenchSetting(db dbmss.Dbms) (time.Duration, error) {
	st := time.Now()
	for i := 0; i < bench_size; i++ {
		index := rand.Intn(5)
		obj := object.Object{
			Type:      object.Types[index],
			Name:      object.Names[index],
			Namespace: object.Namespaces[index]}
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
	for i := 1; i <= bench_size; i++ {
		index := uint(rand.Intn(bench_size))
		db.Delete(index)
	}

	ed := time.Now()
	return ed.Sub(st), nil
}
func BenchFinding(db dbmss.Dbms) (time.Duration, error) {
	st := time.Now()
	for i := 1; i <= bench_size; i++ {
		index := uint(rand.Intn(bench_size))
		db.Find(index)
	}

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
	db_result.replicas = 2

	db_result.setting, _ = BenchSetting(dbs)

	db_result.finding, _ = BenchFinding(dbs)

	db_result.deletion, _ = BenchDeleting(dbs)

	fmt.Println(db_result)
	return nil
}

func main() {

	var master_pstgo dbmss.Dbms

	master_pstgo = &dbmss.Pgs_connection{}

	Bench(master_pstgo, "host=localhost port=5432 user=postgres dbname=pst-go password=paSs", "PostgreSQL")

}

type result struct {
	name     string
	replicas int

	deletion time.Duration
	setting  time.Duration
	finding  time.Duration
}

func (r result) String() string {
	return fmt.Sprintf("Results for Comparing Few DBMSs for %d Random Tests:\n\n-%s(with %d replicas):\n\tSetting: %v\n\tFinding: %v\n\tDeleting: %v", bench_size, r.name, r.replicas, r.setting, r.finding, r.deletion)
}
