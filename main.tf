terraform {
  required_version = ">= 1.0.0"
}

resource "null_resource" "deploy_app" {
  triggers = {
    always_run = timestamp()
  }

  provisioner "remote-exec" {
    connection {
      host        = "148.135.136.11"      # IP VPS kamu
      user        = "root"                # User VPS
      private_key = file("~/.ssh/id_rsa") # Path ke private key lokal
      # optional: kalau SSH pakai port custom
      # port     = 2222
    }

    inline = [
      # Masuk ke folder project
      "cd hris/hris_be/hris_backend/",

      # Pull update dari Git
      "git checkout staging",
      "git pull",

      # Build & run Docker Compose
      "docker compose up -d --build",

      # Bersihkan Docker image/container lama
      "docker system prune -a --volumes -f"
    ]
  }
}
