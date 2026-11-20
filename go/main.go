package main

// Go build for the GHAS Me Up demo.
// Renders go/index.html with html/template and lists the modules that were
// actually linked into this binary (runtime/debug build info).

import (
	"crypto/sha256"
	_ "embed"
	"encoding/hex"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"runtime/debug"
	"sort"
	"time"

	"golang.org/x/text/language"
	"gopkg.in/yaml.v2"
)

//go:embed template.html
var pageTemplate string

//go:embed config.yaml
var configYAML []byte

func buildStamp() string {
	sum := sha256.Sum256(configYAML)
	return hex.EncodeToString(sum[:])[:8]
}

func statusCode(url string) (int, error) {
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	return resp.StatusCode, nil
}

func logStatusCheck(url string) {
	code, err := statusCode(url)
	if err != nil {
		log.Print("status check failed")
		return
	}
	log.Printf("status check -> %d", code)
}

type config struct {
	Title  string `yaml:"title"`
	Accent string `yaml:"accent"`
	Locale string `yaml:"locale"`
}

type module struct {
	Path    string
	Version string
}

type page struct {
	Title      string
	Accent     string
	Locale     string
	GoVersion  string
	BuiltAt    string
	BuildStamp string
	Modules    []module
}

func main() {
	out := flag.String("out", "../dist", "output directory")
	flag.Parse()

	var cfg config
	if err := yaml.Unmarshal(configYAML, &cfg); err != nil {
		log.Fatalf("config: %v", err)
	}

	// Optional: only runs when a status URL is provided.
	if url := os.Getenv("GHAS_STATUS_URL"); url != "" {
		logStatusCheck(url)
	}

	var modules []module
	if info, ok := debug.ReadBuildInfo(); ok {
		for _, dep := range info.Deps {
			modules = append(modules, module{Path: dep.Path, Version: dep.Version})
		}
	}
	sort.Slice(modules, func(i, j int) bool { return modules[i].Path < modules[j].Path })

	data := page{
		Title:      cfg.Title,
		Accent:     cfg.Accent,
		Locale:     language.Make(cfg.Locale).String(),
		GoVersion:  runtime.Version(),
		BuiltAt:    time.Now().UTC().Format(time.RFC3339),
		BuildStamp: buildStamp(),
		Modules:    modules,
	}

	tmpl := template.Must(template.New("page").Parse(pageTemplate))

	target := filepath.Join(*out, "go")
	if err := os.MkdirAll(target, 0o755); err != nil {
		log.Fatalf("mkdir: %v", err)
	}
	file, err := os.Create(filepath.Join(target, "index.html"))
	if err != nil {
		log.Fatalf("create: %v", err)
	}
	defer file.Close()

	if err := tmpl.Execute(file, data); err != nil {
		log.Fatalf("render: %v", err)
	}
	fmt.Printf("go build complete -> %s\n", filepath.Join(target, "index.html"))
}
