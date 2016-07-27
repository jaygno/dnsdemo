// Copyright 2016 The Jaygno Authors. All rights reserved.
// A demo send dns request to nameserver.

package main

import (
        "fmt"
        "net"
        "errors"
       )

func readDNSResponse(c net.Conn) (*dnsMsg, error) {
	b := make([]byte, 512) // see RFC 1035
	n, err := c.Read(b)
	if err != nil {
		return nil, err
	}
	msg := &dnsMsg{}
	if !msg.Unpack(b[:n]) {
		return nil, errors.New("cannot unmarshal DNS message")
	}
	return msg, nil
}

func writeDNSQuery(c net.Conn, msg *dnsMsg) error {
	b, ok := msg.Pack()
	if !ok {
		return errors.New("cannot marshal DNS message")
	}
	if _, err := c.Write(b); err != nil {
		return err
	}
	return nil
}

func queryDNS(name string, conn net.Conn) {
    out := dnsMsg{
		dnsMsgHdr: dnsMsgHdr{
			recursion_desired: true,
		},
		question: []dnsQuestion{
			{name, 1, dnsClassINET},
		},
	}
    if err := writeDNSQuery(conn, &out); err != nil {
		return
	}
    in, _ := readDNSResponse(conn)
    fmt.Println(in)
}

func main() {
    // 创建连接
    conn, err := net.DialUDP("udp4", nil, &net.UDPAddr{IP: net.IPv4(114,114,114,114), Port: 53,})
    if err != nil {
        fmt.Println("连接失败!", err)
        return
    }
    defer conn.Close()

    queryDNS("www.baidu.com", conn)
}
