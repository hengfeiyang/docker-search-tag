# docker-search-tag

list tags for docker search repository results.

## Usage

usage:


```
docker-search-tag NAME [TAG] [NUM]
```

list mysql tags:

```
docker-search-tag mysql
```

list mysql tags by filter 8. (search version 8.x)

```
docker-search-tag mysql 8.
```

or list 10 tags(default is 100):

```
docker-search-tag mysql 8. 10
docker-search-tag mysql _  10
```

> here, can use `_` igor second flag.


results:

```
Tags for mysql:

TAG            SIZE(MB)   LASTPUSH
latest              144   2021-12-02T11:43:26Z
8.0.4-rc             83   2018-03-14T08:42:43Z
8.0.4                83   2018-03-14T08:42:43Z
8.0.3               107   2018-02-17T09:43:00Z
8.0.1                86   2017-06-24T13:36:11Z
8.0.0               126   2017-04-10T23:29:13Z
8.0                 144   2021-12-02T11:43:21Z
8                   144   2021-12-02T11:43:20Z
5.7.36              147   2021-12-02T11:43:09Z
5.7.35              147   2021-10-12T16:42:35Z
5.6.28              106   2016-02-02T22:41:46Z
5.6.27              106   2015-12-17T03:17:56Z
5.6                  98   2021-12-02T11:43:04Z
5.5                  63   2019-05-10T23:43:32Z
5                   147   2021-12-02T11:42:49Z
```

## why

when i use `docker search name` find some images, it only tell me there ere some images like this:

```
NAME                              DESCRIPTION                                     STARS     OFFICIAL   AUTOMATED
mysql                             MySQL is a widely used, open-source relation…   11803     [OK]
mariadb                           MariaDB Server is a high performing open sou…   4492      [OK]
mysql/mysql-server                Optimized MySQL Server Docker images. Create…   885                  [OK]
percona                           Percona Server is a fork of the MySQL relati…   565       [OK]
phpmyadmin                        phpMyAdmin - A web interface for MySQL and M…   391       [OK]
centos/mysql-57-centos7           MySQL 5.7 SQL database server                   92
mysql/mysql-cluster               Experimental MySQL Cluster Docker images. Cr…   89
centurylink/mysql                 Image containing mysql. Optimized to be link…   59                   [OK]
databack/mysql-backup             Back up mysql databases to... anywhere!         54
prom/mysqld-exporter                                                              44                   [OK]
deitch/mysql-backup               REPLACED! Please use http://hub.docker.com/r…   41                   [OK]
tutum/mysql                       Base docker image to run a MySQL database se…   35
widdpim/mysql-client              Dockerized MySQL Client (5.7) including Curl…   1                    [OK]
```

then, i want to pull a image like this: `docker pull mysql:TAG`, but i don't know about tags, i must open browser visit `https://hub.docker.com/_/mysql?tab=tags` to find tags information, after then i can do pull with tag.

it is terible!!! 

i want to know the tag in terminal, like docker search tag xxx!

so here is it.
