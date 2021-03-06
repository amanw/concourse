package testflight_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
)

var _ = Describe("A put step with inputs", func() {
	BeforeEach(func() {
		setAndUnpausePipeline("fixtures/put-inputs.yml")
	})

	Context("when specific inputs are specified", func() {
		It("attaches only the specified inputs to the put container", func() {
			watch := spawnFly("trigger-job", "-j", inPipeline("job-using-specified-inputs"), "-w")
			<-watch.Exited

			interceptS := fly("intercept", "-j", inPipeline("job-using-specified-inputs"), "-s", "some-resource", "--step-type", "put", "--", "ls")
			Expect(interceptS).To(gbytes.Say("specified-input"))
			Expect(string(interceptS.Out.Contents())).ToNot(ContainSubstring("all-input"))
		})
	})

	Context("when it uses all inputs", func() {
		It("attached all inputs to the put container", func() {
			watch := spawnFly("trigger-job", "-j", inPipeline("job-using-all-inputs"), "-w")
			<-watch.Exited

			interceptS := fly("intercept", "-j", inPipeline("job-using-all-inputs"), "-s", "some-resource", "--step-type", "put", "--", "ls")
			Expect(string(interceptS.Out.Contents())).To(ContainSubstring("all-input"))
			Expect(string(interceptS.Out.Contents())).To(ContainSubstring("specified-input"))
		})
	})

	Context("when an empty slice of inputs are specified", func() {
		It("no inputs are attached to the put container", func() {
			watch := spawnFly("trigger-job", "-j", inPipeline("job-using-empty-inputs"), "-w")
			<-watch.Exited

			interceptS := fly("intercept", "-j", inPipeline("job-using-empty-inputs"), "-s", "some-resource", "--step-type", "put", "--", "ls")
			Expect(string(interceptS.Out.Contents())).ToNot(ContainSubstring("all-input"))
			Expect(string(interceptS.Out.Contents())).ToNot(ContainSubstring("specified-input"))
		})
	})

	Context("when no inputs are specified", func() {
		It("attached all inputs to the put container", func() {
			watch := spawnFly("trigger-job", "-j", inPipeline("job-using-all-inputs"), "-w")
			<-watch.Exited

			interceptS := fly("intercept", "-j", inPipeline("job-using-all-inputs"), "-s", "some-resource", "--step-type", "put", "--", "ls")
			Expect(string(interceptS.Out.Contents())).To(ContainSubstring("all-input"))
			Expect(string(interceptS.Out.Contents())).To(ContainSubstring("specified-input"))
		})
	})
})
