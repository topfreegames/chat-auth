package chatauth

type userAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Salt     string `json:"salt"`
}
