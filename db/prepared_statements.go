package db

var PSCheckForUsernameAndPasswordCombination = "SELECT username, password, Id FROM nalog WHERE username = $1 AND password = $2"
var PSAddProducts = "INSERT INTO artikal (sif, bc, naz, kolicina) VALUES ($1, $2, $3, $4)"
var PSSearchProducts = `
SELECT a."sif", a."naz", a."bc", a."kolicina", c."cena"
FROM "artikal" a
INNER JOIN "sadrzi_5" s ON a."sif" = s."artikal_sif"
INNER JOIN "ima_cenu" ic ON s."sadrzi_5_id" = ic."sadrzi_5_id"
INNER JOIN "cenovnik" c ON ic."cenovnik_id_cen" = c."id_cen"
WHERE c."pocetak_vazenja" <= CURRENT_DATE 
  AND c."kraj_vazenja" >= CURRENT_DATE
  AND a."naz" ILIKE $1;
`
var PSAddReceipt = "INSERT INTO fiskalni_racun (radi_kasa_trafika_id, radi_kasa_id) VALUES ($1, $2) RETURNING id"
var PSAddInvoice = "INSERT INTO gotovinski_racun (radi_kasa_trafika_id, radi_kasa_id) VALUES ($1, $2) RETURNING id"
var PSCreateReceiptItem = "INSERT INTO stavka (kol, zaprima_sadrzi_4_artikal_sif, cena) VALUES ($1, $2, $3) RETURNING id"
var PSBindItemWithReceipt = "INSERT INTO sadrzi_3 (fiskalni_racun_id, stavka_id) VALUES ($1, $2)"
var PSBindItemWithInvoice = "INSERT INTO sadrzi_2 (gotovinski_racun_id, stavka_id) VALUES ($1, $2)"
var PSBindReceiptWithCashier = "INSERT INTO izdaje (id_racuna, jmbg_kasira) VALUES ($1, $2)"
var PSBindInvoiceWithCashier = "INSERT INTO izdaje_2 (gotovinski_racun_id, jmbg_kasira, registrovan_pravni_pib) VALUES ($1, $2, $3)"
var PSDeleteReceipt = "DELETE FROM fiskalni_racun WHERE id = $1"
var PSDeleteInvoice = "DELETE FROM gotovinski_racun WHERE id = $1"
var PSGetAllReceipts = `
SELECT
	fr.id,
	fr.createdAt,
	t.naz AS trafika_naz,
	z.ime AS zaposleni_ime,
	z.prz AS zaposleni_prz
FROM fiskalni_racun fr
JOIN trafika t ON fr.radi_kasa_trafika_id = t.id
LEFT JOIN izdaje i ON fr.id = i.id_racuna
LEFT JOIN zaposleni z ON i.jmbg_kasira = z.jmbg
LIMIT $1
OFFSET $2;
`
var PSGetAllInvoices = `
SELECT
	gr.id,
	gr.createdAt,
	t.naz AS trafika_naz,
	z.ime AS zaposleni_ime,
	z.prz AS zaposleni_prz,
	i.registrovan_pravni_pib AS pib
FROM gotovinski_racun gr
JOIN trafika t ON gr.radi_kasa_trafika_id = t.id
LEFT JOIN izdaje_2 i ON gr.id = i.gotovinski_racun_id
LEFT JOIN zaposleni z ON i.jmbg_kasira = z.jmbg
LIMIT $1
OFFSET $2;
`
var PSCountAllReceipts = "SELECT COUNT(*) FROM fiskalni_racun"
var PSCountAllInvoices = "SELECT COUNT(*) FROM gotovinski_racun"
var PSAddRequisition = "INSERT INTO trebovanje (poslovodja_jmbg) VALUES $1 RETURNING broj_trebovanja"
var PSDeleteRequisition = "DELETE FROM trebovanje WHERE poslovodja_jmbg = $1"
