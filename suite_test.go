package gblob_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGBlob(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GBlob Suite")
}
