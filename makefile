test:
	go test -shuffle=on --tags=unit ./... -v

#test:
#	echo $(name)
#	go test -shuffle=on --tags=unit ./... -v -run $(name)
