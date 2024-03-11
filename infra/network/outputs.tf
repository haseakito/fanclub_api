output "vpc_id" {
  value = aws_vpc.this.id
}

output "public_subnet_ids" {
  value = [aws_subnet.public-subnet-1a.id, aws_subnet.public-subnet-1c.id]
}

output "private_subnet_ids" {
  value = [aws_subnet.private-subnet-1a.id, aws_subnet.private-subnet-1c.id]
}

output "db_subnet_group_name" {
  value = aws_db_subnet_group.this.name
}
