package main

import (
	"fmt"
	"log"

	"github.com/miekg/dns"
)

const routerDNS = "192.168.1.1:53"

func handle(w dns.ResponseWriter, r *dns.Msg) {
	m := new(dns.Msg)
	m.SetReply(r)

	if len(r.Question) > 0 {
		q := r.Question[0]
		fmt.Printf("[INFO] %s %s\n", dns.TypeToString[q.Qtype], q.Name)

		if q.Qtype == dns.TypeA && q.Name == "redirect.me." {
			rr, _ := dns.NewRR("redirect.me. 3600 IN A 192.168.1.101")
			m.Answer = append(m.Answer, rr)
			w.WriteMsg(m)
			return
		}

		if q.Name == "tpc.googlesyndication.com." {
			fmt.Println("\033[31m[BLOCKED] tpc.googlesyndication.com\033[0m")
			rr, _ := dns.NewRR("tpc.googlesyndication.com. 60 IN A 0.0.0.0")
			m.Answer = append(m.Answer, rr)
			w.WriteMsg(m)
			return
		}

		if q.Name == "pagead2.googlesyndication.com." {
			fmt.Println("\033[31m[BLOCKED] pagead2.googlesyndication.com\033[0m")
			rr, _ := dns.NewRR("pagead2.googlesyndication.com. 60 IN A 0.0.0.0")
			m.Answer = append(m.Answer, rr)
			w.WriteMsg(m)
			return
		}

		if q.Name == "chat.qwen.ai." {
			fmt.Println("\033[31m[BLOCKED] chat.qwen.ai - NXDOMAIN\033[0m")
			m.Rcode = dns.RcodeNameError // NXDOMAIN
			w.WriteMsg(m)
			return
		}

	}

	c := new(dns.Client)
	resp, _, err := c.Exchange(r, routerDNS)
	if err != nil {
		log.Printf("Failed to forward query to router: %v", err)
		dns.HandleFailed(w, r)
		return
	}

	w.WriteMsg(resp)
}

func main() {
	dns.HandleFunc(".", handle)

	go func() {
		log.Println("Starting DNS server on UDP :53")
		if err := dns.ListenAndServe(":53", "udp", nil); err != nil {
			log.Fatalf("Failed to start UDP server: %v", err)
		}
	}()

	go func() {
		log.Println("Starting DNS server on TCP :53")
		if err := dns.ListenAndServe(":53", "tcp", nil); err != nil {
			log.Fatalf("Failed to start TCP server: %v", err)
		}
	}()

	select {}
}
