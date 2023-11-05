package main

import "fmt"

func main() {

	var bakiye = 5000

	fmt.Println("ATM ye hoşgeldniz. Lütfen yapmak istediğiniz işlemi seçiniz: 1 - Para Yatırma / 2 - Para Çekme / 3 - Bakiye Görüntüleme")

	for i := 0; i < 20; i++ {

		fmt.Println("Lütfen yapmak istediğiniz işlemi seçiniz: 1 - Para Yatırma / 2 - Para Çekme / 3 - Bakiye Görüntüleme")

		var yapilacakIslem int

		fmt.Scanln(&yapilacakIslem)

		if yapilacakIslem == 1 {
			fmt.Println("Lütfen kaç lira yatırmak istediğiniz giriniz:")

			var yatirilacakTutar int

			fmt.Scanln(&yatirilacakTutar)

			bakiye = bakiye + yatirilacakTutar

			fmt.Println("Paranız hesabınıza başarıyla yatırıldı. Güncel hesap bakiyeniz: ", bakiye)

		} else if yapilacakIslem == 2 {
			fmt.Println("lütfen kaç lira çekmek istediğinizi giriniz:")

			var cekilcektutar int

			fmt.Scanln(&cekilcektutar)

			if cekilcektutar > bakiye {
				fmt.Println("Parasız pust. Git para yükle")
			} else {
				bakiye = bakiye - cekilcektutar

				fmt.Println("Çekim işlemi başarıyla tamamlandı. Çekilen Tutar: ", cekilcektutar, "Güncel Bakiye: ", bakiye)
			}

		} else if yapilacakIslem == 3 {

			fmt.Println("Güncel Bakiyeniz:", bakiye)

		} else {
			fmt.Println("Lütfen geçerli bir işlem giriniz")
		}
	}

}
