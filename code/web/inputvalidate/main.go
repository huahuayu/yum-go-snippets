package main

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/asaskevich/govalidator"
)

// main - a simple main entry point
func main() {
	srv := &http.Server{
		Addr:    ":1234",
		Handler: f,
	}
	srv.ListenAndServe()
}

// f - our only application handler to demonstrate validation
var f http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
	// we have an input structure, of type Thing
	var thing *Thing
	// we want to decode and validate Thing from request body
	err := decodeAndValidate(r, thing)
	// there was an error with Thing
	if err != nil {
		// send a bad request back to the caller
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("bad request"))
		return
	}
	// it was decoded, and validated properly, success
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

// decodeAndValidate - entrypoint for deserialization and validation
// of the submission input
func decodeAndValidate(r *http.Request, v InputValidation) error {
	// json decode the payload - obviously this could be abstracted
	// to handle many content types
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return err
	}
	defer r.Body.Close()
	// peform validation on the InputValidation implementation
	return v.Validate(r)
}

// InputValidation - an interface for all "input submission" structs used for
// deserialization.  We pass in the request so that we can potentially get the
// context by the request from our context manager (future net.http will include
// a context in the request)
type InputValidation interface {
	Validate(r *http.Request) error
}

var (
	// ErrInvalidUUID - error when we have a UUID validation issue
	ErrInvalidUUID = errors.New("invalid uuid")
	// ErrInvalidName - error when we have an invalid name
	ErrInvalidName = errors.New("invalid name")
)

// Thing - is our implemenation of InputValidation and the structure we will
// deserialize into
type Thing struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
}

// Validate - implementation of the InputValidation interface
func (t Thing) Validate(r *http.Request) error {
	// validate the ID is a uuid
	if !govalidator.IsUUID(t.ID) {
		return ErrInvalidUUID
	}
	// validate the name is not empty or missing
	if govalidator.IsNull(t.Name) {
		return ErrInvalidName
	}
	return nil
}
