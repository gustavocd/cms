package server

import (
	"errors"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/gorilla/mux"
	"github.com/gustavocd/cms/models"
)

var (
	// ErrInvalidID occurs when the given id is not valid or can not be parse to int64
	ErrInvalidID = errors.New("El id es inválido")
	// ErrNotFound occurs when there is no related data to the query
	ErrNotFound = errors.New("No se encontraron resultados")
	// ErrDeletePage occurs when the delete action fails
	ErrDeletePage = errors.New("No se pudo eliminar la página")
	// ErrCreatePage occurs when the create action fails
	ErrCreatePage = errors.New("No se pudo crear la página")
	// ErrUpdatePage occurs when the update action fails
	ErrUpdatePage = errors.New("No se pudo actualizar la página")
	// ErrDecodePage occurs when the given data can not be decoded
	ErrDecodePage = errors.New("No es posible leer la información enviada")
)

// HandlePagesCreate handles page creation
func (s *Server) HandlePagesCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		page := &models.Page{}
		err := s.decode(w, r, page)
		if err != nil {
			errs := validation.Errors{}
			errs["message"] = ErrDecodePage
			s.respondWithErr(w, r, errs, http.StatusBadRequest)
			return
		}

		err = page.Validate()
		if err != nil {
			s.respondWithErr(w, r, err, http.StatusBadRequest)
			return
		}

		err = s.db.Create(page)
		if err != nil {
			errs := validation.Errors{}
			errs["message"] = ErrCreatePage
			s.respondWithErr(w, r, errs, http.StatusInternalServerError)
			return
		}

		s.respond(w, r, "Página creada exitosamente", http.StatusOK)
	}
}

// HandlePagesGetAll gets all the pages from the database
func (s *Server) HandlePagesGetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pages := []models.Page{}
		err := s.db.All(&pages)
		if err != nil {
			errs := validation.Errors{}
			errs["message"] = ErrNotFound
			s.respondWithErr(w, r, errs, http.StatusNotFound)
			return
		}

		s.respond(w, r, pages, http.StatusOK)
	}
}

// HandlePagesGet gets a single page based on it's id
func (s *Server) HandlePagesGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		page := models.Page{}
		err := s.db.Find(&page, id)
		if err != nil {
			errs := validation.Errors{}
			errs["message"] = ErrNotFound
			s.respondWithErr(w, r, errs, http.StatusNotFound)
			return
		}

		s.respond(w, r, page, http.StatusFound)
	}
}

// HandlePagesDelete deletes a page based on it's id
func (s *Server) HandlePagesDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := s.parseID(r)
		if err != nil {
			errs := validation.Errors{}
			errs["message"] = ErrInvalidID
			s.respondWithErr(w, r, errs, http.StatusBadRequest)
			return
		}
		page := &models.Page{ID: id}
		err = s.db.Destroy(page)
		if err != nil {
			errs := validation.Errors{}
			errs["message"] = ErrDeletePage
			s.respondWithErr(w, r, errs, http.StatusNotFound)
			return
		}

		s.respond(w, r, "Página eliminada exitosamente", http.StatusOK)
	}
}

// HandlePagesUpdate updates a page based on it's id
func (s *Server) HandlePagesUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := s.parseID(r)
		if err != nil {
			errs := validation.Errors{}
			errs["message"] = ErrInvalidID
			s.respondWithErr(w, r, errs, http.StatusBadRequest)
			return
		}
		page := &models.Page{ID: id}
		err = s.decode(w, r, page)
		if err != nil {
			errs := validation.Errors{}
			errs["message"] = ErrDecodePage
			s.respondWithErr(w, r, errs, http.StatusBadRequest)
			return
		}
		err = page.Validate()
		if err != nil {
			s.respondWithErr(w, r, err, http.StatusBadRequest)
			return
		}
		err = s.db.Update(page)
		if err != nil {
			errs := validation.Errors{}
			errs["message"] = ErrUpdatePage
			s.respondWithErr(w, r, errs, http.StatusInternalServerError)
			return
		}

		s.respond(w, r, "Página actualizada exitosamente", http.StatusOK)
	}
}
