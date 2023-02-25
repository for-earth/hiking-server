package mountain

import (
	"io/ioutil"
	"net/http"
	"strings"

	xj "github.com/basgys/goxml2json"
	"github.com/gofiber/fiber/v2"
)

type Mountain struct {
	No      string `xml:"mntilistno"`
	Name    string `xml:"mntiname"`
	SubName string `xml:"mntisname"`
	Summary string `xml:"mntisummary"`
	Address string `xml:"mntiadd"`
	Details string `xml:"mntidetails"`
	Hight   string `xml:"mntihigh"`
}

type Mountains struct {
	NumOfRows  string     `xml:"numOfRows"`
	PageNo     string     `xml:"pageNo"`
	TotalCount string     `xml:"totalCount"`
	Mountains  []Mountain `xml:"items"`
}

const baseURL string = "https://apis.data.go.kr/1400000/service/cultureInfoService2/mntInfoOpenAPI2"
const serviceKey string = "bb3C3v5zF6A%2FocQ998DVgk9DnmxhnqJ7y14o1sWHVSiFRiiXRlehpj%2BxRLDIJoAxOV5tvjzZRIitea7RFASH0A%3D%3D"

func GetMountains(c *fiber.Ctx) error {
	client := &http.Client{}

	req, err := http.NewRequest("GET", baseURL+"?"+"ServiceKey="+serviceKey+"&"+"numOfRows="+"15", nil)
	if err != nil {
		return err
	}

	// send the request and get the response
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	// read the response body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	// close the response body
	defer res.Body.Close()

	// var mountain Mountain
	// var mountains Mountains

	xml := strings.NewReader(string(body))
	json, err := xj.Convert(xml)
	if err != nil {
		panic("That's embarrassing...")
	}

	// set the response content type and send the JSON data back to the client
	c.Type("application/json")
	return c.JSON(json.String())
}
