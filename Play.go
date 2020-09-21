package main

import (
	"time"
	"fmt"
	"os"
	"bufio")

type Ticket struct{
	customerName string
	s Seat
	show Show}

type Theatre struct {
	seats []Seat
	shows []Show}

type Play struct {
	name string
	purchased []Ticket
	showsStart time.Time 
	showEnd time.Time}

type Comedy struct{
	Play
	laughs float32
	deaths int32}

type Tragedy struct {
	Play
	laughs float32
	deaths int32}

type Seat struct{
	number int32
	row int32
	cat Category}

type Category struct{
	name string
	basePrice float32}
	
type Show interface {
	getName() string
	getShowStart() time.Time
	getShowEnd() time.Time
	addPurchase(Ticket) (bool, Ticket)
	isNotPurchased(Ticket) (bool, Ticket)
}

func(a Comedy) getName() string {
	return "Tartuffe"
}

func(a Tragedy) getName() string {
	return "Macbeth"
}

func (a Comedy) getShowStart() time.Time{
	return a.showsStart}

func (a Tragedy) getShowStart() time.Time{
	return a.showsStart}

func (a Comedy) getShowEnd() time.Time{
	return a.showEnd}

func (a Tragedy) getShowEnd() time.Time{
	return a.showEnd}

func (a Comedy) addPurchase(tick Ticket) (bool,Ticket) {
	boo, tick := a.isNotPurchased(tick)
	if boo == true{
		a.purchased = append(a.purchased, tick)	//if not then add the ticket the to the slice
		return true,tick
	}
	return false, tick}

func (a Tragedy) addPurchase(tick Ticket) (bool, Ticket){
	boo, tick := a.isNotPurchased(tick)
	if boo == true{
		a.purchased = append(a.purchased, tick)	//if not then add the ticket the to the slice
		return true, tick
	}
	return false, tick}

func (a Comedy) isNotPurchased(tick Ticket) (bool, Ticket) {
	for i:= 0; i < len(a.purchased); i++ {	//check if the seat has been taken or not
		if a.purchased[i].s.number == tick.s.number {
			if a.purchased[i].s.row == tick.s.row{
				return false, tick
			}
		}
	}
	return true, tick}

func (a Tragedy) isNotPurchased(tick Ticket) (bool, Ticket) {
	for i:= 0; i < len(a.purchased); i++ {	//check if the seat has been taken or not
		if a.purchased[i].s.number == tick.s.number {
			if a.purchased[i].s.row == tick.s.row{
				return false,tick
			}
		}
	}
	return true, tick}



func (a Comedy) NewTicket (name string, seat Seat, show Show) Ticket { //return true when done giving new ticket.
	t2 := Ticket{}

	if seat.number == 0{	//if seat number = 0, this means that the seat is not assigned/ unavailable.
		return t2
	}

	t1 := Ticket {
		customerName: name,
		s: seat,
		show: show,
	}
	a.addPurchase(t1)	//after creating a new ticket, add to the purchase list
	return t1	//return the info of the ticket
}

func (a Tragedy) NewTicket (name string, seat Seat, show Show) (Ticket) { //return true when done giving new ticket.
	t2 := Ticket{}

	if seat.number == 0{
		return t2
	}

	t1 := Ticket {
		customerName: name,
		s: seat,
		show: show,
	}
	a.addPurchase(t1)	//after creating a new ticket, add to the purchase list
	return t1	//return the info of the ticket
}

func (a Comedy) NewSeat (seatNumber int32, rowNumber int32, category Category) Seat {		//return true if new seat is assigned
	//assign new seat
	s1 := Seat{	//create a new seat
		number: seatNumber,
		row: rowNumber,
		cat: category,
	}
	
	s2 := Seat{} //unable to create seat since it is taken
	for i := 0; i < len(a.purchased); i++{
		if a.purchased[i].s.number == seatNumber{
			if a.purchased[i].s.row == rowNumber{
				return s2 	//since the seat is not available therefore, return an empty seat.
			}
		}
	}

	return s1	//if seat is available, return seat varaible
}

