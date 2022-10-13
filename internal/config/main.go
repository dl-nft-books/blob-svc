package config

import (
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/copus"
	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/kit/kv"
)

type Config interface {
	comfig.Logger
	//pgdb.Databaser
	types.Copuser
	comfig.Listenerer
	MimeTypesConfigurator
	AWSConfigurator
	JWTConfigurator
}

type config struct {
	comfig.Logger
	//pgdb.Databaser
	types.Copuser
	comfig.Listenerer
	getter kv.Getter
	MimeTypesConfigurator
	AWSConfigurator
	JWTConfigurator
}

func New(getter kv.Getter) Config {
	return &config{
		getter: getter,
		//Databaser:  pgdb.NewDatabaser(getter),
		Copuser:               copus.NewCopuser(getter),
		Listenerer:            comfig.NewListenerer(getter),
		Logger:                comfig.NewLogger(getter, comfig.LoggerOpts{}),
		MimeTypesConfigurator: NewMimeTypesConfigurator(getter),
		AWSConfigurator:       NewAWSConfigurator(getter),
		JWTConfigurator:       NewJWTConfigurator(getter),
	}
}