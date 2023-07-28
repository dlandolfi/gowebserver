# gowebserver

A repo for practicing backend deployments to an Ubuntu server.

## Steps To Reproduce

### Configure firewall and nginx

Make sure nginx and ufw are installed: `apt update && apt install nginx`

Make sure the nginx systemd service is enabled: `systemctl enable nginx`

Enable firewall: `ufw enable`

Enable ssh: `ufw allow ssh`

Configure ufw to block all incoming connections and allow outgoing connections:

- `ufw default deny incoming`
- `ufw default allow outgoing`

Then, allow all incoming HTTP (port 80) and HTTPS (port 443) connections:

- `ufw allow http`
- `ufw allow https`

Also, allow Nginx HTTP/HTTPS: `ufw allow 'Nginx Full'`

### Create systemd unit file for the backend binary

### Create a reverse proxy

### Set up HTTPS
