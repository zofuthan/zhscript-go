package zhscript

import (
	"crypto/md5"
	"encoding/hex"
)

var codecache_ map[string]*codes___ = make(map[string]*codes___)

func codecache__(code []byte) (*codes___, string) {
	md := md5.New()
	md.Write(code)
	id := hex.EncodeToString(md.Sum(nil))
	if o_liucheng_ {
		o_n__()
		o__(0, "%s", id)
	}
	
	codes, ok := codecache_[id]
	if o_liucheng_ {
		if !ok {
			o__(0, " new")
		}
		o_n__()
	}
	return codes, id
}

func codecache2__(code []byte, id string, up_qv *Qv___) (*codes___, *Errinfo___) {
	codes := new(codes___)
	err2 := codes.z2__([]rune(string(code)), up_qv)
	if err2 != nil {
		return nil, err2
	}
	codecache_[id] = codes
	return codes, nil
}
