// yourproject/internal/application/user.go

package application

import (
	"github.com/domjeff/deploy-to-railway/internal/application/request"
	"github.com/domjeff/deploy-to-railway/internal/application/response"
	"github.com/domjeff/deploy-to-railway/internal/domain"
)

type UserService struct {
	UserRepository domain.UserRepository
}

func NewUserService(repo domain.UserRepository) *UserService {
	return &UserService{UserRepository: repo}
}

func (s *UserService) CreateUser(user *domain.User) error {
	// TODO: check if email is valid

	if err := s.UserRepository.Save(user); err != nil {
		return err
	}

	return nil
}

func (s *UserService) CreateGormStudent(student *domain.Student) error {
	// TODO: check if email is valid

	if err := s.UserRepository.SaveGormStudent(*student); err != nil {
		return err
	}

	return nil
}

func (s *UserService) GetGormStudent(id int) (domain.Student, error) {
	student, err := s.UserRepository.GetGormStudent(id)
	if err != nil {
		return domain.Student{}, err
	}

	return student, nil
}

func (s *UserService) UpdateGormStudent(id int, req request.UpdateGormStudent) error {
	err := s.UserRepository.UpdateGormStudent(domain.Student{
		ID:    uint(id),
		Email: req.Email,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) UpdateUser(user *domain.User) error {
	if err := s.UserRepository.Save(user); err != nil {
		return err
	}

	return nil
}

func (s *UserService) GetUsers() (response.GetUsers, error) {
	res, err := s.UserRepository.GetUsers()
	if err != nil {
		return response.GetUsers{}, err
	}

	var out response.GetUsers

	var users []response.GetUsersUser
	for _, user := range res {
		users = append(users, response.GetUsersUser{
			Id:       user.ID,
			Email:    user.Email,
			Username: user.UserName,
		})
	}

	out.Users = users

	return out, nil
}

// GetUserByID retrieves a user by ID.
func (s *UserService) GetUserByID(id int) (*domain.User, error) {
	// Retrieve the user from the database by ID
	user, err := s.UserRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
