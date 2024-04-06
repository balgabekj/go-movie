package main

import (
	"database/sql"
	"encoding/json"
	"github.com/balgabekj/go_movie/pkg/model"
	"github.com/gorilla/mux"
	"net/http"
)

// Create a handler to create a movie
func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	// Извлекаем данные из тела запроса
	var movie model.Movie
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		app.models.Movies.ErrorLog.Println(err) // Здесь изменение
		http.Error(w, "Ошибка при декодировании данных", http.StatusBadRequest)
		return
	}

	// Вызываем метод модели для добавления фильма в базу данных
	err = app.models.Movies.Insert(&movie)
	if err != nil {
		app.models.Movies.ErrorLog.Println(err) // Здесь изменение
		http.Error(w, "Ошибка при создании фильма", http.StatusInternalServerError)
		return
	}

	// Возвращаем успешный ответ
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(movie)
}

// Create a handler to get a movie
func (app *application) getMovieHandler(w http.ResponseWriter, r *http.Request) {
	// Извлекаем параметр из URL запроса
	params := mux.Vars(r)
	id := params["id"]

	// Вызываем метод модели для получения фильма по его ID
	movie, err := app.models.Movies.Get(id)
	if err != nil {
		app.models.Movies.ErrorLog.Println(err) // Здесь изменение
		if err == sql.ErrNoRows {
			http.Error(w, "Фильм не найден", http.StatusNotFound)
		} else {
			http.Error(w, "Ошибка при получении фильма", http.StatusInternalServerError)
		}
		return
	}

	// Возвращаем найденный фильм
	json.NewEncoder(w).Encode(movie)
}

// Create a handler to update a movie
func (app *application) updateMovieHandler(w http.ResponseWriter, r *http.Request) {
	// Извлекаем параметр из URL запроса
	params := mux.Vars(r)
	id := params["id"]

	// Извлекаем данные из тела запроса
	var movie model.Movie
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		app.models.Movies.ErrorLog.Println(err) // Здесь изменение
		http.Error(w, "Ошибка при декодировании данных", http.StatusBadRequest)
		return
	}

	// Устанавливаем ID фильма для обновления
	movie.ID = id

	// Вызываем метод модели для обновления фильма в базе данных
	err = app.models.Movies.Update(&movie)
	if err != nil {
		app.models.Movies.ErrorLog.Println(err) // Здесь изменение
		http.Error(w, "Ошибка при обновлении фильма", http.StatusInternalServerError)
		return
	}

	// Возвращаем успешный ответ
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(movie)
}

// Create a handler to delete a movie
func (app *application) deleteMovieHandler(w http.ResponseWriter, r *http.Request) {
	// Извлекаем параметр из URL запроса
	params := mux.Vars(r)
	id := params["id"]

	// Вызываем метод модели для удаления фильма из базы данных
	err := app.models.Movies.Delete(id)
	if err != nil {
		app.models.Movies.ErrorLog.Println(err) // Здесь изменение
		http.Error(w, "Ошибка при удалении фильма", http.StatusInternalServerError)
		return
	}

	// Возвращаем успешный ответ
	w.WriteHeader(http.StatusNoContent)
}
