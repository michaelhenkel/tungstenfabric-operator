// Copyright 2019 The Operator-SDK Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ansible

import (
<<<<<<< HEAD
	"github.com/operator-framework/operator-sdk/internal/pkg/scaffold"
	"github.com/operator-framework/operator-sdk/internal/pkg/scaffold/input"
=======
	"fmt"

	"github.com/operator-framework/operator-sdk/internal/pkg/scaffold"
	"github.com/operator-framework/operator-sdk/internal/pkg/scaffold/input"
	"github.com/operator-framework/operator-sdk/internal/pkg/scaffold/internal/deps"
>>>>>>> v0.0.4
)

// GopkgToml - the Gopkg.toml file for a hybrid operator
type GopkgToml struct {
<<<<<<< HEAD
	input.Input
=======
	StaticInput
>>>>>>> v0.0.4
}

func (s *GopkgToml) GetInput() (input.Input, error) {
	if s.Path == "" {
		s.Path = scaffold.GopkgTomlFile
	}
	s.TemplateBody = gopkgTomlTmpl
	return s.Input, nil
}

const gopkgTomlTmpl = `[[constraint]]
  name = "github.com/operator-framework/operator-sdk"
  # The version rule is used for a specific release and the master branch for in between releases.
  branch = "master" #osdk_branch_annotation
<<<<<<< HEAD
  # version = "=v0.6.0" #osdk_version_annotation
=======
  # version = "=v0.7.0" #osdk_version_annotation
>>>>>>> v0.0.4

[[override]]
  name = "k8s.io/api"
  version = "kubernetes-1.13.1"

[[override]]
  name = "k8s.io/apimachinery"
  version = "kubernetes-1.13.1"

[[override]]
  name = "k8s.io/client-go"
  version = "kubernetes-1.13.1"

[prune]
  go-tests = true
  unused-packages = true
`
<<<<<<< HEAD
=======

func PrintDepGopkgTOML(asFile bool) error {
	if asFile {
		_, err := fmt.Println(gopkgTomlTmpl)
		return err
	}
	return deps.PrintDepGopkgTOML(gopkgTomlTmpl)
}
>>>>>>> v0.0.4
