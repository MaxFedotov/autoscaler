//go:build go1.7
// +build go1.7

package awsutil

import (
	"testing"

	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws"
)

type testPrettifyStruct struct {
	Field1 string
	Field2 *string
	Field3 []byte `sensitive:"true"`
	Value  []*string
}

func TestPrettify(t *testing.T) {
	cases := map[string]struct {
		Value  interface{}
		Expect string
	}{
		"general": {
			Value: testPrettifyStruct{
				Field1: "abc123",
				Field2: aws.String("abc123"),
				Field3: []byte("don't show me"),
				Value: []*string{
					aws.String("first"),
					aws.String("second"),
				},
			},
			Expect: `{
  Field1: "abc123",
  Field2: "abc123",
  Field3: <sensitive>,
  Value: ["first","second"],

}`,
		},
	}

	for d, c := range cases {
		t.Run(d, func(t *testing.T) {
			actual := StringValue(c.Value)
			if e, a := c.Expect, actual; e != a {
				t.Errorf("expect:\n%v\nactual:\n%v\n", e, a)
			}
		})
	}
}
