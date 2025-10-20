# dreamdump
A small go project to maybe dump dreamcast games on linux and windows

Usage

```
./dreamdump disc --drive="driveName" --image-name="" --image-path=""
```


Supported commands:
`disc`
`split`

Supported options:   
`--drive=<drivePath>`   
`--sector-order=<sectorOrder>`   
`--read-offset=<number>`   
`--image-path=<name>`   
`--image-name=<name>`   
`--cutoff=<cutoff>`   
`--read-at-once=<number[1-20(Linux allows 40)]>`   
`--speed=<number>`   
`--force-qtoc`   
`--retries=<number[1-255]>`   