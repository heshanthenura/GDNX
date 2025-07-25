# GDNX

A lightweight DNS server written in GO

## 🔧 Description

This is a custom DNS server built using the [miekg/dns](https://github.com/miekg/dns) library.  

## 🚧 Status

**Currently under development.**

## 📦 Dependencies

- [Go](https://golang.org/)
- [miekg/dns](https://github.com/miekg/dns)

```bash
go get github.com/miekg/dns
````

## 🛠️ Run the DNS Server

```bash
go run main.go
```

Make sure to run with sufficient privileges to bind port 53 (or use port forwarding).