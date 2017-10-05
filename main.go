package main

import (
	"InfraManager/auth"
	"InfraManager/tools"
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
	"strconv"
	"strings"
)

func main() {
	http.Handle("/", &Router{config: make(map[string]interface{})})
	http.Handle("/bootstrap/", http.StripPrefix("/", http.FileServer(http.Dir("./static/"))))
	//http.HandleFunc("/login", LoginHandler)
	log.Fatal(http.ListenAndServe(":80", nil))
}

type Server struct {
	ID       int    `json:"id"`
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
	var infra_delServers = regexp.MustCompile(`^/infra/delservers$`)
	var admin_login = regexp.MustCompile(`^/admin/login$`)
	var infra_ping = regexp.MustCompile(`^/infra/ping$`)

	switch {
	case home.MatchString(r.URL.Path):
		HomeHandler(w, r)
	case infra.MatchString(r.URL.Path):
		InfraHandler(w, r)
	case infra_listServer.MatchString(r.URL.Path):
		fmt.Println("Forward to ListServerHandler")
		ListServerHandler(w, r)

	case infra_getServers.MatchString(r.URL.Path):
		GetServersHandler(w, r)
	case infra_delServers.MatchString(r.URL.Path):
		DelServersHandler(w, r)
	case admin_login.MatchString(r.URL.Path):
		LoginHandler(w, r)
	case infra_ping.MatchString(r.URL.Path):
		PingHandler(w, r)
	default:
		//fmt.Fprintf(w, "err:Not Registed to %q", html.EscapeString(r.URL.Path))
		fmt.Fprintf(w, "err:\"Not Registed to "+html.EscapeString(r.URL.Path)+"\"")
	}
	fmt.Println("ServeHTTP finished!")
}

func PingHandler(w http.ResponseWriter, r *http.Request) {
	var data string

	if r.Method == "POST" {
		data = r.PostFormValue("data")
		items := strings.Split(data, ",")
		fmt.Println("itesms:", items)

		pingData:= []tools.PingSt{}
		for i := 0; i < len(items); i++ {
			pingres := tools.Ping(items[i])
			pingData = append(pingData, pingres)


		}
		retData,err := json.Marshal(pingData)
		checkError(err)
		w.Header().Set("Content-type","application/json")
		_,err = w.Write(retData)
		checkError(err)
	}


	if r.Method == "GET" {
		w.Write([]byte("Please use Post!"))
	}
}

const (
	mySigningKey = "WOW,MuchShibe,ToDogge"
)

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
	fmt.Println("ListServer finishied")
}
func GetServersHandler(w http.ResponseWriter, r *http.Request) {
	var sql_select string
	var page, pagesize int
	var offset, limit string
	if r.Method == "GET" {
		queryForm, err := url.ParseQuery(r.URL.RawQuery)
		checkError(err)
		r.ParseForm()
		limit = queryForm["size"][0]
		page, _ = strconv.Atoi(queryForm["page"][0])
	} else if r.Method == "POST" {
		//r.MultipartForm.Value["id"]
		limit = r.PostFormValue("size")
		page, _ = strconv.Atoi(r.PostFormValue("page"))
	}
	pagesize, _ = strconv.Atoi(limit)
	offset = strconv.Itoa(page*pagesize - pagesize)

	fmt.Println("offset: ", offset)
	fmt.Println("limit: ", limit)
	//{sortOrder: "asc", pageSize: 10, pageNumber: 1}
	//sql_select = "select hostname,inet_ntoa(ip1) as ip,os,platform,ip2,ip3 from servers LIMIT 10 OFFSET 10"
	sql_select = "select server_id,hostname,inet_ntoa(ip1) as ip,os,platform,ip2,ip3 from servers LIMIT " + limit + " OFFSET " + offset
	paging := getServers(sql_select, pagesize)
	json_paging, err := json.Marshal(paging)
	checkError(err)
	w.Header().Set("Content-type", "application/json")

	_, err = w.Write(json_paging)
	checkError(err)
}
func DelServersHandler(w http.ResponseWriter, r *http.Request) {
	var data string
	var items []string
	if r.Method == "POST" {
		data = r.PostFormValue("data")
		if data != "" {
			items = strings.Split(data, ",")
		}
	}
	fmt.Println("itesms:", items)
	for i := 0; i < len(items); i++ {
		go delServer(items[i])
	}
	w.Write([]byte("ok"))
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

func InfraHandler(w http.ResponseWriter, r *http.Request) {
	dumx := Authentication{Name: "guest", Token: "guest"}
	param := Parameter{Title: "Infra Manager", Auth: dumx}
	//onlineUser := OnlineUser{User: []*Person{&dumx, &chxd}}
	//t := template.New("Person template")
	//t, err := t.Parse(templ)
	t, err := template.ParseFiles("templates/layout.html", "templates/listserver.html")
	checkError(err)
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
	db, err := sql.Open("mysql", "root:pa55word@tcp(192.168.6.5:3306)/test?charset=utf8")
	checkError(err)
	defer db.Close()
	row := db.QueryRow("SELECT user_name,user_pw FROM auth WHERE user_name=? and user_pw=?", username, userpw)
	err = row.Scan(&uname, &upw)

	return uname
}
func delServer(hostname string) string {
	db, err := sql.Open("mysql", "root:pa55word@tcp(192.168.6.5:3306)/test?charset=utf8")
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
	db, err := sql.Open("mysql", "root:pa55word@tcp(192.168.6.5:3306)/test?charset=utf8")
	checkError(err)
	defer db.Close()

	row := db.QueryRow("select count(*) from servers")
	row.Scan(&totalrow)

	rows, err := db.Query(sql_select)
	checkError(err)

	server := Server{}
	servers := []Server{}
	var hostname, ip, platform, os_filed, ip2, ip3 string
	var server_id int
	for rows.Next() {
		rows.Scan(&server_id, &hostname, &ip, &platform, &os_filed, &ip2, &ip3)
		server.ID = server_id
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
