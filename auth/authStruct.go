package auth

type Authentication struct {
	Name  string `json:"name"`
	Role  string `json:"role"`
	Token string `json:"token"`
}


type Claims map[string]interface{}