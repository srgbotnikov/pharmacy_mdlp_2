cat $1 | base64 > b64.txt && cat b64.txt | tr -d '\r\n' > $2
