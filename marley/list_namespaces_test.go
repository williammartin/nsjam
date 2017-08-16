package marley_test

import (
	"fmt"
	"math"
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
		args    []string

		targetNS string
	)

	BeforeEach(func() {
		targetNS = ""
		args = []string{"list-namespaces", targetNS}
	})

	JustBeforeEach(func() {
		session = execNSJam(args...)
	})

	AfterEach(func() {
		Expect(cmd.Process.Signal(syscall.SIGKILL)).To(Succeed())
	})

	Context("when listing the pid namespace", func() {
		BeforeEach(func() {
			targetNS = "--pid"
		})

		Context("when a valid target pid is provided", func() {
			BeforeEach(func() {
				cmd = exec.Command("sleep", "60")
				Expect(cmd.Start()).To(Succeed())

				args = append(args, "--target", strconv.Itoa(cmd.Process.Pid))
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

		Context("when a target is not provided", func() {
			It("exits with non zero exit code", func() {
				Eventually(session).ShouldNot(gexec.Exit(0))
			})

			It("prints an informative error message", func() {
				Eventually(session.Err).Should(gbytes.Say("Must pass a target pid"))
			})
		})

		Context("when a non-existent pid is provided", func() {
			BeforeEach(func() {
				args = append(args, "--target", strconv.Itoa(math.MaxInt64))
			})

			It("exits with non zero exit code", func() {
				Eventually(session).ShouldNot(gexec.Exit(0))
			})

			It("prints an informative error message", func() {
				Eventually(session.Err).Should(gbytes.Say("Must pass a valid pid"))
			})
		})
	})

})

func parsePidNS(pid int) string {
	pidNS, err := os.Readlink(fmt.Sprintf("/proc/%d/ns/pid", pid))
	Expect(err).NotTo(HaveOccurred())

	return strings.TrimSpace(pidNS)
}