func (a Tragedy) NewSeat (seatNumber int32, rowNumber int32, category Category) Seat {		//return true if new seat is assigned
	s1 := Seat{
		number: seatNumber,
		row: rowNumber,
		cat: category,
	}

	s2 := Seat{} //unable to create seat since it is taken
	for i := 0; i < len(a.purchased); i++{
		if a.purchased[i].s.number == seatNumber{
			if a.purchased[i].s.row == rowNumber{
				return s2 	//since the seat is not available therefore, return an empty seat.
			}
		}
	}
	
	return s1
}

func NewTheatre (seatNumber int32, shows []Show) Theatre{
	t1 := Theatre{
		seats: make([]Seat, seatNumber),	//create a slice with n number of seats
		shows: shows,
	}
	return t1	//create a theatre and return the theatre
}

func initializeComedy() Comedy{
	c1 := Comedy {
		Play: Play {
			name: "Tartuffe",
			purchased: make([]Ticket, 0),
			showsStart: time.Date(2020, time.March,3, 16,0,0,0,time.UTC),
			showEnd: time.Date(2020, time.March, 3, 17, 20, 0, 0 ,time.UTC),
		},
		laughs: 0.2,
		deaths: 0,
	}

	return c1
}

func initializeTragedy () Tragedy {
	t2 := Tragedy{
		Play: Play {
			name: "Macbeth",
			purchased: make([]Ticket, 0),
			showsStart: time.Date(2020, time.April,16, 9,30,0,0,time.UTC),
			showEnd: time.Date(2020, time.April, 16, 12, 30, 0, 0 ,time.UTC),
		},
		laughs: 0.0,
		deaths: 12,
	}
	return t2
}

func initializeSeat() Seat{
	n := Seat{
		number: 1,
		row: 1,
		cat:Category{
			name: "Standard",
			basePrice: 25.0,
		},
	}
	return n
}

func main() {
	c:= initializeComedy()
	t := initializeTragedy()
	n := initializeSeat()

	theatre := Theatre {}
	//create 25 seats in 5 rows evenly, from front to back.
	//   Screen
	//1 6  11 16 21
	//2 7  12 17 22
	//3 8  13 18 23
	//4 9  14 19 24
	//5 10 15 20 25
	var i int32
	for i = 0; i < 5; i++{
		n.row = i + 1
		for o := 0 ; o  < 5; o++ {
			switch o{
			case 0:
				//if o == 0 then its column 1
				n.number = i + 1
				n.cat.name = "Prime"
				n.cat.basePrice = 35.0
			case 4:
				//column 5
				n.number = 21+ i
				n.cat.name = "Special"
				n.cat.basePrice = 15.0
			case 2:
				n.number = i + 11
			case 3:
				n.number = i + 15
			case 1:
				n.number = i + 6
			}
			//after creating the seats, append it to the theatre to create a theatre
			theatre.seats = append(theatre.seats, n)
		}
	}
	theatre.shows = append(theatre.shows, c)
	theatre.shows = append(theatre.shows, t)


	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the name of the play: ")
	text, _ := reader.ReadString('\n')
	for {
		fmt.Print("Enter the seat number: ")   
		var input int32
		fmt.Scanf("%d", &input)
		seatNo := int32(input)
		
		//Play Tartuffe,
		var seatType string
		seatType = ""
		var rowSeat int32
		var price float32
		//which play the customer choose
		if text == c.Play.name{
			for i := 0 ;i < len(theatre.seats); i++{
				//find the type of the ticket
				if theatre.seats[i].number == seatNo{
					seatType = theatre.seats[i].cat.name
					rowSeat = int32(i)
					price = theatre.seats[i].cat.basePrice
				}
			}
			booleanCon := true
			for m:=0 ; m < len(c.purchased); m++{
				//check if the seat is taken or not
				if c.Play.purchased[m].s.number == seatNo{
					//the seat is already taken, give alternative option for the customer.
					newSeat := offerAlter(seatType, theatre, c)
					if newSeat.number != 0 {
						t1 := Ticket{
							s: newSeat,
						}
						add, ticket := theatre.shows[0].addPurchase(t1)
						if add == true {
							c.purchased = append(c.purchased, ticket)
							fmt.Println("The ticket has been added")
						}
					}else {
						fmt.Println("The play is full")
					}
					booleanCon = false
				}
			}
			if booleanCon == true{
				t1 := Ticket{
					s: Seat{
						number: seatNo,
						row: rowSeat,
						cat: Category{
							basePrice: price,
							name: seatType,
						},
					},
				}
				add, ticket := theatre.shows[0].addPurchase(t1)
				if add == true {
					c.purchased = append(c.purchased, ticket)
					fmt.Println("The ticket has been added")
				}
			}
		} else{
			for i := 0 ;i < len(theatre.seats); i++{
				//find the type of the ticket
				if theatre.seats[i].number == seatNo{
					seatType = theatre.seats[i].cat.name
					rowSeat = int32(i)
					price = theatre.seats[i].cat.basePrice
				}
			}
			booleanCon := true
			for m:=0 ; m < len(t.purchased); m++{
				//check if the seat is taken or not
				if t.purchased[m].s.number == seatNo{
					//the seat is already taken, give alternative option for the customer.
					newSeat := offerAlter2(seatType, theatre, t)
					if newSeat.number != 0 {
						t1 := Ticket{
							s: newSeat,
						}
						add, ticket := theatre.shows[0].addPurchase(t1)
						if add == true {
							t.purchased = append(t.purchased, ticket)
							fmt.Print("The ticket ")
							fmt.Print(ticket.s.number) 
							fmt.Println("has been added")
						}
					}else {
						fmt.Println("The play is full")
					}
					booleanCon = false
				}
			}
			if booleanCon == true{
				t1 := Ticket{
					s: Seat{
						number: seatNo,
						row: rowSeat,
						cat: Category{
							basePrice: price,
							name: seatType,
						},
					},
				}
				add, ticket := theatre.shows[0].addPurchase(t1)
				if add == true {
					t.purchased = append(t.purchased, ticket)
					fmt.Print("The ticket ")
					fmt.Print(ticket.s.number) 
					fmt.Println("has been added")
				}
			}
		}
	}
}

