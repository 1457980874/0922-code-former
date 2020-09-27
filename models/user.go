package models

type User struct {
	User string `json:"name"`
	Birthday string `json:"birthday"`
	Address string `json:"address"`
	Nick string `json:"nick"`
	Password string `json:"password"`
}

/*type User struct {
	Name string
	Birthday string
	Address string
	Nick string
}
*/