package service

type INotificationRepository interface {
	PublishMessage(queueName string, message string) error
}

type NotificationService struct {
	repo INotificationRepository
}

func InitNotificationService(repo INotificationRepository) *NotificationService {
	return &NotificationService{
		repo: repo,
	}
}

func (svc *NotificationService) SendNotification(queueName string, message string) error {
	if err := svc.repo.PublishMessage(queueName, message); err != nil {
		return err
	}
	return nil
}
