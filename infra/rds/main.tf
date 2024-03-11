resource "aws_security_group" "this" {
  name   = "rds-sg"
  vpc_id = var.vpc_id

  tags = {
    "Name" = "fanclub-rds-sg"
  }
}

resource "aws_security_group_rule" "this" {
  security_group_id = aws_security_group.this.id

  type                     = "ingress"
  from_port                = 3306
  to_port                  = 3306
  protocol                 = "tcp"
  source_security_group_id = var.alb_security_group_id
}

# Configure Database
resource "aws_db_instance" "this" {
  allocated_storage       = 10
  max_allocated_storage   = 1000
  storage_type            = "gp2"
  engine                  = "MYSQL"
  engine_version          = "8.0"
  instance_class          = "db.t2.micro"
  identifier              = "fanclubdb"
  username                = var.db_user
  password                = var.db_password
  port                    = 3306
  backup_retention_period = 7
  copy_tags_to_snapshot   = true
  skip_final_snapshot     = true
  enabled_cloudwatch_logs_exports = [
    "error",
    "general",
    "slowquery"
  ]
  parameter_group_name   = aws_db_parameter_group.db-pg.name
  vpc_security_group_ids = [aws_security_group.this.id]
  db_subnet_group_name   = var.db_subnet_group_name
}

# Database parameter
resource "aws_db_parameter_group" "db-pg" {
  name   = "db-pg"
  family = "mysql8.0"

  parameter {
    name  = "character_set_client"
    value = "utf8mb4"
  }

  parameter {
    name  = "character_set_connection"
    value = "utf8mb4"
  }

  parameter {
    name  = "character_set_results"
    value = "utf8mb4"
  }

  parameter {
    name  = "character_set_server"
    value = "utf8mb4"
  }
}
