package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/ghazalrafiei/BenchDBMSs/dbmss"
	"github.com/ghazalrafiei/BenchDBMSs/object"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var bench_size int = 1000

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

func Bench(dbs []dbmss.Dbms, address []string, name string) error {
	for i, b := range dbs {

		err := b.Connect(address[i])
		if err != nil {
			return err
		}
		err = b.Create()
		if err != nil {
			return err
		}
	}

	var db_result result
	db_result.name = name
	db_result.replicas = len(dbs)

	var (
		setting_duration  time.Duration
		deletion_duration time.Duration
		finding_duration  time.Duration
		getting_duration  time.Duration
	)
	for _, b := range dbs {
		sd, _ := BenchSetting(b)
		setting_duration += sd
	}
	for _, b := range dbs {
		sd, _ := BenchDeleting(b)
		setting_duration += sd
	}
	for _, b := range dbs {
		sd, _ := BenchFinding(b)
		setting_duration += sd
	}
	for _, b := range dbs {
		sd, _ := BenchGetting(b)
		setting_duration += sd
	}

	db_result.setting = setting_duration
	db_result.deletion = deletion_duration
	db_result.finding = finding_duration
	db_result.getting = getting_duration

	fmt.Println(db_result)
	return nil
}

func main() {
	var (
		//No difference between master and slave in setting and deletion but in finding and getting
		master_pstgo dbmss.Dbms
		slave_pstgo1 dbmss.Dbms
		slave_pstgo2 dbmss.Dbms
	)

	addresses := []string{
		"host=localhost port=5432 user=postgres dbname=pst-go password=paSs",
		"host=localhost port=5432 user=postgres dbname=pst-go1 password=paSs",
		"host=localhost port=5432 user=postgres dbname=pst-go2 password=paSs",
	}

	master_pstgo = &dbmss.Pgs_connection{}
	slave_pstgo1 = &dbmss.Pgs_connection{}
	slave_pstgo2 = &dbmss.Pgs_connection{}

	Bench([]dbmss.Dbms{master_pstgo, slave_pstgo1, slave_pstgo2}, addresses, "PostgreSQL")

}

type result struct {
	name     string
	replicas int

	deletion time.Duration
	setting  time.Duration
	finding  time.Duration
	getting  time.Duration
}

func (r result) String() string {
	return fmt.Sprintf("Results for Comparing Few DBMSs for %d Random Tests:\n\n-%s(with %d replicas):\n\tSetting: %v\n\tDeleting: %v\n\tFinding: %v\n\tGetting: %v", bench_size, r.name, r.replicas, r.setting, r.deletion, r.finding, r.getting)
}
