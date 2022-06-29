/*
Copyright © 2022 Matthew A. Miller <linuxwolf@outer-planes.net>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/linuxwolf/infra-utils/envs/pkg"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var (
	verbosity int = 0
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "envs",
	Short: "Assembles environment variables from files",
	Long:  `envs assembles environment variables from files`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		processor := NewProcessor()

		for _, fname := range args {
			processor.processFile(fname)
		}

		fmt.Fprintln(os.Stdout, processor.envs)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().CountVarP(&verbosity, "verbose", "v", "verbose logging")
}

type Processor struct {
	envs   *pkg.Env
	log    *zap.SugaredLogger
	parser *pkg.Parser
}

func NewProcessor() *Processor {
	logger := pkg.SetupLogging(verbosity)
	proc := Processor{
		envs:   pkg.NewEnvsFromEnviron(),
		log:    logger,
		parser: pkg.NewParser(logger),
	}

	return &proc
}

func (p *Processor) processFile(fname string) {
	var err error
	finfo, err := os.Lstat(fname)
	if err != nil {
		p.log.Warnf("could not stat %s: %v", fname, err)
		return
	}

	// QUESTION: handle directories?
	if !finfo.Mode().IsRegular() {
		p.log.Warnf("%s is not supported", fname)
		return
	}

	fd, err := os.Open(fname)
	if err != nil {
		p.log.Warnf("could not open %s: %v", fname, err)
		return
	}

	next := p.parser.ProcessReader(fd)
	p.envs = p.envs.Including(next)
}
