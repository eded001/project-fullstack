package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v5"
)

func main() {
	e := echo.New()

	var list []string

	e.GET("/", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, map[string]any{
			"message": "Hello World",
		})
	})

	e.GET("/:id", func(c *echo.Context) error {
		id := c.Param("id")

		return c.JSON(http.StatusOK, map[string]any{
			"status": "ok",
			"id":     id,
		})
	})

	e.GET("/list", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, map[string]any{
			"status": "ok",
			"list":   list,
		})
	})

	e.POST("/list/add/:item", func(c *echo.Context) error {
		item := c.Param("item")

		list = append(list, item)

		return c.JSON(http.StatusOK, map[string]any{
			"status": "ok",
			"list":   list,
		})
	})

	e.DELETE("/list/delete/:id", func(c *echo.Context) error {
		id := c.Param("id")

		index, err := strconv.Atoi(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"status":  "error",
				"message": "ID inválido",
			})
		}

		if index < 0 || index >= len(list) {
			return c.JSON(http.StatusNotFound, map[string]any{
				"status":  "error",
				"message": "Índice não encontrado",
			})
		}

		list = append(list[:index], list[index+1:]...)

		return c.JSON(http.StatusOK, map[string]any{
			"status": "ok",
			"list":   list,
		})
	})

	log.Println("server running on :3000")

	if err := e.Start(":3000"); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
