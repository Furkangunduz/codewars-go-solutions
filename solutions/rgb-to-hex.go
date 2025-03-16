package solutions

import (
	"fmt"
	"strings"
)

func RGB(r, g, b int) string {
  if r < 0 {
    r=0
  }else if r>255{
    r=255
  }
   if g < 0 {
    g=0
  }else if g>255{
    g=255
  }
   if b<0{
    b=0
  }else if b>255{
    b=255
  }
  return strings.ToUpper(fmt.Sprintf("%02x%02x%02x",r,g,b))
}