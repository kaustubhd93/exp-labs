# Key Concepts

- Terraform is packed into a single binary 
- It also has a provider specific binary. Provider can be aws, gcp, azure, kubernetes, etc.
- Providers can be of 3 tiers: official(directly by company), verified(actively maintained and compatible with terraform), community(no guaruntee of maintenance)
- Code for terraform is written in Hashicorp Configuration language.
- Terraform is declarative but terraform language features imperative-like functionality.
- Terraform Lifecycle: Code > terraform init > terraform plan > terraform validate > terraform apply > Code. apply > terraform destory (If infra needs to be retired or is ephemeral)
- Terraform is logically split into 2 parts
    - core: uses remote procedure calls to communicate with Terraform plugins
    - plugin: exposes an implementation for a specific service, or provisioner 

## terraform registry 

### public 
- publicly accessible to everyone.

### private
- terraform cloud can be used to publish your private modules used within the organization. 
- just connect a version control system supported by terraform cloud.

## tf file syntax

```
"<BLOCK_TYPE>" "<BLOCK LABEL>" "<BLOCK LABEL>"{
    #block body
    <IDENTIFIER> = <EXPRESSION> #argument
}
```
- BLOCK_TYPE = can have zero or more labels
- BLOCK_LABEL = name of block
- EXPRESSION = represent a value, either literally or by combining other values or by referencing

## tf file alternate syntax

```
{
  "resource": {
    "aws_instance": {
      "example": {
        "instance_type": "t2.micro",
        "ami": "ami-abc123"
      }
    }
  }
}

```


## local values

- A local value assigns a name to an expression so you can use the name multiple times within a module instead of repeating the expression. 
- input variables are like function arguments, output variables are like function return values, local values are like a function's temporary local variables.
- unlike variables which are scoped globally, locals are scoped to their modules.
- local can also be referenced inside a local.

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
- terraform providers (check what providers you are using)
- terraform fmt  (reformat config in standard style - correct indentation)
- terraform validate (validate all configuration)
- terraform plan  (Check the actual status of the infrastructure and plan what is to be created)
- terraform plan -var-file=filename.tfvars (terraform.tfvars will be autoloaded when included in root project directory)
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

- .tfvars file
```
variable "variable_name" {
    type = data_type
}
# data_type can be string, number, list(), etc
```
- `data_type` can be https://developer.hashicorp.com/terraform/language/expressions/types

#### variables via env variables

- `export TF_VAR_variable_name=value`

### Precedence of input variables

> top will override the bottom
- env variables
- terraform.tfvars
- terraform.tfvars.json
- *.auto.tfvars or *.auto.tfvars.json
- -var and -var-file


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
- usual structure
```
- variables.tf
- outputs.tf
- main.tf
- data.tf
```

## terraform provisioners

- provisioners can install software, edit files and provision machines created with terraform.
- allows you to work with 2 different provisioners: cloud-init and packer
- use provisioners only as a last resort
- terraform earlier used to support third party provisioner tools like chef, puppet, saltstack, etc but it is now depracated thinking that these could change the state later on. 

### local-exec

- allows you to execute local commands after a resource is provisioned. 
- `local` here could be your laptop, build-server, terraform cloud run
- example: 

```
resource "aws_instance" "web" {
  # ...

  provisioner "local-exec" {
    command = "echo The server's IP address is ${self.private_ip}"
  }
}
```

### remote-exec

- allows you to execute commands on a target resource after a resource is provisioned.
- target resource is the resource which you provisioned using terraform.
- ensure whatever commands you are using are not very complex.
- example:
```
resource "aws_instance" "web" {
  # ...

  # Establishes connection to be used by all
  # generic remote provisioners (i.e. file/remote-exec)
  connection {
    type     = "ssh"
    user     = "root"
    password = var.root_password
    host     = self.public_ip
  }

  provisioner "remote-exec" {
    inline = [
      "puppet apply",
      "consul join ${aws_instance.web.private_ip}",
    ]
  }
}
```
- `inline`: list of commands that will be executed on the remote server
- `scripts`: list of file_paths that are present on local and will be copied over to the remote system and then execute in order you have mentioned.

### file

- used to copy files or directories from local to remote resource provisioned.
```
resource "aws_instance" "web" {
  # ...

  # Copies the myapp.conf file to /etc/myapp.conf
  provisioner "file" {
    source      = "conf/myapp.conf"
    destination = "/etc/myapp.conf"
  }

  # Copies the string in content into /tmp/file.log
  provisioner "file" {
    content     = "ami used: ${self.ami}"
    destination = "/tmp/file.log"
  }

  # Copies the configs.d folder to /etc/configs.d
  provisioner "file" {
    source      = "conf/configs.d"
    destination = "/etc"
  }
}

```

### null_resource

