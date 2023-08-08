package db

var PSCheckForUsernameAndPasswordCombination = "SELECT username, password, Id FROM nalog WHERE username = $1 AND password = $2"
var PSAddProducts = "INSERT INTO artikal (sif, bc, naz) VALUES ($1, $2, $3)"
var PSSearchProducts = "SELECT * FROM artikal WHERE naz LIKE $1"
