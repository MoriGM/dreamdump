# dreamdump [![Go](https://github.com/MoriGM/dreamdump/actions/workflows/go.yml/badge.svg?branch=main&event=push)](https://github.com/MoriGM/dreamdump/actions/workflows/go.yml)
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


# Supported Drives

Supported drives are in drive/list.go


# Special Thanks

 martin korth (nocash) for General CD-Rom Info https://problemkaputt.de/psxspx-contents.htm   
 superg for redumper (Explanation of Track Splitting and Unscrambling) https://github.com/superg/redumper   
 saramibreak for DiscImageCreator (Explanation of GD-Rom TOC) https://github.com/saramibreak/DiscImageCreator  


 # Gratitude

 The many testers who got this programm to run and work. The nice people at VGPC. Maddog for doing the biggest part of testing and giving most of the infos and helping with feature ideas.