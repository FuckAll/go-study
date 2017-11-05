/**
 *@Description:获取免费ShadowSockets
 *@author izgnod
 *@time 2017-11-05 下午15:44
 */

package main

import (
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/mdp/qrterminal"
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
	surgeConf := flag.Bool("surge", false, "生成surge软件配置")
	limit := flag.Int("limit", 3, "要多少就输入多少，不一定有这么多")
	backupURL := flag.Bool("backup_url", false, "防止被BAN，启用备用")
	method := flag.String("nsm", "", "过滤掉一些不想使用的加密方法. example: chacha20-ietf-poly1305,chacha20")
	qr := flag.Bool("qr", false, "当使用手机端的时候，此选项能够显示配置二维码")
	flag.Parse()

	// config := qrterminal.Config{
	// 	Level:     qrterminal.L,
	// 	Writer:    os.Stdout,
	// 	BlackChar: qrterminal.WHITE,
	// 	WhiteChar: qrterminal.BLACK,
	// }

	// qrterminal.GenerateWithConfig("https://github.com/mdp/qrterminal", config)
	// qrterminal.Generate("https://github.com/mdp/qrterminal", 3, os.Stdout)

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
		if *qr {
			showShadowSockets(best, true)
		} else {
			showShadowSockets(best, false)
		}

	}

}

func showShadowSockets(data []*SS, qr bool) {
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
		if qr {
			uri := fmt.Sprintf("%s:%s@%s:%s", d.Method, d.Password, d.IP, d.Port)
			encodeString := base64.StdEncoding.EncodeToString([]byte(uri))
			encodeString = "ss://" + encodeString + "#" + d.Geo + "-" + d.IP
			fmt.Println(encodeString)
			qrterminal.Generate(encodeString, 3, os.Stdout)
		}
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
