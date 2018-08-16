package main

/**
 * Package: DynDNS for UnitedDomains Reselling
 * Usage on own risk
 * No Support given, feel free to fork and modify =)
 *
 * Done by Bastian Bringenberg <bastian@agentur-pottkinder.de>
 * External IP from: <http://myexternalip.com/#golang>
 */

import (
	"github.com/BurntSushi/toml"
	_ "io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

type Config struct {
	Cnamemaster string // unused
	Subdomain   string
	Domain      string
	User        string
	Pass        string
}

func readConfigfile() Config {
	var configfile = "properties.ini"
	_, err := os.Stat(configfile)
	if err != nil {
		log.Fatal("Config file is missing: ", configfile)
	}

	var config Config
	if _, err := toml.DecodeFile(configfile, &config); err != nil {
		log.Fatal(err)
	}
	return config
}

func dyndns(text string, config Config) {
	v := url.Values{}
	v.Add("s_login", config.User)
	v.Add("s_pw", config.Pass)
	v.Add("command", "UpdateDNSZone")
	v.Add("dnszone", config.Domain+".")
	v.Add("rr0", config.Subdomain+". 600 IN TXT "+text)
	resp, err := http.PostForm("https://api.domainreselling.de/api/call.cgi", v)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
}

func main() {
	config := readConfigfile()
	dyndns("Test0123456789", config)
}
