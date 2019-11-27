package utils

import "regexp"

var (
    isTmallRegex = regexp.MustCompile(`(?i:AliApp\((?:TM|TM\-PD)/[0-9.]+\))`)
)

func IsTmall(ua string) bool {
    return isTmallRegex.MatchString(ua)
}
