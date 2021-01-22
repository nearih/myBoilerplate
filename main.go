package main

import (
	"boilerplate/config"
	"boilerplate/db/sqlitedb"
	"boilerplate/model/example"
	"boilerplate/src/handler"
	"fmt"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

func main() {
	startDb()
	startHttpServer()
}

func startHttpServer() {
	e := echo.New()
	e.Logger.SetOutput(os.Stdout)
	e.Logger.SetLevel(log.INFO)
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: os.Stdout,
		Format: `{"id":"${id}","remote_ip":"${remote_ip}",` +
			`"host":"${host}","method":"${method}","uri":"${uri}","user_agent":"${user_agent}",` +
			`"status":${status},"error":"${error}","latency":${latency},"latency_human":"${latency_human}"` +
			`,"bytes_in":${bytes_in},"bytes_out":${bytes_out}}` + "\n",
	}))

	// register handler
	handler.RegisterHandler(e)

	//start server
	port := fmt.Sprintf(":%+v", config.Config.Server.Port)
	e.Logger.Fatal(e.Start(port))
}

func startGrpcServer() {

}

func startDb() {
	dbCon := sqlitedb.NewConnection()
	dbCon.Db.AutoMigrate(example.Example{})

	x := []example.Example{}
	x = append(x, example.Example{Username: "userA", Password: "pass"})
	x = append(x, example.Example{Username: "userB", Password: "passB"})

	y := example.Example{Username: "userY", Password: "passY"}

	dbCon.Db.Create(y)
	dbCon.Db.CreateInBatches(x, len(x))
}
