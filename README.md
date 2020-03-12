# synapse crawler

# Systemd unit

```
[Unit]
Description=Synapse crawler service
Requires=docker.service
After=docker.service

[Service]
User=synapse-crawler
Restart=always
ExecStart=/usr/bin/docker run --rm --name crawler 242617/synapse-crawler
ExecStop=/usr/bin/docker stop crawler

[Install]
WantedBy=local.target
```
