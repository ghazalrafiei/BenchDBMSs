package main

import (
	"fmt"
	"math/rand"
	"os"
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

func Bench(dbs dbmss.Dbms, address string, name string) error {

	fmt.Printf("Trying to connect to database %s\n", name)

	err := dbs.Connect(address)

	if err != nil {
		fmt.Printf("error: %s \n", err)
		return err
	}

	fmt.Println("Connected to database successfully")

	err = dbs.Create()

	if err != nil {
		fmt.Printf("error: %s \n", err)
		return err
	}

	fmt.Println("Database created successfully")

	var db_result result
	db_result.name = name
	db_result.replicas = 2

	db_result.setting, _ = BenchSetting(dbs)

	fmt.Println("Set test done")

	db_result.finding, _ = BenchFinding(dbs)

	fmt.Println("Find test done")

	db_result.deletion, _ = BenchDeleting(dbs)

	fmt.Println("Delete test done")

	fmt.Println(db_result)

	return nil
}

func main() {

	fmt.Println("Benchmarking Started...")

	db_name := os.Args[1]

	time.Sleep(5 * time.Second)

	switch db_name {
	case "postgres":
		var master_pstgo dbmss.Dbms
		master_pstgo = &dbmss.Pgs_connection_tool{}

		Bench(master_pstgo, "host=postgres-server-master port=5432 user=postgres dbname=gopost password=MasterPass sslmode=disable", "PostgreSQL")

	case "redis":
		var master_redis dbmss.Dbms
		master_redis = &dbmss.Rds_connection_tool{}

		Bench(master_redis, "redis-server-master:6379", "Redis")
	}

}

type result struct {
	name     string
	replicas int

	deletion time.Duration
	setting  time.Duration
	finding  time.Duration
}

func (r result) String() string {
	return fmt.Sprintf("\nResults for Comparing Few DBMSs for %d Random Tests:\n\n-%s(with %d replicas):\n\tSetting: %v\n\tFinding: %v\n\tDeleting: %v", bench_size, r.name, r.replicas, r.setting, r.finding, r.deletion)
}
