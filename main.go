package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/route53"
)

func main() {
	globalIp := getGlobalIp()
	fmt.Println(globalIp)

	domain := os.Getenv("AWS_ROUTE53_DOMAIN")
	record := os.Getenv("AWS_ROUTE53_RECORD")

	client := route53.New(&aws.Config{
		Credentials: credentials.NewStaticCredentials(
			os.Getenv("AWS_ACCESS_KEY_ID"),
			os.Getenv("AWS_SECRET_ACCESS_KEY"),
			""),
	})

	hzresp, err := client.ListHostedZonesByName(&route53.ListHostedZonesByNameInput{DNSName: aws.String(domain)})
	if err != nil {
		panic(err)
	}
	HostedZoneID := hzresp.HostedZones[0].ID
	rrsresp, err := client.ListResourceRecordSets(&route53.ListResourceRecordSetsInput{HostedZoneID: HostedZoneID})
	if err != nil {
		panic(err)
	}

	for i := range rrsresp.ResourceRecordSets {
		fmt.Println(awsutil.StringValue(*rrsresp.ResourceRecordSets[i]))
	}
}

func getGlobalIp() string {
	inetUrl := "http://inet-ip.info/ip"

	body, err := httpRequestBody(inetUrl)
	if err == nil {
		return body
	}
	httpbinUrl := "http://httpbin.org/ip"
	body, err = httpRequestBody(httpbinUrl)

	globalIps := &struct{ Origin string }{}
	err = json.Unmarshal([]byte(body), &globalIps)

	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(globalIps.Origin, ",")[0]
}

func httpRequestBody(url string) (string, error) {
	resp, err := http.Get(url)

	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func listResourses(svc *route53.Route53) {
	listParams := &route53.ListResourceRecordSetsInput{
		HostedZoneId: aws.String(),
	}
	respList, err := svc.ListResourceRecordSets(listParams)
	return respList
}
