package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/google/uuid"
)

type telkayıt struct {
	ID     string
	Isim   string
	Numara string
}

func main() {
	var telrehber []telkayıt

	loadPhonebook(&telrehber)

	for {
		var secim int
		fmt.Println("1. Kişi Ekle")
		fmt.Println("2. Kişileri Listele")
		fmt.Println("3. Kişi Sil")
		fmt.Println("4. Kayıt Et ve Çık")
		fmt.Print("Seçim Yapınız: ")
		fmt.Scanln(&secim)

		switch secim {
		case 1:
			addPerson(&telrehber)
		case 2:
			listPeople(&telrehber)
		case 3:
			deletePerson(&telrehber)
		case 4:
			savePhonebook(&telrehber)
			fmt.Println("Programdan çıkılıyor.")
			return
		default:
			fmt.Println("Geçersiz seçim. Tekrar deneyin.")
		}
	}
}

func loadPhonebook(phonebook *[]telkayıt) {
	file, err := os.Open("rehber.txt")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Rehber dosyası bulunamadı veya bir hata oluştu.")
			return
		}
		log.Fatalf("Dosya yüklenemedi: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) == 3 {
			*phonebook = append(*phonebook, telkayıt{ID: fields[0], Isim: fields[1], Numara: fields[2]})
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Dosya okunurken bir hata oluştu: %v", err)
	}
}

func savePhonebook(phonebook *[]telkayıt) {
	file, err := os.Create("rehber.txt")
	if err != nil {
		fmt.Println("Dosya oluşturulamadı:", err)
		return
	}
	defer file.Close()

	for _, person := range *phonebook {
		_, err := fmt.Fprintf(file, "%s %s %s\n", person.ID, person.Isim, person.Numara)
		if err != nil {
			fmt.Println("Dosyaya yazma hatası:", err)
			return
		}
	}

	fmt.Println("Rehber dosyasına kaydedildi.")
}

func addPerson(phonebook *[]telkayıt) {
	var yeniKayit telkayıt
	yeniKayit.ID = uuid.New().String()
	fmt.Print("Ad Giriniz: ")
	fmt.Scanln(&yeniKayit.Isim)
	fmt.Print("Numara Giriniz: ")
	fmt.Scanln(&yeniKayit.Numara)
	*phonebook = append(*phonebook, yeniKayit)
	fmt.Println("Kişi eklenmiştir:", yeniKayit)

	// Eklenen kişinin rehberde görünüp görünmediğini kontrol edelim
	listPeople(phonebook)
}

func deletePerson(phonebook *[]telkayıt) {
	fmt.Print("Silinecek kişinin ID'sini giriniz: ")
	var silinecekID string
	fmt.Scanln(&silinecekID)
	for i, kisi := range *phonebook {
		if kisi.ID == silinecekID {
			*phonebook = append((*phonebook)[:i], (*phonebook)[i+1:]...)
			fmt.Println("Kişi silindi")
			return
		}
	}
	fmt.Println("ID ile eşleşen kişi bulunamadı.")
}

func listPeople(phonebook *[]telkayıt) {
	fmt.Println("Telefon rehberi:")
	for _, kisi := range *phonebook {
		fmt.Printf("ID: %s\tİsim: %s\tNumara: %s\n", kisi.ID, kisi.Isim, kisi.Numara)
	}
}
