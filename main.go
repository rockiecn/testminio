package main

import (
	"fmt"
	"log"
	"time"

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

	//  创建bucket
	bucketName := "mymusic6"

	fmt.Println("creating bucket..")
	location := "us-east-1"
	err = minioClient.MakeBucket(bucketName, location)
	if err != nil {
		exists, err := minioClient.BucketExists(bucketName)
		if err == nil && exists {
			log.Printf("We already own %s\n", bucketName)
			return
		} else {
			log.Fatalln(err)
		}
	}
	log.Printf("Successfully created %s\n", bucketName)

	fmt.Println("sleep 30 sec for bucket confirm")
	time.Sleep(30 * time.Second)
	// 上传一个文件。
	objectName := "Recovery.txt"
	filePath := "c:/Recovery.txt"

	fmt.Println("put object..")
	n, err := minioClient.FPutObject(bucketName, objectName, filePath, minio.PutObjectOptions{})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Successfully uploaded %s of size %d\n", objectName, n)

	// 显示bucket列表
	fmt.Println("list buckets..")
	buckets, err := minioClient.ListBuckets()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, bucket := range buckets {
		fmt.Println(bucket)
	}
}
