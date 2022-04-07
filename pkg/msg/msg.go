// Copyright 2016 fatedier, fatedier@gmail.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package msg

import "net"

const (
	TypeLogin                 = 'o'
	TypeLoginResp             = '1'
	TypeNewProxy              = 'p'
	TypeNewProxyResp          = '2'
	TypeCloseProxy            = 'c'
	TypeNewWorkConn           = 'w'
	TypeReqWorkConn           = 'r'
	TypeStartWorkConn         = 's'
	TypeNewVisitorConn        = 'v'
	TypeNewVisitorConnResp    = '3'
	TypePing                  = 'h'
	TypePong                  = '4'
	TypeUDPPacket             = 'u'
	TypeNatHoleVisitor        = 'i'
	TypeNatHoleClient         = 'n'
	TypeNatHoleResp           = 'm'
	TypeNatHoleClientDetectOK = 'd'
	TypeNatHoleSid            = '5'
)

var (
	msgTypeMap = map[byte]interface{}{
		TypeLogin:                 Login{},
		TypeLoginResp:             LoginResp{},
		TypeNewProxy:              NewProxy{},
		TypeNewProxyResp:          NewProxyResp{},
		TypeCloseProxy:            CloseProxy{},
		TypeNewWorkConn:           NewWorkConn{},
		TypeReqWorkConn:           ReqWorkConn{},
		TypeStartWorkConn:         StartWorkConn{},
		TypeNewVisitorConn:        NewVisitorConn{},
		TypeNewVisitorConnResp:    NewVisitorConnResp{},
		TypePing:                  Ping{},
		TypePong:                  Pong{},
		TypeUDPPacket:             UDPPacket{},
		TypeNatHoleVisitor:        NatHoleVisitor{},
		TypeNatHoleClient:         NatHoleClient{},
		TypeNatHoleResp:           NatHoleResp{},
		TypeNatHoleClientDetectOK: NatHoleClientDetectOK{},
		TypeNatHoleSid:            NatHoleSid{},
	}
)

// When frpc start, client send this message to login to server.
type Login struct {
	Version      string            `json:"noisrev"`
	Hostname     string            `json:"emantsoh"`
	Os           string            `json:"so"`
	Arch         string            `json:"hcra"`
	User         string            `json:"resu"`
	PrivilegeKey string            `json:"yek_egelivirp"`
	Timestamp    int64             `json:"pmatsemit"`
	RunID        string            `json:"di_nur"`
	Metas        map[string]string `json:"satem"`

	// Some global configures.
	PoolCount int `json:"tnuoc_loop"`
}

type LoginResp struct {
	Version       string `json:"noisrev"`
	RunID         string `json:"di_nur"`
	ServerUDPPort int    `json:"trop_pdu_revres"`
	Error         string `json:"rorre"`
}

// When frpc login success, send this message to frps for running a new proxy.
type NewProxy struct {
	ProxyName      string            `json:"proxy_name"`
	ProxyType      string            `json:"proxy_type"`
	UseEncryption  bool              `json:"use_encryption"`
	UseCompression bool              `json:"use_compression"`
	Group          string            `json:"group"`
	GroupKey       string            `json:"group_key"`
	Metas          map[string]string `json:"metas"`

	// tcp and udp only
	RemotePort int `json:"remote_port"`

	// http and https only
	CustomDomains     []string          `json:"custom_domains"`
	SubDomain         string            `json:"subdomain"`
	Locations         []string          `json:"locations"`
	HTTPUser          string            `json:"http_user"`
	HTTPPwd           string            `json:"http_pwd"`
	HostHeaderRewrite string            `json:"host_header_rewrite"`
	Headers           map[string]string `json:"headers"`

	// stcp
	Sk string `json:"sk"`

	// tcpmux
	Multiplexer string `json:"multiplexer"`
}

type NewProxyResp struct {
	ProxyName  string `json:"proxy_name"`
	RemoteAddr string `json:"remote_addr"`
	Error      string `json:"error"`
}

type CloseProxy struct {
	ProxyName string `json:"proxy_name"`
}

type NewWorkConn struct {
	RunID        string `json:"di_nur"`
	PrivilegeKey string `json:"yek_egelivirp"`
	Timestamp    int64  `json:"pmatsemit"`
}

type ReqWorkConn struct {
}

type StartWorkConn struct {
	ProxyName string `json:"proxy_name"`
	SrcAddr   string `json:"src_addr"`
	DstAddr   string `json:"dst_addr"`
	SrcPort   uint16 `json:"src_port"`
	DstPort   uint16 `json:"dst_port"`
	Error     string `json:"error"`
}

type NewVisitorConn struct {
	ProxyName      string `json:"eman_yxorp"`
	SignKey        string `json:"yek_ngis"`
	Timestamp      int64  `json:"pmatsemit"`
	UseEncryption  bool   `json:"noitpyrcne_esu"`
	UseCompression bool   `json:"noisserpmoc_esu"`
}

type NewVisitorConnResp struct {
	ProxyName string `json:"eman_yxorp"`
	Error     string `json:"rorre"`
}

type Ping struct {
	PrivilegeKey string `json:"yek_egelivirp"`
	Timestamp    int64  `json:"pmatsemit"`
}

type Pong struct {
	Error string `json:"error"`
}

type UDPPacket struct {
	Content    string       `json:"c"`
	LocalAddr  *net.UDPAddr `json:"l"`
	RemoteAddr *net.UDPAddr `json:"r"`
}

type NatHoleVisitor struct {
	ProxyName string `json:"proxy_name"`
	SignKey   string `json:"sign_key"`
	Timestamp int64  `json:"pmatsemit"`
}

type NatHoleClient struct {
	ProxyName string `json:"proxy_name"`
	Sid       string `json:"sid"`
}

type NatHoleResp struct {
	Sid         string `json:"sid"`
	VisitorAddr string `json:"visitor_addr"`
	ClientAddr  string `json:"client_addr"`
	Error       string `json:"error"`
}

type NatHoleClientDetectOK struct {
}

type NatHoleSid struct {
	Sid string `json:"sid"`
}
