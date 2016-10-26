package cqt_test

import (
	"os/exec"
	"strconv"
	"syscall"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("ListNamespaces", func() {

	var (
		session *gexec.Session
	)

	Context("when there is a process running", func() {
		var cmd *exec.Cmd

		BeforeEach(func() {
			cmd = exec.Command("sleep", "60")
			Expect(cmd.Start()).To(Succeed())
		})

		AfterEach(func() {
			Expect(cmd.Process.Signal(syscall.SIGKILL)).To(Succeed())
		})

		Context("and we list namespaces by pid", func() {
			BeforeEach(func() {
				session = execNSJam("list-namespaces", "-p", strconv.Itoa(cmd.Process.Pid))
			})

			It("exits with a zero exit code", func() {
				Eventually(session).Should(gexec.Exit(0))
			})

			It("prints the name of the process", func() {
				Eventually(session.Out).Should(gbytes.Say("sleep"))
			})

			It("prints the namespace inodes", func() {
				Expect(string(session.Wait().Out.Contents())).To(MatchRegexp(`mnt:\d+`))
				Expect(string(session.Wait().Out.Contents())).To(MatchRegexp(`user:\d+`))
				Expect(string(session.Wait().Out.Contents())).To(MatchRegexp(`ipc:\d+`))
				Expect(string(session.Wait().Out.Contents())).To(MatchRegexp(`pid:\d+`))
				Expect(string(session.Wait().Out.Contents())).To(MatchRegexp(`net:\d+`))
			})
		})
	})

	Context("when we list namespaces for a pid that doesn't exist", func() {
		BeforeEach(func() {
			session = execNSJam("list-namespaces", "-p", "-1")
		})

		It("exits with a non-zero exit code", func() {
			Eventually(session).ShouldNot(gexec.Exit(0))
		})

		It("prints an informative error message", func() {
			Eventually(session.Err).Should(gbytes.Say("no process found with pid: -1"))
		})
	})
})
