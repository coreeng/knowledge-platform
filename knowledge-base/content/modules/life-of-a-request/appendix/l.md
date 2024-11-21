+++
title = "L - Nginx configuration"
weight = 12
chapter = false
+++

The following snippet is the nginx configuration. It is configured as a reverse proxy running on port **80** for the **web-app** container running on port **8080**. Nginx is responsible for serving static content. 

```
server {
    listen 80;

    # Disable access logging for this server
    access_log off;

    location / {
        root /usr/share/nginx/html;
        index index.html;
    }

    location /web-app/ {
        proxy_pass http://web-app:8080/;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }
}
```
