package library

import "fmt"

type mahasiswa struct {
  nama string
  umur int
}

func thanks(bahasa string) {
  switch bahasa {
  case "ID":
    fmt.Println("Terima kasih")
  default:
    fmt.Println("Thank you")
  }

}

func (s mahasiswa) perkenalan() {
  fmt.Printf("Perkenalkan Nama Saya %s, umur saya %d tahun ", s.nama, s.umur)
  fmt.Println()
}

func Salam(){
  var saya = mahasiswa{"Dezza", 25}
  fmt.Println("Hallo")
  saya.perkenalan()
  thanks("ID")
}
