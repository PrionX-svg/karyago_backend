package r2

import (
	"bytes"
	"context"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"mime/multipart"
	"net/url"
	"path/filepath"
	"strings"
	"time"

	"github.com/chai2010/webp"
	"github.com/minio/minio-go/v7"
)

type UploadResult struct {
	URL      string `json:"url"`
	FileName string `json:"file_name"`
}

type Uploader struct {
	Client     *minio.Client
	BucketName string
}

func NewUploader(client *minio.Client, bucket string) *Uploader {
	return &Uploader{
		Client:     client,
		BucketName: bucket,
	}
}

func (u *Uploader) Upload(fileHeader *multipart.FileHeader, folder string) (*UploadResult, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return nil, fmt.Errorf("failed to open uploaded file: %w", err)
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf("failed to close file: %v\n", err)
		}
	}(file)

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("failed to decode image: %w", err)
	}

	var webpBuf bytes.Buffer
	if err := webp.Encode(&webpBuf, img, &webp.Options{Lossless: false, Quality: 80}); err != nil {
		return nil, fmt.Errorf("failed to encode to webp: %w", err)
	}

	originalName := strings.TrimSuffix(fileHeader.Filename, filepath.Ext(fileHeader.Filename))
	fileName := fmt.Sprintf("%s/%d_%s.webp", folder, time.Now().UnixNano(), originalName)

	_, err = u.Client.PutObject(
		context.Background(),
		u.BucketName,
		fileName,
		&webpBuf,
		int64(webpBuf.Len()),
		minio.PutObjectOptions{
			ContentType: "image/webp",
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to upload to R2: %w", err)
	}

	reqParams := make(url.Values)
	reqParams.Set("response-content-type", "image/webp")

	presignedURL, err := u.Client.PresignedGetObject(
		context.Background(),
		u.BucketName,
		fileName,
		5*time.Minute,
		reqParams,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to generate presigned URL: %w", err)
	}

	return &UploadResult{
		URL:      presignedURL.String(),
		FileName: fileName,
	}, nil
}

func (u *Uploader) GetPresignedURL(fileName string) (string, error) {
	reqParams := make(url.Values)
	reqParams.Set("response-content-type", "image/webp")

	presignedURL, err := u.Client.PresignedGetObject(
		context.Background(),
		u.BucketName,
		fileName,
		5*time.Minute,
		reqParams,
	)
	if err != nil {
		return "", fmt.Errorf("failed to generate presigned URL: %w", err)
	}
	return presignedURL.String(), nil
}

func (u *Uploader) Delete(fileName string) error {
	fileName = strings.TrimPrefix(fileName, "/")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := u.Client.RemoveObject(ctx, u.BucketName, fileName, minio.RemoveObjectOptions{})
	if err != nil {
		return fmt.Errorf("failed to delete file %s: %w", fileName, err)
	}

	_, statErr := u.Client.StatObject(ctx, u.BucketName, fileName, minio.StatObjectOptions{})
	if statErr == nil {
		return fmt.Errorf("file %s still exists after deletion", fileName)
	}
	if minio.ToErrorResponse(statErr).Code != "NoSuchKey" {
		return fmt.Errorf("unexpected error after deletion: %w", statErr)
	}

	return nil
}
