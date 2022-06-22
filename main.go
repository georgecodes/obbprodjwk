package main

import (
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"flag"
	"fmt"
	"log"
	"crypto/rsa"
	"io/ioutil"
	"crypto/sha256"
	b64 "encoding/base64"
	"strings"
)



var (
	input       = flag.String("in", "", "The public key in PEM format to load")
	alias       = flag.String("alias", "", "The KMS alias")
	outputAsJwks        = flag.Bool("jwks", true, "Output as JWKS")
)

func mustNot(err error) {
	if err != nil {
		panic(err)
	}
}

func loadCert(file string) *x509.Certificate {
	bytes, err := ioutil.ReadFile(file)
	mustNot(err)
	block, _ := pem.Decode(bytes)
	cert, err := x509.ParseCertificate(block.Bytes)
	mustNot(err)

	return cert
}

func kid(cert *x509.Certificate) string {
	h := sha256.New()
	h.Write(cert.Raw)
	return b64.URLEncoding.WithPadding(b64.NoPadding).EncodeToString(h.Sum(nil))
}

func main() {
	flag.Parse()

	if len(*input) == 0 {
		log.Fatalf("Missing required --in parameter")
	}

	if len(*alias) == 0 {
		log.Fatalf("Missing required --alias parameter")
	}

	cert := loadCert(*input)

	key := cert.PublicKey.(*rsa.PublicKey)

	v := key.E
	bs := make([]byte, 3)
	bs[0] = byte(v >> 16)
	bs[1] = byte(v >> 8)
	bs[2] = byte(v)

	n := b64.URLEncoding.WithPadding(b64.NoPadding).EncodeToString(key.N.Bytes())
	e := b64.URLEncoding.EncodeToString(bs)
	a := *alias
	if !strings.HasPrefix(a, "alias/")  {
		a = fmt.Sprintf("alias/%v", a)
	}
	kid := kid(cert)
	d := b64.URLEncoding.EncodeToString([]byte(a))

	jwk := JWK{
		Kty: "RSA",
		Alg: "PS256",
		Alias: a,
		Kid: kid,
		Use: "sig",
		N: n,
		E: e,
		D: d,
	}

	var val []byte
	var err error
	if *outputAsJwks {
		jwks := JWKS{
			Keys: []JWK{jwk},
		}
		val, err = json.MarshalIndent(jwks, "", "    ")
		mustNot(err)

	} else {
		val, err = json.MarshalIndent(jwk, "", "    ")
		mustNot(err)
	}


	fmt.Printf("%v\n", string(val))
}