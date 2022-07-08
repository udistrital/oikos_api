package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/cors"
	_ "github.com/lib/pq"

	_ "github.com/udistrital/oikos_api/routers"
	apistatus "github.com/udistrital/utils_oas/apiStatusLib"
	"github.com/udistrital/utils_oas/auditoria"
	"github.com/udistrital/utils_oas/customerror"
)

func dev() {
	if beego.BConfig.RunMode == "dev" {
		orm.Debug = true
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
}

func config() {
	dev()
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
}

func main() {
	orm.RegisterDataBase("default", "postgres",
		"postgres://"+beego.AppConfig.String("PGuser")+
			":"+beego.AppConfig.String("PGpass")+
			"@"+beego.AppConfig.String("PGhost")+
			":"+beego.AppConfig.String("PGport")+
			"/"+beego.AppConfig.String("PGdb")+
			"?sslmode=disable&search_path="+beego.AppConfig.String("PGschemas")+"")
	config()
	auditoria.InitMiddleware()
	beego.ErrorController(&customerror.CustomErrorController{})
	apistatus.Init()
	beego.Run()
}
