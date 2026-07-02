# 💳 Midtrans Integration — TCC ITPLN Web Platform

> **Dokumen ini menjelaskan alur integrasi pembayaran Midtrans secara detail.**
> Versi: v1.0 | Status: Draft
> Mode: **Sandbox** (development) → **Production** (live)

---

## 1. Gambaran Alur Pembayaran

```
[User] → Klik "Daftar" kelas berbayar
    │
    ▼
[SvelteKit] → POST /api/v1/pendaftaran { kelas_id }
    │
    ▼
[Go Backend]
  1. Validasi: user belum daftar kelas ini
  2. Validasi: jadwal tidak bentrok
  3. Validasi: kuota masih tersedia
  4. Buat record transaksi (status: 'pending')
  5. Request Snap Token ke Midtrans API
  6. Return snap_token ke frontend
    │
    ▼
[SvelteKit] → Buka Midtrans Snap UI pakai snap_token
    │
[User bayar di Snap UI]
    │
    ▼
[Midtrans] → Kirim webhook notification ke Go backend
    │
    ▼
[Go Backend Webhook Handler]
  1. Verifikasi signature key
  2. Update transaksi (status: 'sukses')
  3. Enroll atomik: UPDATE kelas +1 (guard kuota) → INSERT pendaftaran (status: 'aktif')
     (jika kuota penuh → refund, lihat §6.3 edge case)
    │
    ▼
[User] → Diarahkan ke halaman sukses
```

---

## 2. Komponen yang Terlibat

| Komponen | Tugas |
|----------|-------|
| **Frontend (SvelteKit)** | Tampilkan Snap UI, handle callback sukses/gagal dari Snap |
| **Go Backend** | Buat transaksi, request Snap Token, verifikasi webhook |
| **Midtrans Snap** | UI pembayaran (popup/redirect), proses payment |
| **Midtrans Notification** | Kirim webhook ke backend saat status berubah |

---

## 3. Format Order ID

Order ID harus **unik** dan dikirim ke Midtrans. Format yang digunakan:

```
TCC-{6-char-random-uppercase}-{unix-timestamp}

Contoh:
TCC-A3XK91-1751234567
```

**Aturan Order ID:**
- Maksimal 50 karakter
- Hanya alphanumeric, `-`, `_`
- Wajib unik — jika order ID yang sama dikirim ulang ke Midtrans, akan error
- Disimpan di kolom `transaksi.midtrans_order_id`

---

## 4. Detail Langkah di Go Backend

### 4.1 Saat User Klik "Daftar" (Create Transaction)

```
Handler: POST /api/v1/pendaftaran

Usecase — pendaftaran_usecase.go:
  1. Cek duplikat: SELECT dari pendaftaran WHERE user_id=X AND kelas_id=Y
  2. Cek konflik jadwal (lihat bagian 7)
  3. Cek kuota awal (validasi cepat; enforcement sesungguhnya di enroll atomik)

  Jika kelas gratis:
    → Enroll atomik dalam 1 transaksi (lihat 06 §3):
        UPDATE kelas +1 dgn guard kuota → jika 0 baris, return error KUOTA_PENUH
        → INSERT pendaftaran (status: 'aktif')
    → Return { type: "gratis", pendaftaran_id }

  Jika kelas berbayar:
    → Generate order_id: TCC-{random}-{timestamp}
    → INSERT transaksi (status: 'pending', midtrans_order_id = order_id)
    → Request Snap Token ke Midtrans (lihat 4.2)
    → Return { type: "berbayar", transaksi_id, snap_token, redirect_url }
```

### 4.2 Request Snap Token ke Midtrans

Go backend memanggil Midtrans Snap API:

```
POST https://app.sandbox.midtrans.com/snap/v1/transactions
Authorization: Basic base64(ServerKey + ":")
Content-Type: application/json

Body:
{
  "transaction_details": {
    "order_id": "TCC-A3XK91-1751234567",
    "gross_amount": 500000
  },
  "customer_details": {
    "first_name": "Budi",
    "last_name": "Santoso",
    "email": "budi@example.com",
    "phone": "08123456789"
  },
  "item_details": [
    {
      "id": "kelas-uuid",
      "price": 500000,
      "quantity": 1,
      "name": "Workshop K3 Listrik Dasar"
    }
  ],
  "callbacks": {
    "finish": "https://tcc-itpln.id/pembayaran/sukses",
    "error": "https://tcc-itpln.id/pembayaran/gagal",
    "pending": "https://tcc-itpln.id/pembayaran/pending"
  }
}
```

