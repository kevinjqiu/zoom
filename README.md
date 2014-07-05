```
   ____    ___     ___   __  __  
  |_  /   / _ \   / _ \ |  \/  | 
   / /   | (_) | | (_) || |\/| | 
  /___|   \___/   \___/ |_|__|_| 
_|"""""|_|"""""|_|"""""|_|"""""| 
"`-0-0-'"`-0-0-'"`-0-0-'"`-0-0-' 
```

A GeoLocation web service using maxmind.com geoip database

[![Build Status](https://travis-ci.org/kevinjqiu/zoom.svg?branch=master)](https://travis-ci.org/kevinjqiu/zoom)

Command Line Interface
----------------------

### Web API

To start a zoom web api server:

    zoom serve

You can specify the port using `--port` option:

    zoom serve --port 8008

### Update GeoLite2 database

    zoom update

### Query directly on the command line

    zoom query <ip-address>

Attribution
-----------

This product includes GeoLite2 data created by MaxMind, available from [www.maxmind.com](http://www.maxmind.com).

License
-------

Zoom is licensed under MIT.
