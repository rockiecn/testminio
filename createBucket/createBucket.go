package main

import (
	"fmt"
	"log"

	"github.com/minio/minio-go/v6"
)

func main() {
	fmt.Println("connecting gateway..")
	endpoint := "127.0.0.1:5080"
	accessKeyID := "0x9c784fD443Faf95D25F1598F7D2ece581CeA2DEe"
	secretAccessKey := "memoriae"
	//useSSL := true

	fmt.Println("initial minio client..")
	// 初使化minio client对象。
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, false)
	if err != nil {
		log.Fatalln(err)
	}

	// create bucket
	err = minioClient.MakeBucket("06070221", "us-east-1")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Successfully created mybucket.")

}
