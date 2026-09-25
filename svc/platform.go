// Package svc defines platform services, settings, and the registry.
package svc

import (
	"context"
	"fmt"
	"io"
	"sync"

	"github.com/bizshuk/video-downloader/pkg/download"
)

// Options holds common download parameters passed to platform services.
type Options struct {
	URL           string
	Type          download.MediaType
	Quality       int
	OutputDir     string
	CookieBrowser string
}

// Platform defines the capabilities and settings of a video source site.
type Platform interface {
	Name() string
	Label() string
	Match(rawURL string) bool
	NeedsCookies() bool
	DefaultCookieBrowser() string
	ExtraArgs(cookieBrowser string) []string
	Download(ctx context.Context, opts Options, stdout, stderr io.Writer) error
}

var (
	registryMu sync.RWMutex
	platforms  = make(map[string]Platform)
	order      []string
)

// Register registers a platform service. It panics on duplicate names.
func Register(p Platform) {
	registryMu.Lock()
	defer registryMu.Unlock()

	name := p.Name()
	if _, exists := platforms[name]; exists {
		panic(fmt.Sprintf("svc: platform %q already registered", name))
	}
	platforms[name] = p
	order = append(order, name)
}

// Get returns the registered platform by name.
func Get(name string) (Platform, bool) {
	registryMu.RLock()
	defer registryMu.RUnlock()
	p, ok := platforms[name]
	return p, ok
}

// All returns all registered platforms in registration order.
func All() []Platform {
	registryMu.RLock()
	defer registryMu.RUnlock()
	list := make([]Platform, 0, len(order))
	for _, name := range order {
		list = append(list, platforms[name])
	}
	return list
}

// Detect matches a URL to a registered platform.
// If no specific platform matches, it returns the generic platform or nil.
func Detect(rawURL string) Platform {
	registryMu.RLock()
	defer registryMu.RUnlock()

	var generic Platform
	for _, name := range order {
		p := platforms[name]
		if name == "generic" {
			generic = p
			continue
		}
		if p.Match(rawURL) {
			return p
		}
	}
	if generic != nil && generic.Match(rawURL) {
		return generic
	}
	return generic
}
