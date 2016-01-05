package myip

import "regexp"

func ParseIPFromRemoteAddr(remoteAddr string) string {
	r, _ := regexp.Compile(`(.*):[0-9]+`)
	ip := r.FindStringSubmatch(remoteAddr)
	if len(ip) > 1 {
		return ip[1]
	} else {
		return ""
	}
}
