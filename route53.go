package main

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/route53"
)

func getZoneId() string {
	zoneId := os.Getenv("AWS_ROUTE53_DOMAIN")
	return zoneId
}

func getRecordName() string {
	record := os.Getenv("AWS_ROUTE53_RECORD")
	return record
}

func getSession() *session.Session {

	sess, err := session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(
			os.Getenv("AWS_ACCESS_KEY_ID"),
			os.Getenv("AWS_SECRET_ACCESS_KEY"),
			"",
		),
	})

	if err != nil {
		panic(err)
	}
	return sess
}
func getRecordIp() string {
	sess := getSession()
	svc := route53.New(sess)
	params := &route53.ListResourceRecordSetsInput{
		HostedZoneId: aws.String(getZoneId()),
	}
	respList, err := svc.ListResourceRecordSets(params)

	if err != nil {
		log.Fatal(err)
	}
	for _, recordSet := range respList.ResourceRecordSets {
		if *recordSet.Name == getRecordName() {
			log.Println("Record IP: " + *recordSet.ResourceRecords[0].Value)
			return *recordSet.ResourceRecords[0].Value
		}
	}
	log.Println("Record IP: None")
	return ""
}

func changeRecord(globalIp string) {
	sess := getSession()
	svc := route53.New(sess)

	params := &route53.ChangeResourceRecordSetsInput{
		ChangeBatch: &route53.ChangeBatch{
			Changes: []*route53.Change{
				{
					Action: aws.String("UPSERT"),
					ResourceRecordSet: &route53.ResourceRecordSet{
						Name: aws.String(getRecordName()),
						Type: aws.String("A"),
						ResourceRecords: []*route53.ResourceRecord{
							{
								Value: aws.String(globalIp),
							},
						},
						TTL:           aws.Int64(300),
						Weight:        aws.Int64(100),
						SetIdentifier: aws.String("Arbitrary Id describing this change set"),
					},
				},
			},
			Comment: aws.String("update."),
		},
		HostedZoneId: aws.String(getZoneId()),
	}
	resp, err := svc.ChangeResourceRecordSets(params)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Change Response:")
	log.Println(resp)
}
