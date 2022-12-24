package mosaic_server

import (
	"bytes"
	"image/png"
	"io/ioutil"
	"log"
	"main/color_mosaic"
	"main/image_processing"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Server struct {
	address string
	handler *gin.Engine
	timeout time.Duration
}

func getHandler() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("frontend/*.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.POST("/mosaic", func(c *gin.Context) {
		formFile, err := c.FormFile("img")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var mzk color_mosaic.Mosaic
		chose := c.Request.Form.Get("Mosaic")
		size := c.Request.Form.Get("size")
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
		openedFile, err := formFile.Open()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		file, err := ioutil.ReadAll(openedFile)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var im image_processing.Image
		im.LoadImageBytes(file)
		switch size {
		case "original":
		case "1000x1000":
			im = im.ResizeImage(1000, 1000)
		case "500x500":
			im = im.ResizeImage(500, 500)
		case "250x250":
			im = im.ResizeImage(250, 250)
		default:

		}
		res := mzk.MakeMosaic(im)
		buf := new(bytes.Buffer)
		png.Encode(buf, res)
		c.Writer.Write(buf.Bytes())
	})
	return r
}
func NewServer(addr string, timeout time.Duration) Server {
	handler := getHandler()
	return Server{
		address: addr,
		handler: handler,
		timeout: timeout,
	}
}
func (s *Server) Load() {
	serv := http.Server{
		Addr:        s.address,
		ReadTimeout: s.timeout,
		Handler:     s.handler,
	}
	if err := serv.ListenAndServe(); err != nil {
		log.Println("Server is running: ", s.address)
	}
}
