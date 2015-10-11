package aoj

import (
	"encoding/xml"
	"reflect"
	"testing"
)

func TestUserUnmarshalXML(t *testing.T) {
	u := new(User)
	xml.Unmarshal([]byte(userFixture), u)

	want := &User{
		Id:             "zakuro9715",
		Name:           "Takahiro Yaota",
		Affliation:     "",
		RegisterDate:   1354274836353,
		LastSubmitDate: 1443322921511,
		Status: Status{
			Submission:   873,
			Solved:       281,
			Accepted:     301,
			WrongAnswer:  326,
			TimeLimit:    59,
			MemoryLimit:  13,
			OutputLimit:  0,
			RuntimeError: 95,
			CompileError: 79,
		},
		SolvedList: []SolvedProblem{
			{
				JudgeId:        "543436",
				Id:             "0000",
				SubmissionDate: 1355296785075,
				Language:       "C",
				CpuTime:        0,
				Memory:         524,
				CodeSize:       158,
			},
		},
	}
	if !reflect.DeepEqual(u, want) {
		t.Errorf("got %v, want %v", u, want)
	}
}

const userFixture = `
<?xml version="1.0"?>
<user>
<id>zakuro9715</id>
<name>Takahiro Yaota</name>
<affiliation></affiliation>
<registerdate>1354274836353</registerdate>
<lastsubmitdate>1443322921511</lastsubmitdate>
<status>
<submission>873</submission>
<solved>281</solved>
<accepted>301</accepted>
<wronganswer>326</wronganswer>
<timelimit>59</timelimit>
<memorylimit>13</memorylimit>
<outputlimit>0</outputlimit>
<runtimeerror>95</runtimeerror>
<compileerror>79</compileerror>
</status>
<solved_list>
<problem>
<judge_id>543436</judge_id>
<id>0000</id>
<submissiondate>1355296785075</submissiondate>
<language>C</language>
<cputime>0</cputime>
<memory>524</memory>
<code_size>158</code_size>
</problem>
<problem>
</solved_list>
<problem_score>
<score serial="0" category="straight">31.162306</score>
<score serial="1" category="datamanipu">27.328575</score>
<score serial="2" category="string">8.213898</score>
<score serial="3" category="parsing">0.50938934</score>
<score serial="4" category="simulation">2.092038</score>
<score serial="5" category="puzzle">1.3231473</score>
<score serial="6" category="geometry">0.45898613</score>
<score serial="7" category="numerical">0.1807341</score>
<score serial="8" category="math">0.54813427</score>
<score serial="9" category="number">5.994915</score>
<score serial="10" category="graph">1.9430909</score>
<score serial="11" category="game">0.0</score>
<score serial="12" category="combinatorial">6.152039</score>
<score serial="13" category="probability">1.2656037</score>
</problem_score>
</user>
`
