package kinesis

import (
	"net/http"
	"strings"
	"testing"
)

var testSignFactoryDataUsEast1 = []struct {
	AWS_KEY    string
	AWS_SECRET string
	TOKEN      string
	DateHeader string
	AuthHeader string
}{
	{"ASWKEY", "AWSSECRET", "TOKEN1", "Thu, 28 Nov 2013 15:04:05 GMT", "AWS4-HMAC-SHA256 Credential=ASWKEY/20131128/us-east-1/kinesis/aws4_request, SignedHeaders=content-type;date;host;user-agent;x-amz-target, Signature=6c21aca39f1d4afd383fbc45dd3a580192036162f74bf9fda6cad6c6fb7cde2f"},
	{"ASWKEY2", "AWSSECRET2", "TOKEN2", "Thu, 28 Nov 2013 15:04:05 GMT", "AWS4-HMAC-SHA256 Credential=ASWKEY2/20131128/us-east-1/kinesis/aws4_request, SignedHeaders=authorization;content-type;date;host;user-agent;x-amz-target, Signature=0b442a629ffe0a405f718f8e50ffdbe3886574687fe3d60dffcc09d67e4aff5a"},
	{"ASWNEWKEY", "AWSSECRET", "TOKEN3", "Thu, 28 Nov 2013 15:04:05 GMT", "AWS4-HMAC-SHA256 Credential=ASWNEWKEY/20131128/us-east-1/kinesis/aws4_request, SignedHeaders=authorization;content-type;date;host;user-agent;x-amz-target, Signature=f92bd23b4e9f6163779c93dee8b34c673ae394f934b4562dcebfdf4adef9685e"},
	{"ASWKEY", "AWSSECRET", "TOKEN4", "Mon, 25 Nov 2013 15:04:05 GMT", "AWS4-HMAC-SHA256 Credential=ASWKEY/20131125/us-east-1/kinesis/aws4_request, SignedHeaders=authorization;content-type;date;host;user-agent;x-amz-target, Signature=cabe7376fd1b308e8fda031be50a013509dba601445573c5527c1be205c59fa5"},
}

var testSignFactoryDataCustomKinesis = []struct {
	AWS_KEY    string
	AWS_SECRET string
	TOKEN      string
	DateHeader string
	AuthHeader string
}{
	{"ASWKEY", "AWSSECRET", "TOKEN1", "Thu, 28 Nov 2013 15:04:05 GMT", "AWS4-HMAC-SHA256 Credential=ASWKEY/20131128///aws4_request, SignedHeaders=content-type;date;host;user-agent;x-amz-target, Signature=a1a3c571ad100fd1483a5c699475f84b0407c37a76b9c20fa53b579b36033930"},
	{"ASWKEY2", "AWSSECRET2", "TOKEN2", "Thu, 28 Nov 2013 15:04:05 GMT", "AWS4-HMAC-SHA256 Credential=ASWKEY2/20131128///aws4_request, SignedHeaders=authorization;content-type;date;host;user-agent;x-amz-target, Signature=fd47c35be18a84323122fefbe9cccbf251bbf06d4c75d701e4ef89393f89445d"},
	{"ASWNEWKEY", "AWSSECRET", "TOKEN3", "Thu, 28 Nov 2013 15:04:05 GMT", "AWS4-HMAC-SHA256 Credential=ASWNEWKEY/20131128///aws4_request, SignedHeaders=authorization;content-type;date;host;user-agent;x-amz-target, Signature=a4bebba353dfcccf93c8eca519005b9473f8a9f181e07ead591ee9fc1d808287"},
	{"ASWKEY", "AWSSECRET", "TOKEN4", "Mon, 25 Nov 2013 15:04:05 GMT", "AWS4-HMAC-SHA256 Credential=ASWKEY/20131125///aws4_request, SignedHeaders=authorization;content-type;date;host;user-agent;x-amz-target, Signature=cf07b40df788d4428f312147f2eea634aea161a7813a11250d38056c8adb2749"},
}

func TestSign(t *testing.T) {
	request, err := http.NewRequest("POST", "https://kinesis.us-east-1.amazonaws.com", strings.NewReader("{}"))
	if err != nil {
		t.Errorf("NewRequest Error %v", err)
	}

	request.Header.Set("Content-Type", "application/x-amz-json-1.1")
	request.Header.Set("X-Amz-Target", "")
	request.Header.Set("User-Agent", "Golang Kinesis")

	for _, data := range testSignFactoryDataUsEast1 {
		request.Header.Set("Date", data.DateHeader)
		err = Sign(NewAuth(data.AWS_KEY, data.AWS_SECRET, data.TOKEN), request)
		if err != nil {
			t.Errorf("Error on sign (%v)", err)
			continue
		}
		if request.Header.Get("Authorization") != data.AuthHeader {
			t.Errorf("Get this header (%v), but expect this (%v)", request.Header.Get("Authorization"), data.AuthHeader)
		}
	}
}

func TestSignCustom(t *testing.T) {
	request, err := http.NewRequest("POST", "http://kinesis.custom", strings.NewReader("{}"))
	if err != nil {
		t.Errorf("NewRequest Error %v", err)
	}

	request.Header.Set("Content-Type", "application/x-amz-json-1.1")
	request.Header.Set("X-Amz-Target", "")
	request.Header.Set("User-Agent", "Golang Kinesis")

	for _, data := range testSignFactoryDataCustomKinesis {
		request.Header.Set("Date", data.DateHeader)
		err = Sign(NewAuth(data.AWS_KEY, data.AWS_SECRET, data.TOKEN), request)
		if err != nil {
			t.Errorf("Error on sign (%v)", err)
			continue
		}
		if request.Header.Get("Authorization") != data.AuthHeader {
			t.Errorf("Get this header (%v), but expect this (%v)", request.Header.Get("Authorization"), data.AuthHeader)
		}
	}
}
