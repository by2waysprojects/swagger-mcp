package swagger

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/danishjsheikh/swagger-mcp/app/models"
)

func ExtractSchemaName(ref, schemaType string) string {
	if ref != "" {
		parts := strings.Split(ref, "/")
		return parts[len(parts)-1]
	}
	return schemaType
}

func getBaseURL(swaggerSpec models.SwaggerSpec) string {
	// For OpenAPI 3.0
	if swaggerSpec.OpenAPI != "" && len(swaggerSpec.Servers) > 0 {
		return strings.TrimSuffix(swaggerSpec.Servers[0].URL, "/")
	}

	// For Swagger 2.0
	baseURL := swaggerSpec.Host
	if !strings.HasPrefix(baseURL, "http://") && !strings.HasPrefix(baseURL, "https://") {
		baseURL = "https://" + baseURL
	}
	if swaggerSpec.BasePath != "" {
		baseURL = strings.TrimSuffix(baseURL, "/") + "/" + strings.TrimPrefix(swaggerSpec.BasePath, "/")
	}
	return baseURL
}

func ExtractSwagger(swaggerSpec models.SwaggerSpec) {
	resp := models.JsonRpcResponse{
		JSONRPC: "2.0",
		Result:  swaggerSpec,
		ID:      0,
	}
	encoder := json.NewEncoder(os.Stdout)
	if err := encoder.Encode(resp); err != nil {
		_, _ = os.Stderr.WriteString("Error encoding JSON: " + err.Error() + "\n")
	}
}
