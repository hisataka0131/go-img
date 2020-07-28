package controller

import (
	"crypto/sha1"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/olahol/go-imageupload"
	"go-img/model"
	"net/http"
	"path/filepath"
	"time"
)

type ImgHandler struct {
	Db *gorm.DB
}

func (handler *ImgHandler) GetAll(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", c)
}

func (handler *ImgHandler) Upload(c *gin.Context) {

	handler.Db.Create(&model.Image{Filename: ""})
	img, err := imageupload.Process(c.Request, "file")

	if err != nil {
		panic(err)

	}

	thumb, err := imageupload.ThumbnailJPEG(img, 300, 300, 90)

	if err != nil {
		panic(err)
	}

	h := sha1.Sum(thumb.Data)
	imgName := fmt.Sprintf("%s_%x.jpg", time.Now().Format("20060102150405"), h[:4])
	savePath := filepath.Join("./public/uploads", imgName)

	thumb.Save(savePath)
	c.Redirect(http.StatusMovedPermanently, "/")
}
