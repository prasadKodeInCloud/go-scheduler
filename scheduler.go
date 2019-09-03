
package main

import(
	"fmt"
	//"go/build"
	"os"
	"github.com/robfig/cron"
	"net/http"
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
    log "github.com/sirupsen/logrus"
)

const LOG_FILE_PATH = "D:/GO_Projects/src/scheduler/logger/info.log"

func main() {
  fmt.Println("Hello, 世界")

  //logInfo("Started server...")
   // Inspect the cron job entries' next and previous run times.
   //inspect(c.Entries())

   // Echo instance
  e := echo.New()

  // Middleware
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  // Routes
  e.GET("/", hello)

    file, err := os.Create( LOG_FILE_PATH ) //os.OpenFile( LOG_FILE_PATH , os.O_CREATE|os.O_APPEND, 0644)
    if err != nil {
        log.Fatal(err)
    }

    defer file.Close()

    log.SetOutput(file)
    log.SetFormatter(&log.JSONFormatter{})
    log.SetLevel(log.InfoLevel)

    log.WithFields(log.Fields{
        "animal": "walrus",
        "size":   10,
    }).Info("A group of walrus emerges from the ocean")

    log.WithFields(log.Fields{
        "omg":    true,
        "number": 122,
    }).Warn("The group's number increased tremendously!")

    // log.WithFields(log.Fields{
    //     "omg":    true,
    //     "number": 100,
    // }).Fatal("The ice breaks!")

  // file, _ := os.Create( LOG_FILE_PATH )
  // defer file.Close()

  // n3, err := file.WriteString("writes\n")
  // check( err )
  // fmt.Printf("wrote %d bytes\n", n3)
  // //Issue a Sync to flush writes to stable storage.\
  // file.Sync()

  // Start server
  e.Logger.Fatal(e.Start(":1324"))

  c := cron.New()
  c.AddFunc("0 30 * * * *", func() { fmt.Println("Every hour on the half hour") })
  c.AddFunc("0 47 16 * * *", func() { 
    fmt.Println("Executed Scheduler")  
  	runScheduler( SchedulerInput{ "jhajsuyuqyw", "DAILY" } )
  
  })

  c.Start()
  
}

// Handler
func hello(c echo.Context) error {
  return c.String(http.StatusOK, "Hello, World!")
}

type SchedulerInput struct {
    //logFile *os.File
    cronId string
    schedulerType string
}

func runScheduler( e SchedulerInput ){
    fmt.Println("Run scheduler :", e.cronId ) 
    log.WithFields( log.Fields{
        "cronId":    e.cronId,
        "schedulerType": e.schedulerType,
    } ).Info("Executed Scheduler")
}
	
func check(e error) {
    if e != nil {
        panic(e)
    }
}

