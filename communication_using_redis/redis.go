package main 
import ("fmt"
		"github.com/go-redis/redis"
)

func main(){
	fmt.Println("Go Redis Tutorial")

	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "", 
		DB: 0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong,err)

	var i int = 0
	// max := 100000

	for true {
		// fmt.Println(i)
		i += 1
		// err = client.Set("name", i , 0).Err()

		// if err != nil {
		// 	fmt.Println(err)
		// }

		val, err := client.Get("id").Result()
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(val)
	}
	

	
}