- a placeholder for resources with a trigger. used in cases where we want to re-run the provisioner script or block. For example what if an instance gets replaced if it became unhealthy in an autoscaling group. 
```
resource "aws_instance" "cluster" {
  count = 3

  # ...
}

# The primary use-case for the null resource is as a do-nothing container
# for arbitrary actions taken by a provisioner.
#
# In this example, three EC2 instances are created and then a
# null_resource instance is used to gather data about all three
# and execute a single action that affects them all. Due to the triggers
# map, the null_resource will be replaced each time the instance ids
# change, and thus the remote-exec provisioner will be re-run.
resource "null_resource" "cluster" {
  # Changes to any instance of the cluster requires re-provisioning
  triggers = {
    cluster_instance_ids = join(",", aws_instance.cluster.*.id)
  }

  # Bootstrap script can run on any instance of the cluster
  # So we just choose the first in this case
  connection {
    host = element(aws_instance.cluster.*.public_ip, 0)
  }

  provisioner "remote-exec" {
    # Bootstrap script called with private_ip of each node in the cluster
    inline = [
      "bootstrap-cluster.sh ${join(" ",
      aws_instance.cluster.*.private_ip)}",
    ]
  }
}
```

### terraform_data

- similar to null_resources
- use this over null_resources
> Link: https://developer.hashicorp.com/terraform/language/v1.5.x/resources/terraform-data

## terraform data sources

- allows terraform to use information that is defined outside of Terraform, defined by separate configs or by functions
```
# Find the latest available AMI that is tagged with goldenami = true
data "aws_ami" "web" {
  filter {
    name   = "state"
    values = ["available"]
  }

  filter {
    name   = "tag:goldenami"
    values = ["true"]
  }

  most_recent = true
}
```

## terraform resource meta arguments

They can be used with any resource type to change the behaviour of resources.
- `depends_on`: for specifying explicit dependencies  
- `count`: for creating multiple resource instances according to a count  
- `for_each`: to create multiple instances according to a map or a set of strings.  
- `provider`: for selecting a non default provider config.
- `lifecycle`: for lifecycle customizations
- `provisioner`: for taking actions post resource has been provisioned. 

### depends_on

- terraform implicitly can determine the order of provision for resources but there may be some cases where it cannot determine the correct order.
- it allows you explicitly mention the dependency of a resource. 
- When the dependency object is an entire module, depends_on affects the order in which Terraform processes all of the resources and data sources associated with that module.
- try to use it as last resort because it can cause terraform to create more conservative plans that replace more resources than necessary.
```
resource "aws_iam_role" "example" {
  name = "example"

  # assume_role_policy is omitted for brevity in this example. Refer to the
  # documentation for aws_iam_role for a complete example.
  assume_role_policy = "..."
}

resource "aws_iam_instance_profile" "example" {
  # Because this expression refers to the role, Terraform can infer
  # automatically that the role must be created first.
  role = aws_iam_role.example.name
}

resource "aws_iam_role_policy" "example" {
  name   = "example"
  role   = aws_iam_role.example.name
  policy = jsonencode({
    "Statement" = [{
      # This policy allows software running on the EC2 instance to
      # access the S3 API.
      "Action" = "s3:*",
      "Effect" = "Allow",
    }],
  })
}

resource "aws_instance" "example" {
  ami           = "ami-a1b2c3d4"
  instance_type = "t2.micro"

  # Terraform can infer from this that the instance profile must
  # be created before the EC2 instance.
  iam_instance_profile = aws_iam_instance_profile.example

  # However, if software running in this EC2 instance needs access
  # to the S3 API in order to boot properly, there is also a "hidden"
  # dependency on the aws_iam_role_policy that Terraform cannot
  # automatically infer, so it must be declared explicitly:
  depends_on = [
    aws_iam_role_policy.example
  ]
}

```

### count 

