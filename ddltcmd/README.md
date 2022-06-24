# ddltrans

## Install
> 	go install github.com/a631807682/ddltransform/ddltcmd@latest

## Usage
```
NAME:
   ddltcmd start - parse definition language and generate to code

USAGE:
   ddltcmd start [command options] [arguments...]

OPTIONS:
   --parser value, --ps value       use parser for parse ddl file (default: mysql)
   --path value, -p value           path for ddl file
   --transformer value, --tf value  use transformer for code generate (default: gorm)
```
## Examples
> ddltcmd start -ps sqlite -tf gorm -p ./sql.ddl