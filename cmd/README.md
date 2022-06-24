# ddltrans

## Install
> 	go install github.com/a631807682/ddltransform/cmd@latest

## Usage
```
NAME:
   ddltrans start - parse definition language and generate to code

USAGE:
   ddltrans start [command options] [arguments...]

OPTIONS:
   --parser value, --ps value       use parser for parse ddl file (default: mysql)
   --path value, -p value           path for ddl file
   --transformer value, --tf value  use transformer for code generate (default: gorm)
```
## Examples
> ddltrans start -ps sqlite -tf gorm -p ./sql.ddl