package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
)

func main() {
	stat, _ := os.Stdin.Stat()

	if (stat.Mode() & os.ModeCharDevice) != 0 {
		fmt.Fprintln(os.Stderr, "No tokens detected. Hint: cat tokens.txt | aws-test")
		os.Exit(1)
	}

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()

		var count = 0
		var accessKey = ""
		var secretKey = ""

		for _, x := range strings.Fields(line) {
			if count == 0 {
				accessKey = x
			}

			if count == 1 {
				secretKey = x
			}

			count++
		}

		err := testAWSCredentials(accessKey, secretKey)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
	}
}

func testAWSCredentials(accessKey string, secretKey string) error {

	sess, err := session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
	})

	if err != nil {
		return err
	}

	svc := sts.New(sess)
	input := &sts.GetCallerIdentityInput{}

	result, err := svc.GetCallerIdentity(input)

	if err != nil {
		return err
	}

	fmt.Printf("The access key %v has a valid secret key (%v)\n", accessKey, secretKey)
	fmt.Println(result)

	return nil
}
