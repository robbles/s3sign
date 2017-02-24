# s3sign

```
usage: s3sign [--region REGION] [--expiry EXPIRY] BUCKET KEY

positional arguments:
  bucket                 the S3 bucket name
  key                    the S3 key

options:
  --region REGION        the AWS region [default: us-west-1]
  --expiry EXPIRY        how long the URL should stay valid for [default: 24h]
  --help, -h             display this help and exit
```

## Installation

```
go get -u github.com/robbles/s3sign
```

## Sample usage

```
$ export AWS_ACCESS_KEY_ID=XXX AWS_SECRET_ACCESS_KEY=YYY

$ s3sign -expiry 1h -region us-east-1 my-bucket /path/to/key
https://my-bucket.s3.amazonaws.com/path/to/key?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=XXX%2F20170224%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20170224T005818Z&X-Amz-Expires=3600&X-Amz-SignedHeaders=host&X-Amz-Signature=YYY
```
