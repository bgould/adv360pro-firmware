
MEMORY
{
  /* SoftDevice S140. See https://learn.adafruit.com/introducing-the-adafruit-nrf52840-feather/hathach-memory-map. */
  FLASH_TEXT (rw) : ORIGIN = 0x00000000+0x26000, LENGTH = 0xC6000 
  FLASH_APP  (rw) : ORIGIN = 0xEC000, LENGTH = 0x8000 /* Application data */
  RAM (xrw)       : ORIGIN = 0x20004180, LENGTH = 0x20040000-0x20004180 
}

_stack_size = 4K;

/* This value is needed by the Nordic SoftDevice. */
__app_ram_base = ORIGIN(RAM);

INCLUDE "targets/arm.ld"

/* For the flash API */
__flash_data_start = ORIGIN(FLASH_APP);
__flash_data_end = ORIGIN(FLASH_APP) + LENGTH(FLASH_APP);
