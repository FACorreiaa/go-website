# Generate templ files
templ:
	templ generate

# Generate CSS
tailwind:
	npx tailwindcss -i ./assets/css/input.css -o ./assets/css/output.css --minify

# Clean previous builds
clean:
	rm -rf build/

# Build for Cloudflare Workers
build: clean templ tailwind
	mkdir -p build
	go run github.com/syumai/workers/cmd/workers-assets-gen@latest
	GOOS=js GOARCH=wasm go build -o ./build/app.wasm main.go

# Development server
server:
	air \
	--build.cmd "go build -o tmp/bin/main ./main.go" \
	--build.bin "tmp/bin/main" \
	--build.delay "100" \
	--build.exclude_dir "node_modules" \
	--build.include_ext "go" \
	--build.stop_on_error "false" \
	--misc.clean_on_exit true

# Development mode with file watching
dev:
	make tailwind
	make -j3 tailwind-watch templ-watch server

# Watch for templ changes
templ-watch:
	templ generate --watch --proxy="http://localhost:8090" --open-browser=false

# Watch for tailwind changes
tailwind-watch:
	npx tailwindcss -i ./assets/css/input.css -o ./assets/css/output.css --watch

# Development for Cloudflare Workers
dev-wrangler:
	wrangler dev --test-scheduled

# Deploy to Cloudflare Workers
deploy:
	wrangler deploy

# Generate static single-page application
static: templ tailwind
	go run ui/static-gen.go

# Serve static SPA locally
serve-static:
	cd dist && python3 -m http.server 3000

.PHONY: templ tailwind clean build server dev templ-watch tailwind-watch dev-wrangler deploy static serve-static