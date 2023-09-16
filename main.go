package main

import (
	"backend-blogtechv2/db"
	"backend-blogtechv2/handler"
	"backend-blogtechv2/log"
	repoimpl "backend-blogtechv2/repositoty/repo_impl"
	userrouter "backend-blogtechv2/router/user_router"
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

	userHandler := handler.UserHandler{
		UserRepo: repoimpl.NewUserRepo(sql),
	}

	api := userrouter.API{
		Echo:        e,
		UserHandler: userHandler,
	}

	api.SetupRouter()

	e.Logger.Fatal(e.Start(":3000"))
}