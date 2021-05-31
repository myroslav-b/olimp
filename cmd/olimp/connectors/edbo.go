package connectors

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/myroslav-b/olimp/cmd/olimp/catalogs"
)

type tEdboLoader struct {
}

type TEdbo struct {
	url    url.URL
	client *http.Client
}

var schemeEdebo = "http"
var hostEdbo = "registry.edbo.gov.ua"
var pathEdbo = map[string]string{
	"1": "api/universities",
	"2": "api/universities",
	"3": "api/institutions",
	"5": "api/universities",
	"9": "api/universities",
}
var expEdebo = "json"

var transportEdbo = &http.Transport{
	//Proxy: func( *http.Request) (*url.URL, error) {
	//},
	//DialContext: func(ctx context.Context, network string, addr string) (net.Conn, error) {
	//},
	//Dial: func(network string, addr string) (net.Conn, error) {
	//},
	//DialTLSContext: func(ctx context.Context, network string, addr string) (net.Conn, error) {
	//},
	//DialTLS: func(network string, addr string) (net.Conn, error) {
	//},
	//TLSClientConfig:        &tls.Config{},
	//TLSHandshakeTimeout:    0,
	//DisableKeepAlives:      false,
	//DisableCompression:     false,
	//MaxIdleConns:           0,
	//MaxIdleConnsPerHost:    0,
	//MaxConnsPerHost:        0,connectors.tEdboLoader
	//TLSNextProto:           map[string]func(authority string, c *tls.Conn) http.RoundTripper{},
	//ProxyConnectHeader:     map[string][]string{},
	//MaxResponseHeaderBytes: 0,
	//WriteBufferSize:        0,
	//ReadBufferSize:         0,
	//ForceAttemptHTTP2:      false,
}

var clientEdbo = &http.Client{
	Transport: transportEdbo,
	//CheckRedirect: func(req *http.Request, via []*http.Request) error {
	//},
	//Jar:     nil,
	Timeout: 30 * time.Second,
}

func buildUrl(instType, regCode string) url.URL {
	u := url.URL{
		Scheme:      schemeEdebo,
		Opaque:      "",
		User:        nil,
		Host:        hostEdbo,
		Path:        pathEdbo[instType],
		RawPath:     "",
		ForceQuery:  false,
		RawQuery:    fmt.Sprintf("ut=%s&lc=%s&exp=%s", instType, regCode, expEdebo),
		Fragment:    "anchor",
		RawFragment: "",
	}
	return u
}

//ErrPathEdbo - error "Error of path to EDBO"
var ErrPathEdbo = errors.New("Error of path to EDBO")

//ErrRespNot200 - error "The request returned a status code other than 200"
//var ErrRespNot200 = errors.New("The request returned a status code other than 200")

func new(client *http.Client, instType, regCode string) (*TEdbo, error) {
	var edbo TEdbo
	if !checkPathEdebo() {
		return &edbo, ErrPathEdbo
	}
	edbo.url = buildUrl(instType, regCode)
	edbo.client = client

	return &edbo, nil
}

func checkPathEdebo() bool {
	for t := range pathEdbo {
		_, ok := catalogs.MapInstType()[t]
		if !ok {
			return false
		}
	}
	for t := range catalogs.MapInstType() {
		_, ok := pathEdbo[t]
		if !ok {
			return false
		}
	}
	return true
}

func request(edbo *TEdbo) ([]byte, error) {
	req, err := http.NewRequest("GET", edbo.url.String(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := edbo.client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, errors.New(resp.Status)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()

	return body, nil
}

func parse(bytes []byte, tags map[string]string) ([]map[string]string, error) {
	arrMap := make([]map[string]json.RawMessage, 0)
	err := json.Unmarshal(bytes, &arrMap)
	if err != nil {
		return nil, err
	}
	instBundles := make([]map[string]string, len(arrMap))
	for iBundl, rawBundl := range arrMap {
		var m = make(map[string]string)
		for field, tag := range tags {
			var s string
			err = json.Unmarshal(rawBundl[tag], &s)
			if err != nil {
				return nil, err
			}
			m[field] = s
		}
		instBundles[iBundl] = m
	}

	return instBundles, nil
}

func (edboLoader tEdboLoader) LoadBatch(instType, regCode string) ([]map[string]string, error) {
	edbo, err := new(clientEdbo, instType, regCode)
	if err != nil {
		log.Print("> ", err)
		return nil, err
	}

	bytes, err := request(edbo)
	if err != nil {
		log.Print(">> ", err)
		return nil, err
	}

	tags := catalogs.MapEdboFieldTags(instType)
	rez, err := parse(bytes, tags)
	if err != nil {
		log.Print(">>> ", err)
		return nil, err
	}

	return rez, nil
}

func NewEdboLoader() tEdboLoader {
	var edboLoader tEdboLoader
	return edboLoader
}
