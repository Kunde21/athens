package memory

import (
	"github.com/gomods/athens/pkg/storage"
)

type Lister struct{}

func (l *Lister) List(basePath, module string) ([]string, error) {
	key := entries.key(basePath, module)
	entries.RLock()
	defer entries.RUnlock()
	versions, ok := entries.versions[key]
	if !ok {
		return nil, storage.NotFoundErr{BasePath: basePath, Module: module}
	}
	ret := make([]string, len(versions))
	for i, version := range versions {
		ret[i] = version.RevInfo.Version
	}
	return ret, nil
}
