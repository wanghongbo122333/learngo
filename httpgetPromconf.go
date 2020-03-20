package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Config struct {
	Ipv4   string `json:"ipv4"`
	Port   int    `json:"port"`
	Region string `json:"region"`
	Tags   string `json:"tags"`
	Role   string `json:"role"`
}
type Prom struct {
	Conf []Config `json:"conf"`
}

func main() {
	r := gin.Default()
	r.GET("/conf", getPrometheusConf)
	r.Run()
}

func getPrometheusConf(c *gin.Context) {
	b, _ := ioutil.ReadFile("prom.json")
	fmt.Print(string(b))
	var res Prom
	err := json.Unmarshal(b, &res)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println("**Unmarshal result**")
	fmt.Println(&res)

	c.JSON(http.StatusOK, gin.H{
		"RetCode": 0,
		"Msg":     "success",
		"prom":    res,
	})

}