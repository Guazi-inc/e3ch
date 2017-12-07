package client

import (
	"strings"

	"github.com/coreos/etcd/auth/authpb"
	"github.com/coreos/etcd/clientv3"
)

func (clt *EtcdHRCHYClient) permPath(key string) (string, error) {
	if key != "0" && !strings.HasPrefix(key, "/") {
		return "", ErrorInvalidKey
	}

	if clt.rootKey != "" {
		key = "/" + clt.rootKey + key
	}
	return key, nil
}

func (clt *EtcdHRCHYClient) RoleGrantPermission(name string, key, rangeEnd string, ty clientv3.PermissionType) error {
	key, err := clt.permPath(key)
	if err != nil {
		return err
	}

	// rangeEnd == "" means only set key
	if rangeEnd != "" {
		rangeEnd, err = clt.permPath(rangeEnd)
		if err != nil {
			return err
		}
	}

	_, err = clt.client.RoleGrantPermission(clt.ctx, name, key, rangeEnd, ty)
	return err
}

type Perm struct {
	PermType string `json:"perm_type"`
	Key      string `json:"key"`
	RangeEnd string `json:"range_end"`
}

func (clt *EtcdHRCHYClient) GetRolePerms(name string) ([]*Perm, error) {
	resp, err := clt.client.RoleGet(clt.ctx, name)
	if err != nil {
		return nil, err
	}

	perms := []*Perm{}
	for _, p := range resp.Perm {
		perm := &Perm{
			Key:      clt.trimRootKey(string(p.Key)),
			RangeEnd: clt.trimRootKey(string(p.RangeEnd)),
			PermType: authpb.Permission_Type_name[int32(p.PermType)],
		}

		perms = append(perms, perm)
	}
	return perms, nil
}

func (clt *EtcdHRCHYClient) RoleRevokePermission(name string, key, rangeEnd string) error {
	key, err := clt.permPath(key)
	if err != nil {
		return err
	}

	if rangeEnd != "" {
		rangeEnd, err = clt.permPath(rangeEnd)
		if err != nil {
			return err
		}
	}

	_, err = clt.client.RoleRevokePermission(clt.ctx, name, key, rangeEnd)
	return err
}
