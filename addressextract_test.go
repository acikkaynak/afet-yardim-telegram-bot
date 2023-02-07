package main

import "testing"

func TestExtractCity(t *testing.T) {
	address := "Hayrullah Mah. Malik Ejder Cad. 33023 Sokak Hilal Apt. 4. Ankara Kat Onikişubat - gaziantep (Ortahal civarı, Kültürpark Karşısı, 46Ambar Ayakkabının olduğu bina"
	city := ExtractCity(address)

	if city != ANTEP {
		t.Fatalf("%s was waiting %s", city, ANTEP)
	}
}

func TestExtractCityUpperCase(t *testing.T) {
	address := "Hayrullah Mah. Malik Ejder Cad. 33023 Sokak Hilal Apt. 4. Ankara Kat Onikişubat - GAZİANTEP (Ortahal civarı, Kültürpark Karşısı, 46Ambar Ayakkabının olduğu bina"
	city := ExtractCity(address)

	if city != ANTEP {
		t.Fatalf("%s was waiting %s", city, ANTEP)
	}
}

func TestExtractCityMixCaseShort(t *testing.T) {
	address := "Hayrullah Mah. Malik Ejder Cad. 33023 Sokak Hilal Apt. 4. Ankara Kat Onikişubat - AnTeP (Ortahal civarı, Kültürpark Karşısı, 46Ambar Ayakkabının olduğu bina"
	city := ExtractCity(address)

	if city != ANTEP {
		t.Fatalf("%s was waiting %s", city, ANTEP)
	}
}

func TestExtractDistrict(t *testing.T) {
	address := "Hayrullah Mah. Malik Ejder Cad. 33023 Sokak Hilal Apt. 4. Ankara Kat Onikişubat - Kahramanmaraş (Ortahal civarı, Kültürpark Karşısı, 46Ambar Ayakkabının olduğu bina"
	city := MARAS

	district := ExtractDistrict(City(city), address)
	if string(district) != "Onikişubat" {
		t.Fail()
	}
}

func TestExtractDistrictUpperCase(t *testing.T) {
	address := "Hayrullah Mah. Malik Ejder Cad. 33023 Sokak Hilal Apt. 4. Ankara Kat ONİKİŞUBAT - Kahramanmaraş (Ortahal civarı, Kültürpark Karşısı, 46Ambar Ayakkabının olduğu bina"
	city := MARAS

	district := ExtractDistrict(City(city), address)
	if string(district) != "Onikişubat" {
		t.Fail()
	}
}
