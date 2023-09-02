package queries

import (
	"fiber-postgres-api/initializer"
	"fiber-postgres-api/models"
)

func GetBooksQuery() ([]models.Books, error) {
	bookModels := &[]models.Books{}

	DB, err := initializer.Connect()
	if err != nil {
		return nil, err
	}

	err = DB.Find(bookModels).Error
	if err != nil {
		return nil, err
	}
	return *bookModels, nil
}

func GetBookByIDQuery(id string) (*models.Books, error) {
	bookModel := &models.Books{}

	DB, err := initializer.Connect()
	if err != nil {
		return nil, err
	}

	err = DB.Where("id = ?", id).First(bookModel).Error
	if err != nil {
		return nil, err
	}

	return bookModel, nil
}

func CreateBookQuery(book *models.Books) error {
	DB, err := initializer.Connect()
	if err != nil {
		return err
	}

	err = DB.Create(book).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteBookQuery(id string) error {
	bookModel := models.Books{}

	DB, err := initializer.Connect()
	if err != nil {
		return err
	}

	res := DB.Delete(bookModel, id)

	if res.Error != nil {
		return res.Error
	}

	return nil
}

func UpdateBookQuery(updateBook map[string]interface{}, id string) error {
	DB, err := initializer.Connect()
	if err != nil {
		return err
	}

	// Find the book we're updating
	book := &models.Books{}
	DB.First(&book, id)

	// Update the book's information
	DB.Model(&book).Updates(updateBook)

	return nil
}