- By default, a resource block configures one real infrastructure object. (Similarly, a module block includes a child module's contents into the configuration one time.) However, sometimes you want to manage several similar objects (like a fixed pool of compute instances) without writing a separate block for each one. Terraform has two ways to do this: `count` and `for_each`.
- can be used with resource types and modules
- should be a whole number 
```
resource "aws_instance" "server" {
  count = 4 # create four similar EC2 instances

  ami           = "ami-a1b2c3d4"
  instance_type = "t2.micro"

  tags = {
    Name = "Server ${count.index}"
  }
}

```

### for_each

- similar to count but can iterate over multiple related objects and can iterate over a map for more dynamic values. 
- A given resource or module block cannot use both count and for_each
```
resource "azurerm_resource_group" "rg" {
  for_each = {
    a_group = "eastus"
    another_group = "westus2"
  }
  name     = each.key
  location = each.value
}

```
```
resource "aws_iam_user" "the-accounts" {
  for_each = toset( ["Todd", "James", "Alice", "Dottie"] )
  name     = each.key
}

```
- In blocks where for_each is set, an additional each object is available in expressions, so you can modify the configuration of each instance. This object has two attributes:  
    - each.key — The map key (or set member) corresponding to this instance.  
    - each.value — The map value corresponding to this instance. (If a set was provided, this is the same as each.key.)  

## Resource behaviour

- When you execute terraform apply, it will perform below things to a resource depending on your configuration. 
    - create: resources that are not in the state or there is no state but are present in the configuration.  
    - destroy: resources that are present in the state but are not present in the configuration anymore
    - update in place: resources that are in configuration and the state but their arguments have changed
    - destroy and recreate: resources arguments have changed but cannot update in place because of remote API limitations.

## lifecycle in resources

- lifecycle block allows you to change what happens to a resource when: create, destroy,etc
- they are nested withing resources.
    - create_before_destroy (bool)
    - prevent_destroy (bool)
    - ignore_changes (list of attribute names), eg: ignore_changes = ["tags"], The ignore_changes feature is intended to be used when a resource is created with references to data that may change in the future, but should not affect said resource after its creation. In some rare cases, settings of a remote object are modified by processes outside of Terraform, which Terraform would then attempt to "fix" on the next run. In order to make Terraform share management responsibilities of a single object with a separate process, the ignore_changes meta-argument specifies resource attributes that Terraform should ignore when planning updates to the associated remote object.
    ```
    resource "aws_instance" "example" {
        # ...

        lifecycle {
            ignore_changes = [
            # Ignore changes to tags, e.g. because a management agent
            # updates these based on some ruleset managed elsewhere.
            tags,
            ]
        }
    }

    ```
    - replace_triggered_by (list of resource or attribute references)
> Refer: https://developer.hashicorp.com/terraform/language/meta-arguments/lifecycle

## expressions in terraform

### type and values

#### primitive types

- string
- number
- bool

#### no type

- null: null in terraform doesnt mean it's nothing, it means it is going to use the default setting

#### complex/structural/collection

- list(tuple)
- map(object)

## string templates

- string interpolation: `"Hello, ${var.name}!"`
- string directive: `"Hello, %{ if var.name!= "" }${var.name}%{else}unnamed%{endif}"`
- heredoc: 
```
<<EOT
%{ for ip in aws_instance.example[*].private_ip }
server ${ip}
%{ endfor }
EOT
```
- heredoc with whitespace stripping:
```
<<EOT
%{ for ip in aws_instance.example[*].private_ip ~}
server ${ip}
%{ endfor ~}
EOT
```

## conditional expressions

- syntax: `condition ? true_val : false_val`; if condition is true, true_val if false false_val
- A common use of conditional expressions is to define defaults to replace invalid values
- `var.a != "" ? var.a : "default-value-for-a"`
- data type of the value on which condition is being evaluated and return value must be same.

## for expressions

- dont confuse this with `for_each`. `for` is an expression and `for_each` is a meta_argument for a resource in terraform.
- accepts input: `list`, `set`, `tuple`, `map`, `object`
```
# square braces return a tuple
[for s in var.list: upper(s)] ==>  ["HELLO", "WORLD"]
# curly braces return an object
{ for s in var.list: s => upper(s)} ==> {hello = "HELLO", world = "WORLD"}
```
- You can also use the two-symbol form with lists and tuples, in which case the additional symbol is the index of each element starting from zero, which conventionally has the symbol name i or idx unless it's helpful to choose a more specific name:
```
[for i, v in var.list : "${i} is ${v}"]
``` 
- The index or key symbol is always optional. If you specify only a single symbol after the for keyword then that symbol will always represent the value of each element of the input collection

## splat expressions

- a more concise way of expressing a common operation that otherwise could be performed using `for`. 
- `[for s in var.list : s.id]` is equivalent `var.list[*].id`
- sort of looks like jsonpath when lists are involved.

## dynamic blocks

- a dynamic block will allow you to use dynamically nested blocks which are repeatable. usually inside the resource block, a particular setting is quite literal and fixed. In order to make that dynamic and more modular, you can use a dynamic block. 
- For example :
```

locals {
  ingress_rules = [{
    port = 443
    description = "port 443"
  },
  {
    port = 80
    description = "port 80"
  }]
}

resource "aws_security_group" "lb" {
  name = "sg-lb"
  vpc_id = data.aws_vpc.main.id

  # This below otherwise would have just been 
  # ingress {
  #   all_settings here 
  # }
  # refer : https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/security_group
  dynamic "ingress" {

    for_each = locals.ingress_rules

    content {
      description = ingress.value.description
      from_port = ingress.value.port
      to_port = ingress.value.port
      protocol = "tcp"
      cidr_blocks = ["0.0.0.0/0"]
    }
  }
}
```

