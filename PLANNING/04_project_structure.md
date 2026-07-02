# 🏗️ Project Structure — TCC ITPLN Web Platform

> **Dokumen ini mendefinisikan struktur folder untuk Frontend (SvelteKit) dan Backend (Go).**
> Versi: v1.0 | Status: Draft

---

## 1. Overview Monorepo

```
tcc-itpln/
├── frontend/          # SvelteKit App
├── backend/           # Go REST API
├── PLANNING/          # Dokumen perencanaan
│   ├── plan.md
│   ├── ERD.md
│   ├── API.md
│   ├── STRUCTURE.md
│   ├── ENV.md
│   └── SUPABASE.md
├── .gitignore
└── README.md
```

---

## 2. Frontend — SvelteKit

### Struktur Folder Lengkap

```
frontend/
├── src/
│   ├── app.html                    # HTML shell utama
│   ├── app.css                     # Global CSS & design tokens
│   │
│   ├── lib/
│   │   ├── components/             # Komponen UI reusable
│   │   │   ├── common/             # Komponen umum
│   │   │   │   ├── Button.svelte
│   │   │   │   ├── Badge.svelte
│   │   │   │   ├── Card.svelte
│   │   │   │   ├── Modal.svelte
│   │   │   │   ├── Input.svelte
│   │   │   │   ├── Select.svelte
│   │   │   │   ├── Pagination.svelte
│   │   │   │   └── Spinner.svelte
│   │   │   │
│   │   │   ├── layout/             # Komponen layout
│   │   │   │   ├── Navbar.svelte
│   │   │   │   ├── Footer.svelte
│   │   │   │   ├── Sidebar.svelte      # Untuk dashboard & admin
│   │   │   │   └── Breadcrumb.svelte
│   │   │   │
│   │   │   ├── kelas/              # Komponen spesifik Kelas
│   │   │   │   ├── KelasCard.svelte
│   │   │   │   ├── KelasFilter.svelte
│   │   │   │   └── KelasGrid.svelte
│   │   │   │
│   │   │   ├── konsultasi/         # Komponen spesifik Konsultasi
│   │   │   │   └── FormKonsultasi.svelte
│   │   │   │
│   │   │   └── dashboard/          # Komponen untuk halaman dashboard
│   │   │       ├── StatsCard.svelte
│   │   │       └── JadwalCard.svelte
│   │   │
│   │   ├── stores/                 # Svelte stores (state global)
│   │   │   ├── auth.store.ts       # State user session & role
│   │   │   ├── toast.store.ts      # Notifikasi toast UI
│   │   │   └── kelas.store.ts      # Cache data kelas (opsional)
│   │   │
│   │   ├── services/               # Fungsi API call ke Go backend
│   │   │   ├── api.ts              # Base fetch wrapper (set header JWT, handle error)
│   │   │   ├── auth.service.ts     # Supabase auth calls
│   │   │   ├── kelas.service.ts
│   │   │   ├── topik.service.ts
│   │   │   ├── pendaftaran.service.ts
│   │   │   ├── konsultasi.service.ts
│   │   │   ├── transaksi.service.ts
│   │   │   └── admin.service.ts
│   │   │
│   │   ├── types/                  # TypeScript type definitions
│   │   │   ├── auth.types.ts
│   │   │   ├── kelas.types.ts
│   │   │   ├── topik.types.ts
│   │   │   ├── konsultasi.types.ts
│   │   │   ├── transaksi.types.ts
│   │   │   └── api.types.ts        # Generic response wrapper types
│   │   │
│   │   ├── utils/                  # Fungsi helper
│   │   │   ├── format.ts           # Format angka, tanggal, harga
│   │   │   ├── validator.ts        # Validasi form
│   │   │   └── slug.ts             # Slug generator
│   │   │
│   │   └── supabase.ts             # Inisialisasi Supabase client
│   │
│   └── routes/                     # File-based routing SvelteKit
│       │
│       ├── +layout.svelte          # Root layout (Navbar + Footer)
│       ├── +layout.ts              # Root load function (cek session)
│       │
│       ├── (public)/               # Route group — Halaman Publik
│       │   ├── +layout.svelte      # Layout publik
│       │   ├── +page.svelte        # / — Landing Page
│       │   ├── kelas/
│       │   │   ├── +page.svelte    # /kelas — Daftar kelas
│       │   │   └── [slug]/
│       │   │       └── +page.svelte # /kelas/[slug] — Detail kelas
│       │   ├── topik/
│       │   │   ├── +page.svelte    # /topik — Daftar topik
│       │   │   └── [slug]/
│       │   │       └── +page.svelte # /topik/[slug] — Kelas per topik
│       │   ├── konsultasi/
│       │   │   └── +page.svelte    # /konsultasi
│       │   └── tentang/
│       │       └── +page.svelte    # /tentang
│       │
│       ├── auth/                   # Halaman Auth
│       │   ├── login/
│       │   │   └── +page.svelte    # /auth/login
│       │   ├── register/
│       │   │   └── +page.svelte    # /auth/register
│       │   ├── lupa-password/
│       │   │   └── +page.svelte    # /auth/lupa-password — minta link reset
│       │   └── reset-password/
│       │       └── +page.svelte    # /auth/reset-password — set password baru (dari link email)
│       │
│       ├── dashboard/              # Route group — User Dashboard
│       │   ├── +layout.svelte      # Layout dashboard (Sidebar user)
│       │   ├── +layout.ts          # Guard: redirect ke login jika belum auth
│       │   ├── +page.svelte        # /dashboard — Dashboard (halaman utama user yg sudah login)
│       │   ├── kelas/
│       │   │   ├── +page.svelte    # /dashboard/kelas — Kelas saya
│       │   │   └── [id]/
│       │   │       └── +page.svelte # /dashboard/kelas/[id] — Materi kelas
│       │   ├── konsultasi/
│       │   │   └── +page.svelte    # /dashboard/konsultasi — Riwayat konsultasi
│       │   └── profil/
│       │       └── +page.svelte    # /dashboard/profil
│       │
│       ├── admin/                  # Route group — Admin Panel
│       │   ├── +layout.svelte      # Layout admin (Sidebar admin)
│       │   ├── +layout.ts          # Guard: redirect jika bukan admin
│       │   ├── dashboard/
│       │   │   └── +page.svelte    # /admin/dashboard
│       │   ├── kelas/
│       │   │   ├── +page.svelte    # /admin/kelas — List & kelola kelas
│       │   │   ├── baru/
│       │   │   │   └── +page.svelte # /admin/kelas/baru — Form tambah kelas
│       │   │   └── [id]/
│       │   │       └── +page.svelte # /admin/kelas/[id] — Edit kelas
│       │   ├── peserta/
│       │   │   └── +page.svelte    # /admin/peserta
│       │   ├── konsultasi/
│       │   │   ├── +page.svelte    # /admin/konsultasi — List konsultasi
│       │   │   └── [id]/
│       │   │       └── +page.svelte # /admin/konsultasi/[id] — Detail & respon
│       │   ├── topik/
│       │   │   └── +page.svelte    # /admin/topik
│       │   └── transaksi/
│       │       └── +page.svelte    # /admin/transaksi
│       │
│       └── sertifikat/
│           └── [nomor]/
│               └── +page.svelte    # /sertifikat/[nomor] — Verifikasi sertifikat (publik)
│
├── static/                         # Asset statis
│   ├── favicon.ico
│   ├── logo-tcc.png
│   └── og-image.png                # Open Graph image untuk SEO
│
├── .env                            # Environment variables
├── .env.example                    # Template env (di-commit ke git)
├── svelte.config.js
├── vite.config.ts
├── tsconfig.json
└── package.json
```

