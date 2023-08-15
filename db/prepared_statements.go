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
var PSCreateReceiptItem = "INSERT INTO stavka (kol, zaprima_sadrzi_4_artikal_sif, cena) VALUES ($1, $2, $3) RETURNING id"
var PSBindItemWithReceipt = "INSERT INTO sadrzi_3 (fiskalni_racun_id, stavka_id) VALUES ($1, $2)"
var PSBindReceiptWithCashier = "INSERT INTO izdaje (id_racuna, jmbg_kasira) VALUES ($1, $2)"
var PSDeleteReceipt = "DELETE FROM fiskalni_racun WHERE id = $1"
