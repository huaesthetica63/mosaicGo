package mosaic_server

import (
	"bytes"
	"image"
	"image/png"
	"io/ioutil"
	"log"
	"main/color_mosaic"
	"main/image_processing"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
}

func (s *Server) Load() {
	r := gin.Default()
	r.LoadHTMLGlob("frontend/*.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.POST("/mosaic", func(c *gin.Context) {
		formFile, _ := c.FormFile("img")
		openedFile, _ := formFile.Open()
		file, _ := ioutil.ReadAll(openedFile)
		img, _, err := image.Decode(bytes.NewReader(file))
		if err != nil {
			log.Fatalln(err)
		}
		image_processing.SaveToPng(img, "image.png")
		mzk := color_mosaic.NewPeachMosaic()
		var im image_processing.Image
		im.LoadImage("image.png")
		res := mzk.MakeMosaic(im)
		buf := new(bytes.Buffer)
		png.Encode(buf, res)
		c.Writer.Write(buf.Bytes())
	})
	r.Run(":8080")
}
