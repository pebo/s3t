package main

import (
    "fmt"
    "os"
    "io"
    "regexp"
    "errors"
    "path"
    "strings"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/awsutil"
    "github.com/aws/aws-sdk-go/service/s3"
)

type S3Service struct {
    client *s3.S3
}

func IsS3Format(spec string) bool {
    return strings.HasPrefix(spec, "s3:")
}

func ParseS3Spec(spec string)  (bucket string, key string, err error) {
	re := regexp.MustCompile("s3://([^/]+)/(.+)")
    parts := re.FindStringSubmatch(spec)    
    if len (parts) != 3 {
        return "", "", errors.New("Error reading spec")
    }
    return parts[1], parts[2], nil
}

func GetFileName(key string)  (fileName string) {
    _, file := path.Split(key)
    return file
}

func NewS3Service(region string, debug bool) *S3Service {
    var logLevel uint
    if debug {
        logLevel = 1
    } else {
        logLevel = 0
    }
    config := &aws.Config{Region: region, LogLevel: logLevel}
    client := s3.New(config)
    return &S3Service{client: client}
}

func (s S3Service) List(bucket *string) {
    params := &s3.ListObjectsInput{Bucket: bucket}
    resp, err := s.client.ListObjects(params)
    if err != nil {
        panic(err)
    }
    fmt.Println(awsutil.StringValue(resp))
}

func (s S3Service) Get(bucket *string, key *string, to *string) {
    params := &s3.GetObjectInput { Bucket: bucket, Key: key}
    resp, err := s.client.GetObject(params)
    if err != nil {
        panic(err)
    }
    
    // Pretty-print the response data.
    fmt.Println(awsutil.StringValue(resp))
    out, err := os.Create(*to)
    if err != nil {
        panic(err)
    }
    defer out.Close()
    io.Copy(out, resp.Body)
}
