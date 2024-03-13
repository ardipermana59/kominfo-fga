package main

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

var users []User

func main() {
	GinHttp()
}

func GinHttp() {
	// gin => framework HTTP punya golang
	// big community
	engine := gin.New()

	// serve static template
	// engine.LoadHTMLGlob("static/*")
	engine.Static("/static", "./static")

	engine.LoadHTMLGlob("template/*")
	engine.GET("/template/index/:name", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.tmpl", map[string]any{
			"title": ctx.Param("name"),
		})
	})

	// membuat prefix group
	v1 := engine.Group("/api/v1")
	{
		usersGroup := v1.Group("/users")
		{
			// [GET] /api/v1/users
			// filter user by email
			usersGroup.GET("", func(ctx *gin.Context) {
				email := ctx.Query("email")
				if email != "" {
					arr := []User{}
					for _, user := range users {
						// full text search
						if strings.Contains(user.Email, email) {
							arr = append(arr, user)
						}
					}
					ctx.JSON(http.StatusOK, arr)
					return
				}
				ctx.JSON(http.StatusOK, users)
			})

			// [POST] /api/v1/users
			usersGroup.POST("", func(ctx *gin.Context) {
				// binding payload
				user := User{}
				if err := ctx.Bind(&user); err != nil {
					ctx.JSON(http.StatusBadRequest, map[string]any{
						"message": "failed to bind body",
					})
					return
				}
				user.ID = uint(len(users) + 1)
				users = append(users, user)
				ctx.JSON(http.StatusAccepted, map[string]any{
					"message": "user created",
				})
			})

			// [GET] /api/v1/users/:id
			usersGroup.GET("/:id", func(ctx *gin.Context) {
				id, err := strconv.Atoi(ctx.Param("id"))
				if err != nil || id <= 0 {
					ctx.JSON(http.StatusBadRequest, map[string]any{
						"message": "invalid ID",
					})
					return
				}
				for _, user := range users {
					if user.ID == uint(id) {
						ctx.JSON(http.StatusOK, user)
						return
					}
				}

				ctx.JSON(http.StatusNotFound, map[string]any{
					"message": "user not found",
				})
			})

			// [PUT] /api/v1/users/:id
			usersGroup.PUT("/:id", func(ctx *gin.Context) {
				id, err := strconv.Atoi(ctx.Param("id"))
				if err != nil || id <= 0 {
					ctx.JSON(http.StatusBadRequest, map[string]interface{}{
						"message": "invalid ID",
					})
					return
				}

				var updateUser User
				if err := ctx.Bind(&updateUser); err != nil {
					ctx.JSON(http.StatusBadRequest, map[string]interface{}{
						"message": "failed to bind body",
					})
					return
				}

				found := false
				for i, user := range users {
					if user.ID == uint(id) {
						users[i].Username = updateUser.Username
						users[i].Email = updateUser.Email
						found = true
						break
					}
				}

				if found {
					ctx.JSON(http.StatusOK, map[string]interface{}{
						"message": "user updated",
					})
				} else {
					ctx.JSON(http.StatusNotFound, map[string]interface{}{
						"message": "user not found",
					})
				}
			})

			// [DELETE] /api/v1/users/:id
			usersGroup.DELETE("/:id", func(ctx *gin.Context) {
				id, err := strconv.Atoi(ctx.Param("id"))
				if err != nil || id <= 0 {
					ctx.JSON(http.StatusBadRequest, map[string]interface{}{
						"message": "invalid ID",
					})
					return
				}

				found := false
				for i, user := range users {
					if user.ID == uint(id) {
						// remove the user from the slice
						users = append(users[:i], users[i+1:]...)
						found = true
						break
					}
				}

				if found {
					ctx.JSON(http.StatusOK, map[string]interface{}{
						"message": "user deleted",
					})
				} else {
					ctx.JSON(http.StatusNotFound, map[string]interface{}{
						"message": "user not found",
					})
				}
			})

		}
	}

	engine.Run(":80")
}
