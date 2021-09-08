package main

import (
	"fmt"
)

type stats struct{
   level int
   endurance, health int
}

func levelUp(s *stats){
   s.level++
   s.endurance = 42 + (14 * s.level)
   s.health = 5 * s.endurance
}

type character struct{
   name string
   // stats だけでもok
   stats *stats
}


func main(){
   player := character{name: "Mathias", stats: &stats{}}
   levelUp(player.stats)
   fmt.Printf("%+v\n", player.stats)
   fmt.Printf("%+v\n", player.stats.health)
}

// &{level:1 endurance:56 health:280}
// 280
