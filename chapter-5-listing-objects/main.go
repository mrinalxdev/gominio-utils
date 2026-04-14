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

    // list all objects
    objectsCh := client.ListObjects(ctx, bucketName, minio.ListObjectsOptions{Recursive: true})

    log.Println("Objects in bucket:")
    for object := range objectsCh {
        if object.Err != nil {
            log.Fatalln(object.Err)
        }
        log.Printf("Name: %s, Size: %d, LastModified: %s\n", object.Key, object.Size, object.LastModified)
    }

    // list with Prefix
    prefixCh := client.ListObjects(ctx, bucketName, minio.ListObjectsOptions{
        Prefix:    "greet",
        Recursive: false,
    })

    log.Println("\nObjects with prefix 'greet':")
    for object := range prefixCh {
        if object.Err != nil {
            log.Fatalln(object.Err)
        }
        log.Println(object.Key)
    }
}