resource "aws_security_group" "paynment_reconciler_sg" {
  name        = "${var.env}-${var.service}-lambda-into-vpc"
  description = "Outbound rules for payment reconciler lambda"
  vpc_id = "${var.vpc_id}"

  egress {
    from_port       = 0
    to_port         = 0
    protocol        = "-1"
    cidr_blocks     = ["0.0.0.0/0"]
  }
}

output "lambda_into_vpc_id" {
  value = "${aws_security_group.paynment_reconciler_sg.id}"
}
