// Code generated by fileb0x at "2023-09-03 15:33:59.48079419 +0200 CEST m=+0.035439077" from config file "b0x.yaml" DO NOT EDIT.
// modification hash(30798deaa0bd9249c3a2fd00904f93f9.6e40df19812da11a998c932862379ab6)

package static

import (
	"bytes"

	"context"
	"io"
	"net/http"
	"os"
	"path"

	"golang.org/x/net/webdav"
)

var (
	// CTX is a context for webdav vfs
	CTX = context.Background()

	// FS is a virtual memory file system
	FS = webdav.NewMemFS()

	// Handler is used to server files through a http handler
	Handler *webdav.Handler

	// HTTP is the http file system
	HTTP http.FileSystem = new(HTTPFS)
)

// HTTPFS implements http.FileSystem
type HTTPFS struct {
	// Prefix allows to limit the path of all requests. F.e. a prefix "css" would allow only calls to /css/*
	Prefix string
}

// FileWebIndexHTML is "web/index.html"
var FileWebIndexHTML = []byte("\x3c\x21\x44\x4f\x43\x54\x59\x50\x45\x20\x68\x74\x6d\x6c\x3e\x0a\x3c\x68\x74\x6d\x6c\x3e\x0a\x3c\x68\x65\x61\x64\x3e\x0a\x20\x20\x20\x20\x3c\x6d\x65\x74\x61\x20\x63\x68\x61\x72\x73\x65\x74\x3d\x22\x55\x54\x46\x2d\x38\x22\x2f\x3e\x0a\x20\x20\x20\x20\x3c\x6d\x65\x74\x61\x20\x6e\x61\x6d\x65\x3d\x22\x76\x69\x65\x77\x70\x6f\x72\x74\x22\x20\x63\x6f\x6e\x74\x65\x6e\x74\x3d\x22\x77\x69\x64\x74\x68\x3d\x64\x65\x76\x69\x63\x65\x2d\x77\x69\x64\x74\x68\x2c\x20\x69\x6e\x69\x74\x69\x61\x6c\x2d\x73\x63\x61\x6c\x65\x3d\x31\x2e\x30\x22\x2f\x3e\x0a\x20\x20\x20\x20\x3c\x74\x69\x74\x6c\x65\x3e\x57\x68\x6f\x20\x69\x73\x20\x68\x6f\x6d\x65\x3f\x3c\x2f\x74\x69\x74\x6c\x65\x3e\x0a\x20\x20\x20\x20\x3c\x73\x74\x79\x6c\x65\x3e\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x72\x6f\x6f\x74\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x2d\x2d\x68\x65\x61\x64\x65\x72\x2d\x62\x61\x63\x6b\x67\x72\x6f\x75\x6e\x64\x3a\x20\x23\x33\x33\x33\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x2d\x2d\x68\x65\x61\x64\x65\x72\x2d\x63\x6f\x6c\x6f\x72\x3a\x20\x77\x68\x69\x74\x65\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x2d\x2d\x61\x76\x61\x74\x61\x72\x2d\x62\x61\x63\x6b\x67\x72\x6f\x75\x6e\x64\x3a\x20\x23\x64\x31\x64\x35\x64\x62\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x2d\x2d\x61\x76\x61\x74\x61\x72\x2d\x63\x6f\x6c\x6f\x72\x3a\x20\x77\x68\x69\x74\x65\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x2d\x2d\x64\x61\x72\x6b\x2d\x62\x61\x63\x6b\x67\x72\x6f\x75\x6e\x64\x3a\x20\x23\x32\x38\x32\x37\x32\x37\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x2d\x2d\x64\x61\x72\x6b\x2d\x63\x6f\x6c\x6f\x72\x3a\x20\x77\x68\x69\x74\x65\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x7d\x0a\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x2a\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x6d\x61\x72\x67\x69\x6e\x3a\x20\x30\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x70\x61\x64\x64\x69\x6e\x67\x3a\x20\x30\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x7d\x0a\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x62\x6f\x64\x79\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x66\x6f\x6e\x74\x2d\x66\x61\x6d\x69\x6c\x79\x3a\x20\x2d\x61\x70\x70\x6c\x65\x2d\x73\x79\x73\x74\x65\x6d\x2c\x20\x42\x6c\x69\x6e\x6b\x4d\x61\x63\x53\x79\x73\x74\x65\x6d\x46\x6f\x6e\x74\x2c\x20\x22\x53\x65\x67\x6f\x65\x20\x55\x49\x22\x2c\x20\x52\x6f\x62\x6f\x74\x6f\x2c\x20\x48\x65\x6c\x76\x65\x74\x69\x63\x61\x2c\x20\x41\x72\x69\x61\x6c\x2c\x20\x73\x61\x6e\x73\x2d\x73\x65\x72\x69\x66\x2c\x20\x22\x41\x70\x70\x6c\x65\x20\x43\x6f\x6c\x6f\x72\x20\x45\x6d\x6f\x6a\x69\x22\x2c\x20\x22\x53\x65\x67\x6f\x65\x20\x55\x49\x20\x45\x6d\x6f\x6a\x69\x22\x2c\x20\x22\x53\x65\x67\x6f\x65\x20\x55\x49\x20\x53\x79\x6d\x62\x6f\x6c\x22\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x7d\x0a\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x40\x6d\x65\x64\x69\x61\x20\x28\x70\x72\x65\x66\x65\x72\x73\x2d\x63\x6f\x6c\x6f\x72\x2d\x73\x63\x68\x65\x6d\x65\x3a\x20\x64\x61\x72\x6b\x29\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x68\x74\x6d\x6c\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x62\x61\x63\x6b\x67\x72\x6f\x75\x6e\x64\x3a\x20\x76\x61\x72\x28\x2d\x2d\x64\x61\x72\x6b\x2d\x62\x61\x63\x6b\x67\x72\x6f\x75\x6e\x64\x29\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x63\x6f\x6c\x6f\x72\x3a\x20\x76\x61\x72\x28\x2d\x2d\x64\x61\x72\x6b\x2d\x63\x6f\x6c\x6f\x72\x29\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x7d\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x7d\x0a\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x68\x65\x61\x64\x65\x72\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x62\x61\x63\x6b\x67\x72\x6f\x75\x6e\x64\x3a\x20\x76\x61\x72\x28\x2d\x2d\x68\x65\x61\x64\x65\x72\x2d\x62\x61\x63\x6b\x67\x72\x6f\x75\x6e\x64\x29\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x63\x6f\x6c\x6f\x72\x3a\x20\x76\x61\x72\x28\x2d\x2d\x68\x65\x61\x64\x65\x72\x2d\x63\x6f\x6c\x6f\x72\x29\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x70\x61\x64\x64\x69\x6e\x67\x3a\x20\x32\x72\x65\x6d\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x74\x65\x78\x74\x2d\x61\x6c\x69\x67\x6e\x3a\x20\x63\x65\x6e\x74\x65\x72\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x7d\x0a\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x2e\x61\x76\x61\x74\x61\x72\x2d\x77\x72\x61\x70\x70\x65\x72\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x64\x69\x73\x70\x6c\x61\x79\x3a\x20\x66\x6c\x65\x78\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x6a\x75\x73\x74\x69\x66\x79\x2d\x63\x6f\x6e\x74\x65\x6e\x74\x3a\x20\x63\x65\x6e\x74\x65\x72\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x7d\x0a\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x2e\x61\x76\x61\x74\x61\x72\x2d\x6f\x66\x66\x6c\x69\x6e\x65\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x62\x61\x63\x6b\x67\x72\x6f\x75\x6e\x64\x2d\x63\x6f\x6c\x6f\x72\x3a\x20\x72\x67\x62\x28\x33\x33\x20\x33\x33\x20\x33\x33\x29\x20\x21\x69\x6d\x70\x6f\x72\x74\x61\x6e\x74\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x7d\x0a\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x2e\x61\x76\x61\x74\x61\x72\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x61\x6c\x69\x67\x6e\x2d\x69\x74\x65\x6d\x73\x3a\x20\x63\x65\x6e\x74\x65\x72\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x64\x69\x73\x70\x6c\x61\x79\x3a\x20\x66\x6c\x65\x78\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x6a\x75\x73\x74\x69\x66\x79\x2d\x63\x6f\x6e\x74\x65\x6e\x74\x3a\x20\x63\x65\x6e\x74\x65\x72\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x62\x61\x63\x6b\x67\x72\x6f\x75\x6e\x64\x2d\x63\x6f\x6c\x6f\x72\x3a\x20\x76\x61\x72\x28\x2d\x2d\x61\x76\x61\x74\x61\x72\x2d\x62\x61\x63\x6b\x67\x72\x6f\x75\x6e\x64\x29\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x63\x6f\x6c\x6f\x72\x3a\x20\x76\x61\x72\x28\x2d\x2d\x61\x76\x61\x74\x61\x72\x2d\x63\x6f\x6c\x6f\x72\x29\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x62\x6f\x72\x64\x65\x72\x2d\x72\x61\x64\x69\x75\x73\x3a\x20\x35\x30\x25\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x68\x65\x69\x67\x68\x74\x3a\x20\x37\x72\x65\x6d\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x77\x69\x64\x74\x68\x3a\x20\x37\x72\x65\x6d\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x66\x6f\x6e\x74\x2d\x73\x69\x7a\x65\x3a\x20\x33\x72\x65\x6d\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x6d\x61\x72\x67\x69\x6e\x3a\x20\x31\x72\x65\x6d\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x75\x73\x65\x72\x2d\x73\x65\x6c\x65\x63\x74\x3a\x20\x6e\x6f\x6e\x65\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x7d\x0a\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x2e\x75\x73\x65\x72\x73\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x64\x69\x73\x70\x6c\x61\x79\x3a\x20\x66\x6c\x65\x78\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x66\x6c\x65\x78\x2d\x77\x72\x61\x70\x3a\x20\x77\x72\x61\x70\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x66\x6c\x65\x78\x2d\x64\x69\x72\x65\x63\x74\x69\x6f\x6e\x3a\x20\x72\x6f\x77\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x66\x6c\x65\x78\x2d\x62\x61\x73\x69\x73\x3a\x20\x33\x33\x2e\x33\x33\x33\x25\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x6a\x75\x73\x74\x69\x66\x79\x2d\x63\x6f\x6e\x74\x65\x6e\x74\x3a\x20\x63\x65\x6e\x74\x65\x72\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x7d\x0a\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x2e\x75\x73\x65\x72\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x6d\x61\x72\x67\x69\x6e\x3a\x20\x32\x72\x65\x6d\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x74\x65\x78\x74\x2d\x61\x6c\x69\x67\x6e\x3a\x20\x63\x65\x6e\x74\x65\x72\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x64\x69\x73\x70\x6c\x61\x79\x3a\x20\x66\x6c\x65\x78\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x66\x6c\x65\x78\x2d\x64\x69\x72\x65\x63\x74\x69\x6f\x6e\x3a\x20\x63\x6f\x6c\x75\x6d\x6e\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x61\x6c\x69\x67\x6e\x2d\x63\x6f\x6e\x74\x65\x6e\x74\x3a\x20\x63\x65\x6e\x74\x65\x72\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x6a\x75\x73\x74\x69\x66\x79\x2d\x63\x6f\x6e\x74\x65\x6e\x74\x3a\x20\x63\x65\x6e\x74\x65\x72\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x7d\x0a\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x2e\x6f\x6e\x6c\x69\x6e\x65\x2d\x74\x65\x78\x74\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x63\x6f\x6c\x6f\x72\x3a\x20\x23\x32\x32\x37\x39\x32\x32\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x7d\x0a\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x2e\x6f\x66\x66\x6c\x69\x6e\x65\x2d\x74\x65\x78\x74\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x63\x6f\x6c\x6f\x72\x3a\x20\x23\x63\x66\x35\x34\x35\x34\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x7d\x0a\x20\x20\x20\x20\x3c\x2f\x73\x74\x79\x6c\x65\x3e\x0a\x20\x20\x20\x20\x3c\x73\x63\x72\x69\x70\x74\x3e\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x63\x6f\x6e\x73\x74\x20\x67\x65\x6e\x65\x72\x61\x74\x65\x52\x61\x6e\x64\x6f\x6d\x43\x6f\x6c\x6f\x72\x20\x3d\x20\x28\x29\x20\x3d\x3e\x20\x22\x23\x22\x20\x2b\x20\x4d\x61\x74\x68\x2e\x66\x6c\x6f\x6f\x72\x28\x4d\x61\x74\x68\x2e\x72\x61\x6e\x64\x6f\x6d\x28\x29\x20\x2a\x20\x31\x36\x37\x37\x37\x32\x31\x35\x29\x2e\x74\x6f\x53\x74\x72\x69\x6e\x67\x28\x31\x36\x29\x3b\x0a\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x64\x6f\x63\x75\x6d\x65\x6e\x74\x2e\x61\x64\x64\x45\x76\x65\x6e\x74\x4c\x69\x73\x74\x65\x6e\x65\x72\x28\x22\x44\x4f\x4d\x43\x6f\x6e\x74\x65\x6e\x74\x4c\x6f\x61\x64\x65\x64\x22\x2c\x20\x28\x29\x20\x3d\x3e\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x64\x6f\x63\x75\x6d\x65\x6e\x74\x2e\x71\x75\x65\x72\x79\x53\x65\x6c\x65\x63\x74\x6f\x72\x41\x6c\x6c\x28\x22\x2e\x61\x76\x61\x74\x61\x72\x22\x29\x2e\x66\x6f\x72\x45\x61\x63\x68\x28\x61\x76\x61\x74\x61\x72\x20\x3d\x3e\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x61\x76\x61\x74\x61\x72\x2e\x73\x74\x79\x6c\x65\x2e\x62\x61\x63\x6b\x67\x72\x6f\x75\x6e\x64\x43\x6f\x6c\x6f\x72\x20\x3d\x20\x67\x65\x6e\x65\x72\x61\x74\x65\x52\x61\x6e\x64\x6f\x6d\x43\x6f\x6c\x6f\x72\x28\x29\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x7d\x29\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x7d\x29\x3b\x0a\x20\x20\x20\x20\x3c\x2f\x73\x63\x72\x69\x70\x74\x3e\x0a\x3c\x2f\x68\x65\x61\x64\x3e\x0a\x3c\x62\x6f\x64\x79\x3e\x0a\x3c\x68\x65\x61\x64\x65\x72\x3e\x0a\x20\x20\x20\x20\x3c\x68\x31\x3e\x57\x68\x6f\x20\x69\x73\x20\x68\x6f\x6d\x65\x3f\x3c\x2f\x68\x31\x3e\x0a\x3c\x2f\x68\x65\x61\x64\x65\x72\x3e\x0a\x3c\x64\x69\x76\x20\x63\x6c\x61\x73\x73\x3d\x22\x75\x73\x65\x72\x73\x22\x3e\x0a\x20\x20\x20\x20\x7b\x7b\x20\x72\x61\x6e\x67\x65\x20\x24\x75\x73\x65\x72\x2c\x20\x24\x64\x65\x76\x69\x63\x65\x73\x20\x3a\x3d\x20\x2e\x6d\x61\x70\x70\x69\x6e\x67\x20\x7d\x7d\x0a\x20\x20\x20\x20\x3c\x64\x69\x76\x20\x63\x6c\x61\x73\x73\x3d\x22\x75\x73\x65\x72\x22\x3e\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x7b\x7b\x20\x24\x64\x65\x76\x69\x63\x65\x73\x5f\x6c\x65\x6e\x67\x74\x68\x20\x3a\x3d\x20\x6c\x65\x6e\x20\x24\x64\x65\x76\x69\x63\x65\x73\x20\x7d\x7d\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x3c\x64\x69\x76\x20\x63\x6c\x61\x73\x73\x3d\x22\x61\x76\x61\x74\x61\x72\x2d\x77\x72\x61\x70\x70\x65\x72\x22\x3e\x3c\x73\x70\x61\x6e\x20\x63\x6c\x61\x73\x73\x3d\x22\x61\x76\x61\x74\x61\x72\x20\x7b\x7b\x20\x69\x66\x20\x65\x71\x20\x24\x64\x65\x76\x69\x63\x65\x73\x5f\x6c\x65\x6e\x67\x74\x68\x20\x30\x20\x7d\x7d\x61\x76\x61\x74\x61\x72\x2d\x6f\x66\x66\x6c\x69\x6e\x65\x7b\x7b\x20\x65\x6e\x64\x20\x7d\x7d\x22\x3e\x7b\x7b\x20\x73\x6c\x69\x63\x65\x20\x24\x75\x73\x65\x72\x20\x30\x20\x31\x20\x7d\x7d\x3c\x2f\x73\x70\x61\x6e\x3e\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x3c\x2f\x64\x69\x76\x3e\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x3c\x68\x32\x3e\x7b\x7b\x20\x24\x75\x73\x65\x72\x20\x7d\x7d\x3c\x2f\x68\x32\x3e\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x7b\x7b\x20\x69\x66\x20\x67\x74\x20\x24\x64\x65\x76\x69\x63\x65\x73\x5f\x6c\x65\x6e\x67\x74\x68\x20\x30\x20\x7d\x7d\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x3c\x70\x20\x63\x6c\x61\x73\x73\x3d\x22\x6f\x6e\x6c\x69\x6e\x65\x2d\x74\x65\x78\x74\x22\x3e\x6f\x6e\x6c\x69\x6e\x65\x20\x77\x69\x74\x68\x20\x7b\x7b\x20\x6c\x65\x6e\x20\x24\x64\x65\x76\x69\x63\x65\x73\x20\x7d\x7d\x20\x64\x65\x76\x69\x63\x65\x7b\x7b\x20\x69\x66\x20\x67\x74\x20\x24\x64\x65\x76\x69\x63\x65\x73\x5f\x6c\x65\x6e\x67\x74\x68\x20\x31\x20\x7d\x7d\x73\x7b\x7b\x20\x65\x6e\x64\x20\x7d\x7d\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x7b\x7b\x20\x65\x6c\x73\x65\x20\x7d\x7d\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x3c\x70\x20\x63\x6c\x61\x73\x73\x3d\x22\x6f\x66\x66\x6c\x69\x6e\x65\x2d\x74\x65\x78\x74\x22\x3e\x6f\x66\x66\x6c\x69\x6e\x65\x3c\x2f\x70\x3e\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x7b\x7b\x20\x65\x6e\x64\x20\x7d\x7d\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x3c\x2f\x70\x3e\x0a\x20\x20\x20\x20\x3c\x2f\x64\x69\x76\x3e\x0a\x20\x20\x20\x20\x7b\x7b\x20\x65\x6e\x64\x20\x7d\x7d\x0a\x3c\x2f\x64\x69\x76\x3e\x0a\x3c\x2f\x62\x6f\x64\x79\x3e\x0a\x3c\x2f\x68\x74\x6d\x6c\x3e")

