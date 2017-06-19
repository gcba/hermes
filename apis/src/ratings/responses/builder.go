package responses

type Meta struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

const Meta200 = Meta{
	Code: 200, // OK
	Message: "OK"
}

const Meta201 = Meta{
	Code: 201, // Created
	Message: "Created"
}

const Meta400 = Meta{
	Code: 400, // Bad Request
	Message: "Bad Request"
}

const Meta401 = Meta{
	Code: 401, // Unauthorized
	Message: "Unauthorized"
}

const Meta403 = Meta{
	Code: 403, // Forbidden
	Message: "Forbidden"
}

const Meta404 = Meta{
	Code: 404, // Not Found
	Message: "Not Found"
}

const Meta422 = Meta{
	Code: 422, // Unprocessable Entity
	Message: "Unprocessable Entity"
}

const Meta500 = Meta{
	Code: 500, // Internal Server Error
	Message: "Internal Server Error"
}

const Meta501 = Meta{
	Code: 501, // Not Implemented
	Message: "Not Implemented"
}