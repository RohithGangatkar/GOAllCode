package anonymousFilter

type Student struct {
	FirstName string
	LastName  string
	Grade     string
	Country   string
}

func Filter(students []Student, f func(s Student) bool) []Student {
	var studentFiltered []Student
	for _, eachStu := range students {
		if f(eachStu) {
			studentFiltered = append(studentFiltered, eachStu)
		}
	}

	return studentFiltered
}
