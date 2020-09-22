package main

import (
	"fmt"
	"os"

	flag "github.com/spf13/pflag"

	"github.com/idlephysicist/ut"
)

var (
	version, commit string
	versionFlg bool

	format, input string
	noon   bool
)

func main() {
	flag.BoolVarP(&versionFlg, `version`, `v`, false, "Print version and exit")
	flag.BoolVarP(&noon, `noon`, `n`, false, "Make returned datetime noon on that date")
	flag.StringVarP(&format, `format`, `f`, ``, "The format to be used by the program")
	flag.Parse()
	input = flag.Arg(0)

	if versionFlg {
		fmt.Printf("ut version: %s\nbuild commit: %s\n", version, commit)
		os.Exit(0)
	}

	result, err := ut.Ut(input, format, noon)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Println(result)
	}
}
