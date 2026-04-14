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

    // defining a readOnly Policy JSON
    // This allows anonymous public read access to everything in the bucket
    policy := `{
        "Version": "2012-10-17",
        "Statement": [
            {
                "Effect": "Allow",
                "Principal": {"AWS": ["*"]},
                "Action": ["s3:GetObject"],
                "Resource": ["arn:aws:s3:::` + bucketName + `/*"]
            }
        ]
    }`

    // set Policy
    err = client.SetBucketPolicy(ctx, bucketName, policy)
    if err != nil {
        log.Fatalln(err)
    }
    log.Println("Bucket policy set to ReadOnly.")

    // get Policy
    currentPolicy, err := client.GetBucketPolicy(ctx, bucketName)
    if err != nil {
        log.Fatalln(err)
    }
    log.Printf("Current Policy: %s\n", currentPolicy)

    // remove Policy (Set to empty to revoke)
    // err = client.SetBucketPolicy(ctx, bucketName, "")
    // if err != nil { log.Fatalln(err) }
    // log.Println("Policy removed.")
}