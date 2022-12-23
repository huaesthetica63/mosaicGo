package mosaic_server

import (
	"bytes"
	"image/png"
	"io/ioutil"
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
		var mzk color_mosaic.Mosaic
		chose := c.Request.Form.Get("Mosaic")
		switch chose {
		case "grayscale":
			mzk = color_mosaic.NewGrayscaleMosaic(3)
		case "blue":
			mzk = color_mosaic.NewBlueMosaic()
		case "red":
			mzk = color_mosaic.NewRedMosaic()
		case "green":
			mzk = color_mosaic.NewGreenMosaic()
		case "peach":
			mzk = color_mosaic.NewPeachMosaic()
		default:
			mzk = color_mosaic.NewGrayscaleMosaic(3)
		}
		openedFile, _ := formFile.Open()
		file, _ := ioutil.ReadAll(openedFile)
		var im image_processing.Image
		im.LoadImageBytes(file)
		res := mzk.MakeMosaic(im)
		buf := new(bytes.Buffer)
		png.Encode(buf, res)
		c.Writer.Write(buf.Bytes())
	})
	r.Run(":8080")
}
