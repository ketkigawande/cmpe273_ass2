package main

import (  
    "encoding/json"
    "fmt"
    "net/http"
    "github.com/julienschmidt/httprouter"
    "io/ioutil"
    "strings"
    "os"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    "log"
    "strconv"
    

)
const(
  Address = "ds045054.mongolab.com:45054"
  Database ="location"
  Username ="ketkigawande"
  Password ="17@Ketki"
)

var count int



type Request struct {
       Name string `json:"name"`
       Address string `json:"address"`
       City string `json:"city"`
       State string `json:"state"`
       Zip string `json:"zip"`
}

type Request1 struct {
       Address string `json:"address"`
       City string `json:"city"`
       State string `json:"state"`
       Zip string `json:"zip"`
}

type Response struct {
    
    Id string `json:"id"`
    Name string `json:"name"`
    Address string `json:"address"`
    City string `json:"city"`
    State string `json:"state"`
    Zip string `json:"zip"`
       Coordinate struct {
           Lat string `json:"lat"`
           Lng string `json:"lng"`
         } `json:"coordinate"`
   }  


type Response1 struct {
    Statuscode string `json:"statuscode"`

}
func Get(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
       reply:=Response{}
       id1:=p.ByName("id")
        /*abc:=&mgo.DialInfo {
        Address:[]string{Address},
        Database:Database,
        Username:Username,
        Password:Password,
      }*/
      abc:=&mgo.DialInfo{
        Addrs:[]string{Address},
        Database:Database,
        Username:Username,
        Password:Password,
       }
       session, err := mgo.DialWithInfo(abc)
       //session, err := mgo.Dial("127.0.0.1")
       if err != nil {
               panic(err)
       }
       defer session.Close()

       // Optional. Switch the session to a monotonic behavior.
       session.SetMode(mgo.Monotonic, true)

       c := session.DB("location").C("Details")
       fmt.Println(c)

       err = c.Find(bson.M{"id": id1}).One(&reply)
       if err != nil {
       panic(err)
       }

      fmt.Println(reply.Coordinate.Lat)
      //fmt.Println(reply.Lng)

      var res Response
      res.Id=id1
      res.Name=reply.Name
      res.Address=reply.Address
      res.City=reply.City
      res.State=reply.State
      res.Zip=reply.Zip
      res.Coordinate.Lat=reply.Coordinate.Lat
      res.Coordinate.Lng=reply.Coordinate.Lng

     //fmt.Println("Id:", id1)
      uj, _ := json.Marshal(res)
      rw.Header().Set("Content-Type", "application/json")
      rw.WriteHeader(200)
      fmt.Fprintf(rw, "\n\n\n\n%s", uj)
     //fmt.Fprintf(rw, "Hello, %s!\n", p.ByName("name"))
   }

func put(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
      request1:= Request1{}
      reply:=Response{}

      json.NewDecoder(r.Body).Decode(&request1)       
       id1:=p.ByName("id")
       abc:=&mgo.DialInfo{
        Addrs:[]string{Address},
        Database:Database,
        Username:Username,
        Password:Password,
       }
       session, err := mgo.DialWithInfo(abc)
       //session, err := mgo.Dial("mongodb://ketkigawande:17@Ketki@ds045054.mongolab.com:45054/location")
       if err != nil {
               panic(err)
       }
       defer session.Close()

       // Optional. Switch the session to a monotonic behavior.
       session.SetMode(mgo.Monotonic, true)

       c := session.DB("location").C("Details")
       fmt.Println(c)
       forquery := bson.M{"id": id1}
       change := bson.M{"$set": bson.M{"address": request1.Address, "city": request1.City, "state": request1.State, "zip": request1.Zip}}
       err = c.Update(forquery, change)
       if err != nil {
       panic(err)
       }

       err = c.Find(bson.M{"id": id1}).One(&reply)
       if err != nil {
       panic(err)
       }

      fmt.Println(reply.Coordinate.Lat)
      //fmt.Println(reply.Lng)

      var res Response
      res.Id=id1
      res.Name=reply.Name
      res.Address=reply.Address
      res.City=reply.City
      res.State=reply.State
      res.Zip=reply.Zip
      res.Coordinate.Lat=reply.Coordinate.Lat
      res.Coordinate.Lng=reply.Coordinate.Lng

     //fmt.Println("Id:", id1)
      uj, _ := json.Marshal(res)
      w.Header().Set("Content-Type", "application/json")
      w.WriteHeader(201)
      fmt.Fprintf(w, "\n\n\n\n%s", uj)
     //fmt.Fprintf(rw, "Hello, %s!\n", p.ByName("name"))
      
  }


func delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
      
       id1:=p.ByName("id")
       //response1 := Response1{}
       //session, err := mgo.Dial("mongodb://ketkigawande:17@Ketki@ds045054.mongolab.com:45054/location")
       abc:=&mgo.DialInfo{
        Addrs:[]string{Address},
        Database:Database,
        Username:Username,
        Password:Password,
       }
       session, err := mgo.DialWithInfo(abc)
       if err != nil {
               panic(err)
       }
       defer session.Close()

       // Optional. Switch the session to a monotonic behavior.
       session.SetMode(mgo.Monotonic, true)
       c := session.DB("location").C("Details")
       //fmt.Println(c)
       err = c.Remove(bson.M{"id": id1})
       if err != nil {
       panic(err)
       }
    //response1.Statuscode="200"

    //uj, _:= json.Marshal(response1)
    //w.Header().Set("Content-Type", "application/json")
    //w.WriteHeader(200)
    //fmt.Fprintf(w, "\n\n\t%s", uj)

}


func postt(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
   
    request:= Request{}
    response := Response{}
    json.NewDecoder(r.Body).Decode(&request)
    response.Address=request.Address
    response.City=request.City
    response.State=request.State
    request.Address=strings.Replace(request.Address, " ", "+", -1)
    request.City=strings.Replace(request.City," ","+",-1)
    request.State=strings.Replace(request.State," ","+",-1)
    

    a :=fmt.Sprint("http://maps.google.com/maps/api/geocode/json?address="+request.Address+",+"+request.City+",+"+request.State+"&sensor=false")
    fmt.Println("%s",a)
    response1, err := http.Get(a)
    if err != nil {
        fmt.Printf("%s", err)
        os.Exit(1)
    } else {
        defer response1.Body.Close()
        contents, err := ioutil.ReadAll(response1.Body)
        if err != nil {
            fmt.Printf("%s", err)
            os.Exit(1)
        }
      var f interface{}
      err=json.Unmarshal(contents, &f)
      mRes := f.(map[string]interface{})["results"]
      mRes0 := mRes.([]interface{})[0]
      mGeo := mRes0.(map[string]interface{})["geometry"]
      mLoc := mGeo.(map[string]interface{})["location"]
      locLat := mLoc.(map[string]interface{})["lat"].(float64)
      locLng := mLoc.(map[string]interface{})["lng"].(float64)
      //response.Statuscode = "201"
      response.Name=request.Name
      
      response.Zip=request.Zip
      response.Coordinate.Lat=strconv.FormatFloat(locLat, 'f', 6, 64)
      response.Coordinate.Lng=strconv.FormatFloat(locLng, 'f', 6, 64)
      response.Id=strconv.Itoa(count)
      count=count+5
      abc:=&mgo.DialInfo{
        Addrs:[]string{Address},
        Database:Database,
        Username:Username,
        Password:Password,
       }
       session, err := mgo.DialWithInfo(abc)
      //session, err := mgo.Dial("mongodb://ketkigawande:17@Ketki@ds045054.mongolab.com:45054/location")
       if err != nil {
               panic(err)
       }
       defer session.Close()

       // Optional. Switch the session to a monotonic behavior.
       session.SetMode(mgo.Monotonic, true)

       c := session.DB("location").C("Details")
       err = c.Insert(response)
       if err != nil {
               log.Fatal(err)
       }

      

    // Marshal provided interface into JSON structure
    uj, _ := json.Marshal(response)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(201)
    fmt.Fprintf(w, "\n\n\n\n%s", uj)
  }
}



func main() {
    count=7  
    r := httprouter.New()

    server := http.Server{
            Addr:        "0.0.0.0:8080",
            Handler: r,
    }
     r.POST("/locations",postt)
     r.GET("/locations/:id",Get)
     r.PUT("/locations/:id",put)
     r.DELETE("/locations/:id",delete)
    server.ListenAndServe()
    
 }
