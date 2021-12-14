# docker-search-tag

list tags for docker search repository results.

## Usage

usage:

```
docker-search-tag mysql
```

results:

```
Tags for mysql:
-------------------------------
Tag                  Size(MB)
--------------------------------
latest               144
8.0.4-rc             83
8.0.4                83
8.0.3                107
8.0.27               144
8.0.26               143
8.0.25               154
...
5.5.52               83
5.5.51               83
5.5.50               83
5.5.49               83
5.5.46               83
5.5                  63
5                    147
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
