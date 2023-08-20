# Key Concepts

- Terraform is packed into a single binary 
- It also has a provider specific binary. Provider can be aws, gcp, azure, kubernetes, etc.
- Code for terraform is written in Hashicorp Configuration language.
- Terraform is declarative but terraform language features imperative-like functionality.
- Terraform Lifecycle: Code > terraform init > terraform plan > terraform validate > terraform apply > Code. apply > terraform destory (If infra needs to be retired or is ephemeral)
- Terraform is logically split into 2 parts
    - core: uses remote procedure calls to communicate with Terraform plugins
    - plugin: exposes an implementation for a specific service, or provisioner 

## local values

- A local value assigns a name to an expression so you can use the name multiple times within a module instead of repeating the expression. 
- input variables are like function arguments, output variables are like function return values, local values are like a function's temporary local variables.
- unlike variables which are scoped globally, locals are scoped to their modules.

### Syntax

```
locals {
    var_name = expression
}
```

# IaC key concepts 

##  What is infrastructure lifecycle 

- a number of clearly defined and distinct work phases which are used by DevOps Engineers to plan, design, build, test, deliver, maintain and retire cloud infrastructure.
    - Day 0: Plan and Design (IaC starts here.)
    - Day 1: Develop and iterate
    - Day 2: Go live and maintain 

## Advantages of Infrastructure lifecycle

- Reliability 
    - IAC makes changes idempotent(no matter how many times we run IAC, we will always end up with the same state), consistent, repeatable and predictable
- Manageability
- Sensibility
    - can recreate infrastructure if somebody accidentally deletes anything
    - avoid financial and reputational losses.


## What is configuration drift 
- It is when provisioned infrastructure undergoes an unexpected configuration change due to:
    - some team member manually changed something in infrastructure.
    - malicious actors
    - unchecked usage of cloud provider cli or SDK from some other application.

### Correcting configuration drift
- `terraform apply -refresh-only`

### Avoiding configuration drift
- Create immutable infrastructure: containers, blue-green deployment etc.

## Change management

- procedure that will be followed when resources are modified and applied via configuration script.

## Change automation

- a way of automatically creating a consistent, systematic and predictable way of managing change request via controls and policies 
- Terraform uses change automation in the form of execution plans and resource graphs to apply and review complex changesets (a collection of commits that represent changes made to a version control system ) 

## execution plan

- a manual review of what will be added, changed and destroyed before you apply changes.

# Key commands

- terraform init  (Download plugins if required)
- terraform fmt  (reformat config in standard style - correct indentation)
- terraform validate (validate all configuration)
- terraform plan  (Check the actual status of the infrastructure and plan what is to be created)
- terraform plan -var-file=filename.tfvars 
- terraform apply (Create planned resources.)
- terraform apply -auto-approve (Can be used in CICD pipelines)
- terraform apply -refresh-only (for correcting configuration drifts)
- terraform validate (Validates your tf file configuraion, requires terraform init)

# Creating resorces 

## Syntax 

### resource

```
resource "provider_servicename" "resource_name" {
    config ...
    key = value
    key2 = value2
}
```

### input variable syntax

```
variable "variable_name" {
    type = data_type
}
# data_type can be string, number, list(), etc
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

## terraform modules 

- a set of tf files in a directory that allows you to create logical abstraction on top of some resource set. use case: reusable configuration.
- a single tf file in a directory is also a module, it is called as root module.
- 