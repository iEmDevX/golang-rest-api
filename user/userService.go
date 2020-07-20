package user

import (
	"awesomeProject/config"
	"awesomeProject/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetRoute(r *gin.Engine) {
	pathUser := r.Group("/user")
	{
		// Save User
		pathUser.POST("/save", save)
		pathUser.GET("/list", list)
		pathUser.GET("/findID/:ID", findID)
	}
}

/*
	See more http://gorm.io/docs/create.html
*/
func save(c *gin.Context) {
	user := entity.User{Username: "Test Username", FullName: "Name Surname"}
	// Save User
	config.DB.Create(&user)
	// Set Return
	c.SecureJSON(http.StatusOK, gin.H{
		"status":  "SUCCESS",
		"message": "SUCCESS",
		"data":    user.ID,
	})
}

/*
	See more http://gorm.io/docs/query.html
*/
func list(c *gin.Context) {
	var users []entity.User
	find := config.DB.Find(&users)

	if find.Error != nil {
		c.JSON(http.StatusOK, find.Error)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "SUCCESS",
		"message": "SUCCESS",
		"data":    find.Value,
	})
}

/*
	Find By ID
 */
func findID(c *gin.Context) {
	find := config.DB.Where("ID = " + c.Param("ID")).First(&entity.User{})

	c.JSON(http.StatusOK, gin.H{
		"status":  "SUCCESS",
		"message": "SUCCESS",
		"data":    find.Value,
	})
}
