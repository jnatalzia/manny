ps | grep "./[d]ist/main -t" | awk '
{ 
  if($1!="") {
      print "killing gogram: "$1
      system("kill " $1)
  }
}' 