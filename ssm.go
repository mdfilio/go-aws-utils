package main

import (
	"fmt"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

func getSSM(region string, humanregion string, goGroup *sync.WaitGroup) {
	defer goGroup.Done()

	ssmresp, err := ssmsvc.DescribeInstanceInformation(nil)
	if err != nil {
		panic(err)
	}
}

func main() {

	awsregions := map[int][]string{
		0:  {"us-east-1", "N. Virginia"},
		1:  {"us-east-2", "Ohio"},
		2:  {"us-west-1", "N. California"},
		3:  {"us-west-2", "Oregon"},
		4:  {"ca-central-1", "Canada (Central)"},
		5:  {"eu-west-1", "Ireland"},
		6:  {"eu-central-1", "Frankfurt"},
		7:  {"eu-west-2", "London"},
		8:  {"ap-southeast-1", "Singapore"},
		9:  {"ap-southeast-2", "Sydney"},
		10: {"ap-northeast-1", "Tokyo"},
		11: {"ap-northeast-2", "Seoul"},
		12: {"sa-east-1", "São Paulo"},
		13: {"ap-south-1", "Mumbai"},
	}
	goGroup := new(sync.WaitGroup)
	defer goGroup.Wait()

	for i := 0; i < len(awsregions); i++ {
		goGroup.Add(1)
		go getSSM(awsregions[i][0], awsregions[i][1], goGroup)
	}

}
