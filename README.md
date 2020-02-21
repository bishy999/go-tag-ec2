
# go-tag-ec2

go-tag-ec2 is a Go client library for accessing AWS EC2 API.

You can view the AWS API docs here: [https://docs.aws.amazon.com/sdk-for-go/api/service/ec2/](https://docs.aws.amazon.com/sdk-for-go/api/service/ec2/)

You can view the client API docs by serving the docs from this repository : [http://localhost:6060/pkg/](http://localhost:6060/pkg/)
```go
 godoc -http :6060
```


## Status
[![Build Status](https://travis-ci.com/bishy999/go-tag-ec2.svg?branch=master)](https://travis-ci.com/bishy999/go-tag-ec2)
[![Go Report Card](https://goreportcard.com/badge/github.com/bishy999/go-tag-ec2)](https://goreportcard.com/report/github.com/bishy999/go-tag-ec2)



## Usage (package)

### Download package
```go
 go get github.com/bishy999/go-tag-ec2
 ```

### Use package
```go

import 
(
	 "github.com/bishy999/go-tag-ec2/tag"

)
```

### Authentication
You will need AWS credentials to access the AWS API

The credentials by default can be set in ~/.aws/credentials on Linux, macOS, or Unix

```go
[default]
aws_access_key_id = your_access_key_id
aws_secret_access_key = your_secret_access_key
```

You can then use these credentials to create a new client. An example of a client is stored under the cmd directory in this repository

```go


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




```

## Usage (binary)

Download the client binary from the repository and compile it with version 

Go get will download from the master, as such when we download it give it the tag verison from the master

```go
go get -v -race -ldflags "-X main.version=v1.0.0 -X main.buildstamp=`TZ=UTC date -u '+%Y-%m-%dT%H:%M:%SZ'`)" github.com/bishy999/go-tag-ec2cmd/ec2-tags-client

ec2-tags-client -environment=DEV -name=test1928.aws.xcl.ie -team=ATeam -owner=jimmy -region=eu-west-1 -costCentre=00000

```


## Contributing

We love pull requests! Please see the [contribution guidelines](CONTRIBUTING.md).
