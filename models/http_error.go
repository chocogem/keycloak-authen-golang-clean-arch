package models

type HttpError struct {
	Code int `json:"code"`
	Error string `json:"error"`
	Message string `json:"message"`
}

func UnauthorizedError()  HttpError{
	return HttpError{
		401,
		"Unauthorized",
		"Authentication Failed",
	}

}

func TokenNotFoundError()  *HttpError{
	return &HttpError{
		404,
		"Not found",
		"Token not found in the request header",
	}

}


func InternalServerError(message string)  *HttpError{
	return &HttpError{
		500,
		"Internal server error",
		message,
	}

}