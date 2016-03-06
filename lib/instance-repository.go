package flyingSquid

import (
	"fmt"
	"regexp"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type ec2instanceRepository struct {
	instances []map[string]string
}

func (repository *ec2instanceRepository) Fetch(matcher string) []map[string]string {

	if repository.instances == nil {
		repository.fetch()
	} else {
		go repository.fetch()
	}

	if matcher == "" {
		return repository.instances
	}

	var instances []map[string]string

	for _, instance := range repository.instances {
		regex, _ := regexp.Compile(".*" + matcher + ".*")
		if matched := regex.MatchString(instance["name"]); matched {
			instances = append(instances, instance)
		}
	}

	return instances

}

func (repository *ec2instanceRepository) fetch() {

	ec2Service := ec2.New(session.New(), &aws.Config{
		Region: aws.String("ap-southeast-2"),
		)})

	resp, err := ec2Service.DescribeInstances(nil)

	if err != nil {
		fmt.Println("An error occured while fetching instances from AWS")
		fmt.Println(err)
		return
	}

	var instances []map[string]string
	for _, reservation := range resp.Reservations {
		for _, instance := range reservation.Instances {
			launchTime := *instance.LaunchTime
			instances = append(instances, map[string]string{
				"imageId":      *instance.ImageId,
				"instanceId":   *instance.InstanceId,
				"instanceType": *instance.InstanceType,
				"launchTime":   launchTime.Format(time.RFC1123),
				"environment":  repository.find("Environment", instance.Tags),
				"solution":     repository.find("Application", instance.Tags),
				"slice":        repository.find("Slice", instance.Tags)})
		}
	}

	repository.instances = instances

}

func (repository *ec2instanceRepository) find(target string, tags []*ec2.Tag) string {
	for _, tag := range tags {
		if *tag.Key == target {
			return *tag.Value
		}
	}
	return ""
}