func init() {
	err := CTX.Err()
	if err != nil {
		panic(err)
	}

	err = FS.Mkdir(CTX, "web/", 0777)
	if err != nil && err != os.ErrExist {
		panic(err)
	}

	var f webdav.File

	f, err = FS.OpenFile(CTX, "web/index.html", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	_, err = f.Write(FileWebIndexHTML)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}

	Handler = &webdav.Handler{
		FileSystem: FS,
		LockSystem: webdav.NewMemLS(),
	}

}

// Open a file
func (hfs *HTTPFS) Open(path string) (http.File, error) {
	path = hfs.Prefix + path

	f, err := FS.OpenFile(CTX, path, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}

	return f, nil
}

// ReadFile is adapTed from ioutil
func ReadFile(path string) ([]byte, error) {
	f, err := FS.OpenFile(CTX, path, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(make([]byte, 0, bytes.MinRead))

	// If the buffer overflows, we will get bytes.ErrTooLarge.
	// Return that as an error. Any other panic remains.
	defer func() {
		e := recover()
		if e == nil {
			return
		}
		if panicErr, ok := e.(error); ok && panicErr == bytes.ErrTooLarge {
			err = panicErr
		} else {
			panic(e)
		}
	}()
	_, err = buf.ReadFrom(f)
	return buf.Bytes(), err
}

// WriteFile is adapTed from ioutil
func WriteFile(filename string, data []byte, perm os.FileMode) error {
	f, err := FS.OpenFile(CTX, filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
	if err != nil {
		return err
	}
	n, err := f.Write(data)
	if err == nil && n < len(data) {
		err = io.ErrShortWrite
	}
	if err1 := f.Close(); err == nil {
		err = err1
	}
	return err
}

// WalkDirs looks for files in the given dir and returns a list of files in it
// usage for all files in the b0x: WalkDirs("", false)
func WalkDirs(name string, includeDirsInList bool, files ...string) ([]string, error) {
	f, err := FS.OpenFile(CTX, name, os.O_RDONLY, 0)
	if err != nil {
		return nil, err
	}

	fileInfos, err := f.Readdir(0)
	if err != nil {
		return nil, err
	}

	err = f.Close()
	if err != nil {
		return nil, err
	}

	for _, info := range fileInfos {
		filename := path.Join(name, info.Name())

		if includeDirsInList || !info.IsDir() {
			files = append(files, filename)
		}

		if info.IsDir() {
			files, err = WalkDirs(filename, includeDirsInList, files...)
			if err != nil {
				return nil, err
			}
		}
	}

	return files, nil
}
