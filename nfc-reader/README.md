reader
------

https://www.ehuoyan.com/english/products_display.asp?pid=19&id=220&proid=202

supported RFID tech:

- [NTAG213](https://www.nxp.com/products/NTAG213_215_216) was tested to work
  * [purchase link](https://www.amazon.de/-/en/gp/product/B0DSZRCPZX)
- also credit cards like Visa seem to respond



connect reader to USB
---------------------

check it working:

```console
$ lsusb
Bus 001 Device 014: ID 10c4:ea60 Silicon Labs CP210x UART Bridge

$ ls /dev/serial/by-id/
usb-Silicon_Labs_CP2102_USB_to_UART_Bridge_Controller_0001-if00-port0
```


hardware behaviour
------------------

it doesn't print to serial anything / read tags unless you send it command to do so.
(so just reading from serial port and tapping a tag is not expected to produce anything)


get software
-----------

need to use python2 because that's what the code was written against

```shell
git clone https://github.com/mdeverdelhan/YHY523U-driver.git
cd YHY523U-driver/
docker run --rm -it --entrypoint=bash -v "$(pwd):/workspace" --workdir=/workspace --device=/dev/ttyUSB0 python:2
pip install pyserial
python yhy523u.py
```

> [!NOTE]
> the script as-is doesn't do anything. you need to uncomment the example codes from the main entrypoint


todo
----

- port to Go, have daemon that provides server-sent events over HTTP?
