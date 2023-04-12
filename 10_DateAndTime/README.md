# Date and time.

Date and time handling in golang is provided by time package in golang...
Let's see how to work with Data and Time.

How to get the current date and time with timestamp in local and other timezones?
-   ## Overview

    Time or Date is represented in Go using time.Time struct. Time can be also be represented as a Unix Time (Also known as Epoch Time) – It is the number of seconds elapsed since 00:00:00 UTC on 1 January 1970. This time is also known as the Unix epoch.

-   ## Structure
    time.Time object is used to represent a specific point in time. The time.Time struct is as below

            type Time struct {
                // wall and ext encode the wall time seconds, wall time nanoseconds,
                // and optional monotonic clock reading in nanoseconds.
                wall uint64
                ext  int64
                //Location to represent timeZone
                // The nil location means UTC
                loc *Location
            }
    Every time.Time object has an associated location value which is used to determine the minute, hour, month, day and year corresponding to that time.

-   ## Create a new time
    -   ### Using time.Now()
    This function can be used to get the current local timestamp. The signature of the function is

            func Now() Time
    -   ### Using time.Date()
    This function returns the time which is yyyy-mm-dd hh:mm:ss + nsec nanoseconds with the appropriate time zone corresponding to the given location. The signature of the function is:

            func Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location) Time
-   ## Understanding Duration

duration is the time that has elapsed between two instants of time. It is represented as int64 nanosecond count. So duration is nothing in Go but just a number representing time in nanoseconds. So if duration value is  equal to 1000000000 then it represents 1 sec or 1000 milliseconds or 10000000000 nanoseconds

As an example duration between two time values 1 hour apart will be below value which is equal number of nanoseconds in 1 hour.

        1*60*60*1000*1000*1000

It is represented as below in the time package.

        type Duration int64

Below are some common duration which are defined in time package

            const (
                Nanosecond  Duration = 1
                Microsecond          = 1000 * Nanosecond
                Millisecond          = 1000 * Microsecond
                Second               = 1000 * Millisecond
                Minute               = 60 * Second
                Hour                 = 60 * Minute
            )
Some of the function defined on time.Time object that returns the Duration are

func (t Time) Sub(u Time) Duration – It returns the duration t-u
func Since(t Time) Duration – It returns the duration which has elapsed since t
func Until(t Time) Duration – It returns the duration until t

-   ## Add or Subtract to a time
    Now that you have understood what duration , let’s see how we can add or subtract to a time instance

    time package in golang defines two ways of adding or subtracting to a time.

    Add function – It is used to add/subtract a duration to time t. Since duration can be represented in hours, minutes, seconds, milliseconds, microseconds and nanoseconds, therefore Add function can be used to add/subtract hours, minutes, seconds, milliseconds, microseconds and nanoseconds from a time . Its signature is

            func (t Time) Add(d Duration) Time

    AddDate function – It is used to add/subtract years, months and days to time t. Its signature is

            func (t Time) AddDate(years int, months int, days int) Time

    Note: Positive values are used to add to time and negative values are used to subtract. Let’s see a working example of Add and Subtract to time.

    -   ### Add to time
            package main

            import (
                "fmt"
                "time"
            )

            func main() {
                t := time.Now()

                //Add 1 hours
                newT := t.Add(time.Hour * 1)
                fmt.Printf("Adding 1 hour\n: %s\n", newT)

                //Add 15 min
                newT = t.Add(time.Minute * 15)
                fmt.Printf("Adding 15 minute\n: %s\n", newT)

                //Add 10 sec
                newT = t.Add(time.Second * 10)
                fmt.Printf("Adding 10 sec\n: %s\n", newT)

                //Add 100 millisecond
                newT = t.Add(time.Millisecond * 10)
                fmt.Printf("Adding 100 millisecond\n: %s\n", newT)

                //Add 1000 microsecond
                newT = t.Add(time.Millisecond * 10)
                fmt.Printf("Adding 1000 microsecond\n: %s\n", newT)

                //Add 10000 nanosecond
                newT = t.Add(time.Nanosecond * 10000)
                fmt.Printf("Adding 1000 nanosecond\n: %s\n", newT)

                //Add 1 year 2 month 4 day
                newT = t.AddDate(1, 2, 4)
                fmt.Printf("Adding 1 year 2 month 4 day\n: %s\n", newT)
            }
    -   ### Subtract to time
            package main

            import (
                "fmt"
                "time"
            )

            func main() {
                t := time.Now()

                //Add 1 hours
                newT := t.Add(-time.Hour * 1)
                fmt.Printf("Subtracting 1 hour:\n %s\n", newT)

                //Add 15 min
                newT = t.Add(-time.Minute * 15)
                fmt.Printf("Subtracting 15 minute:\n %s\n", newT)

                //Add 10 sec
                newT = t.Add(-time.Second * 10)
                fmt.Printf("Subtracting 10 sec:\n %s\n", newT)

                //Add 100 millisecond
                newT = t.Add(-time.Millisecond * 10)
                fmt.Printf("Subtracting 100 millisecond:\n %s\n", newT)

                //Add 1000 microsecond
                newT = t.Add(-time.Millisecond * 10)
                fmt.Printf("Subtracting 1000 microsecond:\n %s\n", newT)

                //Add 10000 nanosecond
                newT = t.Add(-time.Nanosecond * 10000)
                fmt.Printf("Subtracting 1000 nanosecond:\n %s\n", newT)

                //Add 1 year 2 month 4 day
                newT = t.AddDate(-1, -2, -4)
                fmt.Printf("Subtracting 1 year 2 month 4 day:\n %s\n", newT)
            }
-   ## Time Parsing/Formatting
    Go uses date and time place holders 
    
    -   ### Time Parse Example
    -   ### Time Formatting Example
-   ## Time Diff
-   ## Time Conversion
    -   ### Convert time between different timezones

