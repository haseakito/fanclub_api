variable "vpc_id" {}

variable "alb_security_group_id" {}

variable "alb_target_group_arn" {}

variable "public_subnet_ids" {}

variable "db_url" {}

variable "cloudinary_cloud_name" {}

variable "cloudinary_api_key" {}

variable "cloudinary_api_secret_key" {}

variable "sentry_dsn" {}

variable "clerk_webhook_secret" {}

# variable "ingress_ports" {
#   description = "list of ingress ports"
#   default     = [80, 443]
# }
