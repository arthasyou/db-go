package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/arthasyou/db-go/mysql"
)

// Examples

func main() {
	mysql.Connect("localhost", 3306, "root", "123456", "test")
	// Select
	// mysql.Select("user", []string{"id", "name"}) // SELECT `id`,`name` FROM `user`
	// mysql.Select("tableName")                    // SELECT * FROM `user`

	// Where
	kop := mysql.KeyOpVal{Key: "id", Op: "=", Value: "1"}
	ss := mysql.Select("user", []string{"id", "name"})
	ww := mysql.WhereOr([]*mysql.KeyOpVal{&kop, &kop, &kop})
	fmt.Println(ss + ww)

	// Insert
	kv := mysql.KeyVal{Key: "id", Value: "1"}
	kv1 := mysql.KeyVal{Key: "name", Value: "Jams"}
	a := mysql.Insert("user", []*mysql.KeyVal{&kv, &kv1}) // INSERT INTO `user` (`id`,`name`) VALUES ('1','Jams')
	fmt.Println(a)

	// Update
	v := mysql.Update("user", []*mysql.KeyVal{&kv, &kv1}) // UPDATE `user` SET `id` = '1',`name` = 'Jams'
	fmt.Println(v)

	xx := mysql.Limit(3, 5) // LIMIT 10, 5
	fmt.Println(xx)

	// func Query() {
	// 	u := test.User{}
	// 	a := []test.User{}
	// 	r, err := cli.Query("SELECT * FROM user")
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	for r.Next() {
	// 		err = r.Scan(&u.Id, &u.Appkey, &u.Appsecret, &u.Pools, &u.Created, &u.Remark)
	// 		a = append(a, u)
	// 	}

	// 	fmt.Println(a)
	// }
	waitExit()
}

func waitExit() {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	<-ch
}
