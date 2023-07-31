package utils

import "encoding/base64"

type TLSKey struct {
	CA   []byte
	Cert []byte
	Key  []byte
}

func KeyDecode(ca, cert, key string) (TLSKey, error) {
	var (
		tlsKey TLSKey
		err    error
	)
	if tlsKey.CA, err = base64.StdEncoding.DecodeString(ca); err != nil {
		return tlsKey, err
	}
	if tlsKey.Cert, err = base64.StdEncoding.DecodeString(cert); err != nil {
		return tlsKey, err
	}
	if tlsKey.Key, err = base64.StdEncoding.DecodeString(key); err != nil {
		return tlsKey, err
	}
	return tlsKey, nil
}
