package hadith

import (
	"embed"
	"encoding/json"
	"io/fs"
	"path"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// Repository contains method to interact with data source
type Repository interface {
	GetAllBook() []Book
	GetAllHadith() HadithCollection
}

// repository as a class
type repository struct {
	books   []Book
	hadiths HadithCollection
}

// NewRepository is a function to instantiate new repository object
//
//	this function also retrieve all content of json files that containing collection of hadith.
//	retrieved contents stored to repository object properties and used as the data source.
func NewRepository(data embed.FS) Repository {
	hadiths := map[string][]Hadith{}
	books := []Book{}

	if err := fs.WalkDir(data, "data/tafsirq", func(filename string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}

		b, err := data.ReadFile(filename)
		if err != nil {
			return err
		}

		collection := []Hadith{}

		if err := json.Unmarshal(b, &collection); err != nil {
			return err
		}

		filename = path.Base(filename)
		bookName := filename[:len(filename)-5]

		hadiths[bookName] = collection

		books = append(books, Book{
			Slug: bookName,
			Size: len(collection),
			Name: cases.Title(language.Indonesian).String(strings.ReplaceAll(bookName, "-", " ")),
		})

		return nil
	}); err != nil {
		panic(err)
	}

	return repository{
		books:   books,
		hadiths: hadiths,
	}
}

// GetAllBook to retrieve all hadith books
func (r repository) GetAllBook() []Book {
	return r.books
}

// GetAllHadith to retrieve all hadith collection
func (r repository) GetAllHadith() HadithCollection {
	return r.hadiths
}
