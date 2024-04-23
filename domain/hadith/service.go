package hadith

import (
	"errors"
	"strconv"
)

// Service contains all business logic method
type Service interface {
	GetAllAvailableBooks() []Book
	GetHadithByBook(bookName string, offset, limit int) (Book, []Hadith, error)
	GetHadithByNumber(bookName string, number int) (Book, Hadith, error)
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

// GetHadithByBook to retrieve collection of hadith from the hadith book
func (s service) GetHadithByBook(bookName string, offset, limit int) (Book, []Hadith, error) {
	var book Book

	// Check if book slug is correct and available
	for _, b := range s.repository.GetAllBook() {
		if b.Slug == bookName {
			book = b
		}
	}
	if book.Size == 0 {
		return book, nil, errors.New(bookName + " is not available")
	}

	if offset < 0 {
		return book, nil, errors.New("offset should not less than zero")
	}

	if offset+1 > book.Size {
		return book, nil, errors.New("offset should not larger than total hadith available")
	}

	if limit <= 0 {
		return book, nil, errors.New("limit should not equal zero or less")
	}

	if limit > MAX_LIMIT_PER_PAGE {
		limit = MAX_LIMIT_PER_PAGE
	}

	if limit+offset > book.Size {
		limit = book.Size - offset
	}

	// Retrieve all hadith
	allHadith := s.repository.GetAllHadith()
	hadiths, ok := allHadith[book.Slug]
	if !ok {
		return book, nil, errors.New(bookName + " is not available")
	}
	hadiths = hadiths[offset : offset+limit]

	return book, hadiths, nil
}

// GetHadithByNumber to retrieve a specific hadith
func (s service) GetHadithByNumber(bookName string, number int) (Book, Hadith, error) {
	var book Book
	var hadith Hadith

	if number <= 0 {
		return book, hadith, errors.New("hadith number should not less than 1")
	}

	// Check if book slug is correct and available
	for _, b := range s.repository.GetAllBook() {
		if b.Slug == bookName {
			book = b
		}
	}
	if book.Size == 0 {
		return book, hadith, errors.New(bookName + " is not available")
	}

	// Retrieve all hadith
	allHadith := s.repository.GetAllHadith()
	hadiths, ok := allHadith[book.Slug]
	if !ok {
		return book, hadith, errors.New(bookName + " is not available")
	}

	// Search hadith by number using binary search
	index, low, high := -1, 0, book.Size-1
	for low <= high {
		mid := low + (high-low)/2

		if hadiths[mid].Number == number {
			index = mid
			break
		} else if hadiths[mid].Number < number {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	// index is less than zero means hadith not found
	if index < 0 {
		return book, hadith, errors.New("hadith no. " + strconv.Itoa(number) + " from book of " + book.Name + " is not found")
	}

	hadith = hadiths[index]

	return book, hadith, nil
}
