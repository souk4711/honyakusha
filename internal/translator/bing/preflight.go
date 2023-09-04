package bing

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"

	"github.com/go-resty/resty/v2"

	"github.com/souk4711/honyakusha/internal/conf"
)

func preflight(client *resty.Client, conf conf.ConfTranslator) (map[string]string, error) {
	fail := func(message string) (map[string]string, error) {
		return map[string]string{}, errors.New(message)
	}

	html := ""
	resp, err := client.R().Get(preflightURI(conf))
	if err != nil {
		return fail(err.Error())
	} else {
		html = string(resp.Body())
	}

	RE_IG := regexp.MustCompile("IG:\"([^\"]+)\"")
	IG := RE_IG.FindStringSubmatch(html)[1]
	if IG == "" {
		return fail("Failed to extract IG")
	}

	RE_IID := regexp.MustCompile("data-iid=\"([^\"]+)\"")
	IID := RE_IID.FindStringSubmatch(html)[1]
	if IID == "" {
		return fail("Failed to extract IID")
	}

	RE_TOKEN_PAYLOAD := regexp.MustCompile(`params_AbusePreventionHelper\s?=\s?([^\]]+\])`)
	TOKEN_PAYLOAD := RE_TOKEN_PAYLOAD.FindStringSubmatch(html)[1]
	if TOKEN_PAYLOAD == "" {
		return fail("Failed to extract params_AbusePreventionHelper")
	}

	var TOKENS []interface{}
	if err := json.Unmarshal([]byte(TOKEN_PAYLOAD), &TOKENS); err != nil {
		return fail("Failed to extract key/token")
	}
	if len(TOKENS) != 3 {
		return fail("Failed to extract key/token")
	}
	KEY, _ := TOKENS[0].(float64)
	TOKEN, _ := TOKENS[1].(string)
	TOKEN_EXP, _ := TOKENS[2].(float64)

	client.SetHeader("Referer", preflightURI((conf)))
	for _, header := range resp.Header()["Set-Cookie"] {
		client.SetHeader("Cookie", header)
	}

	return map[string]string{
		"ig":       IG,
		"iid":      IID,
		"key":      fmt.Sprintf("%.f", KEY),
		"token":    TOKEN,
		"tokenExp": fmt.Sprintf("%.f", TOKEN_EXP),
	}, nil
}

func preflightURI(conf conf.ConfTranslator) string {
	if conf.URI == "" {
		return "https://www.bing.com/translator"
	} else {
		return conf.URI + "/translator"
	}
}
