package main

import (
	"database/sql"
	"github.com/codegangsta/cli"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "initUpdater"
	app.Usage = "initUpdater is to init user table."
	app.Action = func(c *cli.Context) {
		/* -- DB -- */
		db, err := sql.Open("mysql", "root:@/test_bbs")
		if err != nil {
			panic(err.Error())
		}
		defer db.Close()

		/* -- SQL -- */
		sql_1 := "UPDATE user SET name = 'Paul', modified = NOW() WHERE id = 1;"
		data1, _ := db.Exec(sql_1)

		sql_2 := "UPDATE user SET name = 'John', modified = NOW() WHERE id = 2;"
		data2, _ := db.Exec(sql_2)

		println(data1)
		println(data2)
		println("initUdpater Finished!")
	}

	app.Run(os.Args)
}
