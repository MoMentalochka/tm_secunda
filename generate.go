package tm_secunda

//@TODO migrate to TASKFILE
//go:generate echo "  🧪  Собираем OpenAPI бандл для генерации"
//go:generate npx @redocly/cli bundle shared/api/tm/v1/tm.openapi.yaml --output shared/api/bundles/openapi-bundle.yaml

//go:generate  echo "  📁  Генерируем из OpenAPI"
//go:generate go run github.com/ogen-go/ogen/cmd/ogen --target ./shared/pkg/openapi/tm/v1 --package auth --clean ./shared/api/bundles/openapi-bundle.yaml
