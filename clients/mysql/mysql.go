package mysql

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
)

const (
	NpxMysqlMcpServer = "npx-mysql-mcp-server"
)

type MysqlParam struct {
	MysqlHost               string
	MysqlPort               string
	MysqlUser               string
	MysqlPass               string
	MysqlDb                 string
	Path                    string
	MysqlPoolSize           string
	MysqlQueryTimeout       string
	MysqlCacheTTL           string
	MysqlRateLimit          string
	MysqlMaxQueryComplexity string
	MysqlSSL                string
	MysqlEnableLogging      string
	MysqlLogLevel           string
	MysqlMetricsEnable      string
	AllowInsertOperation    string
	AllowUpdateOperation    string
	AllowDeleteOperation    string
}

func InitMysqlMCPClient(p *MysqlParam, options ...param.Option) *param.MCPClientConf {

	mysqlMCPClient := &param.MCPClientConf{
		Name: NpxMysqlMcpServer,
		StdioClientConf: &param.StdioClientConfig{
			Command: "npx",
			Env: []string{
				// Basic connection settings
				"MYSQL_HOST=" + p.MysqlHost,
				"MYSQL_PORT=" + p.MysqlPort,
				"MYSQL_USER=" + p.MysqlUser,
				"MYSQL_PASS=" + p.MysqlPass,
				"MYSQL_DB=" + p.MysqlDb,
				"PATH=" + p.Path,

				// Performance settings
				"MYSQL_POOL_SIZE=" + p.MysqlPoolSize,
				"MYSQL_QUERY_TIMEOUT=" + p.MysqlQueryTimeout,
				"MYSQL_CACHE_TTL=" + p.MysqlCacheTTL,

				// Security settings
				"MYSQL_RATE_LIMIT=" + p.MysqlRateLimit,
				"MYSQL_MAX_QUERY_COMPLEXITY=" + p.MysqlMaxQueryComplexity,
				"MYSQL_SSL=" + p.MysqlSSL,

				// Monitoring settings
				"MYSQL_ENABLE_LOGGING=" + p.MysqlEnableLogging,
				"MYSQL_LOG_LEVEL=" + p.MysqlLogLevel,
				"MYSQL_METRICS_ENABLED=" + p.MysqlMetricsEnable,

				// Write operation flags
				"ALLOW_INSERT_OPERATION=" + p.AllowInsertOperation,
				"ALLOW_UPDATE_OPERATION=" + p.AllowUpdateOperation,
				"ALLOW_DELETE_OPERATION=" + p.AllowDeleteOperation,
			},
			Args: []string{
				"-y",
				"@benborla29/mcp-server-mysql",
			},
			InitReq: mcp.InitializeRequest{},
		},
	}

	for _, o := range options {
		o(mysqlMCPClient)
	}

	if mysqlMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion == "" {
		mysqlMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	}

	if mysqlMCPClient.StdioClientConf.InitReq.Params.ClientInfo.Name == "" {
		mysqlMCPClient.StdioClientConf.InitReq.Params.ClientInfo = mcp.Implementation{
			Name:    "mcp-server/mysql",
			Version: "0.1.0",
		}
	}

	return mysqlMCPClient
}
