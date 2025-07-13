func matchPlayersAndTrainers(players []int, trainers []int) int {
    res := 0
	i,j := 0, 0;
    
      sort.Ints(players)
      sort.Ints(trainers)
    for i < len(players) && j< len(trainers) {
          if players[i] <= trainers[j] {
               res+= 1;
               i++;
               j++
          }else{
              j++;
          }
    }
      return res;
}