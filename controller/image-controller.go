package controller


import (
	// "fmt"
	"net/http"
	"os"

	"goapi/utilService"
	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/gin-gonic/gin"
	"encoding/base64"
	"bytes"
)

// image upload controller
func UploadImage(ctx *gin.Context) {
    // 1. Get the image data from the request body
    var inputData struct {
        Input struct {
            Arg1 struct {
                Images []string `json:"images"`
            } `json:"arg1"`
        } `json:"input"`
    }

    if err := ctx.ShouldBindJSON(&inputData); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // 2. Set up the Cloudinary configuration
    cld, _ := cloudinary.NewFromParams(os.Getenv("CLOUDINARY_CLOUD_NAME"), os.Getenv("CLOUDINARY_API_KEY"), os.Getenv("CLOUDINARY_SECRET"))
    var images = inputData.Input.Arg1.Images

    // 3. Upload images to Cloudinary and store the URLs in an array
    var urls []string
    for _, imageBase64 := range images {
        // Decode the base64 image data into binary
        imageBinary, err := base64.StdEncoding.DecodeString(imageBase64)
        if err != nil {
            ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        // Upload the binary image to Cloudinary
        response, err := cld.Upload.Upload(ctx.Request.Context(), bytes.NewReader(imageBinary), uploader.UploadParams{
            PublicID: utilService.PublicID(),
        })
        if err != nil {
            ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        urls = append(urls, response.SecureURL)
    }

    // 4. Send the URLs to the client
    ctx.JSON(200, gin.H{"urls": urls})
}