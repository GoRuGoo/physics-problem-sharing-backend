package infrastructure

import (
	"context"
	"errors"
	"io"
	"log"
	"mime/multipart"
	"physics/interfaces/bucket"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

type BucketHandler struct {
	ctx    context.Context
	client *storage.Client
}

func NewBucketHandler() bucket.BucketHandler {

	credentionFilePath := "./physics-bucket.json"
	//bucketName := "physics-problem-sharing-bucket"
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile((credentionFilePath)))
	if err != nil {
		log.Fatal(err.Error())
	}

	return &BucketHandler{ctx: ctx, client: client}
}

func (h *BucketHandler) WriteExecute(objectName string, fileData multipart.File) error {
	obj := h.client.Bucket("physics-problem-sharing-bucket").Object(objectName)
	writer := obj.NewWriter(h.ctx)
	_, err := io.Copy(writer, fileData)
	if err != nil {
		return errors.New(err.Error())
	}
	err = writer.Close()
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

func (h *BucketHandler) DeleteExecute(objectName string) error {
	obj := h.client.Bucket("physics-problem-sharing-bucket").Object(objectName)
	attrs, err := obj.Attrs(h.ctx)
	if err != nil {
		return errors.New(err.Error())
	}
	obj = obj.If(storage.Conditions{GenerationMatch: attrs.Generation})

	err = obj.Delete(h.ctx)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}
