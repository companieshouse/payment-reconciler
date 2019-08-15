resource "aws_security_group" "paynment_reconciler_sg" {
  name        = "payment-reconciler-security-group"
  description = "Outbound rules for payment reconciler lambda"

  egress {
    from_port       = 0
    to_port         = 0
    protocol        = "-1"
    cidr_blocks     = ["0.0.0.0/0"]
  }
}

output "security_group_ids" {
  value = "${aws_security_group.paynment_reconciler_sg.id}"
}
