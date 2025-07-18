package r2

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"time"

	"github.com/minio/minio-go/v7"
)

type UploadResult struct {
	URL      string
	FileName string
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
	defer func() {
		if cerr := file.Close(); cerr != nil {
			fmt.Printf("Error closing file: %v\n", cerr)
		}
	}()

	var buf bytes.Buffer
	if _, err := io.Copy(&buf, file); err != nil {
		return nil, fmt.Errorf("failed to read file into buffer: %w", err)
	}

	fileName := fmt.Sprintf("%s/%d_%s", folder, time.Now().UnixNano(), fileHeader.Filename)

	_, err = u.Client.PutObject(context.Background(), u.BucketName, fileName, &buf, int64(buf.Len()), minio.PutObjectOptions{
		ContentType: fileHeader.Header.Get("Content-Type"),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to upload to R2: %w", err)
	}

	url := fmt.Sprintf("https://%s/%s/%s", u.Client.EndpointURL().Host, u.BucketName, fileName)

	return &UploadResult{
		URL:      url,
		FileName: fileName,
	}, nil
}

func (u *Uploader) Delete(fileName string) error {
	err := u.Client.RemoveObject(context.Background(), u.BucketName, fileName, minio.RemoveObjectOptions{})
	if err != nil {
		return fmt.Errorf("failed to delete file %s: %w", fileName, err)
	}
	return nil
}
