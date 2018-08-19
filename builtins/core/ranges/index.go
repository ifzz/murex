package ranges

import (
	"errors"
	"strconv"
)

type rfIndex struct {
	start int
	end   int
	i     int
}

func (rf *rfIndex) Start(_ []byte) bool {
	rf.i++
	return rf.i > rf.start
}

func (rf *rfIndex) End(_ []byte) bool {
	if rf.end > -1 {
		rf.i++
		return rf.i > rf.end+1
	} else {
		return false
	}
}

func newNumber(r *rangeParameters) (err error) {
	rf := new(rfIndex)

	if r.Start == "" {
		r.Start = "0"
	}

	if r.End == "" {
		r.End = "-1"
	}

	rf.start, err = strconv.Atoi(r.Start)
	if err != nil {
		return errors.New("Cannot convert start value to integer: " + err.Error())
	}

	rf.end, err = strconv.Atoi(r.End)
	if err != nil {
		return errors.New("Cannot convert end value to integer: " + err.Error())
	}

	r.Match = rf

	return
}
