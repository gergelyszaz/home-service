generate:	
	openapi-generator-cli generate \
	-i api/openapi.yaml \
	-g go-server \
	--git-repo-id home-service \
	--git-user-id gergelyszaz \
	--additional-properties=sourceFolder=gen \
	--additional-properties=packageName=openapi \
	--additional-properties=router=chi \


clean:
	rm -rd gen

serve:
	go run main.go

