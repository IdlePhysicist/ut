package ut

import "testing"

func TestDate(t *testing.T) {
  date   := [4]string{`20191016`, `2019/10/16`, `16-2019-10`, `10.16.2019`}
  format := [4]string{`YYYYMMDD`, `YYYY/MM/DD`, `DD-YYYY-MM`, `MM.DD.YYYY`}
  expect := int64(1571227200)

  for i := range format {
    result, err := Ut(date[i], format[i], true)
    if err != nil {
      t.Errorf(`FAILED: %v`, err)
    }

    if result != expect {
      t.Errorf(`FAILED: expected %v got %v`, expect, result)
    }
  }
}
