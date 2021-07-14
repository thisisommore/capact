package cmd

import (
	"fmt"
	"os"
	"time"

	"capact.io/capact/internal/cli"
	"capact.io/capact/internal/cli/environment/create"
	"capact.io/capact/internal/cli/heredoc"

	"capact.io/capact/internal/cli/capact"

	"capact.io/capact/internal/cli/install"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

// NewInstall returns a cobra.Command for installing Capact in the created env
func NewInstall() *cobra.Command {
	opts := capact.Options{}

	installCmd := &cobra.Command{
		Use:   "install [OPTIONS]",
		Short: "install Capact into a given environment",
		Long:  "Use this command to install the Capact version in the environment.",
		Example: heredoc.WithCLIName(`
			# Install latest Capact version from main branch
			<cli> install

			# Install Capact 0.1.0 version
			<cli> install --version 0.1.0

			# Install Capact from local git repository. Needs to be run from the main directory
			<cli> install --version @local`, cli.Name),
		RunE: func(cmd *cobra.Command, args []string) error {
			k8sCfg, err := config.GetConfig()
			if err != nil {
				return errors.Wrap(err, "while creating k8s config")
			}

			err = opts.Validate()
			if err != nil {
				return errors.Wrap(err, "while validating options")
			}

			return install.Install(cmd.Context(), os.Stdout, k8sCfg, opts)
		},
	}

	flags := installCmd.Flags()

	flags.StringVar(&opts.Parameters.Version, "version", capact.LatestVersionTag, "Capact version. Possible values @latest, @local, 0.3.0, ...")
	flags.StringVar(&opts.Name, "name", create.KindDefaultClusterName, "Cluster name, overrides config.")
	flags.StringVar(&opts.Namespace, "namespace", capact.Namespace, "Capact namespace.")
	flags.StringVar(&opts.Environment, "environment", capact.KindEnv, "Capact environment.")
	flags.StringSliceVar(&opts.SkipComponents, "skip-component", []string{}, "Components names that should not be installed. Takes comma-separated list.")
	flags.StringSliceVar(&opts.SkipImages, "skip-image", []string{}, "Local images names that should not be build when using local build. Takes comma-separated list.")
	flags.StringSliceVar(&opts.FocusImages, "focus-image", []string{}, "Local images to build, all if not specified. Takes comma-separated list.")
	flags.StringVar(&opts.Parameters.Override.HelmRepoURL, "helm-repo-url", capact.HelmRepoStable, fmt.Sprintf("Capact Helm chart repository URL. Use %s tag to select repository which holds the latest Helm chart versions.", capact.LatestVersionTag))
	flags.StringVar(&opts.Parameters.Override.CapactValues.Global.ContainerRegistry.Tag, "override-capact-image-tag", "", "Allows you to override Docker image tag for Capact components. By default, Docker image tag from Helm chart is used.")
	flags.StringVar(&opts.Parameters.Override.CapactValues.Global.ContainerRegistry.Path, "override-capact-image-repo", "", "Allows you to override Docker image repository for Capact components. By default, Docker image repository from Helm chart is used.")
	flags.BoolVar(&opts.Parameters.IncreaseResourceLimits, "increase-resource-limits", true, "Enables higher resource requests and limits for components.")
	flags.BoolVar(&opts.Parameters.Override.CapactValues.Engine.TestSetup.Enabled, "enable-test-setup", false, "Enables test setup for the Capact E2E validation scenarios.")
	flags.BoolVar(&opts.Parameters.Override.CapactValues.Notes.PrintInsecure, "print-insecure-helm-release-notes", false, "Prints the base64-encoded Gateway password directly in Helm release notes.")
	flags.BoolVar(&opts.Parameters.Override.CapactValues.HubPublic.Populator.Enabled, "enable-populator", true, "Enables Public Hub data populator")
	flags.DurationVar(&opts.Timeout, "timeout", 10*time.Minute, `Maximum time during which the upgrade process is being watched, where "0" means "infinite". Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".`)
	flags.BoolVar(&opts.UpdateHostsFile, "update-hosts-file", true, "Updates /etc/hosts with entry for Capact GraphQL Gateway.")
	flags.BoolVar(&opts.UpdateTrustedCerts, "update-trusted-certs", true, "Add Capact GraphQL Gateway certificate.")
	flags.BoolVar(&opts.Verbose, "verbose", false, "Prints more verbose output.")

	return installCmd
}