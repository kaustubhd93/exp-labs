locals {
  mode_of_creation = "terraform"
}

module "vpc" {
  source = "terraform-aws-modules/vpc/aws"

  for_each = var.vpc_config

  name = var.vpc_config[each.key].vpc_name
  cidr = var.vpc_config[each.key].vpc_cidr

  azs             = ["ap-south-1a", "ap-south-1b"]
  private_subnets = var.vpc_config[each.key].private_subnets
  public_subnets  = var.vpc_config[each.key].public_subnets

  enable_nat_gateway = false
  enable_vpn_gateway = false

  tags = {
    Name = var.vpc_config[each.key].vpc_name
    environment = var.environment
    mode_of_creation = local.mode_of_creation

  }

  private_subnet_tags = merge({ Name = "private-${var.vpc_config[each.key].vpc_name}" }, { environment = var.environment, mode_of_creation = local.mode_of_creation})
  public_subnet_tags = merge({ Name = "public-${var.vpc_config[each.key].vpc_name}" }, {environment = var.environment, mode_of_creation = local.mode_of_creation })

}


