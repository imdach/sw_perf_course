package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func sendRequest() {

	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("POST", "http://192.168.31.197:8000/route0", nil)

	// Headers
	req.Header.Add("Content-Type", "application/pdf")
	auth := "GA1.2.1928432377.1606920046; MEIQIA_TRACK_ID=1wLP26KD1raftAexEOW5QnqqReF; 1wLP26KD1raftAexEOW5QnqqReF=undefined; Hm_lvt_f83d162eeb1abea371d41d4b60d345db=1633961833; LF_ID=1639560954040-1147233-9739072; GCID=5c2289f-7eff658-ec91eac-2ed7f7f; GRID=5c2289f-7eff658-ec91eac-2ed7f7f; MEIQIA_VISIT_ID=268xi5wtCOFDNKUjxehW1d3YioF; gksskpitn=d302190a-a56c-4f51-8e84-eb0dc497cd48; Hm_lvt_59c4ff31a9ee6263811b23eb921a5083=1646268115,1646819131,1647914637,1648200878; Hm_lvt_022f847c4e3acd44d4a2481d9187f1e6=1646268116,1646819131,1647914638,1648200878; GCESS=BgYEq0yVlwUEAAAAAA0BAQcEKiw25goEAAAAAAgBAwsCBgACBLiMPWIBCGZHDwAAAAAAAwS4jD1iCQEBBAQALw0ADAEB; sensorsdata2015jssdkcross=%7B%22distinct_id%22%3A%221001318%22%2C%22first_id%22%3A%2217ed2f655d920f-09db8a329f749f-133a6253-1296000-17ed2f655db9c3%22%2C%22props%22%3A%7B%22%24latest_traffic_source_type%22%3A%22%E7%9B%B4%E6%8E%A5%E6%B5%81%E9%87%8F%22%2C%22%24latest_search_keyword%22%3A%22%E6%9C%AA%E5%8F%96%E5%88%B0%E5%80%BC_%E7%9B%B4%E6%8E%A5%E6%89%93%E5%BC%80%22%2C%22%24latest_referrer%22%3A%22%22%2C%22%24latest_landing_page%22%3A%22https%3A%2F%2Ftime.geekbang.org%2F%22%2C%22%24latest_utm_source%22%3A%22u_nav_web%22%2C%22%24latest_utm_medium%22%3A%22u_nav_web%22%2C%22%24latest_utm_term%22%3A%22u_nav_web%22%2C%22%24latest_utm_campaign%22%3A%22timewebmenu%22%2C%22%24latest_utm_content%22%3A%22menu%22%7D%2C%22%24device_id%22%3A%2217623e60541670-0903496e91057d-18356153-1296000-17623e60542624%22%7D; _gid=GA1.2.877925763.1648615646; _gat=1; gk_process_ev={%22count%22:45%2C%22utime%22:1648200884250%2C%22referrer%22:%22https://time.geekbang.org/%22%2C%22target%22:%22%22%2C%22referrerTarget%22:%22page_geektime_login%22}; Hm_lpvt_59c4ff31a9ee6263811b23eb921a5083=1648624436; Hm_lpvt_022f847c4e3acd44d4a2481d9187f1e6=1648624437; SERVERID=3431a294a18c59fc8f5805662e2bd51e|1648624437|1648621706\r\n"
	req.Header.Add("Authorization", auth)

	// Fetch Request
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Failure : ", err)
	}

	// Read Response Body
	respBody, _ := ioutil.ReadAll(resp.Body)

	// Display Results
	fmt.Println("response Status : ", resp.Status)
	fmt.Println("response Headers : ", resp.Header)
	fmt.Println("response Body : ", string(respBody))
}

func main() {
	sendRequest()
}
