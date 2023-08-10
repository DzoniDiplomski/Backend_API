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
