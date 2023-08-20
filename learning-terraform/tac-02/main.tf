locals {
  mode_of_creation = "terraform"
}

# resource "aws_instance" "tf_inst" {
#   ami = var.ami_id
#   instance_type = var.instance_type

#   tags = {
#     Name             = "terraform_poc"
#     mode_of_creation = local.mode_of_creation
#     environment = var.environment
#   }
# }

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


# 01:31:33

