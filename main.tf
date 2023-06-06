terraform {
  required_providers {
    discord-interactions = {
      source = "roleypoly/discord-interactions"
      version = "0.1.0"
    }
  }
}

provider "discord-interactions" {
  # Configuration options
  credentials = file("variables.tf")
  application_id = credentials.application_id
  token = credentials.token
}