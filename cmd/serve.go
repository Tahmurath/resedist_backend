package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"net/http"
	"resedist/pkg/config"
	"resedist/pkg/routing"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve app on dev server",
	Long:  `Application will be served on host and port defined in config.yml`,
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

func serve() {
	config.Set()

	routing.Init()
	router := routing.GetRouter()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"appName": viper.Get("App.Name"),
		})
	})

	routing.Serve()
}