func offerAlter (seatType string, theatre Theatre, c Comedy) (Seat) {
	var listOfFreeSeats []Seat
	var catCheckList []bool
	catCheckList = make([]bool, 3)
	loop := true

	for loop {
		List1 := 0
		List2 := 0
		List3 := 0
		if seatType == "Prime"{ //if seatType is Prime, therefore it is row 1
			
			for i := 0; i < len(theatre.seats)/5; i ++{
				listOfFreeSeats = append (listOfFreeSeats, theatre.seats[i])
			}
			boolean := true
			var alterSeat Seat
			for w := 0 ; w < len(listOfFreeSeats) ; w++{
				for q:=0;q< len(c.purchased);q++{
					if listOfFreeSeats[w].number == c.purchased[q].s.number {
						boolean = false
						List1++
					}
				}
				if boolean == true{
					//then we give this alternative option
					alterSeat = listOfFreeSeats[w]
					return alterSeat
				} else{
					if List1 == 5{
						catCheckList[0] = true
					}
					if catCheckList[0] == true{
						if catCheckList[1] == true{
							seatType = "Special"
						}else{
							seatType = "Standard"
						}
					} 
				}
			}
		} else if seatType == "Standard"{
			
			//if standard will be from 2 to 4
			for i := 0; i < len(theatre.seats); i ++{
				if theatre.seats[i].number >= 6 && theatre.seats[i].number <= 20{
					listOfFreeSeats = append (listOfFreeSeats, theatre.seats[i])
				}
			}
			boolean := true
			var alterSeat Seat
			for w := 0 ; w < len(listOfFreeSeats) ; w++{
				for q:=0;q< len(c.purchased);q++{
					if listOfFreeSeats[w].number == c.purchased[q].s.number {
						boolean = false
						List2++
						
					}
				}
				if boolean == true{
					//then we give this alternative option
					alterSeat = listOfFreeSeats[w]
					return alterSeat
				}else{
					if List2 == 15{
						catCheckList[1] = true
					}
					if catCheckList[1] == true{
						if catCheckList[0] == true{
							seatType = "Special"
						}else{
							seatType = "Prime"
						}
					} 
				}
			}
		} else{	//type special row = 5
			catCheckList[3] = true
			for i := 0; i < len(theatre.seats); i ++{
				if theatre.seats[i].number >= 21{
					listOfFreeSeats = append (listOfFreeSeats, theatre.seats[i])
				}
			}
			boolean := true
			var alterSeat Seat
			for w := 0 ; w < len(listOfFreeSeats) ; w++{
				for q:=0;q< len(c.purchased);q++{
					if listOfFreeSeats[w].number == c.purchased[q].s.number {
						boolean = false
						List3++
					}
				}
				if boolean == true{
					//then we give this alternative option
					alterSeat = listOfFreeSeats[w]
					return alterSeat
				} else{
					//if boolean = false, this means all of the seats in that cat has been purchased.
					if List3 == 5{
						catCheckList[2] = true
					}
					if catCheckList[2] == true{
						if catCheckList[1] == true{
							seatType = "Standard"
						}else{
							seatType = "Prime"
						}
					} 
				}
			}
		}
	}
	var emptySeat Seat
	return emptySeat
}

