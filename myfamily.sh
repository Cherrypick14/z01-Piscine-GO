curl -s https://learn.zone01kisumu.ke/assets/superhero/all.json | jq --raw-output --arg id "$HERO_ID" ' .[] | select (.id==($id | tonumber)) | .connections.relatives | sub("\n";"\\n")'

