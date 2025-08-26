package models

// MCPConfig represents the top-level MCP server configuration
type MCPConfig struct {
	ToolSet *ToolSetConfig `yaml:"toolSet,omitempty" json:"toolSet,omitempty"`
	Server  ServerConfig   `yaml:"server,omitempty" json:"server,omitempty"`
	Tools   []Tool         `yaml:"tools,omitempty" json:"tools,omitempty"`
}

// ToolSetConfig defines the configuration for a toolset.
type ToolSetConfig struct {
	Name        string             `json:"name,omitempty"`
	ServerTools []ServerToolConfig `json:"serverTools,omitempty"`
}

// ServerToolConfig specifies which tools from a server to include in a toolset.
type ServerToolConfig struct {
	ServerName string   `json:"serverName,omitempty"`
	Tools      []string `json:"tools,omitempty"`
}

// ServerConfig represents the MCP server configuration
type ServerConfig struct {
	Name            string           `yaml:"name" json:"name"`
	BaseURL         string           `yaml:"baseURL,omitempty" json:"baseURL,omitempty"`
	Config          map[string]any   `yaml:"config,omitempty" json:"config,omitempty"`
	AllowTools      []string         `yaml:"allowTools,omitempty" json:"allowTools,omitempty"`
	SecuritySchemes []SecurityScheme `yaml:"securitySchemes,omitempty" json:"securitySchemes,omitempty"`
}

// SecurityScheme defines a security scheme that can be used by the tools.
type SecurityScheme struct {
	ID                string `yaml:"id" json:"id"`
	Type              string `yaml:"type" json:"type"`                         // e.g., "http", "apiKey", "oauth2", "openIdConnect"
	Scheme            string `yaml:"scheme,omitempty" json:"scheme,omitempty"` // e.g., "basic", "bearer" for "http" type
	In                string `yaml:"in,omitempty" json:"in,omitempty"`         // e.g., "header", "query", "cookie" for "apiKey" type
	Name              string `yaml:"name,omitempty" json:"name,omitempty"`     // Name of the header, query parameter or cookie for "apiKey" type
	DefaultCredential string `yaml:"defaultCredential,omitempty" json:"defaultCredential,omitempty"`
}

// Tool represents an MCP tool configuration
type Tool struct {
	Name                  string                   `yaml:"name" json:"name"`
	Description           string                   `yaml:"description" json:"description"`
	Args                  []Arg                    `yaml:"args" json:"args"`
	RequestTemplate       RequestTemplate          `yaml:"requestTemplate" json:"requestTemplate,omitempty"`
	ResponseTemplate      ResponseTemplate         `yaml:"responseTemplate" json:"responseTemplate,omitempty"`
	ErrorResponseTemplate *string                  `yaml:"errorResponseTemplate,omitempty" json:"errorResponseTemplate,omitempty"`
	Security              *ToolSecurityRequirement `yaml:"security,omitempty" json:"security,omitempty"`
}

// Arg represents an MCP tool argument
type Arg struct {
	Name        string         `yaml:"name" json:"name"`
	Description string         `yaml:"description" json:"description"`
	Type        string         `yaml:"type,omitempty" json:"type,omitempty"`
	Required    bool           `yaml:"required,omitempty" json:"required,omitempty"`
	Default     any            `yaml:"default,omitempty" json:"default,omitempty"`
	Enum        []any          `yaml:"enum,omitempty" json:"enum,omitempty"`
	Items       map[string]any `yaml:"items,omitempty" json:"items,omitempty"`
	Properties  map[string]any `yaml:"properties,omitempty" json:"properties,omitempty"`
	Position    string         `yaml:"position,omitempty" json:"position,omitempty"`
}

// RequestTemplate represents the MCP request template
type RequestTemplate struct {
	URL            string                   `yaml:"url" json:"url"`
	Method         string                   `yaml:"method" json:"method"`
	Headers        []Header                 `yaml:"headers,omitempty" json:"headers,omitempty"`
	Body           string                   `yaml:"body,omitempty" json:"body,omitempty"`
	ArgsToJsonBody bool                     `yaml:"argsToJsonBody,omitempty" json:"argsToJsonBody,omitempty"`
	ArgsToUrlParam bool                     `yaml:"argsToUrlParam,omitempty" json:"argsToUrlParam,omitempty"`
	ArgsToFormBody bool                     `yaml:"argsToFormBody,omitempty" json:"argsToFormBody,omitempty"`
	Security       *ToolSecurityRequirement `yaml:"security,omitempty" json:"security,omitempty"`
}

// ToolSecurityRequirement specifies a security scheme requirement for a tool.
type ToolSecurityRequirement struct {
	ID          string `yaml:"id" json:"id"`                                       // References a SecurityScheme ID defined in ServerConfig.SecuritySchemes
	Passthrough bool   `yaml:"passthrough,omitempty" json:"passthrough,omitempty"` // Whether to pass through the security credentials
}

// Header represents an HTTP header
type Header struct {
	Key   string `yaml:"key" json:"key"`
	Value string `yaml:"value" json:"value"`
}

// ResponseTemplate represents the MCP response template
type ResponseTemplate struct {
	Body        string `yaml:"body,omitempty" json:"body,omitempty"`
	PrependBody string `yaml:"prependBody,omitempty" json:"prependBody,omitempty"`
	AppendBody  string `yaml:"appendBody,omitempty" json:"appendBody,omitempty"`
}

// ConvertOptions represents options for the conversion process
type ConvertOptions struct {
	ServerName     string                 `json:"serverName"`
	ServerConfig   map[string]interface{} `json:"serverConfig"`
	ToolNamePrefix string                 `json:"toolNamePrefix"`
	TemplatePath   string                 `json:"templatePath"`
}

// ToolTemplate represents a template for applying to all tools
type ToolTemplate struct {
	RequestTemplate  *RequestTemplate         `yaml:"requestTemplate,omitempty" json:"requestTemplate,omitempty"`
	ResponseTemplate *ResponseTemplate        `yaml:"responseTemplate,omitempty" json:"responseTemplate,omitempty"`
	Security         *ToolSecurityRequirement `yaml:"security,omitempty" json:"security,omitempty"`
}

// MCPConfigTemplate represents a template for patching the generated config
type MCPConfigTemplate struct {
	Server ServerConfig `yaml:"server" json:"server"`
	Tools  ToolTemplate `yaml:"tools,omitempty" json:"tools,omitempty"`
}
