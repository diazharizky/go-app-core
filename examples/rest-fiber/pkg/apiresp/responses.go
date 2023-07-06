package apiresp

type ErrDetail struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}

type Response struct {
	OK     bool        `json:"ok"`
	Data   interface{} `json:"data,omitempty"`
	Errors []ErrDetail `json:"errors"`
}

func Default() Response {
	return Response{
		OK:     true,
		Errors: []ErrDetail{},
	}
}

func FatalError() Response {
	return Response{
		OK: false,
		Errors: []ErrDetail{
			{
				Code:        ErrFatal,
				Description: "Internal server error",
			},
		},
	}
}

func BadRequestError() Response {
	return Response{
		OK: false,
		Errors: []ErrDetail{
			{
				Code:        ErrBadRequest,
				Description: "Bad request",
			},
		},
	}
}
