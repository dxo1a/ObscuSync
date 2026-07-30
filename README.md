# ObscuSync

ObscuSync syncs game mods, configs and other files between computers.

Supported games live in the [registry](internal/config/registry.go).  
Maybe later I'll let people define their own configs of supported games, but not yet.

How it works is pretty simple:  
Host runs `update-manifest [profile]` - this builds a manifest with file list, sizes and SHA256 hashes.  
Client runs `sync [profile]`, grabs that manifest, compares it with local files and applies the changes.

**Important:** profile name must be the same on both host and client.

> [!WARNING]  
> When syncing, the client pulls files straight from the host machine.  
> Only use this with people you actually trust. There's no real security here.  
> I'm not responsible if your PC gets infected or hacked. The tool just moves files around.

## Table of Contents
- [ObscuSync](#obscusync)
  - [Table of Contents](#table-of-contents)
  - [Commands](#commands)
  - [Statistics](#statistics)

## Commands
- **serve** — start the HTTP server
- **sync** — sync local files with the remote manifest (server address is in `config.yaml`)
- **update-manifest [profile]** — rebuild the manifest for a profile
- **help** — show help

## Statistics
![GreenLuma downloads](https://img.shields.io/github/downloads/dxo1a/ObscuSync/total?style=flat-square) ![GreenLuma repo stars](https://img.shields.io/github/stars/dxo1a/ObscuSync?style=flat-square)