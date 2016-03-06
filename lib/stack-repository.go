package flyingSquid

import (
	"fmt"
	"regexp"
	"time"

	"github.com/aws/aws-sdk-go/service/cloudformation"
)

type stackRepository struct {
	stacks []map[string]string
}

func (repository *stackRepository) Fetch(matcher string) []map[string]string {

	if repository.stacks == nil {
		repository.fetch()
	} else {
		go repository.fetch()
	}

	if matcher == "" {
		fmt.Println("no matcher, sending back all")
		return repository.stacks
	}

	var stacks []map[string]string

	for _, stack := range repository.stacks {
		regex, _ := regexp.Compile(".*" + matcher + ".*")
		if matched := regex.MatchString(stack["stackId"]); matched {
			stacks = append(stacks, stack)
		}
	}

	return stacks

}

func (repository *stackRepository) Resources(matcher string) []map[string]string {
	return []map[string]string{
		{"name": "MyLoadbalancer", "type": "AWS::ElasticLoadBalancing::LoadBalancer", "id": "0000"},
		{"name": "MyECSService", "type": "AWS::ECS::Service", "id": "1111"},
		{"name": "MyTaskDef", "type": "AWS::ECS::TaskDefinition", "id": "2222"},
		{"name": "MyCluster", "type": "AWS::ECS:Cluster", "id": "3333"},
		{"name": "MyAutoScalingGroup", "type": "AWS::AutoScaling::AutoScalingGroup", "id": "4444"},
		{"name": "MYLaunchConfig", "type": "AWS::AutoScaling::LaunchConfiguration", "id": "5555"},
	}
}

func (repository *stackRepository) fetch() {

	// service := cloudformation.New(session.New(), &aws.Config{
	// 	Region: aws.String("ap-southeast-2"),
	// 	Credentials: credentials.NewStaticCredentials(
	// 	)})
	//
	// resp, err := service.DescribeStacks(nil)
	//
	// if err != nil {
	// 	fmt.Println("An error occured while fetching stacks from AWS")
	// 	fmt.Println(err)
	// 	return
	// }
	//
	// var stacks []map[string]string
	// for _, stack := range resp.Stacks {
	// 	creationTime := *stack.CreationTime
	// 	stacks = append(stacks, map[string]string{
	// 		"status":       *stack.StackStatus,
	// 		"stackName":    *stack.StackName,
	// 		"stackId":      *stack.StackId,
	// 		"creationTime": creationTime.Format(time.RFC1123),
	// 		"environment":  repository.find("Environment", stack.Tags),
	// 		"solution":     repository.find("Application", stack.Tags),
	// 		"slice":        repository.find("Slice", stack.Tags)})
	// }
	//
	// repository.stacks = stacks
	repository.stacks = []map[string]string{
		{
			"stackId":        "000000",
			"stackName":      "public-sales-api",
			"creationTime":   time.Now().Format(time.RFC1123),
			"environment":    "kaos",
			"slice":          "master",
			"lineOfBusiness": "ARHI",
			"status":         cloudformation.StackStatusCreateComplete,
		},
		{
			"stackId":        "111111",
			"stackName":      "points-plus-pay",
			"creationTime":   time.Now().Format(time.RFC1123),
			"environment":    "kaos",
			"slice":          "ethan",
			"lineOfBusiness": "ARHI",
			"status":         cloudformation.StackStatusUpdateRollbackFailed,
		},
		{
			"stackId":        "222222",
			"stackName":      "photos-on-cards",
			"creationTime":   time.Now().Format(time.RFC1123),
			"environment":    "kaos",
			"slice":          "shea",
			"lineOfBusiness": "ISHI",
			"status":         cloudformation.StackStatusUpdateRollbackComplete,
		},
	}

}

func (repository *stackRepository) find(target string, tags []*cloudformation.Tag) string {
	for _, tag := range tags {
		if *tag.Key == target {
			return *tag.Value
		}
	}
	return ""
}
