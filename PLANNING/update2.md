# INTEGRASI AI Chat Bot
Tujuan dari penggunaan integrasi AI Chat bot ini adalah untuk menggantikan FAQ yang biasanya dibuat secara statis. 
kali ini, kami ganti dengan sistem AI dengan tentunya menggunakan sistem RAG sebagai otak nya. 

sehingga, AI akan menjawab berdasarkan pengetahuannya secara umum dan juga secara sistem RAG 


- provider = bynara

## Cara penggunaan
baseUrl = `https://router.bynara.id/v1`

API KEY = `<disimpan di backend/.env sebagai BYNARA_API_KEY — JANGAN taruh di dokumen>`

- Cara penggunaan 

```
curl https://router.bynara.id/v1/chat/completions -H "Authorization: Bearer $BYNARA_API_KEY" -H "Content-Type: application/json" -d '{"model":"mistral-medium-3-5","messages":[{"role":"user","content":"Hello"}]}'
```

aku ada limit dari API Key tersebut adaalah 60 req/min 
