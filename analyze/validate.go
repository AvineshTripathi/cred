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

func Validate(matchedId [][]string, matchedKey [][]string) (*[]utils.Truth, error) {

	var result []utils.Truth
	if len(matchedKey) > 100 {
		chunkSize := 100
		numChunks := len(matchedKey) / chunkSize
		if len(matchedKey)%chunkSize != 0 {
			numChunks++
		}
		r := make(chan []utils.Truth, numChunks)

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
			result = append(result, ele...)
		}

	} else {
		resultsSmall, err := checkSmallArr(matchedId, matchedKey)
		if err != nil {
			fmt.Println(err)
		}	
		fmt.Println(*resultsSmall)
		result = append(result, *resultsSmall...)
	}

	return &result, nil
}

// Danger Area! Under Constructions
func check(matchedId [][]string, matchedKey [][]string, result chan []utils.Truth) {

	region := "us-west-2"
	res := []utils.Truth{}
	var r utils.Truth
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

			r.Id = tid
			r.Key = tkey
			r.IdPath = id[0]
			r.KeyPath = key[0]

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
				res = append(res, r)
				continue
			}

			

			if output.Account != nil {
				r.Found = true

			}

		}
	}
	result <- res
}

func checkSmallArr(matchedId [][]string, matchedKey [][]string) (*[]utils.Truth, error) {

	region := "us-west-2"
	var res []utils.Truth
	var r utils.Truth
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
			
			r.Id = tid
			r.Key = tkey

			r.IdPath = id[0]
			r.KeyPath = key[0]
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
				res = append(res, r)
				continue
			}

			// extra check 
			if output.Account != nil {
				r.Found = true
			}

			res = append(res, r)
		}
	}

	return &res, nil
}
