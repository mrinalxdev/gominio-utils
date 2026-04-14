package main

import (
    "context"
    "log"

    "github.com/minio/minio-go/v7"
    "github.com/minio/minio-go/v7/pkg/credentials"
    "github.com/minio/minio-go/v7/pkg/notification"
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

    events := []string{
        string(notification.ObjectCreatedAll),   // "s3:ObjectCreated:*"
        string(notification.ObjectRemovedAll),   // "s3:ObjectRemoved:*"
        string(notification.ObjectAccessedAll),  // "s3:ObjectAccessed:*"
    }

    // optional : filter by prefix/suffix (also from notification package)
    prefix := ""        // e.g., "uploads/"
    suffix := ".txt"    // e.g., only listen for .txt files
    // you can also use notification.FilterRule{Key: "prefix", Value: "uploads/"} for easy configs

    log.Printf("listening for events on bucket '%s'...\n", bucketName)

    // listen for bucket notifications (long-running)
    eventInfoCh := client.ListenBucketNotification(ctx, bucketName, prefix, suffix, events)

    for notificationInfo := range eventInfoCh {
        if notificationInfo.Err != nil {
            log.Printf("notification error: %v\n", notificationInfo.Err)
            continue
        }

        for _, record := range notificationInfo.Records {
            switch notification.EventType(record.EventName) {
            case notification.ObjectCreatedPut:
                log.Printf("[PUT]     %s (size: %d bytes)\n", 
                    record.S3.Object.Key, record.S3.Object.Size)
            case notification.ObjectCreatedCopy:
                log.Printf("[COPY]    %s\n", record.S3.Object.Key)
            case notification.ObjectRemovedDelete:
                log.Printf("[DELETE]  %s\n", record.S3.Object.Key)
            case notification.ObjectAccessedGet:
                log.Printf("[GET]     %s\n", record.S3.Object.Key)
            default:
                log.Printf("[%s] %s\n", record.EventName, record.S3.Object.Key)
            }
        }
    }
}