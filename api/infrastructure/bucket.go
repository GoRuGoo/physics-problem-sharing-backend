package infrastructure

import (
	"context"
	"errors"
	"io"
	"log"
	"physics/interfaces/bucket"

	"os"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

type BucketHandler struct {
	ctx    context.Context
	bucket *storage.BucketHandle
}

func NewBucketHandler() bucket.BucketHandler {

	credentionFilePath := "./physics-bucket.json"
	bucketName := "physics-problem-sharing-bucket"
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile((credentionFilePath)))
	if err != nil {
		log.Fatal(err.Error())
	}

	bucket := client.Bucket(bucketName)

	return &BucketHandler{ctx: ctx, bucket: bucket}
}

func (h *BucketHandler) WriteExecute(objectName string, fileData os.File) error {
	obj := h.bucket.Object(objectName)
	writer := obj.NewWriter(h.ctx)
	_, err := io.Copy(writer, &fileData)
	if err != nil {
		return errors.New(err.Error())
	}
	err = writer.Close()
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}
