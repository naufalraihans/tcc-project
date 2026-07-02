# 🎨 UI/UX Design Reference — TCC ITPLN Web Platform

> **Dokumen ini memetakan prototipe desain (dari folder `designUI/`) ke kebutuhan project TCC.**
> Versi: v1.0 | Status: Draft
> Sumber prototipe: **v0.dev** (Next.js + TailwindCSS + shadcn/ui + KokonutUI)

---

## 1. Ringkasan Prototipe yang Tersedia

| Prototipe | Folder | Framework UI | Gaya Desain |
|-----------|--------|-------------|-------------|
| **Landing Page** | `designUI/landingPage/` | shadcn/ui, custom animations | Modern minimalis, dark mode, animated sphere, marquee stats, grid lines, noise texture |
| **Login Page** | `designUI/loginPage/` | shadcn/ui | Glassmorphism card, gradient background, floating orbs, social login buttons |
| **Dashboard** | `designUI/dashboard/` | KokonutUI + shadcn/ui | Sidebar layout, dark mode toggle, card-based content, clean data lists |

---

## 2. Apa yang Diadopsi vs Diubah

### 2.1 Yang DIADOPSI (Konsep Layout & Interaksi)

| Dari Prototipe | Apa yang Diambil |
|----------------|-----------------|
| **Landing Page** | Layout section-based: Hero → Features → How It Works → Metrics → CTA → Footer |
| **Landing Page** | Micro-animations: fade-in on scroll, hover-lift, animated text rotation |
| **Landing Page** | Stats marquee di hero section |
| **Landing Page** | Noise texture overlay untuk kesan premium |
| **Login Page** | Glassmorphism card (backdrop blur + semi-transparent background) |
| **Login Page** | Floating decorative glass orbs di background |
| **Login Page** | Gradient background full-page |
| **Dashboard** | Sidebar + Top Nav + Content area layout |
| **Dashboard** | Collapsible sidebar dengan mobile hamburger menu |
| **Dashboard** | Card-based content grid (2-column grid → full-width) |
| **Dashboard** | Dark/Light mode toggle |

### 2.2 Yang DIUBAH (Warna, Logo, Konten)

| Aspek | Prototipe (Lama) | TCC ITPLN (Baru) |
|-------|------------------|------------------|
| **Warna utama** | Hitam/Putih netral (oklch-based) | Navy Teal `#0C4F6A` + Sky Blue `#1A8DB2` |
| **Warna CTA** | `#0C115B` (dark navy) | `#0C4F6A` (navy teal) |
| **Gradient Hero** | `foreground` → `transparent` | `#0C4F6A` → `#1A8DB2` (navy-teal gradient) |
| **Gradient Login BG** | Generic gradient image | Navy Teal → Sky Blue gradient |
| **Logo sidebar** | KokonutUI logo | Logo TCC (`logo tcc baru (1).png`) |
| **Logo header** | Generic | Logo TCC + "Training & Consulting Center" sub-text |
| **Font display** | `Instrument Serif` | `Plus Jakarta Sans` (Bold/Semibold) |
| **Font body** | `Instrument Sans` | `Inter` (Regular/Medium) |
| **Font mono** | `JetBrains Mono` | Tidak digunakan (project non-tech facing) |
| **Social Login** | Google, Apple, Meta | **Dihapus** — hanya Email/Password via Supabase Auth |
| **Teks hero** | "The platform to create/build/scale" | Konten TCC: kompetensi ketenagalistrikan |
| **Dashboard content** | Accounts, Transactions, Events | Kelas Saya, Jadwal, Konsultasi, Profil |
| **Sidebar menu** | Overview/Finance/Team | Disesuaikan per role (User vs Admin) |
| **Bahasa UI** | English | **Bahasa Indonesia** |

---

## 3. Design Tokens — TCC ITPLN

Mapping dari prototipe ke CSS variables TCC:

### 3.1 Palet Warna (Pengganti oklch prototipe)

