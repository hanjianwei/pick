package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os/exec"
	"strings"

	"github.com/miekg/dns"
)

func readURL(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	content, err := ioutil.ReadAll(res.Body)
	return content, err
}

func runCommand(cmdStr string) ([]byte, error) {
	cmd := exec.Command("bash", "-c", cmdStr)
	return cmd.Output()
}

func trimRADBOutput(out string) []string {
	lines := strings.Split(out, "\n")

	if len(lines) > 1 {
		return strings.Split(lines[1], " ")
	}

	return nil
}

func fetchAWS() []string {
	// Fetch IP ranges
	url := "https://ip-ranges.amazonaws.com/ip-ranges.json"
	content, err := readURL(url)
	if err != nil {
		log.Fatal(err)
	}

	// Parse JSON
	var res struct {
		SyncToken  string `json:"syncToken"`
		CreateDate string `json:"createDate"`
		Prefixes   []struct {
			IPPrefix string `json:"ip_prefix"`
			Region   string `json:"region"`
			Service  string `json:"service"`
		} `json:"prefixes"`
	}
	if err := json.Unmarshal(content, &res); err != nil {
		log.Fatal(err)
	}

	// Extract IP ranges
	inets := make([]string, len(res.Prefixes))
	for i, inet := range res.Prefixes {
		inets[i] = inet.IPPrefix
	}

	return inets
}

func fetchCF() []string {
	url := "https://www.cloudflare.com/ips-v4"
	content, err := readURL(url)
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(string(content), "\n")
}

func fetchDomain(domain string) net.IP {
	m := new(dns.Msg)
	m.SetQuestion(domain, dns.TypeA)

	in, err := dns.Exchange(m, "8.8.8.8:53")
	if err != nil {
		log.Fatal(err)
	}

	return in.Answer[0].(*dns.A).A
}

// Query ASN information from radb.
// See: http://www.radb.net/support/query2.php
func fetchCompany(company string) []string {
	out, err := runCommand(fmt.Sprintf("whois -h whois.radb.net '!i%s,1'", company))
	if err != nil {
		log.Fatal(err)
	}

	return trimRADBOutput(string(out))
}

func fetchASN(asn string) []string {
	out, err := runCommand(fmt.Sprintf("whois -h whois.radb.net '!g%s'", asn))
	if err != nil {
		log.Fatal(err)
	}

	return trimRADBOutput(string(out))
}
