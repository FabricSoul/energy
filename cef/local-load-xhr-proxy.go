//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package cef

import (
	"bytes"
	"crypto/tls"
	"embed"
	"errors"
	. "github.com/energye/energy/v2/consts"
	"github.com/energye/energy/v2/logger"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strconv"
)

// cookies jar
var jar, _ = cookiejar.New(nil)

// IXHRProxy
//  本地资源加载 XHR 请求代理接口
type IXHRProxy interface {
	Send(request *ICefRequest) (*XHRProxyResponse, error) // 发送请求，在浏览器进程同步执行
}

// XHRProxy
//  数据请求代理
type XHRProxy struct {
	Scheme LocalProxyScheme // http/https/tcp default: http
	IP     string           // default: localhost
	Port   int              // default: 80
	SSL    XHRProxySSL
	client *httpClient
}

// XHRProxySSL
//  https证书配置，如果其中某一配置为空，则跳过ssl检查, 如果证书配置错误则请求失败
type XHRProxySSL struct {
	FS        *embed.FS // 证书到内置执行文件时需要设置
	RootDir   string    // 根目录, FS不为空时是内置资源目录名，否则是本地硬盘目录
	SSLCert   string    // RootDir/to/path/cert.crt
	SSLKey    string    // RootDir/to/path/key.key
	SSLCARoot string    // RootDir/to/path/ca.crt
}

type httpClient struct {
	tr     *http.Transport
	client http.Client
}

// XHRProxyResponse
//  代理响应数据
type XHRProxyResponse struct {
	Data       []byte              // 响应数据
	DataSize   int                 // 响应数据大小
	StatusCode int32               // 响应状态码
	Header     map[string][]string // 响应头
}

func (m *XHRProxySSL) skipVerify() bool {
	return m.RootDir == "" || m.SSLCert == "" || m.SSLKey == "" || m.SSLCARoot == ""
}

func (m *XHRProxy) Send(request *ICefRequest) (*XHRProxyResponse, error) {
	if m.Scheme == LpsHttp {
		return m.sendHttp(request)
	} else if m.Scheme == LpsHttps {
		return m.sendHttps(request)
	} else if m.Scheme == LpsTcp {
		return m.sendTcp(request)
	}
	return nil, errors.New("incorrect scheme")
}

func (m *XHRProxy) sendHttp(request *ICefRequest) (*XHRProxyResponse, error) {
	reqUrl, err := url.Parse(request.URL())
	if err != nil {
		return nil, err
	}
	targetUrl := new(bytes.Buffer)
	targetUrl.WriteString("http://")
	targetUrl.WriteString(m.IP)
	if m.Port > 0 {
		targetUrl.WriteString(":")
		targetUrl.WriteString(strconv.Itoa(m.Port))
	}
	targetUrl.WriteString(reqUrl.Path)
	targetUrl.WriteString(reqUrl.RawQuery)
	// 读取请求数据
	requestData := new(bytes.Buffer)
	postData := request.GetPostData()
	if postData.IsValid() {
		dataCount := int(postData.GetElementCount())
		elements := postData.GetElements()
		for i := 0; i < dataCount; i++ {
			element := elements.Get(uint32(i))
			switch element.GetType() {
			case PDE_TYPE_EMPTY:
			case PDE_TYPE_BYTES:
				if byt, c := element.GetBytes(); c > 0 {
					requestData.Write(byt)
				}
			case PDE_TYPE_FILE:
				if f := element.GetFile(); f != "" {
					if byt, err := ioutil.ReadFile(f); err == nil {
						requestData.Write(byt)
					}
				}
			}
			element.Free()
		}
		postData.Free()
	}
	logger.Debug("XHRProxy TargetURL:", targetUrl.String(), "method:", request.Method(), "dataLength:", len(requestData.Bytes()))
	httpRequest, err := http.NewRequest(request.Method(), targetUrl.String(), requestData)
	if err != nil {
		return nil, err
	}
	// 设置请求头
	header := request.GetHeaderMap()
	if header.IsValid() {
		size := header.GetSize()
		for i := 0; i < int(size); i++ {
			key := header.GetKey(uint32(i))
			c := header.FindCount(key)
			for j := 0; j < int(c); j++ {
				value := header.GetEnumerate(key, uint32(j))
				httpRequest.Header.Add(key, value)
			}
		}
		header.Free()
	}
	// 创建 client
	cli := &http.Client{
		Jar: jar,
	}
	cli.CloseIdleConnections()
	httpResponse, err := cli.Do(httpRequest)
	if err != nil {
		return nil, err
	}
	defer httpResponse.Body.Close()
	// 读取响应头
	responseHeader := make(map[string][]string)
	for key, value := range httpResponse.Header {
		for _, vs := range value {
			if header, ok := responseHeader[key]; ok {
				responseHeader[key] = append(header, vs)
			} else {
				responseHeader[key] = []string{vs}
			}
		}
	}
	// 读取响应数据
	buf := new(bytes.Buffer)
	c, err := buf.ReadFrom(httpResponse.Body)
	result := &XHRProxyResponse{
		Data:       buf.Bytes(),
		DataSize:   int(c),
		StatusCode: int32(httpResponse.StatusCode),
		Header:     responseHeader,
	}
	return result, nil
}

