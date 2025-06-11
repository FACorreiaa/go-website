# Fernando Correia - Personal Website

Modern personal portfolio website built with Go and templ, deployable as both dynamic Cloudflare Workers and static SPA.

## 🛠️ Tech Stack

- **Go** - Backend logic and templating
- **templ** - Type-safe HTML templating
- **Tailwind CSS** - Utility-first styling
- **Alpine.js** - Lightweight JavaScript framework
- **Cloudflare Workers** - Edge deployment

## 🚀 Quick Start

### Cloudflare Workers (Recommended)
```bash
# Development
make dev-wrangler    # → http://localhost:8787

# Deploy
make deploy         # → Live on Cloudflare
```

### Static SPA (Universal)
```bash
# Generate single-file website
make static         # Creates dist/index.html

# Test locally  
make serve-static   # → http://localhost:3000
```

### Local Development
```bash
# Traditional server with hot reload
make dev           # → http://localhost:8090
```

## 📁 Project Structure

```
├── ui/                 # templ templates
│   ├── components/     # Reusable components
│   ├── layouts/        # Page layouts  
│   ├── modules/        # Feature modules
│   └── pages/          # Page templates
├── assets/css/         # Tailwind styles
├── static-gen.go       # Static SPA generator
└── main.go            # Main application
```

## 📋 Available Commands

- `make build` - Build for Cloudflare Workers
- `make static` - Generate static SPA
- `make deploy` - Deploy to Cloudflare
- `make dev` - Local development server
- `make dev-wrangler` - Cloudflare Workers dev
- `make serve-static` - Serve static files

See [DEPLOYMENT.md](./DEPLOYMENT.md) for detailed deployment options.

## ✨ Features

- 📱 Responsive design with dark/light theme
- ⚡ Fast loading with optimized assets
- 🌍 Global edge deployment via Cloudflare
- 📄 Dynamic PDF CV generation
- 🔧 Type-safe templates with templ
- 🎨 Modern UI with Tailwind CSS
