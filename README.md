# Power Manager
I use this tiny go app for checking a few my local ip addresses. When they are all down, i am running a bash script to shutdown energy consuming virtual machines to increase my uptime of UPS.

When all ips are back online and pingable, it runs another script to restore all machine.

## Env
Create a .env file with the following content:

- HOSTS="ip1,ip2,ip3"
- SCRIPT_DOWN=path_for_script
- SCRIPT_UP=path_for_script

SCRIPT_DOWN will be runned one time when the all ips are down and UP script will be runned when they are back online
