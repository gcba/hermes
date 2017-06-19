package responses

import (
	"net/http"
)

type Meta struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

const (
	Meta200 = Meta{
		Code: http.StatusOK,
		Message: "OK"
	}

	Meta201 = Meta{
		Code: http.StatusCreated,
		Message: "Created"
	}

	Meta400 = Meta{
		Code: http.StatusBadRequest,
		Message: "Bad Request"
	}

	Meta401 = Meta{
		Code: http.StatusUnauthorized,
		Message: "Unauthorized"
	}

	Meta403 = Meta{
		Code: http.StatusForbidden,
		Message: "Forbidden"
	}

	Meta404 = Meta{
		Code: http.StatusNotFound,
		Message: "Not Found"
	}

	Meta422 = Meta{
		Code: http.StatusUnprocessableEntity,
		Message: "Unprocessable Entity"
	}

	Meta500 = Meta{
		Code: http.StatusInternalServerError,
		Message: "Internal Server Error"
	}

	Meta501 = Meta{
		Code: http.NotImplemented,
		Message: "Not Implemented"
	}
)