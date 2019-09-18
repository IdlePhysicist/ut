package main

import (
  "flag"
  "fmt"
  "time"
  "strings"
)

type Args struct {
  input string
  noon  bool
}

func parse() *Args {
  // Parse args 
  var args Args
  flag.Parse()
  flag.BoolVar(&args.noon,`n`, true, `fudge the time`) 
  args.input = flag.Args()[0]
  return &args
}

func main() {
  in := parse()

  layout := `2006-01-02`
  if in.noon {
    layout = `2006-01-02T15:04:05Z`
    newSlice := []string{in.input, `12:00:00Z`}
    in.input = strings.Join(newSlice, `T`)
  }
  d, err := time.Parse(layout, in.input)
  if err != nil {
    //d, err = time.Parse(time.RFC3339, in.input)
    //if err != nil {
    panic(err)
    //}
  }

  fmt.Println(d.Unix())
}

