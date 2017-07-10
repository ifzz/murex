package typemgmt

import (
	"bufio"
	"errors"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
	"strconv"
	"strings"
)

func init() {
	proc.GoFunctions["table"] = proc.GoFunction{Func: cmdTable, TypeIn: types.Generic, TypeOut: types.Csv}
	proc.GoFunctions["format"] = proc.GoFunction{Func: cmdFormat, TypeIn: types.Generic, TypeOut: types.Generic}
}

func cmdTable(p *proc.Process) (err error) {
	p.Stdout.SetDataType(types.Csv)

	separator, err := p.Parameters.String(0)
	if err != nil {
		return
	}

	/*dt := p.Stdin.GetDataType()
	if dt == types.String || dt == types.Generic {
		t := table.NewTable()
		var copyFailed bool
		go func() {
			_, err = io.Copy(t, p.Stdin)
			if err != nil {
				copyFailed = true
				return
			}
		}()

		for {
			if copyFailed {
				return err
			}
			row, eof := t.GetRow()
			if eof || copyFailed {
				return err
			}

			_, err = p.Stdout.Writeln([]byte(strings.Join(row, separator)))
			if err != nil {
				t.Close()
				return err
			}
		}
	}*/

	var (
		a []string
		s string
	)

	join := func(b []byte) {
		a = append(a, string(b))
	}

	if p.IsMethod {
		p.Stdin.ReadArray(join)
		s = strings.Join(a, separator)
	} else {
		s = strings.Join(p.Parameters.StringArray()[1:], string(separator))
	}

	_, err = p.Stdout.Writeln([]byte(s))
	return
}

func cmdFormat(p *proc.Process) (err error) {
	format, err := p.Parameters.String(0)
	if err != nil {
		return
	}

	p.Stdout.SetDataType(format)

	dt := p.Stdin.GetDataType()

	if dt == types.String || dt == types.Generic {

		if format == types.Json {
			var (
				headings []string
				jObj     []map[string]string
			)

			scanner := bufio.NewScanner(p.Stdin)
			for scanner.Scan() {
				s := scanner.Text()
				split := strings.Split(s, " ")
				if len(headings) == 0 {
					headings = split
				} else {
					m := make(map[string]string)
					for i := range split {
						if i >= len(headings) {
							m[strconv.Itoa(i)] = split[i]
						} else {
							m[headings[i]] = split[i]
						}
						jObj = append(jObj, m)
					}
				}
			}
			if err := scanner.Err(); err != nil {
				return err
			}

			b, err := utils.JsonMarshal(jObj)
			if err != nil {
				return err
			}

			_, err = p.Stdout.Writeln(b)
			return err
		}
	}

	return errors.New("I don't know how to convert this data")
}
