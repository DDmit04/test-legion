package host

import (
	"fmt"
	"net"
	"regexp"

	exceptions2 "github.com/DDmit04/test-legion/src/internal/models/exceptions"
	"github.com/DDmit04/test-legion/src/internal/services/readers"
)

const (
	ipv6Regex   = `^(([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))$`
	ipv4Regex   = `^(((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.|$)){4})`
	domainRegex = `^(?:[a-z0-9](?:[a-z0-9-]{0,61}[a-z0-9])?\.)+[a-z0-9][a-z0-9-]{0,61}[a-z0-9]$`
)

type BaseHostReader struct {
	readers.BaseReader
}

func (r *BaseHostReader) validateHost(host string) *exceptions2.ModuleException {
	match, _ := regexp.MatchString(ipv4Regex+`|`+ipv6Regex+`|`+domainRegex, host)
	if !match {
		return exceptions2.InvalidHost(fmt.Sprintf(": %s", host))
	}
	return nil
}

func (r *BaseHostReader) validateIp(ip string) (net.IP, *exceptions2.ModuleException) {
	res := net.ParseIP(ip)
	if res == nil {
		return nil, exceptions2.InvalidIp(fmt.Sprintf(": %s", ip))
	}
	return res, nil
}

func (r *BaseHostReader) extractDomainIps(domain string) ([]net.IP, error) {
	var res []net.IP
	ips, err := net.LookupIP(domain)
	if err != nil {
		return nil, exceptions2.HostIpsReadError(fmt.Sprintf("error host - \"%s\"", domain))
	}
	for _, ip := range ips {
		if ipv4 := ip.To4(); ipv4 != nil {
			res = append(res, ipv4)
		} else {
			res = append(res, ip)
		}
	}
	return res, nil
}
