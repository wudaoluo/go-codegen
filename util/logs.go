package util

import (
	"github.com/spf13/viper"
	"github.com/wudaoluo/golog"
	"github.com/wudaoluo/golog/conf"
)

func InitLogs() {
	golog.SetLogger(
		golog.ZAPLOG,
		conf.WithLogType(conf.LogNormalType),
		conf.WithProjectName("go-codegen"),
		conf.WithLogType(conf.LogJsontype),
		conf.WithFilename(viper.GetString("default.logfile")),
	)
}

func FlushLogs() {
	_ = golog.Sync() //ignore error
}
