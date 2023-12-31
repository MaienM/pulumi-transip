// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package transip

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/ettle/strcase"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	shim "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim"
	shimv1 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v1"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"

	"github.com/MaienM/pulumi-transip/provider/pkg/version"
	"github.com/aequitas/terraform-provider-transip/provider"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

// all of the token components used below.
const (
	// This variable controls the default name of the package in the package
	mainMod = "index" // the transip module
)

func convertName(name string) string {
	idx := strings.Index(name, "_")
	contract.Assertf(idx > 0 && idx < len(name)-1, "Invalid snake case name %s", name)
	name = name[idx+1:]
	contract.Assertf(len(name) > 0, "Invalid snake case name %s", name)
	return strcase.ToPascal(name)
}

func makeDataSource(mod string, name string) tokens.ModuleMember {
	name = convertName(name)
	return tfbridge.MakeDataSource("transip", mod, "get"+name)
}

func makeResource(mod string, res string) tokens.Type {
	return tfbridge.MakeResource("transip", mod, convertName(res))
}

// preConfigureCallback is called before the providerConfigure function of the underlying provider.
// It should validate that the provider can be configured, and provide actionable errors in the case
// it cannot be. Configuration variables can be read from `vars` using the `stringValue` function -
// for example `stringValue(vars, "accessKey")`.
func preConfigureCallback(vars resource.PropertyMap, c shim.ResourceConfig) error {
	return nil
}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := shimv1.NewProvider(provider.Provider())
	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:    p,
		Name: "transip",
		// DisplayName is a way to be able to change the casing of the provider
		// name when being displayed on the Pulumi registry
		DisplayName: "TransIP",
		// The default publisher for all packages is Pulumi.
		// Change this to your personal name (or a company name) that you
		// would like to be shown in the Pulumi Registry if this package is published
		// there.
		Publisher: "MaienM",
		// LogoURL is optional but useful to help identify your package in the Pulumi Registry
		// if this package is published there.
		//
		// You may host a logo on a domain you control or add an SVG logo for your package
		// in your repository and use the raw content URL for that file as your logo URL.
		LogoURL: "https://raw.githubusercontent.com/MaienM/pulumi-transip/main/docs/transip.png",
		// PluginDownloadURL is an optional URL used to download the Provider
		// for use in Pulumi programs
		// e.g https://github.com/org/pulumi-provider-name/releases/
		PluginDownloadURL: "github://api.github.com/MaienM/pulumi-transip",
		Description:       "A Pulumi package for creating and managing TransIP resources",
		// category/cloud tag helps with categorizing the package in the Pulumi Registry.
		// For all available categories, see `Keywords` in
		// https://www.pulumi.com/docs/guides/pulumi-packages/schema/#package.
		Keywords: []string{
			"pulumi",
			"transip",
			"category/cloud",
		},
		License:    "Apache-2.0",
		Homepage:   "https://github.com/MaienM/pulumi-transip",
		Repository: "https://github.com/MaienM/pulumi-transip",
		// The GitHub Org for the provider - defaults to `terraform-providers`. Note that this
		// should match the TF provider module's require directive, not any replace directives.
		Version:              version.Version,
		GitHubOrg:            "aequitas",
		Config:               map[string]*tfbridge.SchemaInfo{},
		PreConfigureCallback: preConfigureCallback,
		Resources: map[string]*tfbridge.ResourceInfo{
			"transip_dns_record":                 {Tok: makeResource(mainMod, "transip_dns_record")},
			"transip_domain":                     {Tok: makeResource(mainMod, "transip_domain")},
			"transip_domain_dnssec":              {Tok: makeResource(mainMod, "transip_domain_dnssec")},
			"transip_domain_nameservers":         {Tok: makeResource(mainMod, "transip_domain_nameservers")},
			"transip_openstack_project":          {Tok: makeResource(mainMod, "transip_openstack_project")},
			"transip_openstack_user":             {Tok: makeResource(mainMod, "transip_openstack_user")},
			"transip_private_network":            {Tok: makeResource(mainMod, "transip_private_network")},
			"transip_private_network_attachment": {Tok: makeResource(mainMod, "transip_private_network_attachment")},
			"transip_sshkey":                     {Tok: makeResource(mainMod, "transip_sshkey")},
			"transip_vps":                        {Tok: makeResource(mainMod, "transip_vps")},
			"transip_vps_firewall":               {Tok: makeResource(mainMod, "transip_vps_firewall")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"transip_domain":            {Tok: makeDataSource(mainMod, "transip_domain")},
			"transip_domains":           {Tok: makeDataSource(mainMod, "transip_domains")},
			"transip_openstack_project": {Tok: makeDataSource(mainMod, "transip_openstack_project")},
			"transip_openstack_user":    {Tok: makeDataSource(mainMod, "transip_openstack_user")},
			"transip_private_network":   {Tok: makeDataSource(mainMod, "transip_private_network")},
			"transip_sshkey":            {Tok: makeDataSource(mainMod, "transip_sshkey")},
			"transip_vps":               {Tok: makeDataSource(mainMod, "transip_vps")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			PackageName: "@maienm/pulumi-transip",

			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
			// See the documentation for tfbridge.OverlayInfo for how to lay out this
			// section, or refer to the AWS provider. Delete this section if there are
			// no overlay files.
			//Overlay: &tfbridge.OverlayInfo{},
		},
		Python: &tfbridge.PythonInfo{
			PackageName: "pulumi_transip",

			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/MaienM/pulumi-%[1]s/sdk/", "transip"),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				"transip",
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			RootNamespace: "MaienM",

			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
		Java: &tfbridge.JavaInfo{
			BasePackage: "com.maienm",
		},
	}

	prov.SetAutonaming(255, "-")

	return prov
}
