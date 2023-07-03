package config

type Config struct {
	Email        string `yaml:"email" validate:"required|email"`
	Timezone     string `yaml:"timezone" validate:"required"`
	AgePublicKey string `yaml:"agePublicKey" validate:"required"`
	Apps         Apps
	Network      Network          `yaml:"network" validate:"required"`
	GitHub       GitHubConfig     `yaml:"github"`
	Cloudflare   CloudflareConfig `yaml:"cloudflare"`
	Ansible      AnsibleConfig    `yaml:"ansible"`
}

type Network struct {
	ClusterCidr    string `yaml:"clusterCidr" validate:"required|cidr"`
	ServiceCidr    string `yaml:"serviceCidr" validate:"required|cidr"`
	K8sGatewayAddr string `yaml:"k8sGatewayAddr" validate:"ip"`
	IngressAddr    string `yaml:"ingressAddr" validate:"ip"`
	KubeVIPAddr    string `yaml:"kubeVIPAddr" validate:"ip"`
}

type GitHubConfig struct {
	Public      bool   `yaml:"public" validate:"required|bool"`
	URL         string `yaml:"url" validate:"required|url"`
	FluxWebhook FluxWebhook
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

type Apps struct {
	WeaveGitOps WeaveGitOps
	Grafana     Grafana
}

type WeaveGitOps struct {
	AdminPassword string `yaml:"adminPassword" validate:"required"`
}

type Grafana struct {
	AdminPassword string `yaml:"adminPassword" validate:"required"`
}

type FluxWebhook struct {
	Secret string `yaml:"secret" validate:"required"`
}
