resource "aws_security_group" "payment_reconciler" {
  name        = "${var.environment}-${var.service}-lambda-into-vpc"
  description = "Outbound rules for payment reconciler lambda"
  vpc_id = var.vpc_id

  egress {
    from_port       = 0
    to_port         = 0
    protocol        = "-1"
    cidr_blocks     = ["0.0.0.0/0"]
  }

  tags = {
    Name = "${var.environment}-${var.service}-sg"
  }

}

output "lambda_into_vpc_id" {
  value = aws_security_group.payment_reconciler.id
}
