package main

import (
	"bytes"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jung-kurt/gofpdf"
	"github.com/signintech/gopdf"
	"github.com/unidoc/unipdf/v3/model"
)

const (
	logoPath     = "./assets/logo.png"
	uploadDir    = "./uploads"
	thumbnailDir = "./thumbnails"
)

func main() {
	r := gin.Default()
	gin.SetMode(gin.DebugMode)

	r.POST("/upload", uploadFile)

	r.Run(":8080")
}

func uploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check the file type
	fileType := getFileType(file.Filename)

	// Convert to PDF if the file is not a PDF
	var pdfFilePath string
	if fileType != "pdf" {
		pdfFilePath, err = convertToPDF(file)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	} else {
		pdfFilePath, err = saveUploadedFile(c, file)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	// Check if the PDF is password protected
	isPasswordProtected, err := isPDFPasswordProtected(pdfFilePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Handle password-protected PDF
	if isPasswordProtected {
		c.JSON(http.StatusOK, gin.H{"message": "File Protected"})
	}

	// Generate the QR Code
	grCode := generateQRCode()

	// Add QR Code watermark to the PDF
	err = addWatermarkToPDF(pdfFilePath, grCode, 470, 750)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Generate PDF thumbnail
	// err = generateThumbnail(pdfFilePath)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

	// Override PDF metadata
	// err = overrideMetadata(pdfFilePath)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded and processed successfully"})
}

func getFileType(fileName string) string {
	fileExt := strings.ToLower(filepath.Ext(fileName))
	if len(fileExt) > 1 && fileExt[0] == '.' {
		return fileExt[1:]
	}
	return fileExt
}

func saveUploadedFile(c *gin.Context, file *multipart.FileHeader) (string, error) {
	filePath := filepath.Join(uploadDir, file.Filename)
	err := c.SaveUploadedFile(file, filePath)
	if err != nil {
		return "", err
	}
	return filePath, nil
}

func convertToPDF(file *multipart.FileHeader) (string, error) {
	// Open the uploaded file
	srcFile, err := file.Open()
	if err != nil {
		return "", err
	}
	defer srcFile.Close()

	// Create the output PDF file
	pdfFilePath := filepath.Join(uploadDir, file.Filename+".pdf")
	destFile, err := os.Create(pdfFilePath)
	if err != nil {
		return "", err
	}
	defer destFile.Close()

	// Create a new PDF object
	pdf := gofpdf.New("P", "mm", "A4", "")

	// Add a new page to the PDF
	pdf.AddPage()

	// Load the file data into a byte slice
	fileData, err := io.ReadAll(srcFile)
	if err != nil {
		return "", err
	}

	// Convert the image to PDF and write it to the page
	pdf.RegisterImageOptionsReader(file.Filename, gofpdf.ImageOptions{ImageType: file.Header.Get("Content-Type")}, bytes.NewReader(fileData))
	pdf.ImageOptions(file.Filename, 0, 0, 0, 0, true, gofpdf.ImageOptions{}, 0, "")

	// Output the PDF content to the file
	err = pdf.OutputFileAndClose(destFile.Name())
	if err != nil {
		return "", err
	}

	return pdfFilePath, nil
}

func isPDFPasswordProtected(filePath string) (bool, error) {
	// Open the PDF file
	file, err := os.Open(filePath)
	if err != nil {
		return false, err
	}
	defer file.Close()

	// Create a PDF reader
	pdfReader, err := model.NewPdfReader(file)
	if err != nil {
		return false, err
	}

	// Check if the PDF is encrypted
	isEncrypted, err := pdfReader.IsEncrypted()
	if err != nil {
		return false, err
	}

	return isEncrypted, nil
}

func generateQRCode() string {
	// Generate a random UUID
	uuid := uuid.New().String()

	// Construct the full QR Code URL
	grCodeURL := "https://privy.id/verifyKrandom-" + uuid

	// Add PrivyID logo to the QR Code
	grCodeWithLogo := addLogoToQRCode(grCodeURL)

	return grCodeWithLogo
}

func addLogoToQRCode(grCodeURL string) string {
	// Implement code to add the PrivyID logo to the QR Code URL
	// You can use a third-party library or API to overlay the logo onto the QR Code image

	// Placeholder implementation
	return grCodeURL + " [PrivyID Logo]"
}

func convertToGrayscales(img image.Image) image.Image {
	bounds := img.Bounds()
	gray := image.NewGray(bounds)

	draw.Draw(gray, bounds, img, bounds.Min, draw.Src)
	return gray
}

func addWatermarkToPDF(filePath, grCode string, x, y float64) error {
	// Generate the barcode image from the QR Code
	qrCode, err := qr.Encode(grCode, qr.M, qr.Auto)
	if err != nil {
		return err
	}

	qrCode, err = barcode.Scale(qrCode, 125, 125)
	if err != nil {
		return err
	}

	// Convert the barcode image to grayscale
	grayQrCode := convertToGrayscale(qrCode)

	// Load the logo image
	logoFile, err := os.Open("./assets/logo.png")
	if err != nil {
		return err
	}
	defer logoFile.Close()

	logoImg, _, err := image.Decode(logoFile)
	if err != nil {
		return err
	}

	// Calculate the position to place the logo at the center
	logoWidth := logoImg.Bounds().Max.X
	logoHeight := logoImg.Bounds().Max.Y

	qrCodeWidth := grayQrCode.Bounds().Max.X
	qrCodeHeight := grayQrCode.Bounds().Max.Y

	logoX := (qrCodeWidth - logoWidth) / 2
	logoY := (qrCodeHeight - logoHeight) / 2

	// Create a new image with the QR code and logo
	qrCodeWithLogo := image.NewRGBA(image.Rect(0, 0, qrCodeWidth, qrCodeHeight))
	draw.Draw(qrCodeWithLogo, qrCodeWithLogo.Bounds(), grayQrCode, image.Point{}, draw.Src)
	draw.Draw(qrCodeWithLogo, qrCodeWithLogo.Bounds().Add(image.Pt(logoX, logoY)), logoImg, image.Point{}, draw.Over)

	// Save the QR code with the logo as a PNG file
	file, err := os.Create("qrcode_with_logo.png")
	if err != nil {
		return err
	}
	defer file.Close()

	err = png.Encode(file, qrCodeWithLogo)
	if err != nil {
		return err
	}

	// Create a new PDF document
	pdf := gopdf.GoPdf{}

	// Initialize a new configuration
	config := gopdf.Config{
		PageSize: *gopdf.PageSizeA4,
	}

	// Set the configuration
	pdf.Start(config)

	// Open the existing PDF file
	pdf.AddPage()

	// Add the QR code with the logo as a watermark
	err = pdf.Image("qrcode_with_logo.png", x, y, nil)
	if err != nil {
		return err
	}

	// Add a TrueType font
	err = pdf.AddTTFFont("arial", "fonts/arial.ttf")
	if err != nil {
		return err
	}

	// Set the font
	err = pdf.SetFont("arial", "", 12)
	if err != nil {
		return err
	}

	// Save the modified PDF with the watermark
	err = pdf.WritePdf(fmt.Sprintf("%s_with_watermark.pdf", filePath))
	if err != nil {
		return err
	}

	return nil
}

func convertToGrayscale(img image.Image) *image.Gray {
	bounds := img.Bounds()
	grayImg := image.NewGray(bounds)

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			grayImg.Set(x, y, img.At(x, y))
		}
	}

	return grayImg
}

