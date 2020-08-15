package dbmss

import (
	"fmt"

	"encoding/json"

	"github.com/ghazalrafiei/BenchDBMSs/object"
	"github.com/go-redis/redis"
)

type Rds_connection struct {
	client  *redis.Client
	last_id uint
}

func (rd *Rds_connection) Connect(adr string) error {
	fmt.Println("called")
	client := redis.NewClient(&redis.Options{
		Addr:     adr,
		Password: "",
		DB:       0,
	})
	_, err := client.Ping().Result()
	rd.client = client
	rd.last_id = 0
	return err
}

func (rd *Rds_connection) Create() error {

	return nil
}

func (rd *Rds_connection) Set(obj *object.Object) error {
	js, err := json.Marshal(obj)
	rd.last_id++
	rd.client.Set(string(rd.last_id), js, 0).Err()
	if err != nil {
		fmt.Println(err)
	}

	return nil
}

func (rd *Rds_connection) Delete(id uint) error {
	rd.client.Del(string(id))
	return nil
}

func (rd *Rds_connection) Find(id uint) error {
	_, err := rd.client.Get(string(id)).Result()
	return err
}
