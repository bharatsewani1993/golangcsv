package main

import(
  "fmt"
  "encoding/csv"
  "encoding/json"
  "log"
  "os"
  "io"
  "bufio"
)

var data = [][]string{{"Bharat", "Sewani", "Bharatsewani1993@gmail.com"}, {"Hitesh", "Sewani", "Hiteshsewani.com"}}

type Userinfo struct{
   Firstname string `json:"firstname"`
   Lastname string `json:"lastname"`
   Creds *Creds `json:"creds,omitempty"`
}

type Creds struct{
  Username string `json:"username"`
  Password string `json:"password"`
  Email string `json:"email"`
}

//function to readfile
func readfile()  {
    csvfile, _ := os.Open("userinfo.csv")
    reader := csv.NewReader(bufio.NewReader(csvfile))
    var users []Userinfo
    //infinite for loop
    for {
       line, error := reader.Read()
       //if end of file
       if error == io.EOF {
         break
       } else if error != nil{
         // if any other error
         log.Fatal(error)
       }
         //process file info
          users = append(users, Userinfo{
            Firstname: line[0],
            Lastname: line[1],
            Creds: &Creds{
              Username: line[2],
              Password: line[3],
              Email: line[4],
            },
          })
    }
    userinfojson, _ := json.Marshal(users)
    fmt.Println(string(userinfojson))
}

//function to write file
func writefile(){
   file, err := os.Create("result.csv")
   if err != nil{
     log.Fatal(err)
   }
   defer file.Close()

   writer := csv.NewWriter(file)
   defer writer.Flush()

    for _, value := range data{
      err := writer.Write(value)
      if err!= nil{
        log.Fatal(err)
      }
    }
}

func main(){
    readfile()
    writefile()
}
