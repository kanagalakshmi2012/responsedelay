package main
import (
	"fmt"
	"github.com/miekg/dns"
	"os"
	"time"
)

func queryWithAccessControl(name string, qtype uint16, server string, nodes int) *dns.Msg {
	client := new(dns.Client)
	client.Net = "udp"
	msg := new(dns.Msg)
	msg.SetQuestion(dns.Fqdn(name), qtype)
	delay := accessControlDelay(nodes)
	time.Sleep(delay)
	resp, _, err := client.Exchange(msg, server)
	if err != nil {
		fmt.Println("DNS query error:", err)
		os.Exit(1)
	}
	return resp
}

func accessControlDelay(nodes int) time.Duration {
	baseDelay := 50 * time.Millisecond
	increment := 20 * time.Millisecond

	if nodes <= 0 {
		return 0
	}
	return baseDelay + time.Duration(nodes-1)*increment
}

func main() {
	server := "****:53"
	domain := "**."
	clusterNodes := 7

	response := queryWithAccessControl(domain, dns.TypeA, server, clusterNodes)
	for _, ans := range response.Answer {
		fmt.Println(ans)
	}
}
