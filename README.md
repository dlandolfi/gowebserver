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

Create systemd unit at`/lib/systemd/system/service_name.service`

```
[Unit]  
Description=backend binary
    
[Service]
Type=simple  
Restart=always  
RestartSec=5s  
ExecStart=/path/to/binary
	
[Install]
WantedBy=multi-user.target
```

Afterwards, be sure to run: 

- `systemctl enable service_name` to enable the systemd service so it always starts at startup 
- `systemctl start service_name` to start it 

This will start the go backend binary at start up and will also try restarting every 5 seconds if it happens to crash

### Create a reverse proxy

Create reverse proxy at `etc/nginx/conf.d/gowebserver.conf`

``` 
server {  
  server_name <domain or ip>;  
 
  location / {  
    proxy_pass http://localhost:8080;  
  }  
}
```
This will allow nginx to forward requests to the go server

### Set up HTTPS

HTTPS was set up using cerbot
https://certbot.eff.org/
