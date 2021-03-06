# MyHue
Test go cli app to interact with the api of the philips hue 
using cobra and viper libs

Just to Learn go 
source: 
https://github.com/amimof/huego

## Usage ##

List lights:
```
Usage:
  myhue list [flags]

Flags:
  -f, --filter string   specify keyword to filter name
  -h, --help            help for list
  -r, --reachable       display reachable lights only
  -w, --width           width list
```

set light 
```
Usage:
  myhue set [flags]

Flags:
  -b, --bright int   brightness level
  -h, --help         help for set
  -l, --light int    light id
  -o, --on           Light on or off
```
After first run a user token wil be created and and stored in the config file \
before the first run you have to press the button on the hue bridge 

config file is stored in: **home\config.myhue.yaml**

## Examples ## 

```
myhue list
myhue list --width
myhue list -w --reachable --filter "name"

myhue set --light 1 --on
myhue set -l 1 --on -br 60
myhue set -l 1 

```
## environment vars ##

```
"HUE_USER":"someusername"
"HUE_IP":"1.1.1.1" (optional wil search for bridge)
"HUE_TOKEN":"token after first user creation" (after first run it will give you a token to set with this var)
```

# build
build exe: 
go build

# refs
```
https://github.com/amimof/huego
https://golang.org
https://github.com/spf13/cobra
https://github.com/spf13/viper
```