func (m *XHRProxy) sendHttps(request *ICefRequest) (*XHRProxyResponse, error) {
	//cert, err := tls.LoadX509KeyPair(m.SSLCert, m.SSLKey)
	//if err != nil {
	//	return nil, err
	//}
	//crt, err := ioutil.ReadFile("ca.crt")
	//if err != nil {
	//	return nil, err
	//}
	//certPool := x509.NewCertPool()
	//certPool.AppendCertsFromPEM(crt)
	//if m.client == nil {
	//	m.client = &httpClient{}
	//	var config *tls.Config
	//	if m.SSL.skipVerify() {
	//
	//	} else {
	//
	//	}
	//
	//	m.client.tr = &http.Transport{
	//		TLSClientConfig: &tls.Config{
	//			InsecureSkipVerify: true,
	//			//Certificates:       []tls.Certificate{cert},
	//			//RootCAs:            certPool,
	//		},
	//	}
	//}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
			//Certificates:       []tls.Certificate{cert},
			//RootCAs:            certPool,
		},
	}
	// 创建 client
	cli := &http.Client{
		Jar:       jar,
		Transport: tr,
	}
	reqUrl, err := url.Parse(request.URL())
	if err != nil {
		return nil, err
	}
	targetUrl := new(bytes.Buffer)
	targetUrl.WriteString("https://")
	targetUrl.WriteString(m.IP)
	if m.Port > 0 {
		targetUrl.WriteString(":")
		targetUrl.WriteString(strconv.Itoa(m.Port))
	}
	targetUrl.WriteString(reqUrl.Path)
	targetUrl.WriteString(reqUrl.RawQuery)
	// 读取请求数据
	requestData := new(bytes.Buffer)
	postData := request.GetPostData()
	if postData.IsValid() {
		dataCount := int(postData.GetElementCount())
		elements := postData.GetElements()
		for i := 0; i < dataCount; i++ {
			element := elements.Get(uint32(i))
			switch element.GetType() {
			case PDE_TYPE_EMPTY:
			case PDE_TYPE_BYTES:
				if byt, c := element.GetBytes(); c > 0 {
					requestData.Write(byt)
				}
			case PDE_TYPE_FILE:
				if f := element.GetFile(); f != "" {
					if byt, err := ioutil.ReadFile(f); err == nil {
						requestData.Write(byt)
					}
				}
			}
			element.Free()
		}
		postData.Free()
	}
	logger.Debug("XHRProxy TargetURL:", targetUrl.String(), "method:", request.Method(), "dataLength:", len(requestData.Bytes()))
	httpRequest, err := http.NewRequest(request.Method(), targetUrl.String(), requestData)
	if err != nil {
		return nil, err
	}
	// 设置请求头
	header := request.GetHeaderMap()
	if header.IsValid() {
		size := header.GetSize()
		for i := 0; i < int(size); i++ {
			key := header.GetKey(uint32(i))
			c := header.FindCount(key)
			for j := 0; j < int(c); j++ {
				value := header.GetEnumerate(key, uint32(j))
				httpRequest.Header.Add(key, value)
			}
		}
		header.Free()
	}
	//httpRequest.Header.Add("Host", "energy.yanghy.cn")
	//httpRequest.Header.Add("Origin", "https://energy.yanghy.cn")
	//httpRequest.Header.Add("Referer", "https://energy.yanghy.cn/")

	httpResponse, err := cli.Do(httpRequest)
	if err != nil {
		return nil, err
	}
	defer httpResponse.Body.Close()
	// 读取响应头
	responseHeader := make(map[string][]string)
	for key, value := range httpResponse.Header {
		for _, vs := range value {
			if header, ok := responseHeader[key]; ok {
				responseHeader[key] = append(header, vs)
			} else {
				responseHeader[key] = []string{vs}
			}
		}
	}
	// 读取响应数据
	buf := new(bytes.Buffer)
	c, err := buf.ReadFrom(httpResponse.Body)
	result := &XHRProxyResponse{
		Data:       buf.Bytes(),
		DataSize:   int(c),
		StatusCode: int32(httpResponse.StatusCode),
		Header:     responseHeader,
	}
	return result, nil
}

func (m *XHRProxy) tcpListen() {

}

func (m *XHRProxy) sendTcp(request *ICefRequest) (*XHRProxyResponse, error) {
	return nil, errors.New("tcp unrealized")
}