func offerAlter2 (seatType string, theatre Theatre, c Tragedy) (Seat) {
	var listOfFreeSeats []Seat
	var catCheckList []bool
	catCheckList = make([]bool, 3)
	loop := true

	for loop {
		List1 := 0
		List2 := 0
		List3 := 0
		if seatType == "Prime"{ //if seatType is Prime, therefore it is row 1
			
			for i := 0; i < len(theatre.seats)/5; i ++{
				listOfFreeSeats = append (listOfFreeSeats, theatre.seats[i])
			}
			boolean := true
			var alterSeat Seat
			for w := 0 ; w < len(listOfFreeSeats) ; w++{
				for q:=0;q< len(c.purchased);q++{
					if listOfFreeSeats[w].number == c.purchased[q].s.number {
						boolean = false
						List1++
					}
				}
				if boolean == true{
					//then we give this alternative option
					alterSeat = listOfFreeSeats[w]
					return alterSeat
				} else{
					if List1 == 5{
						catCheckList[0] = true
					}
					if catCheckList[0] == true{
						if catCheckList[1] == true{
							seatType = "Special"
						}else{
							seatType = "Standard"
						}
					} 
				}
			}
		} else if seatType == "Standard"{
			
			//if standard will be from 2 to 4
			for i := 0; i < len(theatre.seats); i ++{
				if theatre.seats[i].number >= 6 && theatre.seats[i].number <= 20{
					listOfFreeSeats = append (listOfFreeSeats, theatre.seats[i])
				}
			}
			boolean := true
			var alterSeat Seat
			for w := 0 ; w < len(listOfFreeSeats) ; w++{
				for q:=0;q< len(c.purchased);q++{
					if listOfFreeSeats[w].number == c.purchased[q].s.number {
						boolean = false
						List2++
						break
					}
				}
				if boolean == true{
					//then we give this alternative option
					alterSeat = listOfFreeSeats[w]
					return alterSeat
				}else{
					if List2 == 15{
						catCheckList[1] = true
					}
					if catCheckList[1] == true{
						if catCheckList[0] == true{
							seatType = "Special"
						}else{
							seatType = "Prime"
						}
					} 
				}
			}
		} else{	//type special row = 5
			catCheckList[3] = true
			for i := 0; i < len(theatre.seats); i ++{
				if theatre.seats[i].number >= 21{
					listOfFreeSeats = append (listOfFreeSeats, theatre.seats[i])
				}
			}
			boolean := true
			var alterSeat Seat
			for w := 0 ; w < len(listOfFreeSeats) ; w++{
				for q:=0;q< len(c.purchased);q++{
					if listOfFreeSeats[w].number == c.purchased[q].s.number {
						boolean = false
						List3++
					}
				}
				if boolean == true{
					//then we give this alternative option
					alterSeat = listOfFreeSeats[w]
					return alterSeat
				} else{
					//if boolean = false, this means all of the seats in that cat has been purchased.
					if List3 == 5{
						catCheckList[2] = true
					}
					if catCheckList[2] == true{
						if catCheckList[1] == true{
							seatType = "Standard"
						}else{
							seatType = "Prime"
						}
					} 
				}
			}
		}
	}
	var emptySeat Seat
	return emptySeat
}
