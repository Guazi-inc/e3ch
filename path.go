package client

import (
	"path"
	"strings"
)

func isRoot(key string) bool {
	return key == "/"
}

func checkRootKey(rootKey string) bool {
	return !strings.Contains(rootKey, "/")
	//return rootKey != "" && !strings.HasSuffix(rootKey, "/")
}

// ensure key, return (realKey, parentKey)
func (clt *EtcdHRCHYClient) ensureKey(key string) (r1 string, r2 string, err error) {
	//if !strings.HasPrefix(key, "/") {
	//	return "", "", ErrorInvalidKey
	//}

	if isRoot(key) {
		return "/", "/", nil
		//return clt.rootKey, clt.rootKey, nil
	} else {
		realKey := WithRootKey(clt.rootKey, key)
		return realKey, path.Clean(realKey + "/../"), nil
	}
}

func WithRootKey(rootKey, key string) string {
	key = strings.Trim(key, "/")
	//key = strings.TrimPrefix(key, "/")
	//key = strings.TrimSuffix(key, "/")
	ret := rootKey + "/" + key
	if rootKey != "" {
		ret = "/" + ret
	}
	return ret
}
