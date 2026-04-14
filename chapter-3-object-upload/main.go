package main

import (
    "bytes"
    "context"
    "fmt"
    "io"
    "log"

    "github.com/minio/minio-go/v7"
    "github.com/minio/minio-go/v7/pkg/credentials"
)

type progress struct{}

func newProgressListener() io.Writer { return &progress{} }
func (p *progress) Write(b []byte) (int, error) {
    fmt.Print(".")
    return len(b), nil
}

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
    
    //usual check for bucket
    err = client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
    if err != nil {
        log.Println("Bucket note:", err)
    }

    // Small upload
    data := []byte("Hello World from Go!")
    _, err = client.PutObject(ctx, bucketName, "greeting.txt", 
        bytes.NewReader(data), int64(len(data)), 
        minio.PutObjectOptions{ContentType: "text/plain"})
    if err != nil {
        log.Fatalln(err)
    }
    log.Println("\nuploaded greeting.txt")

    // Large upload with progress
    largeData := bytes.Repeat([]byte("x"), 10*1024*1024)
    _, err = client.PutObject(ctx, bucketName, "large-file.bin",
        bytes.NewReader(largeData), int64(len(largeData)),
        minio.PutObjectOptions{
            ContentType: "application/octet-stream",
            Progress:    newProgressListener(),
        })
    if err != nil {
        log.Fatalln(err)
    }
    log.Println("\nuploaded large-file.bin")
}