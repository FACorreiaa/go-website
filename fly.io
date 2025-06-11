# fly.toml app configuration file generated for skyvisor on 2024-04-25T00:02:00+02:00
#
# See https://fly.io/docs/reference/configuration/ for information about how to use this file.
#

app = 'go-facorreia-website'
primary_region = 'lhr'

[build]
  dockerfile = "Dockerfile"
  build-target = "production"

[env]
  MODE="production"
  SERVER_PORT = 6968
  SERVER_WRITE_TIMEOUT=15
  SERVER_READ_TIMEOUT=15
  SERVER_IDLE_TIMEOUT=150
  SERVER_GRACEFUL_TIMEOUT=15

[build.args]
  MODE="production"

[http_service]
  internal_port = 6968
  force_https = true
  auto_stop_machines = true
  auto_start_machines = true
  min_machines_running = 0
  processes = ['app']

[[services.ports]]
  handlers = ["tls", "http"]
  port = 443

[[services.tls.domains]]
  domain = "singuwe.com"


