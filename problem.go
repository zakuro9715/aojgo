package aoj

import (
	"encoding/xml"
	"github.com/zakuro9715/aojgo/util/http"
	"net/url"
	"strconv"
	"strings"
)

type Problem struct {
	Id                 string
	Name               string
	Available          int
	ProblemTimeLimit   int
	ProblemMemoryLimit int
	Status             Status
	SolvedList         []SolvedUser
}

type SolvedUser struct {
	Id             string
	SubmissionDate int64
	Language       string
	CpuTime        int
	Memory         int
	CodeSize       int
}

const (
	ProblemSearchUrl     = "http://judge.u-aizu.ac.jp/onlinejudge/webservice/problem"
	ProblemListSearchUrl = "http://judge.u-aizu.ac.jp/onlinejudge/webservice/problem_list"
)

func (p *Problem) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	pp := struct {
		Id                 string       `xml:"id"`
		Name               string       `xml:"name"`
		Available          string       `xml:"available"`
		ProblemTimeLimit   string       `xml:"problemtimelimit"`
		ProblemMemoryLimit string       `xml:"problemmemorylimit"`
		Status             Status       `xml:"status"`
		SolvedList         []SolvedUser `xml:"solved_list>user"`
	}{}
	if err := d.DecodeElement(&pp, &start); err != nil {
		return err
	}

	available, err := strconv.Atoi(strings.Trim(pp.Available, "\n"))
	problemTimeLimit, err := strconv.Atoi(strings.Trim(pp.ProblemTimeLimit, "\n"))
	problemMemoryLimit, err := strconv.Atoi(strings.Trim(pp.ProblemMemoryLimit, "\n"))
	if err != nil {
		return err
	}

	*p = Problem{
		Id:                 strings.Trim(pp.Id, "\n"),
		Name:               strings.Trim(pp.Name, "\n"),
		Available:          available,
		ProblemTimeLimit:   problemTimeLimit,
		ProblemMemoryLimit: problemMemoryLimit,
		Status:             pp.Status,
		SolvedList:         pp.SolvedList,
	}
	return nil
}

func (s *SolvedUser) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	ss := struct {
		Id             string `xml:"id"`
		SubmissionDate string `xml:"submissiondate"`
		Language       string `xml:"language"`
		CpuTime        string `xml:"cputime"`
		Memory         string `xml:"memory"`
		CodeSize       string `xml:"code_size"`
	}{}

	if err := d.DecodeElement(&ss, &start); err != nil {
		return err
	}

	submissionDate, _ := strconv.ParseInt(strings.Trim(ss.SubmissionDate, "\n "), 10, 64)
	cpuTime, _ := strconv.Atoi(strings.Trim(ss.CpuTime, "\n "))
	memory, _ := strconv.Atoi(strings.Trim(ss.Memory, "\n "))
	codeSize, _ := strconv.Atoi(strings.Trim(ss.CodeSize, "\n "))
	*s = SolvedUser{
		Id:             strings.Trim(ss.Id, "\n "),
		SubmissionDate: submissionDate,
		Language:       strings.Trim(ss.Language, "\n "),
		CpuTime:        cpuTime,
		Memory:         memory,
		CodeSize:       codeSize,
	}
	return nil
}

func ProblemSearchApi(id string, status bool) (*Problem, error) {
	q := url.Values{"id": {id}, "status": {strconv.FormatBool(status)}}.Encode()
	data, err := http.Get(ProblemSearchUrl, q)
	if err != nil {
		return nil, err
	}
	p := new(Problem)
	err = xml.Unmarshal(data, p)
	return p, err
}

func ProblemListSearchApi(volume int) ([]Problem, error) {
	q := url.Values{"volume": {strconv.Itoa(volume)}}.Encode()
	data, err := http.Get(ProblemListSearchUrl, q)
	if err != nil {
		return nil, err
	}
	pl := &struct {
		Problems []Problem `xml:"problem"`
	}{}
	err = xml.Unmarshal(data, pl)
	return pl.Problems, err
}
