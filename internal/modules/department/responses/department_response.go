package responses

type Department struct {
	ID             uint
	Title          string
	DepartmentType string
	Parent         string
}

type Departments struct {
	Data []Department
}
