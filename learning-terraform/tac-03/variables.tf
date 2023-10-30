# variable "profile" {
#   type = string
# }

variable "environment" {
  type = string
}

variable "vpc_config" {
  type = map(object({
    vpc_name = string
    vpc_cidr = string
    private_subnets = list(string)
    public_subnets = list(string)
  }))
}