```css
:root {
  /* ── Primary ─────────────────────── */
  --navy-teal:    #0C4F6A;    /* Warna utama — header, CTA, sidebar active */
  --sky-blue:     #1A8DB2;    /* Aksen — badge, link, ikon aktif */
  --cool-slate:   #2E4A5A;    /* Hover state, elemen depth */

  /* ── Neutral ─────────────────────── */
  --off-white:    #F4F7FA;    /* Background section, card subtle */
  --white:        #FFFFFF;    /* Background utama */
  --charcoal:     #1A1A2E;    /* Teks heading utama */
  --muted:        #6B7280;    /* Teks sekunder, caption */
  --border:       #E2E8F0;    /* Border card, divider */

  /* ── Semantic ────────────────────── */
  --success:      #1E7B45;    /* Badge "Gratis", status berhasil */
  --warning:      #D97706;    /* Status pending */
  --danger:       #DC2626;    /* Status gagal, error */

  /* ── Glassmorphism (Login Page) ──── */
  --glass-bg:     rgba(255, 255, 255, 0.25);
  --glass-border: rgba(255, 255, 255, 0.4);
  --glass-blur:   blur(40px) saturate(250%);

  /* ── Radius ──────────────────────── */
  --radius:       16px;
  --radius-sm:    8px;
  --radius-full:  9999px;     /* Untuk pill/badge */
}
```

### 3.2 Dark Mode (Dashboard & Admin)

```css
.dark {
  --bg-primary:     #0F0F12;
  --bg-card:        #0F0F12;
  --bg-sidebar:     #0F0F12;
  --border-color:   #1F1F23;
  --text-primary:   #FFFFFF;
  --text-secondary: #A1A1AA;
}
```

> Dark mode hanya aktif di halaman **Dashboard** dan **Admin Panel** — Landing Page tetap light mode.

---

## 4. Mapping Komponen Prototipe → TCC

### 4.1 Landing Page

| Section Prototipe | Fungsi di TCC | Konten |
|-------------------|---------------|--------|
| `Navigation` | Header publik | Logo TCC + nav links (Kelas, Topik, Konsultasi, Tentang) + Masuk/Daftar |
| `HeroSection` | Hero Landing Page | Tagline TCC + stats (peserta, bidang, program) + CTA "Lihat Program" |
| `FeaturesSection` | Kenapa TCC ITPLN | Bento grid: Tiga Pilar (S/W-B/W-H/W) + output Competence-Performance-Certification (data: `profile.md` §4) |
| `HowItWorksSection` | Alur Pendaftaran | Step 1-4: Pilih Kelas → Daftar → Bayar → Ikuti Kelas |
| `MetricsSection` | Statistik Platform | 6.042 peserta, 11 bidang, 100+ program, 6 sertifikasi (data: `profile.md` §3) |
| `IntegrationsSection` | Mitra & Partner | Logo strip — **perlu konfirmasi daftar & file logo resmi** (lihat `profile.md` §10); logo partner belum tersedia |
| `PricingSection` | Kelas Unggulan | Grid kartu kelas — pilih ~3-6 program dari `profile.md` §7 |
| `TestimonialsSection` | *Opsional* — Bisa dihapus atau isi testimonial peserta |
| `SecuritySection` | *Dihapus* — Tidak relevan |
| `DevelopersSection` | *Dihapus* — Tidak relevan |
| `CtaSection` | CTA Konsultasi | "Butuh Konsultansi?" + tombol "Ajukan Konsultasi" |
| `FooterSection` | Footer | Kontak TCC, link navigasi, alamat |

### 4.2 Login Page

| Elemen Prototipe | Adaptasi TCC |
|------------------|-------------|
| Gradient background | Gradient `#0C4F6A` → `#1A8DB2` |
| Floating glass orbs | Tetap, ubah warna ke navy-teal tones |
| Card glassmorphism | Tetap, teks diganti Bahasa Indonesia |
| "Welcome Back" | "Selamat Datang" |
| Email + Password input | Tetap |
| Social login (Google/Apple/Meta) | **Dihapus** — hanya email/password |
| "Forgot password?" | "Lupa password?" + link ke `/auth/lupa-password` |
| *Tambahan* | Link "Belum punya akun? Daftar" → `/auth/register` |
| Button color `#0C115B` | Diganti `#0C4F6A` (navy teal) |

### 4.3 Dashboard (User)

| Elemen Prototipe | Adaptasi TCC |
|------------------|-------------|
| Sidebar — "Overview" group | **Kelas Saya**: Dashboard, Kelas Saya, Riwayat Konsultasi |
| Sidebar — "Finance" group | **Dihapus** — tidak relevan untuk user biasa |
| Sidebar — "Team" group | **Dihapus** — diganti "Akun": Profil, Logout |
| Sidebar — Logo "KokonutUI" | Logo TCC + "Training & Consulting Center" |
| Top Nav — Profile button | Tetap, tampilkan nama user + avatar |
| Content — "Accounts" card | **Kelas Aktif** — daftar kelas yang sedang diikuti |
| Content — "Recent Transactions" card | **Jadwal Mendatang** — kelas yang akan datang |
| Content — "Upcoming Events" card | **Riwayat Konsultasi** — status pengajuan konsultasi |
| Dark/Light toggle | Tetap |

