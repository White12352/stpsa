package option

type ClashAPIOptions struct {
	ExternalController       string `json:"external_controller,omitempty"`
	ExternalUI               string `json:"external_ui,omitempty"`
	ExternalUIDownloadURL    string `json:"external_ui_download_url,omitempty"`
	ExternalUIDownloadDetour string `json:"external_ui_download_detour,omitempty"`
	Secret                   string `json:"secret,omitempty"`
	DefaultMode              string `json:"default_mode,omitempty"`
	StoreSelected            bool   `json:"store_selected,omitempty"`
	StoreFakeIP              bool   `json:"store_fakeip,omitempty"`
	CacheFile                string `json:"cache_file,omitempty"`
}

type Provider struct {
	Tag            string   `json:"tag"`
	URL            string   `json:"url"`
	Interval       Duration `json:"interval,omitempty"`
	CacheFile      string   `json:"cache_file,omitempty"`
	DownloadDetour string   `json:"download_detour,omitempty"`

	Exclude string `json:"exclude,omitempty"`
	Include string `json:"include,omitempty"`

	DialerOptions
}

type GroupCommonOption struct {
	Outbounds []string `json:"outbounds"`
	Providers []string `json:"providers"`
}

type SelectorOutboundOptions struct {
	GroupCommonOption
	Default string `json:"default,omitempty"`
}

type URLTestOutboundOptions struct {
	GroupCommonOption
	URL       string   `json:"url,omitempty"`
	Interval  Duration `json:"interval,omitempty"`
	Tolerance uint16   `json:"tolerance,omitempty"`
}

// HealthCheckOptions is the settings for health check
type HealthCheckOptions struct {
	Interval     Duration `json:"interval"`
	Sampling     uint     `json:"sampling"`
	Destination  string   `json:"destination"`
	Connectivity string   `json:"connectivity"`
	DetourOf     string   `json:"detour_of,omitempty"`
}
