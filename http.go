package paggr

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
	ex "github.com/wolvex/go/error"
)

type Client interface {
	Post(version, msgid string, request *Message) (*Message, *ex.AppError)
	Get(version, msgid, params map[string]string) (*HttpResponse, *ex.AppError)
}

type HttpResponse struct {
	Status int
	Header map[string]string
	Body   string
	Length int64
	Raw    []byte
}

type HttpClient struct {
	Url        string
	OriginHost string
	Session    *http.Client
	Signer     *Signer
	Unsigners  map[string]*Unsigner
}

func NewClient(url, originHost string, signer *Signer, unsigners map[string]*Unsigner, timeout int64) *HttpClient {
	transport := &http.Transport{}
	if strings.HasPrefix(url, "https") {
		transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}
	/**
	proxyURL, err := url.Parse("http://localhost:8888")
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}*/

	return &HttpClient{
		Url:        url,
		OriginHost: originHost,
		Session: &http.Client{
			Timeout:   time.Duration(timeout) * time.Millisecond,
			Transport: transport,
		},
		Signer:    signer,
		Unsigners: unsigners,
	}
}

func (c *HttpClient) Post(body *Message) (response *Message, err *ex.AppError) {
	var req []byte
	var e error

	//if e = c.SignRequest(body); e != nil {
	if c.Signer == nil {
		err = ex.Errorc(ERR_OTHERS).Rem("Failed getting signer object")
		return
	}

	if e = c.Signer.Set(body); e != nil {
		err = ex.Error(e, ERR_OTHERS).Rem("Unable to sign the outgoing request")
		return
	}

	if req, e = json.Marshal(body); err != nil {
		err = ex.Error(e, ERR_INVALID_FORMAT).Rem("Unable to marshal request to json format")
		return
	}

	header := make(map[string]string)
	header["Content-Type"] = "application/json"
	header["X-Version"] = body.Version
	header["X-Msg-ID"] = body.MsgID
	header["X-Origin-Host"] = c.OriginHost

	if res, e := c.Submit("POST", header, req); e != nil {
		if isTimeout(e) || isEOF(e) {
			err = ex.Error(e, ERR_TIMEOUT).Rem("Timeout/EOF detected")
		} else {
			err = ex.Error(e, ERR_OTHERS).Rem("Unable to send POST to %s", c.Url)
		}
		return
	} else {
		if res.Length <= 0 {
			if res.Status != 200 && res.Status != 202 {
				err = ex.Errorc(ERR_OTHERS).Rem("Received %d http status", res.Status)
			} else {
				err = ex.Errorc(ERR_INVALID_FORMAT).Rem("Received blank or unknown response from server")
			}
			return
		} else {
			if e := json.Unmarshal(res.Raw, &response); e != nil {
				if !strings.Contains(res.Header["Content-Type"], "json") {
					err = ex.Error(e, ERR_INVALID_FORMAT).Rem("Unable to decode because response is not in json format")
					return
				}
			}

			if c.Unsigners != nil {
				//validate response signature here
			}
		}
	}
	return
}

func (c *HttpClient) Get(version, msgid string, params map[string]string) (response *HttpResponse, err *ex.AppError) {

	return
}

func (c *HttpClient) Submit(method string, header map[string]string, body []byte) (response *HttpResponse, err error) {
	//initiliaze request
	req, err := http.NewRequest(method, c.Url, bytes.NewBuffer(body))
	if err != nil {
		return
	}

	//assign headers
	//req.Header.Add("Content-Type", "application/json; charset=utf-8")
	for k, v := range header {
		req.Header.Add(k, v)
	}

	if dump, e := httputil.DumpRequestOut(req, true); e != nil {
		log.Error(e)
	} else {
		log.WithFields(log.Fields{
			"request": string(dump),
			"url":     c.Url,
			"msg_id":  header["X-Msg-ID"],
		}).Info("Sending HTTP request")
		//fmt.Printf("HTTP Request: \n %q \n", dump)
	}

	var res *http.Response
	res, err = c.Session.Do(req)
	if err != nil {
		log.WithField("error", err).Error("Exception caught while sending HTTP package")

		req.Close = true
		c.Session.CloseIdleConnections()

		return
	}
	defer res.Body.Close()

	if dump, e := httputil.DumpResponse(res, true); e != nil {
		log.Error(e)
	} else {
		log.WithFields(log.Fields{
			"response": string(dump),
			"msg_id":   header["X-Msg-ID"],
		}).Info("Receiving HTTP response")
		//fmt.Printf("HTTP Response: \n %q \n", dump)
	}

	response = &HttpResponse{
		Status: res.StatusCode,
	}
	response.Header = make(map[string]string)
	for name, value := range res.Header {
		response.Header[name] = value[0]
	}
	response.Length = res.ContentLength
	response.Raw, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	return
}

func isTimeout(err error) bool {
	if err, ok := err.(net.Error); ok && err.Timeout() {
		return true
	} else {
		return false
	}
}

func isEOF(err error) bool {
	return strings.Contains(err.Error(), "EOF")
}

/**
func (c *HttpClient) SignRequest(msg *Message) error {
	if payload, err := json.Marshal(msg.Request); err != nil {
		return err
	} else {
		signed, err := c.Signer.Sign([]byte(payload))
		if err != nil {
			return err
		}
		msg.Signature = base64.StdEncoding.EncodeToString(signed)
	}
	return nil
}
*/
