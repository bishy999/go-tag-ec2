package main

import (
	"flag"
	"log"

	"github.com/bishy999/go-tag-ec2/pkg/tag"
)

var (
	environmentPtr = flag.String("environment", "", "Type of instance created")
	namePtr        = flag.String("name", "", "The name of instance")
	teamPtr        = flag.String("team", "", "The team the instance belongs to")
	ownerPtr       = flag.String("owner", "", "User who created and is responsible for the instance")
	regioinPtr     = flag.String("region", "", "The region in which the instance resides")
	costCentrePtr  = flag.String("costCentre", "", "The costCentre this instance belongs to")
)

var (
	version    string
	buildstamp string
)

func main() {

	log.Printf("Version    : %s\n", version)
	log.Printf("Build Time : %s\n", buildstamp)

	flag.Parse()

	client, err := tag.NewClient(regioinPtr)
	if err != nil {
		log.Printf("error creating new client %v", err.Error())
		return
	}

	input := tag.UserInfo{
		Client: *client,
		Tags:   map[string]string{"Name": *namePtr, "Region": *regioinPtr, "Environment": *environmentPtr, "Team": *teamPtr, "Owner": *ownerPtr, "Cost-Centre": *costCentrePtr},
	}

	err = input.ContainsValidTags()
	if err != nil {
		log.Printf("Error: %s", err.Error())
		return
	}

	instances, err := input.GetInstance()
	if err != nil {
		log.Printf("error getting instance: %s", err.Error())
		return
	}

	ok, err := input.CreateTags(instances)
	if err != nil {
		log.Printf("error getting instances: %s", err.Error())
		return
	}

	if ok {
		log.Printf("Successfully tagged instance id: %q", instances)
	}

}
