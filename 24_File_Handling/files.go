package main

import (
	"fmt"
	"os"
)

func main() {
	// f, err := os.Open("example.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// fileInfo, err := f.Stat()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("file name:", fileInfo.Name())
	// fmt.Println("file or folder:", fileInfo.IsDir())
	// fmt.Println("file size:", fileInfo.Size())
	// fmt.Println("file permissions:", fileInfo.Mode())
	// fmt.Println("file modification:", fileInfo.ModTime())

	//read file
	//  f, err := os.Open("example.txt")
	//  if err != nil {
	// 	 panic(err)
	//  }

	//  defer f.Close() // Ensure the file is closed after we're done

	//  buff := make([]byte, 10)
	//  d, err := f.Read(buff)
	//  if err != nil {
	// 	 panic(err)
	//  }
	//  for i:= 0; i<len(buff); i++ {
	// 	fmt.Println("data:", d, string(buff[i]))
	//  }

	// read file using ReadFile will read the entire file at once so not suitable for large files
	// data, err := os.ReadFile("example.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("data:", string(data))

	// read folders
	// dir, err := os.Open(".")
	// if err != nil {
	// 	panic(err)
	// }
	// defer dir.Close() // Ensure the directory is closed after we're done
	// fileInfo, err := dir.Readdir(0) // 0 means read all files
	// for _, fi := range fileInfo {
	// 	fmt.Println(fi.Name(), fi.IsDir(), fi.Size(), fi.Mode(), fi.ModTime())
	// }

	// create a file
	// f, err := os.Create("example2.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer f.Close()

	// f.WriteString("Hi, Golang!\n")
	// f.WriteString("This is a new file created using os.Create.\n")

	// Another way of writing to a file
	// bytes := []byte("Hello World!")
	// f.Write(bytes)

	// Read from and copy to another file (stream fashion)
	// sourceFile, err := os.Open("example.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer sourceFile.Close()

	// destFile, err := os.Create("example_copy.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer destFile.Close()

	// reader := bufio.NewReader(sourceFile)
	// writer := bufio.NewWriter(destFile)

	// for {
	// 	b, err := reader.ReadByte()
	// 	if err != nil {
	// 		if err.Error() == "EOF" {
	// 			break // End of file reached
	// 		}
	// 		panic(err) // Handle other errors
	// 	}

	// 	e := writer.WriteByte(b)
	// 	if e != nil {
	// 		panic(e) // Handle write error
	// 	}
	// }
	// writer.Flush() // Ensure all buffered data is written to the file
	// fmt.Println("File copied successfully!")

	// delete a file
	err := os.Remove("example_copy.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("File deleted successfully!")
}
