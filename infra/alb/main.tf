resource "aws_security_group" "this" {
  vpc_id = var.vpc_id

  dynamic "ingress" {
    for_each = var.ingress_ports

    iterator = port

    content {
      from_port   = port.value
      to_port     = port.value
      protocol    = "tcp"
      cidr_blocks = ["0.0.0.0/0"]
    }
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    "Name" = "fanclub-alb-sg"
  }
}

resource "aws_lb" "this" {
  name               = "fanclub-lb"
  load_balancer_type = "application"

  security_groups = [aws_security_group.this.id]
  subnets         = var.public_subnet_ids

  tags = {
    "Environment" = "production"
  }
}

resource "aws_lb_target_group" "this" {
  name = "fanclub-lb-tg"
  vpc_id = var.vpc_id
  port = 80
  protocol = "HTTP"
  target_type = "ip"

  health_check {
    port = 80
    path = "/"
    timeout = 20
  }
}

resource "aws_lb_listener" "http" {
  port              = 80
  protocol          = "HTTP"
  load_balancer_arn = aws_lb.this.arn

  default_action {
    type = "fixed-response"
    fixed_response {
      content_type = "text/plain"
      message_body = "503 Service Unavailable"
      status_code  = "503"
    }
  }
}

resource "aws_lb_listener_rule" "http_rule" {
  listener_arn = aws_lb_listener.http.arn

  action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.this.arn
  }

  condition {
    path_pattern {
      values = ["*"]
    }
  }
}