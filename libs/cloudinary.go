package libs

import (
	"context"

	"mime/multipart"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/google/uuid"
)

type CloudinaryCredentials struct {
	Name   string
	APIKey string
	Secret string
}

func GetCloudinaryCredentials() CloudinaryCredentials {
	return CloudinaryCredentials{
		Name:   os.Getenv("CLOUD_NAME"),
		APIKey: os.Getenv("CLOUD_APIKEY"),
		Secret: os.Getenv("CLOUD_SECRET"),
	}
}

func UploadImage(file multipart.File) (string, string, error) {

	ctx := context.Background()
	imageId := uuid.New().String()

	creds := GetCloudinaryCredentials()

	cld, err := cloudinary.NewFromParams(creds.Name, creds.APIKey, creds.Secret)
	if err != nil {
		return "", "", err
	}

	resp, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{PublicID: imageId})
	if err != nil {
		return "", "", err
	}
	return resp.SecureURL, imageId, nil

}

func DeleteImage(image_id string) (string, error) {

	ctx := context.Background()
	creds := GetCloudinaryCredentials()

	cld, err := cloudinary.NewFromParams(creds.Name, creds.APIKey, creds.Secret)
	if err != nil {
		return "", err
	}
	resp, err := cld.Upload.Destroy(ctx, uploader.DestroyParams{PublicID: image_id})
	if err != nil {
		return "", err
	}

	return resp.Result, nil

}
