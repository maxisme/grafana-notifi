Install caddy:
```
$ git clone -b v2 "https://github.com/caddyserver/caddy.git"
$ cd caddy/cmd/caddy/
$ go build -o /usr/bin/caddy
$ groupadd --system caddy
$ useradd --system \
  	--gid caddy \
  	--create-home \
  	--home-dir /var/lib/caddy \
  	--shell /usr/sbin/nologin \
  	--comment "Caddy web server" \
  	caddy
```
Then move [`caddy.service`](https://raw.githubusercontent.com/caddyserver/dist/master/init/caddy.service) to `/etc/systemd/system/caddy.service`:
```
$ cp caddy.service /etc/systemd/system/caddy.service
$ systemctl daemon-reload
$ systemctl enable caddy
$ systemctl start caddy
$ systemctl status caddy
```

then to create the service
```
cd /root/
git clone https://github.com/maxisme/grafana-notifi
go build . -o /usr/local/bin/grafana-notifi
```
then to start the service
```
$ cp grafananotifi.service /etc/systemd/system/
$ cp grafananotifi.socket /etc/systemd/system/
```
```
$ systemctl daemon-reload
$ systemctl enable grafananotifi.socket
$ systemctl start grafananotifi.socket
$ systemctl status grafananotifi.socket
```