package controllers

type H map[string]interface{}

func SuccessResponse(data interface{}) H {
	return H{
		"status": "success",
		"data":   data,
	}
}
