package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"time"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

func main() {
	credentialFilePath := "./physics-bucket.json"
	bucketName := "physics-problem-sharing-bucket"
	objectName := "ore/tettttttttttt.txt"
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(credentialFilePath))
	if err != nil {

		log.Fatal(err.Error())
	}
	defer client.Close()

	// Open local file.
	f, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer f.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	o := client.Bucket(bucketName).Object(objectName)

	// Optional: set a generation-match precondition to avoid potential race
	// conditions and data corruptions. The request to upload is aborted if the
	// object's generation number does not match your precondition.
	// For an object that does not yet exist, set the DoesNotExist precondition.
	o = o.If(storage.Conditions{DoesNotExist: true})
	// If the live object already exists in your bucket, set instead a
	// generation-match precondition using the live object's generation number.
	// attrs, err := o.Attrs(ctx)
	// if err != nil {
	//      return fmt.Errorf("object.Attrs: %w", err)
	// }
	// o = o.If(storage.Conditions{GenerationMatch: attrs.Generation})

	// Upload an object with storage.Writer.
	wc := o.NewWriter(ctx)
	if _, err = io.Copy(wc, f); err != nil {
		log.Fatal(err.Error())
	}
	if err := wc.Close(); err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("done")
}
