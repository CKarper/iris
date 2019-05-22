package spec

type (
	Filter struct {
		Name string `yaml:"name"`
	}

	Integration struct {
		Name         string   `yaml:"name"`
		Integrations []string `yaml:"integrations"`
		Destinations []string `yaml:"destination"`
	}

	Destination struct {
		Name string `yaml:"name"`
		URL  string `yaml:"url"`
	}

	Spec struct {
		Filters      []Filter      `yaml:"filters"`
		IntegrationS []Integration `yaml:"integration"`
		Destinations []Destination `yaml:"destination"`
	}
)
