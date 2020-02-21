package tag

import (
	"errors"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
)

// Client for accessing ec2 api
type Client struct {
	Svc ec2iface.EC2API
}

// NewClient Creates a new client for ec2 interaction
func NewClient(region *string) (*Client, error) {

	sess, err := session.NewSession(&aws.Config{Region: aws.String(*region)})
	if err != nil {
		return nil, err
	}
	return &Client{
		Svc: ec2.New(sess),
	}, nil

}

// GetInstance Get instance id matching name provided
func (t *UserInfo) GetInstance() (string, error) {

	var count int
	var id string

	resp, err := t.Client.Svc.DescribeInstances(&ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			&ec2.Filter{
				Name:   aws.String("tag:Name"),
				Values: []*string{aws.String(t.Tags["Name"])},
			},
			&ec2.Filter{
				Name:   aws.String("instance-state-name"),
				Values: []*string{aws.String("running")},
			},
		},
	})

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				log.Fatalf(aerr.Error())
			}
		} else {
			log.Fatalf(err.Error())
		}
	}

	for _, r := range resp.Reservations {

		if len(r.Instances) != 1 {
			log.Fatalf("We have found [%d] instance ids. Expecting exactly 1 ⚠️", len(r.Instances))
		} else {
			instance := r.Instances[0].InstanceId
			id = *instance
			log.Printf("Instance id: %q", id)
			count++
		}

	}

	if count == 0 {
		msg := "\"No instance id was found for name provided: " + t.Tags["Name"] + "\""
		return "", errors.New(msg)
	}

	return id, nil

}

//CreateTags Reports whether tags have been successfully added to the instance
func (t *UserInfo) CreateTags(name string) (bool, error) {

	var tags []*ec2.Tag

	for k, v := range t.Tags {
		tags = append(tags, &ec2.Tag{
			Key:   aws.String(k),
			Value: aws.String(v),
		})
	}

	input := &ec2.CreateTagsInput{
		Resources: []*string{
			aws.String(name),
		},
		Tags: tags,
	}

	_, err := t.Client.Svc.CreateTags(input)

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				log.Fatalf(aerr.Error())
			}
		} else {
			log.Fatalf(err.Error())
		}
	}

	return true, nil
}
