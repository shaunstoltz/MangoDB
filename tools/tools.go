// Copyright 2021 FerretDB Inc.
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

//go:build go1.18
// +build go1.18

package tools // import "github.com/FerretDB/FerretDB/tools"

import (
	_ "github.com/BurntSushi/go-sumtype"
	_ "github.com/go-task/task/v3/cmd/task"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/quasilyte/go-consistent"
	_ "github.com/reviewdog/reviewdog/cmd/reviewdog"
	_ "golang.org/x/perf/cmd/benchstat"
	_ "golang.org/x/tools/cmd/stringer"
	_ "mvdan.cc/gofumpt"
)

// Check that `go` in $PATH have the right version.
// Catches problems like `/some/path/go generate` invocations where `/some/path/go` is 1.18+
// (that's checked by the build tags above), but just `go` in $PATH (typically something like `/usr/bin/go`)
// is an earlier version.

//go:generate go run check.go

//go:generate go build -v -o ../bin/ github.com/BurntSushi/go-sumtype
//go:generate go build -v -o ../bin/ github.com/go-task/task/v3/cmd/task
//go:generate go build -v -o ../bin/ github.com/golangci/golangci-lint/cmd/golangci-lint
//go:generate go build -v -o ../bin/ github.com/quasilyte/go-consistent
//go:generate go build -v -o ../bin/ github.com/reviewdog/reviewdog/cmd/reviewdog
//go:generate go build -v -o ../bin/ golang.org/x/perf/cmd/benchstat
//go:generate go build -v -o ../bin/ golang.org/x/tools/cmd/stringer
//go:generate go build -v -o ../bin/ mvdan.cc/gofumpt
