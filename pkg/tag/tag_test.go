package tag_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"

	"github.com/bishy999/go-tag-ec2/pkg/tag"
)

const (
	instanceId = "i-abcdefghx"
)

type mockEC2Client struct {
	ec2iface.EC2API
	Reservations []*ec2.Reservation
	InstanceId   string
}

func (c *mockEC2Client) DescribeInstances(in *ec2.DescribeInstancesInput) (*ec2.DescribeInstancesOutput, error) {

	//mock response functionality
	reservations := []*ec2.Reservation{}
	instances := []*ec2.Instance{}

	instances = append(instances, &ec2.Instance{
		InstanceId: aws.String(instanceId),
	})

	reservations = append(reservations, &ec2.Reservation{Instances: instances})

	return &ec2.DescribeInstancesOutput{
		Reservations: reservations,
	}, nil
}

func (c *mockEC2Client) CreateTags(in *ec2.CreateTagsInput) (*ec2.CreateTagsOutput, error) {

	if aws.StringValue(in.Resources[0]) == instanceId {
		match := mockEC2Client{InstanceId: aws.StringValue(in.Resources[0])}

		fmt.Printf("%T", match)

	}

	return &ec2.CreateTagsOutput{}, nil

}

func TestGetInstances(t *testing.T) {

	tt := []struct {
		name        string
		id          string
		region      string
		environment string
		team        string
		owner       string
		costCenter  string
		result      string
	}{
		{"Test 001 get instance", "i-abcdefgh", "eu-west-1", "EDDIE_DEV", "System_Team", "lassie", "0000", instanceId},
	}

	for _, tc := range tt {
		fmt.Println(tc.name)
		t.Run(tc.name, func(t *testing.T) {

			mockSvc := &mockEC2Client{}

			input := tag.UserInfo{
				Client: tag.Client{mockSvc},
				Tags:   map[string]string{"Name": tc.id, "Region": tc.region, "Environment": tc.environment, "Team": tc.team, "Owner": tc.owner, "Cost-Centre": tc.costCenter},
			}

			res, err := input.GetInstance()
			if err != nil {
				log.Printf("error getting instance: %s", err.Error())
				return
			}
			if res != tc.result {
				t.Errorf("Test %v result should be `%v`, got  `%v`", tc.name, tc.result, err)
			}
		})
	}

}

func ExampleTagData_GetInstance() {

	mockSvc := &mockEC2Client{}

	input := tag.UserInfo{
		Client: tag.Client{mockSvc},
		Tags:   map[string]string{"Name": "dummyvalue", "Region": "eu-west-1", "Environment": "Denvironment", "Team": "Ateam", "Owner": "timmy", "Cost-Centre": "0000"},
	}

	res, err := input.GetInstance()
	if err != nil {
		log.Printf("error getting instance: %s", err.Error())
		return
	}
	fmt.Printf("\n%v", res)

	// Output:
	// i-abcdefghx
}

func TestCreateTags(t *testing.T) {

	tt := []struct {
		name        string
		id          string
		region      string
		environment string
		team        string
		owner       string
		costCenter  string
		result      bool
	}{
		{"Test 001 create tag", "dummyvalue", "eu-west-1", "Denvironment", "Ateam", "lassie", "0000", true},
	}

	for _, tc := range tt {
		fmt.Println(tc.name)
		t.Run(tc.name, func(t *testing.T) {

			mockSvc := &mockEC2Client{}

			input := tag.UserInfo{
				Client: tag.Client{mockSvc},
				Tags:   map[string]string{"Name": tc.id, "Region": tc.region, "Environment": tc.environment, "Team": tc.team, "Owner": tc.owner, "Cost-Centre": tc.costCenter},
			}

			res, err := input.CreateTags(tc.id)
			if err != nil {
				log.Printf("error getting instance: %s", err.Error())
				return
			}
			if res != tc.result {
				t.Errorf("Test %v result should be `%v`, got  `%v`", tc.name, tc.result, err)
			}
		})
	}

}

func ExampleTagData_CreateTags() {

	mockSvc := &mockEC2Client{}

	input := tag.UserInfo{
		Client: tag.Client{mockSvc},
		Tags:   map[string]string{"Name": "dummyvalue", "Region": "eu-west-1", "Environment": "Denvironment", "Team": "Ateam", "Owner": "timmy", "Cost-Centre": "0000"},
	}

	res, err := input.CreateTags(input.Tags["Name"])
	if err != nil {
		log.Printf("error getting instance: %s", err.Error())
		return
	}
	fmt.Printf("\n%v", res)

	// Output:
	//true

}

func BenchmarkTagData_CreateTags(b *testing.B) {

	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		mockSvc := &mockEC2Client{}

		input := tag.UserInfo{
			Client: tag.Client{mockSvc},
			Tags:   map[string]string{"Name": "dummyvalue", "Region": "eu-west-1", "Environment": "Denvironment", "Team": "Ateam", "Owner": "timmy", "Cost-Centre": "0000"},
		}

		res, err := input.CreateTags(input.Tags["Name"])
		if err != nil {
			log.Printf("error getting instance: %s", err.Error())
			return
		}

		fmt.Println(res)

		// Output:
		//true

	}

}
