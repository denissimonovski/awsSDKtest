package main

import (
	"bufio"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"os"
	"strings"
)

func main() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	ec2Svc := ec2.New(sess)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Choose:\n1) Start VM\n2) Stop VM\n3) Describe VM\n")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	switch text {
	case "1":
		input := &ec2.StartInstancesInput{
			InstanceIds: []*string{
				aws.String("i-083ba1a258935288c"),
			},
		}
		result, err := ec2Svc.StartInstances(input)
		if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
				switch aerr.Code() {
				default:
					fmt.Println(aerr.Error())
				}
			} else {
				// Print the error, cast err to awserr.Error to get the Code and
				// Message from an error.
				fmt.Println(err.Error())
			}
			return
		}

		fmt.Println(result)
	case "2":
		input := &ec2.StopInstancesInput{
			InstanceIds: []*string{
				aws.String("i-083ba1a258935288c"),
			},
		}
		result, err := ec2Svc.StopInstances(input)
		if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
				switch aerr.Code() {
				default:
					fmt.Println(aerr.Error())
				}
			} else {
				// Print the error, cast err to awserr.Error to get the Code and
				// Message from an error.
				fmt.Println(err.Error())
			}
			return
		}

		fmt.Println(result)
	case "3":
		input := &ec2.DescribeInstancesInput{
			InstanceIds: []*string{
				aws.String("i-083ba1a258935288c"),
			},
		}
		result, err := ec2Svc.DescribeInstances(input)
		if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
				switch aerr.Code() {
				default:
					fmt.Println(aerr.Error())
				}
			} else {
				// Print the error, cast err to awserr.Error to get the Code and
				// Message from an error.
				fmt.Println(err.Error())
			}
			return
		}

		fmt.Println(result)
	}
}
