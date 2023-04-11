package analyze

import (
	"context"
	"fmt"
	"strings"

	"github.com/AvineshTripathi/cred/utils"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

func validate(matchedId [][]string, matchedKey [][]string) (*[]utils.Truth, error) {

	var result []utils.Truth
	if len(matchedKey) > 100 {
		chunkSize := 100
		numChunks := len(matchedKey) / chunkSize
		if len(matchedKey)%chunkSize != 0 {
			numChunks++
		}
		r := make(chan utils.Truth, numChunks)

		for i := 0; i < numChunks; i++ {
			start := i * chunkSize
			end := start + chunkSize
			if end > len(matchedKey) {
				end = len(matchedKey)
			}
			go check(matchedId, matchedKey[start:end], r)
		}

		for i := 0; i < numChunks; i++ {
			ele := <-r
			result = append(result, ele)
		}

	} else {
		resultsSmall, err := checkSmallArr(matchedId, matchedKey)
		if err != nil {
			fmt.Println(err)
		}

		result = append(result, *resultsSmall)
	}

	return &result, nil
}

func check(matchedId [][]string, matchedKey [][]string, result chan utils.Truth) {

	region := "us-west-2"
	res := utils.Truth{}

	for _, id := range matchedId {
		if len(id) != 2 {
			continue
		}

		tid := strings.TrimSpace(id[1])

		for _, key := range matchedKey {
			if len(key) != 2 {
				continue
			}

			tkey := strings.TrimSpace(key[1])
			ctx := context.TODO()

			// do we really need repeat all the below steps again and again
			cfg, err := config.LoadDefaultConfig(ctx, config.WithCredentialsProvider(
				credentials.NewStaticCredentialsProvider(tid, tkey, ""),
			),
				config.WithRegion(region),
			)
			if err != nil {
				fmt.Println(err, "first")
			}

			svc := sts.NewFromConfig(cfg)
			input := &sts.GetCallerIdentityInput{}

			output, err := svc.GetCallerIdentity(ctx, input)
			if err != nil {
				continue
			}

			res.Id = tid
			res.Key = tkey
			res.IdPath = id[0]
			res.KeyPath = key[0]

			if output.Account != nil {
				res.Found = true

			}

		}
	}
	result <- res
}

func checkSmallArr(matchedId [][]string, matchedKey [][]string) (*utils.Truth, error) {

	region := "us-west-2"
	res := &utils.Truth{}

	for _, id := range matchedId {
		if len(id) != 2 {
			continue
		}

		tid := strings.TrimSpace(id[1])

		for _, key := range matchedKey {
			if len(key) != 2 {
				continue
			}

			tkey := strings.TrimSpace(key[1])
			ctx := context.TODO()
			// do we really need repeat all the below steps again and again
			cfg, err := config.LoadDefaultConfig(ctx, config.WithCredentialsProvider(
				credentials.NewStaticCredentialsProvider(tid, tkey, ""),
			),
				config.WithRegion(region),
			)
			if err != nil {
				fmt.Println(err, "first")
			}

			svc := sts.NewFromConfig(cfg)
			input := &sts.GetCallerIdentityInput{}

			output, err := svc.GetCallerIdentity(ctx, input)
			if err != nil {
				continue
			}

			res.Id = tid
			res.Key = tkey

			res.IdPath = id[0]
			res.KeyPath = key[0]

			if output.Account != nil {
				res.Found = true
			}

		}
	}

	return res, nil
}
