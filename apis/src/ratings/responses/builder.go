package responses

type Meta struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

const Response200 = Meta{
	Code: 200, // OK
	Message: "OK"
}

const Response201 = Meta{
	Code: 201, // Created
	Message: "Created"
}

const Response400 = Meta{
	Code: 400, // Bad Request
	Message: "Bad Request"
}

const Response401 = Meta{
	Code: 401, // Unauthorized
	Message: "Unauthorized"
}

const Response403 = Meta{
	Code: 403, // Forbidden
	Message: "Forbidden"
}

const Response404 = Meta{
	Code: 404, // Not Found
	Message: "Not Found"
}

const Response422 = Meta{
	Code: 422, // Unprocessable Entity
	Message: "Unprocessable Entity"
}

const Response500 = Meta{
	Code: 500, // Internal Server Error
	Message: "Internal Server Error"
}

const Response501 = Meta{
	Code: 501, // Not Implemented
	Message: "Not Implemented"
}