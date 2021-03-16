TOKEN := d36c4ebe-91a4-4c30-8ce7-b87ed39fe6a2


create:
	curl -X 'POST' \
	'http://35.72.193.240:8080/user/create' \
	-H 'accept: application/json' \
	-H 'Content-Type: application/json' \
	-d '{"name": "string"}' | jq

get:
	curl -X 'GET' \
	'http://35.72.193.240:8080/user/get' \
	-H 'accept: application/json' \
	-H 'x-token: $(TOKEN)'

status:
	curl -X 'PUT' \
	'http://35.72.193.240:8080/user/update' \
	-H 'accept: */*' \
	-H 'x-token: $(TOKEN)' \
	-H 'Content-Type: application/json' \
	-d '{"name": "string"}'

game:
	curl -X 'POST' \
	'http://35.72.193.240:8080/game/finish' \
	-H 'accept: application/json' \
	-H 'x-token: $(TOKEN)' \
	-H 'Content-Type: application/json' \
	-d '{"score": 10001}'| jq

gacha:
	curl -X 'POST' \
	'http://35.72.193.240:8080/gacha/draw' \
	-H 'accept: application/json' \
	-H 'x-token: $(TOKEN)' \
	-H 'Content-Type: application/json' \
	-d '{"times": 1}' | jq

ranking:
	curl -X 'POST' \
	'http://35.72.193.240:8080/ranking' \
	-H 'accept: application/json' \
	-d '{"start": 1}' | jq

collection:
	curl -X 'GET' \
	'http://35.72.193.240:8080/collection' \
	-H 'accept: application/json' \
	-H 'x-token: $(TOKEN)' | jq

g_status:
	curl -X 'GET' \
	'http://35.72.193.240:8080/user/get_status' \
	-H 'accept: */*' \
	-H 'x-token: $(TOKEN)' \
	-H 'Content-Type: application/json' | jq

