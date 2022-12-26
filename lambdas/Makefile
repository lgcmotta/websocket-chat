#!make
include .env

.PHONY: build deploy


build:
	$(MAKE) -C connect build
	$(MAKE) -C disconnect build
	$(MAKE) -C join build
	$(MAKE) -C broadcast build
	$(MAKE) -C direct build
	$(MAKE) -C active-members build

deploy: build
	aws lambda update-function-code --function-name arn:aws:lambda:us-east-1:${ACCOUNT_ID}:function:chat-connect \
		--zip-file fileb://connect/main.zip \
		--publish
	aws lambda update-function-code --function-name arn:aws:lambda:us-east-1:${ACCOUNT_ID}:function:chat-disconnect \
		--zip-file fileb://disconnect/main.zip \
		--publish
	aws lambda update-function-code --function-name arn:aws:lambda:us-east-1:${ACCOUNT_ID}:function:chat-join \
		--zip-file fileb://join/main.zip \
		--publish
	aws lambda update-function-code --function-name arn:aws:lambda:us-east-1:${ACCOUNT_ID}:function:chat-broadcast \
		--zip-file fileb://broadcast/main.zip \
		--publish
	aws lambda update-function-code --function-name arn:aws:lambda:us-east-1:${ACCOUNT_ID}:function:chat-direct	\
		--zip-file fileb://direct/main.zip \
		--publish
	aws lambda update-function-code --function-name arn:aws:lambda:us-east-1:${ACCOUNT_ID}:function:chat-members	\
		--zip-file fileb://active-members/main.zip \
		--publish