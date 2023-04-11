package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

func try() {
	ctx := context.TODO()
	// do we really need repeat all the below steps again and again
	cfg, err := config.LoadDefaultConfig(ctx, config.WithCredentialsProvider(
		credentials.NewStaticCredentialsProvider("AKIATZHQTJTIWLER3L6S", "VO21ptAfMZiNe+/7pWE3t1cBbgEctE7RkSNQkHf9", ""),
	),
		config.WithRegion("us-west-2"),
	)
	if err != nil {
		fmt.Println(err)
	}

	svc := sts.NewFromConfig(cfg)
	input := &sts.GetCallerIdentityInput{}

	output, err := svc.GetCallerIdentity(ctx, input)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(output.Account)
}