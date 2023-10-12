package main

// "Z" (UTC+0)
//#This is our datetime stored in DB
//2019-10-12T07:20:50.52Z
//#This is how it should be converted to UTC+7
//2019-10-12T14:20:50.52+07:00

//“Z”: stands for Zero timezone (UTC+0). Or equal to +00:00 in the RFC 3339.
//RFC 3339 is following the ISO 8601 DateTime format. The only difference is RFC allows us to replace “T” with “space”.

func main() {

}