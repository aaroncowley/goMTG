#!/bin/bash

# grab from mtgjson
curl https://mtgjson.com/json/AllCards.json.zip > AllCards.json.zip
# unzip
unzip AllCards.json.zip
#use jq to chop leading name off cards and convert to csv
jq -r 'to_entries[] | [(.value) ] | @json' AllCards.json > AllCards.csv
#use mongo import to put into your db
mongoimport --db mtgdb --collection cards --type csv --file AllCards.csv --headerline
