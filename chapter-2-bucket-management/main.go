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

    /*
     * creating buckets and validating if the same name exists
     * after that will be just listing out all the buckets available
     */
    err = client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
    if err != nil {
        exists, errBucketExists := client.BucketExists(ctx, bucketName)
        if errBucketExists == nil && exists {
            log.Printf("bucket %s already exists.\n", bucketName)
        } else {
            log.Fatalln(err)
        }
    } else {
        log.Printf("done creating %s\n", bucketName)
    }
    buckets, err := client.ListBuckets(ctx)
    if err != nil {
        log.Fatalln(err)
    }

    for _, bucket := range buckets {
        log.Printf("bucket -- %s, createdd -- %s\n", bucket.Name, bucket.CreationDate)
    }
}