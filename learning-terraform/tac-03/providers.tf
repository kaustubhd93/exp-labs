# Terraform settings block
terraform {
  # cloud {
  #   organization = "kd-tf-org"
  #   workspaces {
  #     name = "poc-tac-02"
  #   }
  # }
#   backend "remote" {
#     hostname = "app.terraform.io"
#     organization = "kd-tf-org"

#     workspaces {
#       name = "poc-tac-02"
#     }
#   }
  backend "s3" {
    bucket = "poc-deploy"
    key    = "tf-poc/tac-02/terraform.tfstate"
    #key    = "terraform.tfstate"
    #workspace_key_prefix = "poc-tac02"
    region = "ap-south-1"
  }
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
  profile = "default"
}