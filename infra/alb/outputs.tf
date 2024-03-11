output "alb_security_group_id" {
  value = aws_security_group.this.id
}

output "alb_target_group_arn" {
  value = aws_lb_target_group.this.arn
}