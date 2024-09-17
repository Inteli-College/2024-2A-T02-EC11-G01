#!/bin/sh

. /etc/environment

echo "${AWS_ACCESS_KEY_ID}:${AWS_SECRET_ACCESS_KEY}" > /etc/passwd-s3fs
chmod 600 /etc/passwd-s3fs

s3fs $AWS_BUCKET_NAME $MOUNTED_DIR \
    -o url=https://s3.$AWS_REGION.amazonaws.com \
    -o use_path_request_style \
    -o allow_other \
    -o passwd_file=/etc/passwd-s3fs

tail -f /dev/null
