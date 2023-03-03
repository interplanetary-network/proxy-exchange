#!/bin/sh

DOMAIN=$1
DNSLINK=$2

hostnames=$(curl -s \
    -X GET \
    -H "Authorization: Bearer ${CF_API_TOKEN}" \
    -H "Content-Type: application/json" \
    "https://api.cloudflare.com/client/v4/zones/${CF_ZONE_ID}/web3/hostnames")

for i in $(echo $hostnames |jq -r -c '.result | .[]'); do
    target=$(echo $i |jq -r -c '.target')
    name=$(echo $i |jq -r -c '.name')
    id=$(echo $i |jq -r -c '.id')
    if [ $target == "ipfs" -a $name == "$DOMAIN" ]; then
        curl -s \
            -X PATCH \
            -H "Authorization: Bearer ${CF_API_TOKEN}" \
            -H "Content-Type: application/json" \
            -d "{\"dnslink\":\"${DNSLINK}\"}" \
            "https://api.cloudflare.com/client/v4/zones/${CF_ZONE_ID}/web3/hostnames/${id}" | jq -r -c '.result.status'
    fi
done
