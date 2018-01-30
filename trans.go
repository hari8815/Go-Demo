
package main

import (
	"gopkg.in/gographics/imagick.v2/imagick"
	"strconv"
	"fmt"
)
var imgName,finalImg string
var pagNo int

func main() {
	
	//Intialize ImageMagick
	imagick.Initialize()
    defer imagick.Terminate()
    mw := imagick.NewMagickWand()
    defer mw.Destroy()
    
    //Read PDF file
    mw.ReadImage("example1.pdf")
    
	pagNo := int(mw.GetNumberImages())

    if( pagNo > 0 ){
	    //Convert PDF into Images
	    for i := 0; i < pagNo; i++ {
		    mw.SetIteratorIndex(i)        
		    mw.SetImageFormat("jpg")

		    imgName := "./Exported/Exported-image-"+strconv.Itoa(i)+".jpg"
		    mw.WriteImage(imgName)
		}
		fmt.Println("Pdf to Image Export Done!")

		//Initialize trim background
		target := imagick.NewPixelWand()
		target.SetColor("white")

		for j := 0; j < pagNo; j++ {
			imgName := "./Exported/Exported-image-"+strconv.Itoa(j)+".jpg"
			mw.ReadImage(imgName)
		    
		    mw.TrimImage(10)

		    finalImg := "./Final/Column-image-"+strconv.Itoa(j)+".jpg"
		    mw.WriteImage(finalImg)
		}

		fmt.Println("Final Image Trim Done!")
	}else{
		fmt.Println("Invalid PDF File!")
	}
}
