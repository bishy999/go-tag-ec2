package tag_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/bishy999/go-tag-ec2/pkg/tag"
)

func TestValidInput(t *testing.T) {

	tt := []struct {
		name      string
		flagValue string
		result    string
	}{
		{"Test 001 correct input", "test01", ""},
		{"Test 002 blank field", "", "all fields are required and can not be blank"},
		{"Test 003 bad symbols", "%^&$%", "input needs to be alphanumeric can include an underscore and has to be less than 15 characters"},
		{"Test 004 too many characters", "qwertyuiopasdfghjklzxcvbnm12345", "input needs to be alphanumeric can include an underscore and has to be less than 15 characters"},
	}

	for _, tc := range tt {
		fmt.Println(tc.name)
		t.Run(tc.name, func(t *testing.T) {
			input := tag.UserInfo{
				Tags: map[string]string{"Name": tc.flagValue, "Region": tc.flagValue, "Environment": tc.flagValue, "Team": tc.flagValue, "Owner": tc.flagValue},
			}

			err := input.ContainsValidTags()
			if err != nil {
				if err.Error() != tc.result {
					t.Errorf("Test %v result should be `%v`, got  `%v`", tc.name, tc.result, err)
				}
			}

		})
	}

}

func ExampleTagData_ContainsValidTags() {

	input := tag.UserInfo{
		Tags: map[string]string{"Name": "test01", "Region": "eu-west-1", "Environment": "EDDIE_DEV", "Team": "System_Team", "Owner": "jimmy"},
	}

	err := input.ContainsValidTags()
	if err != nil {
		log.Fatal(err)
	}

	// Output:
	//
}
