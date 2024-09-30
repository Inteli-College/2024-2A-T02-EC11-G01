#![allow(clippy::result_large_err)]

use std::error::Error;
use aws_sdk_s3::Client;
use uuid::Uuid;
use std::fs;
use std::time::Instant;

pub async fn file_to_upload(image_folder: &str) -> Result<(), Box<dyn Error>> {
    let image_path = image_folder;

    if fs::metadata(&image_folder).is_ok() {
        println!("File exists at: {}", image_folder);
    } else {
        eprintln!("File does not exist at: {}", image_folder);
        return Err(Box::from("File not found"));
    }

    // Load the AWS config from environment variables
    // Tip: Use the aws cli and export the credentials to the environment
    // I may change the from_env() function, since it is deprecated
    let shared_config = aws_config::from_env()
        .load()
        .await;
    let client = Client::new(&shared_config);

    let bucket_name = std::env::var("BUCKET_NAME")?;
    let file_name = image_path.to_string();
    let key = Uuid::new_v4().to_string();

    // Start measuring the upload time
    let start_time = Instant::now();

    // Try to upload the file and track success or failure
    match upload_object(&client, &bucket_name, &file_name, &key).await {
        Ok(_) => {
            let elapsed_time = start_time.elapsed();
            println!("Upload successful in {:.2?} seconds.", elapsed_time);
        }
        Err(e) => {
            let elapsed_time = start_time.elapsed();
            eprintln!("Upload failed after {:.2?} seconds. Error: {:?}", elapsed_time, e);
        }
    };

    Ok(())
}

pub async fn upload_object(
    client: &aws_sdk_s3::Client,
    bucket_name: &str,
    file_name: &str,
    key: &str,
) -> Result<aws_sdk_s3::operation::put_object::PutObjectOutput, Box<dyn std::error::Error>> {
    // Load the file as a byte stream
    let body = aws_sdk_s3::primitives::ByteStream::from_path(std::path::Path::new(file_name)).await?;

    // Upload the object to S3
    client
        .put_object()
        .bucket(bucket_name)
        .key(key)
        .body(body)
        .send()
        .await
        .map_err(|err| err.into())
}