// func generateThumbnail(filePath string) error {
// 	// Open the PDF file
// 	file, err := os.Open(filePath)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	// Create a new PDF reader
// 	pdfReader, err := model.NewPdfReader(file)
// 	if err != nil {
// 		return err
// 	}

// 	// Get the number of pages in the PDF
// 	numPages, err := pdfReader.GetNumPages()
// 	if err != nil {
// 		return err
// 	}

// 	// Check if the PDF has any pages
// 	if numPages < 1 {
// 		return fmt.Errorf("PDF has no pages")
// 	}

// 	// Load the first page of the PDF
// 	pageNum := 1
// 	page, err := pdfReader.GetPage(pageNum)
// 	if err != nil {
// 		return err
// 	}

// 	// Render the page to an image
// 	imgWidth := 720
// 	imgHeight := 360

// 	// Scale the page content to fit the thumbnail dimensions
// 	pageContentWidth, pageContentHeight, err := page.GetContentBox()
// 	if err != nil {
// 		return err
// 	}
// 	scaleX := float64(imgWidth) / float64(pageContentWidth)
// 	scaleY := float64(imgHeight) / float64(pageContentHeight)

// 	img := image.NewRGBA(image.Rect(0, 0, imgWidth, imgHeight))
// 	gopdfPage := model.GopdfPageFromPage(page)
// 	err = gopdfPage.Draw(page.GetResources(), img, nil, nil, scaleX, scaleY)
// 	if err != nil {
// 		return err
// 	}

// 	// Create a new file for the thumbnail
// 	thumbnailPath := filePath + "_thumbnail.jpg"
// 	thumbnailFile, err := os.Create(thumbnailPath)
// 	if err != nil {
// 		return err
// 	}
// 	defer thumbnailFile.Close()

// 	// Encode the image as JPEG and write it to the file
// 	err = jpeg.Encode(thumbnailFile, img, nil)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func overrideMetadata(filePath string) error {
// 	// Implement code to override the PDF metadata
// }
