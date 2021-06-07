package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olzhas-b/hello_world_crud/database"
	"github.com/olzhas-b/hello_world_crud/models"
	"log"
	"net/http"
)
func GetUser(c *gin.Context)  {
	db := database.GetDataBase()
	var users []models.User
	db.Find(&users)
	response := ""
	for _, user := range users {
		response += fmt.Sprintf("username: %s, id: %d Phone: %s", user.Name, user.ID, user.Phone)
		response += "\n"
	}
	c.String(200, response)
	return
}

func GetUserByID(c *gin.Context) {
	id := c.Param("id")

	db := database.GetDataBase()
	var user models.User
	db.First(&user, "ID=?", id)
	c.String(200, fmt.Sprintf("username: %s, id: %d Phone: %s", user.Name, user.ID, user.Phone))
	return
}

func PostUser(c *gin.Context)  {
	db := database.GetDataBase()
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&user)
	fmt.Printf("id: %d; name: %s; login: %s; phone: %s", user.ID, user.Name, user.Email, user.Phone)
	c.String(200, fmt.Sprintf("id: %d; name: %s; login: %s; phone: %s", user.ID, user.Name, user.Email, user.Phone))
	return
}

func PutUser(c *gin.Context) {
	db := database.GetDataBase()
	var updateInfo models.UpdateInfo
	if err := c.ShouldBindJSON(&updateInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var user models.User
	db.First(&user, "id=?", &updateInfo.ID)
	if user.Password == updateInfo.Password {
		result := db.Model(&user).Update("Name", updateInfo.Name)
		if result.Error != nil {
			log.Println("Failed to change Name")
		}
	} else {
		log.Println("Wrong Password")
	}
}

func DeleteUser(c *gin.Context) {
	db := database.GetDataBase()
	var id models.ID
	if err := c.ShouldBindJSON(&id); err != nil {
		c.JSON(200, gin.H{"error": err.Error()})
	}
	db.Delete(&models.User{}, id)
	c.JSON(200, gin.H{"deleted item id: ": id})
	return
}
