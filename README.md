# Programmes du livre : [Apprenez à maîtriser le langage de programmation Go](https://www.amazon.fr/Apprenez-ma%C3%AEtriser-langage-programmation-Go/dp/B0BW345PWR/ref=sr_1_7?__mk_fr_FR=%C3%85M%C3%85%C5%BD%C3%95%C3%91&crid=2QJL21OFZPZOB&keywords=programmation+go&qid=1676891590&sprefix=programmation+go%2Caps%2C115&sr=8-7)
[![Couverture du livre](couverture.png)](https://www.amazon.fr/Apprenez-ma%C3%AEtriser-langage-programmation-Go/dp/B0BW345PWR/ref=sr_1_7?__mk_fr_FR=%C3%85M%C3%85%C5%BD%C3%95%C3%91&crid=2QJL21OFZPZOB&keywords=programmation+go&qid=1676891590&sprefix=programmation+go%2Caps%2C115&sr=8-7)

Auteur : [Hamdy Abou El Anein](https://www.amazon.fr/Hamdy-Abou-El-Anein/e/B0BM89T88X/ref=aufs_dp_fta_dsk)


## Maintenance and tests

Maintainer: Hamdy Abou El Anein — hamdy.aea@protonmail.com

Root examples are independent programs: run or build each file individually. The password lesson now uses crypto/rand with bounded lengths and no indexing past the output buffer. Test it with `go test PassGen.go PassGen_test.go`. `listes2.go` now prints its array and the previously empty `main.go` is a valid placeholder. All standalone root programs compile independently, and all six nested modules pass `go test ./...` (most have compile checks only). No HTTP example or web server was run. Third-party references, including the LogRocket example attribution, are retained.
