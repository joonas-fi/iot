Henkan poreamme
===============

```mermaid
flowchart TD
    subgraph "Poreamme"
        ctrl[Balboa control board]
        serialbridge[RS485 <-> Ethernet\nbridge]
    end
    subgraph "Raspberry Pi"
        homeassistant[Home Assistant]
        mqtt[MQTT\nmosquitto]
        bwalink
    end
    lan[Ethernet\nLAN]
    phonecontrol[Ohjaus puhelimesta] --> homeassistant
    ctrl --> serialbridge --> lan
    homeassistant --> bwalink --> lan
    mqtt <--> homeassistant
    mqtt <--> bwalink
```


Hardware
--------

### Wiring diagram

```mermaid
flowchart TD
    subgraph "Balboa control board"
        ctrlB[B/-]
        ctrlA[A/+]
        ctrlDC[+15.9 V]
        ctrlGND[Ground]
    end
    subgraph "RS485 bridge (data side)"
        bridgeSerialGND[GND]
        bridgeB[485B]
        bridgeA[485A]
        bridgeNC1[NC]
        bridgeNC2[NC]
    end
    subgraph "RS485 bridge (power & LAN side)"
        bridgeV[V+]
        bridgeGND[V-]
        bridgeRJ45[RJ45]
    end
    ctrlA --> bridgeA
    ctrlB --> bridgeB
    ctrlDC --> bridgeV
    ctrlGND --> bridgeGND
```

![](bridge.jpg)


### Bridge configuration

Looks like one can access it over web interface in port 80.

The URL is `http://192.168.1.254/`

| Configuration | Value  |
|---------------|--------|
| Baud rate     | 115200 |
| Data bits     | 8      |
| Parity        | None   |
| Stop bits     | 1      |
| Flow control  | None   |

The TCP port for serial port is `4196`

Troubleshooting:

- The "pwr" light should stay always lit
- When serial is connected properly and receiving data from the control board, the "act" light should blink in rapid succession
- The "link" light should be stably lit. It seemed to only light up after we connected ethernet, so it might indicate ethernet link?


Software
--------

Home Assistant -> MQTT

bwalink -> MQTT

### Install steps

- install mqtt
- install home assistant
- configure MQTT to home assistant
- install bwalink


### TODO

- static IP reservation for RS485 bridge
- Tailscale for raspberry pi
- Tailscale for Henkka


Links
-----

- [Balboa hottub integration without wifi module](https://community.openhab.org/t/balboa-hottub-integration-without-wifi-module/147110)
- [Homepage: RS485 <-> Ethernet](https://www.waveshare.com/wiki/RS485_TO_ETH_(B))
- [Amazon: RS485 <-> Ethernet](https://www.amazon.de/RS485-Ethernet-Converter-Industrial-Auto-Negotiation-transparent/dp/B09QMNWYLQ)
- [Protocol notes](https://github.com/ccutrer/balboa_worldwide_app/blob/main/doc/protocol.md)
- [Ready app](https://github.com/ccutrer/balboa_worldwide_app)
