package helpers

import (
	"LazarusPoC/configs"

	"github.com/go-resty/resty/v2"
)

var headers = map[string]string{
	"accept":             "application/json",
	"X-Appwrite-key":     configs.APPWRITE_KEY,
	"X-Appwrite-Project": configs.APPWRITE_APP_ID,
}

var Resty = resty.New().SetHostURL(configs.APPWRITE_BASE_URL).SetHeaders(headers)
