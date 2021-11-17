package main

import(
//"os"
"fmt"
"io/ioutil"
"io/fs"
"strings"
"os"
)

type file string

func main(){

    /*
    getFileNames
    for name in filenames{
        if nameContainsCharacters(){
            rename file
        }
   */ 



   // origPath := "test.txt"
   // newPath := "test1.txt"
   // err := os.Rename(origPath, newPath)
   // if err != nil {
   //     log.Fatal(err)
   // }
   var files []fs.FileInfo = GetFileNames()
   for _, fs := range(files) {
       fmt.Println(fs.Name())
       if HasBadCharacter(fs.Name()){
           RenameFile(fs.Name())
       }
   }
}


func GetFileNames() []fs.FileInfo{
   files, err := ioutil.ReadDir(".")
   if err != nil {
       panic(err)
   }
   return files
}

func HasBadCharacter(filename string) bool{
    fmt.Println(strings.Contains(filename, "_") || strings.Contains(filename, " ") || hasUpperCase(filename))
    fmt.Println()
    return strings.Contains(filename, "-") || strings.Contains(filename, " ") || hasUpperCase(filename)


}

func hasUpperCase(filename string) bool {
    return filename != strings.ToLower(filename)
 }

func RenameFile(oldName string){
    fmt.Println("OLDNAME: ", oldName)
    fmt.Println("renaming")
    newName  := strings.ToLower(oldName)
    newName = strings.ReplaceAll(newName, "_", "-")
    newName = strings.ReplaceAll(newName, " ", "-")
    os.Rename(oldName, newName)
}

/*
func RenameFile(oldName string){
    newName := []byte{}
    for _, v := range oldName {
        if v == 32 || (v > 64 && v < 91){
            v -= 32
            newName = append(newName,  byte(v))
            continue
        }
            newName = append(newName,  byte(v))
    }
    fmt.Println(string(newName))
    os.Rename(oldName, string(newName))
}
*/
