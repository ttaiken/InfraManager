package main

import (
	"InfraManager/auth"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
	"html"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"

	//"TalkingData-owl/Godeps/_workspace/src/github.com/gorilla/mux"

	"strconv"
)

func main() {

	//http.Handle("/user/", &Router{config: make(map[string]interface{})})
	//http.HandleFunc("/infra", InfraHandler)
	//http.HandleFunc("/login", LoginHandler)
	//router := mux.NewRouter()
	//router := mux.NewRouter()
	//router := mux.NewRouter()
	http.Handle("/", &Router{config: make(map[string]interface{})})
	http.Handle("/bootstrap/", http.StripPrefix("/", http.FileServer(http.Dir("./static/"))))

	log.Fatal(http.ListenAndServe(":8000", nil))
}

type Server struct {
	Hostname string `json:"hostname"`
	IP       string `json:"ip"`
	OS       string `json:"os"`
	Platform string `json:"platform"`
	IP2      string `json:"ip2"`
	IP3      string `json:"ip3"`
}

type Paging struct {
	Last_page int      `json:"last_page"`
	Data      []Server `json:"data"`
}

type Router struct {
	config map[string]interface{}
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var home = regexp.MustCompile(`^/$`)
	var infra = regexp.MustCompile(`^/infra$`)
	var infra_listServer = regexp.MustCompile(`^/infra/listserver$`)
	var infra_getServers = regexp.MustCompile(`^/infra/getservers$`)
	var infra_delServer = regexp.MustCompile(`^/infra/delserver$`)
	var admin_login = regexp.MustCompile(`^/admin/login$`)

	//var logout = regexp.MustCompile(`^/user/logout`)
	//var addUser = regexp.MustCompile(`^/user/adduser`)
	//var delUser = regexp.MustCompile(`^/user/deluser`)

	switch {

	case home.MatchString(r.URL.Path):
		HomeHandler(w, r)
	case infra.MatchString(r.URL.Path):
		InfraHandler(w, r)
	case infra_listServer.MatchString(r.URL.Path):
		ListServerHandler(w, r)
	case infra_getServers.MatchString(r.URL.Path):
		GetServersHandler(w, r)
	case infra_delServer.MatchString(r.URL.Path):
		DelServerHandler(w, r)
	case admin_login.MatchString(r.URL.Path):
		LoginHandler(w, r)

	default:
		fmt.Fprintf(w, "No Registed form %q", html.EscapeString(r.URL.Path))

	}

}

const (
	mySigningKey = "WOW,MuchShibe,ToDogge"
)

//type Info struct {
//	user string
//	exp  int64
//}

type Claims map[string]interface{}

type Authentication struct {
	Name  string
	Token string
}

type Parameter struct {
	Title   string
	Auth    Authentication
	Servers []Server
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	//dir, _ := os.Getwd()
	//t, err := template.ParseFiles(filepath.Join(dir, "templates", "home.html"))
	//checkError(err)
	//err = t.Execute(w, param)
	dumx := Authentication{Name: "guest", Token: "guest"}
	param := Parameter{Title: "Infra Manager", Auth: dumx}
	t, err := template.ParseFiles("templates/layout.html", "templates/home.html")
	checkError(err)
	err = t.ExecuteTemplate(w, "layout", param)
	checkError(err)
}

func ListServerHandler(w http.ResponseWriter, r *http.Request) {
	dumx := Authentication{Name: "guest", Token: "guest"}
	param := Parameter{Title: "Infra Manager", Auth: dumx}
	t, err := template.ParseFiles("templates/layout.html", "templates/listserver.html")
	checkError(err)
	err = t.ExecuteTemplate(w, "layout", param)
	checkError(err)
}
func DelServerHandler(w http.ResponseWriter, r *http.Request) {
	queryForm, err := url.ParseQuery(r.URL.RawQuery)
	checkError(err)
	r.ParseForm()
	var hostname string
	hostname = queryForm["hostname"][0]
	println("hostname: ", hostname)
	info := delServer(hostname)
	w.Write([]byte(info))

}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	cookieToken, err := r.Cookie("token")
	if err == nil {
		fmt.Println("cookie：", cookieToken)
		fmt.Println("cookie value: ", cookieToken.Value)
		token, ok := auth.TokenParse(cookieToken.Value, mySigningKey)
		if ok == "ok" {
			fmt.Println("token: ", token)
			claims, _ := token.Claims.(jwt.MapClaims)
			fmt.Println("claims: ", claims)
			tmpname := claims["user"].(string)
			fmt.Println("tmpuser: ", tmpname)
			dumx := Authentication{Name: tmpname, Token: cookieToken.Value}
			fmt.Println("dumx: ", dumx)
			param := Parameter{Title: "Login Page", Auth: dumx}
			t := template.New("hello template")
			t, _ = t.Parse("Hi {{.Name}}!, Do you want to logout and login again?")
			err := t.Execute(w, param.Auth)
			checkError(err)
			return
		}
	}
	if r.Method == "GET" {
		dumx := Authentication{Name: "guest", Token: "guest"}
		param := Parameter{Title: "Infra Manager", Auth: dumx}
		t, err := template.ParseFiles("templates/layout.html", "templates/login.html")
		checkError(err)
		err = t.ExecuteTemplate(w, "layout", param)
		checkError(err)
	} else if r.Method == "POST" {
		username := r.FormValue("username")
		userpw := r.FormValue("userpw")
		tmpuser := checkUser(string(username), string(userpw))
		fmt.Println("user is : %s", tmpuser)

		if tmpuser == "" {
			tmpuser = "nouser"
		}
		token, err := auth.TokenNew([]byte(mySigningKey), tmpuser)

		userCredential := Authentication{Name: tmpuser, Token: token}
		param := Parameter{Title: tmpuser, Auth: userCredential}

		cookie1 := http.Cookie{
			Name:  "token",
			Value: token,
		}
		w.Header().Set("Set-Cookie", cookie1.String())

		t := template.New("hello template")
		t, _ = t.Parse("welecome {{.Token}}!")
		checkError(err)
		err = t.Execute(w, param.Auth)
		checkError(err)
	}
}

