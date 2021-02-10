# -----------------------------------------------------------------------------
# OUTPUTS
# This stack supports the following output values.
# Documentation: https://www.terraform.io/docs/configuration/outputs.html
# -----------------------------------------------------------------------------

output "name" {
  description = "The name of the stack."
  value       = var.name
}

output "meta_computer_system_name" {
  description = "The name of the computer system."
  value       = var.meta_computer_system_name
}

output "meta_contact_email_address" {
  description = "An email address of a contact person."
  value       = var.meta_contact_email_address
}

output "meta_environment_type" {
  description = "The type of the environment."
  value       = var.meta_environment_type
}

output "data_bucket_name" {
  description = "The data S3 bucket name."
  value       = module.s3data.id
}

output "data_bucket_arn" {
  description = "The data S3 bucket ARN."
  value       = module.s3data.arn
}