package main

import (
	"log"
	"net/http"
	"slices"
	"strconv"
	"sync"

	"github.com/labstack/echo/v5"
)

type store struct {
	mu    sync.RWMutex
	items []string
}

func main() {
	e := echo.New()
	s := &store{items: []string{}}

	e.GET("/", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, map[string]any{"message": "Hello World"})
	})

	e.GET("/list", s.list)
	e.POST("/list/add/:item", s.add)
	e.DELETE("/list/delete/:id", s.remove)

	log.Println("server running on :3000")

	if err := e.Start(":3000"); err != nil {
		log.Fatalf("server error: %v", err)
	}
}

func (s *store) list(c *echo.Context) error {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return c.JSON(http.StatusOK, map[string]any{
		"status": "ok",
		"list":   s.items,
	})
}

func (s *store) add(c *echo.Context) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.items = append(s.items, c.Param("item"))

	return c.JSON(http.StatusOK, map[string]any{
		"status": "ok",
		"list":   s.items,
	})
}

func (s *store) remove(c *echo.Context) error {
	index, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  "error",
			"message": "ID inválido",
		})
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if index < 0 || index >= len(s.items) {
		return c.JSON(http.StatusNotFound, map[string]any{
			"status":  "error",
			"message": "Índice não encontrado",
		})
	}

	s.items = slices.Delete(s.items, index, index+1)

	return c.JSON(http.StatusOK, map[string]any{
		"status": "ok",
		"list":   s.items,
	})
}
