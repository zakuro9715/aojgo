package aoj

import (
	"encoding/xml"
	"fmt"
	"github.com/zakuro9715/aojgo/util/http"
	"net/url"
	"strconv"
	"strings"
)

type SolvedRecord struct {
	RunId     string `xml:"run_id"`
	UserId    string `xml:"user_id"`
	ProblemId string `xml:"problem_id"`
	Date      int64  `xml:"date"`
	Language  string `xml:"language"`
	CpuTime   int    `xml:"cputime"`
	Memory    int    `xml:"memory"`
	CodeSize  int    `xml:"code_size"`
}

const (
	SolvedRecordSearchUrl = "http://judge.u-aizu.ac.jp/onlinejudge/webservice/solved_record"
)

func (r *SolvedRecord) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	rr := struct {
		RunId     string `xml:"run_id"`
		UserId    string `xml:"user_id"`
		ProblemId string `xml:"problem_id"`
		Date      string `xml:"date"`
		Language  string `xml:"language"`
		CpuTime   string `xml:"cputime"`
		Memory    string `xml:"memory"`
		CodeSize  string `xml:"code_size"`
	}{}
	if err := d.DecodeElement(&rr, &start); err != nil {
		return err
	}

	date, err := strconv.ParseInt(strings.Trim(rr.Date, "\n"), 10, 0)
	cpuTime, err := strconv.Atoi(strings.Trim(rr.CpuTime, "\n"))
	memory, err := strconv.Atoi(strings.Trim(rr.Memory, "\n"))
	codeSize, err := strconv.Atoi(strings.Trim(rr.CodeSize, "\n"))
	if err != nil {
		return err
	}

	*r = SolvedRecord{
		RunId:     strings.Trim(rr.RunId, "\n"),
		UserId:    strings.Trim(rr.UserId, "\n"),
		ProblemId: strings.Trim(rr.ProblemId, "\n"),
		Date:      date,
		Language:  strings.Trim(rr.Language, "\n"),
		CpuTime:   cpuTime,
		Memory:    memory,
		CodeSize:  codeSize,
	}
	return nil
}

func SolvedRecordSearchApi(user_id, problem_id, language string, dateBegin, dateEnd int) ([]SolvedRecord, error) {
	q := url.Values{}
	if user_id != "" {
		q.Add("user_id", user_id)
	}
	if problem_id != "" {
		q.Add("problem_id", problem_id)
	}
	if language != "" {
		q.Add("language", language)
	}
	if dateBegin != -1 {
		q.Add("date_begin", strconv.Itoa(dateBegin))
	}
	if dateEnd != -1 {
		q.Add("date_end", strconv.Itoa(dateEnd))
	}
	data, err := http.Get(SolvedRecordSearchUrl, q.Encode())
	if err != nil {
		return nil, err
	}
	r := &struct {
		SolvedRecords []SolvedRecord `xml:"solved"`
	}{}
	fmt.Printf("%v\n", string(data))
	err = xml.Unmarshal(data, r)
	return r.SolvedRecords, err
}
