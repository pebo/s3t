package main

import (
    "fmt"
    "os"
    "net/http"
    "io/ioutil"
    "errors"
    "github.com/codegangsta/cli"
)

func findRegion() (string, error) {
    resp, err := http.Get("http://169.254.169.254/latest/meta-data/placement/availability-zone/")
    if err != nil {
        return "", errors.New("Error reading region")
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return "", errors.New("Error reading region")
    }
    bodyStr := string(body)
    return bodyStr[:len(bodyStr)-1], nil
}

func download(s3svc *S3Service, args []string) {
    from := args[0]
    bucket, key, err := ParseS3Spec(from)
    if err != nil {
        panic(err)
    }
    
    var to string
    if len(args) == 2 {
        to = args[1]
    } else {
        to = GetFileName(key)
    }
    s3svc.Get(&bucket, &key, &to)
}


func getRegion(c *cli.Context) string {
    region := c.GlobalString("region")
    if region == "" {
        awsRegion, err := findRegion()
        if err != nil {
            panic(err)
        }
        region = awsRegion
    }
    return region
}

func NewLsCommand() cli.Command {
    return cli.Command {
        Name:  "ls",
        Usage: "list <bucket>",
        Action: func(c *cli.Context) {
            if len(c.Args()) == 0 {
                cli.ShowAppHelp(c)
                os.Exit(1)
            }
            region := getRegion(c)                
            debug := c.GlobalBool("debug")
            s3svc := NewS3Service(region, debug)
            bucket := c.Args().First()
            s3svc.List(&bucket)
        },
    }
}

func NewDlCommand() cli.Command {
    return cli.Command {
        Name:      "dl",
        Usage:     "<from> [<to>]\n   from/to: <s3://bucket/key>",
        Action: func(c *cli.Context) {
            fmt.Println("dl")
            fmt.Println(c.GlobalString("region"))
                
            region := getRegion(c)                
            debug := c.GlobalBool("debug")
            s3svc := NewS3Service(region, debug)
            fmt.Println(c.Args())
            if len(c.Args()) == 0 {
                cli.ShowAppHelp(c)
                os.Exit(1)
            }
                 
            download(s3svc, c.Args())
        },
    }
}


func main() {
    app := cli.NewApp()
    app.Name = "s3t"
    app.Usage = "Usage: %v [options] <command>\n"
    app.Action = func(c *cli.Context) {
        println("s3t -h for usage")
    }

    // global level flags
    app.Flags = []cli.Flag{
        cli.BoolFlag{
            Name:  "debug",
            Usage: "Show more output",
        },
        cli.StringFlag{
            Name:  "region",
            Usage: "aws region",
        },
    }
    
    // Commands
    app.Commands = []cli.Command{
        NewLsCommand(),
        NewDlCommand(),
    }

    app.Run(os.Args)    
}
