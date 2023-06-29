package config

type Config struct {
	Core       CoreConfig       `yaml:"core"`
	GitHub     GitHubConfig     `yaml:"github"`
	Cloudflare CloudflareConfig `yaml:"cloudflare"`
	Ansible    AnsibleConfig    `yaml:"ansible"`
}

type CoreConfig struct {
	MetalLBRange             string `yaml:"metalLBRange" validate:"required|ipRange"`
	AgePublicKey             string `yaml:"agePublicKey" validate:"required"`
	Timezone                 string `yaml:"timezone" validate:"required"`
	WeaveGitOpsAdminPassword string `yaml:"weaveGitOpsAdminPassword" validate:"required|password"`
}

type GitHubConfig struct {
	URL               string `yaml:"url" validate:"required|url"`
	FluxWebhookSecret string `yaml:"fluxWebhookSecret" validate:"required"`
}

type CloudflareConfig struct {
	Domain string           `yaml:"domain" validate:"required|url"`
	Email  string           `yaml:"email" validate:"required|email"`
	APIKey string           `yaml:"apiKey" validate:"required"`
	Tunnel CloudflareTunnel `yaml:"tunnel"`
}

type CloudflareTunnel struct {
	AccountTag   string `yaml:"accountTag" validate:"required"`
	TunnelSecret string `yaml:"tunnelSecret" validate:"required"`
	TunnelID     string `yaml:"tunnelID" validate:"required"`
}

type AnsibleConfig struct {
	ControlNodeHostnamePrefix string        `yaml:"controlNodeHostnamePrefix" validate:"required"`
	NodeHostnamePrefix        string        `yaml:"nodeHostnamePrefix" validate:"required"`
	Hosts                     []AnsibleHost `yaml:"hosts"`
}

type AnsibleHost struct {
	IPAddress    string `yaml:"ipAddress" validate:"required|ip"`
	SSHUsername  string `yaml:"sshUsername" validate:"required"`
	SudoPassword string `yaml:"sudoPassword" validate:"required"`
	ControlNode  bool   `yaml:"controlNode" validate:"required|bool"`
	Hostname     string `yaml:"hostname"`
}
