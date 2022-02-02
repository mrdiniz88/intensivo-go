package entity

type CourseRepository interface {
	Insert(Course) error
}