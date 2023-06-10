---
title: Transip Installation & Configuration
meta_desc: Information on how to install the Transip provider.
layout: installation
---

## Installation

The Pulumi Transip provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@maienm/pulumi-transip`](https://www.npmjs.com/package/@maienm/pulumi-transip)
* Python: [`pulumi_transip`](https://pypi.org/project/pulumi_transip/)
* Go: [`github.com/MaienM/pulumi-transip/sdk/go/transip`](https://github.com/MaienM/pulumi-transip/sdk/go/transip)
* .NET: [`MaienM.Transip`](https://www.nuget.org/packages/MaienM.Transip)

### Provider Binary

The Transip provider binary is a third party binary. It can be installed using the `pulumi plugin` command.

```bash
pulumi plugin install resource transip
```

Replace the version string with your desired version.
