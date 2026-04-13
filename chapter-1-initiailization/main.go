package main

import (
    "context"
    "log"
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
    log.Printf("client initialized successfully: %T", client)
    bucketName := "myTesting-bucket" // setting up bucket name and verfiying the connection
    exists, err := client.BucketExists(ctx, bucketName)
    if err != nil {
        log.Fatalln(err)
    }
    log.Printf("connections verified. bucket '%v' exists: %v", bucketName, exists)
}