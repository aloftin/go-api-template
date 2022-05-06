package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"template-service/models"

	"github.com/gorilla/mux"
)

const sortDelimeter = ":"

var (
	// Get the realm
	realm = os.Getenv("REALM")
)

// RelayResponse relays an HTTP response along to a response writer
func RelayResponse(w http.ResponseWriter, res *http.Response) {
	w.Header().Set("Content-Type", res.Header.Get("Content-Type"))
	w.Header().Set("Content-Length", res.Header.Get("Content-Length"))
	w.WriteHeader(res.StatusCode)
	io.Copy(w, res.Body)
	res.Body.Close()
}

// RespondWithJSON sends a JSON response
func RespondWithJSON(ctx context.Context, w http.ResponseWriter, code int, payload interface{}) error {
	response, err := json.Marshal(payload)

	if err != nil {
		log.Println(err)
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err = w.Write(response)

	return err
}

// RespondWithRestError sends an error response based on a given service error
func RespondWithRestError(ctx context.Context, w http.ResponseWriter, serviceError error) {
	switch serviceError.(type) {
	case *models.InvalidInputError:
		respondWithError(ctx, w, http.StatusBadRequest, serviceError, serviceError.Error())
	case *models.ResourceNotFoundError:
		respondWithError(ctx, w, http.StatusNotFound, serviceError, serviceError.Error())
	default:
		respondWithError(ctx, w, http.StatusInternalServerError, serviceError, "Server error")
	}
}

const (
	defaultLimit  = "25"
	defaultOffset = "0"
	maxLimit      = 100
)

// GetPagingParamsFromRequest validates and returns paging parameters from an HTTP request
// This func tries to use sensible defaults in the absence of params: defaultLimit, defaultOffset, maxLimit
func GetPagingParamsFromRequest(ctx context.Context, r *http.Request) (models.PagingParameters, error) {
	var pagingParams models.PagingParameters
	var limit int
	var offset int

	limitParam := r.URL.Query().Get("limit")
	if limitParam == "" {
		limitParam = defaultLimit
	}

	offsetParam := r.URL.Query().Get("offset")
	if offsetParam == "" {
		offsetParam = defaultOffset
	}

	limit, limitErr := strconv.Atoi(limitParam)
	offset, offsetErr := strconv.Atoi(offsetParam)

	if limitErr != nil || offsetErr != nil {
		return pagingParams, &models.InvalidInputError{Message: "Paging parameters must be integers"}
	}

	if limit < 0 || offset < 0 {
		return pagingParams, &models.InvalidInputError{Message: "Paging parameters must be positive integers"}
	}

	if limit > maxLimit {
		limit = maxLimit
	}

	pagingParams.Limit = limit
	pagingParams.Offset = offset
	return pagingParams, nil
}

// GetInt64URLParameter parses a URL variable into an int64
func GetInt64URLParameter(r *http.Request, paramName string) (int64, error) {
	vars := mux.Vars(r)
	idstr := vars[paramName]

	return strconv.ParseInt(idstr, 10, 64)
}

// GetIntURLParameter parses a URL variable into an int64
func GetIntURLParameter(r *http.Request, paramName string) (int, error) {
	vars := mux.Vars(r)
	idstr := vars[paramName]

	return strconv.Atoi(idstr)
}

// ParseSortParameter parse the provided sort parameter into sort and sort direction
func ParseSortParameter(sortParam string) (string, string) {

	// Default values for sort and sort direction
	sort := ""
	sortDir := ""

	if sortParam != "" {

		// Split the sort parameter on the delimiter
		sortParts := strings.Split(sortParam, sortDelimeter)
		sort = sortParts[0]

		// Set the sort direction if it's specified
		if len(sortParts) > 1 {
			sortDir = sortParts[1]
		}
	}

	return sort, sortDir
}

// IsSuccessCode checks if a status code is 200-level
func IsSuccessCode(code int) bool {
	return code >= 200 && code <= 299
}

// respondWithNotFound sends a 404 response with an empty payload
func respondWithNotFound(ctx context.Context, w http.ResponseWriter, err error) {
	RespondWithJSON(ctx, w, http.StatusNotFound, map[string]string{"error": err.Error()})
}

// respondWithError sends an error response
func respondWithError(ctx context.Context, w http.ResponseWriter, code int, err error, message string) {
	if err != nil {
		log.Println(err)
		fmt.Printf("%s\n%s", message, err)
	} else {
		log.Println(err)
		fmt.Println(message)
	}

	RespondWithJSON(ctx, w, code, map[string]string{"error": message})
}
