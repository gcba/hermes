package responses

import (
	"net/http"
)

type Status struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var Statuses = map[int]Status{
	http.StatusOK: Status{
		Code:    http.StatusOK,
		Message: "OK"},
	http.StatusCreated: Status{
		Code:    http.StatusCreated,
		Message: "Created"},
	http.StatusBadRequest: Status{
		Code:    http.StatusBadRequest,
		Message: "Bad Request"},
	http.StatusForbidden: Status{
		Code:    http.StatusForbidden,
		Message: "Forbidden"},
	http.StatusNotFound: Status{
		Code:    http.StatusNotFound,
		Message: "Not Found"},
	http.StatusNotAcceptable: Status{
		Code:    http.StatusNotAcceptable,
		Message: "Not Acceptable"},
	http.StatusUnsupportedMediaType: Status{
		Code:    http.StatusUnsupportedMediaType,
		Message: "Unsupported Media Type"},
	http.StatusUnprocessableEntity: Status{
		Code:    http.StatusUnprocessableEntity,
		Message: "Unprocessable Entity"},
	http.StatusInternalServerError: Status{
		Code:    http.StatusInternalServerError,
		Message: "Internal Server Error"},
	http.StatusNotImplemented: Status{
		Code:    http.StatusNotImplemented,
		Message: "Not Implemented"}}
