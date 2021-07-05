#!/bin/sh
/opt/cprocsp/bin/amd64/cryptcp -sign -pin 12345 -detached -nochain -thumbprint $1 $2 $3 && cat $3 | tr -d '\r\n' > $4