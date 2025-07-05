package controllers



import (
	"go-crud/models"
	"go-crud/database"
	"net/http"

	"github.com/gin-gonic/gin"

)


func CreateUser(c *gin.Context){
	db := database.DB
	var user models.User
	if err := c.ShouldBindJSON(&user); err !=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()} )
		return
	}
	db.Create(&user)
	c.JSON(http.StatusOK, gin.H{"message" : "User Creaed Successfully"})
}


func GetUsers(c *gin.Context){
	db :=database.DB	
	var users []models.User
	db.Find(&users)
	c.JSON(http.StatusOK, users)
}


func GetUserByQuery(c *gin.Context) {
	db := database.DB
	id := c.Query("id")
	name := c.Query("name")
	email := c.Query("email")

	var user models.User

	if id != "" {
		if err := db.Where("id = ?", id).First(&user).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
	} else if name != "" {
		if err := db.Where("name = ?", name).First(&user).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
	} else if email != "" {
		if err := db.Where("email = ?", email).First(&user).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Please provide id, name, or email as query parameter"})
		return
	}

	c.JSON(http.StatusOK, user)
}



func UpdateUser(c *gin.Context) {
	db := database.DB
	id := c.Query("id")

	var user models.User
	if err := db.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found!"})
		return
	}

	type UpdateUserInput struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	var input UpdateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	user.Name = input.Name
	user.Email = input.Email

	db.Save(&user)

	c.JSON(http.StatusOK, gin.H{
		"message": "User updated successfully",
		"user":    user,
	})
}


func DeleteUser(c * gin.Context){
	var user models.User
	db := database.DB
	id := c.Query("id")
	if err := db.First(&user, id).Error; err != nil{
		c.JSON(http.StatusNotFound, gin.H{"error" : "User not found!"})
		return
	}
	db.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"message" : "User deleted successfully"})

}