package main
import (
	"fmt"
	"github.com/miekg/dns"
	"os"
)
func queryDNS(name string, qtype uint16, server string) *dns.Msg {
	client := new(dns.Client)
	client.Net = "udp"
	msg := new(dns.Msg)
	msg.SetQuestion(dns.Fqdn(name), qtype)
	msg.SetEdns0(4096, true)
.	resp, _, err := client.Exchange(msg, server)
	if err != nil {
		fmt.Println("DNS query error:", err)
		os.Exit(1)
	}
	return resp
}
func main() {
	server := "8.8.8.8:53"
	domain := "example.com."
	aResp := queryDNS(domain, dns.TypeA, server)
	fmt.Println("A Record Response:")
	for _, ans := range aResp.Answer {
		fmt.Println(ans)
	}
	fmt.Println("Authenticated Data (AD bit):", aResp.AuthenticatedData)
	fmt.Println("\nRequesting DNSKEY Record with DNSSEC:")
	dnskeyResp := queryDNS(domain, dns.TypeDNSKEY, server)
	for _, rr := range dnskeyResp.Answer {
		fmt.Println(rr)
	}
	fmt.Println("\nRequesting RRSIG Record for A Record:")
	rrsigResp := queryDNS(domain, dns.TypeRRSIG, server)
	for _, rr := range rrsigResp.Answer {
		fmt.Println(rr)
	}
}
