# Using MinIO with golang


### Chapter 1 : Initialization

Setting up SDK and evironment variables with handling context
firstly import `github.com/minio/minio-go/v7` and define a struct to hold configuration (Endpoint, AccessKey, SecretKey, UseSSL)

Initialize a new minio.Client using minio.New and implement a simple "Health Check" function that pings the MinIO server to ensure the connection is successful.

### Chapter 2 : Bucket Management 
- Make Bucket: Write a function to create a new bucket named my-user-bucket. Handle the specific error case where the bucket already exists.
- Bucket Exists: Write a check using BucketExists to verify if a bucket is present before attempting operations.
- List Buckets: Iterate through the list of all buckets on the server and print their creation dates and names.

### Chapter 3 : Object Upload (PutObject)

- Use FPutObject to upload a local file (e.g., image.jpg) directly to the bucket. Explicitly set the ContentType (e.g., image/jpeg). This is *Simple Upload* 

- Use PutObject to upload data from an in-memory buffer (bytes.Buffer) or a stream which makes *Stream Upload*. A good way of handling **io.Reader**

### Chapter 4: Object Download (GetObject)

- Download to File: Use FGetObject to download an object from the bucket and save it to the local filesystem.
- Stream to Memory: Use GetObject to retrieve an object as an io.Reader. Read the first 10 bytes of the file to verify the file type (magic numbers) without downloading the whole file.
- Handle "Object Not Found" errors specifically.

### Chapter 5: Listing Objects

- Basic List: Use ListObjects to list all items in a specific bucket.
- Filtered List: Use ListObjectsOptions to list only objects with a specific prefix (e.g., images/) and set the `Recursive` flag to false to simulate folder structures.
- Iterate through the object channel ch := minioClient.ListObjects(...) and print the size and last modified date of each object.

### Chapter 6: Metadata & Copying

- Copy Object: Use CopyObject to duplicate a file within the same bucket (e.g., source.txt -> backup.txt).
- Custom Metadata: During the copy operation, add custom user metadata (e.g., X-Amz-Meta-Author: GoStudent).
- Set Tags: Implement a function to apply object tags (Key-Value pairs) to an uploaded file.

### Chapter 7: Presigned URLs

- Shareable Link (GET): Generate a Presigned URL for downloading an object that expires in 24 hours using PresignedGetObject.
- Upload Link (PUT): Generate a Presigned URL for uploading an object using PresignedPutObject. This allows someone to upload to your bucket without having your Secret Key, only via this specific URL.
- Print the URLs to the console.


### Chapter 8: Deletion Operations

- Remove Single: Use RemoveObject to delete a specific file. Handle cases where the object does not exist.
- Remove Multiple: Use RemoveObjects to delete a list of objects in a single API call. Iterate through the error channel to see if any specific deletions failed.
- Clean Up: Use RemoveIncompleteUpload to clear out any "stuck" multipart upload data from failed transfers.

### Chapter 9: Bucket Policies & Public Access

- Set Policy: Write a function that sets a "ReadOnly" policy on a specific bucket using SetBucketPolicy. This allows anonymous users (public internet) to list and download files from this bucket.
- Get Policy: Retrieve the current policy configuration using GetBucketPolicy and print the JSON structure.
- Remove Policy: Revoke public access by setting the policy back to none.


### Chapter 10: Bucket Notifications (Events)

- Set Notification: Configure a bucket to send a notification (e.g., to a Webhook URL or an ARN) whenever a file is created (s3:ObjectCreated:*).
- Listen to Notifications: Use the ListenBucketNotification method. This opens a long-running connection that prints events to the console whenever a file is uploaded or deleted in real-time.
- Run this in a separate Goroutine to keep the application alive while listening.


