package hadith

// Book model to contain hadith book
type Book struct {
	Slug string `json:"book"`
	Name string `json:"name"`
	Size int    `json:"total"`
}

// Hadith model to contain hadith text
type Hadith struct {
	Number    int    `json:"no"`
	Text      string `json:"ar"`
	Translate string `json:"id"`
}

// HadithCollextion an alias to collection of hadith from all hadith book
type HadithCollection map[string][]Hadith
