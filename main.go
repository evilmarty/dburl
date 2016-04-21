package dburl

import (
	"database/sql"
	"net/url"
	"strings"
)

var (
	formatters = map[string]func(*url.URL) string{
		"mysql": formatMySQL,
	}
)

func Open(rawurl string) (*sql.DB, error) {
	url, err := url.Parse(rawurl)
	if err != nil {
		return nil, err
	}

	driverName := strings.ToLower(url.Scheme)
	formatter := formatters[driverName]
	if formatter == nil {
		formatter = formatDefault
	}

	dataSourceName := formatter(url)

	return sql.Open(driverName, dataSourceName)
}

func formatDefault(url *url.URL) string {
	return url.String()
}

func formatMySQL(url *url.URL) string {
	url.Scheme = ""
	if url.Host != "" {
		url.Host = "tcp(" + url.Host + ")"
	}
	return url.String()[2:]
}
