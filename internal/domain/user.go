package domain

type User struct {
	ID       int
	UserName string
	Email    string
	// Add other user fields here
}

type UserRepository interface {
	Save(user *User) error
	GetUsers() (users []User, err error)
	FindByID(id int) (*User, error)

	SaveGormStudent(gormStudent Student) error
	GetGormStudent(id int) (Student, error)

	UpdateGormStudent(gormStudent Student) error
}
