package GraGO

import (
	"log"
	"testing"
)

func Test_getCommitHash(t *testing.T) {
	//type args struct {
	//	label string
	//}
	//tests := []struct {
	//	name string
	//	args args
	//}{
	//	// TODO: Add test cases.
	//}
	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//	})
	//}
	oid, err := getOidFromPos("HEAD")
	if err != nil {
		t.Error(err)
	}
	log.Printf(oid.String())
}
