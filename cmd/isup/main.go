package isup

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/dchest/validator"
	"github.com/goware/urlx"
)

func checkDNS(domain string) {
	ips, err := net.LookupIP(domain)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not get IPs: %v\n", err)
		os.Exit(1)
	}

	for _, ip := range ips {
		fmt.Printf("%s. IN A %s\n", domain, ip.String())
	}
}

func getIp(domain string) {
	ips, _ := net.LookupIP(domain)
	for _, ip := range ips {
		if ipv4 := ip.To4(); ipv4 != nil {
			fmt.Println("IPv4: ", ipv4)
		}
	}
}

func getUrl(domain string) {
	url, _ := urlx.Parse(domain)
	normalizedDomain, _ := urlx.Normalize(url)
	ip, _ := urlx.Resolve(url)

	response, errors := http.Head(normalizedDomain)

	if errors != nil {
		log.Fatal(errors)
	}

	result := "Domain [" + ip.IP.String() + "] " + normalizedDomain + " reposned with status " + response.Status
	if response.StatusCode >= 400 {
		result = result + " NOT REACHEBLE ❌"
	} else {
		result = result + " REACHEBLE ✅"
	}

	fmt.Println(result)
}

func IsUp(domain string) {
	log.Println("Processing: ", domain)
	isDomainValid := validator.IsValidDomain(domain)
	if !isDomainValid {
		domainNotValid := errors.New("domain not valid")
		log.Fatal(domainNotValid)
	}

	checkDNS(domain)
	getIp(domain)
	getUrl(domain)
}
