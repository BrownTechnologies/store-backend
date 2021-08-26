package controller

import (
	"encoding/csv"
	"store-backend/dynamodb"
	"store-backend/modals"

	"io"

	"log"
	"net/http"
	"os"
	"path/filepath"

	"store-backend/utils"

	"github.com/gin-gonic/gin"
)

func UpdateZipCodes(context *gin.Context) {

	url := context.PostForm("url")
	if url == "" {
		// set default
		url = "https://www.post.japanpost.jp/zipcode/dl/roman/ken_all_rome.zip?210622"
	}

	zipPath, err := utils.DownloadFile(url)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}
	defer os.Remove(zipPath)

	contentDir := filepath.Dir(zipPath) + "/zip-data/"
	err = utils.Unzip(zipPath, contentDir)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}

	files, _ := os.ReadDir(contentDir)
	reader, _ := os.Open(contentDir + files[0].Name())
	r := csv.NewReader(reader)

	db := dynamodb.New()
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil || len(record) != 7 {
			log.Fatal("error while reading record or the record is improper.")
		} else {
			entryJp := modals.ZipCodeEntry{
				Pincode:    record[0],
				Lang:       "jp",
				City:       utils.TransformIntoKana(record[3]),
				Area:       utils.TransformIntoKana(record[2]),
				Prefecture: utils.TransformIntoKana(record[1]),
			}
			entryEn := modals.ZipCodeEntry{
				Pincode:    record[0],
				Lang:       "en",
				City:       record[6],
				Area:       record[5],
				Prefecture: record[4],
			}

			db.InsertIntoZipcode(entryJp)
			db.InsertIntoZipcode(entryEn)
		}
	}

	context.JSON(http.StatusOK, gin.H{
		"message": url,
	})
}
