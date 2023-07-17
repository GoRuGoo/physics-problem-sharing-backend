package infrastructure

import (
	"context"
	"errors"
	"io"
	"log"
	"mime/multipart"
	"physics/interfaces/bucket"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
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

func (h *BucketHandler) WriteExecute(objectName string, fileData multipart.File) error {
	obj := h.bucket.Object(objectName)
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
	obj := h.bucket.Object(objectName)
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

func (h *BucketHandler) SelectAllObjectsExecute() ([]string, error) {
	var return_object_list []string
	obj := h.bucket.Objects(h.ctx, nil)

	for {
		attrs, err := obj.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return return_object_list, errors.New(err.Error())
		}

		return_object_list = append(return_object_list, attrs.Name)
	}
	return return_object_list, nil

}

func (h *BucketHandler) SelectSpecificObjectsExecute(dirName string) ([]string, error) {
	var return_object_list []string
	obj := h.bucket.Objects(h.ctx, &storage.Query{
		Prefix: dirName,
	})
	for {
		attrs, err := obj.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return return_object_list, errors.New(err.Error())
		}

		return_object_list = append(return_object_list, attrs.Name)
	}
	return return_object_list, nil
}
