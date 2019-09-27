package paggr

import (
	"encoding/base64"
	"encoding/json"

	crypt "github.com/wolvex/go/crypto"
)

type Sign interface {
	Set(msg *Message) error
	Get(msg *Message) error
	Check(msg *Message) error
}

type Signer struct {
	Signer crypt.Signer
}

type Unsigner struct {
	Unsigner crypt.Unsigner
}

func NewSignerFromFile(file string) (*Signer, error) {
	if signer, err := crypt.LoadPrivateKey(file); err != nil {
		return nil, err
	} else {
		return &Signer{
			Signer: signer,
		}, nil
	}
}

func NewSigner(privateKey string) (*Signer, error) {
	if signer, err := crypt.ParsePrivateKey([]byte(privateKey)); err != nil {
		return nil, err
	} else {
		return &Signer{
			Signer: signer,
		}, nil
	}
}

func NewUnsignerFromFile(file string) (*Unsigner, error) {
	if unsigner, err := crypt.LoadPublicKey(file); err != nil {
		return nil, err
	} else {
		return &Unsigner{
			Unsigner: unsigner,
		}, nil
	}
}

func NewUnsigner(publicKey string) (*Unsigner, error) {
	if unsigner, err := crypt.ParsePublicKey([]byte(publicKey)); err != nil {
		return nil, err
	} else {
		return &Unsigner{
			Unsigner: unsigner,
		}, nil
	}
}

func (c *Signer) Set(msg *Message) error {
	if payload, err := json.Marshal(msg.Request); err != nil {
		return err
	} else {
		signed, err := c.Signer.Sign(payload)
		if err != nil {
			return err
		}
		//fmt.Printf("Generated sign: %s\n", base64.StdEncoding.EncodeToString(signed))
		msg.Signature = base64.StdEncoding.EncodeToString(signed)
	}
	return nil
}

func (c *Signer) Get(msg string) (string, error) {
	if signed, err := c.Signer.Sign([]byte(msg)); err != nil {
		return "", err
	} else {
		return base64.StdEncoding.EncodeToString(signed), nil
	}
}

func (c *Unsigner) Check(msg string, sign string) error {
	if s, e := base64.StdEncoding.DecodeString(sign); e != nil {
		return e
	} else {
		return c.Unsigner.Unsign([]byte(msg), s)
	}
}

func (c *Unsigner) CheckRequest(msg Message) error {
	//if payload, err := json.Marshal(msg.Payload); err != nil {
	//	return err
	//} else {
	payload := msg.Payload
	//fmt.Println(string(payload))
	//fmt.Println(msg.Signature)
	if sign, err := base64.StdEncoding.DecodeString(msg.Signature); err != nil {
		return err
	} else {
		return c.Unsigner.Unsign(payload, sign)
	}
	//}
}

func (c *Unsigner) CheckResponse(msg Message) error {
	if payload, err := json.Marshal(msg.Response); err != nil {
		return err
	} else {
		if sign, err := base64.StdEncoding.DecodeString(msg.Signature); err != nil {
			return err
		} else {
			return c.Unsigner.Unsign(payload, []byte(sign))
		}
	}
}
