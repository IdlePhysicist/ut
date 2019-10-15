package main

import (
  "errors"
  "fmt"
  "time"
  "strings"
  "os"

  flag "github.com/spf13/pflag"
)


// Mon Jan 2 15:04:05 -0700 MST 2006
var (
  ref =   map[string]string{`YYYY`: `2006`, `MM`: `01`, `DD`: `02`, `hh`: `15`, `mm`: `04`, `ss`: `05`}
  version string
  commit  string
  versionFlg bool
)

type Args struct {
  format string
  input  string
  noon   bool
}

var args Args

func init() {
  // Parse args 
  flag.BoolVarP(&args.noon, `noon`, `n`, false, "Make returned datetime noon on that date")
  flag.StringVarP(&args.format, `format`, `f`, ``, "The format to be used by the program")
  flag.BoolVarP(&versionFlg, `version`, `v`, false, "Print version and exit")
  flag.Parse()
  args.input = flag.Arg(0)

  if versionFlg {
    fmt.Printf("ut version: %s\nbuild commit: %s\n", version, commit)
    os.Exit(0)
  }
}

func main() {
  format, err := convertFormat(args.format)
  if err != nil {
    fmt.Println("Error malformed format")
    os.Exit(1)
  }

  if args.noon {
    formatSlice := []string{format, `15:04:05Z`}
    format = strings.Join(formatSlice, `T`)

    inSlice := []string{args.input, `12:00:00Z`}
    args.input = strings.Join(inSlice, `T`)
  }

  d, err := time.Parse(format, args.input)
  if err != nil {
    panic(err)
  }

  fmt.Println(d.Unix())
}

func convertFormat(format string) (string, error) {
  if strings.Count(format, `h`) != 2 {
    return ``, errors.New(`Length of day format is not two`)
  }
  format = strings.Replace(format, `hh`, ref[`hh`], 1)

  if strings.Count(format, `m`) != 2 {
    return ``, errors.New(`Length of day format is not two`)
  }
  format = strings.Replace(format, `mm`, ref[`mm`], 1)

  if strings.Count(format, `s`) != 2 {
    return ``, errors.New(`Length of day format is not two`)
  }
  format = strings.Replace(format, `ss`, ref[`ss`], 1)

  if strings.Count(format, `D`) != 2 {
    return ``, errors.New(`Length of day format is not two`)
  }
  format = strings.Replace(format, `DD`, ref[`DD`], 1)

  if strings.Count(format, `M`) != 2 {
    return ``, errors.New(`Length of month format is not two`)
  }
  format = strings.Replace(format, `MM`, ref[`MM`], 1)

  if strings.Count(format, `Y`) != 4 {
    return ``, errors.New(`Length of year format is not four`)
  }
  format = strings.Replace(format, `YYYY`, ref[`YYYY`], 1)

  return format, nil
}
