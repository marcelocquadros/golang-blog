package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/marcelocquadros/blog/internal/app/infra/repositories/user"
	usecase "github.com/marcelocquadros/blog/internal/app/usecases/user"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	r := gin.Default()
	db, err := sqlx.Connect("sqlite3", ":memory:")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	userRepo := user.NewUserRepository(db)

	uc := usecase.NewCreateUser(userRepo)

	r.POST("/", func(c *gin.Context) {
		var cmd usecase.CreateUserCmd
		if err := c.ShouldBindJSON(&cmd); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		id, err := uc.Execute(&cmd)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		c.JSON(201, id)
	})

	r.Run()
}
