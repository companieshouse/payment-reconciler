# Role trust relationship
data "aws_iam_policy_document" "trust_relationship" {
  statement {
    sid     = ""
    effect  = "Allow"
    actions = ["sts:AssumeRole"]
    principals {
      type = "AWS"
      identifiers = ["arn:aws:iam::${var.aws_account_id}:root"]
    }
  }
}

# Create roles
resource "aws_iam_role" "iam_role" {
  for_each           = var.iam_roles
  name               = each.value
  assume_role_policy = data.aws_iam_policy_document.trust_relationship.json
}

# Attach policies to roles
resource "aws_iam_role_policy_attachment" "iam_role_policies" {
  for_each   = var.iam_role_policies
  role       = each.value
  policy_arn = lookup(var.iam_policies_arns,each.key)
}

# Attach AWS Managed policies to roles
resource "aws_iam_role_policy_attachment" "iam_role_managed_policies" {
  for_each   = var.iam_role_aws_policies
  role       = each.value
  policy_arn = "arn:aws:iam::aws:policy/${each.key}"
  depends_on = [aws_iam_role.iam_role]
}
