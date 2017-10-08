package conf



type Config struct {
	Hostname        string `json:"hostname"`
	IP 				string	`json:"ip"`
	Port         	int	`json:"port"`
	Templates       string	`json:"templates"`
	StaticFiles     string	`json:"staticfiles"`
	MenuList      	[]Menu	`json:"menulist"`
}


type Menu struct {
	Name   string `json:"name"`
	Link  string `json:"link"`
}
