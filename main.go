package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/domainr/whois"
	"golang.org/x/net/publicsuffix"
)

func domainToOrg(domain string) string {
	req, err := whois.NewRequest(domain)
	if err != nil {
		return ""
	}

	res, err := whois.DefaultClient.Fetch(req)
	if err != nil {
		return ""
	}

	// find the organization name in the WHOIS response
	orgRegex := regexp.MustCompile(`(?i)Registrant Organization:\s*(.*)`)
	matches := orgRegex.FindStringSubmatch(res.String())

	if len(matches) > 1 {
		return strings.TrimSpace(matches[1])
	}

	// fallback to base domain name if no org found
	baseDomain, err := publicsuffix.EffectiveTLDPlusOne(domain)
	if err != nil {
		return ""
	}
	return strings.Title(strings.Split(baseDomain, ".")[0])
}

func main() {
	verbose := flag.Bool("v", false, "Enable verbose output")
	var file string
	flag.StringVar(&file, "f", "", "Input file containing domain names")
	flag.Parse()

	if file != "" {
		f, err := os.Open(file)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			domain := scanner.Text()
			orgName := domainToOrg(domain)
			if orgName != "" {
				if *verbose {
					fmt.Printf("%s > %s\n", domain, orgName)
				} else {
					fmt.Println(orgName)
				}
			}
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading file:", err)
		}
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			domain := scanner.Text()
			orgName := domainToOrg(domain)
			if orgName != "" {
				if *verbose {
					fmt.Printf("%s > %s\n", domain, orgName)
				} else {
					fmt.Println(orgName)
				}
			}
		}
	}
}