**Response dari Midtrans:**
```json
{
  "token": "snap-token-abc123xyz",
  "redirect_url": "https://app.sandbox.midtrans.com/snap/v2/vtweb/snap-token-abc123xyz"
}
```

---

## 5. Frontend — Tampilkan Snap UI

### 5.1 Load Midtrans Snap.js

Di `app.html` atau load dinamis:
```html
<script
  src="https://app.sandbox.midtrans.com/snap/snap.js"
  data-client-key="SB-Mid-client-xxxxx">
</script>
```
> Untuk production: ganti ke `https://app.midtrans.com/snap/snap.js`

### 5.2 Buka Popup Snap

```typescript
// Di halaman detail kelas atau setelah backend return snap_token
const handleDaftar = async (kelasId: string) => {
  // 1. Hit backend untuk buat transaksi
  const res = await fetch('/api/v1/pendaftaran', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      'Authorization': `Bearer ${session.access_token}`
    },
    body: JSON.stringify({ kelas_id: kelasId })
  })

  const data = await res.json()

  if (data.data.type === 'gratis') {
    // Langsung redirect ke dashboard
    goto('/dashboard/kelas')
    return
  }

  // 2. Buka Snap UI
  window.snap.pay(data.data.snap_token, {
    onSuccess: (result) => {
      // User berhasil bayar — redirect ke halaman sukses
      goto(`/pembayaran/sukses?order_id=${result.order_id}`)
    },
    onPending: (result) => {
      // Pembayaran pending (transfer bank belum dikonfirmasi)
      goto(`/pembayaran/pending?order_id=${result.order_id}`)
    },
    onError: (result) => {
      // Pembayaran gagal
      goto(`/pembayaran/gagal?order_id=${result.order_id}`)
    },
    onClose: () => {
      // User tutup popup tanpa bayar
      // Tidak redirect — biarkan user di halaman yang sama
      // Tampilkan toast: "Pembayaran dibatalkan. Kamu bisa mencoba lagi."
    }
  })
}
```

---

## 6. Webhook Handler di Go Backend

### 6.1 Endpoint Webhook

```
POST /webhook/midtrans
```

Midtrans akan mengirim notifikasi ke URL ini setiap kali status transaksi berubah (pending → sukses, pending → gagal, dsb).

> **Penting:** URL ini harus bisa diakses publik oleh Midtrans server. Untuk development, gunakan **ngrok** atau **Cloudflare Tunnel** untuk expose localhost.

### 6.2 Payload dari Midtrans

```json
{
  "transaction_time": "2026-08-01 10:30:00",
  "transaction_status": "settlement",
  "transaction_id": "midtrans-internal-txn-id",
  "order_id": "TCC-A3XK91-1751234567",
  "merchant_id": "G12345678",
  "gross_amount": "500000.00",
  "currency": "IDR",
  "payment_type": "bank_transfer",
  "bank": "bca",
  "signature_key": "hash_sha512_dari_order_id+status_code+gross_amount+server_key"
}
```

### 6.3 Logika Webhook Handler

```
1. VERIFIKASI SIGNATURE (wajib, jangan skip)
   ─────────────────────────────────────────
   Expected signature = SHA512(order_id + status_code + gross_amount + server_key)
   Jika tidak cocok → return 400, abaikan request
   (Ini mencegah request palsu dari pihak luar)

2. AMBIL DATA TRANSAKSI
   ─────────────────────
   SELECT * FROM transaksi WHERE midtrans_order_id = order_id
   Jika tidak ditemukan → return 404

3. CEK IDEMPOTENCY
   ────────────────
   Jika transaksi.status sudah 'sukses' → return 200 (jangan proses ulang)
   Midtrans bisa kirim webhook lebih dari sekali untuk order yang sama

4. PROSES BERDASARKAN STATUS
   ──────────────────────────
   transaction_status = 'settlement' atau 'capture'
     → UPDATE transaksi SET status='sukses', midtrans_txn_id=..., metode_pembayaran=...
     → panggil RUTIN ENROLL ATOMIK (sama persis dgn kelas gratis, lihat 06 §3):
         UPDATE kelas +1 dgn guard kuota (WHERE peserta_terdaftar < kuota)
         → INSERT pendaftaran (status='aktif')
       (counter dinaikkan DI SINI saja — JANGAN tambah manual, tidak ada trigger increment)
     → UPDATE transaksi SET pendaftaran_id = [id pendaftaran baru]

   transaction_status = 'deny' atau 'cancel' atau 'expire'
     → UPDATE transaksi SET status='gagal'

   transaction_status = 'refund'
     → UPDATE transaksi SET status='refund'
     → panggil RUTIN UN-ENROLL ATOMIK (jika pendaftaran ada):
         UPDATE kelas -1 (+ balikin status 'penuh'→'aktif')
         → UPDATE pendaftaran SET status='dibatalkan'

5. RETURN 200 OK ke Midtrans
   ──────────────────────────
   Jika tidak return 200, Midtrans akan retry webhook (hingga 5x)
```

