# Deployment Options

This project supports multiple deployment strategies:

## 🚀 Cloudflare Workers (Current Setup)
**Best for**: Dynamic server-side functionality with global edge deployment

```bash
# Development
make dev-wrangler          # → http://localhost:8787

# Deploy
make deploy               # → https://go-website.fernandocorreia316.workers.dev
```

**Features:**
- Server-side Go application compiled to WebAssembly
- Dynamic route handling
- PDF generation API endpoint
- Global edge distribution

---

## 📁 Static Single-Page Application (SPA)
**Best for**: Simple static hosting (Netlify, Vercel, GitHub Pages, etc.)

```bash
# Generate static SPA
make static               # Creates dist/index.html with everything inlined

# Test locally
make serve-static         # → http://localhost:3000
```

**Features:**
- Single `index.html` file with all CSS inlined
- Client-side routing with Alpine.js
- All pages embedded in one file
- Works on any static hosting platform

**Generated Files:**
```
dist/
├── index.html           # Complete SPA (includes all CSS & JS)
└── assets/css/
    └── output.css       # Separate CSS file (optional)
```

---

## 🖥️ Traditional Server (Development)
**Best for**: Local development and testing

```bash
# Development with hot reload
make dev                 # → http://localhost:8090

# Build and run server
make server             # Traditional Go server
```

---

## 📋 Quick Deploy to Popular Platforms

### Netlify / Vercel / GitHub Pages
1. Run `make static`
2. Upload the `dist` folder
3. Set `dist/index.html` as the entry point

### Cloudflare Pages
1. Run `make static` 
2. Upload `dist` folder to Cloudflare Pages
3. Or use the existing Cloudflare Workers setup

### AWS S3 / Static Hosting
1. Run `make static`
2. Upload `dist/index.html` to your bucket
3. Configure as static website

---

## 🎯 Recommended Usage

- **Production**: Cloudflare Workers (`make deploy`)
- **Quick demos**: Static SPA (`make static`)
- **Development**: Local server (`make dev`)

Each approach maintains the same functionality while optimizing for different hosting environments.