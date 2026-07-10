# Project Rules & Development Reference

## Application Description
This is a school management application ("warga-sekolah").
- **Backend**: Go (do not modify)
- **Frontend**: SvelteKit, embedded in the Go app (located in the `frontend` folder)

## Strict Rules
- **DO NOT** edit, modify, or adjust any files outside of the `frontend/` directory (such as Go backend code, databases, internal configurations, etc.) unless you are directly and explicitly instructed by the user to modify the backend code.

## Development Commands
All commands must be executed within the `frontend/` directory:
- Run development server: `bun run dev`
- Build frontend for Go embed: `bun run build`
- Run type check: `bun run check`
- Format frontend code: `bun run format`
- Lint frontend code: `bun run lint`