> **Edge case — bayar tapi slot habis.** Kuota TIDAK di-reserve selama transaksi `pending`. Cek kuota di §4.1 hanya validasi awal; slot baru benar-benar diambil saat webhook sukses menjalankan enroll atomik. Jadi dua orang bisa sama-sama bayar seat terakhir. Untuk pembayar kedua, rutin enroll atomik mengembalikan **0 baris** (kuota penuh) → **jangan** buat pendaftaran.
> Penanganan: proses refund Midtrans → set `transaksi.status='refund'` + notifikasi user (uang balik, gagal masuk kelas). Tetap **return 200** ke Midtrans (webhook sudah diterima; masalah kuota bukan urusan retry Midtrans).
> *Alternatif kalau overbooking mau dihindari total:* reserve slot saat buat transaksi pending & lepas saat expire — lebih rumit, tunda sampai perlu.

---

## 7. Validasi Konflik Jadwal

Dilakukan di Go backend **sebelum** buat transaksi/pendaftaran:

```sql
-- Cek apakah user punya kelas lain yang jadwalnya overlap
SELECT k.judul, k.jadwal_mulai, k.jadwal_selesai
FROM pendaftaran p
JOIN kelas k ON p.kelas_id = k.id
WHERE p.user_id = $1
  AND p.status = 'aktif'
  AND k.jadwal_mulai < $3   -- jadwal_selesai kelas baru
  AND k.jadwal_selesai > $2 -- jadwal_mulai kelas baru
LIMIT 1;
```

Jika ada result → return error:
```json
{
  "success": false,
  "error": "SCHEDULE_CONFLICT",
  "message": "Jadwal bentrok dengan kelas 'Workshop Smart Grid' (01 Agt 2026, 08.00–17.00)"
}
```

---

## 8. Halaman Status Pembayaran (Frontend)

| Route | Keterangan |
|-------|------------|
| `/pembayaran/sukses` | Tampilkan ringkasan pembayaran + CTA ke dashboard |
| `/pembayaran/pending` | Informasi cara transfer + instruksi bank |
| `/pembayaran/gagal` | Pesan gagal + tombol "Coba Lagi" |

### Catatan Penting untuk Halaman Sukses
> Jangan langsung tampilkan data kelas dari sisi frontend berdasarkan callback Snap. Halaman sukses hanya menampilkan ringkasan order. **Sumber kebenaran adalah webhook** — pendaftaran di DB baru terbentuk setelah webhook diterima dan diproses backend.
>
> Akibatnya, ada jeda singkat antara Snap callback "onSuccess" dan data pendaftaran muncul di `/dashboard/kelas`. Ini normal — tampilkan pesan: *"Pendaftaran kamu sedang diproses, cek kembali dalam beberapa saat."*

---

## 9. Mode Sandbox vs Production

| Setting | Sandbox (Dev) | Production (Live) |
|---------|--------------|-------------------|
| Snap.js URL | `app.sandbox.midtrans.com/snap/snap.js` | `app.midtrans.com/snap/snap.js` |
| API URL | `app.sandbox.midtrans.com` | `app.midtrans.com` |
| Server Key prefix | `SB-Mid-server-` | `Mid-server-` |
| Client Key prefix | `SB-Mid-client-` | `Mid-client-` |
| Kartu test | Ada (dari docs Midtrans) | Kartu nyata |

> Di Go backend, mode diatur lewat env var `MIDTRANS_IS_PRODUCTION=false/true`. SDK Go Midtrans akan otomatis pakai URL yang sesuai.

---

## 10. Checklist Integrasi Midtrans

```
[ ] Daftar akun Midtrans Sandbox di dashboard.sandbox.midtrans.com
[ ] Ambil Server Key & Client Key → isi ke .env
[ ] Set Notification URL di Midtrans Dashboard → URL webhook Go backend
[ ] Untuk dev: setup ngrok/tunnel agar webhook bisa hit localhost
[ ] Test flow lengkap dengan kartu test Midtrans
[ ] Verifikasi idempotency: kirim webhook duplikat → pastikan tidak double-insert
[ ] Verifikasi signature check: kirim request tanpa signature → pastikan ditolak
[ ] Ganti ke Production keys sebelum deploy
```

---

*Dokumen ini akan diperbarui seiring implementasi pembayaran.*
