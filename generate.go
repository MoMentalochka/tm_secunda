package tm_secunda

//go:generate echo "  🧪  Собираем OpenAPI бандл для генерации"
//go:generate npx @redocly/cli bundle shared/api/auth/v1/auth.openapi.yaml --output shared/api/bundles/gen-auth-openapi.yaml

//go:generate  echo "  📁  Генерируем из OpenAPI"
//go:generate go run github.com/ogen-go/ogen/cmd/ogen --target ./shared/pkg/openapi/auth/v1 --package auth --clean ./shared/api/bundles/gen-auth-openapi.yaml
