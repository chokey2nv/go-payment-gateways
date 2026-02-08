package paystack_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestPaystack(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Paystack Suite")
}
