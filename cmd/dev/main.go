package main

import (
	"backend-blogtechv2/db"
	"backend-blogtechv2/handler"
	"backend-blogtechv2/helper"
	"backend-blogtechv2/log" 
	repoimpl "backend-blogtechv2/repositoty/repo_impl"
	"backend-blogtechv2/router"
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
		Host:     "35.184.240.100", //"localhost"
		Port:     5432,
		UserName: "postgres",
		PassWord: "Hoangkimluong192@",
		DbName:   "blogtech-data",
	} 
	sql.Connect()
	defer sql.Close()
	e := echo.New()

	structValidator := helper.NewStructValidator()
	structValidator.RegisterValidate() 
	e.Validator = structValidator

	userHandler := handler.UserHandler{
		UserRepo: repoimpl.NewUserRepo(sql),
	}

	postHandler := handler.PostHandler{
		PostRepo: repoimpl.NewPostRepo(sql),
	}

	 
	api := router.API{
		Echo:        e,
		UserHandler: userHandler,
		PostHandler: postHandler,
	}

	api.SetupRouter()
 
	e.Logger.Fatal(e.Start(":3000"))
}