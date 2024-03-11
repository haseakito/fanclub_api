terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
  backend "s3" {
    bucket = "own-terraform-tfstate"
    key    = "terraform.tfstate"
    region = "ap-northeast-1"
  }
}

# Configure the AWS Provider
provider "aws" {}

# Set up VPC and Subnet
module "network" {
  source = "./network"

  db_name = var.db_name
}

# Set up Load Balancer
module "alb" {
  source = "./alb"

  vpc_id            = module.network.vpc_id
  public_subnet_ids = module.network.public_subnet_ids
}

# Set up RDS
module "rds" {
  source = "./rds"

  vpc_id                = module.network.vpc_id
  db_name               = var.db_name
  db_user               = var.db_user
  db_password           = var.db_password
  db_subnet_group_name  = module.network.db_subnet_group_name
  alb_security_group_id = module.alb.alb_security_group_id
}

module "ecs" {
  source = "./ecs"

  vpc_id                    = module.network.vpc_id
  db_url                    = var.DATABE_URL
  cloudinary_cloud_name     = var.CLOUDINARY_CLOUD_NAME
  cloudinary_api_key        = var.CLOUDINARY_API_KEY
  cloudinary_api_secret_key = var.CLOUDINARY_API_SECRET_KEY
  sentry_dsn                = var.SENTRY_DSN
  clerk_webhook_secret      = var.CLERK_WEBHOOK_SECRET
  alb_security_group_id     = module.alb.alb_security_group_id
  public_subnet_ids         = module.network.public_subnet_ids
  alb_target_group_arn      = module.alb.alb_target_group_arn
}
