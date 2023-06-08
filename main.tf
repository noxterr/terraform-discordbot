terraform {
  required_providers {
    discord-interactions = {
      source  = "roleypoly/discord-interactions"
      version = "0.1.0"
    }
  }
}

provider "discord-interactions" {
  # Configuration options
  application_id = var.application_id
  bot_token      = var.token
}

resource "discord-interactions_global_command" "example" {
  name        = "ping"
  description = "An example guild-specific command"

  option {
    type        = 6
    name        = "user"
    description = "Tell this person hello!"
  }

  option {
    type        = 3
    name        = "message"
    description = "What message do I send?"
  }
}