### 4.4 Dashboard (Admin)

| Elemen | Adaptasi TCC |
|--------|-------------|
| Sidebar — Group 1 | **Dashboard**: Statistik global |
| Sidebar — Group 2 | **Kelola**: Kelas, Peserta, Topik, Konsultasi, Transaksi |
| Sidebar — Group 3 | **Akun**: Profil Admin, Logout |
| Content area | Statistik cards (total peserta, kelas aktif, pendapatan) + tabel data |

---

## 5. Animasi & Interaksi yang Dipertahankan

| Animasi | Dari Prototipe | Diterapkan di |
|---------|---------------|--------------|
| `hover-lift` | Landing Page | Semua card (kelas, topik, bento) |
| `animate-char-in` | Landing Page Hero | Teks hero rotating words (opsional) |
| `marquee` | Landing Page Stats | Stats marquee di hero section |
| `noise-overlay` | Landing Page | Background landing page |
| `line-reveal` | Landing Page | Section title pada scroll |
| `backdrop-filter blur` | Login Page | Card login glassmorphism |
| `animate-pulse` | Login Page | Floating glass orbs |
| Sidebar slide transition | Dashboard | Mobile sidebar open/close |

---

## 6. Halaman Register (Baru — Tidak Ada Prototipe)

Mengambil style dari **Login Page** dengan modifikasi:

| Aspek | Detail |
|-------|--------|
| Background | Sama dengan login (gradient + orbs) |
| Card | Glassmorphism sama |
| Title | "Buat Akun Baru" |
| Fields | Nama Lengkap, Email, Password, Konfirmasi Password |
| CTA | "Daftar" (navy teal) |
| Footer | "Sudah punya akun? Masuk" → link ke `/auth/login` |

---

## 7. Responsive Breakpoints

Mengikuti breakpoints dari prototipe:

| Breakpoint | Perilaku |
|------------|----------|
| `>= 1024px` (lg) | Full layout: sidebar visible, 2-column grid |
| `768px – 1023px` (md) | Sidebar hidden (hamburger), content full-width |
| `< 768px` (sm) | Single column, stacked cards, mobile nav |

---

## 8. Daftar Asset yang Diperlukan

| Asset | Sumber | Format |
|-------|--------|--------|
| Logo TCC (warna) | `FILETAMBAH/logo tcc baru (1).png` | PNG |
| Logo TCC (putih) | `FILETAMBAH/Logo TCC putih baru transparan (2).png` | PNG (transparan) |
| Logo ITPLN | `FILETAMBAH/logoITPLN.png` | PNG |
| Logo partner (Danantara, PLN NP, dll) | **Belum tersedia** — perlu file resmi (lihat `profile.md` §10) | PNG/SVG |
| Gradient background login | **Generate** — gradient CSS, tidak perlu gambar | CSS |
| Animated sphere | Dari prototipe landing page | React component (port ke Svelte) |

---

## 9. Catatan Penting untuk Implementasi

### Migrasi Next.js → SvelteKit
Prototipe dibuat di **Next.js + React**, tapi project TCC menggunakan **SvelteKit**. Yang perlu diperhatikan:
- Semua `useState` / `useEffect` → diganti Svelte reactivity (`$:`, `onMount`)
- `shadcn/ui` components → port ke Svelte atau gunakan library Svelte equivalent
- Tailwind CSS tetap bisa dipakai langsung
- `lucide-react` → `lucide-svelte`
- `next/image` → tag `<img>` biasa atau Svelte image component
- `next/link` → Svelte `<a>` atau `goto()` dari `$app/navigation`

### KokonutUI Dashboard
- KokonutUI adalah component library React — **tidak bisa dipakai langsung di Svelte**
- Ambil **konsep layout dan styling-nya** saja, lalu implementasi ulang di Svelte
- Yang perlu di-port: Sidebar, TopNav, Content grid layout

---

*Dokumen ini akan diperbarui saat masuk fase implementasi UI.*
