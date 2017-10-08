package conf
import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"strings"

	"fmt"
)

// Opening (or creating) config file in JSON format
func ReadConfig(filename string) Config {
	config := Config{}
	//if _, err := os.Stat(GetRoot() + "/conf/" + "config.json"); os.IsNotExist(err) {}
	fmt.Print("filename: ",filename)
	file, err := os.Open(filename)

	defer file.Close()
	if err != nil {
		log.Fatalln(filename, " is not existent.")
	} else {
		err = json.NewDecoder(file).Decode(&config)
		if err != nil {
			log.Fatal(err)
		}
	}
	return config
}

func ParseConfig() (Config) {

	cfile := "config.json"
	cfg := ReadConfig(GetRoot() + "/" + cfile)
	if cfg.Port == 0 {
		cfg.Port = 80
	}
	if cfg.Hostname == "" {
		cfg.Hostname, _ = os.Hostname()
	}
	if cfg.IP == "" {
		cfg.IP = "127.0.0.1"
	}

	log.Printf("Config loaded")
	//db, err := sql.Open("sqlite3", cfg.Db)
	return cfg
}

func GetRoot() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	directory := strings.Replace(dir, "\\", "/", -1)

	return directory
}

