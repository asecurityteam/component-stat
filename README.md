# component-stat - Settings component for generating a metrics client
[![GoDoc](https://godoc.org/github.com/asecurityteam/component-stat?status.svg)](https://godoc.org/github.com/asecurityteam/component-stat)


[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=asecurityteam_component-stat&metric=bugs)](https://sonarcloud.io/dashboard?id=asecurityteam_component-stat)
[![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=asecurityteam_component-stat&metric=code_smells)](https://sonarcloud.io/dashboard?id=asecurityteam_component-stat)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=asecurityteam_component-stat&metric=coverage)](https://sonarcloud.io/dashboard?id=asecurityteam_component-stat)
[![Duplicated Lines (%)](https://sonarcloud.io/api/project_badges/measure?project=asecurityteam_component-stat&metric=duplicated_lines_density)](https://sonarcloud.io/dashboard?id=asecurityteam_component-stat)
[![Lines of Code](https://sonarcloud.io/api/project_badges/measure?project=asecurityteam_component-stat&metric=ncloc)](https://sonarcloud.io/dashboard?id=asecurityteam_component-stat)
[![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=asecurityteam_component-stat&metric=sqale_rating)](https://sonarcloud.io/dashboard?id=asecurityteam_component-stat)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=asecurityteam_component-stat&metric=alert_status)](https://sonarcloud.io/dashboard?id=asecurityteam_component-stat)
[![Reliability Rating](https://sonarcloud.io/api/project_badges/measure?project=asecurityteam_component-stat&metric=reliability_rating)](https://sonarcloud.io/dashboard?id=asecurityteam_component-stat)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=asecurityteam_component-stat&metric=security_rating)](https://sonarcloud.io/dashboard?id=asecurityteam_component-stat)
[![Technical Debt](https://sonarcloud.io/api/project_badges/measure?project=asecurityteam_component-stat&metric=sqale_index)](https://sonarcloud.io/dashboard?id=asecurityteam_component-stat)
[![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=asecurityteam_component-stat&metric=vulnerabilities)](https://sonarcloud.io/dashboard?id=asecurityteam_component-stat)


<!-- TOC -->autoauto- [component-stat - Settings component for generating a metrics client](#component-stat---settings-component-for-generating-a-metrics-client)auto    - [Overview](#overview)auto    - [Quick Start](#quick-start)auto    - [Status](#status)auto    - [Contributing](#contributing)auto        - [Building And Testing](#building-and-testing)auto        - [License](#license)auto        - [Contributing Agreement](#contributing-agreement)autoauto<!-- /TOC -->

## Overview

This is a [`settings`](https://github.com/asecurityteam/settings) that enables
constructing a metrics client. The resulting client is powered by
[`xstats`](https://github.com/rs/xstats) and can output to a variety of metrics
collecting systems.

## Quick Start

```golang
package main

import (
    "context"

    stat "github.com/asecurityteam/component-stat"
    "github.com/asecurityteam/settings/v2"
)

func main() {
    ctx := context.Background()
    envSource := settings.NewEnvSource(os.Environ())

    s := stat.New(ctx, envSource)
    s.Count("my_metric", 1, "tag:value", "tag2:value2")
}
```

## Status

This project is in incubation which means we are not yet operating this tool in
production and the interfaces are subject to change.

## Contributing

### Building And Testing

We publish a docker image called [SDCLI](https://github.com/asecurityteam/sdcli) that
bundles all of our build dependencies. It is used by the included Makefile to help
make building and testing a bit easier. The following actions are available through
the Makefile:

-   make dep

    Install the project dependencies into a vendor directory

-   make lint

    Run our static analysis suite

-   make test

    Run unit tests and generate a coverage artifact

-   make integration

    Run integration tests and generate a coverage artifact

-   make coverage

    Report the combined coverage for unit and integration tests

### License

This project is licensed under Apache 2.0. See LICENSE.txt for details.

### Contributing Agreement

Atlassian requires signing a contributor's agreement before we can accept a patch. If
you are an individual you can fill out the [individual
CLA](https://na2.docusign.net/Member/PowerFormSigning.aspx?PowerFormId=3f94fbdc-2fbe-46ac-b14c-5d152700ae5d).
If you are contributing on behalf of your company then please fill out the [corporate
CLA](https://na2.docusign.net/Member/PowerFormSigning.aspx?PowerFormId=e1c17c66-ca4d-4aab-a953-2c231af4a20b).
