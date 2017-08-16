package responses

import (
	"net/http"
)

type meta struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var metas = map[int]meta{
	http.StatusOK: meta{
		Code:    http.StatusOK,
		Message: "OK"},
	http.StatusCreated: meta{
		Code:    http.StatusCreated,
		Message: "Created"},
	http.StatusBadRequest: meta{
		Code:    http.StatusBadRequest,
		Message: "Bad Request"},
	http.StatusForbidden: meta{
		Code:    http.StatusForbidden,
		Message: "Forbidden"},
	http.StatusNotFound: meta{
		Code:    http.StatusNotFound,
		Message: "Not Found"},
	http.StatusNotAcceptable: meta{
		Code:    http.StatusNotAcceptable,
		Message: "Not Acceptable"},
	http.StatusUnsupportedMediaType: meta{
		Code:    http.StatusUnsupportedMediaType,
		Message: "Unsupported Media Type"},
	http.StatusUnprocessableEntity: meta{
		Code:    http.StatusUnprocessableEntity,
		Message: "Unprocessable Entity"},
	http.StatusInternalServerError: meta{
		Code:    http.StatusInternalServerError,
		Message: "Internal Server Error"},
	http.StatusNotImplemented: meta{
		Code:    http.StatusNotImplemented,
		Message: "Not Implemented"}}
