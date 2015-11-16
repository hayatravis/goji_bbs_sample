package main

import (
	"database/sql"
	"fmt"
	"github.com/codegangsta/cli"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"time"
	// "github.com/elgs/gosqljson"
	"math/rand"
)

func main() {
	app := cli.NewApp()
	app.Name = "updater"
	app.Usage = "updater is update display flag gradually."
	app.Action = func(c *cli.Context) {
		/* -- DB -- */
		db, err := sql.Open("mysql", "root:@/test_bbs")
		if err != nil {
			panic(err.Error())
		}
		defer db.Close()

		/* -- Params -- */
		offset := c.Args()[0]
		limit := c.Args()[1]

		/* -- SQL -- */
		sql := "SELECT id, name FROM user LIMIT ?, ?;"
		update_sql := "UPDATE user SET name = 'test', modified = NOW() WHERE id = ?;"
		data, _ := db.Query(sql, offset, limit)
		defer data.Close()

		var randNum int

		for data.Next() {
			var id int
			var name string
			if err := data.Scan(&id, &name); err != nil {
				fmt.Println("error")
			}
			randNum = randInt(15, 5)
			time.Sleep(time.Duration(randNum) * time.Second)
			db.Exec(update_sql, id)
			fmt.Printf("%d is %s \n", id, name)
		}

		println("Udpater Finished!")
	}

	app.Run(os.Args)
}

func randInt(max int, min int) int {
	return min + rand.Intn(max-min)
}
