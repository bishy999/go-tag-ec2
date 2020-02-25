package tag

import (
	"errors"
	"log"
	"regexp"
)

const (
	usage = `
	######################################################################################################################################################
	#                                                                                                                                                    #
	# Enter the required tagging information                                                                                                             #
	#                                                                                                                                                    #
	# Usage:                                                                                                                                             #
	#     ec2-tags-client -environment=DEV -name=mytestenv -team=ATeam -owner=jimmy -region=eu-west-1 -costCentre=00000                                  #
    #                                                                                                                                                    #
	#                                                                                                                                                    #
	######################################################################################################################################################
	`

	err1 = `input needs to be alphanumeric can include an underscore and has to be less than 15 characters`
	err2 = `all fields are required and can not be blank`

	matcher = "^[a-zA-Z0-9_.-]{1,30}$"
)

// UserInfo struct contains info provided by the user
type UserInfo struct {
	Client Client
	Tags   map[string]string
}

// ContainsValidTags whether sufficient input has been provided to be able to tag instance
func (t *UserInfo) ContainsValidTags() error {

	r, _ := regexp.Compile(matcher)

	log.Printf("######## Tags ######## \n")

	for k, v := range t.Tags {
		if v == "" {
			log.Print(usage)
			log.Printf("Error caused by key \"%s\" with value \"%s\"\n", k, v)
			return errors.New(err2)
		}

		match := r.FindString(v)
		if match == "" {
			log.Printf("Error caused by key \"%s\" with value \"%s\"\n", k, v)
			return errors.New(err1)
		}
		log.Printf("# %s: %s\n", k, v)

	}
	log.Printf("########################## \n")

	return nil

}
