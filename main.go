package main

import (
    "fmt"
    "os"
    "log"
    "strconv"
    "io/ioutil"
)

func main() {
    if len(os.Args) < 4 {
        fmt.Println( `Usage: binary_splitter split  <file name> <number of output files>
     or binary_splitter rejoin <output name> <input files>`)
        panic(1)
    }

    operation := os.Args[1]
    switch operation {
    case "split":
        fileName := os.Args[2]
        outputCount, err := strconv.Atoi(os.Args[3])
        if err != nil {
            log.Fatal(err)
        }
        splitFile(fileName, outputCount)
    case "rejoin":
        outputName := os.Args[2]
        fileNames := os.Args[3:]
        joinFiles(outputName, fileNames)
    default:
        fmt.Println("Supported Operations: 'split', 'rejoin'")
    }
}

func joinFiles(outputName string, fileNames []string) {
    fmt.Println("Joining", fileNames, "into", outputName)
    os.Remove(outputName)
    output, err := os.OpenFile(outputName, os.O_APPEND | os.O_WRONLY | os.O_CREATE, 0600)
    if err != nil {
        log.Fatal(err)
    }
    defer output.Close()

    for _, fileName := range fileNames {
        fileBytes, err := ioutil.ReadFile(fileName)
        if err != nil {
            log.Fatal(err)
        }
        output.Write(fileBytes)
    }
}

func splitFile(fileName string, outputCount int) {
    fmt.Println("splitting", fileName, "into", outputCount, "files")
    bytes, err := ioutil.ReadFile(fileName)
    if err != nil {
        log.Fatal(err)
    }
    splitBytes(fileName, bytes, outputCount)
}

func splitBytes(fileName string, bs []byte, outputCount int) {
    bytesPerFile := len(bs) / outputCount
    byteIndex := 0
    fileNumber := 0
    for ; fileNumber < outputCount - 1; fileNumber++ {
        writeBytes("./" + fileName + fmt.Sprintf("%d", fileNumber), bs[byteIndex : byteIndex + bytesPerFile])
        byteIndex += bytesPerFile
    }
    writeBytes("./" + fileName + fmt.Sprintf("%d", fileNumber), bs[byteIndex:])
}

func writeBytes(fileName string, bs []byte) {
    err := os.WriteFile(
        fileName,
        bs,
        0666,
    )
    if err != nil {
        log.Fatal(err)
    }
}
