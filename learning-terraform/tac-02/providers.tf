# Terraform settings block
terraform {
  backend "s3" {
    bucket = "poc-deploy"
    key    = "tf-poc/tac-03/terraform.tfstate"
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