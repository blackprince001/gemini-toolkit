package types

type ToolKitConfig struct {
	ApiKey    string
	BaseModel string
}

type ToolKitService struct {
	Config ToolKitConfig
}

func NewToolKitService(cfg ToolKitConfig) *ToolKitService {
	return &ToolKitService{
		Config: cfg,
	}
}
