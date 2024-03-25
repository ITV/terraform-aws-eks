/*
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package scale_test

import (
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/aws/karpenter/test/pkg/environment/aws"
)

var env *aws.Environment

func TestScale(t *testing.T) {
	RegisterFailHandler(Fail)
	BeforeSuite(func() {
		env = aws.NewEnvironment(t)
		SetDefaultEventuallyTimeout(time.Hour)
	})
	AfterSuite(func() {
		env.Stop()
	})
	RunSpecs(t, "Alpha/Scale")
}

var _ = BeforeEach(func() {
	env.ExpectPrefixDelegationEnabled()
	env.BeforeEach()
})
var _ = AfterEach(func() { env.Cleanup() })
var _ = AfterEach(func() {
	env.AfterEach()
	env.ExpectPrefixDelegationDisabled()
})
