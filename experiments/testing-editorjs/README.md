# Content Composer — A Dev-First Content Engine for Go Backends

> **Composable content, block-based editing, and static-site speed — all in a drop-in Go module.**

---

## The Problem

Modern web development is stuck between two extremes:

- **Visual builders** (Wix, Framer, Webflow) offer sleek editing experiences — but limited backend flexibility and performance tuning.
- **Code-first tools** (WordPress, headless CMSs, static site generators) give developers control — but are bloated, complex, or too opinionated.

Developers using **Go** for backend development are especially underserved. There’s no streamlined, modern, performance-oriented content management system designed specifically for the Go ecosystem.

---

## The Solution

**Content Composer** is a lightweight, Go-native content rendering toolkit that gives developers:

- **Drop-in integration**: Add a route, configure a few settings, and instantly start composing pages.
- **Structured content editing**: Visual, block-based editing powered by [Editor.js](https://editorjs.io/).
- **Flexible data sources**: Pull content from APIs, databases, or flat-file collections (Framer-style).
- **Composable templates**: Use idiomatic Go templating and a layout engine inspired by `golang.org/x/website`.
- **Developer-first UX**: Designed for engineers who want fast, scalable, versionable content systems — with none of the WordPress bloat.

---

## Who Is It For?

- **Go developers** building landing pages, blogs, dashboards, or micro-sites.
- **Startup teams** that want content agility without spinning up a separate CMS.
- **Agencies & freelancers** tired of bloated systems and looking for high-performance, maintainable sites.

---

## What It Can Do

- Render pages from structured content blocks
- Integrate with REST/GraphQL APIs or direct DB queries
- Handle layout and content composition dynamically
- Allow non-devs to edit content through a clean, JSON-based visual editor
- Deliver performance equal to or better than static site generators — with dynamic flexibility

---

## Roadmap (WIP)

- Editor.js integration for visual content blocks
- Layout system based on Go template inheritance
- Config-based route and page registration
- Admin interface for managing content collections
- Plugin system for content types, SEO tags, and caching strategies

---

## Inspiration

- `golang.org/x/website` static rendering architecture
- Framer’s concept of **collections** and visual modularity
- Hugo’s performance, without the rigid SSG pipeline
- WordPress's extensibility — minus the PHP bloat

---

## Status & Call to Action

We’re in early development and currently onboarding **pilot users**.

If you:

- Build with Go,
- Need flexible content management,
- Care about performance and dev UX,

**We’d love to collaborate.**

> [Reach out or follow progress here](https://github.com/Fako-Drop/Content-Composer)

---
