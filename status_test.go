package aoj
import (
  "encoding/xml"
  "testing"
)


func TestStatusUnmarshal1(t *testing.T) {
  s := new(Status)
  err := xml.Unmarshal([]byte(statusTestFixture1), s)
  if err != nil {
    t.Errorf("Faild to unmarshal: %v", err)
  }
  if s.Submission != 17576 {
    t.Errorf("Submission is %v, want 17575", s.Submission)
  }
  if s.Accepted != 7020 {
    t.Errorf("Accepted is %v, want 7020", s.Accepted)
  }
  if s.WrongAnswer != 6993 {
    t.Errorf("WrongAnswer is %v, want 6993", s.WrongAnswer)
  }
  if s.TimeLimit != 1414 {
    t.Errorf("TimeLimit is %v, want 1414", s.TimeLimit)
  }
  if s.MemoryLimit != 15 {
    t.Errorf("MemoryLimit is %v, want 15", s.MemoryLimit)
  }
  if s.OutputLimit != 0 {
    t.Errorf("OutputLimit is %v, want 0", s.OutputLimit)
  }
  if s.RuntimeError != 2110 {
    t.Errorf("RuntimeError is %v, want 2110", s.RuntimeError)
  }
}

func TestStatusUnmarshal2(t *testing.T) {
  s := new(Status)
  err := xml.Unmarshal([]byte(statusTestFixture2), s)
  if err != nil {
    t.Errorf("Faild to unmarshal: %v", err)
  }
  if s.Submission != 2589 {
    t.Errorf("Submission is %v, want 2589", s.Submission)
  }
  if s.Solved != 716 {
    t.Errorf("Solved is %v, want 716", s.Solved)
  }
  if s.Accepted != 864 {
    t.Errorf("Accepted is %v, want 864", s.Accepted)
  }
  if s.WrongAnswer != 1095 {
    t.Errorf("WrongAnswer is %v, want 1095", s.WrongAnswer)
  }
  if s.TimeLimit != 338 {
    t.Errorf("TimeLimit is %v, want 338", s.TimeLimit)
  }
  if s.MemoryLimit != 34 {
    t.Errorf("MemoryLimit is %v, want 34", s.MemoryLimit)
  }
  if s.OutputLimit != 0 {
    t.Errorf("OutputLimit is %v, want 0", s.OutputLimit)
  }
  if s.RuntimeError != 181 {
    t.Errorf("RuntimeError is %v, want 181", s.RuntimeError)
  }
  if s.CompileError != 76 {
    t.Errorf("CompileError is %v, want 76", s.CompileError)
  }
}



// http://judge.u-aizu.ac.jp/onlinejudge/webservice/problem?id=0002
const statusTestFixture1 = `
<Status>
<submission>
17576
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
</Status>`

const statusTestFixture2 = `
<status>
<submission>2589</submission>
<solved>716</solved>
<accepted>864</accepted>
<wronganswer>1095</wronganswer>
<timelimit>338</timelimit>
<memorylimit>34</memorylimit>
<outputlimit>0</outputlimit>
<runtimeerror>181</runtimeerror>
<compileerror>76</compileerror>
</status>`
