package cmd

// Copyright © 2019 oleg2807@gmail.com
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

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"

	"github.com/olegsu/iris/pkg/util"
	"github.com/olegsu/iris/pkg/v2/js"
	"github.com/olegsu/iris/pkg/v2/spec"
)

var v1Cmd = &cobra.Command{
	Use:  "v1",
	Long: "Start the server",
	Run: func(cmd *cobra.Command, args []string) {
		logger := buildLogger("v1")
		engine := js.NewJSEngine()

		var script string
		{
			f, err := ioutil.ReadFile("/Users/olsynt/workspace/personal/iris/iris.js")
			if err != nil {
				logger.Error("error reading file", "file", "f", "err", err.Error())
				os.Exit(1)
			}
			logger.Debug("File loaded")
			script = string(f)
		}

		_, err := engine.Load([]string{script}, nil)
		if err != nil {
			logger.Error("error loading script to engine", "err", err.Error())
		}
		res, err := engine.CallFn("buidIRISObject", nil)
		if err != nil {
			logger.Error("error getting result from getFilters function", "err", err.Error())
		}
		spec := &spec.Spec{}
		util.UnmarshalOrDie([]byte(res.String()), spec, logger)
		fmt.Println(spec)
		for _, f := range spec.Filters {
			engine.CallFn(f.Name, nil)
		}
	},
}

func init() {
	rootCmd.AddCommand(v1Cmd)
}
