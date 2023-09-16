package main

import (
	"backend-blogtechv2/db"
	"backend-blogtechv2/log"
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
)
func init() {
	fmt.Println(" >>>>>> ", os.Getenv("APP_NAME")) 
	//os.Setenv("APP_NAME", "github")
	log.InitLogger(false)
}

func main() {

	sql := &db.Sql{
		Host:     "localhost", //"localhost"
		Port:     5432,
		UserName: "postgres",
		PassWord: "Hoangkimluong192@",
		DbName:   "blogtech-data",
	} 
	sql.Connect()
	defer sql.Close()
	e := echo.New()

	e.Logger.Fatal(e.Start(":3000"))
}