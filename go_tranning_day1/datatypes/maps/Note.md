- The zero value of a map is nil. If you try to add elements to a nil map, a run-time paiic will occur. Hence the map has to be initialize before adding elements.
- Map are reference types: when a map is assigned to a new variable, they both poin to the same internal data structure. Hence changes made in one will reflect in other
- When maps are passed as parameters to function, any change is made to the map inside the function, it will be visible to the caller also
- Maps can't compared using the == operator
