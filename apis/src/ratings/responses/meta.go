package responses

import (
	"net/http"
)

type meta struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

const (
	meta200 = meta{
		Code: http.StatusOK,
		Message: "OK"
	}

	meta201 = meta{
		Code: http.StatusCreated,
		Message: "Created"
	}

	meta400 = meta{
		Code: http.StatusBadRequest,
		Message: "Bad Request"
	}

	meta401 = meta{
		Code: http.StatusUnauthorized,
		Message: "Unauthorized"
	}

	meta403 = meta{
		Code: http.StatusForbidden,
		Message: "Forbidden"
	}

	meta404 = meta{
		Code: http.StatusNotFound,
		Message: "Not Found"
	}

	meta422 = meta{
		Code: http.StatusUnprocessableEntity,
		Message: "Unprocessable Entity"
	}

	meta500 = meta{
		Code: http.StatusInternalServerError,
		Message: "Internal Server Error"
	}

	meta501 = meta{
		Code: http.NotImplemented,
		Message: "Not Implemented"
	}

	metas = map[int]meta{
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