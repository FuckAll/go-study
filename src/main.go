/**
 *@Description:获取免费ShadowSockets
 *@author izgnod
 *@time 2017-11-05 下午15:44
 */

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
)

const (
	url       = "https://ss.weirch.com/ss.php"
	backupURL = "http://mirror.weirch.com/ss.php"
	testURL   = "http://www.gstatic.com/generate_204"
	ssModule  = "https://github.com/R0uter/ss.conf-for-surge/raw/master/ss.module"
)

// SS shadowsockets info
type SS struct {
	Health   float64 `json:"health"`
	IP       string  `json:"ip"`
	Port     string  `json:"port"`
	Password string  `json:"password"`
	Method   string  `json:"method"`
	Verified string  `json:"verified"`
	Geo      string  `json:"geo"`
}

type respData struct {
	Data []interface{} `json:"data"`
}

func getData(backup bool) []*SS {
	client := &http.Client{}
	timeStamp := time.Now().Unix()
	random := random(100, 999)

	parameter := strconv.Itoa(int(timeStamp)) + strconv.Itoa(random)
	useURL := url
	if backup {
		useURL = backupURL
	}

	req, err := http.NewRequest("GET", useURL, nil)
	if err != nil {
		panic(err)
	}

	q := req.URL.Query()
	q.Add("_", parameter)
	req.URL.RawQuery = q.Encode()

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)

	}
	re := &respData{}
	err = json.Unmarshal(body, &re)
	if err != nil {
		panic(err)
	}

	ret := []*SS{}
	for _, v := range re.Data {
		tmp := new(SS)
		for k, vv := range v.([]interface{}) {
			switch k {
			case 0:
				tmp.Health = vv.(float64)
			case 1:
				tmp.IP = vv.(string)
			case 2:
				tmp.Port = vv.(string)
			case 3:
				tmp.Password = vv.(string)
			case 4:
				tmp.Method = vv.(string)
			case 5:
				tmp.Verified = vv.(string)
			case 6:
				tmp.Geo = vv.(string)
			}
		}
		ret = append(ret, tmp)
	}
	return ret
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func main() {
	surgeConf := flag.Bool("surge", false, "get surge conf.")
	limit := flag.Int("limit", 3, "you need how many data.")
	backupURL := flag.Bool("backup_url", false, "use backup url.")
	method := flag.String("nsm", "", "method non support. example: chacha20-ietf-poly1305,chacha20")
	flag.Parse()

	nonsupportMethod := []string{}
	nonsupportMethod = append(nonsupportMethod, strings.Split(*method, ",")...)
	data := getData(*backupURL)
	best := []*SS{}
	var count int
	// filter just need health 100
	for _, v := range data {
		if v.Health == 100 {
			var ok bool
			for _, m := range nonsupportMethod {
				if m == v.Method {
					ok = true
				}
			}
			if ok {
				continue
			}
			if count >= *limit {
				break
			}

			best = append(best, v)
			count++
		}
	}
	if *surgeConf {
		generateSurgeConf(best)
	} else {
		showShadowSockets(best)
	}

}

func showShadowSockets(data []*SS) {
	showStr := `
	=======================
	Area: %s
	Host: %s 
	Port: %s
	Method: %s
	Password: %s 
	=======================
	`
	for _, d := range data {
		color.Green(showStr, d.Geo, d.IP, d.Port, d.Method, d.Password)
	}
}

func generateSurgeConf(data []*SS) {
	geo := map[string]int{}
	geoName := []string{}
	showStr := "[Proxy]\n💊DIRECT = direct\n"
	customStr := "%s = custom, %s,%s,%s,%s,%s\n"
	for _, d := range data {
		name := d.Geo + strconv.Itoa(geo[d.Geo]+1)
		geo[d.Geo] = geo[d.Geo] + 1
		geoName = append(geoName, name)
		showStr = showStr + fmt.Sprintf(customStr, name, d.IP, d.Port, d.Method, d.Password, ssModule)
	}
	chinaProxyStr := "[Proxy Group]\nChinaProxy = select, 💊DIRECT, "
	proxyStr := "Proxy = select, 💊DIRECT, @Auto, "
	autoStr := "@Auto = url-test, "

	for _, d := range geoName {
		chinaProxyStr = chinaProxyStr + d + ","
		proxyStr = proxyStr + d + ","
		autoStr = autoStr + d + ","
	}

	chinaProxyStr = strings.TrimSuffix(chinaProxyStr, ",") + "\n"
	proxyStr = strings.TrimSuffix(proxyStr, ",") + "\n"
	autoStr = autoStr + "url =" + testURL + "\n"

	color.Green(showStr + "\n")

	color.Green(chinaProxyStr)
	color.Green(proxyStr)
	color.Green(autoStr)
}
