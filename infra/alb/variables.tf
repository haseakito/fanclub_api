variable "vpc_id" {}

variable "public_subnet_ids" {}

variable "ingress_ports" {
  description = "list of ingress ports"
  default     = [80, 443]
}