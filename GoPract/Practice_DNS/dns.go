package main

import (
	"fmt"
	"net"
)

func main() {
	//DNS name
	iprecords, _ := net.LookupIP("amazon.com") //The net.LookupIP() function accepts a string(domain-name) and returns a slice of net.IP objects that contains host's IPv4 and IPv6 addresses.
	for _, ip := range iprecords {
		fmt.Println(ip)
	}

	//Canonical name (CNAME)
	cname, _ := net.LookupCNAME("m.google.com") //The net.LookupCNAME() function accepts a host-name(m.facebook.com) as a string and returns a single canonical domain name for the given host.
	fmt.Println(cname)

	//PTR (pointer)
	ptr, _ := net.LookupAddr("49.204.141.1") //performs a reverse finding for the address and returns a list of names that map to the given address.
	for _, ptrvalue := range ptr {
		fmt.Println(ptrvalue)
	}

	//Name Server(NS)
	nameserver, _ := net.LookupNS("amazon.com") //describe the authorized name servers for the zone
	for _, ns := range nameserver {
		fmt.Println(ns)
	}

	//MX Records
	mxrecords, _ := net.LookupMX("gmail.com") // identify the servers that can exchange emails.
	for _, mx := range mxrecords {
		fmt.Println(mx.Host, mx.Pref)
	}

	//SRV Service record
	cname, srvs, err := net.LookupSRV("xmpp-server", "tcp", "golang.org") // tries to resolve an SRV query of the given service, protocol, and domain name.
	if err != nil {
		panic(err)
	}

	fmt.Printf("\ncname: %s \n\n", cname)

	for _, srv := range srvs {
		fmt.Printf("%v:%v:%d:%d\n", srv.Target, srv.Port, srv.Priority, srv.Weight)
	}

	//TXT records
	txtrecords, _ := net.LookupTXT("amazon.com") //stores information about the SPF that can identify the authorized server to send email on behalf of your organization.

	for _, txt := range txtrecords {
		fmt.Println(txt)
	}
}
