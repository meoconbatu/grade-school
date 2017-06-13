package school

import (
	"sort"
)

const testVersion = 1

type Grade struct {
	grade    int
	students []string
}
type School []Grade

func New() *School {
	s := make(School, 0)
	return &s
}
func (s *School) Add(student string, grade int) {
	students := s.Grade(grade)
	if students == nil {
		students = append(students, student)
		g := Grade{grade: grade, students: students}
		*s = append(*s, g)
	} else {
		for i := 0; i < len(*s); i++ {
			if (*s)[i].grade == grade {
				// ss := append((*s)[i].students, student)
				// sort.Strings(ss)
				(*s)[i].students = append((*s)[i].students, student)
			}
		}
	}

}
func (s *School) Grade(grade int) []string {
	for _, g := range *s {
		if g.grade == grade {
			return g.students
		}
	}
	return nil
}
func (s *School) Enrollment() []Grade {
	sort.Sort(*s)
	for i := 0; i < len(*s); i++ {
		sort.Strings((*s)[i].students)
	}
	return *s
}
func (s School) Len() int {
	return len(s)
}
func (s School) Less(i, j int) bool {
	return s[i].grade < s[j].grade
}
func (s School) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
