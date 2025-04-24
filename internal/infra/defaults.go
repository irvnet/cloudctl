package infra

type ComponentConfig struct {
	ServerType string
	Image      string
	Location   string
	SSHKey     string
	Labels     map[string]string
}

var ComponentDefaults = map[string]ComponentConfig{
	"devbox": {
		ServerType: "cax11", Image: "ubuntu-22.04", Location: "fsn1", SSHKey: "virtuser-bluefin",
		Labels: map[string]string{"role": "devbox"},
	},
	"ctrl": {
		ServerType: "cax11", Image: "ubuntu-22.04", Location: "fsn1", SSHKey: "virtuser-bluefin",
		Labels: map[string]string{"role": "ctrl"},
	},
	"worker": {
		ServerType: "cax11", Image: "ubuntu-22.04", Location: "fsn1", SSHKey: "virtuser-bluefin",
		Labels: map[string]string{"role": "worker"},
	},
}
