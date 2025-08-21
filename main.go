package main

import (
	"context"
	"go-test/config"
	"go-test/servs"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// CORS com configuração customizada
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200", "https://premiss.io"}, // origens permitidas
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour, // cache do preflight
	}))

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// GET /pokemon/:nameOrID  (ex.: /pokemon/ditto ou /pokemon/132)
	router.GET("/pokemon/:nameOrID", func(c *gin.Context) {
		nameOrID := c.Param("nameOrID")
		ctx, cancel := context.WithTimeout(c.Request.Context(), config.HttpTimeout)
		defer cancel()

		p, code, err := servs.FetchPokemon(ctx, nameOrID)
		if err != nil {
			c.JSON(code, gin.H{
				"error":   err.Error(),
				"message": "não foi possível obter dados da PokéAPI",
			})
			return
		}
		c.JSON(http.StatusOK, p)
	})
	router.GET("/pokemon", func(c *gin.Context) {
		name := c.Query("name")
		if name == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "param 'name' é obrigatório"})
			return
		}
		ctx, cancel := context.WithTimeout(c.Request.Context(), config.HttpTimeout)
		defer cancel()

		p, code, err := servs.FetchPokemon(ctx, name)
		if err != nil {
			c.JSON(code, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, p)
	})
	router.Run() // listen and serve on 0.0.0.0:8080
}
