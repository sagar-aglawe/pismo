package providers

type IProvider interface {
	New(c *ProviderConfig)
	IsInitiated() bool
}

type Registry struct {
	Config    *ProviderConfig
	Providers map[string]IProvider
}

var GlobalRegistry = Registry{
	Providers: make(map[string]IProvider),
}

func New() Registry {
	GlobalRegistry.Config = &ProviderConfig{
		PostgresClientConfig: getPostgresConfig(),
	}

	return GlobalRegistry
}

func (r *Registry) Get(providerName string) IProvider {
	if provider, exists := r.Providers[providerName]; exists {
		return provider
	}
	return nil
}

func (r *Registry) GetOrResolve(providerName string) IProvider {
	if provider, exists := r.Providers[providerName]; exists {
		if provider.IsInitiated() {
			return provider
		}
	}

	r.Resolve(providerName)
	return r.Get(providerName)

}

func (r *Registry) Set(providerName string, provider IProvider) {
	r.Providers[providerName] = provider
}

func (r *Registry) Resolve(providerNames ...string) {
	for _, providerName := range providerNames {
		providerClient := r.Get(providerName)
		if !providerClient.IsInitiated() {
			providerClient.New(r.Config)
		}
	}
}