### Konvensi Penamaan Frontend

| Jenis | Konvensi | Contoh |
|-------|----------|--------|
| Komponen Svelte | PascalCase | `KelasCard.svelte` |
| File TypeScript | camelCase + suffix | `kelas.service.ts`, `kelas.types.ts` |
| Route folder | kebab-case | `lupa-password/` |
| Store variable | camelCase | `authStore`, `toastStore` |

---

## 3. Backend — Go (Clean Architecture)

### Filosofi Arsitektur
Menggunakan **Clean Architecture** dengan 4 lapisan:

```
Handler (HTTP) → Usecase (Business Logic) → Repository (DB Query) → Database
```

Tiap lapisan hanya boleh berkomunikasi dengan lapisan di bawahnya. **Handler tidak boleh langsung akses DB.**

### Struktur Folder Lengkap

```
backend/
├── cmd/
│   └── server/
│       └── main.go                 # Entry point — init server, DB, router
│
├── internal/
│   │
│   ├── domain/                     # Layer: Entity / Model
│   │   ├── profile.go
│   │   ├── topik.go
│   │   ├── instruktur.go
│   │   ├── kelas.go
│   │   ├── materi_kelas.go
│   │   ├── pendaftaran.go
│   │   ├── konsultasi.go
│   │   ├── transaksi.go
│   │   └── sertifikat.go
│   │
│   ├── repository/                 # Layer: Database Query
│   │   ├── interfaces.go           # Interface definitions untuk semua repo
│   │   ├── profile_repo.go
│   │   ├── topik_repo.go
│   │   ├── instruktur_repo.go
│   │   ├── kelas_repo.go
│   │   ├── pendaftaran_repo.go
│   │   ├── konsultasi_repo.go
│   │   ├── transaksi_repo.go
│   │   └── sertifikat_repo.go
│   │
│   ├── usecase/                    # Layer: Business Logic
│   │   ├── interfaces.go           # Interface definitions untuk semua usecase
│   │   ├── auth_usecase.go
│   │   ├── topik_usecase.go
│   │   ├── instruktur_usecase.go
│   │   ├── kelas_usecase.go
│   │   ├── pendaftaran_usecase.go  # Termasuk validasi konflik jadwal
│   │   ├── konsultasi_usecase.go
│   │   ├── transaksi_usecase.go    # Termasuk verifikasi webhook Midtrans
│   │   └── sertifikat_usecase.go
│   │
│   ├── handler/                    # Layer: HTTP Handler (Controller)
│   │   ├── auth_handler.go
│   │   ├── topik_handler.go
│   │   ├── instruktur_handler.go
│   │   ├── kelas_handler.go
│   │   ├── materi_handler.go
│   │   ├── pendaftaran_handler.go
│   │   ├── konsultasi_handler.go
│   │   ├── transaksi_handler.go
│   │   ├── sertifikat_handler.go
│   │   ├── webhook_handler.go      # Midtrans webhook handler
│   │   └── admin_handler.go        # Admin-specific handlers
│   │
│   ├── middleware/                 # HTTP Middleware
│   │   ├── auth.go                 # Verifikasi JWT dari Supabase
│   │   ├── role.go                 # Cek role (RequireAdmin, RequireUser)
│   │   ├── cors.go                 # CORS config
│   │   └── logger.go              # Request logging
│   │
│   ├── router/
│   │   └── router.go              # Definisi semua route & middleware chain
│   │
│   └── dto/                        # Data Transfer Objects (Request & Response shape)
│       ├── auth_dto.go
│       ├── kelas_dto.go
│       ├── topik_dto.go
│       ├── pendaftaran_dto.go
│       ├── konsultasi_dto.go
│       ├── transaksi_dto.go
│       └── response.go             # Generic response wrapper
│
├── pkg/                            # Package utility yang reusable
│   ├── database/
│   │   └── postgres.go            # Koneksi ke Supabase PostgreSQL
│   ├── supabase/
│   │   └── jwt.go                 # Verifikasi & parse JWT Supabase
│   ├── midtrans/
│   │   └── client.go              # Midtrans SDK wrapper
│   └── utils/
│       ├── slug.go
│       ├── response.go            # Helper build JSON response
│       └── validator.go
│
├── config/
│   └── config.go                  # Load & parse environment variables
│
├── migrations/                    # SQL migration files (urutan eksekusi)
│   ├── 001_create_profiles.sql
│   ├── 002_create_topik.sql
│   ├── 003_create_instruktur.sql
│   ├── 004_create_kelas.sql
│   ├── 005_create_materi_kelas.sql
│   ├── 006_create_pendaftaran.sql
│   ├── 007_create_konsultasi.sql
│   ├── 008_create_transaksi.sql
│   ├── 009_create_sertifikat.sql
│   └── 010_create_triggers.sql    # Supabase trigger: auto-create profile
│
├── .env
├── .env.example
├── go.mod
├── go.sum
├── Makefile                       # Shortcut commands (run, build, migrate, test)
└── README.md
```

