package main

import (
	"context"
	"log"
	"net/http"

	"flag"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/jszwec/s3fs/v2"
)

func main() {
	bucket := flag.String("bucket", "your-bucket-name", "Bucket name")
	profile := flag.String("profile", "default", "Profile that used for authentification")
	region := flag.String("region", "us-east-1", "Region that used for authentification")

	flag.Parse()

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithSharedConfigProfile(*profile),
		config.WithRegion(*region),
	)

	if err != nil {
		log.Fatalln(err)
	}

	client := s3.NewFromConfig(cfg)
	s3fs := s3fs.New(client, *bucket)

	http.Handle("/", http.FileServer(http.FS(s3fs)))

	port := ":8080"
	log.Fatal(http.ListenAndServe(port, nil))
}
