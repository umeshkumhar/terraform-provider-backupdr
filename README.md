# Terraform Provider Scaffolding (Terraform Plugin Framework)

_This template repository is built on the [Terraform Plugin Framework](https://github.com/hashicorp/terraform-plugin-framework). The template repository built on the [Terraform Plugin SDK](https://github.com/hashicorp/terraform-plugin-sdk) can be found at [terraform-provider-scaffolding](https://github.com/hashicorp/terraform-provider-scaffolding). See [Which SDK Should I Use?](https://developer.hashicorp.com/terraform/plugin/framework-benefits) in the Terraform documentation for additional information._

This repository is a *template* for a [Terraform](https://www.terraform.io) provider. It is intended as a starting point for creating Terraform providers, containing:

- A resource and a data source (`internal/provider/`),
- Examples (`examples/`) and generated documentation (`docs/`),
- Miscellaneous meta files.

These files contain boilerplate code that you will need to edit to create your own Terraform provider. Tutorials for creating Terraform providers can be found on the [HashiCorp Developer](https://developer.hashicorp.com/terraform/tutorials/providers-plugin-framework) platform. _Terraform Plugin Framework specific guides are titled accordingly._

Please see the [GitHub template repository documentation](https://help.github.com/en/github/creating-cloning-and-archiving-repositories/creating-a-repository-from-a-template) for how to create a new repository from this template on GitHub.

Once you've written your provider, you'll want to [publish it on the Terraform Registry](https://developer.hashicorp.com/terraform/registry/providers/publishing) so that others can use it.

## Requirements

- [Terraform](https://developer.hashicorp.com/terraform/downloads) >= 1.0
- [Go](https://golang.org/doc/install) >= 1.20

## Building The Provider

1. Clone the repository
1. Enter the repository directory
1. Build the provider using the Go `install` command:

```shell
go install
```

## Adding Dependencies

This provider uses [Go modules](https://github.com/golang/go/wiki/Modules).
Please see the Go documentation for the most up to date information about using Go modules.

To add a new dependency `github.com/author/dependency` to your Terraform provider:

```shell
go get github.com/author/dependency
go mod tidy
```

Then commit the changes to `go.mod` and `go.sum`.

## Using the provider

Fill this in for each provider

## Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (see [Requirements](#requirements) above).

To compile the provider, run `go install`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

To generate or update documentation, run `go generate`.

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```shell
make testacc
```


{
  "azp": "764086051850-6qr4p6gpi6hn506pt8ejuq83di341hur.apps.googleusercontent.com",
  "aud": "764086051850-6qr4p6gpi6hn506pt8ejuq83di341hur.apps.googleusercontent.com",
  "sub": "105624511008430578532",
  "scope": "https://www.googleapis.com/auth/cloud-platform https://www.googleapis.com/auth/sqlservice.login https://www.googleapis.com/auth/userinfo.email openid",
  "exp": "1712041876",
  "expires_in": "3319",
  "email": "umeshkumhar@google.com",
  "email_verified": "true"
}

{
  "azp": "32555940559.apps.googleusercontent.com",
  "aud": "32555940559.apps.googleusercontent.com",
  "sub": "105624511008430578532",
  "scope": "https://www.googleapis.com/auth/accounts.reauth https://www.googleapis.com/auth/appengine.admin https://www.googleapis.com/auth/cloud-platform https://www.googleapis.com/auth/compute https://www.googleapis.com/auth/sqlservice.login https://www.googleapis.com/auth/userinfo.email openid",
  "exp": "1712039205",
  "expires_in": "611",
  "email": "umeshkumhar@google.com",
  "email_verified": "true"
}



ya29.a0Ad52N3_ffqwurrIYBQB218W4YberJjNPCQgSStxBlnXQVH0B8pFtUDia1LE9Ok9xqohfXzUJ_BGvrilrDHmOhIM5FQ_--7CiIgCyqVqHTco8M6qK8-q85d8nuOagD8OUhMz9pDjl1uUfIqVZv8Uoqeu_W0wq6jjdZthqnwaCgYKAaESARISFQHGX2Mi54cH_trNQf70S7SeAn6tjQ0173

ya29.a0Ad52N38THmsa8HwxMwGYcHGr3KS_srcewsC2tE0dLG2oqDDR9HRzWL87fMLIrfzbazeusTuv01_6K3-UZ07EWJsO-tlixMZSxxAmW03wJYeEJHDkw2wUp-t28L80yNDEsjov4cwTu6VM6WMzeMPL_luir2c8GKBdF0BYhQaCgYKAWISARISFQHGX2Mi2jT0J6OnOv3cJeI6ctJCew0173# terraform-provider-backupdr
# terraform-provider-backupdr
