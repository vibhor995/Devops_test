package main

import (
    "fmt"
    "os"
    "time"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/ec2"
)



func TestWriteListInstance(t *testing.T) {
       thisTime := time.Now()
       nowString := thisTime.Format("2006-01-02 15:04:05 Monday")
       t.Log("Starting unit test at " + nowString)

       api := &aws.Config 
        
       awsConnect, err := session.NewSession(api{
           Region: aws.String("us-east-2")},
       )
       if err != nil {
        fmt.Println(err)
       }

       ec2sess := ec2.New(awsConnect)
       instanceInfo := &ec2.DescribeInstancesInput{
           InstanceIds: []*string{aws.String("i-007170d8ee4f99dee")},
       }

       result, err :=  ec2sess.DescribeInstances(instanceInfo)
       if err != nil {
		t.Log("Wrong instance ID or Got an error retrieving information about your Amazon EC2 instances:")
		t.Log(err)
		return
	}

          
}   
