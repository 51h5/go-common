package utils

import "regexp"

var (
    isTmallRegex    = regexp.MustCompile(`(?i:AliApp\((?:TM|TM\-PD)/[0-9.]+\))`)
    isDingTalkRegex = regexp.MustCompile(`(?i:AliApp\((?:DingTalk)/[0-9.]+\))`)
)

func IsTmall(ua string) bool {
    return isTmallRegex.MatchString(ua)
}

func IsDingTalk(ua string) bool {
    return isDingTalkRegex.MatchString(ua)
}
