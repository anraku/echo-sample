package response

func CreateResponse(data interface{}) interface{} {
	return map[string]interface{}{
		"result": data,
	}
}
