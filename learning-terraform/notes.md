# Key Concepts

- Terraform is packed into a single binary 
- It also has a provider specific binary. Provider can be aws, gcp, azure, kubernetes, etc.
- Code for terraform is written in Hashicorp Configuration language.

# Key commands

- terraform init  (Download plugins if required)
- terraform plan  (Check the actual status of the infrastructure and plan what is to be created)
- terraform apply (Create planned resources.)

# Creating resorces 

## Syntax 

```
resource "provider_servicename" "resource_name" {
    config ...
    key = value
    key2 = value2
}
```