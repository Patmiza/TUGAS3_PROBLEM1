package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users []User

// -------------------- controller --------------------

// get all users

func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "sukses",
		"users":   users,
	})
}

func CreateUserController(c echo.Context) error {
	newUser := User{}
	c.Bind(&newUser)

	//Id increment
	if len(users) == 0 {
		newUser.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		newUser.Id = newId
	}

	users = append(users, newUser)
	return c.JSON(http.StatusOK, users)
}

//get user by Id
func GetUserController(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "tidak ada ID",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"id":   id,
		"name": users[id-1].Name,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "tidak ada ID",
		})
	}

	var newUsers []User
	for i := 0; i < len(users); i++ {
		if i == id {
			continue
		} else {
			newUsers = append(newUsers, users[i])
		}
	}

	return c.JSON(http.StatusOK, newUsers)
}

// update user by id
func UpdateUserController(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "tidak ada ID",
		})
	}
	updateUser := new(User)
	c.Bind(updateUser)
	users[id].Name = updateUser.Name
	return c.JSON(http.StatusOK, users[id])
}

// ---------------------------------------------------
func main() {

	users = []User{}
	e := echo.New()
	//e.Server.ListenAndServe()
	e.GET("/users", GetUsersController)
	e.POST("/users", CreateUserController)
	e.GET("/users/:id", GetUserController)
	e.DELETE("/users/:id", DeleteUserController)
	e.PUT("/users/:id", UpdateUserController)
	e.Logger.Fatal(e.Start(":8000"))

}
