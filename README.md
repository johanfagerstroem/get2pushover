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


Build
-----

    make
    make dist


Install (the traditional systemd way)
-------------------------------------

1. Download pre-built packages from [Releases](https://github.com/johanfagerstroem/get2pushover/releases) or build according to steps above.

2. Extract it under /opt:

        sudo tar xf get2pushover-VERSION-OS-ARCH.tar.gz -C /opt/ 

3. Modify the configuration file `/opt/get2pushover/config`.

4. Create a user under which the service should run:

        sudo useradd --system --no-create-home --shell /bin/false get2pushover

5. Setup systemd service and start it:

        sudo cp /opt/get2pushover/get2pushover.service /etc/systemd/system/
        sudo systemctl enable get2pushover.service
        sudo systemctl start get2pushover.service

6. Verify that the service is running:

        sudo systemctl status get2pushover.service


Build and run using Docker
--------------------------

1. Build Docker image:
    
        docker build -t get2pushover:$(git describe --always) .

2. Run:

        docker run -e "LISTEN_PORT=3333" -p 3333:3333 get2pushover

