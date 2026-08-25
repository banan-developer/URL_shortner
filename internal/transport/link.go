package transport

import "URL_shortner/internal/service"

type LinksTransport struct {
	service *service.LinksService
}

func NewLinksTransport(service *service.LinksService) *LinksTransport {
	return &LinksTransport{
		service: service,
	}
}