### Konvensi Penamaan Backend

| Jenis | Konvensi | Contoh |
|-------|----------|--------|
| File Go | snake_case | `kelas_repo.go` |
| Struct / Interface | PascalCase | `KelasRepository`, `KelasUsecase` |
| Method | PascalCase (exported) | `GetBySlug()`, `Create()` |
| Variable lokal | camelCase | `kelasRepo`, `ctx` |
| Route | kebab-case | `/admin/kelas-baru` |

### Contoh Makefile Commands
```makefile
run:        go run ./cmd/server/main.go
build:      go build -o bin/server ./cmd/server
migrate:    psql $DATABASE_URL -f migrations/...
test:       go test ./...
```

---

## 4. Dependency yang Direncanakan

### Frontend
| Package | Kegunaan |
|---------|----------|
| `@supabase/supabase-js` | Supabase client (Auth, Storage) |
| `@midtrans/midtrans-js` | Midtrans Snap UI |
| `lucide-svelte` | Icon library |
| `date-fns` | Formatting tanggal |

### Backend (Go)
| Package | Kegunaan |
|---------|----------|
| `github.com/gin-gonic/gin` | HTTP Router & framework |
| `github.com/jackc/pgx/v5` | PostgreSQL driver |
| `github.com/golang-jwt/jwt/v5` | JWT parsing & verifikasi |
| `github.com/go-playground/validator/v10` | Request validation |
| `github.com/veritrans/go-midtrans` | Midtrans SDK |
| `github.com/joho/godotenv` | Load .env file |

---

*Dokumen ini akan diperbarui seiring perkembangan project.*
