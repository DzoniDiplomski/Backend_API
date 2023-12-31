package db

var PSCheckForUsernameAndPasswordCombination = `SELECT n.username, n.password, n.zaposleni_jmbg, z.spec
FROM nalog n
JOIN zaposleni z ON n.zaposleni_jmbg = z.jmbg
WHERE n.username = $1 AND n.password = $2;
`
var PSAddProducts = "INSERT INTO artikal (sif, bc, naz) VALUES ($1, $2, $3) RETURNING sif"
var PSSearchProducts = `
SELECT a."sif", a."naz", a."bc", s."kolicina", c."cena"
FROM "artikal" a
INNER JOIN "sadrzi_5" s ON a."sif" = s."artikal_sif"
INNER JOIN "ima_cenu" ic ON s."artikal_sif" = ic."sif"
INNER JOIN "cenovnik" c ON ic."cenovnik_id_cen" = c."id_cen"
WHERE c."pocetak_vazenja" <= CURRENT_DATE 
  AND c."kraj_vazenja" >= CURRENT_DATE
  AND a."naz" ILIKE $1
  AND ic."status" = true;
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
WHERE fr.createdAt::date = CURRENT_DATE
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
WHERE gr.createdAt::date = CURRENT_DATE
LIMIT $1
OFFSET $2;
`
var PSCountAllReceipts = "SELECT COUNT(*) FROM fiskalni_racun WHERE createdAt::date = CURRENT_DATE"
var PSCountAllInvoices = "SELECT COUNT(*) FROM gotovinski_racun WHERE createdAt::date = CURRENT_DATE"
var PSCountAllRequisitions = "SELECT COUNT(*) FROM trebovanje"
var PSAddRequisition = "INSERT INTO trebovanje (poslovodja_jmbg) VALUES ($1) RETURNING broj_trebovanja"
var PSDeleteRequisition = "DELETE FROM trebovanje WHERE poslovodja_jmbg = $1"
var PSCreateRequisitionItem = "INSERT INTO stavka_trebovanja (naz, kol, broj_trebovanja) VALUES ($1, $2, $3) RETURNING id"
var PSReduceItemQuantity = "UPDATE sadrzi_5 SET kolicina = kolicina - $1 WHERE artikal_sif = $2;"
var PSIncreaseItemQuantity = "UPDATE sadrzi_5 SET kolicina = kolicina + $1 WHERE artikal_sif = $2;"
var PSUpdateProductPrice = "INSERT INTO cenovnik (cena, pocetak_vazenja, kraj_vazenja) VALUES ($1, $2, $3) RETURNING id_cen"
var PSBindPriceWithProduct = "INSERT INTO ima_cenu (cenovnik_id_cen, sif, status) VALUES ($1, $2, $3)"
var PSRevokeAllPrices = `UPDATE ima_cenu
SET status = false
WHERE sif = $1 AND cenovnik_id_cen <> $2;
`
var PSAddSumToReceipt = `UPDATE fiskalni_racun SET suma = $1 WHERE id = $2`
var PSAddSumToInvoice = `UPDATE gotovinski_racun SET suma = $1 WHERE id = $2`
var PSGetTodaysMarket = `SELECT * from dnevni_pazar WHERE datum = CURRENT_DATE`
var PSCreateTodaysMarket = `INSERT INTO dnevni_pazar DEFAULT VALUES  RETURNING id`
var PSUpdateTodaysMarketSum = `UPDATE dnevni_pazar SET suma = suma + $1 WHERE id = $2`
var PSGetProductPricesOverTime = `SELECT c.cena, c.pocetak_vazenja, c.kraj_vazenja
FROM ima_cenu ic
JOIN cenovnik c ON ic.cenovnik_id_cen = c.id_cen
WHERE ic.sif = $1;`
var PSGetAllRequisitions = `
SELECT t.broj_trebovanja, t.createdAt
FROM trebovanje t
ORDER BY t.createdAt DESC
LIMIT $1
OFFSET $2;
`
var PSGetRequisitionItems = `
SELECT st.kol, st.naz
FROM stavka_trebovanja st
WHERE broj_trebovanja = $1`
var PSGetReceiptItems = `
SELECT a."sif", a."naz", s."kol", s."cena"
FROM fiskalni_racun fr
JOIN sadrzi_3 s3 ON fr."id" = s3."fiskalni_racun_id"
JOIN stavka s ON s3."stavka_id" = s."id"
JOIN artikal a ON s."zaprima_sadrzi_4_artikal_sif" = a."sif"
WHERE fr."id" = $1;
`
var PSGetInvoiceItems = `
SELECT a."sif", a."naz", s."kol", s."cena"
FROM gotovinski_racun gr
JOIN sadrzi_2 s2 ON gr."id" = s2."gotovinski_racun_id"
JOIN stavka s ON s2."stavka_id" = s."id"
JOIN artikal a ON s."zaprima_sadrzi_4_artikal_sif" = a."sif"
WHERE gr."id" = $1;`
var PSCreateCalculation = `INSERT INTO kalkulacija DEFAULT VALUES RETURNING id`
var PSCreateCalculationItem = `INSERT INTO stavka_kalkulacije (sif, neto_cena, marza, kolicina, pdv_stopa) VALUES ($1, $2, $3, $4, $5) RETURNING id`
var PSBindCalculationItemWithCalculation = `
INSERT INTO sadrzi_6 (id_kalkulacije, id_stavke) VALUES ($1, $2)`
var PSAddProductToStorage = `
INSERT INTO sadrzi_5 (artikal_sif, magacin_id, kolicina) VALUES ($1, $2, $3)
`
var PSGetAllCalculations = `
SELECT * FROM kalkulacije`
var PSCountAllCalculations = `SELECT COUNT(*) FROM kalkulacija
`
