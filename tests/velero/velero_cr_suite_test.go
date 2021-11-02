package oadp

import (
	"fmt"
	"path/filepath"
	"testing"

	. "oadp/lib"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestVelero(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Velero Suite")
}

var _ = Describe("Test Velero CR with", func() {
	It("Noobaa storage", func() {
		// pods, err := GetPodsInNamespace("")
		// if err != nil {
		// 	fmt.Printf(" Error : %s \n", err.Error())
		// }
		// fmt.Printf("Pods in namespace default \n")
		// for _, pod := range pods.Items {
		// 	fmt.Printf("  %s\n", pod.Name)
		// }

		credsFilePath, _ := filepath.Abs("../../data/aws_credentials")

		err := CreateSecret(credsFilePath, "nitish", "cloud-credentials")
		if err != nil {
			fmt.Printf("error occured %s", err.Error())
			Expect(err).NotTo(HaveOccurred())
		}
	})
})
