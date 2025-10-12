variable "resource_group_name" {
  description = "The name of the resource group"
  type        = string
  default     = "terraform-iac-rg"
}

variable "location" {
  description = "The Azure region to deploy resources"
  type        = string
  default     = "Canada Central"
}

variable "acr_name" {
  description = "The name of the Azure Container Registry"
  type        = string
  default     = "myacr12345"
}

variable "app_service_name" {
  description = "The name of the Azure App Service"
  type        = string
  default     = "my-app-service-plan"
}

variable "app_service_plan_name" {
  description = "The name of the Azure App Service"
  type        = string
  default     = "my-app-service"
}

variable "subscription_id" {
  description = "The Azure Subscription ID"
  type        = string
  default     = "1cc729ee-fbd5-4b09-95e8-64fa2a6f2b8b"
}