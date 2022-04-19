/*
 * Copyright (c) 2022 XXIV
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */
package main

import "C"
import (
	"github.com/fatih/color"
)

//export _Black
func _Black(format *C.char) {
  color.Black(C.GoString(format))
}

//export _Red
func _Red(format *C.char) {
  color.Red(C.GoString(format))
}

//export _Green
func _Green(format *C.char) {
  color.Green(C.GoString(format))
}

//export _Yellow
func _Yellow(format *C.char) {
  color.Yellow(C.GoString(format))
}

//export _Blue
func _Blue(format *C.char) {
  color.Blue(C.GoString(format))
}

//export _Magenta
func _Magenta(format *C.char) {
  color.Magenta(C.GoString(format))
}

//export _Cyan
func _Cyan(format *C.char) {
  color.Cyan(C.GoString(format))
}

//export _White
func _White(format *C.char) {
  color.White(C.GoString(format))
}

//export _BlackString
func _BlackString(format *C.char) *C.char {
  return C.CString(color.BlackString(C.GoString(format)))
}

//export _RedString
func _RedString(format *C.char) *C.char {
  return C.CString(color.RedString(C.GoString(format)))
}

//export _GreenString
func _GreenString(format *C.char) *C.char {
  return C.CString(color.GreenString(C.GoString(format)))
}

//export _YellowString
func _YellowString(format *C.char) *C.char {
  return C.CString(color.YellowString(C.GoString(format)))
}

//export _BlueString
func _BlueString(format *C.char) *C.char {
  return C.CString(color.BlueString(C.GoString(format)))
}

//export _MagentaString
func _MagentaString(format *C.char) *C.char {
  return C.CString(color.MagentaString(C.GoString(format)))
}

//export _CyanString
func _CyanString(format *C.char) *C.char {
  return C.CString(color.CyanString(C.GoString(format)))
}

//export _WhiteString
func _WhiteString(format *C.char) *C.char {
  return C.CString(color.WhiteString(C.GoString(format)))
}

//export _HiBlack
func _HiBlack(format *C.char) {
  color.HiBlack(C.GoString(format))
}

//export _HiRed
func _HiRed(format *C.char) {
  color.HiRed(C.GoString(format))
}

//export _HiGreen
func _HiGreen(format *C.char) {
  color.HiGreen(C.GoString(format))
}

//export _HiYellow
func _HiYellow(format *C.char) {
  color.HiYellow(C.GoString(format))
}

//export _HiBlue
func _HiBlue(format *C.char) {
  color.HiBlue(C.GoString(format))
}

//export _HiMagenta
func _HiMagenta(format *C.char) {
  color.HiMagenta(C.GoString(format))
}

//export _HiCyan
func _HiCyan(format *C.char) {
  color.HiCyan(C.GoString(format))
}

//export _HiWhite
func _HiWhite(format *C.char) {
  color.HiWhite(C.GoString(format))
}

//export _HiBlackString
func _HiBlackString(format *C.char) *C.char {
  return C.CString(color.HiBlackString(C.GoString(format)))
}

//export _HiRedString
func _HiRedString(format *C.char) *C.char {
  return C.CString(color.HiRedString(C.GoString(format)))
}

//export _HiGreenString
func _HiGreenString(format *C.char) *C.char {
  return C.CString(color.HiGreenString(C.GoString(format)))
}

//export _HiYellowString
func _HiYellowString(format *C.char) *C.char {
  return C.CString(color.HiYellowString(C.GoString(format)))
}

//export _HiBlueString
func _HiBlueString(format *C.char) *C.char {
  return C.CString(color.HiBlueString(C.GoString(format)))
}

//export _HiMagentaString
func _HiMagentaString(format *C.char) *C.char {
  return C.CString(color.HiMagentaString(C.GoString(format)))
}

//export _HiCyanString
func _HiCyanString(format *C.char) *C.char {
  return C.CString(color.HiCyanString(C.GoString(format)))
}

//export _HiWhiteString
func _HiWhiteString(format *C.char) *C.char {
  return C.CString(color.HiWhiteString(C.GoString(format)))
}

func main() {}