package dbmss

import (
	"encoding/json"
	"fmt"

	"github.com/ghazalrafiei/BenchDBMSs/object"
	"github.com/go-redis/redis"
)

type Rds_connection_tool struct {
	client  *redis.Client
	last_id uint
}

func (rd *Rds_connection_tool) Connect(adr string) error {

	client := redis.NewClient(&redis.Options{
		Addr:     adr,
		Password: "redis",
		DB:       0,
	})
	_, err := client.Ping().Result()
	rd.client = client
	rd.last_id = 0

	return err
}

func (rd *Rds_connection_tool) Create() error {

	rd.client.FlushAll()

	return nil
}

func (rd *Rds_connection_tool) Set(obj *object.Object) error {

	js, err := json.Marshal(obj)
	rd.last_id++
	rd.client.Set(string(rd.last_id), js, 0).Err()

	if err != nil {
		fmt.Println(err)
	}

	return nil
}

func (rd *Rds_connection_tool) Delete(id uint) error {

	rd.client.Del(string(id))
	return nil

}

func (rd *Rds_connection_tool) Find(id uint) error {

	_, err := rd.client.Get(string(id)).Result()

	return err
}
