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
-d '{"source": "0.0.0.0/32", "mask": "255.255.255.255", "connectionLimit": 1000}'
https://192.168.13.91/mgmt/tm/ltm/virtual/~Common~hello-vs1


curl -k -u your_username:your_password -H "Content-Type: application/json" \
-X PATCH \
-d '[{"op": "replace", "path": "/connectionLimit", "value": 1000}]' \
https://your_bigip_ip_address/mgmt/tm/ltm/virtual/~Common~your_virtual_server_name


curl -k -u admin:MsTac@2001 \
-H "Content-Type: application/json" \
-X PATCH  \
-d '{"disabled": true}' \
https://192.168.13.91/mgmt/tm/ltm/virtual/~Common~hello-vs1

export BIGIP_USERNAME=admin
export BIGIP_PASSWORD=MsTac@2001
export BIGIP_ADDRESS=192.168.13.91

#curl -sk -u ${BIGIP_USERNAME}:${BIGIP_PASSWORD} -H "Content-Type: application/json" -X POST -d '{"name": "<irule_name>", "apiAnonymous": "<irule_definition>"}' https://${BIGIP_ADDRESS}/mgmt/tm/ltm/rule

curl -sk -u ${BIGIP_USERNAME}:${BIGIP_PASSWORD} \
-H "Content-Type: application/json" \
-X POST \
-d '{"name": "MyRedirectRule", "apiAnonymous": "when HTTP_REQUEST {\n SWITCH -glob [string tolower [HTTP::uri]] {\n        \"/example*\" {\n             HTTP::redirect http://example.org[HTTP::uri]\n        }\n\n        default {\n            HTTP::redirect http://example.com[HTTP::uri]\n        }\n    }\n}"}' \
https://${BIGIP_ADDRESS}/mgmt/tm/ltm/rule

curl -sk -u ${BIGIP_USERNAME}:${BIGIP_PASSWORD} -H "Content-Type: application/json" -X DELETE https://${BIGIP_ADDRESS}/mgmt/tm/ltm/pool/~Common~hello-pool



curl -sku admin:MsTac@2001 -X POST -H "Content-Type: application/json" -d '{"name": "example_ifile", "sourcePath": "http://example.com/path/to/your-file"}' https://192.168.13.91/mgmt/tm/ltm/ifile

curl -sk -u ${BIGIP_USERNAME}:${BIGIP_PASSWORD} -H "Content-Type: application/json" -X PUT -d '{"monitor": "<monitor_state>"}' https://${BIGIP_ADDRESS}/mgmt/tm/ltm/pool/~Common~<pool_name>/members/~Common~<member_address>:<member_port>

curl -sk -u ${BIGIP_USERNAME}:${BIGIP_PASSWORD} -H "Content-Type: application/json" -X POST -d '{"name": "<member_address>:<member_port>"}' https://${BIGIP_ADDRESS}/mgmt/tm/ltm/pool/~Common~<pool_name>/members

curl -sk -u ${BIGIP_USERNAME}:${BIGIP_PASSWORD} -H "Content-Type: application/json" -X POST -d '{"name": "MyRedirectRule", "apiAnonymous": "when HTTP_REQUEST {\n set req_uri [string tolower [HTTP::uri]]\n if { [string match \"/example*\" \$req_uri] } {\n        HTTP::redirect http://example.org[HTTP::uri]\n    } else {\n        HTTP::redirect http://example.com[HTTP::uri]\n    }\n}"}' https://${BIGIP_ADDRESS}/mgmt/tm/ltm/rule