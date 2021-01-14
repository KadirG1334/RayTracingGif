package main

import "testing"
import "math"
import "fmt"


var DivideF int
var MultiplyF int
var SubsF int
var AddF int
var LengthSquF int
var LengthF int
var not int
var nof int
var Fnumb int
func Bool2int(b bool) int {

	if b {
		not = not + 1
		return not
	} else {
		nof = nof +1
		return nof
	}
	
}
func float64Eq(x, y float64) bool { return math.Abs(x-y) < 1e-14 }
// ApproxEqual reports whether v and ov are equal within a small epsilon,mutlak değer

func TestDot(t *testing.T){
tests := []struct {
		v1, v2 Vec3
		formed   float64
	}{
		{Vec3{1, 0, 0}, Vec3{1, 0, 0}, 1},
		{Vec3{1, 0, 0}, Vec3{0, 1, 0}, 0},
		{Vec3{1, 0, 0}, Vec3{0, 1, 1}, 0},
		{Vec3{1, 1, 1}, Vec3{-1, -1, -1}, -3},
		{Vec3{1, 2, 2}, Vec3{-0.3, 0.4, -1.2}, -14.9},
	}
	for _, test := range tests {
		v1 := Vec3{test.v1.X, test.v1.Y, test.v1.Z}
		v2 := Vec3{test.v2.X, test.v2.Y, test.v2.Z}
		Bool2int(float64Eq(v1.Dot(v2), test.formed))

		if (float64Eq(v1.Dot(v2), test.formed))==false{
			t.Errorf("%v · %v = %v, formed %v", v1, v2, v1.Dot(v2), test.formed)
		}



}	
	
	fmt.Printf("%d test applied  %v SUCCEED " ,  len(tests), not)
	fmt.Printf("and %d FAILED", nof)
   	fmt.Println("\n=================================== ")
  
}

