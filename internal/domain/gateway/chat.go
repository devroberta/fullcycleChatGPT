package gateway

import (
	"context"

	"github.com/devroberta/fullcycleChatGPT/internal/domain/entity"
)

type ChatGateway interface {
	CreateChat(ctx context.Context, chat *entity.Chat) error
	FindChatByID(ctx context.Context, chaID string) (*entity.Chat, error)
	SaveChat(ctx context.Context, chat *entity.Chat) error
}
