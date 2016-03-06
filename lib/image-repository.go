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

type imageRepository interface {
	Fetch(matcher string) []map[string]string
}

type ec2imageRepository struct {
	images []map[string]string
}

func (repository *ec2imageRepository) Fetch(matcher string) []map[string]string {

	if repository.images == nil {
		fmt.Println("fresh cache")
		repository.fetch()
	} else {
		fmt.Println("re-priming cache")
		go repository.fetch()
	}

	if matcher == "" {
		return repository.images
	}

	var images []map[string]string

	for _, image := range repository.images {
		regex, _ := regexp.Compile(".*" + matcher + ".*")
		if matched := regex.MatchString(image["name"]); matched {
			images = append(images, image)
		}
	}

	return images

}

func (repository *ec2imageRepository) fetch() {

	ec2Service := ec2.New(session.New(), &aws.Config{
		Region: aws.String("ap-southeast-2"),
		Credentials: credentials.NewStaticCredentials(

	resp, err := ec2Service.DescribeImages(nil)

	if err != nil {
		fmt.Println("An error occured while fetching images from AWS")
		fmt.Println(err)
		return
	}

	var images []map[string]string
	for _, image := range resp.Images {
		if image.Name != nil {
			images = append(images, map[string]string{
				"name":        *image.Name,
				"created":     *image.CreationDate,
				"id":          *image.ImageId,
				"environment": repository.find("Environment", image.Tags),
				"solution":    repository.find("Application", image.Tags),
				"slice":       repository.find("Slice", image.Tags)})
		}
	}

	repository.images = images

}

func (repository *ec2imageRepository) find(target string, tags []*ec2.Tag) string {
	for _, tag := range tags {
		if *tag.Key == target {
			return *tag.Value
		}
	}
	return ""
}

//FakeImageRepository ...
type FakeImageRepository struct{}

//Fetch ...
func (f FakeImageRepository) Fetch(matcher string) []map[string]string {
	return []map[string]string{
		{"name": "foo", "created": time.Now().Format(time.RFC1123), "id": "boop"},
		{"name": "bar", "created": time.Now().Format(time.RFC1123), "id": "zoop"}}
}
