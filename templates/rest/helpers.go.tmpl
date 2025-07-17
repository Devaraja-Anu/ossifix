package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type envelope map[string]any

func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {

	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	js = append(js, '\n')

	//loop  through the header map and add each header to the http.ResponseWriter header map.

	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}

func (app *application) ReadJSON(w http.ResponseWriter, r *http.Request, destination any) error {

	//The underscores are ignored by the Go compiler and are only there to help with readability.
	maxBytes := 1_048_576
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	dec := json.NewDecoder(r.Body)

	//If the JSON from the client now includes any  field which cannot be mapped to the target destination,
	//  the decoder will return an error instead of just ignoring the field
	dec.DisallowUnknownFields()

	err := dec.Decode(destination)
	if err != nil {
		var syntaxError *json.SyntaxError
		var umarshalTypeError *json.UnmarshalTypeError
		var invalidUmarshalError *json.InvalidUnmarshalError

		switch {
		case errors.As(err, &syntaxError):
			return fmt.Errorf("body contains badly formed character at %d", syntaxError.Offset)
		case errors.Is(err, io.ErrUnexpectedEOF):
			return errors.New("body contains badly-formed JSON")
		case errors.As(err, &umarshalTypeError):
			if umarshalTypeError.Field != "" {
				return fmt.Errorf("body contains incorrect JSON type for field %q", umarshalTypeError.Field)
			}
			return fmt.Errorf("body contains incorrect JSON type at character %d", umarshalTypeError.Offset)
		case errors.Is(err, io.EOF):
			return errors.New("body must not be empty")
		case strings.HasPrefix(err.Error(), "json: unknown field"):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field")
			return fmt.Errorf("body contains unknown key %s", fieldName)
		case err.Error() == "http: request body too large":
			return fmt.Errorf("body must not be larger than %d MB", (maxBytes / 1048756))
		case errors.As(err, &invalidUmarshalError):
			panic(err)
		default:
			return err
		}
	}

	// once decoded decode again this time to an anonymous empty struct. If this is not an EOF error that means
	// there is more than one JSON value
	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must contain only a single JSON value")
	}
	return nil
}

func (app *application) readString(qs url.Values, key, defaultvalue string) string {

	s := qs.Get(key)
	if s == "" {
		return defaultvalue
	}
	return s
}
