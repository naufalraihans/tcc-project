# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## What this repo is (read this first)

This is the **TCC ITPLN Web Platform** at the planning stage. The production app is **not built yet** — the repo currently holds specs, content, UI references, and source assets.

- **The plan (the actual work):** `PLANNING/*.md` — the intended production architecture and requirements. This is the focus. Start here.
- **UI reference only:** `designUI/{landingPage,dashboard,loginPage}/` — three separate **v0-generated Next.js 16 + React 19 + Tailwind v4 + shadcn/ui** apps. These are look-and-feel references, **not** the production stack (which will be SvelteKit + Go — see below). Use them for visual direction; don't build production code inside them. Their `node_modules/`/`.next/` are gitignored.
- **Content source of truth:** `PLANNING/profile.md` — all real facts (numbers, names, programs, contacts) come from here, sourced from the official PDF in `FILETAMBAH/`. When writing UI copy, pull data from `profile.md`, not from memory or dummy values.
- **Source assets:** `FILETAMBAH/` — official logos + the profile PDF/PPTX. Real brand assets live here.

Language of the docs and UI copy is **Indonesian**.

## Planned production architecture

**SvelteKit (TS) frontend + Go (Clean Architecture) backend + Supabase (Postgres/Auth/Storage) + Midtrans.** The `frontend/` and `backend/` folders described in the planning docs are a spec, not reality — nothing is scaffolded yet. Reference `PLANNING/` before implementing:

- `01_project_plan.md` — roles/access matrix, class model, roadmap, tech stack decisions.
- `02_erd_database.md` / `06_supabase_setup.md` — schema & migrations.
- `03_api_contract.md` — Go REST API contract.
- `04_project_structure.md` — target monorepo layout & naming conventions.
- `05_environment.md` — env vars for frontend/backend.
- `07_auth_flow.md` — **key decision:** Supabase Auth owns the session; the Go backend only *verifies* the JWT (`SUPABASE_JWT_SECRET`). The Supabase JWT `role` claim is always `"authenticated"` — the real app role (`user`/`admin`) lives in the `profiles` table, so the backend **must query the DB** to check role, per-request.
- `08_midtrans_payment.md` — paid-class flow: pending transaction → Midtrans Snap → webhook callback to Go backend confirms enrollment.
- `09_ui_reference.md` — maps the `designUI/` prototypes to intended screens.

## Design rules (from `PLANNING/01_project_plan.md` §6 — enforce these)

- **No emoji anywhere in the UI.** No decorative icons — icons only when they carry a navigation/action function.
- **No orange.** Palette is cool blue-teal only: Navy Teal `#0C4F6A`, Sky Blue `#1A8DB2`, Cool Slate `#2E4A5A`, Off White `#F4F7FA`, Charcoal `#1A1A2E`, Muted `#6B7280`, Border `#E2E8F0`.
- Minimalist + bento grid + formal. Generous whitespace, consistent rounded corners, chip/pill badges.
- Fonts: `Plus Jakarta Sans` (headings), `Inter` (body).
