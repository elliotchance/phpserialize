package phpserialize_test

import (
	"github.com/crowley-io/macchiato"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestPhpserialize(t *testing.T) {
	RegisterFailHandler(Fail)
	macchiato.RunSpecs(t, "Phpserialize Suite")
}
