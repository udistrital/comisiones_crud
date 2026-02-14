package main

import (
	_ "github.com/udistrital/comisiones_estudio_crud/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/udistrital/utils_oas/customerrorv2"

)

func main() {
  orm.Debug = true
  orm.RegisterDataBase("default", "postgres", beego.AppConfig.String("sqlconn"))
  //orm.RegisterDataBase("default", "postgres", "postgres://postgres:1234@192.168.160.1:5432/comisiones?sslmode=disable&search_path=comisiones")
  if beego.BConfig.RunMode == "dev" {
  	beego.BConfig.WebConfig.DirectoryIndex = true
  	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
  }
  beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
    AllowOrigins: []string{"*"},
    AllowMethods: []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
    AllowHeaders: []string{"Origin", "x-requested-with",
      "content-type",
      "accept",
      "origin",
      "authorization",
      "x-csrftoken"},
    ExposeHeaders:    []string{"Content-Length"},
    AllowCredentials: true,
  }))
  beego.ErrorController(&customerrorv2.CustomErrorController{})
  beego.Run()
}
