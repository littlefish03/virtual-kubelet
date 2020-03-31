package main

import (
	"github.com/virtual-kubelet/virtual-kubelet/cmd/virtual-kubelet/internal/provider"
	"github.com/virtual-kubelet/virtual-kubelet/cmd/virtual-kubelet/internal/provider/mock"
        "github.com/virtual-kubelet/virtual-kubelet/cmd/virtual-kubelet/internal/provider/openstack"
)

func registerMock(s *provider.Store) {
	s.Register("mock", func(cfg provider.InitConfig) (provider.Provider, error) { //nolint:errcheck
		return mock.NewMockProvider(
			cfg.ConfigPath,
			cfg.NodeName,
			cfg.OperatingSystem,
			cfg.InternalIP,
			cfg.DaemonPort,
		)
	})
}

func registerOpenstack(s *provider.Store) {
        s.Register("openstack", func(cfg provider.InitConfig) (provider.Provider, error) { //nolint:errcheck
                return openstack.NewZunProvider(
                        cfg.ConfigPath,
                        cfg.ResourceManager,
                        cfg.NodeName,
                        cfg.OperatingSystem,
                        cfg.DaemonPort,
                )
        })
}
