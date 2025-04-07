package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/url"

	"os"
	"sort"
	"strings"
)

var TG_BOT_TOKEN = os.Getenv("TG_BOT_TOKEN")

var TG_BOT_TOKEN_HMAC = generateHMACSHA256(TG_BOT_TOKEN, []byte("WebAppData"))

type UserData struct {
	AllowsWriteToPm bool   `json:"allows_write_to_pm"`
	FirstName       string `json:"first_name"`
	ID              int64  `json:"id"`
	LanguageCode    string `json:"language_code"`
	LastName        string `json:"last_name"`
	PhotoUrl        string `json:"photo_url"`
	Username        string `json:"username"`
}

func toHash(val []byte) string {
	return hex.EncodeToString(val)
}

func generateHMACSHA256(message string, key []byte) []byte {
	messageBytes := []byte(message)
	hmacSHA256 := hmac.New(sha256.New, key)
	hmacSHA256.Write(messageBytes)
	return hmacSHA256.Sum(nil)
}

func VerifyInitData(initData string) (bool, *url.Values) {
	init_data_query, err := url.ParseQuery(initData)

	if err != nil {
		return false, nil
	}

	var check_str_tokens []string = []string{}

	var hash_str string

	for key, val := range init_data_query {
		if key == "hash" {
			hash_str = val[0]
			continue
		}
		check_str_tokens = append(check_str_tokens, fmt.Sprintf("%s=%s", key, val[0]))
	}

	sort.Slice(check_str_tokens, func(i, j int) bool {
		return check_str_tokens[i] < check_str_tokens[j]
	})

	check_str := strings.Join(check_str_tokens, "\n")

	user_data_hmac := generateHMACSHA256(check_str, TG_BOT_TOKEN_HMAC)

	dataValid := toHash(user_data_hmac) == hash_str

	if dataValid {
		return dataValid, &init_data_query
	}

	return dataValid, nil
}

func ParseUserDataFromQuery(initDataQuery *url.Values) (*UserData, error) {
	var userData UserData
	err := json.Unmarshal([]byte(initDataQuery.Get("user")), &userData)
	if err != nil {
		return nil, err
	}

	return &userData, nil
}
