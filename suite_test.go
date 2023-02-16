package gblob_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGblob(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gblob Suite")
}
