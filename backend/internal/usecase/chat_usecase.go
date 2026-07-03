package usecase

import (
	"context"
	"fmt"
	"strings"

	"tcc-itpln/backend/internal/dto"
	"tcc-itpln/backend/internal/repository"
	"tcc-itpln/backend/pkg/bynara"
)

var ErrChatDisabled = errSentinelUC("chatbot belum aktif (BYNARA_API_KEY belum diset)")

// Pengetahuan statik TCC (RAG-lite dari PLANNING/profile.md).
const tccKnowledge = `Kamu adalah asisten AI resmi Training & Consulting Center ITPLN (TCC ITPLN). Jawab pertanyaan pengunjung website secara ringkas, ramah, dan akurat dalam Bahasa Indonesia.

Aturan:
- Jawab berdasarkan PENGETAHUAN dan KATALOG KELAS di bawah. Pengetahuan umum boleh, tapi untuk fakta spesifik TCC (harga, jadwal, nama, angka) jangan mengarang di luar yang tersedia.
- Jika informasi tidak ada, katakan belum tersedia dan arahkan ke kontak resmi.
- Tanpa emoji. Hindari format markdown berlebihan. Ringkas 2-5 kalimat kecuali diminta detail.

PENGETAHUAN TCC ITPLN:
- Identitas: Training & Consulting Center ITPLN (TCC ITPLN), unit bisnis di bawah Institut Teknologi PLN. Berdiri 7 Februari 2023, bertransformasi menjadi Training & Consulting Center pada Mei 2026.
- Bidang: layanan pelatihan & konsultasi di energi, teknologi berwawasan lingkungan, dan manajemen.
- Direktur TCC: Ir. Suharto, M.T., IPU.
- Statistik: peserta tumbuh dari 1.024 (tahun 1) menjadi 6.042 (tahun 3). 11 bidang pelatihan, sekitar 100 program di katalog, 6 program sertifikasi.
- 11 bidang pelatihan: Engineering, Digital Platform, Manajemen Aset, Manajemen Risiko, Manajemen Keuangan, Sumber Daya Manusia, EPC, EBT (Renewable Energy), OHS/K3, Bimtek, Workshop.
- Framework 3 pilar: Software (kurikulum/konten), Brainware (instruktur/mentor), Hardware (smart classroom & infrastruktur). Output: Competence, Performance, Certification.
- Produk: Training/Workshop, Konsultansi/Pendampingan, Kemitraan/Bimtek. Pelatihan dapat tailor-made sesuai kebutuhan klien.
- Contoh program unggulan: Manajemen Aset Pembangkit ISO 55001 (16 JP), Workshop K3 & 5S (24 JP), Basic Theory Power Plant (8 JP), PLTS (8 JP), Waste to Energy (24 JP), International Training AMI/Smart Grid (40 JP).
- Program sertifikasi: PSPPI (Insinyur), Ketenagalistrikan, Microsoft Office Specialist (MOS), MikroTik (MTCNA), TKBI (Bahasa Inggris), TKDA (Kemampuan Dasar Akademik).
- Metode belajar: Online / Offline / Hybrid / Blended.
- Kontak: email trainingcenter@itpln.ac.id; kantor Graha YPK PLN, Jl. Lebak Bulus Tengah No.5, Cilandak, Jakarta Selatan 12430; Anna Agustina +62 813-1084-5077; Dwi Listiawati +62 815-1405-6864.
- Untuk melihat & mendaftar kelas, arahkan pengunjung ke menu Kelas di website atau login ke dashboard.`

type chatUsecase struct {
	kelas repository.KelasRepository
	bot   *bynara.Client
}

func NewChatUsecase(kelas repository.KelasRepository, bot *bynara.Client) ChatUsecase {
	return &chatUsecase{kelas, bot}
}

func (u *chatUsecase) Reply(ctx context.Context, msgs []dto.ChatMessage) (string, error) {
	if !u.bot.Enabled() {
		return "", ErrChatDisabled
	}

	system := tccKnowledge + "\n\n" + u.katalogContext(ctx)
	out := []bynara.Message{{Role: "system", Content: system}}

	// hanya 12 pesan terakhir, role user/assistant, batasi panjang konten
	start := 0
	if len(msgs) > 12 {
		start = len(msgs) - 12
	}
	for _, m := range msgs[start:] {
		if m.Role != "user" && m.Role != "assistant" {
			continue
		}
		c := m.Content
		if len(c) > 2000 {
			c = c[:2000]
		}
		out = append(out, bynara.Message{Role: m.Role, Content: c})
	}

	return u.bot.Chat(ctx, out)
}

// katalogContext menarik kelas aktif LIVE dari DB agar bot tahu katalog terkini.
func (u *chatUsecase) katalogContext(ctx context.Context) string {
	paged, err := u.kelas.List(ctx, dto.KelasFilter{Status: "aktif", Page: 1, Limit: 50})
	if err != nil || len(paged.Items) == 0 {
		return "KATALOG KELAS AKTIF SAAT INI: belum ada kelas yang terdaftar di sistem."
	}
	var b strings.Builder
	b.WriteString("KATALOG KELAS AKTIF SAAT INI (data live dari sistem):\n")
	for _, k := range paged.Items {
		topik := "-"
		if k.Topik != nil {
			topik = k.Topik.Nama
		}
		biaya := "Gratis"
		if k.TipeHarga == "berbayar" {
			biaya = fmt.Sprintf("Rp %.0f", k.Harga)
		}
		b.WriteString(fmt.Sprintf("- %s | bidang: %s | format: %s | biaya: %s\n", k.Judul, topik, k.Format, biaya))
	}
	return b.String()
}
