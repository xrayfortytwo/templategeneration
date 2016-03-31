grammar Calc ;

@header {
    // calc is for calculating
}

calc : expre* EOF ;
expre : num0 op num1 (op num2)? '=' num3 ;
num0 : NUMBER ;
num1 : NUMBER ;
num2 : NUMBER ;
ob : OP ;
OP : (*|+|-|) ; 
NUMBER : [0-9]+ ;

