get2pushover
============

get2pushover is an API proxy that converts HTTP GET requests to Pushover's
POST-based API. The typical use-case is when someone wants to send a Pushover
notification from a device on the network (e.g. an Axis camera, a QNAP NAS,
etc.) that doesn't support HTTP POST in its event notification system. 

Deploy this API proxy on a machine on the network and let the devices send
notifications through it.

Pre-built packages are provided for amd64.


HTTP API
--------
    
    GET /?token=APPTOKEN&user=USERTOKEN&title=MESSAGETITLE&message=MESSAGE HTTP/1.1

If *token* or *user* isn't provided in the request, the defaults from the
configuration file is used (environment variables PUSHOVER_DEFAULT_USER and
PUSHOVER_DEFAULT_TOKEN).

If *title* isn't provided in the request, the proxy defaults to the caller's
FQDN. If FQDN can't be resolved, it defaults to *get2pushover*


Install (the traditional systemd way)
-------------------------------------

1. Build the distribution package or download pre-built packages from Releases:

        make && make dist

2. Extract it under /opt:

        sudo tar czvf -C /opt/ get2pushover-VERSION.tar.gz

3. Modify the configuration file `/opt/get2pushover/config`.

4. Create a user under which the service should run:

        useradd -r get2pushover

5. Setup systemd service and start it:

        sudo cp /opt/get2pushover-VERSION/get2pushover.service /etc/systemd/system/
        sudo systemctl enable get2pushover.service
        sudo systemctl start get2pushover.service

6. Verify that the service is running:

        sudo systemctl status get2pushover.service


Build and run using Docker
--------------------------

Coming.
