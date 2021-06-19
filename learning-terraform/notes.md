# Key Concepts

- Terraform is packed into a single binary 
- It also has a provider specific binary. Provider can be aws, gcp, azure, kubernetes, etc.
- Code for terraform is written in Hashicorp Configuration language.

# Key commands

- terraform init  (Download plugins if required)
- terraform plan  (Check the actual status of the infrastructure and plan what is to be created)
- terraform apply (Create planned resources.)
- terraform validate (Validates your tf file configuraion, requires terraform init)

# Creating resorces 

## Syntax 

```
resource "provider_servicename" "resource_name" {
    config ...
    key = value
    key2 = value2
}
```

## Key obervations

- Terraform does not care about the order in which you have defined resource in the tf file. 
- We cannot run only one file with terraform like we do with python or general purpose programming files. Below statement will not run, it will just spit out the help page.
    ```
    terraform plan filename.tf
    ```
- The provider block can exist only in one .tf file. 

## Gotchas

- Do not try to tinker with the terraform.tfstate file.
- Do not try to copy an existing .terraform folder for a different tf project. It does not recognise it. Run `terraform init` for separate projects. 