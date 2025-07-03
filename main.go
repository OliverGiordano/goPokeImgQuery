package main
import (
    "io"
    "os"
    "image"
    "image/png"
    "io/ioutil"
    "net/http"
    "math/rand"
    "fmt"
    "strconv"
    "encoding/json"
    "sync"
)

type SubImager interface {
	SubImage(r image.Rectangle) image.Image
}

func main(){
    pokeNum := rand.Intn(1025)
    if pokeNum == 0{
	pokeNum++
    }
    url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", strconv.Itoa(pokeNum))
    imgUrl := fmt.Sprintf("https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/%s.png", strconv.Itoa(pokeNum))
    var wg sync.WaitGroup
    wg.Add(2)
    go getPokeName(url, &wg)
    go getPokeImg(imgUrl, &wg)
    wg.Wait()
}
func getPokeName(url string, wg *sync.WaitGroup) {
    defer wg.Done()
    resp, _ := http.Get(url)
    body, _ := ioutil.ReadAll(resp.Body)
    var pBody pokeStruct
    _ = json.Unmarshal(body, &pBody)
    file, _ := os.Create("pokeName.txt")
    defer file.Close()
    dressUp := fmt.Sprintf("	-------------%s-------------	\n", pBody.Name)
    file.WriteString(dressUp)

}

func getPokeImg(url string, wg *sync.WaitGroup){
    defer wg.Done()
    imgPull, err := http.Get(url)
    if err != nil {
	panic(err)
    }
    defer imgPull.Body.Close()

    img, err := os.Create("poke.png") 
    defer img.Close()
    if err != nil {
	panic(err)
    }
    _, _ = io.Copy(img, imgPull.Body)
    img, err = os.Open("poke.png") //I need to re-open file so last changes are saved
    cropImg(img)
}


func cropImg(img *os.File){
    decodedImg, err := png.Decode(img)
    if err != nil {
	panic(err)
    }
    bounds := decodedImg.Bounds()
    width := bounds.Dx()
    cropSize := image.Rect(0, 0, width/2+20, width/2+20)
    cropSize = cropSize.Add(image.Point{18, 14})
    cImg := decodedImg.(SubImager).SubImage(cropSize)
    nImg, err := os.Create("CPoke.png")
    if err != nil{
	panic(err)
    }
    err = png.Encode(nImg, cImg)
    if err != nil{
	panic(err)
    }
}
