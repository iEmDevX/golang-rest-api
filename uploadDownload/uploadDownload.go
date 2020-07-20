package uploadDownload

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var path = "public/"

func GetRoute(r *gin.Engine) {
	/*
		Check Path And Create
	*/
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, os.ModePerm)
	}

	pathPDF := r.Group("/file")
	{
		pathPDF.POST("/upload-single", singleUpload)
		pathPDF.POST("/upload-multi", multiUpload)
		pathPDF.GET("/preview/:filename", preview)
		pathPDF.GET("/download/:filename", download)
	}
}

/*
	Send body form-data => file (must single file)
*/
func singleUpload(c *gin.Context) {
	file, _ := c.FormFile("file")
	log.Println(file.Filename)

	filename := filepath.Base(file.Filename)
	// Upload the file to specific dst.
	c.SaveUploadedFile(file, path+filename)
	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}

/*
	Send body form-data => files (can multi)
*/
func multiUpload(c *gin.Context) {
	// Multipart form
	form, err := c.MultipartForm()
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}
	files := form.File["files"]

	for _, file := range files {
		filename := filepath.Base(file.Filename)
		if err := c.SaveUploadedFile(file, path+filename);
			err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
			return
		}
	}
	c.String(http.StatusOK, fmt.Sprintf("Uploaded successfully %d files with fields.", len(files)))
}

/*
	preview Image or other file
*/
func preview(c *gin.Context) {
	fullPath := path + c.Param("filename")
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		c.String(http.StatusGone, "File Not Found ")
		return
	}
	c.File(fullPath)
}

/*
	download file
*/
func download(c *gin.Context) {
	fileName := c.Param("filename")
	baf, err := ioutil.ReadFile(path + fileName)
	if err != nil {
		c.String(http.StatusGone, "File Not Found ")
		return
	}

	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Data(http.StatusOK, "application/octet-stream", baf)
}