func TestCross(t *testing.T){
	tests := []struct {
		v1, v2, formed Vec3
	}{
		{Vec3{1, 0, 0}, Vec3{1, 0, 0}, Vec3{0, 0, 0}},
		{Vec3{1, 0, 0}, Vec3{0, 1, 0}, Vec3{0, 0, 1}},
		{Vec3{0, 1, 0}, Vec3{1, 0, 0}, Vec3{0, 0, -1}},
		{Vec3{1, 2, 3}, Vec3{-4, 5, -6}, Vec3{-27, -6, 12}},
	}
	
	for _, test := range tests {
		 Mycross := test.v1.Cross(test.v2)
			
	
	 if !(Mycross==(test.formed)) {
			t.Errorf("%v X %v = %v, formed %v ", test.v1, test.v2, Mycross, test.formed)
			Fnumb +=1
			
    }
	
   	}
   	       
   	fmt.Printf("%d test applied  %d SUCCEED ", len(tests), len(tests)-Fnumb)
   	fmt.Println("\n=================================== ")
}
	func TestLength(t *testing.T){
		tests:=[]struct {
		v1 Vec3 
		formed float64
	}{
		{Vec3{3, 0, 4}, 5},
		{Vec3{1, 1, 1}, 25},
		{Vec3{8, 15, 0}, 17},
		}
	
	for _, test := range tests {
		Mylength := test.v1.Length()
		if (Mylength != test.formed){
			t.Errorf("%v length is = %v, formed %v", test.v1, Mylength, test.formed)
			LengthF += LengthF +1  
		}
	}
	fmt.Printf("%d test applied  %d SUCCEED and %d FAILED ", len(tests), len(tests)-LengthF,LengthF)
   	fmt.Println("\n=================================== ")
}
	func TestSquaredLength(t *testing.T){
		tests:=[]struct {
		v1 Vec3 
		formed float64
		}{
		{Vec3{3, 0, 4}, 25},
		{Vec3{5, 7, 8}, 138},
		{Vec3{6, 3, 0}, 45},
		}

	for _, test := range tests {
		MySqu := test.v1.SquaredLength()
		if (MySqu != test.formed){
			t.Errorf("%v squared length is = %v, formed %v", test.v1, MySqu, test.formed)
			LengthSquF = LengthSquF +1  
		}
}
    fmt.Printf("%d test applied  %d SUCCEED and %d FAILED ", len(tests), len(tests)-LengthSquF,LengthSquF)
   	fmt.Println("\n=================================== ")
}
	func TestAdd(t *testing.T){
		tests :=[]struct{
			v1, v2, formed Vec3
	}{
	
		{Vec3{0, 0, 0}, Vec3{0, 0, 0}, Vec3{0, 0, 0}},
		{Vec3{1, 2, 3}, Vec3{4, 5, 6}, Vec3{5, 7, 9}},
		{Vec3{1, -3, 5}, Vec3{1, -6, -6}, Vec3{2, -9, -1}},
	}
	for _, test := range tests {
		 MyAdd := test.v1.Add(test.v2)
	 if !(MyAdd==(test.formed)) {
			t.Errorf("%v + %v = %v, formed %v ", test.v1, test.v2, MyAdd, test.formed)
			AddF += 1
			
    }

}		
		fmt.Printf("%d test applied  %d SUCCEED and %d FAILED", len(tests), len(tests)-AddF,AddF)
     	fmt.Println("\n=================================== ")
}
func TestSubstract(t *testing.T){
		tests :=[]struct{
			v1, v2, formed Vec3
	}{
		{Vec3{10, 7, 12}, Vec3{4, 5, 8}, Vec3{6, 2, 4}},
		{Vec3{0, 0, 0}, Vec3{0, 0, 0}, Vec3{0, 0, 0}},
		{Vec3{11, 23, 34}, Vec3{4, 5, 6}, Vec3{7, 18, 28}},
		{Vec3{1, -3, 5}, Vec3{1, -6, -6}, Vec3{0, 3, 11}},
	}
	for _, test := range tests {
		 MySubs := test.v1.Subtract(test.v2)
	 if !(MySubs==(test.formed)) {
			t.Errorf("%v - %v = %v, formed %v ", test.v1, test.v2,  MySubs, test.formed)
			SubsF += 1
			
    }

}		
		fmt.Printf("%d test applied  %d SUCCEED and %d FAILED", len(tests), len(tests)-SubsF,SubsF)
     	fmt.Println("\n=================================== ")
}
	func TestMultiply(t *testing.T){
		tests :=[]struct{
			v1, v2, formed Vec3
		}{
		{Vec3{10, 7, 12}, Vec3{4, 5, 8}, Vec3{40, 35, 96}},
		{Vec3{0, 0, 0}, Vec3{0, 0, 0}, Vec3{0, 0, 0}},
		{Vec3{11, 9, 13}, Vec3{4, 5, 6}, Vec3{7, 18, 28}},	
		}
		for _, test := range tests {
		 MyMultiply:= test.v1.Multiply(test.v2)

	
		if !(MyMultiply==(test.formed)) {
			t.Errorf("%v * %v = %v, formed %v ", test.v1, test.v2,  MyMultiply, test.formed)
			MultiplyF += 1
	}	}
	fmt.Printf("%d test applied  %d SUCCEED and %d FAILED", len(tests), len(tests)-MultiplyF,MultiplyF)
     fmt.Println("\n=================================== ")
 }
 func TestDivide(t *testing.T){
		tests :=[]struct{
			v1, v2, formed Vec3
		}{
		{Vec3{10, 15, 12}, Vec3{4, 5, 10}, Vec3{2.5, 3, 1.2}},
		{Vec3{13, 13, 13}, Vec3{13, 13, 13}, Vec3{1, 1, 1}},
		{Vec3{24, 9, 50}, Vec3{4, 3, 10}, Vec3{6, 3, 5}},	
		}
		for _, test := range tests {
		 MyDivide:= test.v1.Divide(test.v2)

	
		if !(MyDivide==(test.formed)) {
			t.Errorf("%v / %v = %v, formed %v ", test.v1, test.v2,  MyDivide, test.formed)
			DivideF += 1
	}	}
	fmt.Printf("%d test applied  %d SUCCEED and %d FAILED", len(tests), len(tests)-DivideF,DivideF)
     fmt.Println("\n=================================== ")
 }
  func TestAddScalar(t *testing.T){
		var ScalF int
		tests :=[]struct{
			v1 Vec3
			m float64 
			formed Vec3
		}{
		{Vec3{14, 8, 9.7}, 1.3, Vec3{15.3, 9.3, 11}},
		{Vec3{5, 4, 9}, 5, Vec3{10, 9, 14}},
		}
		for _, test := range tests {
		 AddScal:= test.v1.AddScalar(test.m)

	
		if !(AddScal==(test.formed)) {
			t.Errorf("%v + %v = %v, formed %v ", test.v1, test.m,  AddScal, test.formed)
			ScalF += 1
	}	}
	fmt.Printf("%d test applied  %d SUCCEED and %d FAILED", len(tests), len(tests)-ScalF,ScalF)
     fmt.Println("\n=================================== ")
 }
 func TestSubtractScalar(t *testing.T){
		var SubScalF int
		tests :=[]struct{
			v1 Vec3
			m float64 
			formed Vec3
		}{
		{Vec3{14, 8, 9.7}, 1.5, Vec3{12.5, 6.5, 8.2}},
		{Vec3{8, 7, 9}, 5, Vec3{3, 2, 4}},
		}
		for _, test := range tests {
		 SubScal:= test.v1.SubtractScalar(test.m)

	
		if !(SubScal==(test.formed)) {
			t.Errorf("%v - %v = %v, formed %v ", test.v1, test.m,  SubScal, test.formed)
			SubScalF += 1
	}	}
	fmt.Printf("%d test applied  %d SUCCEED and %d FAILED", len(tests), len(tests)-SubScalF,SubScalF)
     fmt.Println("\n=================================== ")
 }
 func TestMultiplyScalar(t *testing.T){
		var MultScalF int
		tests :=[]struct{
			v1 Vec3
			m float64 
			formed Vec3
		}{
		{Vec3{2, 4, 6}, 1.5, Vec3{3, 6, 9}},
		{Vec3{8, 7, 9}, 2, Vec3{16, 14, 18}},
		}
		for _, test := range tests {
		 MulScal:= test.v1.MultiplyScalar(test.m)

	
		if !(MulScal==(test.formed)) {
			t.Errorf("%v * %v = %v, formed %v ", test.v1, test.m,   MulScal, test.formed)
			MultScalF += 1
	}	}
	fmt.Printf("%d test applied  %d SUCCEED and %d FAILED", len(tests), len(tests)-MultScalF,MultScalF)
     fmt.Println("\n=================================== ")
 }
 func TestDivideScalar(t *testing.T){
		var DivScalF int
		tests :=[]struct{
			v1 Vec3
			m float64 
			formed Vec3
		}{
		{Vec3{2, 4, 6}, 0.5, Vec3{4, 8, 12}},
		{Vec3{8, 7, 9}, 2, Vec3{4, 3.5, 4.5}},
		}
		for _, test := range tests {
		 DivScal:= test.v1.DivideScalar(test.m)

	
		if !(DivScal==(test.formed)) {
			t.Errorf("%v / %v = %v, formed %v ", test.v1, test.m,   DivScal, test.formed)
			DivScalF += 1
	}	}
	fmt.Printf("%d test applied  %d SUCCEED and %d FAILED", len(tests), len(tests)-DivScalF,DivScalF)
     fmt.Println("\n=================================== ")
 }
 func TestNormalize(t *testing.T){
		var NormF int
		tests :=[]struct{
			v1 Vec3
			L float64 
			formed Vec3
		
		}{
		{Vec3{3, 4, 0}, 5, Vec3{0.6, 0.8, 0}},
		{Vec3{8, 0, 15}, 17, Vec3{4, 3.5, 4.5}},
		}
		for _, test := range tests {
		 Norm:= test.v1.Normalize()

	
		if !(Norm==(test.formed)) {
			t.Errorf("%v / %v = %v, formed %v ", test.v1, test.L,  Norm, test.formed)
			NormF += 1
	}	}
	fmt.Printf("%d test applied  %d SUCCEED and %d FAILED", len(tests), len(tests)-NormF,NormF)
     fmt.Println("\n=================================== ")
 }
