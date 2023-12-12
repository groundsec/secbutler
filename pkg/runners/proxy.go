package runners

import (
	"github.com/groundsec/secbutler/pkg/logger"
	"github.com/soluchok/freeproxy"
)

func GetRandomProxy() {
	// TODO: Write custom package
	gen := freeproxy.New()
	randomProxy := gen.Get()
	logger.Println(randomProxy)
}
