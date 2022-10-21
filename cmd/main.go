package main

import (
	"github.com/akhilan/gin-lms/pkg/books"
	"github.com/akhilan/gin-lms/pkg/common/db"
	"github.com/akhilan/gin-lms/pkg/users"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)

	r := gin.Default()
	h := db.Init(dbUrl)

	books.RegisterRoutes(r, h)
	users.RegisterRoutes(r, h)
	// register more routes here

	r.Run(port)
}
