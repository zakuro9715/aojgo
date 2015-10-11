package aoj

import (
	"encoding/xml"
	"github.com/zakuro9715/aojgo/util/http"
	"net/url"
)

const (
	UserSearchUrl = "http://judge.u-aizu.ac.jp/onlinejudge/webservice/user"
)

type User struct {
	Id             string          `xml:"id"`
	Name           string          `xml:"name"`
	Affliation     string          `xml:"affliation"`
	RegisterDate   int64           `xml:"registerdate"`
	LastSubmitDate int64           `xml:"lastsubmitdate"`
	Status         Status          `xml:"status"`
	SolvedList     []SolvedProblem `xml:"solved_list>problem"`
}
type SolvedProblem struct {
	Id             string `xml:"id"`
	JudgeId        string `xml:"judge_id"`
	SubmissionDate int64  `xml:"submissiondate"`
	Language       string `xml:"language"`
	CpuTime        int    `xml:"cputime"`
	Memory         int    `xml:"memory"`
	CodeSize       int    `xml:"code_size"`
}

func UserSearchApi(userId string) (*User, error) {
	data, err := http.Get(UserSearchUrl, url.Values{"id": {userId}}.Encode())
	if err != nil {
		return nil, err
	}

	u := new(User)
	err = xml.Unmarshal(data, u)
	return u, err
}
