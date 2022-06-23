package utils

import (
	"regexp"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var (
	// https://github.com/golang/lint/blob/master/lint.go#L770
	commonInitialisms = []string{"API", "ASCII", "CPU", "CSS", "DNS", "EOF", "GUID", "HTML", "HTTP", "HTTPS", "ID", "IP", "JSON", "LHS", "QPS", "RAM", "RHS", "RPC", "SLA", "SMTP", "SSH", "TLS", "TTL", "UID", "UI", "UUID", "URI", "URL", "UTF8", "VM", "XML", "XSRF", "XSS"}
	caser             = cases.Title(language.Und)
)

func ToFormatName(name string) string {
	result := strings.ReplaceAll(caser.String(strings.ReplaceAll(name, "_", " ")), " ", "")
	for _, initialism := range commonInitialisms {
		result = regexp.MustCompile(caser.String(strings.ToLower(initialism))+"([A-Z]|$|_)").ReplaceAllString(result, initialism+"$1")
	}
	return result
}
