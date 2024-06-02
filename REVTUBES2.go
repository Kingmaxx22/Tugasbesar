package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Login struct {
	uid    string
	pwd    string
	bio    string
	teman  [1024]string
	temanN int
}
type Post struct {
	main  string
	comm  [1024]KOM
	commN int
	op    string
}

type KOM struct {
	main string
	op   string
}

var login [1024]Login
var accN int
var postN int
var post [10000]Post
var aktif bool = true
var cuser string

func main() {
	for aktif {
		uiawal()
	}
}

func uiawal() {

	var a int
	fmt.Println("=================")
	fmt.Println("  SOCIAL MEDIA")
	fmt.Println("=================")
	fmt.Printf("1.Login\n2.Daftar Akun\n3.Keluar Aplikasi\n")
	fmt.Print("Pilih 1/2/3 :")
	fmt.Scan(&a)
	switch a {
	case 1:
		masuk()
	case 2:
		daftar()
	case 3:
		aktif = false
	default:
		fmt.Println("Pilihan Anda Tidak Valid")
		uiawal()
	}
}

func daftar() {
	var tuid, tpwd string
	var test bool
	test = true
	fmt.Println("Silahkan input Username dan Password")
	fmt.Print("Username:  ")
	fmt.Scan(&tuid)
	fmt.Println()
	fmt.Print("Password:  ")
	fmt.Scan(&tpwd)
	if val(tuid) < 0 {
		fmt.Printf("Maaf Username Anda Tidak Valid \nMasukan Minimal 1 Huruf\n")
		return
	}
	for i := 0; i < 1024 && test; i++ {
		if tuid == login[i].uid {
			fmt.Println("User Sudah ada")
			test = false
			
		}

	}

	if test {
		login[accN].uid = tuid
		login[accN].pwd = tpwd
		fmt.Println("Pendaftaran Berhasil")
		accN++
	}
}

func masuk() {
	var tuid, tpwd string
	var test bool
	test = true
	fmt.Println("Silahkan Masukan Username dan Password Anda")
	fmt.Print("Username:  ")
	fmt.Scan(&tuid)
	fmt.Println()
	fmt.Print("Password:  ")
	fmt.Scan(&tpwd)
	for i := 0; i < 1024 && test; i++ {
		if tuid == login[i].uid && tpwd == login[i].pwd {
			test = false
			cuser = tuid
			MainMenu()
			
		}

	}
	if test == true {
		fmt.Println("Maaf Username atau Password Anda Salah")
	}
}

func MainMenu() {

	var run bool
	run = true
	for run {
		var a int
		fmt.Println("=================")
		fmt.Println("      MENU")
		fmt.Println("=================")
		fmt.Printf("1.Lihat Postingan\n2.Lihat Profil\n3.Buat Postingan\n4.Berikan Komentar\n5.Tambahkan teman\n6.Hapus teman\n7.Cetak teman ascending\n8.Cetak teman descending\n9.Mencari nama teman\n10.Logout\n")
		fmt.Print("Pilih (1-10) :")
		fmt.Scan(&a)
		switch a {
		case 1:
			jabarpost()
		case 2:
			currentuser()
		case 3:
			buatpost()
		case 4:
			komen()
		case 5:
			frn()
		case 6:
			delfrn()
		case 7:
			frnlst()
		case 8:
			frnlsd()
		case 9:
			SearchFriendlist()
		case 10:
			return
		default:
			fmt.Print("Pilihan Anda Tidak Valid")
			MainMenu()
			return
		}
	}
}

func currentuser() {
	var id int = index(cuser)
	var a string
	if id != -1 {
		var bio string = login[id].bio
		if bio == "" {
			bio = "Bio Masih Kosong"
		}
		fmt.Printf("Username anda :%s \nBio anda: %s\n", login[id].uid, bio)
		fmt.Print("Edit Bio Anda y/n? ")
		fmt.Scan(&a)
		if a == "y" {
			fmt.Print("Ketik Bio Baru: ")
			fmt.Scan(&login[id].bio)
		}
		fmt.Println("Bio Anda Telah Diperbaharui")
	}
}

func index(uid string) int {
	var index int = -1
	for i := 0; i <= accN; i++ {
		if uid == login[i].uid {
			index = i
		}
	}
	return index
}

func jabarpost() {
	if postN == 0 {
		fmt.Println()
		fmt.Println("Anda Belum Memiliki Postingan/Belum Ada  Postingan")
		fmt.Println()
	}
	for i := 0; i < postN; i++ {
		fmt.Printf("%d. Diposting Oleh : %s\n", i+1, post[i].op)
		fmt.Println(post[i].main)
		for j := 0; j < post[i].commN; j++ {
			fmt.Println("Dikomentari oleh :", post[i].comm[j].op)
			fmt.Println(post[i].comm[j].main)
		}
	}
}

func buatpost() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Membuat Postingan \n")
	fmt.Println("Masukkan Postingan:")
	conten, _ := reader.ReadString('\n')
	content, _ := reader.ReadString('\n')
	// Trim spaces from the beginning and end of the input
	content = strings.TrimSpace(content)
	post[postN].main = conten
	post[postN].main = content
	post[postN].op = cuser
	postN++
	fmt.Println("Anda Telah Berhasil Membuat Postingan")

}
func komen() {
	reader := bufio.NewReader(os.Stdin)
	var a int
//	var txt string
	if postN > 0 {
		jabarpost()
		fmt.Println("Postingan Keberapa Yang Ingin Anda Beri Komentar?")
		fmt.Scan(&a)
	//	fmt.Println("Apa komentarmu?")
	//	fmt.Scan(&txt)
	fmt.Print("Membuat Komentar \n")
	fmt.Println("Masukkan komentar:")
	conten, _ := reader.ReadString('\n')
	content, _ := reader.ReadString('\n')
		fmt.Println("Komentar Berhasil Ditambahkan")
		post[a-1].comm[post[a-1].commN].main = conten
		post[a-1].comm[post[a-1].commN].main = content
		post[a-1].comm[post[a-1].commN].op = cuser
		post[a-1].commN++
	}
}

