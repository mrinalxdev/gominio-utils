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

    // remove single Object
    err = client.RemoveObject(ctx, bucketName, "greeting-backup.txt", minio.RemoveObjectOptions{})
    if err != nil {
        log.Println("Error removing greeting-backup.txt (might not exist):", err)
    } else {
        log.Println("Removed greeting-backup.txt")
    }

    // remove multiple Objects
    objectsToDelete := []minio.ObjectInfo{
        {Key: "greeting.txt"},
        {Key: "large-file.bin"},
    }

    // RemoveObjects returns a channel of error information
    errorsCh := client.RemoveObjects(ctx, bucketName, objectsToDelete, minio.RemoveObjectsOptions{})

    for err := range errorsCh {
        if err.Err != nil {
            log.Printf("Failed to remove %s: %v\n", err.ObjectName, err.Err)
        } else {
            log.Printf("Successfully removed %s\n", err.ObjectName)
        }
    }
}