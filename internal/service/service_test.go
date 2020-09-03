package service

import (
	"context"
	"flag"
	"os"
	"testing"

	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/testing/lich"
)

var (
	testService *Service
	ctx         = context.Background()
)

func TestMain(m *testing.M) {
	flag.Set("conf", "../../test")
	//flag.Set("f", "../../test/docker-compose.yaml")
	flag.Parse()
	os.Setenv("DISABLE_LICH", "true")
	disableLich := os.Getenv("DISABLE_LICH") != ""
	if !disableLich {
		if err := lich.Setup(); err != nil {
			panic(err)
		}
	}
	var err error
	if err = paladin.Init(); err != nil {
		panic(err)
	}
	var cf func()
	if testService, cf, err = newTestService(); err != nil {
		panic(err)
	}

	ret := m.Run()
	if cf != nil {
		cf()
	}

	if !disableLich {
		_ = lich.Teardown()
	}
	os.Exit(ret)
}
