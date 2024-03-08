package impl

import (
	"github.com/hdkef/hadoop/pkg/repository"
	pkRepoImpl "github.com/hdkef/hadoop/pkg/repository/badger"
)

func NewkeyValueRepo(storareRoot string) repository.KeyValueRepository {
	return pkRepoImpl.NewBadgerRepo(storareRoot)
}
