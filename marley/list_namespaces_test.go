package marley_test

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("ListNamespaces", func() {

	var (
		session *gexec.Session
		cmd     *exec.Cmd

		targetNS string
	)

	BeforeEach(func() {
		targetNS = ""
	})

	JustBeforeEach(func() {
		cmd = exec.Command("sleep", "60")
		Expect(cmd.Start()).To(Succeed())

		session = execNSJam("list-namespaces", targetNS, "--target", strconv.Itoa(cmd.Process.Pid))
	})

	AfterEach(func() {
		Expect(cmd.Process.Signal(syscall.SIGKILL)).To(Succeed())
	})

	Context("when listing the pid namespace", func() {
		BeforeEach(func() {
			targetNS = "--pid"
		})

		It("exits with a zero exit code", func() {
			Eventually(session).Should(gexec.Exit(0))
		})

		It("prints out the name of the target", func() {
			Eventually(session).Should(gbytes.Say("sleep"))
		})

		It("prints out the pid namespace of the target", func() {
			Expect(string(session.Wait().Out.Contents())).To(ContainSubstring(parsePidNS(cmd.Process.Pid)))
		})
	})
})

func parsePidNS(pid int) string {
	pidNS, err := os.Readlink(fmt.Sprintf("/proc/%d/ns/pid", pid))
	Expect(err).NotTo(HaveOccurred())

	return strings.TrimSpace(pidNS)
}
