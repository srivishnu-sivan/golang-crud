package main


 import (
    "go-crud/controllers"
    "go-crud/database"
    "go-crud/models"

    "github.com/gin-gonic/gin"
)


func main(){
	database.ConnectDB()
	r := gin.Default()

	// initialize automigration
	database.DB.AutoMigrate(&models.User{})


	// router initlization
	r.POST("/createUser", controllers.CreateUser)
	r.GET("/getusers", controllers.GetUsers)
	r.GET("/users", controllers.GetUserByQuery)
	r.PUT("/users/", controllers.UpdateUser)
	r.DELETE("/users/", controllers.DeleteUser)





	r.Run(":8082")
}