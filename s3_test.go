package main

import "testing"

func TestIsS3Format(t *testing.T) {
    if !IsS3Format("s3:/hjadsg") {
        t.Error("Expected s3 format")
    }
    if IsS3Format("file.dat") {
        t.Error("Expected not s3 format")
    }
}

func TestGetFileNameWithPath(t *testing.T) {
    fileName := GetFileName("path/key.dat")

    if fileName != "key.dat" {
        t.Error("Expected key.dat got", fileName)
    }
}

func TestGetFileNameWithOutPath(t *testing.T) {
    fileName := GetFileName("key.dat")

    if fileName != "key.dat" {
        t.Error("Expected key.dat got", fileName)
        
    }
}
