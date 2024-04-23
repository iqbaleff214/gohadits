package hadith

// Service contains all business logic method
type Service interface {
	GetAllAvailableBooks() []Book
}

// service as a class
type service struct {
	repository Repository
}

// NewService is a function to instantiate new service object
func NewService(repository Repository) Service {
	return service{repository}
}

// GetAllAvailableBooks to retrieve all hadith books
func (s service) GetAllAvailableBooks() []Book {
	return s.repository.GetAllBook()
}
