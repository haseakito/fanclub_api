locals {
  account_id = data.aws_caller_identity.user.account_id
}

resource "aws_ecs_cluster" "this" {
  name = "fanclub-cluster"
}

resource "aws_iam_role" "ecs_task_execution" {
  name = "ecs_task_execution"
  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Action = "sts:AssumeRole"
        Principal = {
          Service = "ecs-tasks.amazonaws.com"
        }
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "ecs_task_execution" {
  role       = aws_iam_role.ecs_task_execution.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy"
}

resource "aws_iam_role_policy_attachment" "ecs_task_execution_ssm" {
  role       = aws_iam_role.ecs_task_execution.name
  policy_arn = "arn:aws:iam::aws:policy/AmazonSSMReadOnlyAccess"
}

resource "aws_ecs_task_definition" "this" {
  family                   = "fanclub-task"
  cpu                      = 256
  memory                   = 2048
  network_mode             = "awsvpc"
  requires_compatibilities = ["FARGATE"]

  execution_role_arn = aws_iam_role.ecs_task_execution.arn
  container_definitions = templatefile("./ecs/container_definitions.json",
    {
      account_id                = local.account_id
      db_url                    = var.db_url
      cloudinary_cloud_name     = var.cloudinary_cloud_name
      cloudinary_api_key        = var.cloudinary_api_key
      cloudinary_api_secret_key = var.cloudinary_api_secret_key
      sentry_dsn                = var.sentry_dsn
      clerk_webhook_secret      = var.clerk_webhook_secret
    }
  )
}

resource "aws_ecs_service" "this" {
  name            = "fanclub-service"
  launch_type     = "FARGATE"
  desired_count   = 1
  cluster         = aws_ecs_cluster.this.name
  task_definition = aws_ecs_task_definition.this.arn

  network_configuration {
    security_groups  = [var.alb_security_group_id]
    subnets          = var.public_subnet_ids
    assign_public_ip = true
  }

  load_balancer {
    target_group_arn = var.alb_target_group_arn
    container_name   = "nginx"
    container_port   = 80
  }
}
