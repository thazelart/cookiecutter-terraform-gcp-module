# cookiecutter-terraform-gcp-module

Cookiecutter template to create Terraform modules.


## Usage

```bash
$ pip install cookiecutter
$ cookiecutter https://github.com/thazelart/cookiecutter-terraform-gcp-module.git
```

You will be prompted for basic info (your name, module name, etc.) which will be used in the template.

## Output
A basic terraform module with example and test

```
└── tf-mod-name
    ├── examples
    │   └── default_example
    │       ├── backend.tf
    │       ├── main.tf
    │       ├── provider.tf
    │       └── variables.tf
    ├── test
    │   └── simple_example_test.go
    ├── CHANGELOG.md
    ├── LICENSE
    ├── main.tf
    ├── outputs.tf
    ├── simple_example.tf
    └── variables.tf

```

To generate you documentation, update the `main.tf`, `variables.tf` and `outputs.tf`. Once done, please use [terraform-docs](https://github.com/segmentio/terraform-docs) 
```
terraform-docs markdown tables . > README.mdcd
```

## Authors
[Thibault Hazelart](https://github.com/thazelart)

## License
[Aache 2.0](/LICENSE)
