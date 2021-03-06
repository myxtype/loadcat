# Loadcat

loadcat是一个nginx配置程序，允许您使用nginx作为负载平衡器。这个项目的灵感来自于各种在线的nginx负载平衡教程文章，以及linode的负载平衡服务nodebalancers的存在。到目前为止，该工具涵盖了一些HTTP和HTTPS负载平衡功能，例如SSL终止、动态添加服务器、将它们标记为不可用或根据需要进行备份，以及设置它们的权重以公平地分配负载。

## 安装

### Arch Linux

manually:

~~~
$ git clone https://aur.archlinux.org/loadcat.git
$ cd loadcat
$ makepkg
# pacman -U loadcat-0.1_alpha.1-1-x86_64.pkg.tar.xz
~~~

### From Source

Install Loadcat using the go get command:

```
$ go get github.com/myxtype/loadcat/cmd/loadcatd
```

## 使用

If you installed Loadcat from source, you can launch it with:

```
$ cd $GOPATH/src/github.com/myxtype/loadcat
# $GOPATH/bin/loadcatd
```

Loadcat parses a TOML encoded configuration file. In case one is not found, Loadcat will create one with same sane defaults. The location of the configuration file can be specified with the _-config_ flag.

Loadcat works by generating Nginx configuration files under a particular directory (a directory named "out" under Loadcat's working directory, as set in loadcat.conf). Nginx must be configured to load configuration files from this directory. For example on Arch Linux, when installed from AUR, Loadcat uses /var/lib/loadcat/out as the directory where generated Nginx configuration files are stored. You must include the following line inside the `http {}` block of /etc/nginx/nginx.conf to load configuration files generated by Loadcat:

```
include  /var/lib/loadcat/out/*/nginx.conf
```

Once Loadcat is running, you can navigate to http://{host}:26590 on your web browser, where _{host}_ is the domain name or IP address of the machine where Loadcat is running (for example http://localhost:26590 when running locally). To save a thousand words, here is a (kind of big) picture:

![4 steps primer](http://i.imgur.com/7l6zN5n.png)

Make sure that Nginx is running as a systemd service and is configured to load generated configuration files from the appropriate directory.

## Caution

As this is a very young project, and pretty experimental, you may encounter bugs and issues here and there. It would be really appreciated if you could open an issue outlining details of the bug or issue, or any feature that you would like to request. Any contribution or criticism (constructive or destructive) is really appreciated.

A lot of Nginx load balancing features is still not covered by this tool and at this moment that makes this rather limited in context of practical applications. However, solving this problem is just a matter of time. Although Nginx Plus specific features may have to wait for a while - at least until I get my hands on an instance or someone else with access to Nginx Plus starts contributing.

## Documentation

- [Reference](http://godoc.org/github.com/hjr265/loadcat)

## Resources

- [DigitalOcean's Tutorial on Using Nginx as a Load Balancer](https://www.digitalocean.com/community/tutorials/how-to-set-up-nginx-load-balancing)
- [Nginx's HTTP Load Balancing Documentation](http://nginx.org/en/docs/http/load_balancing.html)

## Contributing

Contributions are welcome.

## License

Loadcat is available under the [BSD (3-Clause) License](http://opensource.org/licenses/BSD-3-Clause).
