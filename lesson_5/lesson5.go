package main

import "fmt"

func main() {
	var x int = 10
	var p *int
	p = &x

	fmt.Println("x: ", x)
	fmt.Println("x address: ", &x)
	fmt.Println("p: ", p)
	fmt.Println("p points to: ", *p)

	//1
	fmt.Println("Before: ", x)
	updateValue(p)
	fmt.Println("After: ", x)

	//2
	y, z := 3, 5
	swap(&y, &z) 


	//Middle
	//8
	var math, mathNumb, physics, physicsNum, chem, chemNum int
	mathP, physicsP, chemP := &math, &physics, &chem

	addMark(mathP, &mathNumb, 5)
	addMark(physicsP, &physicsNum, 5)
	addMark(chemP, &chemNum, 5)

	
}

//1
func updateValue(p *int) {
	*p =  *p + 10
}

//2
func swap(a, b *int) {
	p := *a
	*a = *b
	*b = p
}

//4
func increment(p *int) {
	*p =  *p + 1
}

//5
func isPosititve(p *int) bool {
	if *p > 0 {
		return true
	} else {
		return false
	}
}

//6
func changeString(p *string) {
	*p = *p + " Go"
}

//7
func double(p *int) {
	*p = *p * 2	
}

//8
func isEven(p *int) bool{
	if *p % 2 == 0 {
		return true
	} else {
		return false
	}	
}

//9
func checkNil(p *int) bool{
	if p == nil {
		return true
	} else {
		return false
	}
}

//10


//middle level 

//1
func storeData(nameP *string, ageP *int, name string, age int) {
	*nameP = name
	*ageP = age
	fmt.Println("Name is ", *nameP, " Age is ", *ageP)
}

//3
func addScore(studentScoreP *int, studentScore int, studentsTotalP, studentsNumberP *int) {
	*studentScoreP = studentScore
	*studentsTotalP += studentScore
	*studentsNumberP ++
}

func scoreSummary(studentsTotalP, studentsNumberP *int) {
	fmt.Println("Average Score is ", *studentsTotalP / *studentsNumberP)
}


//4
func vote(votesP *int) {
	*votesP ++
}

func winner(Candidate1Votes, Candidate2Votes *int) {
	if *Candidate1Votes > *Candidate2Votes {
		fmt.Println("Winner is Candidate 1")
	} else if *Candidate1Votes < *Candidate2Votes {
		fmt.Println("Winner is Candidate 2")
	} else {
		fmt.Println("It's a tie!")
	}
}

//6
func addBalance(totalSum *int, addedSum int) {
	*totalSum += addedSum
}

func checkFalance(totalSum *int) {
	fmt.Println("Current balance is ", *totalSum)
}

//7
//func addExpense(expenseBook *int, )

//8
func addMark(subjectMarksP *int, totalMarksP *int, mark int) {
	*subjectMarksP += mark
	*totalMarksP +=1
}

func calculateMean(subjectMarksP *int, totalMarksP *int) int {
	return *subjectMarksP / *totalMarksP
}

//9