package decimal

import (
	"encoding/json"
	"math/rand"
	"strconv"
	"testing"
)
import "fmt"

/*
	输入一个10进制的数和一个想要转换的进制，输出转换后的结果
*/

type x int

func (tmp x) xxxx() int {
	tmp = 1
	return 0
}

func (tmp *x) yyyy() int {
	var a x = 100
	tmp = &a
	fmt.Println(&a)
	fmt.Println(tmp)
	return 1
}

func TestX(t *testing.T) {
	var xx x = 2
	a := xx.xxxx()
	fmt.Println("--------", a)
	fmt.Println(xx)

	xx.yyyy()
	fmt.Println(&xx)
}

type T struct {
	name string
}

func (tmp T) fun1() {
	tmp.name = "李四"
}

func (tmp *T) fun2() {
	tmp.name = "王五"
}

func TestT(t *testing.T) {
	tmp := T{name: "张三"}
	tmp.fun1()
	fmt.Println(tmp)

	tmp.fun2()
	fmt.Println(tmp)

	//tmp2 := &T{name:"张三"}
	//tmp2.fun1()
	//fmt.Println(tmp2)
	//
	//tmp2.fun2()
	//fmt.Println(tmp2)
}

func TestDecimalChange(t *testing.T) {
	decimalChange(100, 6)
}

func TestReserveString(t *testing.T) {
	fmt.Println("翻转字符串之后的结果为：", reserveString("g1od"))
}

func TestRand(t *testing.T) {

	var a string
	fmt.Println(a)
	fmt.Println("=============")
	ran := rand.New(rand.NewSource(516987700626817125))
	fmt.Println(ran.Float32())
}

type tmp struct {
	id   int
	name string
}

type CaInfoPart struct {
	Lat                float64 `json:"lat"`                  // 纬度
	Lng                float64 `json:"lng"`                  // 经度
	LocationCityID     int     `json:"location_city_id"`     // GPS城市ID
	LocationDistrictID int     `json:"location_district_id"` // GPS区县ID
	CityID             int     `json:"city_id"`              // 客户选择城市ID（类型int）
	SiteCityID         int     `json:"site_city_id"`         // 客户选择站点ID（类型int)
}

func TestPoint(t *testing.T) {
	caInfo := `{"ca_s":"pz_sogou", "ca_n":"pz_bt", "ca_medium":"-", "ca_term":"-", "ca_content":"", "ca_campaign":"", "ca_kw":"-", "ca_i":"-", "scode":"-", "keyword":"-", "ca_keywordid":"-", "ca_transid":"", "platform":"2", "version":1, "track_id":"8312984297971712", "guid":"ac07f38a-4d05-47ee-c5a4-1725b5470874", "display_finance_flag":"-", "client_ab":"-", "sessionid":"214c5302-227e-4c6f-8fb6-dbcc71668e35", "ca_a":"-", "ca_b":"-", "customer_type":3, "location_city_id":0, "location_district_id":0, "city_id":146, "site_city_id":146, "lat":39.7736, "lng":-1, "coupon_ab":"0"}`
	var caInfoPart CaInfoPart
	err := json.Unmarshal([]byte(caInfo), &caInfoPart)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(fmt.Sprintf("%+v", caInfoPart))
	fmt.Println(strconv.FormatFloat(caInfoPart.Lng, 'f', -1, 64))
	type xxx struct {
		name string
	}
	var data struct {
		Distance float64 `json:"distance"`
		Data     *xxx
	}

	data.Data = (*xxx)(nil)
	fmt.Println(data.Distance)
	fmt.Println(data.Data.name)

	var _ = ([]int)(nil) == nil
}
