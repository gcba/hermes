package responses

import (
	"net/http"
)

type Meta struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

const (
	meta200 = Meta{
		Code: http.StatusOK,
		Message: "OK"
	}

	meta201 = Meta{
		Code: http.StatusCreated,
		Message: "Created"
	}

	meta400 = Meta{
		Code: http.StatusBadRequest,
		Message: "Bad Request"
	}

	meta401 = Meta{
		Code: http.StatusUnauthorized,
		Message: "Unauthorized"
	}

	meta403 = Meta{
		Code: http.StatusForbidden,
		Message: "Forbidden"
	}

	meta404 = Meta{
		Code: http.StatusNotFound,
		Message: "Not Found"
	}

	meta422 = Meta{
		Code: http.StatusUnprocessableEntity,
		Message: "Unprocessable Entity"
	}

	meta500 = Meta{
		Code: http.StatusInternalServerError,
		Message: "Internal Server Error"
	}

	meta501 = Meta{
		Code: http.NotImplemented,
		Message: "Not Implemented"
	}

	metas = map[int]Meta{
		http.StatusOK: meta200,
		http.StatusCreated: meta201,
		http.StatusBadRequest: meta400,
		http.StatusUnauthorized: meta401,
		http.StatusForbidden: meta403,
		http.StatusNotFound: meta404,
		http.StatusUnprocessableEntity: meta422,
		http.StatusInternalServerError: meta500,
		http.NotImplemented: meta501
	}
)