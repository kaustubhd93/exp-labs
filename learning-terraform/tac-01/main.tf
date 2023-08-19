# terraform {
#   required_providers {
#     aws = {
#       source  = "hashicorp/aws"
#       version = "~> 4.0"
#     }
#   }
# }

# # Configure the AWS Provider
# provider "aws" {
#   # if not mentioned explicity, picks up default profile
#   #profile = var.profile
#   region = "ap-south-1"
# }

terraform {
  required_providers {
    aws = {
      source = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

provider "aws" {
  # Configuration options
  region = "ap-south-1"
  #profile = "default"
}

locals {
  mode_of_creation = "terraform"
}

variable "instance_type" {
  type = string
}

variable "environment" {
  type = string
}

# variable "profile" {
#   type = string
# }

resource "aws_instance" "tf_inst" {
  ami = "ami-0f5ee92e2d63afc18"
  instance_type = "t3a.nano"
  #instance_type = var.instance_type

  tags = {
    Name             = "terraform_poc"
    mode_of_creation = local.mode_of_creation
    environment = var.environment
  }
}

output "public_ip" {
  value = aws_instance.tf_inst.public_ip
}

output "private_ip" {
  value = aws_instance.tf_inst.private_ip
}

# 01:31:33

