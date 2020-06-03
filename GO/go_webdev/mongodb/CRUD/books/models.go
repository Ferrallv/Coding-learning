package books

import (
	"errors"
	"ExercisesAndProjectsInC/GO/go_webdev/mongodb/CRUD/config"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"strconv"
	"context"
)

type Book struct {
	Isbn string
	Title string
	Author string
	Price float32 `truncate`
}

func AllBooks() ([]Book, error) {
	bks := []Book{}
	cursor, err := config.Books.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}

	if err = cursor.All(context.TODO(), &bks); err != nil {
		return nil, err
	}

	return bks, nil
}

func OneBook(r *http.Request) (Book, error) {
	bk := Book{}
	isbn := r.FormValue("isbn")
	if isbn == "" {
		return bk, errors.New("400. Bad Request.")
	}
	err := config.Books.FindOne(context.TODO(), bson.M{"isbn": isbn}).Decode(&bk)
	if err != nil {
		return bk, err
	}
	return bk, nil
}

func PutBook(req *http.Request) (Book, error) {
	bk := Book{}
	bk.Isbn = req.FormValue("isbn")
	bk.Title = req.FormValue("title")
	bk.Author = req.FormValue("Author")
	p := req.FormValue("price")

	if bk.Isbn == "" || bk.Title == "" || bk.Author == "" || p == "" {
		return bk, errors.New("400. Bad request. All fields must be complete")
	}

	f64, err := strconv.ParseFloat(p, 32)
	if err != nil {
		return bk, errors.New("406. Not Acceptable. Price must be a number")
	}
	bk.Price = float32(f64)

	_, err = config.Books.InsertOne(context.TODO(), bk)
	if err != nil {
		return bk, errors.New("500. Internal Server Error." + err.Error())
	}
	return bk, nil
}

func UpdateBook(req *http.Request) (Book, error) {
	bk := Book{}
	bk.Isbn = req.FormValue("isbn")
	bk.Title = req.FormValue("title")
	bk.Author = req.FormValue("Author")
	p := req.FormValue("price")

	if bk.Isbn == "" || bk.Title == "" || bk.Author == "" || p == "" {
		return bk, errors.New("400. Bad request. All fields must be complete")
	}

	f64, err := strconv.ParseFloat(p, 32)
	if err != nil {
		return bk, errors.New("406. Not Acceptable. Price must be a number")
	}
	bk.Price = float32(f64)

	_, err = config.Books.UpdateOne(context.TODO(), bson.M{"isbn": bk.Isbn}, &bk)
	if err != nil {
		return bk, err
	}
	return bk, nil
}

func DeleteBook(req *http.Request) error {
	isbn := req.FormValue("isbn")
	if isbn == "" {
		return errors.New("400. Bad Request.")
	}

	_, err := config.Books.DeleteOne(context.TODO(), bson.M{"isbn": isbn})
	if err != nil {
		return errors.New("500. Internal Server Error")
	}
	return nil
}
