terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.0"
    }
  }
}

# Configure the AWS Provider
provider "aws" {
  # if not mentioned explicity, picks up default profile
  #profile = var.profile
  region = "ap-south-1"
}

variable "instance_type" {
  type = string
}

# variable "profile" {
#   type = string
# }

resource "aws_instance" "tf_inst" {
  ami = "ami-06984ea821ac0a879"
  #instance_type = "t3a.micro"
  instance_type = var.instance_type

  tags = {
    Name             = "terraform_poc"
    mode_of_creation = "terraform"
  }
}

