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

	"github.com/PuerkitoBio/goquery"
	"github.com/fatih/color"
	"github.com/mdp/qrterminal"
)

const (
	url       = "https://free-ss.site/ss.php"
	backupURL = "https://free-ss.gq/ss.php"
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

	time.Sleep(8 * time.Second)
	resp, err = client.Do(req)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	body = []byte(`{"data": [[100, "162.254.201.254", "4436", "^CT&zd7*Ra", "aes-256-cfb", "10:32:04", "US"], [90, "134.119.214.215", "4436", "^CT&zd7*Ra", "aes-256-cfb", "10:32:04", "FR"], [95, "139.59.105.216", "18053", "54164519", "aes-256-cfb", "10:32:05", "SG"], [100, "128.199.143.153", "443", "Zjc789456123?", "aes-256-cfb", "10:32:09", "SG"], [100, "172.247.33.184", "12345", "ya2uz2et", "aes-256-cfb", "10:32:11", "US"], [100, "64.137.247.221", "8989", "0123456789", "chacha20", "10:32:03", "CA"], [100, "164.132.47.151", "4436", "^CT&zd7*Ra", "aes-256-cfb", "10:32:05", "FR"], [100, "138.68.129.114", "443", "Zjc789456123?", "aes-256-cfb", "10:32:08", "GB"], [100, "sg-ipv6.free-ss.site", "11111", "43967365", "aes-128-gcm", "10:32:06", "SG"], [100, "159.89.133.233", "443", "Zjc789456123?", "aes-256-cfb", "10:32:08", "US"], [100, "us-ipv6.free-ss.site", "11111", "74707513", "aes-128-gcm", "10:32:04", "US"], [100, "128.199.109.223", "443", "Zjc789456123?", "aes-256-cfb", "10:32:10", "SG"], [100, "213.183.53.157", "80", "byebye_1218", "chacha20-ietf-poly1305", "10:32:06", "RU"], [100, "45.76.142.30", "443", "root123", "aes-256-cfb", "10:32:08", "GB"], [100, "64.137.246.174", "8989", "0123456789", "chacha20", "10:32:12", "CA"], [100, "46.101.79.76", "443", "Zjc789456123?", "aes-256-cfb", "10:32:08", "GB"], [100, "5.101.179.106", "20353", "9264392837", "aes-128-gcm", "10:32:05", "EE"], [100, "188.166.155.75", "443", "Zjc789456123?", "aes-256-cfb", "10:32:09", "GB"], [100, "104.192.81.17", "3333", "bf0a7253", "aes-256-cfb", "10:32:04", "US"], [95, "107.174.205.197", "60799", "ssmgr3193904396", "aes-256-cfb", "10:32:04", "US"], [100, "128.199.193.158", "443", "Zjc789456123?", "aes-256-cfb", "10:32:10", "SG"], [100, "64.137.248.154", "8989", "0123456789", "chacha20", "10:32:03", "CA"], [100, "138.68.176.92", "443", "Zjc789456123?", "aes-256-cfb", "10:32:09", "GB"], [100, "185.145.130.214", "4436", "^CT&zd7*Ra", "aes-256-cfb", "10:32:04", "NL"], [75, "77.81.105.31", "8088", "Zjc789456123?", "aes-256-cfb", "10:32:10", "RO"], [100, "45.32.130.112", "11111", "10733332", "aes-128-gcm", "10:32:03", "US"], [100, "64.137.237.199", "8989", "0123456789", "chacha20", "10:32:10", "CA"], [100, "45.76.161.215", "443", "root123", "aes-256-cfb", "10:32:10", "SG"], [100, "146.185.159.72", "443", "Zjc789456123?", "aes-256-cfb", "10:32:09", "NL"], [100, "128.199.76.10", "443", "Zjc789456123?", "aes-256-cfb", "10:32:10", "SG"], [100, "159.65.0.201", "2666", "QAZXSW12345PLM987", "rc4-md5", "10:32:09", "SG"], [100, "94.237.64.38", "4436", "^CT&zd7*Ra", "aes-256-cfb", "10:32:06", "SG"], [100, "213.183.48.10", "15025", "16801532", "rc4-md5", "10:32:04", "RU"], [100, "207.148.11.160", "443", "root123", "aes-256-cfb", "10:32:08", "US"], [100, "54.37.177.131", "7001", "freess.pw", "aes-256-cfb", "10:32:04", "FR"], [100, "209.250.234.209", "443", "root123", "aes-256-cfb", "10:32:08", "DE"], [90, "138.68.231.150", "443", "Zjc789456123?", "aes-256-cfb", "10:32:08", "US"], [95, "172.247.33.238", "12345", "ya2uz2et", "aes-256-cfb", "10:32:16", "US"], [100, "172.247.33.170", "12345", "ya2uz2et", "aes-256-cfb", "10:32:11", "US"], [100, "144.217.163.62", "4436", "^CT&zd7*Ra", "aes-256-cfb", "10:32:04", "CA"], [100, "168.1.149.13", "30538", "ff.tn", "aes-256-cfb", "10:32:08", "AU"], [100, "139.59.106.134", "2333", "nutgeek", "rc4-md5", "10:32:09", "IN"], [100, "94.46.164.140", "4436", "^CT&zd7*Ra", "aes-256-cfb", "10:32:05", "PT"], [100, "45.77.239.36", "443", "root123", "aes-256-cfb", "10:32:10", "AU"], [100, "104.238.180.212", "443", "root123", "aes-256-cfb", "10:32:08", "US"], [95, "64.137.236.177", "9966", "123qwe", "aes-256-cfb", "10:32:08", "CA"], [100, "185.103.110.172", "4436", "^CT&zd7*Ra", "aes-256-cfb", "10:32:05", "FI"], [100, "45.77.34.244", "11111", "37997663", "aes-128-gcm", "10:32:05", "SG"], [100, "146.185.153.205", "443", "Zjc789456123?", "aes-256-cfb", "10:32:08", "NL"], [100, "64.137.251.231", "8989", "0123456789", "chacha20", "10:32:16", "CA"], [100, "62.113.196.43", "40010", "vpnnest!@#123d", "aes-256-cfb", "10:32:04", "DE"], [100, "138.197.167.61", "2666", "QAZXSW12345PLM987", "rc4-md5", "10:32:08", "CA"], [100, "138.68.69.99", "2666", "QAZXSW12345PLM987", "rc4-md5", "10:32:09", "DE"], [100, "45.63.37.61", "443", "root123", "aes-256-cfb", "10:32:08", "US"], [100, "45.32.149.14", "443", "root123", "aes-256-cfb", "10:32:09", "FR"]]}`)

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

func readFile() []*SS {
	f, err := os.Open("ss.html")
	if err != nil {
		panic(err)
	}

	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		panic(err)
	}
	ret := []*SS{}

	xx, _ := doc.Find("body div.container").Find("div.main").Html()
	fmt.Println("this is div.main", xx)
	return nil
	//doc.Find("body div.dataTables_wrapper").Eq(1).Find("tbody tr").Each(func(i int, s *goquery.Selection) {
	doc.Find("body div.container").Find("div.main").Find("div#tbe788_wrapper.dataTables_wrapper.dt-foundation.no-footer").Find("div#tbe788.compact.stripe.hover.nowrap.dataTable.no-footer").Find("tbody tr").Each(func(i int, s *goquery.Selection) {
		test, _ := s.Html()
		fmt.Println("this is html,", test)

		tmp := &SS{}
		s.Find("td").Each(func(i int, ss *goquery.Selection) {
			switch i {
			case 0:
				tmp.Health, _ = strconv.ParseFloat(strings.Split(ss.Text(), "/")[0], 64)
			case 1:
				tmp.IP = ss.Text()
			case 2:
				tmp.Port = ss.Text()
			case 3:
				tmp.Password = ss.Text()
			case 4:
				tmp.Method = ss.Text()
			case 5:
				tmp.Verified = ss.Text()
			case 6:
				tmp.Geo = ss.Text()
			}
			if tmp.Health == 0 || tmp.IP == "" || tmp.Port == "" || tmp.Password == "" || tmp.Method == "" || tmp.Verified == "" || tmp.Geo == "" {
				return
			}
			ret = append(ret, tmp)
		})
	})
	return ret
}

func main() {
	surgeConf := flag.Bool("surge", false, "生成surge软件配置")
	limit := flag.Int("limit", 3, "要多少就输入多少，不一定有这么多")
	//backupURL := flag.Bool("backup_url", false, "防止被BAN，启用备用")
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
	//data := getData(*backupURL)
	data := readFile()
	best := []*SS{}
	var count int
	// filter just need health 100
	for _, v := range data {
		if v.Health == 10 {
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
