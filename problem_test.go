package aoj

import (
	"encoding/xml"
	"testing"
)

func TestProblemUnmarshal(t *testing.T) {
	p := new(Problem)
	err := xml.Unmarshal([]byte(problemFixture), p)
	if err != nil {
		t.Errorf("Faild to unmarshal: %v", err)
	}

	want := Problem{
		Id:                 "0000",
		Name:               "Name",
		Available:          1,
		ProblemTimeLimit:   1,
    ProblemMemoryLimit: 65536,
		Status: Status{
			Submission:   17575,
			Accepted:     7020,
			WrongAnswer:  6993,
			TimeLimit:    1414,
			MemoryLimit:  15,
			OutputLimit:  0,
			RuntimeError: 2110,
		},
		SolvedList: []SolvedUser{
			{
				Id:             "mango",
				SubmissionDate: 1380965484922,
				Language:       "JAVA",
				CpuTime:        6,
				Memory:         20416,
				CodeSize:       2197,
			},
		},
	}

	if p.Id != want.Id {
    t.Errorf("Id: got %v, want %v", p.Id, want.Id)
	}
	if p.Name != want.Name {
    t.Errorf("Name: got %v, want %v", p.Name, want.Name)
	}
	if p.Available != want.Available {
    t.Errorf("Available: got %v, want %v", p.Available, want.Available)
	}
	if p.ProblemTimeLimit != want.ProblemTimeLimit {
    t.Errorf("ProblemTimeLimit: got %v, want %v", p.ProblemTimeLimit, want.ProblemTimeLimit)
  }
	if p.ProblemMemoryLimit != want.ProblemMemoryLimit {
    t.Errorf("ProblemMemoryLimit: got %v, want %v", p.ProblemMemoryLimit, want.ProblemMemoryLimit)
	}
	if p.Status != want.Status {
    t.Errorf("Status: got %v, want %v", p.Status, want.Status)
	}
	if len(p.SolvedList) != len(want.SolvedList) {
    t.Errorf("len(SolvedList): got %v, want %v", len(p.SolvedList), len(want.SolvedList))
	}
	if p.SolvedList[0] != want.SolvedList[0] {
    t.Errorf("SolvedList[0]: got %v, want %v", p.SolvedList[0], want.SolvedList[0])
	}
}

const problemFixture = `
<?xml version="1.0"?>
<problem>
<id>
0000
</id>
<name>
Name
</name>
<available>
1
</available>
<problemtimelimit>
1
</problemtimelimit>
<problemmemorylimit>
65536
</problemmemorylimit>
<status>
<submission>
17575
</submission>
<accepted>
7020
</accepted>
<wronganswer>
6993
</wronganswer>
<timelimit>
1414
</timelimit>
<memorylimit>
15
</memorylimit>
<outputlimit>
0
</outputlimit>
<runtimeerror>
2110
</runtimeerror>
</status>
<solved_list>
<user>
<id>
mango
</id>
<submissiondate>
1380965484922
</submissiondate>
<language>
JAVA
</language>
<cputime>
6
</cputime>
<memory>
20416
</memory>
<code_size>
2197
</code_size>
</user>
</solved_list>
</problem>`
