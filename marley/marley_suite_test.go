package marley_test

import (
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"

	"testing"
)

var (
	nsjamBinPath string
)

var _ = BeforeSuite(func() {
	var err error
	nsjamBinPath, err = gexec.Build("github.com/williammartin/nsjam")
	Expect(err).NotTo(HaveOccurred())
})

var _ = AfterSuite(func() {
	gexec.CleanupBuildArtifacts()
})

var execNSJam = func(args ...string) *gexec.Session {
	cmd := exec.Command(nsjamBinPath, args...)
	session, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
	Expect(err).NotTo(HaveOccurred())
	return session
}

func TestMarley(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Marley Suite")
}
