// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build !windows

// Package gapicgen provides some helpers for gapicgen binaries.
package gapicgen

import (
	"fmt"
	"os"

	"cloud.google.com/go/internal/gapicgen/execv"
)

// VerifyAllToolsExist ensures that all required tools exist on the system.
func VerifyAllToolsExist(toolsNeeded []string) error {
	for _, t := range toolsNeeded {
		c := execv.Command("which", t)
		c.Stderr = os.Stderr
		if c.Run() != nil {
			return fmt.Errorf("%s does not appear to be installed. please install it. all tools needed: %v", t, toolsNeeded)
		}
	}
	return nil
}
