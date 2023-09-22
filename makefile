RabbitStart:
	docker start e-petitions_rabbitMQ

NotificationBuild:
	go build -o ./build/Notification.exe ./Notification/main.go

NotificationStart:
	./build/Notification.exe

StartService: RabbitStart NotificationStart