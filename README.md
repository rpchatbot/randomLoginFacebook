# randomLoginFacebook

## Description
### Stop service:
```bash
sudo systemctl stop myapp.service
```

### Build app:
```bash
go build -o myapp
```

### Move to bin:
```bash
sudo mv myapp /usr/local/bin/
```

### Reload:
```bash
sudo systemctl daemon-reload
```

### Start service:
```bash
sudo systemctl start myapp.service
```

### Check status:
```bash
sudo systemctl status myapp.service
```