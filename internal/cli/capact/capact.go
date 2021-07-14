package capact

import (
	"time"

	"github.com/pkg/errors"
)

const (
	// LatestVersionTag tag used to select latest version
	LatestVersionTag = "@latest"
	// LocalVersionTag tag used to select local charts and images
	LocalVersionTag = "@local"
	// LocalDockerTag tag used when building local images
	LocalDockerTag = "dev"
	// LocalDockerPath path used when building local images
	LocalDockerPath = "local"
	// KindEnv default name for kind environment
	KindEnv = "kind"

	// LocalChartsPath path to Helm charts in Capact repo
	LocalChartsPath = "deploy/kubernetes/charts"
	// HelmRepoLatest URL of the latest Capact charts repository
	HelmRepoLatest = "https://storage.googleapis.com/capactio-latest-charts"
	// HelmRepoStable URL of the stable Capact charts repository
	HelmRepoStable = "https://storage.googleapis.com/capactio-stable-charts"

	// CRDUrl Capact CRD URL
	CRDUrl = "https://raw.githubusercontent.com/capactio/capact/main/deploy/kubernetes/crds/core.capact.io_actions.yaml"

	// Name Capact name
	Name = "capact"
	// Namespace Capact default namespace to install
	Namespace = "capact-system"

	// RepositoryCache Helm cache for repositories
	RepositoryCache = "/tmp/helm"

	// CertFile Capact Gateway certificate file name
	CertFile = "capact-local-ca.crt"
	// LinuxCertsPath path to Linux certificates directory
	LinuxCertsPath = "/usr/local/share/ca-certificates"
)

// Options to set when interacting wit Capact
type Options struct {
	Name               string
	Namespace          string
	Environment        string
	SkipComponents     []string
	SkipImages         []string
	FocusImages        []string
	DryRun             bool
	Timeout            time.Duration
	Parameters         InputParameters
	UpdateHostsFile    bool
	UpdateTrustedCerts bool
	Verbose            bool
}

// Validate validates provided options
func (o Options) Validate() error {
	if len(o.SkipImages) != 0 && len(o.FocusImages) != 0 {
		return errors.New("cannot skip and focus images at the same time")
	}

	return nil
}