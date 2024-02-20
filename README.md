Advantage 360 Pro Keyboard Firmware
===================================

This is open-source firmware for the Kinesis Advantage 360 Pro keyboard. It is
developed with [TinyGo][tinygo] and written in the [Go Programming Language][golang].

Compared to the ZMK fork that is maintained by Kinesis, major functionality is
lacking, such as Bluetooth and macros.  

Building Firmware
-----------------

In order to build firmware in this repository, please install [Go][golang] 1.21+ and [TinyGo][tinygo] 0.31+.

The two halves of the keyboard each have separate firmware files (.uf2) to use for flashing.

```
tinygo build -target=./targets/adv360pro-left.json -size=short -o adv360pro-left.uf2
tinygo build -target=./targets/adv360pro-right.json -size=short -o adv360pro-right.uf2
```

License
-----------------------

Licensed GPLv3 or later.

Copyright © 2019-2023 Benjamin Gould

Licensed GPLv3 or later.
[golang]: https://golang.org/
[tinygo]: https://tinygo.org/