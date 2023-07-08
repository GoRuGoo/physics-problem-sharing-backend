package test

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

func test() {
	credentialFilePath := "./physics-bucket.json"

	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(credentialFilePath))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	bucketName := "physics-problem-sharing-bucket"
	objectName := "ore/test.txt"

	// バケットの参照を取得
	bucket := client.Bucket(bucketName)

	// ファイルを開く
	file, err := os.Open("sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// オブジェクトを作成し、データを書き込む
	obj := bucket.Object(objectName)
	writer := obj.NewWriter(ctx)
	if _, err := writer.Write([]byte{}); err != nil {
		log.Fatal(err)
	}
	if err := writer.Close(); err != nil {
		log.Fatal(err)
	}

	log.Printf("File %s uploaded to bucket %s.", objectName, bucketName)
	log.Println("done")
}
