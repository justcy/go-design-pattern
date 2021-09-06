package main

func main() {
	c := &CustomerCol{}
	c.Add(NewEnterpriseCustomer("A company"))
	c.Add(NewEnterpriseCustomer("B company"))
	c.Add(NewIndividualCustomer("bob"))
	c.Accept(&ServiceRequestVisitor{})

	c1 := &CustomerCol{}
	c1.Add(NewEnterpriseCustomer("A company"))
	c1.Add(NewIndividualCustomer("bob"))
	c1.Add(NewEnterpriseCustomer("B company"))
	c1.Accept(&AnalysisVisitor{})
}