func frn() {
	var id int = index(cuser)
	var tmp string
	var test bool = true
	fmt.Println("Ketik Username Teman Anda")
	fmt.Scan(&tmp)
	for i := 0; i <= accN; i++ {
		if tmp == login[i].uid && duplikat(tmp, id) {
			test = false
			fmt.Println("teman berhasil ditambah")
			login[id].teman[login[id].temanN] = tmp
			login[id].temanN++
		}
	}
	if test {
		fmt.Println("maaf user tidak ditemukan atau sudah menjadi teman anda")
	}
}

func frnlst() {
	var id int = index(cuser)
	for i := 0; i < login[id].temanN; i++ {
		for j := i; j > 0 && val(login[id].teman[j-1]) > val(login[id].teman[j]); j-- {
			login[id].teman[j], login[id].teman[j-1] = login[id].teman[j-1], login[id].teman[j]
		}

	}

	if login[id].temanN == 0 {
		fmt.Println("Anda belum memiliki teman")
	} else {
		fmt.Println("daftar teman anda")
		for n := 0; n < login[id].temanN; n++ {
			fmt.Println(n+1, login[id].teman[n])
		}
	}
}

func frnlsd() {
	// unutk dirubah menjadi selection sort
	var id int = index(cuser)
//	for i := 0; i < login[id].temanN; i++ {
//		for j := i; j > 0 && val(login[id].teman[j-1]) < val(login[id].teman[j]); j-- {
//			login[id].teman[j], login[id].teman[j-1] = login[id].teman[j-1], login[id].teman[j]
//		}
//
//	}

/* sudah jadi selection sort */
	for i:=0;i<login[id].temanN;i++ {
		min := i
		for j:=i ; j<login[id].temanN;j++{
			if val(login[id].teman[min]) < val(login[id].teman[j]) {
				min = j
			}
		}
		temp := login[id].teman[min]
		login[id].teman[min] = login[id].teman[i]
		login[id].teman[i] = temp
	}

	if login[id].temanN == 0 {
		fmt.Println("Anda belum memiliki teman")
	} else {
		fmt.Println("daftar teman anda")
		for n := 0; n < login[id].temanN; n++ {
			fmt.Println(n+1, login[id].teman[n])
		}
	}
}

func duplikat(x string, uid int) bool {
	for i := 0; i <= accN; i++ {
		if x == login[uid].teman[i] {
			return false
		}
	}
	return true
}

func delfrn() {
	var id int = index(cuser)
	var tmp string
	var test bool = true
	var pos int
	var l,r,m int 
	pos = -1
for i := 0; i < login[id].temanN; i++ {
		for j := i; j > 0 && val(login[id].teman[j-1]) > val(login[id].teman[j]); j-- {
			login[id].teman[j], login[id].teman[j-1] = login[id].teman[j-1], login[id].teman[j]
		}

	}

	l = 0
	r = login[id].temanN 

	if login[id].temanN == 0 {
		fmt.Println("Anda belum memiliki teman")
	} else {
		fmt.Println("input teman yang ingin dihapus")
		fmt.Scan(&tmp)

	for l <= r && test == true{

/*	fmt.Println("", m,pos)
	fmt.Println(login[id].teman[m])
*/
		m = (l + r) / 2
		if val(login[id].teman[m]) == val(tmp) {
			pos = m
			test = false
			//break
			
		} else if val(login[id].teman[m]) > val(tmp){
			r = m-1
		} else if val(login[id].teman[m]) < val(tmp){
			l = m+1
		} 
		//else {
		//	break
		//}

	}

/*		for i := 0; i < login[id].temanN; i++ {
			if tmp == login[id].teman[i] {
				pos = i
				test = false
				login[id].temanN--
			}
		}

*/

		if pos != -1 {
		login[id].temanN--
		for j := pos; j <= login[id].temanN; j++ {
			login[id].teman[j] = login[id].teman[j+1]

		}
	}
			if test==true {
			fmt.Println("Maaf data tidak ditemukan")
		}
	}
}

func val(a string) int {
	var val int
	var b string
	for i := len(a) - 1; i >= 0; i-- {
		b += string(a[i])
	}
	for ind, char := range b {
		val += pang(ind) * (int(char) - 64)
	}
	return val
}

func pang(n int) int {
	var x int = 1
	for i := 0; i < n; i++ {
		x *= 100
	}
	return x
}

func SearchFriendlist() {
	var id int = index(cuser)
	var searchTerm string
	fmt.Println("Masukkan Username Teman yang Ingin Dicari:")
	fmt.Scan(&searchTerm)

	var found bool = false
	for i := 0; i < login[id].temanN; i++ {
		if strings.Contains(login[id].teman[i], searchTerm) {
			if !found {
				fmt.Println("Teman yang ditemukan:")
				found = true
			}
			fmt.Println(login[id].teman[i])
		}
	}

	if !found {
		fmt.Println("Tidak ada teman yang ditemukan dengan kata kunci tersebut.")
	}
}
