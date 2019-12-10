package sentry

import "fmt"

type BotState uint8

const (
	BotStateInitializing BotState = iota
	BotStateStopped
	BotStateRunning
	BotStateStopping
)

func (bs BotState) String() string {
	return []string{
		"initializing",
		"stopped",
		"running",
		"stopping",
	}[bs]
}

func InitState() BotState {
	return BotStateInitializing
}

func nextState(bs BotState) (BotState, error) {
	switch bs {
	case BotStateInitializing:
		return BotStateStopped, nil
	case BotStateStopped:
		return BotStateRunning, nil
	case BotStateRunning:
		return BotStateStopping, nil
	case BotStateStopping:
		return BotStateStopped, nil
	default:
		return BotStateInitializing, fmt.Errorf("暂无下一状态: %s", bs.String())

	}
}
