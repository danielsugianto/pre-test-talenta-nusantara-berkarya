package database

import (
	"errors"
	"strconv"

	"github.com/danielsugianto/pre-test/test-talenta-nusantara-berkarya/models"
)

var languages []models.Language

func init() {
	// languages = append(languages, models.Language{ID: 1, Name: "language 1", Year: 2010})
	// languages = append(languages, models.Language{ID: 2, Name: "language 2", Year: 2011})
	// languages = append(languages, models.Language{ID: 3, Name: "language 3", Year: 2012})
	// languages = append(languages, models.Language{ID: 4, Name: "language 4", Year: 2013})
	// languages = append(languages, models.Language{ID: 5, Name: "language 5", Year: 2014})
}

// get language hard coded
func GetLanguageHardCoded() (models.Language, error) {
	language := models.Language{
		Language: "C",
		Appeared: 1972,
		Created: []string{
			"Dennis Ritchie",
		},
		Functional:     true,
		ObjectOriented: false,
		Relation: models.Relation{
			InfluencedBy: []string{
				"B",
				"ALGOL 68",
				"Assembly",
				"FORTRAN",
			},
			Influences: []string{
				"C++",
				"Objective-C",
				"C#",
				"Java",
				"JavaScript",
				"PHP",
				"Go",
			},
		},
	}
	return language, nil
}

// get all languages
func GetLanguages() ([]models.Language, error) {
	return languages, nil
}

// create new language
func CreateLanguage(language models.Language) (models.Language, error) {
	languages = append(languages, language)
	return language, nil
}

// get language by ID
func GetLanguage(id string) (models.Language, error) {
	idInt, _ := strconv.Atoi(id)
	language := languages[idInt]
	return language, nil
}

// delete language by ID
func DeleteLanguage(id string) error {
	idInt, _ := strconv.Atoi(id)

	if idInt >= len(languages) {
		return errors.New("language not found")
	}

	languages[idInt] = languages[len(languages)-1]
	languages[len(languages)-1] = models.Language{}
	languages = languages[:len(languages)-1]

	return nil
}

// update language by ID
func UpdateLanguage(id string, LanguageParam models.Language) (models.Language, error) {
	idInt, _ := strconv.Atoi(id)

	if idInt >= len(languages) {
		return models.Language{}, errors.New("language not found")
	}

	languages[idInt] = LanguageParam

	return languages[idInt], nil
}
