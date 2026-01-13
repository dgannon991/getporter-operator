package v1

import (
	"testing"

	"k8s.io/apimachinery/pkg/runtime"
)

func TestAddKnownTypes(t *testing.T) {
	scheme := runtime.NewScheme()
	err := AddToScheme(scheme)
	if err != nil {
		t.Fatalf("failure to add to Scheme %v", err)
	}
	err = addKnownTypes(scheme)
	if err != nil {
		t.Fatalf("failure to add known types %v", err)
	}
}
