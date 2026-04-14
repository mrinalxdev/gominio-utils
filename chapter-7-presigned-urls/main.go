package main

import (
    "context"
    "log"
    "time"
    "github.com/minio/minio-go/v7"
    "github.com/minio/minio-go/v7/pkg/credentials"
)

func main() {
    ctx := context.Background()
    client, err := minio.New("localhost:9000", &minio.Options{
        Creds:  credentials.NewStaticV4("minioadmin", "minioadmin", ""),
        Secure: false,
    })
    if err != nil {
        log.Fatalln(err)
    }

    bucketName := "my-course-bucket"
    objectName := "greeting.txt"

    // presigned GET URL (Download)
    // Valid for 24 hours
    presignedGetURL, err := client.PresignedGetObject(ctx, bucketName, objectName, time.Duration(24*time.Hour), nil)
    if err != nil {
        log.Fatalln(err)
    }
    log.Printf("download URL: \n%s\n", presignedGetURL)

    // presigned PUT URL (Upload)
    // allows anyone with this link to upload to 'new-upload.txt' without credentials
    newObject := "new-upload.txt"
    presignedPutURL, err := client.PresignedPutObject(ctx, bucketName, newObject, time.Duration(24*time.Hour))
    if err != nil {
        log.Fatalln(err)
    }
    log.Printf("upload URL: \n%s\n", presignedPutURL)
}