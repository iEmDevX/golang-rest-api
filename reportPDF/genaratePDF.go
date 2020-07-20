package reportPDF

import "github.com/gin-gonic/gin"

func GetRoute(r *gin.Engine) {
	pathPDF := r.Group("/pdf")
	{
		pathPDF.GET("/gen", genarate)
	}
}

func genarate(r *gin.Context) {

}