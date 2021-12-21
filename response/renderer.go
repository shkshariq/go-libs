package response

import (
	"context"
	"encoding/json"
	"net/http"

	error_handler "github.com/shkshariq/go-util/error-handler"
)

type errorResponse struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Details interface{} `json:"details"`
	Debug   interface{} `json:"debug"`
}

func HandleError(ctx context.Context, err error, w http.ResponseWriter) {

	if !isPublicVisible(err) {
		errResponse := errorResponse{
			Message: `Something went wrong`,
			Code:    1000,
			Details: `Not available`,
			Debug:   err.Error(),
		}

		genericError(ctx, errResponse, w)
		return
	}

	if domainError, ok := isDomainError(err); ok {

		body := errorResponse{
			Message: domainError.Message,
			Code:    domainError.Code,
			Details: domainError.Details,
		}

		body.Debug = domainError.Error()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(body)
		return
	}

	if applicationError, ok := isApplicationError(err); ok {

		body := errorResponse{
			Message: applicationError.Message,
			Code:    10000,
			Details: applicationError.Details,
		}

		body.Debug = applicationError.Error()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(body)
	}
}

func isPublicVisible(err error) bool {

	_, isDomain := err.(error_handler.DomainError)
	isApplication := false

	if !isDomain {
		_, isApplication = err.(error_handler.ApplicationError)
	}

	return isDomain || (isApplication)
}

func isDomainError(err error) (error_handler.DomainError, bool) {
	domainError, isDomain := err.(error_handler.DomainError)
	return domainError, isDomain
}

func isApplicationError(err error) (error_handler.ApplicationError, bool) {
	applicationError, isApplication := err.(error_handler.ApplicationError)
	return applicationError, isApplication
}

//If application Debug is enabled then only show application errors
func genericError(_ context.Context, err errorResponse, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(err)
}
