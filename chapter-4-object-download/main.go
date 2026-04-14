package main

import (
    "context"
    "io"
    "log"
    "os"
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

    // download directly to file
    err = client.FGetObject(ctx, bucketName, objectName, "downloaded-greeting.txt", minio.GetObjectOptions{})
    if err != nil {
        log.Fatalln(err)
    }
    log.Println("Successfully downloaded to downloaded-greeting.txt")

    // download to memory / Stream
    obj, err := client.GetObject(ctx, bucketName, objectName, minio.GetObjectOptions{})
    if err != nil {
        log.Fatalln(err)
    }
    defer obj.Close()

    // Read first 5 bytes
    buffer := make([]byte, 5)
    n, err := obj.Read(buffer)
    if err != nil && err != io.EOF {
        log.Fatalln(err)
    }
    log.Printf("Read first %d bytes: %s\n", n, string(buffer))

    // Clean up local file
    os.Remove("downloaded-greeting.txt")
}