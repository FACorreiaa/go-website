# Generate templ files
templ:
	templ generate

# Generate CSS
tailwind:
	tailwindcss -i ./assets/css/input.css -o ./assets/css/output.css --minify

# Clean previous builds
clean:
	rm -rf dist/

# Build static site
build: clean templ tailwind
	go run main.go

# Serve static files locally for testing
serve:
	cd dist && python3 -m http.server 8080

# Development mode with file watching
dev:
	make tailwind
	templ generate --watch --proxy="http://localhost:8080" --open-browser=false &
	tailwindcss -i ./assets/css/input.css -o ./assets/css/output.css --watch &
	air --build.cmd "go run main.go" --build.bin "" --build.delay "100" --build.include_ext "go,templ" --build.stop_on_error "false" --misc.clean_on_exit true

.PHONY: templ tailwind clean build serve dev