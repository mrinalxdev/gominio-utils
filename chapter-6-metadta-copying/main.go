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

    bucketName := "my-course-bucket"
    src := "greeting.txt"
    dst := "greeting-backup.txt"

    // source Options
    srcOpts := minio.CopySrcOptions{
        Bucket: bucketName,
        Object: src,
    }

    // destination Options with Custom Metadata
    dstOpts := minio.CopyDestOptions{
        Bucket: bucketName,
        Object: dst,
        UserMetadata: map[string]string{
            "Original-File": src,
            "Author":        "GoStudent",
        },
    }

    // perform Copy
    _, err = client.CopyObject(ctx, dstOpts, srcOpts)
    if err != nil {
        log.Fatalln(err)
    }
    log.Printf("Copied %s to %s with metadata.", src, dst)

    // set User Tags
    // note -- this requires the object to exist
    err = client.PutObjectTagging(ctx, bucketName, dst, map[string]string{
        "Environment": "Development",
        "Project":     "MinIO-Course",
    }, minio.PutObjectTaggingOptions{})
    if err != nil {
        log.Fatalln(err)
    }
    log.Println("Tags applied")
}