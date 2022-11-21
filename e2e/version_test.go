// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package e2e

import (
	"os/exec"

	"github.com/onsi/ginkgo/v2"
	"github.com/runfinch/common-tests/option"
)

var testVersion = func(o *option.Option) {
	ginkgo.Context("Version", func() {
		ginkgo.Specify("Test version", func() {
			exec.Command("finch", "version").Run()
			//gomega.Expect(command.StdoutStr(o, "version")).To(gomega.Equal(fmt.Sprintf("Finch version: %s", "v0.1.0")))
		})
	})
}
