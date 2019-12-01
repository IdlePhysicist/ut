package ut

import (
  "errors"
  "time"
  "strings"
)


// Mon Jan 2 15:04:05 -0700 MST 2006
var ref = map[string]string{
  `YYYY`: `2006`, `MM`: `01`, `DD`: `02`,
  `hh`: `15`, `mm`: `04`, `ss`: `05`,
}



func Ut(input, format string, noon bool) (int64, error) {
  format, err := convertFormat(format)
  if err != nil {
    return 0, err
  }

  if noon {
    format = strings.Join([]string{format, `15:04:05Z`}, `T`)
    input  = strings.Join([]string{input,  `12:00:00Z`}, `T`)
  }

  d, err := time.Parse(format, input)
  if err != nil {
    return 0, err
  }

  return d.Unix(), err
}

func convertFormat(format string) (string, error) {
  if strings.Contains(format, `D`) {
    if strings.Count(format, `D`) != 2 {
      return ``, errors.New(`Length of day format is not two`)
    }
    format = strings.Replace(format, `DD`, ref[`DD`], 1)
  }

  if strings.Contains(format, `M`) {
    if strings.Count(format, `M`) != 2 {
      return ``, errors.New(`Length of month format is not two`)
    }
    format = strings.Replace(format, `MM`, ref[`MM`], 1)
  }

  if strings.Contains(format, `Y`) {
    if strings.Count(format, `Y`) != 4 {
      return ``, errors.New(`Length of year format is not four`)
    }
    format = strings.Replace(format, `YYYY`, ref[`YYYY`], 1)
  }

  if strings.Contains(format, `h`) {
    if strings.Count(format, `h`) != 2 {
      return ``, errors.New(`Length of hour format is not two`)
    }
    format = strings.Replace(format, `hh`, ref[`hh`], 1)
  }

  if strings.Contains(format, `m`) {
    if strings.Count(format, `m`) != 2 {
      return ``, errors.New(`Length of minute format is not two`)
    }
    format = strings.Replace(format, `mm`, ref[`mm`], 1)
  }

  if strings.Contains(format, `s`) {
    if strings.Count(format, `s`) != 2 {
      return ``, errors.New(`Length of second format is not two`)
    }
    format = strings.Replace(format, `ss`, ref[`ss`], 1)
  }

  return format, nil
}
