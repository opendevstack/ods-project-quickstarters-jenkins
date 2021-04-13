# -----------------------------------------------------------------------------
# ENVIRONMENT VARIABLES
# This stack supports the following secrets as environment variables.
# -----------------------------------------------------------------------------

# AWS_ACCESS_KEY_ID
# AWS_SECRET_ACCESS_KEY
# AWS_DEFAULT_REGION

# -----------------------------------------------------------------------------
# REQUIRED PARAMETERS
# The following stack parameters require a value.
# Documentation: https://www.terraform.io/docs/configuration/variables.html
# -----------------------------------------------------------------------------

# -----------------------------------------------------------------------------
# OPTIONAL PARAMETERS
# The following stack parameters are optional with sensible defaults.
# Documentation: https://www.terraform.io/docs/configuration/variables.html
# -----------------------------------------------------------------------------

variable "name" {
  description = "The name of the stack."
  type        = string
  default     = "stack-aws-quickstarter"
}

variable "meta_business_application_name" {
  description = "The name of the business application system."
  type        = string
  default     = "quickstarter"
}

variable "meta_environment" {
  description = "The type of the environment. Can be any of DEVELOPMENT, EVALUATION, PRODUCTIVE, QUALITYASSURANCE, TRAINING, VALIDATION."
  type        = string
  default     = "DEVELOPMENT"
}

variable "meta_contact_email_address" {
  description = "An email address of a contact person."
  type        = string
  default     = "changeme@phoenix.com"
}

variable "meta_dynamic_resource_name" {
  description = "The name of the dynamic resource."
  type        = string
  default     = "undefined"
}

# ---------------------------------------------
# S3 Bucket Variables
# ---------------------------------------------

variable "data_bucket_name" {
  description = "The name of the S3 data bucket."
  type        = string
  default     = "quickstarter"
}

