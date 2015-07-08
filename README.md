# s3t
Amazon S3 Tool - a basic tool for performing operations on Amazon S3.

## Features

- download files from s3
- list files in bucket

Examples:

    s3t --region us-west-2 dl s3://bucket/path/file destpath/destfile  
 
# run
Running inside ec2: 

    ./s3t

Running from machine outside ec2:

    go run main.go --region us-west-2 --bucket my-bucket --key my-key
 
# dev

    git clone git@github.com:pebo/s3t.git $GOPATH/src/github.com/pebo/s3t
    cd $GOPATH/src/github.com/pebo/s3t
    
    go install
    godep save
    
    
Building a linux amd64 statically linked executable from mac osx

    boot2docker up
    export ...
    ./build.sh


