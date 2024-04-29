#!/usr/bin/env bash

#https://192.168.13.91/xui/
#
#admin
#MsTac@2001


# send to post request, as the follows:
curl -k -u admin:MsTac@2001 \
-H "Content-Type: application/json" \
-X POST \
-d '{"name": "vs-go-instance3", "destination": "192.159.14.35:350", "mask": "255.255.255.255",  "sourceAddressTranslation": {"type": "automap"}}' \
https://192.168.13.91/mgmt/tm/ltm/virtual



curl -k -u admin:MsTac@2001 \
-H "Content-Type: application/json" \
-X PUT \
-d '{"connectionLimit": 1000}'
https://192.168.13.91/mgmt/tm/ltm/virtual/~Common~hello-vs1