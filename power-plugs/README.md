Provisioning completely new device
---------------------------------

- Plug it in the wall
- From phone or computer connect to wifi named `athom-smart-plug-...`
- Give it the wifi details to which it will connect to (the home automation wifi network)
- Look for the IP from your network and then use the flashing

Flashing
--------

to flash:

```shell
docker run --rm -v "$(pwd):/config" -it esphome/esphome run --device 192.168.1.8 DEVICE_ID.yaml
```
