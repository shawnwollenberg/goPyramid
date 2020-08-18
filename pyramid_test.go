package pyramid

import (
	"testing"
)

func TestStepsGood(t *testing.T) {
	want := `   #   
  ###  
 ##### 
#######`
	if got := pyramid(4); got != want {
		t.Errorf("pyramid(4) = %q, want %q", got, want)
	}
}

func TestStepsRecursion(t *testing.T) {
	want := `   #   
  ###  
 ##### 
#######`
	if got := pyramidRecursion(4); got != want {
		t.Errorf("pyramidRecursion(4) = %q, want %q", got, want)
	}
}
