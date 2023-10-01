locals {
  mode_of_creation = "terraform"
}

module "vpc" {
  source = "terraform-aws-modules/vpc/aws"

  name = var.vpc_name
  cidr = var.vpc_cidr

  azs             = ["ap-south-1a", "ap-south-1b"]
  private_subnets = var.private_subnets
  public_subnets  = var.public_subnets

  enable_nat_gateway = false
  enable_vpn_gateway = false

  tags = {
    Name = var.vpc_name
    environment = var.environment
    mode_of_creation = local.mode_of_creation

  }
}


