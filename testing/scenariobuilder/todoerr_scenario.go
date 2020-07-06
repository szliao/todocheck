package scenariobuilder

import "fmt"

// TodoErrType is an enum specifying the possible todo err type one could expect
type TodoErrType string

// Possible types of todo errors to expect
const (
	TodoErrTypeMalformed        TodoErrType = "ERROR: Malformed todo."
	TodoErrTypeIssueClosed                  = "ERROR: Issue is closed."
	TodoErrTypeIssueNonExistent             = "ERROR: Issue doesn't exist."
)

// TodoErrScenario encapsulates a test scenario for an expected todo err the program should return.
type TodoErrScenario struct {
	TodoErrType
	sourceFile    string
	sourceLineNum int
	contents      []string
}

// NewTodoErr returns a new todo err scenario
func NewTodoErr() *TodoErrScenario {
	return &TodoErrScenario{}
}

// WithType specifies the expected todo err type for the given scenario
func (s *TodoErrScenario) WithType(t TodoErrType) *TodoErrScenario {
	s.TodoErrType = t
	return s
}

// WithSourceFile specifies the expected source file for the given todo err scenario
func (s *TodoErrScenario) WithSourceFile(sf string) *TodoErrScenario {
	s.sourceFile = sf
	return s
}

// WithLineNum specifies the expected starting line number for the given todo err scenario
func (s *TodoErrScenario) WithLineNum(n int) *TodoErrScenario {
	s.sourceLineNum = n
	return s
}

// ExpectLine appends an expected error line to the todo err scenario. more than one line can be specified.
func (s *TodoErrScenario) ExpectLine(line string) *TodoErrScenario {
	s.contents = append(s.contents, line)
	return s
}

func (s *TodoErrScenario) String() string {
	str := fmt.Sprintf("%s\n", s.TodoErrType)
	for i := 0; i < len(s.contents); i++ {
		str += fmt.Sprintf("%s:%d: %s\n", s.sourceFile, i+s.sourceLineNum, s.contents[i])
	}

	if s.TodoErrType == TodoErrTypeMalformed {
		str += "\t> TODO should match pattern - \"TODO [TASK_ID]:\""
	}

	return str
}