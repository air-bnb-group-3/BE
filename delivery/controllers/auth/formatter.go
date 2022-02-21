package auth

type LoginRespFormat struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

type LoginReqFormat struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
