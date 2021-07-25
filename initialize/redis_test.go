package initialize

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/go-redis/redis"
)

func TestRedis(t *testing.T) {

	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "123456", // no password set
		DB:       0,        // use default DB
	})
	_, err := client.Ping().Result()
	if err != nil {
		t.Error(err)
	}

	/* client.Set("key", "val", time.Second*10)
	mapInstance := make(map[string]interface{})
	mapInstance["Name"] = "jqw"
	mapInstance["Age"] = 18 */

	type Student struct {
		Num  int
		Name string
		Age  int
	}
	student := Student{10, "jqw", 18}
	data := StructToMapDemo(student)
	t.Log(data)
	fmt.Println(data)
	val, err := client.HMSet("hmset", data).Result()
	if err != nil {
		t.Error(err)
	}
	t.Log(val)

}

func StructToMapDemo(obj interface{}) map[string]interface{} {
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		data[obj1.Field(i).Name] = obj2.Field(i).Interface()
	}
	return data
}
