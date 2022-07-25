package read

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type result struct {
	CurrentTime string `json:"currentTime"`
	CurrentTime2 string `json:"currentTime2"`
	ReturnMsg string `json:"returnMsg"`
	Code string `json:"code"`
	SubCode string `json:"subCode"`
}

func Connect() {
	client := &http.Client{}
	req,err := http.NewRequest(http.MethodGet,"http://api.m.jd.com/client.action?functionId=queryMaterialProducts&client=wh5",nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("user-agent","Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36")
	req.AddCookie(&http.Cookie{Name:"__jdu", Value:"162038205346622430613"})
	req.AddCookie(&http.Cookie{Name:"shshshfpa", Value:"53614689-9b1f-d708-b5b2-5aadda216aaa-1620382144"})
	req.AddCookie(&http.Cookie{Name:"pinId", Value:"DbF2aakPjokAJZhnWZQBQw"})
	req.AddCookie(&http.Cookie{Name:"jcap_dvzw_fp", Value:"mhIBEP9hu7i4lvQsgguPDcqisRljrjtAKMI0Ipa45iBFvTLHYOD38DKVR-K-oa-EZngZdg=="})
	req.AddCookie(&http.Cookie{Name:"pin", Value:"15873143784_p"})
	req.AddCookie(&http.Cookie{Name:"_pst", Value:"15873143784_p"})
	req.AddCookie(&http.Cookie{Name:"__jda", Value:"76161171.162038205346622430613.1620382053.1654567414.1654853559.141"})
	req.AddCookie(&http.Cookie{Name:"__jdv", Value:"76161171|direct|-|none|-|1654567413727"})
	req.AddCookie(&http.Cookie{Name:"shshshfp", Value:"9d85afe1f8a8697a4e2bdab643d0434c"})
	req.AddCookie(&http.Cookie{Name:"3AB9D23F7A4B3C9B", Value:"BCDNDNKT2RWZLWGCGDKPHMRM6GPOQHFDOENDZEVGGQCSZDNQWFUE754AFAPVU652EHFYVY5ZHL4TZY446ZJWPY4S5Q"})
	req.AddCookie(&http.Cookie{Name:"user-key", Value:"b4ce7a6d-a88c-4be3-8db2-87b199fdae4d"})
	req.AddCookie(&http.Cookie{Name:"shshshfpb", Value:"c5nU3Fy8I16V0xY5%2BUjgmzw%3D%3D"})
	req.AddCookie(&http.Cookie{Name:"unick", Value:"jd158731iqx"})
	req.AddCookie(&http.Cookie{Name:"_tp", Value:"r5Bxl0uVX6LvlA5pQMUecg%3D%3D"})
	req.AddCookie(&http.Cookie{Name:"whwswswws", Value:""})
	req.AddCookie(&http.Cookie{Name:"areaId", Value:"18"})
	req.AddCookie(&http.Cookie{Name:"ipLoc-djd", Value:"18-1482-48942-49058"})
	req.AddCookie(&http.Cookie{Name:"TrackID", Value:"1A8_V7joKtkTp0gLPs-QCPrbUMBaluAcMZj4qetWwRc-DXfdF6IVNixWLPss9Mw-ALNSDLr-z-ZvdUjQfXwZyVR0haQn0K_q4sv2fVB7WFa8"})
	resp,_ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	var res result
	_ = json.Unmarshal(body,&res)
	fmt.Printf("%#v", res)
}
