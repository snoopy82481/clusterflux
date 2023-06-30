package config

type Config struct {
	Email                    string           `yaml:"email" validate:"required|email"`
	Timezone                 string           `yaml:"timezone" validate:"required"`
	AgePublicKey             string           `yaml:"agePublicKey" validate:"required"`
	WeaveGitOpsAdminPassword string           `yaml:"weaveGitOpsAdminPassword" validate:"required|password"`
	Network                  Network          `yaml:"network" validate:"required"`
	GitHub                   GitHubConfig     `yaml:"github"`
	Cloudflare               CloudflareConfig `yaml:"cloudflare"`
	Ansible                  AnsibleConfig    `yaml:"ansible"`
}

type Network struct {
	LoadBalancerRange string `yaml:"loadBalancerRange" validate:"required"`
}

type GitHubConfig struct {
	Public            bool   `yaml:"public" validate:"required|bool"`
	URL               string `yaml:"url" validate:"required|url"`
	FluxWebhookSecret string `yaml:"fluxWebhookSecret" validate:"required"`
}

type CloudflareConfig struct {
	Domain   string           `yaml:"domain" validate:"required|url"`
	APIToken string           `yaml:"apiToken" validate:"required"`
	Tunnel   CloudflareTunnel `yaml:"tunnel"`
}

type CloudflareTunnel struct {
	AccountTag   string `yaml:"accountTag" validate:"required"`
	TunnelSecret string `yaml:"tunnelSecret" validate:"required"`
	TunnelID     string `yaml:"tunnelID" validate:"required"`
}

type AnsibleConfig struct {
	Enabled                   bool          `yaml:"enable" validate:"required|bool"`
	ControlNodeHostnamePrefix string        `yaml:"controlNodeHostnamePrefix" validate:"requiredif:Enabled,true"`
	NodeHostnamePrefix        string        `yaml:"nodeHostnamePrefix" validate:"requiredif:Enabled,true"`
	Hosts                     []AnsibleHost `yaml:"hosts"`
}

type AnsibleHost struct {
	IPAddress    string `yaml:"ipAddress" validate:"requiredif:Enabled,true|ip"`
	SSHUsername  string `yaml:"sshUsername" validate:"requiredif:Enabled,true"`
	SudoPassword string `yaml:"sudoPassword" validate:"requiredif:Enabled,true"`
	ControlNode  bool   `yaml:"controlNode" validate:"requiredif:Enabled,true|bool"`
	Hostname     string `yaml:"hostname"`
}
