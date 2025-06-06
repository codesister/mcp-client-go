package googlemap

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
)

const (
	NpxGooglemapMcpServer    = "npx-googlemap-mcp-server"
	DockerGooglemapMcpServer = "docker-googlemap-mcp-server"
)

type GoogleMapParam struct {
	GooglemapApiKey string
}

func InitGooglemapMCPClient(p *GoogleMapParam, options ...param.Option) *param.MCPClientConf {

	googlemapMCPClient := &param.MCPClientConf{
		Name: NpxGooglemapMcpServer,
		StdioClientConf: &param.StdioClientConfig{
			Command: "npx",
			Env: []string{
				"GOOGLE_MAPS_API_KEY=" + p.GooglemapApiKey,
			},
			Args: []string{
				"-y",
				"@modelcontextprotocol/server-google-maps",
			},
			InitReq: mcp.InitializeRequest{},
		},
	}

	for _, o := range options {
		o(googlemapMCPClient)
	}

	if googlemapMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion == "" {
		googlemapMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	}

	if googlemapMCPClient.StdioClientConf.InitReq.Params.ClientInfo.Name == "" {
		googlemapMCPClient.StdioClientConf.InitReq.Params.ClientInfo = mcp.Implementation{
			Name:    "mcp-server/googlemap",
			Version: "0.1.0",
		}
	}

	return googlemapMCPClient
}

func InitDockerGooglemapMCPClient(p *GoogleMapParam, options ...param.Option) *param.MCPClientConf {

	googlemapMCPClient := &param.MCPClientConf{
		Name: DockerGooglemapMcpServer,
		StdioClientConf: &param.StdioClientConfig{
			Command: "docker",
			Env: []string{
				"GOOGLE_MAPS_API_KEY=" + p.GooglemapApiKey,
			},
			Args: []string{
				"run",
				"-i",
				"--rm",
				"-e",
				"GOOGLE_MAPS_API_KEY",
				"mcp/google-maps",
			},
			InitReq: mcp.InitializeRequest{},
		},
	}

	for _, o := range options {
		o(googlemapMCPClient)
	}

	if googlemapMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion == "" {
		googlemapMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	}

	if googlemapMCPClient.StdioClientConf.InitReq.Params.ClientInfo.Name == "" {
		googlemapMCPClient.StdioClientConf.InitReq.Params.ClientInfo = mcp.Implementation{
			Name:    "mcp-server/googlemap",
			Version: "0.1.0",
		}
	}

	return googlemapMCPClient
}