func GetServersHandler(w http.ResponseWriter, r *http.Request) {
	//page - the page number being requested
	//size - the number of rows to a page (if paginationSize is set)
	//sort - the first currently sorted field (if any)
	//sort_dir - the first currently sort direction (if any)
	//filter - the first currently filtered field (if any)
	//filter_value - the first current filter value (if any)
	//filter_type - the first current filter type (if any)

	var sql_select string
	var page, pagesize int
	var offset, limit string
	queryForm, err := url.ParseQuery(r.URL.RawQuery)
	r.ParseForm()

	pagesize, _ = strconv.Atoi(queryForm["size"][0])
	page, _ = strconv.Atoi(queryForm["page"][0])
	limit = queryForm["size"][0]
	offset = strconv.Itoa(page*pagesize - pagesize)

	fmt.Println("offset: ", offset)
	fmt.Println("limit: ", limit)
	//{sortOrder: "asc", pageSize: 10, pageNumber: 1}
	//sql_select = "select hostname,inet_ntoa(ip1) as ip,os,platform,ip2,ip3 from servers LIMIT 10 OFFSET 10"
	sql_select = "select hostname,inet_ntoa(ip1) as ip,os,platform,ip2,ip3 from servers LIMIT " + limit + " OFFSET " + offset
	paging := getServers(sql_select, pagesize)
	println("GetServerHandler")
	println(sql_select)
	println()
	//servers := paging["Rows"]
	//json_server, err := json.Marshal(servers)
	json_paging, err := json.Marshal(paging)
	checkError(err)
	//fmt.Println("servers: ", paging)
	//fmt.Println("dict: ", json_paging)
	//w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Content-type", "application/json")

	_, err = w.Write(json_paging)
	checkError(err)
}

func InfraHandler(w http.ResponseWriter, r *http.Request) {
	dumx := Authentication{Name: "guest", Token: "guest"}
	param := Parameter{Title: "Infra Manager", Auth: dumx}
	//onlineUser := OnlineUser{User: []*Person{&dumx, &chxd}}
	//t := template.New("Person template")
	//t, err := t.Parse(templ)
	t, err := template.ParseFiles("templates/nav.html", "templates/infra.html")
	checkError(err)

	//err = t.Execute(w, param)
	err = t.ExecuteTemplate(w, "layout", param)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
func checkUser(username string, userpw string) string {
	var uname, upw string
	db, err := sql.Open("mysql", "root:pa55word@tcp(192.168.6.1:3306)/test?charset=utf8")
	checkError(err)
	defer db.Close()
	row := db.QueryRow("SELECT user_name,user_pw FROM auth WHERE user_name=? and user_pw=?", username, userpw)
	err = row.Scan(&uname, &upw)

	return uname
}
func delServer(hostname string) string {
	db, err := sql.Open("mysql", "root:pa55word@tcp(192.168.6.1:3306)/test?charset=utf8")
	checkError(err)
	defer db.Close()
	stmt, err := db.Prepare(`DELETE FROM servers WHERE hostname=?`)
	checkError(err)
	res, err := stmt.Exec(hostname)
	checkError(err)
	num, err := res.RowsAffected()
	checkError(err)
	return "Deleted: " + fmt.Sprintf("%d", num) + " row"
}

func getServers(sql_select string, pagesize int) Paging {
	paging := Paging{}
	var totalrow, totalpage int
	db, err := sql.Open("mysql", "root:pa55word@tcp(192.168.6.1:3306)/test?charset=utf8")
	checkError(err)
	defer db.Close()

	row := db.QueryRow("select count(*) from servers")
	row.Scan(&totalrow)

	rows, err := db.Query(sql_select)
	checkError(err)

	server := Server{}
	servers := []Server{}
	var hostname, ip, platform, os_filed, ip2, ip3 string
	for rows.Next() {
		rows.Scan(&hostname, &ip, &platform, &os_filed, &ip2, &ip3)
		server.Hostname = hostname
		server.IP = ip
		server.OS = os_filed
		server.Platform = platform
		server.IP2 = ip2
		server.IP3 = ip3
		servers = append(servers, server)
	}
	db.Close()

	if totalrow%pagesize == 0 {
		totalpage = totalrow / pagesize
	} else {
		totalpage = totalrow/pagesize + 1
	}
	paging = Paging{Last_page: totalpage, Data: servers}
	return paging
}
