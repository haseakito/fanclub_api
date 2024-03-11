# Configure AWS VPC
resource "aws_vpc" "this" {
  cidr_block           = "10.0.0.0/16"
  instance_tenancy     = "default"
  enable_dns_support   = true
  enable_dns_hostnames = true

  tags = {
    "Name" = "fanclub-vpc"
  }
}

# Public subnet 1a
resource "aws_subnet" "public-subnet-1a" {
  vpc_id            = aws_vpc.this.id
  cidr_block        = "10.0.0.0/24"
  availability_zone = "ap-northeast-1a"

  tags = {
    "Name" = "fabclub-public-subnet-1a"
  }
}

# Public subnet 1c
resource "aws_subnet" "public-subnet-1c" {
  vpc_id            = aws_vpc.this.id
  cidr_block        = "10.0.1.0/24"
  availability_zone = "ap-northeast-1c"

  tags = {
    "Name" = "fabclub-public-subnet-1c"
  }
}

# Private subnet 1a
resource "aws_subnet" "private-subnet-1a" {
  vpc_id            = aws_vpc.this.id
  cidr_block        = "10.0.2.0/24"
  availability_zone = "ap-northeast-1a"

  tags = {
    "Name" = "fabclub-private-subnet-1a"
  }
}

# Private subnet 1c
resource "aws_subnet" "private-subnet-1c" {
  vpc_id            = aws_vpc.this.id
  cidr_block        = "10.0.3.0/24"
  availability_zone = "ap-northeast-1c"

  tags = {
    "Name" = "fabclub-private-subnet-1c"
  }
}

# Configure AWS Internet Gateway
resource "aws_internet_gateway" "this" {
  vpc_id = aws_vpc.this.id

  tags = {
    "Name" = "fanclub-igw"
  }
}

# Private Route Table
resource "aws_route_table" "public" {
  vpc_id = aws_vpc.this.id

  tags = {
    "Name" = "fanclub-rtb-public"
  }
}

# Configure AWS Route
resource "aws_route" "public" {
  route_table_id         = aws_route_table.public.id
  gateway_id             = aws_internet_gateway.this.id
  destination_cidr_block = "0.0.0.0/0"
}


resource "aws_route_table_association" "public-1a" {
  subnet_id      = aws_subnet.public-subnet-1a.id
  route_table_id = aws_route_table.public.id
}

resource "aws_route_table_association" "public-1c" {
  subnet_id      = aws_subnet.public-subnet-1c.id
  route_table_id =  aws_route_table.public.id
}

# Private Route Table
resource "aws_route_table" "private" {
  vpc_id = aws_vpc.this.id
  tags = {
    Name = "fanclub-rtb-private"
  }
}

resource "aws_route_table_association" "private-1a" {
  subnet_id      = aws_subnet.private-subnet-1a.id
  route_table_id = aws_route_table.private.id
}

resource "aws_route_table_association" "private-1c" {
  subnet_id      = aws_subnet.private-subnet-1c.id
  route_table_id = aws_route_table.private.id
}

# Subnet group for database
resource "aws_db_subnet_group" "this" {
  name        = var.db_name
  description = "db subnet group of ${var.db_name}"
  subnet_ids  = [aws_subnet.private-subnet-1a.id, aws_subnet.private-subnet-1c.id]
}
