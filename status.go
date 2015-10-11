package aoj

import (
	"encoding/xml"
	"strconv"
	"strings"
)

type Status struct {
	Submission   int `xml:"submission"`
	Solved       int `xml:"solved"`
	Accepted     int `xml:"accepted"`
	WrongAnswer  int `xml:"wronganswer"`
	TimeLimit    int `xml:"timelimit"`
	MemoryLimit  int `xml:"memorylimt"`
	OutputLimit  int `xml:"outputlimit"`
	RuntimeError int `xml:"runtimeerror"`
	CompileError int `xml:"compileerror"`
}

func (s *Status) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	ss := struct {
		Submission   string `xml:"submission"`
		Solved       string `xml:"solved"`
		Accepted     string `xml:"accepted"`
		WrongAnswer  string `xml:"wronganswer"`
		TimeLimit    string `xml:"timelimit"`
		MemoryLimit  string `xml:"memorylimit"`
		OutputLimit  string `xml:"outputlimit"`
		RuntimeError string `xml:"runtimeerror"`
		CompileError string `xml:"compileerror"`
	}{}
	if err := d.DecodeElement(&ss, &start); err != nil {
		return err
	}
	submission, _ := strconv.Atoi(strings.Trim(ss.Submission, "\n "))
	solved, _ := strconv.Atoi(strings.Trim(ss.Solved, "\n "))
	accepted, _ := strconv.Atoi(strings.Trim(ss.Accepted, "\n "))
	wrongAnswer, _ := strconv.Atoi(strings.Trim(ss.WrongAnswer, "\n "))
	timeLimit, _ := strconv.Atoi(strings.Trim(ss.TimeLimit, "\n "))
	memoryLimit, _ := strconv.Atoi(strings.Trim(ss.MemoryLimit, "\n "))
	outputLimit, _ := strconv.Atoi(strings.Trim(ss.OutputLimit, "\n "))
	runtimeError, _ := strconv.Atoi(strings.Trim(ss.RuntimeError, "\n "))
	compileError, _ := strconv.Atoi(strings.Trim(ss.CompileError, "\n "))

	*s = Status{
		Submission:   submission,
		Solved:       solved,
		Accepted:     accepted,
		WrongAnswer:  wrongAnswer,
		TimeLimit:    timeLimit,
		MemoryLimit:  memoryLimit,
		OutputLimit:  outputLimit,
		RuntimeError: runtimeError,
		CompileError: compileError,
	}
	return nil
}
