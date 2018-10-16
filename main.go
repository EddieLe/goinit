package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/mikemintang/go-curl"
	"github.com/spf13/viper"
)

var defaultYaml = []byte(`
	SERVER_HOST: "http://127.0.0.1"
	SERVER_PORT: ":8888"
	`)

func main() {
	//先使用ENV
	viper.AutomaticEnv()
	//yaml設定
	viper.SetConfigType("yaml")
	var configFile string
	flag.StringVar(&configFile, "c", "config.yaml", "a string var")

	if configFile != "" {
		content, err := ioutil.ReadFile(configFile)
		if err != nil {
			content = defaultYaml
		}

		viper.ReadConfig(bytes.NewBuffer(content))
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName("config")
		if err := viper.ReadInConfig(); err == nil {
			fmt.Println("Using config file:", viper.ConfigFileUsed())
		} else {
			// load default config
			viper.ReadConfig(bytes.NewBuffer(defaultYaml))
		}
	}
	// code := flag.Int("code", 0, "code")
	var (
		systemToken        string
		gameAdminToken     string
		gameToken          string
		platformAdminToken string
		groupID            string
	)
	flag.StringVar(&systemToken, "systemToken", "", "systemToken")
	flag.StringVar(&gameAdminToken, "gameAdminToken", "", "gameAdminToken")
	flag.StringVar(&gameToken, "gameToken", "", "gameToken")
	flag.StringVar(&platformAdminToken, "platformAdminToken", "", "platformAdminToken")
	flag.StringVar(&groupID, "groupID", "", "groupID")
	flag.Parse()

	// fmt.Println("systemToken:", systemToken)
	// fmt.Println("gameAdminToken:", gameAdminToken)
	// fmt.Println("gameToken:", gameToken)
	// fmt.Println("platformAdminToken:", platformAdminToken)
	name := []string{
		"eddie",
		"yaomin",
		"max",
		"evelyn",
		"rdekyle",
		"cooper",
		"james",
		"edwin",
		"cchang",
		"beckham",
		"ken",
		"chiawen",
		"miya",
		"sean",
		"linyen",
		"front1",
		"front2",
		"front3",
		"front4",
		"front5",
		"front6",
		"front7",
		"front8",
		"front9",
		"front0",
		"leo",
		"shawn",
		"tony",
		"jingan",
		"jennifer",
		"rain",
		"jim",
		"eddielee"}

	if systemToken != "" {
		addGameAdminUser(systemToken)
	} else if gameAdminToken != "" {
		for _, v := range name {
			addGameUser(gameAdminToken, v, groupID)
		}
	} else if gameToken != "" {
		addPlatformAdminUser(gameToken)
	} else if platformAdminToken != "" {
		for _, v := range name {
			addPlatformUser(platformAdminToken, v, groupID)
		}
	}
}

func addGameAdminUser(token string) {
	addGameAdminUserApi := viper.GetString("SERVER_HOST") + viper.GetString("SERVER_PORT") + "/api/account/v2/system/gameAdmin"

	headers := map[string]string{
		"validity_token": token,
	}

	postData := map[string]interface{}{
		"username":          "gameAdmin",
		"password":          "54321",
		"businessCode":      "INRD2",
		"currency":          "CNY",
		"availableCurrency": []string{"CNY", "USD"},
		"timezone":          "zh-TW",
	}
	req := curl.NewRequest()
	resp, err := req.
		SetUrl(addGameAdminUserApi).
		SetHeaders(headers).
		SetPostData(postData).
		Post()

	if err != nil {
		fmt.Println(err)
	} else {
		if resp.IsOk() {
			fmt.Println(resp.Body)
		} else {
			fmt.Println(resp.Raw)
			fmt.Println(resp.Body)
		}
	}
}

func addGameUser(token string, username string, groupID string) {
	addGameUserApi := viper.GetString("SERVER_HOST") + viper.GetString("SERVER_PORT") + "/api/account/v2/game"

	headers := map[string]string{
		"validity_token": token,
	}
	postData := map[string]interface{}{
		"username":    username,
		"password":    "54321",
		"groupIDList": []string{groupID},
	}
	req := curl.NewRequest()
	resp, err := req.
		SetUrl(addGameUserApi).
		SetHeaders(headers).
		SetPostData(postData).
		Post()

	if err != nil {
		fmt.Println(err)
	} else {
		if resp.IsOk() {
			fmt.Println(resp.Body)
		} else {
			fmt.Println(resp.Raw)
			fmt.Println(resp.Body)
		}
	}
}

func addPlatformAdminUser(token string) {
	addPlatAdminUserApi := viper.GetString("SERVER_HOST") + viper.GetString("SERVER_PORT") + "/api/account/v2/game/platformAdmin"
	headers := map[string]string{
		"validity_token": token,
	}
	postData := map[string]interface{}{
		"username":          "platformAdmin",
		"password":          "54321",
		"businessCode":      "BBOS",
		"gameManagerID":     1,
		"currency":          "CNY",
		"availableCurrency": []string{"CNY", "USD"},
		"timezone":          "zh-TW",
		"lobbyDomain":       "1",
	}
	req := curl.NewRequest()
	resp, err := req.
		SetUrl(addPlatAdminUserApi).
		SetHeaders(headers).
		SetPostData(postData).
		Post()

	if err != nil {
		fmt.Println(err)
	} else {
		if resp.IsOk() {
			fmt.Println(resp.Body)
		} else {
			fmt.Println(resp.Raw)
			fmt.Println(resp.Body)
		}
	}
}

func addPlatformUser(token string, username string, groupID string) {
	addPlatUserApi := viper.GetString("SERVER_HOST") + viper.GetString("SERVER_PORT") + "/api/account/v2/platform"
	headers := map[string]string{
		"validity_token": token,
	}
	postData := map[string]interface{}{
		"username":    username,
		"password":    "54321",
		"groupIDList": []string{groupID},
	}
	req := curl.NewRequest()
	resp, err := req.
		SetUrl(addPlatUserApi).
		SetHeaders(headers).
		SetPostData(postData).
		Post()

	if err != nil {
		fmt.Println(err)
	} else {
		if resp.IsOk() {
			fmt.Println(resp.Body)
		} else {
			fmt.Println(resp.Raw)
			fmt.Println(resp.Body)
		}
	}
}
