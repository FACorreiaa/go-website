package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/FACorreiaa/go-website/ui/pages"
	"github.com/a-h/templ"
)

func main() {
	// Create dist directory
	distPath := "dist"
	os.RemoveAll(distPath) // Clean existing dist
	os.MkdirAll(distPath, 0755)

	// Generate single page application
	generateSPA(distPath)

	// Copy CSS
	copyCSS(distPath)

	fmt.Println("✅ Static site generated in ./dist/")
	fmt.Println("📁 Files created:")
	fmt.Println("   - index.html (complete SPA)")
	fmt.Println("   - assets/css/output.css")
}

func generateSPA(distPath string) {
	// Create HTML file
	file, err := os.Create(filepath.Join(distPath, "index.html"))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Generate complete SPA HTML
	html := generateCompleteHTML()

	_, err = file.WriteString(html)
	if err != nil {
		panic(err)
	}
}

func generateCompleteHTML() string {
	// Get CSS content
	cssContent := readCSS()

	// Render each page to string
	landingHTML := renderComponentToString(pages.Landing())
	aboutHTML := renderComponentToString(pages.About())
	projectsHTML := renderComponentToString(pages.Projects())

	// Create complete SPA HTML
	html := fmt.Sprintf(`<!DOCTYPE html>
<html lang="en" class="h-full dark">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Fernando Correia - Portfolio</title>

    <!-- Inline CSS -->
    <style>
%s
    </style>

    <!-- Alpine.js -->
    <script defer src="https://cdn.jsdelivr.net/npm/alpinejs@3.x.x/dist/cdn.min.js"></script>

    <!-- SPA Router Script -->
    <script>
        // Theme handler
        document.documentElement.classList.toggle('dark', localStorage.getItem('appTheme') === 'dark');

        document.addEventListener('alpine:init', () => {
            Alpine.data('themeHandler', () => ({
                isDark: localStorage.getItem('appTheme') === 'dark',
                themeClasses() {
                    return this.isDark ? 'text-white' : 'bg-white text-black'
                },
                toggleTheme() {
                    this.isDark = !this.isDark;
                    localStorage.setItem('appTheme', this.isDark ? 'dark' : 'light');
                    document.documentElement.classList.toggle('dark', this.isDark);
                }
            }));

            Alpine.data('themeSwitcherHandler', () => ({
                isDarkMode() {
                    return this.isDark
                },
                isLightMode() {
                    return !this.isDark
                }
            }));

            // SPA Router
            Alpine.data('router', () => ({
                currentPage: 'home',

                init() {
                    // Set initial page based on URL
                    this.handleRoute();

                    // Listen for browser back/forward
                    window.addEventListener('popstate', () => {
                        this.handleRoute();
                    });
                },

                handleRoute() {
                    const path = window.location.pathname;
                    if (path === '/about') {
                        this.currentPage = 'about';
                    } else if (path === '/projects') {
                        this.currentPage = 'projects';
                    } else {
                        this.currentPage = 'home';
                    }
                },

                navigate(page) {
                    this.currentPage = page;
                    const path = page === 'home' ? '/' : '/' + page;
                    window.history.pushState({}, '', path);
                }
            }));
        });
    </script>
</head>
<body x-data="{ ...themeHandler(), ...router() }" x-bind:class="themeClasses">
    <!-- Home Page -->
    <div x-show="currentPage === 'home'" x-transition>
%s
    </div>

    <!-- About Page -->
    <div x-show="currentPage === 'about'" x-transition>
%s
    </div>

    <!-- Projects Page -->
    <div x-show="currentPage === 'projects'" x-transition>
%s
    </div>
</body>
</html>`, cssContent, landingHTML, aboutHTML, projectsHTML)

	// Fix navigation links to use Alpine.js router
	html = strings.ReplaceAll(html, `href="/"`, `href="#" @click.prevent="navigate('home')"`)
	html = strings.ReplaceAll(html, `href="/about"`, `href="#" @click.prevent="navigate('about')"`)
	html = strings.ReplaceAll(html, `href="/projects"`, `href="#" @click.prevent="navigate('projects')"`)

	return html
}

func renderComponentToString(component templ.Component) string {
	var builder strings.Builder
	err := component.Render(context.Background(), &builder)
	if err != nil {
		panic(err)
	}
	return builder.String()
}

func readCSS() string {
	cssPath := "assets/css/output.css"
	content, err := os.ReadFile(cssPath)
	if err != nil {
		fmt.Printf("Warning: Could not read CSS file: %v\n", err)
		return ""
	}
	return string(content)
}

func copyCSS(distPath string) {
	// Also copy CSS file for external hosting compatibility
	assetsDir := filepath.Join(distPath, "assets", "css")
	os.MkdirAll(assetsDir, 0755)

	srcFile, err := os.Open("assets/css/output.css")
	if err != nil {
		fmt.Printf("Warning: Could not copy CSS file: %v\n", err)
		return
	}
	defer srcFile.Close()

	dstFile, err := os.Create(filepath.Join(assetsDir, "output.css"))
	if err != nil {
		fmt.Printf("Warning: Could not create CSS file: %v\n", err)
		return
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		fmt.Printf("Warning: Could not copy CSS content: %v\n", err)
	}
}
