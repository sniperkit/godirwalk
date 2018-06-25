package cli

import (
	"time"

	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

type logFields = logrus.Fields

var (
	log = logrus.New()
)

func NewLogger() {
	log.Formatter = new(prefixed.TextFormatter)
	log.Level = logrus.DebugLevel
	log.Debugln("Started cli logger at", time.Now().Format("2015-06-16-0431 UTC"))
}
