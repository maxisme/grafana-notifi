Install caddy:
```
$ git clone -b v2 "https://github.com/caddyserver/caddy.git"
$ cd caddy/cmd/caddy/
$ go build -o /usr/local/bin/caddy
$ groupadd --system caddy
$ useradd --system \
  	--gid caddy \
  	--create-home \
  	--home-dir /var/lib/caddy \
  	--shell /usr/sbin/nologin \
  	--comment "Caddy web server" \
  	caddy
```
Then move `caddy.service` to `/etc/systemd/system/caddy.service`:
```
$ cp caddy.service /etc/systemd/system/caddy.service
$ systemctl daemon-reload
$ systemctl enable caddy
$ systemctl start caddy
$ systemctl status caddy
```