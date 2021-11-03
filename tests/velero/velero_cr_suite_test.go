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

var _ = BeforeSuite(func() {
	// Create cloud-credentials secret in oadp namespace
	credsFilePath, _ := filepath.Abs("../../data/aws_credentials")
	err := CreateSecret(credsFilePath, "openshift-adp", "cloud-credentials")
	if err != nil {
		fmt.Printf("An error occured while creating secret : %s \n", err.Error())
		Expect(err).NotTo(HaveOccurred())
	}
})

var _ = Describe("Test Velero CR with", func() {
	It("Noobaa storage", func() {
		// Create Velero CR and verify it is successful
		fmt.Print("Running test case to create Velero CR with noobaa storage ...\n")
		veleroFilePath, _ := filepath.Abs("../../data/velero_cr_noobaa.yaml")
		err := CreateVeleroCRD(veleroFilePath, "openshift-adp")
		if err != nil {
			Expect(err).NotTo(HaveOccurred())
		}
	})
})

var _ = AfterSuite(func() {
	// TODO Cleanup logic
})
